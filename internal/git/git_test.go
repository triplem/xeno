// SPDX-License-Identifier: Apache-2.0

package git

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// repo builds a real repository, because what is under test is what git reports and a fake
// would be a second implementation of the thing being read. A machine without git skips.
func repo(t *testing.T) string {
	t.Helper()
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not on the path")
	}
	root := t.TempDir()
	for _, args := range [][]string{
		{"init", "--initial-branch=main"},
		{"config", "user.name", "A Committer"},
		{"config", "user.email", "committer@example.test"},
		{"config", "commit.gpgsign", "false"},
	} {
		run(t, root, args...)
	}
	return root
}

func run(t *testing.T, root string, args ...string) string {
	t.Helper()
	cmd := exec.Command("git", append([]string{"-C", root}, args...)...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("git %s: %v\n%s", strings.Join(args, " "), err, out)
	}
	return strings.TrimSpace(string(out))
}

// commit writes a file and commits it with the message given, which may carry trailers. The
// file is named after the whole message rather than a prefix of it: two subjects that share
// eight characters would otherwise write the same file and collide on a merge, which is how
// this helper first went wrong.
func commit(t *testing.T, root, message string) string {
	t.Helper()
	safe := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9' {
			return r
		}
		return '-'
	}, message)
	if err := os.WriteFile(filepath.Join(root, safe+".txt"), []byte(message), 0o644); err != nil {
		t.Fatal(err)
	}
	run(t, root, "add", ".")
	run(t, root, "commit", "-m", message)
	return run(t, root, "rev-parse", "HEAD")
}

func TestTheRangeIsReadOldestFirst(t *testing.T) {
	root := repo(t)
	commit(t, root, "first: the base")
	base := run(t, root, "rev-parse", "HEAD")
	commit(t, root, "second: one")
	commit(t, root, "third: two")

	cs, err := Commits(root, base, "HEAD")
	if err != nil {
		t.Fatal(err)
	}
	if len(cs) != 2 {
		t.Fatalf("read %d commits, want the two after the base", len(cs))
	}
	if cs[0].Subject != "second: one" || cs[1].Subject != "third: two" {
		t.Fatalf("subjects are %q and %q, want oldest first", cs[0].Subject, cs[1].Subject)
	}
	if cs[0].Email != "committer@example.test" || cs[0].Author != "A Committer" {
		t.Fatalf("author is %q <%q>, want the configured one", cs[0].Author, cs[0].Email)
	}
	if cs[0].Parents != 1 || cs[0].IsMerge() {
		t.Fatalf("a commit with one parent reads as %d parents, merge %v", cs[0].Parents, cs[0].IsMerge())
	}
}

// A subject is whatever somebody typed, including the characters a format string would
// otherwise be cut on.
func TestASubjectWithAwkwardCharactersSurvives(t *testing.T) {
	root := repo(t)
	commit(t, root, "first: the base")
	base := run(t, root, "rev-parse", "HEAD")
	awkward := `fix(parser): handle a | pipe, a \t tab and a "quote"`
	commit(t, root, awkward)

	cs, err := Commits(root, base, "HEAD")
	if err != nil {
		t.Fatal(err)
	}
	if len(cs) != 1 || cs[0].Subject != awkward {
		t.Fatalf("subject is %q, want %q", cs[0].Subject, awkward)
	}
}

func TestTrailersAreReadByKeyAndCaseInsensitively(t *testing.T) {
	root := repo(t)
	commit(t, root, "first: the base")
	base := run(t, root, "rev-parse", "HEAD")
	commit(t, root, "feat: something\n\nXeno-Intent: XENO-0219\nCo-Authored-By: Somebody <s@example.test>")

	cs, err := Commits(root, base, "HEAD")
	if err != nil {
		t.Fatal(err)
	}
	c := cs[0]
	for _, key := range []string{"Xeno-Intent", "xeno-intent", "XENO-INTENT", "Xeno-Intent:"} {
		if !c.HasTrailer(key) {
			t.Errorf("HasTrailer(%q) is false", key)
		}
	}
	if c.HasTrailer("Signed-off-by") {
		t.Error("a trailer the commit does not carry reads as present")
	}
}

// Git answers with a letter and the rule needs a boolean. An unsigned commit is N, which is not
// signed for this purpose; G and U are, because trust lives in the verifier's keyring (A70).
func TestAnUnsignedCommitIsNotSigned(t *testing.T) {
	root := repo(t)
	commit(t, root, "first: the base")
	base := run(t, root, "rev-parse", "HEAD")
	commit(t, root, "feat: unsigned")

	cs, err := Commits(root, base, "HEAD")
	if err != nil {
		t.Fatal(err)
	}
	if cs[0].Signed() {
		t.Fatalf("an unsigned commit reads as signed, status %q", cs[0].Signature)
	}
	for _, status := range []string{"G", "U"} {
		if !(Commit{Signature: status}).Signed() {
			t.Errorf("status %s does not read as signed", status)
		}
	}
	for _, status := range []string{"B", "X", "Y", "R", "E", "N", ""} {
		if (Commit{Signature: status}).Signed() {
			t.Errorf("status %q reads as signed", status)
		}
	}
}

func TestAMergeCommitIsRecognised(t *testing.T) {
	root := repo(t)
	commit(t, root, "first: the base")
	base := run(t, root, "rev-parse", "HEAD")
	run(t, root, "checkout", "-b", "side")
	commit(t, root, "feat: on the side")
	run(t, root, "checkout", "main")
	commit(t, root, "feat: on main")
	run(t, root, "merge", "--no-ff", "-m", "merge: the side branch", "side")

	cs, err := Commits(root, base, "HEAD")
	if err != nil {
		t.Fatal(err)
	}
	var merges int
	for _, c := range cs {
		if c.IsMerge() {
			merges++
		}
	}
	if merges != 1 {
		t.Fatalf("%d merge commits in the range, want 1", merges)
	}
}

// Section 12 forbids inferring a range, so there is no default to fall back to and an absent
// end is an error the caller words as a finding.
func TestAnAbsentEndOfTheRangeIsAnError(t *testing.T) {
	root := repo(t)
	commit(t, root, "first: the base")
	for _, c := range [][2]string{{"", "HEAD"}, {"HEAD", ""}, {"", ""}} {
		if _, err := Commits(root, c[0], c[1]); err == nil {
			t.Errorf("base %q head %q produced no error", c[0], c[1])
		}
	}
}

// The error says which ref did not resolve, because the exit status alone says only that
// something did not.
func TestARefThatDoesNotResolveSaysSo(t *testing.T) {
	root := repo(t)
	commit(t, root, "first: the base")
	_, err := Commits(root, "HEAD", "a-branch-nobody-made")
	if err == nil {
		t.Fatal("an unresolvable ref produced no error")
	}
	if !strings.Contains(err.Error(), "a-branch-nobody-made") {
		t.Fatalf("the error is %q, want the ref named", err)
	}
}

func TestADirectoryThatIsNoRepositoryIsAnError(t *testing.T) {
	if _, err := exec.LookPath("git"); err != nil {
		t.Skip("git is not on the path")
	}
	if _, err := Commits(t.TempDir(), "HEAD~1", "HEAD"); err == nil {
		t.Fatal("a directory that is no repository produced no error")
	}
}

func TestAnEmptyRangeIsNoCommitsAndNoError(t *testing.T) {
	root := repo(t)
	commit(t, root, "first: the base")
	cs, err := Commits(root, "HEAD", "HEAD")
	if err != nil {
		t.Fatal(err)
	}
	if len(cs) != 0 {
		t.Fatalf("read %d commits from HEAD..HEAD, want none", len(cs))
	}
}
