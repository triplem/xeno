// SPDX-License-Identifier: Apache-2.0

package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/triplem/xeno/internal/mcp"
	"github.com/triplem/xeno/internal/runner"
)

// cmdMCP serves the six process operations over standard input and output, for as long as
// the client keeps the pipe open.
//
// It is answered before the dispatch table, as `version` is, and for a related reason:
// `mcp` is one word where every entry of that table is two, and nothing about it takes an
// intent, a phase or a finding. The flags are the one it needs and no more — a server
// started by a client has nobody to read a suggestion or a verdict line.
//
// It is not in the `usage` constant, which is the one thing here that is awkward and is so
// deliberately: docs/commands.md carries that constant verbatim and a test holds the two
// identical, and the page is not the agent's to write. The line it needs is proposed on
// #285 for a person to add to both at once.
//
// Standard output is the transport and carries JSON and nothing else, so everything meant
// for a person goes to standard error. That is the stdio transport's rule rather than a
// preference: one stray line on standard output is a parse error in the client.
func cmdMCP(args []string, out, errw io.Writer) int {
	fs := flag.NewFlagSet("mcp", flag.ContinueOnError)
	fs.SetOutput(errw)
	root := fs.String("root", ".", "repository root")
	if err := fs.Parse(args); err != nil {
		return 2
	}
	s := &mcp.Server{R: runner.New(*root), In: os.Stdin, Out: out, Log: errw}
	if err := s.Serve(); err != nil {
		fmt.Fprintln(errw, "error:", err)
		return 2
	}
	return 0
}
