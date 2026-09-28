---
intent: github.com/triplem/xeno#111
phase: 03-implementation
created: "2026-09-28T17:35:35Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0e77cc2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 3a18fc03cf47ba42d1f1732bc00b0f6a6c86dfe2a0ebc6a2fa0371c55968d503
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

`cmd/xeno/main.go`. The paragraph about what `init` did moves above `printInit`, with a
blank line before it and the closing brace of `printGate` above that. `printEnforcement`
keeps its own comment and nothing else. `gofmt` decided the blank line: without it the
comment lands against the previous function's brace and the file is no longer formatted,
which is a check the repository already runs.

`internal/gates/gates.go`. `undeclaredEvidence` takes the shape P2 decided.

`insideEvidence` is new, five lines, and turns a declared path into the key this check
compares: the path relative to `evidence/`, or nothing. It answers the empty path, the
path that is exactly `evidence/`, a path pointing elsewhere in the phase directory, and
one trying to climb out with `../`, all through the same test rather than through four
special cases. Its comment records what the old key did with an empty path, because an
inert entry in a map is the kind of thing a reader restores while tidying.

The directory read becomes `filepath.Walk` from `evidence/`, skipping directories, keyed
by `filepath.Rel` of each file with `filepath.ToSlash` over it. `filepath.Walk` and not
`WalkDir`, for the reason P2 gives: this package calls its findings variable `fs`.

The existence check is now a `Stat` for a directory rather than a `ReadDir` that failed.
Same outcome, one syscall, and it says what it is asking.

`internal/gates/schema_test.go`. Four tests and three helpers. `evidenceFile` writes a
file at a path relative to `evidence/`, so a test can nest one; `declare` appends an
evidence block to the corpus `output.md`; `undeclared` returns only the findings of this
check, so a hand written declaration cannot pass a test by producing some other finding.

Three of the four fail against the tree before the change. The fourth, the pending
declaration, passes there too, and its comment says why: the entry keyed `.` was inert
rather than wrong, so the test guards against its return instead of proving its removal.

<!-- xeno:section:deviations -->
## Deviations from the design

None. The writer, the key, the walk, the skipped directories and the choice of `Walk`
over `WalkDir` are as P2 decided them, and the two rejected alternatives stayed
rejected.

One thing P2 did not mention and the code needed: `filepath.Walk` from a root that does
not exist reports an error through the callback rather than before it, so the existence
of `evidence/` is checked first. P2 said `Stat` for the root and the reason is the same
either way, so this is a detail of the same decision rather than a deviation from it.

The order of the record is the order of the work this time. P0 to P2 were written before
any code, and the code was written inside P3. The deviation this section carried for
XENO-0108 does not recur.
