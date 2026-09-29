---
intent: github.com/triplem/xeno#134
phase: 03-implementation
created: "2026-09-29T17:59:30Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c05acae.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2ec969e5d3f2810492c214a76d8f637bdf50c022e28403642b78788caa4d59ed
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: by-hand
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`cmd/xeno`. Two format strings, `listRow` and `phaseRow`, each shared by its heading and
its rows, which is what makes the alignment hold by construction. `--all`, a boolean
beside the three that already exist. `tail`, which says how many rows a listing shows
and how many it leaves out. A notice after the table naming the count and the flag,
printed only when something was left out. The heading printed with the first row rather
than before the loop, so a repository holding no intents prints nothing.

The one-intent form gets the same treatment, and its heading is the one worth most:
`state` and `verdict` side by side, a position and a judgement.

`cmd/xeno/list_test.go`, the package's first test file. Three tests: the listing's
heading over the widest row each column can take, the same for the phase form, and
`tail` over nothing, three, ten, eleven and sixty-three intents, with and without
`--all`.

The alignment tests compare a heading against a row built from the same format string,
so they assert the property rather than the string. `tail` is separated from the
printing for the same reason: the arithmetic can be checked without capturing output.

`internal/runner` is untouched, which is AC8 by construction rather than by assertion.

**What it prints now.** Ten rows under a heading, the newest last, and `53 older, --all
to see them`.

<!-- xeno:section:deviations -->
## Deviations from the design

None in the design. The shared format strings, the tail in the command, the notice after
the table and the heading with the first row are as P2 decided.

Two things the design did not reach, both small and both from writing the tests.

`tail` exists because the arithmetic had to be testable. P2 said the tail is taken in
the command and left it inline; a test would then have had to capture standard output to
check that eleven intents show ten. The function is three lines and its absence would
have cost a worse test.

`sprintRow` exists for the same reason, so that a test can build a heading and a row
from one format string and compare them. It is a wrapper around `fmt.Sprintf` and it
earns its place by making the alignment assertable at all.

This is `cmd/xeno`'s first test file. The package had none, which #110 is about, and
three tests do not close that.
