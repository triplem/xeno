// SPDX-License-Identifier: Apache-2.0

package model_test

import (
	"path/filepath"
	"testing"

	"github.com/triplem/xeno/internal/model"
)

// XENO_PLUGIN_DATA was the last entry in section 7's normalised environment with no reader: the
// entry point exported it and six places in the runner wrote `.xeno/local/` as a literal instead
// (#205). These cover the resolution, and the first of them is the one that protects every
// repository that exists — unset has to mean exactly what the literals meant.

// TestUnsetResolvesToWhatTheLiteralsMeant is the compatibility criterion. Every repository has
// state under .xeno/local/ and no variable set, so this is the case that must not move.
func TestUnsetResolvesToWhatTheLiteralsMeant(t *testing.T) {
	t.Setenv(model.LocalDataEnv, "")
	root := "/srv/project"

	if got, want := model.LocalDir(root), filepath.Join(root, ".xeno/local"); got != want {
		t.Fatalf("unset resolves to %s, want %s", got, want)
	}
	for _, c := range []struct{ parts, want []string }{
		{[]string{"phase.env"}, []string{".xeno/local/phase.env"}},
		{[]string{"cost-ledger.yaml"}, []string{".xeno/local/cost-ledger.yaml"}},
		{[]string{"enforcement.yaml"}, []string{".xeno/local/enforcement.yaml"}},
		{[]string{"runs", "PROJ-1", "00-intake.lock"}, []string{".xeno/local/runs/PROJ-1/00-intake.lock"}},
	} {
		if got, want := model.LocalPath(root, c.parts...), filepath.Join(root, c.want[0]); got != want {
			t.Errorf("LocalPath(%v) is %s, want %s", c.parts, got, want)
		}
	}
}

// An absolute value is taken as given, because the entry point exports
// `${XENO_PLUGIN_DATA:=$root/.xeno/local}` and joining that to the root would nest one path
// inside another — which is the detail neither the issue nor section 7 contains.
func TestAnAbsoluteValueIsUsedAsItStands(t *testing.T) {
	t.Setenv(model.LocalDataEnv, filepath.FromSlash("/var/lib/xeno"))

	if got, want := model.LocalDir("/srv/project"), filepath.FromSlash("/var/lib/xeno"); got != want {
		t.Fatalf("an absolute value resolved to %s, want %s", got, want)
	}
	if got, want := model.LocalPath("/srv/project", "phase.env"), filepath.FromSlash("/var/lib/xeno/phase.env"); got != want {
		t.Fatalf("LocalPath under an absolute value is %s, want %s", got, want)
	}
}

// Relative to the repository root and not to the working directory: every other path the runner
// handles is root-relative, and a value meaning different things depending on where the command
// ran would be a worse promise than none.
func TestARelativeValueIsRelativeToTheRoot(t *testing.T) {
	t.Setenv(model.LocalDataEnv, "build/xeno-local")
	root := filepath.FromSlash("/srv/project")

	if got, want := model.LocalDir(root), filepath.Join(root, "build/xeno-local"); got != want {
		t.Fatalf("a relative value resolved to %s, want %s", got, want)
	}
}

// Empty reads as unset. os.Getenv cannot tell them apart, `export XENO_PLUGIN_DATA=` produces
// one, and resolving it to the root would put the ledger and the run marker at the top of the
// tree. Whitespace goes the same way, since a shell variable set from an empty substitution is
// as likely to be a space as nothing.
func TestAnEmptyOrBlankValueReadsAsUnset(t *testing.T) {
	root := filepath.FromSlash("/srv/project")
	want := filepath.Join(root, ".xeno/local")

	for _, v := range []string{"", "   ", "\t"} {
		t.Setenv(model.LocalDataEnv, v)
		if got := model.LocalDir(root); got != want {
			t.Errorf("value %q resolved to %s, want %s", v, got, want)
		}
	}
}

// The resolver creates nothing: the writers already make what they need, and a resolver that
// made directories would do it on every read, including the reads that only report a path.
func TestTheResolverCreatesNothing(t *testing.T) {
	root := t.TempDir()
	t.Setenv(model.LocalDataEnv, "")

	dir := model.LocalDir(root)
	if _, err := filepath.Rel(root, dir); err != nil {
		t.Fatalf("the default is not under the root: %v", err)
	}
	if entries, err := filepath.Glob(filepath.Join(root, "*")); err != nil || len(entries) != 0 {
		t.Fatalf("resolving created something: %v %v", entries, err)
	}
}
