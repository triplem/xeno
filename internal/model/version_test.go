// SPDX-License-Identifier: Apache-2.0

package model

import (
	"regexp"
	"runtime/debug"
	"strings"
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
		{"no revision leaves the version alone", "", false, "dev"},
		{"no revision, and modified says nothing either", "", true, "dev"},
		{"a clean tree names its commit", "356db60abcdef0123456789", false, "dev+356db60"},
		{"a modified tree says so", "356db60abcdef0123456789", true, "dev+356db60.dirty"},
		{"a short revision is not padded", "abc", false, "dev+abc"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := stamp(devVersion, c.revision, c.modified); got != c.want {
				t.Fatalf("stamp(%q, %v) = %q, want %q", c.revision, c.modified, got, c.want)
			}
		})
	}
}

// Build metadata is "+" followed by dot separated alphanumerics and hyphens, and that part
// is still well formed. The part before it deliberately is not a semver version: a
// development build cannot learn the newest tag, so it claims no number rather than the
// wrong one, which is what `0.1.0-dev` was for twenty-nine minor releases (#177).
func TestTheStampClaimsNoVersionAndItsMetadataIsWellFormed(t *testing.T) {
	metadata := regexp.MustCompile(`^[0-9A-Za-z-]+(\.[0-9A-Za-z-]+)*$`)
	semver := regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+`)
	for _, v := range []string{
		stamp(devVersion, "356db60abcdef", false),
		stamp(devVersion, "356db60abcdef", true),
	} {
		base, meta, found := strings.Cut(v, "+")
		if !found {
			t.Fatalf("%q carries no build metadata", v)
		}
		if !metadata.MatchString(meta) {
			t.Errorf("%q is not valid build metadata", meta)
		}
		if semver.MatchString(base) {
			t.Errorf("%q names a version, which a development build cannot know", base)
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
