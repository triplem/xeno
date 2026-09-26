// SPDX-License-Identifier: Apache-2.0

package model

import "testing"

// The patterns are the ones section 12 writes, and `**` is the part filepath.Match does
// not have: it matches any number of segments, including none.
func TestMatchPath(t *testing.T) {
	for _, tc := range []struct {
		pattern, path string
		want          bool
	}{
		{"src/payment/**", "src/payment/card.go", true},
		{"src/payment/**", "src/payment/iso/4217.go", true},
		{"src/payment/**", "src/payment", true},
		{"src/payment/**", "src/shipping/box.go", false},
		{"**/testdata/**", "internal/gates/testdata/phase/output.md", true},
		{"**/testdata/**", "testdata/x", true},
		{"**/testdata/**", "internal/gates/gates.go", false},
		{"docs/adr/*.md", "docs/adr/0012-payments.md", true},
		{"docs/adr/*.md", "docs/adr/old/0001.md", false},
		{"**", "anything/at/all", true},
		{"*.go", "main.go", true},
		{"*.go", "cmd/main.go", false},
	} {
		if got := MatchPath(tc.pattern, tc.path); got != tc.want {
			t.Errorf("MatchPath(%q, %q) = %v, want %v", tc.pattern, tc.path, got, tc.want)
		}
	}
}
