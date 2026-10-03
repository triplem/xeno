// SPDX-License-Identifier: Apache-2.0

//go:build ignore

// plugin-digest prints the digest G-Supply compares, computed by the same code that
// compares it, so the release and the check cannot disagree about the definition.
//
//	go run scripts/plugin-digest.go
//
// The release passes it to the build through -ldflags, which is the anchor section 13
// describes: "Nothing in the repository states what the expected value is."
package main

import (
	"fmt"
	"os"

	"github.com/triplem/xeno/internal/plugin"
)

func main() {
	h := plugin.Hash(".")
	if h == "" {
		fmt.Fprintln(os.Stderr, "no vendored plugin under "+plugin.Dir)
		os.Exit(1)
	}
	fmt.Println(h)
}
