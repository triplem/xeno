// SPDX-License-Identifier: Apache-2.0

package mcp

import (
	"encoding/json"
	"fmt"

	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/runner"
)

// The six operations section 13 fixes: "read intent, record assumption, write artifact,
// fetch template, run gates locally, and query the symbol index. Six, and the surface is a
// budget rather than a list".
//
// The names are those six phrases and nothing else, so that what a client shows the model
// is what the specification enumerates. A seventh entry here needs the argument section 13
// asks for, which is why `order` is written out and `TestTheSurfaceIsSix` holds it: the cost
// of a definition is paid on every request of every phase whether the tool is called or not,
// so the thing to notice is a tool being added, not a tool being slow.
const (
	readIntent  = "read_intent"
	recordAssum = "record_assumption"
	writeArt    = "write_artifact"
	fetchTmpl   = "fetch_template"
	runGates    = "run_gates"
	querySymbol = "query_symbol_index"
)

// order is the order tools/list sends, which is the order section 13 names them in.
var order = []string{readIntent, recordAssum, writeArt, fetchTmpl, runGates, querySymbol}

// tool is one operation: the sentence the model reads, the schema the client validates
// against, and the call. The call takes the runner because every one of them is a front for
// a runner method a command also calls.
type tool struct {
	description string
	schema      string
	call        func(*runner.Runner, json.RawMessage) (string, error)
}

// The descriptions are one line each and say what the operation answers or writes, and
// nothing about when to reach for it. A paragraph of guidance belongs in the phase skill,
// which is loaded when the phase runs; a description is resident for the whole session.
var tools = map[string]tool{
	readIntent: {
		description: "The state of one intent: its phases, their verdicts, and what it has asked and decided.",
		schema:      schema(`"intent":{"type":"string"}`, "intent"),
		call:        callReadIntent,
	},
	recordAssum: {
		description: "Record one assumption of a phase, with where it came from and how much weight it carries.",
		schema: schema(`"intent":{"type":"string"},"phase":{"type":"string"},"text":{"type":"string"},`+
			`"origin":{"type":"string"},"confidence":{"type":"string"},"resolves":{"type":"string"}`,
			"intent", "phase", "text", "origin", "confidence"),
		call: callRecordAssumption,
	},
	writeArt: {
		description: "Write one section of a phase result; the file is rendered again from the template.",
		schema: schema(`"intent":{"type":"string"},"phase":{"type":"string"},`+
			`"section":{"type":"string"},"content":{"type":"string"}`,
			"intent", "phase", "section", "content"),
		call: callWriteArtifact,
	},
	fetchTmpl: {
		description: "The sections a phase owes, in order, with their headings and which are required.",
		schema:      schema(`"phase":{"type":"string"}`, "phase"),
		call:        callFetchTemplate,
	},
	runGates: {
		description: "Run the gates of a phase here and now and return the verdict with its findings.",
		schema:      schema(`"intent":{"type":"string"},"phase":{"type":"string"}`, "intent", "phase"),
		call:        callRunGates,
	},
	querySymbol: {
		description: "Where a name is defined, from the index the project produced; empty where there is none.",
		schema:      schema(`"name":{"type":"string"}`, "name"),
		call:        callQuerySymbol,
	},
}

// schema renders one input schema from its properties and the names of the required ones.
// Written as text rather than built from maps because a map marshals its keys in sorted
// order, and the order the properties are declared in is the order a client shows them.
func schema(properties string, required ...string) string {
	req, err := json.Marshal(required)
	if err != nil { // unreachable for a slice of strings, and silence would be worse than a panic
		panic(err)
	}
	return fmt.Sprintf(`{"type":"object","properties":{%s},"required":%s}`, properties, req)
}

// answer is what every operation returns: the result as compact JSON, because the consumer
// is a model and a sentence formatted for a terminal would be a second rendering of data
// the command already prints its own way.
func answer(v any) (string, error) {
	b, err := json.Marshal(v)
	return string(b), err
}

// phaseOf resolves a phase argument the way the dispatch in cmd/xeno does, so that "00" and
// "00-intake" mean here what they mean on the command line.
func phaseOf(arg string) (string, error) { return model.ResolvePhase(arg) }

func callReadIntent(r *runner.Runner, raw json.RawMessage) (string, error) {
	var a struct{ Intent string }
	if err := decode(raw, &a); err != nil {
		return "", err
	}
	states, err := r.Status(a.Intent)
	if err != nil {
		return "", err
	}
	s := r.Summarise(a.Intent)
	type phase struct {
		Phase   string `json:"phase"`
		State   string `json:"state"`
		Verdict string `json:"verdict,omitempty"`
		Stale   bool   `json:"stale,omitempty"`
	}
	out := struct {
		Intent    string  `json:"intent"`
		Created   string  `json:"created,omitempty"`
		State     string  `json:"state,omitempty"`
		Questions int     `json:"questions"`
		Decisions int     `json:"decisions"`
		Problem   string  `json:"problem,omitempty"`
		Phases    []phase `json:"phases"`
	}{Intent: s.Key, Created: s.Created, State: s.State,
		Questions: s.Questions, Decisions: s.Decisions, Problem: s.Problem}
	for _, st := range states {
		out.Phases = append(out.Phases, phase{st.Phase, st.State, st.Status, st.Stale})
	}
	return answer(out)
}

func callRecordAssumption(r *runner.Runner, raw json.RawMessage) (string, error) {
	var a struct{ Intent, Phase, Text, Origin, Confidence, Resolves string }
	if err := decode(raw, &a); err != nil {
		return "", err
	}
	phase, err := phaseOf(a.Phase)
	if err != nil {
		return "", err
	}
	rec, err := r.RecordAssumption(a.Intent, phase, a.Text, a.Origin, a.Confidence, a.Resolves)
	if err != nil {
		return "", err
	}
	// The id is what the answer is for: confirming or rejecting the assumption later names
	// it, and it is assigned by the writer rather than by the caller.
	return answer(struct {
		ID         string `json:"id"`
		Phase      string `json:"phase"`
		Assumption string `json:"assumption"`
		Origin     string `json:"origin,omitempty"`
		Confidence string `json:"confidence,omitempty"`
		Status     string `json:"status"`
		Resolves   string `json:"resolves,omitempty"`
	}{rec.ID, rec.Phase, rec.Assumption, rec.Origin, rec.Confidence, rec.Status, rec.Resolves})
}

func callWriteArtifact(r *runner.Runner, raw json.RawMessage) (string, error) {
	var a struct{ Intent, Phase, Section, Content string }
	if err := decode(raw, &a); err != nil {
		return "", err
	}
	phase, err := phaseOf(a.Phase)
	if err != nil {
		return "", err
	}
	t, err := r.SectionSet(a.Intent, phase, a.Section, a.Content)
	if err != nil {
		return "", err
	}
	// What the command prints after the same call: which template the file was rendered
	// from and whether that template came from the plugin or from the project.
	return answer(struct {
		Intent   string `json:"intent"`
		Phase    string `json:"phase"`
		Section  string `json:"section"`
		Template string `json:"template"`
		Source   string `json:"template_source"`
	}{a.Intent, phase, a.Section, t.Ref(), string(t.Source)})
}

func callFetchTemplate(r *runner.Runner, raw json.RawMessage) (string, error) {
	var a struct{ Phase string }
	if err := decode(raw, &a); err != nil {
		return "", err
	}
	phase, err := phaseOf(a.Phase)
	if err != nil {
		return "", err
	}
	t, err := r.Template(phase)
	if err != nil {
		return "", err
	}
	type section struct {
		ID       string `json:"id"`
		Heading  string `json:"heading"`
		Required bool   `json:"required"`
	}
	out := struct {
		Phase    string    `json:"phase"`
		Template string    `json:"template"`
		Source   string    `json:"template_source"`
		Language string    `json:"language"`
		Title    string    `json:"title,omitempty"`
		Sections []section `json:"sections"`
	}{Phase: phase, Template: t.Ref(), Source: string(t.Source),
		Language: t.Bundle.Language, Title: t.Bundle.Title}
	for _, s := range t.Template.Sections {
		out.Sections = append(out.Sections, section{s.ID, t.Bundle.Headings[s.ID], s.Required})
	}
	return answer(out)
}

// callRunGates is the local run, and it is the same GateRun the command calls: the verdict
// is written to gate.yaml by the runner exactly as it would be, because a gate that wrote
// nothing when called this way would make a phase driven through the operations unverifiable.
//
// The gate path makes no network call, which is the property the verdicts rest on, and
// nothing here changes that: this adds a caller to GateRun and no capability to it.
func callRunGates(r *runner.Runner, raw json.RawMessage) (string, error) {
	var a struct{ Intent, Phase string }
	if err := decode(raw, &a); err != nil {
		return "", err
	}
	phase, err := phaseOf(a.Phase)
	if err != nil {
		return "", err
	}
	g, err := r.GateRun(a.Intent, phase)
	if err != nil {
		return "", err
	}
	return answer(verdict(g))
}

// verdict is gate.yaml's verdict in the shape a model reads it: the status, the hash it was
// taken against, and per check its findings with the remedy each carries.
//
// It is projected rather than marshalled from model.Gate because that type carries yaml
// tags and no json ones, so marshalling it directly would send Go field names and an
// un-inlined Common block — a shape nothing else in this repository speaks, invented by
// the absence of a tag rather than chosen.
func verdict(g *model.Gate) any {
	type finding struct {
		ID       string `json:"id"`
		File     string `json:"file,omitempty"`
		Cause    string `json:"cause"`
		Next     string `json:"next,omitempty"`
		Advisory bool   `json:"advisory,omitempty"`
		Decision string `json:"decision,omitempty"`
	}
	type check struct {
		Gate     string    `json:"gate"`
		Result   string    `json:"result"`
		Findings []finding `json:"findings,omitempty"`
	}
	out := struct {
		Intent        string  `json:"intent"`
		Phase         string  `json:"phase"`
		Status        string  `json:"status"`
		ArtifactsHash string  `json:"artifacts_hash"`
		Checks        []check `json:"checks"`
	}{Intent: g.Intent, Phase: g.Phase, Status: g.Status, ArtifactsHash: g.ArtifactsHash}
	for _, c := range g.Checks {
		k := check{Gate: c.Gate, Result: c.Result}
		for _, f := range c.Findings {
			d := ""
			if f.Decision != nil {
				d = f.Decision.Type
			}
			k.Findings = append(k.Findings, finding{f.ID, f.File, f.Cause, f.Next, f.Advisory, d})
		}
		out.Checks = append(out.Checks, k)
	}
	return out
}

func callQuerySymbol(r *runner.Runner, raw json.RawMessage) (string, error) {
	var a struct{ Name string }
	if err := decode(raw, &a); err != nil {
		return "", err
	}
	s := r.Symbols(a.Name)
	return answer(struct {
		Name string `json:"name"`
		// The locations, and the provenance of whatever answered. Absent where there is no
		// index, with `why` carrying the sentence: an answer with no locations and no reason
		// cannot be told from a name the index does not hold.
		Found       []symbol `json:"found"`
		Tool        string   `json:"tool,omitempty"`
		ToolVersion string   `json:"tool_version,omitempty"`
		AgeMinutes  int      `json:"index_age_minutes,omitempty"`
		Why         string   `json:"why,omitempty"`
	}{Name: s.Name, Found: symbols(s), Tool: s.Tool, ToolVersion: s.ToolVersion,
		AgeMinutes: int(s.Age.Minutes()), Why: s.Why})
}

// symbol is one location, flattened to the four fields section 5 fixes plus the container.
type symbol struct {
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	File      string `json:"file"`
	Line      int    `json:"line"`
	Container string `json:"container,omitempty"`
}

func symbols(a *runner.SymbolAnswer) []symbol {
	out := make([]symbol, 0, len(a.Found))
	for _, s := range a.Found {
		out = append(out, symbol{s.Name, s.Kind, s.File, s.Line, s.Container})
	}
	return out
}

// decode reads a call's arguments. Absent arguments decode to the zero value rather than to
// an error, because the runner's own refusals say which field is missing and in the words the
// command says it in; an error here would be a second vocabulary for the same mistake.
func decode(raw json.RawMessage, into any) error {
	if len(raw) == 0 {
		return nil
	}
	if err := json.Unmarshal(raw, into); err != nil {
		return fmt.Errorf("the arguments are not what this tool takes: %v", err)
	}
	return nil
}
