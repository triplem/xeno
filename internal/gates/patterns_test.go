// SPDX-License-Identifier: Apache-2.0

package gates

import "testing"

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
