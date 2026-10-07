// SPDX-License-Identifier: Apache-2.0

// Package cost implements the cost record of section 11: the token counts a session reports,
// attributed to the phase that was open when they were spent.
//
// Section 11 assumes one session per phase and the agent working inside it. Neither holds by
// itself, so attribution is not inferred here. A hook appends one line per turn naming the phase
// the runner said was open, and a turn with none is recorded as such: measured against this
// repository, filtering a transcript by each phase's own window captures about seven per cent of
// what a session spent, because the work is done before `phase start` is called. A rule that
// produced a number for every turn would hide that; a ledger that says `none` shows it.
//
// Nothing here reads a message. A transcript is a whole conversation and a repository is not a
// place for one, so only usage counts and identifiers are taken from it.
package cost

import (
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"go.yaml.in/yaml/v3"

	"github.com/triplem/xeno/internal/model"
)

// LedgerFile is the name of the turn ledger, inside the local data location: it is a
// machine's own record and not part of the trail, so it lies outside artifacts_hash and is
// gitignored with the rest of that directory.
//
// A file name rather than a path from #205. The directory is XENO_PLUGIN_DATA's to decide and
// model.LocalDir resolves it, so this package and the runner cannot disagree about where a
// repository's local data is.
const LedgerFile = "cost-ledger.yaml"

// LedgerPath is the ledger of one repository, resolved.
func LedgerPath(root string) string { return model.LocalPath(root, LedgerFile) }

// NoPhase is what a turn spent with no phase open records.
const NoPhase = "none"

// Totals are the three counts section 11 keeps apart, because output costs several times input
// and a cached prefix a fraction of either.
type Totals struct {
	In     int `yaml:"tokens_in"`
	Out    int `yaml:"tokens_out"`
	Cached int `yaml:"tokens_cached"`
}

func (t Totals) add(o Totals) Totals {
	return Totals{t.In + o.In, t.Out + o.Out, t.Cached + o.Cached}
}

func (t Totals) sub(o Totals) Totals {
	return Totals{t.In - o.In, t.Out - o.Out, t.Cached - o.Cached}
}

// Zero reports whether anything was spent, which is what decides that a phase writes no file.
func (t Totals) Zero() bool { return t.In == 0 && t.Out == 0 && t.Cached == 0 }

// Turn is one line of the ledger: what a session had spent in total when a turn ended, and which
// phase was open. Cumulative rather than a delta, so that a line can be checked against the
// transcript it came from and a missed turn widens a gap instead of losing one.
type Turn struct {
	At      string `yaml:"at"`
	Session string `yaml:"session"`
	Intent  string `yaml:"intent,omitempty"`
	Phase   string `yaml:"phase"`
	Totals  `yaml:",inline"`
}

// TranscriptTotals sums the usage of every assistant message in a JSONL transcript. A record
// without usage, or a line that will not parse, is skipped: a transcript is written while a
// session runs and its last line may be half there.
func TranscriptTotals(path string) (Totals, error) {
	f, err := os.Open(path)
	if err != nil {
		return Totals{}, err
	}
	defer f.Close()

	var t Totals
	s := bufio.NewScanner(f)
	s.Buffer(make([]byte, 0, 1024*1024), 16*1024*1024) // a turn's record can be large
	for s.Scan() {
		var rec struct {
			Message struct {
				Usage struct {
					Input       int `json:"input_tokens"`
					Output      int `json:"output_tokens"`
					CacheRead   int `json:"cache_read_input_tokens"`
					CacheCreate int `json:"cache_creation_input_tokens"`
				} `json:"usage"`
			} `json:"message"`
		}
		if err := json.Unmarshal(s.Bytes(), &rec); err != nil {
			continue
		}
		u := rec.Message.Usage
		t = t.add(Totals{u.Input, u.Output, u.CacheRead + u.CacheCreate})
	}
	return t, nil
}

// Append adds one line to the ledger, creating it where it does not exist.
func Append(root string, turn Turn) error {
	path := LedgerPath(root)
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	turns, err := Read(root)
	if err != nil {
		return err
	}
	turns = append(turns, turn)
	b, err := yaml.Marshal(turns)
	if err != nil {
		return err
	}
	return os.WriteFile(path, b, 0o644)
}

// Read returns the ledger, or nothing where there is none.
func Read(root string) ([]Turn, error) {
	b, err := os.ReadFile(LedgerPath(root))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	var turns []Turn
	if err := yaml.Unmarshal(b, &turns); err != nil {
		return nil, err
	}
	return turns, nil
}

// ForPhase totals what was attributed to one phase of one intent, and names the sessions it came
// from.
//
// The cost of a turn is the difference between its line and the line before it in the same
// session, so a phase's total is the sum of those differences over the lines naming it. The first
// line of a session has no predecessor and contributes nothing, which is correct: whatever it
// counts was spent before the ledger existed.
func ForPhase(root, intent, phase string) (Totals, []string, error) {
	turns, err := Read(root)
	if err != nil {
		return Totals{}, nil, err
	}
	bySession := map[string][]Turn{}
	for _, t := range turns {
		bySession[t.Session] = append(bySession[t.Session], t)
	}
	var total Totals
	seen := map[string]bool{}
	for session, list := range bySession {
		for i := 1; i < len(list); i++ {
			if list[i].Phase != phase || list[i].Intent != intent {
				continue
			}
			total = total.add(list[i].Totals.sub(list[i-1].Totals))
			seen[session] = true
		}
	}
	sessions := make([]string, 0, len(seen))
	for s := range seen {
		sessions = append(sessions, s)
	}
	sort.Strings(sessions)
	return total, sessions, nil
}

// LivePhase reads the intent and phase the runner said were open, from the file phase start
// writes and phase finish removes. Empty where none is.
//
// The path is complete and is not joined to anything here. It used to be a root and a path
// relative to it, and the caller resolved the second with model.LocalPath, which returns the
// absolute value of XENO_PLUGIN_DATA as it stands — so the join turned an absolute marker into
// a relative one and every turn of every hooked session recorded `none`. A resolver this
// function calls and a resolver its caller calls are one resolver too many, so there is one
// argument and nothing to disagree about (#287).
func LivePhase(path string) (intent, phase string) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", ""
	}
	for _, line := range strings.Split(string(b), "\n") {
		k, v, ok := strings.Cut(strings.TrimSpace(line), "=")
		if !ok {
			continue
		}
		v = strings.Trim(strings.TrimSpace(v), `"`)
		// The file is shell, so a line reads `export XENO_PHASE="02-design"`.
		switch strings.TrimPrefix(strings.TrimSpace(k), "export ") {
		case "XENO_INTENT":
			intent = v
		case "XENO_PHASE":
			phase = v
		}
	}
	return intent, phase
}
