// Package hashing implements Appendix B of the process definition to the byte.
// Both values are recomputed by whoever verifies the trail, so nothing here may
// depend on the machine, the clock or the file system beyond file content and path.
package hashing

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// Excluded names per scope. A phase excludes gate.yaml and cost.yaml, an intent
// excludes gate.yaml only.
var (
	PhaseExcluded  = map[string]bool{"gate.yaml": true, "cost.yaml": true}
	IntentExcluded = map[string]bool{"gate.yaml": true}
)

// Normalise replaces every CRLF with LF and changes nothing else.
func Normalise(b []byte) []byte {
	return bytes.ReplaceAll(b, []byte("\r\n"), []byte("\n"))
}

// Hex returns the lowercase hex sha256 of b.
func Hex(b []byte) string {
	s := sha256.Sum256(b)
	return hex.EncodeToString(s[:])
}

// FileHash is the sha256 of a file's normalised content, used for evidence items.
func FileHash(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return Hex(Normalise(b)), nil
}

// DirHash computes artifacts_hash over the files lying directly in dirRel,
// relative to repoRoot, without descending, excluding the given names.
func DirHash(repoRoot, dirRel string, excluded map[string]bool) (string, error) {
	entries, err := os.ReadDir(filepath.Join(repoRoot, dirRel))
	if err != nil {
		return "", err
	}
	type line struct{ path, sum string }
	var lines []line
	for _, e := range entries {
		if e.IsDir() || excluded[e.Name()] {
			continue
		}
		rel := filepath.ToSlash(filepath.Join(dirRel, e.Name()))
		b, err := os.ReadFile(filepath.Join(repoRoot, rel))
		if err != nil {
			return "", err
		}
		lines = append(lines, line{rel, Hex(Normalise(b))})
	}
	// Byte order on the path, which is what LC_ALL=C sort produces.
	sort.Slice(lines, func(i, j int) bool { return lines[i].path < lines[j].path })
	var stream strings.Builder
	for _, l := range lines {
		fmt.Fprintf(&stream, "%s  %s\n", l.sum, l.path)
	}
	return Hex([]byte(stream.String())), nil
}

// FindingID is F- plus the first six hex characters of the sha256 over gate, rule,
// path and cause, each terminated with \n, the empty rule id included.
func FindingID(gate, rule, path, cause string) string {
	return "F-" + Hex([]byte(gate + "\n" + rule + "\n" + path + "\n" + cause + "\n"))[:6]
}
