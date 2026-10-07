// SPDX-License-Identifier: Apache-2.0

// Package mcp serves the process operations over the Model Context Protocol, on standard
// input and output.
//
// It is the server section 13 describes, and the description is most of the design: "not
// something anybody stands up: it is a subcommand of the binary that is on the machine
// anyway, started by the client over standard input and output, with no host, no port and
// no deployment". So there is no listener here, no transport choice and no configuration —
// one goroutine reading lines from a reader and writing lines to a writer.
//
// Every operation is a front for the runner method the command of the same name calls, and
// that is the property to keep: "Every operation the server exposes has a command that does
// the same thing, because the server is a thin shell around the same code." Nothing in this
// package writes a file, hashes anything or decides a verdict. Where an operation would need
// something the runner does not expose, the fix is a method on the runner that the command
// can reach too, never a second path to the same artifact: two writers of one artifact is
// how a trail driven through the operations stops being comparable with one driven by
// commands.
//
// The framing is written against encoding/json rather than taken from an SDK, because this
// repository has one dependency and it is a YAML parser. What that costs is the scope below:
// the revision this implements is stated rather than negotiated across eras, and anything
// outside it is answered with an error rather than guessed at.
package mcp

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/runner"
)

// Revision is the protocol revision this server implements, and it is one revision and not
// a range.
//
// 2025-06-18 is the handshake based revision, what the specification's compatibility page
// calls legacy: a client opens with `initialize`, the version is agreed once and the session
// is scoped to this process. It is chosen because it is what the clients this plugin is
// written for speak — Claude Code implements 2025-06-18 (verified 2026-07-01), and Codex
// gained the 2026-07-28 revision in v0.147.0 as an opt-in a project turns on in
// `.codex/config.toml` rather than as its default.
//
// A modern client probing a stdio server sends `server/discover` first and falls back to
// `initialize` on any error that is not a recognised modern error, which is what the
// method-not-found answer below is: a dual-era client therefore reaches the handshake by
// the route the specification lays out for it, and a client that speaks only 2026-07-28 is
// told in those words rather than being served something that looks like an answer. When a
// client's default moves, the work is to add the modern era beside this one, not to change
// what this constant says.
const Revision = "2025-06-18"

// Name and Version identify the server to the client. The name is the plugin's own, so that
// a tool appears in a client as this plugin's and not as a second thing to configure.
const Name = "xeno"

// Server is one conversation: a runner to act through, a reader the client writes to, a
// writer the client reads, and a log nothing but a person reads.
//
// Log is standard error and never standard output. The stdio transport is one JSON message
// per line on the one stream, so a line printed for a person is a parse error for the client.
type Server struct {
	R        *runner.Runner
	In       io.Reader
	Out, Log io.Writer
}

// Serve reads requests until the client closes its end, which is how a stdio server stops:
// the client owns the process and ends it by closing the pipe.
//
// A line that is not a request gets a parse error and the loop carries on. The alternative,
// exiting, would take the session's other tools down with it over one malformed line.
func (s *Server) Serve() error {
	in := bufio.NewScanner(s.In)
	// A section of prose is one field of one request, so the default 64 KiB line limit is
	// reachable by ordinary use rather than by abuse.
	in.Buffer(make([]byte, 0, 64*1024), 8*1024*1024)
	out := bufio.NewWriter(s.Out)
	defer out.Flush()
	for in.Scan() {
		line := in.Bytes()
		if len(trimSpace(line)) == 0 {
			continue
		}
		var req request
		if err := json.Unmarshal(line, &req); err != nil {
			if werr := s.write(out, errorFor(nil, codeParse, "the line is not JSON")); werr != nil {
				return werr
			}
			continue
		}
		res, answer := s.handle(&req)
		if !answer {
			continue // a notification: the protocol forbids a response
		}
		if err := s.write(out, res); err != nil {
			return err
		}
	}
	return in.Err()
}

// write emits one message and flushes it. Flushing per message is the point rather than a
// cost: the client is waiting on this line before it sends another, so a buffered response
// is a hung session.
func (s *Server) write(out *bufio.Writer, v any) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	if _, err := out.Write(append(b, '\n')); err != nil {
		return err
	}
	return out.Flush()
}

// handle answers one request. The second return says whether there is anything to send: a
// notification carries no id and is answered with silence, which the protocol requires.
func (s *Server) handle(req *request) (any, bool) {
	if req.ID == nil {
		return nil, false
	}
	switch req.Method {
	case "initialize":
		return s.initialize(req), true
	case "ping":
		// The base protocol's liveness check. An empty result, because that is what it is
		// for: a client learns the process is still there.
		return resultFor(req.ID, struct{}{}), true
	case "tools/list":
		return resultFor(req.ID, listing()), true
	case "tools/call":
		return s.call(req), true
	default:
		// Also the answer to `server/discover`, deliberately and not by omission. A modern
		// client probes with it and reads anything that is not a recognised modern error as
		// a legacy server, so method-not-found is what sends it to the handshake this
		// server does implement.
		return errorFor(req.ID, codeNoMethod, req.Method+" is not a method of this server; "+
			"it implements MCP "+Revision+", where the conversation opens with initialize"), true
	}
}

// initialize is the handshake. The reply names the revision this server implements, whatever
// the client asked for: the revision is not negotiable here, and a server that echoed a
// version it does not implement would be claiming to speak it.
func (s *Server) initialize(req *request) any {
	var p struct {
		ProtocolVersion string `json:"protocolVersion"`
	}
	if len(req.Params) > 0 {
		_ = json.Unmarshal(req.Params, &p)
	}
	if p.ProtocolVersion != "" && p.ProtocolVersion != Revision && s.Log != nil {
		fmt.Fprintf(s.Log, "the client asked for MCP %s; this server implements %s\n",
			p.ProtocolVersion, Revision)
	}
	return resultFor(req.ID, initializeResult{
		ProtocolVersion: Revision,
		// Tools and nothing else. Resources and prompts are capabilities this server does
		// not have, and section 13 says why a prompt is not among them: the phases are
		// skills until both clients can take a prompt as a command.
		Capabilities: capabilities{Tools: &struct{}{}},
		ServerInfo:   serverInfo{Name: Name, Version: model.RunnerVersion},
	})
}

// call runs one tool. A tool that refuses or fails answers with isError inside a result and
// not with a JSON-RPC error: the protocol reserves those for the call not happening, and a
// refusal the model has to read is the call happening and saying no.
func (s *Server) call(req *request) any {
	var p struct {
		Name      string          `json:"name"`
		Arguments json.RawMessage `json:"arguments"`
	}
	if len(req.Params) > 0 {
		if err := json.Unmarshal(req.Params, &p); err != nil {
			return errorFor(req.ID, codeBadParams, "the parameters of tools/call are not an object")
		}
	}
	t, ok := tools[p.Name]
	if !ok {
		return errorFor(req.ID, codeBadParams, "no tool "+p.Name)
	}
	text, err := t.call(s.R, p.Arguments)
	if err != nil {
		return resultFor(req.ID, callResult{Content: textContent(err.Error()), IsError: true})
	}
	return resultFor(req.ID, callResult{Content: textContent(text)})
}

// listing is tools/list, in a fixed order so that two runs of one client send the same bytes.
func listing() toolsResult {
	out := toolsResult{Tools: make([]definition, 0, len(order))}
	for _, name := range order {
		t := tools[name]
		out.Tools = append(out.Tools, definition{
			Name: name, Description: t.description, InputSchema: json.RawMessage(t.schema),
		})
	}
	return out
}

// trimSpace is strings.TrimSpace over bytes, for the one use above: a transport that writes
// a bare newline between messages is not an error to report.
func trimSpace(b []byte) []byte {
	i, j := 0, len(b)
	for i < j && (b[i] == ' ' || b[i] == '\t' || b[i] == '\r' || b[i] == '\n') {
		i++
	}
	for j > i && (b[j-1] == ' ' || b[j-1] == '\t' || b[j-1] == '\r' || b[j-1] == '\n') {
		j--
	}
	return b[i:j]
}
