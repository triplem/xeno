package hashing

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func write(t *testing.T, root, rel, content string) {
	t.Helper()
	p := filepath.Join(root, rel)
	if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(p, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

const dir = ".xeno/intents/PROJ-1/phases/01-requirements"

func TestExclusionsAndSubdirectories(t *testing.T) {
	root := t.TempDir()
	write(t, root, dir+"/output.md", "a\n")
	base, _ := DirHash(root, dir, PhaseExcluded)

	write(t, root, dir+"/gate.yaml", "x")
	write(t, root, dir+"/cost.yaml", "y")
	write(t, root, dir+"/evidence/attached.yaml", "z")
	after, _ := DirHash(root, dir, PhaseExcluded)
	if base != after {
		t.Fatal("gate.yaml, cost.yaml or a subdirectory entered the hash")
	}
	write(t, root, dir+"/digest.md", "b\n")
	if h, _ := DirHash(root, dir, PhaseExcluded); h == base {
		t.Fatal("a covered file did not change the hash")
	}
}

func TestCRLFIsNormalised(t *testing.T) {
	a, b := t.TempDir(), t.TempDir()
	write(t, a, dir+"/output.md", "line\nline\n")
	write(t, b, dir+"/output.md", "line\r\nline\r\n")
	ha, _ := DirHash(a, dir, PhaseExcluded)
	hb, _ := DirHash(b, dir, PhaseExcluded)
	if ha != hb {
		t.Fatal("CRLF and LF produced different hashes")
	}
}

func TestPathIsPartOfTheHash(t *testing.T) {
	root := t.TempDir()
	write(t, root, dir+"/output.md", "same\n")
	other := ".xeno/intents/PROJ-2/phases/01-requirements"
	write(t, root, other+"/output.md", "same\n")
	h1, _ := DirHash(root, dir, PhaseExcluded)
	h2, _ := DirHash(root, other, PhaseExcluded)
	if h1 == h2 {
		t.Fatal("moving an intent must change its hash")
	}
}

// The appendix says the stream can be reproduced with standard tools. This test does
// exactly that, so a reader who distrusts the runner can trust the definition.
func TestReproducibleWithStandardTools(t *testing.T) {
	if _, err := exec.LookPath("sha256sum"); err != nil {
		t.Skip("sha256sum not available")
	}
	root := t.TempDir()
	write(t, root, dir+"/output.md", "b\n")
	write(t, root, dir+"/context.lock.yaml", "a\n")
	write(t, root, dir+"/digest.md", "c\n")
	script := `cd "$1" && for f in ` + dir + `/*; do [ -f "$f" ] || continue; ` +
		`case "$(basename "$f")" in gate.yaml|cost.yaml) continue;; esac; ` +
		`sed 's/\r$//' "$f" | sha256sum | sed "s|-\$|$f|"; done | LC_ALL=C sort -k2 | sha256sum | cut -d' ' -f1`
	out, err := exec.Command("sh", "-c", script, "sh", root).Output()
	if err != nil {
		t.Fatal(err)
	}
	got, _ := DirHash(root, dir, PhaseExcluded)
	if strings.TrimSpace(string(out)) != got {
		t.Fatalf("standard tools %s, runner %s", strings.TrimSpace(string(out)), got)
	}
}

func TestFindingIDKeepsEmptyRuleTerminator(t *testing.T) {
	if FindingID("G-Schema", "", "p", "c") == FindingID("G-Schema", "p", "", "c") {
		t.Fatal("field boundaries are not preserved")
	}
	if id := FindingID("G", "", "p", "c"); len(id) != 8 || id[:2] != "F-" {
		t.Fatalf("unexpected id form %q", id)
	}
}
