// SPDX-License-Identifier: Apache-2.0

package model

import (
	"strings"
	"testing"
)

// TestTheQualifiedIdIsCompletedFromTheConfiguration covers the one rule Qualified holds:
// what a person left out comes from the tracker block, whatever they left out.
func TestTheQualifiedIdIsCompletedFromTheConfiguration(t *testing.T) {
	github := Tracker{Adapter: "github", Project: "triplem/xeno",
		BaseURL: "https://api.github.com"}
	cases := []struct {
		name    string
		tracker Tracker
		issue   string
		want    string
	}{
		{"a bare number", github, "176", "github.com/triplem/xeno#176"},
		{"the whole id", github, "github.com/triplem/xeno#176", "github.com/triplem/xeno#176"},
		{"the repository without the host", github, "triplem/xeno#176",
			"github.com/triplem/xeno#176"},
		{"another repository on the same host", github, "triplem/other#4",
			"github.com/triplem/other#4"},
		{"surrounding space", github, "  176  ", "github.com/triplem/xeno#176"},
		{"a key that is not a number", Tracker{Adapter: "jira", Project: "PROJ",
			BaseURL: "https://jira.example.com"}, "PROJ-123", "jira.example.com/PROJ#PROJ-123"},
		{
			// A self managed deployment serves its API under its own host, so there is no
			// api. prefix to take off and the host is the address as configured.
			name:    "a self managed deployment",
			tracker: Tracker{Adapter: "gitlab", Project: "group/proj", BaseURL: "https://git.example.com/api/v4"},
			issue:   "7",
			want:    "git.example.com/group/proj#7",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.tracker.Qualified(c.issue)
			if err != nil {
				t.Fatalf("Qualified(%q): %v", c.issue, err)
			}
			if got != c.want {
				t.Errorf("Qualified(%q) = %q, want %q", c.issue, got, c.want)
			}
		})
	}
}

// TestAnIdThatCannotBeCompletedIsRefused asserts the three cases where there is nothing to
// derive from. Each names the field or the flag, because the person reading it is the one
// who has to supply the part that is missing.
func TestAnIdThatCannotBeCompletedIsRefused(t *testing.T) {
	cases := []struct {
		name    string
		tracker Tracker
		issue   string
		says    string
	}{
		{"no issue at all", Tracker{Project: "triplem/xeno", BaseURL: "https://api.github.com"},
			"  ", "--for"},
		{"no project configured and no host given", Tracker{BaseURL: "https://api.github.com"},
			"176", "tracker.project"},
		{"no address to read the host off", Tracker{Project: "triplem/xeno"}, "176",
			"tracker.base_url"},
		{"a space inside the key", Tracker{Project: "triplem/xeno",
			BaseURL: "https://api.github.com"}, "176 and 177", "holds a space"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := c.tracker.Qualified(c.issue)
			if err == nil {
				t.Fatalf("Qualified(%q) = %q, want an error", c.issue, got)
			}
			if !strings.Contains(err.Error(), c.says) {
				t.Errorf("Qualified(%q) says %q, which does not name %q", c.issue, err, c.says)
			}
		})
	}
}

// TestTheNextKeyContinuesTheSequenceOnDisk covers the padding and the two schemes this
// repository has: keys up to XENO-0121 are issue numbers under the older one and the
// current sequence starts beyond them, so the highest number wins regardless of its width.
func TestTheNextKeyContinuesTheSequenceOnDisk(t *testing.T) {
	cases := []struct {
		name  string
		names []string
		want  string
	}{
		{"one key", []string{"XENO-0226"}, "XENO-0227"},
		{"the highest wins, not the last read", []string{"XENO-0226", "XENO-0105"}, "XENO-0227"},
		{"an unpadded key of the older scheme does not win", []string{"XENO-9", "XENO-0226"},
			"XENO-0227"},
		{"the padding is the one the highest key has", []string{"P-7"}, "P-8"},
		{"three digits stay three", []string{"P-007"}, "P-008"},
		{"a width is grown rather than truncated", []string{"P-999"}, "P-1000"},
		{"a prefix may hold a hyphen", []string{"TWO-WORDS-3"}, "TWO-WORDS-4"},
		{"names that are not keys are skipped", []string{"README", "notes", "XENO-0226"},
			"XENO-0227"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := NextKey(c.names)
			if err != nil {
				t.Fatalf("NextKey(%v): %v", c.names, err)
			}
			if got != c.want {
				t.Errorf("NextKey(%v) = %q, want %q", c.names, got, c.want)
			}
		})
	}
}

// TestASequenceThatCannotBeContinuedIsRefused asserts that the two cases where the
// directory does not answer ask for the key rather than inventing a prefix.
func TestASequenceThatCannotBeContinuedIsRefused(t *testing.T) {
	for _, names := range [][]string{nil, {}, {"README"}, {"XENO-0226", "OTHER-3"}} {
		got, err := NextKey(names)
		if err == nil {
			t.Fatalf("NextKey(%v) = %q, want an error", names, got)
		}
		if !strings.Contains(err.Error(), "--intent") {
			t.Errorf("NextKey(%v) says %q, which does not name --intent", names, err)
		}
	}
}

// Locate is Qualified read backwards, so the round trip is what asserts it: what the
// configuration completed, the same configuration takes apart again.
func TestLocateIsQualifiedReadBackwards(t *testing.T) {
	for _, tc := range []struct {
		tracker Tracker
		issue   string
		project string
		key     string
	}{
		{Tracker{Project: "triplem/xeno", BaseURL: "https://api.github.com"},
			"288", "triplem/xeno", "288"},
		{Tracker{Project: "group/proj", BaseURL: "https://git.example.com/api/v4"},
			"PROJ-7", "group/proj", "PROJ-7"},
		// A nested namespace, which is the case the host's own path separator makes
		// ambiguous if the split were on the first slash rather than on the host prefix.
		{Tracker{Project: "group/sub/proj", BaseURL: "https://git.example.com/api/v4"},
			"4", "group/sub/proj", "4"},
	} {
		id, err := tc.tracker.Qualified(tc.issue)
		if err != nil {
			t.Fatal(err)
		}
		project, key, err := tc.tracker.Locate(id)
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		if project != tc.project || key != tc.key {
			t.Errorf("%s locates %q %q, want %q %q", id, project, key, tc.project, tc.key)
		}
	}
}

// An issue on a host the configuration does not name cannot be reached by the adapter the
// configuration selects, and the error names both so that whoever reads it can tell which
// of the two is wrong.
func TestAnIssueOnAnotherHostNamesBothHosts(t *testing.T) {
	tr := Tracker{Project: "triplem/xeno", BaseURL: "https://api.github.com"}
	_, _, err := tr.Locate("git.example/group/proj#1")
	if err == nil {
		t.Fatal("an issue on another host was located anyway")
	}
	for _, want := range []string{"git.example/group/proj", "github.com"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("the error does not name %q: %v", want, err)
		}
	}
}

func TestSomethingThatIsNotAQualifiedIdIsRefused(t *testing.T) {
	tr := Tracker{Project: "triplem/xeno", BaseURL: "https://api.github.com"}
	for _, id := range []string{"", "288", "github.com/triplem/xeno", "github.com/triplem/xeno#"} {
		if _, _, err := tr.Locate(id); err == nil {
			t.Errorf("%q was read as a qualified id", id)
		}
	}
}

// ---- section 12's "Starting an intent": what approval means, judged here and not by an
// adapter, so that two hosts cannot answer it differently.

func TestApprovalNeedsTheLabelAndTheWord(t *testing.T) {
	word := Comment{Author: "m", At: "2026-10-08T15:10:28Z", Body: "Approved.\n\nbecause it was decided\non the issue"}
	cases := []struct {
		name    string
		issue   Issue
		missing int
		reason  string
	}{
		{"neither", Issue{}, 2, ""},
		{"the label alone", Issue{Labels: []string{"approved"}}, 1, ""},
		{"the word alone", Issue{Comments: []Comment{word}}, 1, ""},
		{"both, with case and punctuation ignored", Issue{Labels: []string{"wp12", "Approved"},
			Comments: []Comment{word}}, 0, "because it was decided\non the issue"},
		{"a comment that merely contains the word", Issue{Labels: []string{"approved"},
			Comments: []Comment{{Body: "this could be approved later"}}}, 1, ""},
		{"the word with nothing after it", Issue{Labels: []string{"approved"},
			Comments: []Comment{{Author: "m", Body: "approved"}}}, 0, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			a, missing := c.issue.Approval()
			if len(missing) != c.missing {
				t.Fatalf("missing %v, want %d things", missing, c.missing)
			}
			if c.missing == 0 && (a.By != "m" || a.Reason != c.reason) {
				t.Errorf("read %+v", a)
			}
		})
	}
}

// The last approving comment stands, because an approval is withdrawn and given again by
// writing another, and the one that stands is the latest.
func TestTheLastApprovingCommentStands(t *testing.T) {
	i := Issue{Labels: []string{"approved"}, Comments: []Comment{
		{Author: "first", Body: "approved\nfor the wrong reason"},
		{Author: "second", Body: "approved\nfor the right one"},
	}}
	a, _ := i.Approval()
	if a.By != "second" || a.Reason != "for the right one" {
		t.Errorf("read %+v", a)
	}
}

// Ahead orders the open milestones by due date, undated ones last, then by number, and
// answers independently of the order the host listed them in.
func TestAheadIsTheEarliestOpenMilestoneThatIsNotTheIssues(t *testing.T) {
	open := []Milestone{
		{Title: "later", Number: 3},
		{Title: "1.1", Due: "2026-12-31", Number: 2},
		{Title: "1.0", Due: "2026-11-30", Number: 1},
		{Title: "done", Due: "2026-01-01", Number: 0, Closed: true},
	}
	cases := []struct {
		name string
		on   *Milestone
		want string
	}{
		{"none", nil, ""},
		{"the earliest", &Milestone{Title: "1.0"}, ""},
		{"the second", &Milestone{Title: "1.1"}, "1.0"},
		{"the undated one", &Milestone{Title: "later"}, "1.0"},
		{"a closed one", &Milestone{Title: "0.9", Closed: true}, ""},
		{"with no open milestones at all", &Milestone{Title: "1.1"}, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			i := Issue{Milestone: c.on, OpenMilestones: open}
			if c.name == "with no open milestones at all" {
				i.OpenMilestones = nil
			}
			got := i.Ahead()
			switch {
			case c.want == "" && got != nil:
				t.Errorf("held by %+v", got)
			case c.want != "" && (got == nil || got.Title != c.want):
				t.Errorf("got %+v, want %s", got, c.want)
			}
		})
	}
	// Two undated milestones: the lower number first.
	i := Issue{Milestone: &Milestone{Title: "b"}, OpenMilestones: []Milestone{{Title: "b", Number: 2}, {Title: "a", Number: 1}}}
	if got := i.Ahead(); got == nil || got.Title != "a" {
		t.Errorf("between two undated milestones got %+v, want a", got)
	}
}
