// SPDX-License-Identifier: Apache-2.0

package gates

import (
	"strings"
	"testing"
)

func TestConventionalCommitsPattern(t *testing.T) {
	ok := []string{
		"feat: a thing",
		"fix(gates): a thing",
		"feat(x)!: a breaking thing",
		"docs: a thing (#8)",
		"chore(release): 0.8.0 [skip ci]",
		"revert: feat: a thing",
		"feat: a thing\n\nand a body\n\nRefs #7\n",
	}
	for _, m := range ok {
		if err := CheckMessage("conventional-commits", m); err != nil {
			t.Errorf("rejected a well formed subject %q: %v", m, err)
		}
	}
	bad := []string{
		"Bump actions/checkout from 4 to 7",
		"feat a thing",
		"feat:",
		"feat: ",
		"Feat: a thing",
		"feature: a thing",
		"",
		"\n\nonly a body",
	}
	for _, m := range bad {
		if err := CheckMessage("conventional-commits", m); err == nil {
			t.Errorf("accepted %q", m)
		}
	}
}

// Only the subject is read. A body that would not pass as a subject changes nothing.
func TestOnlyTheSubjectIsRead(t *testing.T) {
	if err := CheckMessage("conventional-commits", "feat: a thing\nBump something from 4 to 7\n"); err != nil {
		t.Fatalf("a body was judged as if it were a subject: %v", err)
	}
}

func TestAnUnknownPatternSaysWhatThereIs(t *testing.T) {
	err := CheckMessage("nope", "feat: a thing")
	if err == nil {
		t.Fatal("an unknown pattern was accepted")
	}
	if got := err.Error(); got == "" || !contains(got, "conventional-commits") {
		t.Fatalf("the error does not name what is shipped: %v", err)
	}
}

func contains(s, sub string) bool {
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}

// The second shipped pattern: the same form carrying an issue reference in the subject, which
// is what survives a squash merge. Both hosts' spellings are accepted (A71).
func TestTheWithIssuePattern(t *testing.T) {
	for _, tc := range []struct {
		subject string
		want    bool
	}{
		{"feat: the thing (#162)", true},
		{"fix(rules): the other thing (!41)", true},
		{"feat!: a breaking change (#1)", true},
		{"chore(deps)!: bump something (#99999)", true},
		{"feat: the thing", false},
		{"feat: the thing (#)", false},
		{"feat: the thing (#abc)", false},
		{"feat: (#162)", false},
		{"feat: the thing (#162) and more", false},
		{"the thing (#162)", false},
	} {
		err := CheckMessage("conventional-commits-with-issue", tc.subject)
		if (err == nil) != tc.want {
			t.Errorf("%q: matched %v, want %v (%v)", tc.subject, err == nil, tc.want, err)
		}
	}
}

// One map answers at the gate and in xeno check commit-message, which is the property section 9
// asks for by name: a hook cannot start rejecting what a gate accepts.
func TestBothPatternsAreNamedInTheListing(t *testing.T) {
	names := strings.Join(PatternNames(), ",")
	for _, want := range []string{"conventional-commits", "conventional-commits-with-issue"} {
		if !strings.Contains(names, want) {
			t.Errorf("PatternNames is %q, want %s in it", names, want)
		}
	}
	if len(PatternNames()) != 2 {
		t.Errorf("%d shipped patterns, want the two section 9 names", len(PatternNames()))
	}
}
