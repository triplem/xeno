// SPDX-License-Identifier: Apache-2.0

package cost

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func transcript(t *testing.T, lines ...string) string {
	t.Helper()
	p := filepath.Join(t.TempDir(), "t.jsonl")
	if err := os.WriteFile(p, []byte(strings.Join(lines, "\n")+"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	return p
}

// line is one assistant record of a transcript, built by hand rather than with fmt so that the
// test reads as the JSONL it is testing.
func line(in, out, read, create int) string {
	return `{"type":"assistant","message":{"usage":{"input_tokens":` + itoa(in) +
		`,"output_tokens":` + itoa(out) + `,"cache_read_input_tokens":` + itoa(read) +
		`,"cache_creation_input_tokens":` + itoa(create) + `}}}`
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}

// The cached share is the two cache fields together, because section 11 keeps input, output and
// cached apart and a cached prefix costs a fraction of either.
func TestTranscriptTotalsSumsEveryAssistantMessage(t *testing.T) {
	p := transcript(t, line(10, 5, 100, 20), line(1, 2, 3, 4))
	got, err := TranscriptTotals(p)
	if err != nil {
		t.Fatal(err)
	}
	want := Totals{In: 11, Out: 7, Cached: 127}
	if got != want {
		t.Fatalf("totals %+v, want %+v", got, want)
	}
}

// A transcript is written while a session runs, so its last line may be half there, and a record
// without usage is most of the file.
func TestAHalfWrittenLineIsSkippedRatherThanFatal(t *testing.T) {
	p := transcript(t, line(10, 5, 0, 0), `{"type":"user","message":{"content":"x"}}`,
		`{"type":"assistant","message":{"usa`)
	got, err := TranscriptTotals(p)
	if err != nil {
		t.Fatalf("a truncated line failed the read: %v", err)
	}
	if got != (Totals{In: 10, Out: 5}) {
		t.Fatalf("totals %+v, want the one complete record", got)
	}
}

// The ledger records cumulative totals, so a turn's cost is the difference between two lines and
// a phase's total is the sum of those differences. The first line of a session has no predecessor
// and contributes nothing: what it counts was spent before the ledger existed.
func TestAPhaseTotalsTheDifferencesBetweenItsLines(t *testing.T) {
	root := t.TempDir()
	for _, turn := range []Turn{
		{At: "1", Session: "s", Phase: NoPhase, Totals: Totals{In: 1, Out: 100, Cached: 1000}},
		{At: "2", Session: "s", Intent: "i", Phase: "02-design", Totals: Totals{In: 2, Out: 300, Cached: 3000}},
		{At: "3", Session: "s", Intent: "i", Phase: "02-design", Totals: Totals{In: 3, Out: 600, Cached: 6000}},
		{At: "4", Session: "s", Phase: NoPhase, Totals: Totals{In: 9, Out: 900, Cached: 9000}},
	} {
		if err := Append(root, turn); err != nil {
			t.Fatal(err)
		}
	}
	got, sessions, err := ForPhase(root, "i", "02-design")
	if err != nil {
		t.Fatal(err)
	}
	// Two lines name the phase: 300-100 and 600-300 of output.
	want := Totals{In: 2, Out: 500, Cached: 5000}
	if got != want {
		t.Fatalf("totals %+v, want %+v", got, want)
	}
	if len(sessions) != 1 || sessions[0] != "s" {
		t.Errorf("sessions %v, want the one it drew on", sessions)
	}
}

// A turn spent with no phase open is recorded and attributed to nothing, which is what makes the
// gap visible instead of absent (A63).
func TestAnUnattributedTurnIsRecordedAndCountedNowhere(t *testing.T) {
	root := t.TempDir()
	for i, ph := range []string{NoPhase, NoPhase, NoPhase} {
		if err := Append(root, Turn{At: itoa(i), Session: "s", Phase: ph,
			Totals: Totals{Out: 100 * (i + 1)}}); err != nil {
			t.Fatal(err)
		}
	}
	turns, err := Read(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(turns) != 3 {
		t.Fatalf("%d lines in the ledger, want every turn", len(turns))
	}
	got, _, err := ForPhase(root, "i", "02-design")
	if err != nil {
		t.Fatal(err)
	}
	if !got.Zero() {
		t.Fatalf("unattributed turns were counted somewhere: %+v", got)
	}
}

// Sessions are kept apart, or a difference would be taken across two transcripts.
func TestTwoSessionsAreNotSubtractedFromEachOther(t *testing.T) {
	root := t.TempDir()
	for _, turn := range []Turn{
		{At: "1", Session: "a", Intent: "i", Phase: "p", Totals: Totals{Out: 100}},
		{At: "2", Session: "b", Intent: "i", Phase: "p", Totals: Totals{Out: 9000}},
		{At: "3", Session: "a", Intent: "i", Phase: "p", Totals: Totals{Out: 150}},
		{At: "4", Session: "b", Intent: "i", Phase: "p", Totals: Totals{Out: 9100}},
	} {
		if err := Append(root, turn); err != nil {
			t.Fatal(err)
		}
	}
	got, sessions, err := ForPhase(root, "i", "p")
	if err != nil {
		t.Fatal(err)
	}
	if got.Out != 150 { // 150-100 within a, 9100-9000 within b
		t.Fatalf("output %d, want 150: the sessions were mixed", got.Out)
	}
	if len(sessions) != 2 {
		t.Errorf("sessions %v, want both", sessions)
	}
}

// The file phase start writes is shell, so a line reads export XENO_PHASE="02-design".
func TestTheLivePhaseIsReadFromTheRunnersOwnMarker(t *testing.T) {
	root := t.TempDir()
	rel := "phase.env"
	if err := os.WriteFile(filepath.Join(root, rel),
		[]byte("export XENO_INTENT=\"git.example/p#1\"\nexport XENO_PHASE=\"02-design\"\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	intent, phase := LivePhase(root, rel)
	if intent != "git.example/p#1" || phase != "02-design" {
		t.Fatalf("read %q and %q", intent, phase)
	}
	// No marker is no phase, which is the ordinary state between phases.
	i2, p2 := LivePhase(root, "absent.env")
	if i2 != "" || p2 != "" {
		t.Errorf("an absent marker produced %q and %q", i2, p2)
	}
}

// No ledger is not an error: a repository where the hook has never run has nothing to total.
func TestNoLedgerIsEmptyRatherThanAnError(t *testing.T) {
	turns, err := Read(t.TempDir())
	if err != nil {
		t.Fatalf("an absent ledger failed the read: %v", err)
	}
	if len(turns) != 0 {
		t.Fatalf("%d turns from nothing", len(turns))
	}
}
