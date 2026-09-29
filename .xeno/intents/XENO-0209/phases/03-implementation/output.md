---
intent: github.com/triplem/xeno#109
phase: 03-implementation
created: "2026-09-29T18:39:42Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+3ec2429.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 53e8d7f3cf7df27115b1164816c43a4fea95089766edc4cb6f85a05773c18c59
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

`internal/model`. `KnownIntentFiles`, the four names section 4 lists at the intent
level, and `PhasesDir`, the one directory it lists beside them. Two maps rather than
one, and the comment says why: a single map would accept `output.md` in an intent
directory and `assumptions.yaml` in a phase, so neither check would distinguish the
levels section 4 distinguishes.

`internal/gates`. `intentDirectoryFindings`, the twin of `directoryFindings` one level
up, with the same shape and the same two wordings for a file and a directory. Its
comment carries the reason Appendix B gives and the reason a directory is reported
although `DirHash` cannot descend into one: a directory nobody wrote is where files
appear next.

`CompleteOnClose` calls it first, before the reason and the learning record, so that a
verdict lists what the directory contained before what the intent asserted. The hash
that verdict carries covers that directory.

`internal/gates/carryforward_test.go`. Four tests: a stray file, a stray directory with
its own wording, the files section 4 names staying silent including the `gate.yaml` a
second close would find, and that the two levels do not share their lists.

Nothing else. `hashing.DirHash`, `IntentExcluded`, the normalisation and the phase level
check are untouched, because the gap was never in the hash: it was that nothing said
what the hash was allowed to cover.

<!-- xeno:section:deviations -->
## Deviations from the design

None in the design. The map, the twin function, its placement first in
`CompleteOnClose`, the reported directory and the untouched hash are as P2 decided.

One thing about the verification rather than the change, and it is the second time. The
new tests reference `model.KnownIntentFiles` and `model.PhasesDir`, so they do not
compile against the tree without the change and the before-and-after cannot be run as a
test. XENO-0201 hit this when `Finish` gained a parameter and recorded it; here it is a
new symbol rather than a changed signature, and the substitute is better: two binaries,
one built from `main` and one from this branch, run against the same scratch intent with
a stray file in it. That is a real before and after at the command level and it is in
the verification.

The order of the work was right. P0 to P2 before any code.
