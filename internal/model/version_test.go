// SPDX-License-Identifier: Apache-2.0

package model

import (
	"regexp"
	"runtime/debug"
	"testing"
)

// The stamp is tested as a function of its inputs. Building a binary and reading what
// came out would test the toolchain rather than this.
func TestStamp(t *testing.T) {
	cases := []struct {
		name     string
		revision string
		modified bool
		want     string
	}{
		{"no revision leaves the version alone", "", false, "0.1.0-dev"},
		{"no revision, and modified says nothing either", "", true, "0.1.0-dev"},
		{"a clean tree names its commit", "356db60abcdef0123456789", false, "0.1.0-dev+356db60"},
		{"a modified tree says so", "356db60abcdef0123456789", true, "0.1.0-dev+356db60.dirty"},
		{"a short revision is not padded", "abc", false, "0.1.0-dev+abc"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := stamp("0.1.0-dev", c.revision, c.modified); got != c.want {
				t.Fatalf("stamp(%q, %v) = %q, want %q", c.revision, c.modified, got, c.want)
			}
		})
	}
}

// Build metadata is "+" followed by dot separated alphanumerics and hyphens. A version
// that is not parseable is worse than the constant it replaced.
func TestStampIsValidSemverBuildMetadata(t *testing.T) {
	valid := regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+-[0-9A-Za-z-]+\+[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*$`)
	for _, v := range []string{
		stamp("0.1.0-dev", "356db60abcdef", false),
		stamp("0.1.0-dev", "356db60abcdef", true),
	} {
		if !valid.MatchString(v) {
			t.Fatalf("%q is not a version with valid build metadata", v)
		}
	}
}

// The design rests on this: a test binary carries no vcs settings, so the placeholder
// survives here and every fixture that names it keeps working. If the toolchain ever
// starts stamping test binaries, the fixtures move and this says why.
func TestATestBinaryCarriesNoVCSSettings(t *testing.T) {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		t.Skip("no build info in this binary")
	}
	for _, s := range info.Settings {
		if s.Key == "vcs.revision" || s.Key == "vcs.modified" {
			t.Fatalf("a test binary now carries %s=%s, so RunnerVersion is stamped in tests too", s.Key, s.Value)
		}
	}
	if RunnerVersion != devVersion {
		t.Fatalf("RunnerVersion is %q in a test binary, want the unstamped %q", RunnerVersion, devVersion)
	}
}
