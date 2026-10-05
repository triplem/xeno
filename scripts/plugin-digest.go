// SPDX-License-Identifier: Apache-2.0

//go:build ignore

// plugin-digest prints the digest G-Supply compares, computed by the same code that
// compares it, so the release and the check cannot disagree about the definition.
//
//	go run scripts/plugin-digest.go [--require-shipped]
//
// The release passes it to the build through -ldflags, which is the anchor section 13
// describes: "Nothing in the repository states what the expected value is."
//
// With --require-shipped the digest is printed only if the tree the binaries will carry,
// at plugin.ShippedDir, exists and hashes equal to the vendored one. The release passes
// the flag, because the digest is taken over one path and the binaries carry the other
// and until #201 nothing compared them: a copy that succeeded while producing a different
// tree would ship a digest describing a tree nobody has, and every phase of every project
// installing that release would fail G-Supply with two hex strings and nothing to act on.
//
// Without the flag nothing is verified, and it says so on standard error rather than
// letting a person assume otherwise. That is the default because a development tree
// carries no shipped copy by design — internal/plugin/embedded.go says why — so a script
// that always required one would fail in the tree where it is most often run by hand.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/triplem/xeno/internal/plugin"
)

func main() {
	requireShipped := flag.Bool("require-shipped", false,
		"refuse unless the tree at "+plugin.ShippedDir+" is the one the digest describes")
	flag.Parse()

	h := plugin.Hash(".")
	if h == "" {
		fmt.Fprintln(os.Stderr, "no vendored plugin under "+plugin.Dir)
		os.Exit(1)
	}

	if !*requireShipped {
		fmt.Fprintln(os.Stderr, "not verified against "+plugin.ShippedDir+
			": pass --require-shipped in a release, where the binaries carry that tree")
		fmt.Println(h)
		return
	}

	shipped := plugin.HashTree(plugin.ShippedDir)
	if shipped == "" {
		fmt.Fprintln(os.Stderr, "no tree at "+plugin.ShippedDir+
			": the release copies "+plugin.Dir+" there before taking the digest")
		os.Exit(1)
	}
	if shipped != h {
		fmt.Fprintln(os.Stderr, "the tree the binaries would carry is not the one this digest describes:")
		fmt.Fprintln(os.Stderr, "  "+plugin.Dir+" is "+h)
		fmt.Fprintln(os.Stderr, "  "+plugin.ShippedDir+" is "+shipped)
		for _, d := range plugin.Differences(plugin.Dir, plugin.ShippedDir) {
			fmt.Fprintln(os.Stderr, "  "+d)
		}
		os.Exit(1)
	}
	fmt.Println(h)
}
