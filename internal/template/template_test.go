// SPDX-License-Identifier: Apache-2.0

package template

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// repoRoot is this repository, which carries the shipped set. The tests use it rather
// than a fixture, so that the set itself is what is being checked.
const repoRoot = "../.."

// shippedOnly is a root holding this repository's plugin templates and no project ones, so
// that Load resolves to the shipped copy whatever .xeno/config/templates happens to carry.
// This repository now carries an override of its own, for #117's measurement, and the
// subject of the test below is the shipped set rather than what this project resolves to.
func shippedOnly(t *testing.T) string {
	t.Helper()
	abs, err := filepath.Abs(filepath.Join(repoRoot, pluginDir))
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, filepath.Dir(pluginDir)), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(abs, filepath.Join(root, pluginDir)); err != nil {
		t.Fatal(err)
	}
	return root
}

func TestEveryShippedTemplateHasBothBundles(t *testing.T) {
	root := shippedOnly(t)
	entries, err := os.ReadDir(filepath.Join(repoRoot, pluginDir))
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 6 {
		t.Fatalf("the set has %d templates, want the six phases", len(entries))
	}
	for _, e := range entries {
		for _, lang := range []string{"en", "de"} {
			r, err := Load(root, e.Name(), lang)
			if err != nil {
				t.Errorf("%s in %s: %v", e.Name(), lang, err)
				continue
			}
			if r.Source != FromPlugin {
				t.Errorf("%s resolved to %s, want the shipped one", e.Name(), r.Source)
			}
			if r.Bundle.Language != lang {
				t.Errorf("%s bundle says %q, want %q", e.Name(), r.Bundle.Language, lang)
			}
		}
	}
}

// The budget is three to four required sections per template, which is the lever this
// package is really about. A fifth means arguing another away, and a test is where that
// argument has to happen.
func TestTheRequiredSectionBudgetHolds(t *testing.T) {
	root := shippedOnly(t)
	entries, _ := os.ReadDir(filepath.Join(repoRoot, pluginDir))
	for _, e := range entries {
		r, err := Load(root, e.Name(), "en")
		if err != nil {
			t.Fatal(err)
		}
		n := 0
		for _, s := range r.Template.Sections {
			if s.Required {
				n++
			}
		}
		if n < 2 || n > 4 {
			t.Errorf("%s requires %d sections, the budget is three to four", e.Name(), n)
		}
	}
}

// WP2's first acceptance: the same artifact re-renders in another language without
// changing what it asserts. The headings move, the ids and the content do not.
func TestReRenderingInAnotherLanguageChangesOnlyTheHeadings(t *testing.T) {
	content := map[string]string{
		"decisions":    "We chose the boring one.",
		"alternatives": "The exciting one.",
		"impact":       "Two packages.",
	}
	en, err := Load(shippedOnly(t), "design", "en")
	if err != nil {
		t.Fatal(err)
	}
	de, err := Load(shippedOnly(t), "design", "de")
	if err != nil {
		t.Fatal(err)
	}
	if Parse(en.Render(content)) == nil {
		t.Fatal("nothing parsed back")
	}
	a, b := Parse(en.Render(content)), Parse(de.Render(content))
	if len(a) != len(b) {
		t.Fatalf("different sections: %v against %v", a, b)
	}
	for id, text := range a {
		if b[id] != text {
			t.Errorf("section %s differs between languages:\n  en %q\n  de %q", id, text, b[id])
		}
	}
	if !strings.Contains(de.Render(content), "## Auswirkungen") {
		t.Error("the German rendering does not use the German heading")
	}
	if strings.Contains(de.Render(content), "## Impact") {
		t.Error("the German rendering kept an English heading")
	}
}

// WP2's second acceptance. Falling back would render in a language nobody asked for and
// record that it was rendered in the one they did.
func TestAMissingBundleFailsRatherThanFallingBack(t *testing.T) {
	if _, err := Load(shippedOnly(t), "design", "fr"); err == nil {
		t.Fatal("a missing bundle fell back instead of failing")
	}
}

func TestAProjectTemplateBeatsTheShippedOneForThatIdAlone(t *testing.T) {
	root := t.TempDir()
	// Two shipped templates, one of them overridden.
	for _, id := range []string{"design", "intake"} {
		copyTemplate(t, filepath.Join(repoRoot, pluginDir, id), filepath.Join(root, pluginDir, id))
	}
	over := filepath.Join(root, projectDir, "design")
	if err := os.MkdirAll(over, 0o755); err != nil {
		t.Fatal(err)
	}
	write(t, filepath.Join(over, "template.yaml"),
		"id: design\nversion: 9.9.9\nphase: 02-design\nsections:\n  - { id: only-this, required: true }\n")
	write(t, filepath.Join(over, "strings.en.yaml"),
		"language: en\ntitle: Ours\nheadings:\n  only-this: Only this\n")

	r, err := Load(root, "design", "en")
	if err != nil {
		t.Fatal(err)
	}
	if r.Source != FromProject || r.Ref() != "design@9.9.9" {
		t.Fatalf("the override did not win: %s %s", r.Source, r.Ref())
	}
	other, err := Load(root, "intake", "en")
	if err != nil {
		t.Fatal(err)
	}
	if other.Source != FromPlugin {
		t.Fatal("overriding one template took the rest of the set with it")
	}
}

// An absent optional section is fine, a present empty one is not, so the renderer writes
// an optional section only where there is something in it. Required ones are rendered
// whether or not they carry anything, because a missing one is the finding.
func TestOptionalSectionsAppearOnlyWhenFilled(t *testing.T) {
	r, err := Load(shippedOnly(t), "intake", "en")
	if err != nil {
		t.Fatal(err)
	}
	body := r.Render(map[string]string{"problem": "x"})
	if strings.Contains(body, "open-questions") {
		t.Error("an empty optional section was rendered")
	}
	if !strings.Contains(body, AnchorPrefix+"scope -->") {
		t.Error("an empty required section was left out")
	}
	if len(r.Missing(map[string]string{"problem": "x"})) != 2 {
		t.Errorf("Missing did not report the two empty required sections: %v",
			r.Missing(map[string]string{"problem": "x"}))
	}
	body = r.Render(map[string]string{"problem": "x", "open-questions": "Q-1"})
	if !strings.Contains(body, AnchorPrefix+"open-questions -->") {
		t.Error("a filled optional section was left out")
	}
}

// The agent never writes an anchor, so a round trip has to survive one.
func TestRenderAndParseRoundTrip(t *testing.T) {
	r, _ := Load(shippedOnly(t), "verification", "en")
	content := map[string]string{
		"test-mapping": "AC-1 -> TestOne\nAC-2 -> TestTwo",
		"results":      "pass",
		"gaps":         "AC-3 is verified by hand.",
	}
	got := Parse(r.Render(content))
	for id, want := range content {
		if got[id] != want {
			t.Errorf("section %s came back as %q, want %q", id, got[id], want)
		}
	}
}

func copyTemplate(t *testing.T, from, to string) {
	t.Helper()
	if err := os.MkdirAll(to, 0o755); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(from)
	if err != nil {
		t.Fatal(err)
	}
	for _, e := range entries {
		b, err := os.ReadFile(filepath.Join(from, e.Name()))
		if err != nil {
			t.Fatal(err)
		}
		write(t, filepath.Join(to, e.Name()), string(b))
	}
}

func write(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
