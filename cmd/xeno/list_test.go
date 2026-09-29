// SPDX-License-Identifier: Apache-2.0

package main

import (
	"strings"
	"testing"
)

// The heading and every row share one format string, which is what keeps a heading from drifting
// from the column it labels. This asserts the property rather than the string: the heading's own
// columns have to start where a row's do, at the widest value each column can take.
func TestTheHeadingLinesUpWithTheWidestRow(t *testing.T) {
	head := sprintRow(listRow, "CREATED", "INTENT", "STATE", "PHASE")
	// The widest values this listing can print: a date, a key as long as the longest here, the
	// longest state, and a phase with its verdict.
	row := sprintRow(listRow, "2026-09-29", "XENO-0207", "03-implementation", "05-review  green")
	for _, col := range []string{"INTENT", "STATE", "PHASE"} {
		at := strings.Index(head, col)
		if at < 0 {
			t.Fatalf("the heading does not name %q: %q", col, head)
		}
		if at >= len(row) || row[at] == ' ' {
			t.Errorf("the %q heading at column %d labels whitespace in the widest row:\n  %s\n  %s",
				col, at, head, row)
		}
	}
}

// The one-intent form puts a position beside a judgement, which is the pair a heading is worth most
// for.
func TestThePhaseHeadingLinesUpToo(t *testing.T) {
	head := sprintRow(phaseRow, "PHASE", "STATE", "VERDICT")
	row := sprintRow(phaseRow, "03-implementation", "changed-after-verdict", "provisional")
	for _, col := range []string{"STATE", "VERDICT"} {
		at := strings.Index(head, col)
		if at < 0 || at >= len(row) || row[at] == ' ' {
			t.Errorf("the %q heading does not sit over its column:\n  %s\n  %s", col, head, row)
		}
	}
}

// Ten is what a listing shows, and the number is the one the notice counts against.
func TestTheDefaultIsTenAndTheNoticeCountsTheRest(t *testing.T) {
	if listDefault != 10 {
		t.Fatalf("the default is %d, and the issue asked for ten", listDefault)
	}
	for _, tc := range []struct{ total, wantShown, wantHidden int }{
		{0, 0, 0},
		{3, 3, 0},
		{10, 10, 0},
		{11, 10, 1},
		{63, 10, 53},
	} {
		shown, hidden := tail(tc.total, false)
		if shown != tc.wantShown || hidden != tc.wantHidden {
			t.Errorf("%d intents: shows %d and hides %d, want %d and %d",
				tc.total, shown, hidden, tc.wantShown, tc.wantHidden)
		}
		// --all hides nothing, whatever the count.
		if shown, hidden := tail(tc.total, true); shown != tc.total || hidden != 0 {
			t.Errorf("%d intents with --all: shows %d and hides %d", tc.total, shown, hidden)
		}
	}
}
