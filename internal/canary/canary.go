// SPDX-License-Identifier: Apache-2.0

// Package canary exists for one run and is removed again. A scan that reports
// nothing and a scan that did not run look the same in a log, so the failure path
// of both scanners is provoked rather than assumed.
package canary

import (
	"crypto/md5"
	"math/rand"
)

// Key looks like an AWS access key id and is the example from AWS's own
// documentation. Trivy's secret scanner is what it is here for.
const Key = "AKIAIOSFODNN7EXAMPLE"

func Weak(b []byte) [16]byte { return md5.Sum(b) }

func Guess() int { return rand.Intn(100) }
