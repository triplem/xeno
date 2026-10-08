// SPDX-License-Identifier: Apache-2.0

package mcp

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"

	"go.yaml.in/yaml/v3"

	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/runner"
)

// serve runs the lines through one server and returns what it wrote, line by line. The
// server is driven over buffers rather than over a pipe: the transport is a reader and a
// writer, so a test needs no process.
func serve(t *testing.T, r *runner.Runner, lines ...string) []string {
	t.Helper()
	var out, log bytes.Buffer
	s := &Server{R: r, In: strings.NewReader(strings.Join(lines, "\n") + "\n"), Out: &out, Log: &log}
	if err := s.Serve(); err != nil {
		t.Fatalf("the server stopped on an error: %v", err)
	}
	var got []string
	for _, l := range strings.Split(strings.TrimSpace(out.String()), "\n") {
		if l != "" {
			got = append(got, l)
		}
	}
	return got
}

// call sends one tools/call and returns the text of the single content block, and whether
// the tool reported an error.
func call(t *testing.T, r *runner.Runner, name string, args map[string]any) (string, bool) {
	t.Helper()
	a, err := json.Marshal(args)
	if err != nil {
		t.Fatal(err)
	}
	req := `{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"name":"` + name +
		`","arguments":` + string(a) + `}}`
	lines := serve(t, r, req)
	if len(lines) != 1 {
		t.Fatalf("one call, %d responses: %v", len(lines), lines)
	}
	var res struct {
		Result struct {
			Content []content `json:"content"`
			IsError bool      `json:"isError"`
		} `json:"result"`
		Error *rpcError `json:"error"`
	}
	if err := json.Unmarshal([]byte(lines[0]), &res); err != nil {
		t.Fatalf("the response is not a response: %v: %s", err, lines[0])
	}
	if res.Error != nil {
		t.Fatalf("%s failed at the protocol level, which is for a call that did not happen: %v",
			name, res.Error.Message)
	}
	if len(res.Result.Content) != 1 {
		t.Fatalf("%s returned %d content blocks", name, len(res.Result.Content))
	}
	return res.Result.Content[0].Text, res.Result.IsError
}

// The budget section 13 fixes, and the whole of what this test is for: six operations, named
// as that section names them. A seventh needs an argument of the kind the symbol index has,
// and the argument is made on an issue before it is made here.
func TestTheSurfaceIsSix(t *testing.T) {
	if len(tools) != 6 || len(order) != 6 {
		t.Fatalf("%d tools and %d in the listing order; section 13 fixes the surface at six",
			len(tools), len(order))
	}
	want := []string{"read_intent", "record_assumption", "write_artifact",
		"fetch_template", "run_gates", "query_symbol_index"}
	for i, name := range want {
		if order[i] != name {
			t.Errorf("tool %d is %q, and section 13's %d is %q", i, order[i], i, name)
		}
		if _, ok := tools[name]; !ok {
			t.Errorf("%s is in the order and not in the table", name)
		}
	}
}

// A definition is sent with every request of every phase whether it is called or not, so the
// thing to hold is its size. One sentence and a schema that parses.
func TestEveryDefinitionIsOneSentenceAndASchema(t *testing.T) {
	for _, d := range listing().Tools {
		if d.Description == "" || len(d.Description) > 110 {
			t.Errorf("%s carries %d characters of description: %q",
				d.Name, len(d.Description), d.Description)
		}
		if strings.Count(d.Description, ". ") > 0 {
			t.Errorf("%s carries more than one sentence: %q", d.Name, d.Description)
		}
		var parsed struct {
			Type       string         `json:"type"`
			Properties map[string]any `json:"properties"`
			Required   []string       `json:"required"`
		}
		if err := json.Unmarshal(d.InputSchema, &parsed); err != nil {
			t.Errorf("%s has no schema a client can read: %v", d.Name, err)
			continue
		}
		if parsed.Type != "object" || len(parsed.Properties) == 0 || len(parsed.Required) == 0 {
			t.Errorf("%s has a schema that asks for nothing: %s", d.Name, d.InputSchema)
		}
	}
}

// The handshake, and the one sentence this server makes about itself: which revision it
// implements. It answers with its own revision whatever was asked for, because answering
// with the client's would be claiming to speak it.
func TestTheHandshakeNamesTheRevisionTheServerImplements(t *testing.T) {
	lines := serve(t, runner.New(t.TempDir()),
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2026-07-28"}}`)
	if len(lines) != 1 {
		t.Fatalf("one initialize, %d responses", len(lines))
	}
	var res struct {
		Result initializeResult `json:"result"`
	}
	if err := json.Unmarshal([]byte(lines[0]), &res); err != nil {
		t.Fatal(err)
	}
	if res.Result.ProtocolVersion != Revision {
		t.Errorf("the server answered %q and implements %q",
			res.Result.ProtocolVersion, Revision)
	}
	if res.Result.Capabilities.Tools == nil {
		t.Error("the server declares no tools capability, and tools are all it has")
	}
	if res.Result.ServerInfo.Name != Name || res.Result.ServerInfo.Version != model.RunnerVersion {
		t.Errorf("the server identifies itself as %+v", res.Result.ServerInfo)
	}
}

// A modern client probes a stdio server with server/discover and falls back to the handshake
// on any error that is not a recognised modern one. So the answer to it is method-not-found,
// which is the route the specification lays out, and the message says what is implemented
// because a legacy-only client has nothing but that message to show a person.
func TestServerDiscoverSendsAModernClientToTheHandshake(t *testing.T) {
	lines := serve(t, runner.New(t.TempDir()), `{"jsonrpc":"2.0","id":7,"method":"server/discover"}`)
	var res struct {
		ID    json.RawMessage `json:"id"`
		Error *rpcError       `json:"error"`
	}
	if err := json.Unmarshal([]byte(lines[0]), &res); err != nil {
		t.Fatal(err)
	}
	if res.Error == nil || res.Error.Code != codeNoMethod {
		t.Fatalf("server/discover was answered with %+v", res)
	}
	if !strings.Contains(res.Error.Message, Revision) {
		t.Errorf("the error does not say which revision this is: %q", res.Error.Message)
	}
	if string(res.ID) != "7" {
		t.Errorf("the response carries id %s and the request carried 7", res.ID)
	}
}

// A notification has no id and the protocol forbids a response to it. The client sends
// notifications/initialized immediately after the handshake, so a server that answered it
// would put a response with a null id in front of every session.
func TestANotificationIsAnsweredWithSilence(t *testing.T) {
	lines := serve(t, runner.New(t.TempDir()),
		`{"jsonrpc":"2.0","method":"notifications/initialized"}`,
		`{"jsonrpc":"2.0","id":2,"method":"ping"}`)
	if len(lines) != 1 {
		t.Fatalf("a notification and a ping produced %d responses: %v", len(lines), lines)
	}
	if !strings.Contains(lines[0], `"id":2`) {
		t.Errorf("the one response is not the ping's: %s", lines[0])
	}
}

// One malformed line is a parse error and the session carries on. Exiting instead would take
// the rest of the client's tools down with it.
func TestALineThatIsNotJSONDoesNotEndTheSession(t *testing.T) {
	lines := serve(t, runner.New(t.TempDir()), `not json at all`,
		`{"jsonrpc":"2.0","id":3,"method":"ping"}`)
	if len(lines) != 2 {
		t.Fatalf("expected a parse error and a pong, got %v", lines)
	}
	if !strings.Contains(lines[0], `"code":-32700`) || !strings.Contains(lines[0], `"id":null`) {
		t.Errorf("the parse error is not the protocol's: %s", lines[0])
	}
	if !strings.Contains(lines[1], `"id":3`) {
		t.Errorf("the session did not carry on: %s", lines[1])
	}
}

// A refusal is the call happening and saying no, so it comes back inside the result with
// isError and not as a JSON-RPC error. The words are the runner's own, which is the point of
// the server being a shell: a caller is told what a person at a terminal would be told.
func TestARefusalComesBackInsideTheResult(t *testing.T) {
	f := drive(t)
	text, isError := call(t, f.r, recordAssum, map[string]any{
		"intent": fixtureKey, "phase": "00", "text": "something",
		"origin": "nowhere", "confidence": "high",
	})
	if !isError {
		t.Fatalf("an assumption with an invented origin was accepted: %s", text)
	}
	if !strings.Contains(text, "--origin is one of") {
		t.Errorf("the refusal is not the runner's: %q", text)
	}
}

// An absent index is ordinary and is an answer rather than a failure, which is the property
// everything about the index rests on: no phase and no verdict depends on it.
func TestAQueryWithNoIndexIsAnAnswerAndNotAFailure(t *testing.T) {
	f := drive(t)
	text, isError := call(t, f.r, querySymbol, map[string]any{"name": "Serve"})
	if isError {
		t.Fatalf("a repository with no index was reported as a failure: %s", text)
	}
	var res struct {
		Found []symbol `json:"found"`
		Why   string   `json:"why"`
	}
	if err := json.Unmarshal([]byte(text), &res); err != nil {
		t.Fatal(err)
	}
	if len(res.Found) != 0 || res.Why == "" {
		t.Errorf("the answer does not say why there is no index: %s", text)
	}
}

// ---- the acceptance: the artifacts are the same artifacts

const fixtureKey = "XENO-0001"

// approvedHost is one fake host for every fixture of this package, answering an approved
// issue. One and not one per fixture, because the qualified intent id carries the host it
// was derived from, and two fixtures on two ports would leave two intents with different
// ids — which the test comparing the routes by their artifacts would then read as the
// routes differing. It is never closed; it lives as long as the test binary does.
var approvedHost = sync.OnceValue(func() *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if strings.HasSuffix(req.URL.Path, "/comments") {
			_ = json.NewEncoder(w).Encode([]map[string]any{{
				"user": map[string]string{"login": "m"}, "body": "approved\n\nwanted"}})
			return
		}
		_ = json.NewEncoder(w).Encode(map[string]any{"title": "t", "body": "b",
			"labels": []map[string]string{{"name": "approved"}}})
	}))
})

// fixture is a repository with a vendored plugin, an intent and a started first phase, which
// is the state a phase's writes happen in whichever way they are made.
type fixture struct {
	root string
	r    *runner.Runner
}

func drive(t *testing.T) *fixture {
	t.Helper()
	t.Setenv(runner.HarnessVersionEnv, "")
	root := t.TempDir()
	r := runner.New(root)
	r.PluginSource = filepath.Join("..", "..", ".xeno", "plugin")
	r.Now = func() time.Time { return time.Date(2026, 10, 7, 9, 0, 0, 0, time.UTC) }
	if _, err := r.Init(runner.InitOptions{Host: "github", Vendor: true,
		TrackerKey: "triplem/xeno", Language: "en"}); err != nil {
		t.Fatal(err)
	}
	// From #330 a tracker block means `intent start` reads the issue and starts only on an
	// approval, so the block init wrote is pointed at a host that answers one.
	srv := approvedHost()
	t.Setenv("XENO_TRACKER_TOKEN", "a-token")
	cfg := filepath.Join(root, ".xeno", "config", "project.yaml")
	b, err := os.ReadFile(cfg)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(cfg, bytes.ReplaceAll(b, []byte("https://api.github.com"), []byte(srv.URL)), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := r.IntentStart(fixtureKey, "285", false); err != nil {
		t.Fatal(err)
	}
	// And no token from here on, so that the phase start reads nothing: the two routes are
	// compared by the artifacts they leave, and an intake quoting an issue from a host whose
	// port differs per fixture would differ for a reason that is not the route's.
	t.Setenv("XENO_TRACKER_TOKEN", "")
	// Written here because no command writes it: `xeno intent start` creates the intent and
	// the suggestion it prints says the register is written by hand. An assumption recorded
	// through either route needs it to exist, and the two routes need the same file.
	if err := os.WriteFile(filepath.Join(root, model.IntentDir(fixtureKey), "assumptions.yaml"),
		[]byte("assumptions: []\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := r.Start(fixtureKey, model.Phases[0]); err != nil {
		t.Fatal(err)
	}
	if _, err := r.ScopeSet(fixtureKey, []byte("include:\n  - internal/**\n")); err != nil {
		t.Fatal(err)
	}
	return &fixture{root: root, r: r}
}

// The clause of #285 that is the whole point of the server being a shell: a phase driven
// through the operations produces artifacts a gate cannot tell apart from one driven by
// commands.
//
// Two repositories, the same acts in the same order against the same clock, one through the
// tool calls and one through the runner methods the commands call. The comparison is over the
// bytes of output.md and over the verdict's artifacts_hash, which is what a gate reads: a
// difference anywhere inside that hash is a difference in the trail, and this is the test that
// would catch the server growing a second way to write an artifact.
func TestAPhaseDrivenThroughTheOperationsIsTheSameArtifact(t *testing.T) {
	phase := model.Phases[0]
	section := "problem"

	viaTools := drive(t)
	if text, isError := call(t, viaTools.r, writeArt, map[string]any{
		"intent": fixtureKey, "phase": "00", "section": section, "content": "The server is not built.",
	}); isError {
		t.Fatalf("write_artifact refused: %s", text)
	}
	if text, isError := call(t, viaTools.r, runGates, map[string]any{
		"intent": fixtureKey, "phase": "00",
	}); isError {
		t.Fatalf("run_gates refused: %s", text)
	}

	viaCommands := drive(t)
	if _, err := viaCommands.r.SectionSet(fixtureKey, phase, section, "The server is not built."); err != nil {
		t.Fatal(err)
	}
	if _, err := viaCommands.r.GateRun(fixtureKey, phase); err != nil {
		t.Fatal(err)
	}

	for _, rel := range []string{"output.md", "context.lock.yaml"} {
		a := read(t, viaTools.root, phase, rel)
		b := read(t, viaCommands.root, phase, rel)
		if a != b {
			t.Errorf("%s differs between the two runs:\n--- operations\n%s\n--- commands\n%s", rel, a, b)
		}
	}
	a, b := gate(t, viaTools.root, phase), gate(t, viaCommands.root, phase)
	if a.ArtifactsHash != b.ArtifactsHash {
		t.Errorf("the verdicts were taken over different artifacts: %s and %s",
			a.ArtifactsHash, b.ArtifactsHash)
	}
	if a.Status != b.Status {
		t.Errorf("the verdicts differ: %s and %s", a.Status, b.Status)
	}
	if a.ArtifactsHash == "" {
		t.Error("the verdict carries no artifacts_hash, so this compared nothing")
	}
}

// Recording an assumption through the operation puts it in the same register the command
// writes, which is the same property one file along: the register is inside the intent and is
// read by the gates.
func TestAnAssumptionRecordedThroughTheOperationIsInTheRegister(t *testing.T) {
	f := drive(t)
	text, isError := call(t, f.r, recordAssum, map[string]any{
		"intent": fixtureKey, "phase": "00", "text": "Both clients speak the handshake revision.",
		"origin": "user-input", "confidence": "medium",
	})
	if isError {
		t.Fatalf("record_assumption refused: %s", text)
	}
	var rec struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}
	if err := json.Unmarshal([]byte(text), &rec); err != nil {
		t.Fatal(err)
	}
	if rec.ID == "" || rec.Status != "open" {
		t.Fatalf("the answer names no open assumption: %s", text)
	}
	b, err := os.ReadFile(filepath.Join(f.root, model.IntentDir(fixtureKey), "assumptions.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(b), rec.ID) {
		t.Errorf("%s is not in the register:\n%s", rec.ID, b)
	}
}

// fetch_template answers with the sections the phase owes, and write_artifact is judged
// against the same template: the ids one returns are the ids the other accepts.
func TestTheTemplateFetchedIsTheTemplateWrittenAgainst(t *testing.T) {
	f := drive(t)
	text, isError := call(t, f.r, fetchTmpl, map[string]any{"phase": "00"})
	if isError {
		t.Fatalf("fetch_template refused: %s", text)
	}
	var res struct {
		Template string `json:"template"`
		Sections []struct {
			ID       string `json:"id"`
			Heading  string `json:"heading"`
			Required bool   `json:"required"`
		} `json:"sections"`
	}
	if err := json.Unmarshal([]byte(text), &res); err != nil {
		t.Fatal(err)
	}
	if res.Template == "" || len(res.Sections) == 0 {
		t.Fatalf("the intake template came back empty: %s", text)
	}
	for _, s := range res.Sections {
		if s.Heading == "" {
			t.Errorf("section %s came back with no heading", s.ID)
		}
	}
	// The first section the template names is one write_artifact accepts, which is the
	// agreement the two operations have to keep.
	if text, isError := call(t, f.r, writeArt, map[string]any{
		"intent": fixtureKey, "phase": "00", "section": res.Sections[0].ID, "content": "x",
	}); isError {
		t.Fatalf("the template's own first section was refused: %s", text)
	}
}

// read_intent reports the phases of the intent and what it has asked, which is what the
// status command prints from the same two calls.
func TestReadIntentReportsThePhasesOfTheIntent(t *testing.T) {
	f := drive(t)
	text, isError := call(t, f.r, readIntent, map[string]any{"intent": fixtureKey})
	if isError {
		t.Fatalf("read_intent refused: %s", text)
	}
	var res struct {
		Intent string `json:"intent"`
		Phases []struct {
			Phase string `json:"phase"`
			State string `json:"state"`
		} `json:"phases"`
	}
	if err := json.Unmarshal([]byte(text), &res); err != nil {
		t.Fatal(err)
	}
	if res.Intent != fixtureKey || len(res.Phases) != len(model.Phases) {
		t.Fatalf("read_intent answered %s with %d phases", res.Intent, len(res.Phases))
	}
	if res.Phases[0].State != "running" {
		t.Errorf("the started phase reads %q", res.Phases[0].State)
	}
}

func read(t *testing.T, root, phase, rel string) string {
	t.Helper()
	b, err := os.ReadFile(filepath.Join(root, model.PhaseDir(fixtureKey, phase), rel))
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func gate(t *testing.T, root, phase string) model.Gate {
	t.Helper()
	var g model.Gate
	b, err := os.ReadFile(filepath.Join(root, model.PhaseDir(fixtureKey, phase), "gate.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	if err := yaml.Unmarshal(b, &g); err != nil {
		t.Fatal(err)
	}
	return g
}
