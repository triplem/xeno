// SPDX-License-Identifier: Apache-2.0

package secrets

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// samples are the secret shapes this repository carries on purpose: the fixtures that assert
// the filter fires, and the records that describe it. gitleaks makes the same carve-out
// upstream for the AWS documentation key, with an allowlist section 4's file has no room for.
var samples = []string{
	"AKIAIOSFODNN7EXAMPLE",
	"ghp_aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	"xoxb-111111111111",
	"-----BEGIN OPENSSH PRIVATE KEY-----",
	"Authorization: Bearer AAAA",
	"AKIA|ASIA", // the shape quoted in prose about the filter
}

// The bar the shipped set has to clear, and the reason the import was a selection rather than
// a copy: a pattern that fires on prose teaches whoever reads a digest to distrust the
// redaction. Two upstream rules were left out because they failed this, and the filter's own
// header names them.
//
// The subject is this repository's whole text, which is the largest body of prose available
// that also discusses secret shapes constantly, so it is a harder corpus than most projects'.
func TestTheShippedFilterLeavesThisRepositoryAlone(t *testing.T) {
	// Two hundred patterns over a megabyte and a half of text takes about twenty seconds, which
	// is the cost of the corpus and not of the filter: redacting a summary is two hundred passes
	// over a few kilobytes and is not measurable. The bar still runs in CI, which does not pass
	// -short, and stays out of the way of a developer running the suite in a loop.
	if testing.Short() {
		t.Skip("the corpus bar is slow by construction; CI runs it")
	}
	f, err := Load("../..")
	if err != nil {
		t.Fatal(err)
	}
	if len(f.Patterns) < 100 {
		t.Fatalf("only %d patterns in force; the shipped set is not loading", len(f.Patterns))
	}

	var scanned, flagged int
	err = filepath.Walk("../..", func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		rel := filepath.ToSlash(p)
		switch {
		case strings.Contains(rel, "/.git/"), strings.Contains(rel, "/vendor/"):
			return nil
		case strings.HasSuffix(rel, "secrets.yaml"): // the filter is not its own subject
			return nil
		}
		switch filepath.Ext(rel) {
		case ".md", ".go", ".yaml", ".yml", ".json", ".sh", ".toml":
		default:
			return nil
		}
		b, readErr := os.ReadFile(p)
		if readErr != nil {
			return nil
		}
		scanned++
		// One pass per pattern over the whole file rather than per line: the same answer, and
		// two hundred patterns against every line of this repository takes twenty seconds.
		text := string(b)
		redacted := f.Redact(text)
		if redacted == text {
			return nil
		}
		was, now := strings.Split(text, "\n"), strings.Split(redacted, "\n")
		for i := range was {
			if i >= len(now) || was[i] == now[i] || carriesSample(was[i]) {
				continue
			}
			flagged++
			if flagged <= 10 {
				t.Errorf("%s:%d fires on ordinary content:\n  %s", rel, i+1, strings.TrimSpace(was[i]))
			}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if scanned < 200 {
		t.Fatalf("scanned only %d files; the corpus is not being read", scanned)
	}
	t.Logf("%d patterns over %d files, %d lines flagged", len(f.Patterns), scanned, flagged)
}

func carriesSample(line string) bool {
	for _, s := range samples {
		if strings.Contains(line, s) {
			return true
		}
	}
	return false
}
