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

// writeAt writes one file at a path inside the repository, creating the directories it needs,
// and commits it. Paths is about which paths changed, so the test needs paths rather than the
// flat files commit above produces.
func writeAt(t *testing.T, root, path, content string) {
	t.Helper()
	full := filepath.Join(root, path)
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	run(t, root, "add", ".")
	run(t, root, "commit", "-m", "writes "+path)
}

func TestThePathsOfARangeAreListedUnderTheDirectoryAsked(t *testing.T) {
	root := repo(t)
	writeAt(t, root, "README.md", "first")
	base := run(t, root, "rev-parse", "HEAD")
	writeAt(t, root, ".xeno/intents/PROJ-1/intent.yaml", "key: PROJ-1\n")
	writeAt(t, root, "internal/thing.go", "package thing\n")

	under, err := Paths(root, base, "HEAD", ".xeno/intents")
	if err != nil {
		t.Fatal(err)
	}
	if len(under) != 1 || under[0] != ".xeno/intents/PROJ-1/intent.yaml" {
		t.Fatalf("under the directory asked: %v", under)
	}
	all, err := Paths(root, base, "HEAD", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(all) != 2 {
		t.Fatalf("over the whole tree, want both: %v", all)
	}
}

// A directory created and dropped again inside one branch never existed at either end, which
// is why this is a tree comparison and not a log over the range.
func TestAPathAddedAndRemovedInsideTheRangeIsNotReported(t *testing.T) {
	root := repo(t)
	writeAt(t, root, "README.md", "first")
	base := run(t, root, "rev-parse", "HEAD")
	writeAt(t, root, ".xeno/intents/PROJ-9/intent.yaml", "key: PROJ-9\n")
	run(t, root, "rm", "-r", ".xeno/intents/PROJ-9")
	run(t, root, "commit", "-m", "and dropped again")

	paths, err := Paths(root, base, "HEAD", ".xeno/intents")
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) != 0 {
		t.Fatalf("nothing existed at either end, got %v", paths)
	}
}

// Rename detection is off, so a move reports both ends. A caller asking which intents a change
// touches is answerable for the one it moved away from as much as for the one it moved to.
func TestAMoveReportsBothPaths(t *testing.T) {
	root := repo(t)
	writeAt(t, root, ".xeno/intents/PROJ-1/intent.yaml", "key: PROJ-1\n")
	base := run(t, root, "rev-parse", "HEAD")
	run(t, root, "mv", ".xeno/intents/PROJ-1", ".xeno/intents/PROJ-2")
	run(t, root, "commit", "-m", "moved")

	paths, err := Paths(root, base, "HEAD", ".xeno/intents")
	if err != nil {
		t.Fatal(err)
	}
	if len(paths) != 2 {
		t.Fatalf("both ends of the move, got %v", paths)
	}
}

func TestPathsRefusesAnAbsentEndOfTheRange(t *testing.T) {
	root := repo(t)
	writeAt(t, root, "README.md", "first")
	if _, err := Paths(root, "HEAD", "", ""); err == nil {
		t.Fatal("an empty head is a caller error, not an empty range")
	}
	if _, err := Paths(root, "", "HEAD", ""); err == nil {
		t.Fatal("an empty base is a caller error, not an empty range")
	}
}

func TestPathsSaysWhichRefDidNotResolve(t *testing.T) {
	root := repo(t)
	writeAt(t, root, "README.md", "first")
	_, err := Paths(root, "no-such-ref", "HEAD", "")
	if err == nil {
		t.Fatal("want an error")
	}
	if !strings.Contains(err.Error(), "no-such-ref") {
		t.Fatalf("the error names the ref: %v", err)
	}
}
