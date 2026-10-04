// SPDX-License-Identifier: Apache-2.0

package model

import (
	"path/filepath"
	"strings"
)

// MatchPath answers whether a repository relative path matches a scope pattern of the
// form section 12 writes: `src/payment/**`, `**/testdata/**`, `docs/adr/*.md`.
//
// Written here rather than taken from a library because the project has one dependency
// and adding a second is a decision rather than a step. What is needed is small: `**`
// matches any number of path segments including none, and within a segment the rules are
// filepath.Match's, which is what a reader of the pattern expects.
func MatchPath(pattern, path string) bool {
	return matchSegments(strings.Split(pattern, "/"), strings.Split(path, "/"))
}

func matchSegments(pattern, path []string) bool {
	for len(pattern) > 0 {
		if pattern[0] == "**" {
			// Zero segments, then one, then two: the first match wins, and a trailing
			// `**` therefore matches the rest of the path however deep it goes.
			rest := pattern[1:]
			if len(rest) == 0 {
				return true
			}
			for i := 0; i <= len(path); i++ {
				if matchSegments(rest, path[i:]) {
					return true
				}
			}
			return false
		}
		if len(path) == 0 {
			return false
		}
		ok, err := filepath.Match(pattern[0], path[0])
		if err != nil || !ok {
			return false
		}
		pattern, path = pattern[1:], path[1:]
	}
	return len(path) == 0
}
