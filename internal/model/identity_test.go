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
