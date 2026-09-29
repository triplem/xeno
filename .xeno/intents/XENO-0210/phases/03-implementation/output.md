---
intent: github.com/triplem/xeno#110
phase: 03-implementation
created: "2026-09-29T19:11:21Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+454cfbf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 75e97d40075ea82adaa13deb284e58fe31320425932cbcc3a495e81c2881cd65
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

`cmd/xeno/main.go`. `opts` gains `out` and `errw`; `run` takes them and `main` supplies
the real files and is now the only line in the package that names them. `parse` sets
them on the `opts` it builds. `report` and `suggest` become methods on `opts`, so the
thirty command functions that already take one need nothing threaded. Fifty-five writes
become writes to one of the two.

`verifyCode` is extracted from `cmdGateVerify`, holding the one branch a CI wrapper
depends on, with section 6's reason beside it.

`printInit` and `printEnforcement` take a writer, being the two printers with no `opts`
in reach.

No word, no format string and no exit code changed. Every hunk of the diff is a writer
or a signature.

`cmd/xeno/main_test.go`. Thirteen tests over `run` with a buffer each: the three exit
codes, a red verdict on standard output, the dispatch against the usage string, `init`
as one word, the positional argument, `--export` carrying no suggestion, `version`,
`gate verify` over this repository, and `verifyCode` over seven results.

**Two things the refactor found.** A loop in `suggest` binds `o` over the owed
obligations, so `o.out` inside it resolved to a field of a string and the package
stopped compiling; the variable is `owed` now. And `parse` built its `opts` without the
writers, which the first run caught as a nil pointer in `fmt.Fprintf` — the kind of
thing a package at zero coverage has no way to notice.

**One thing the tests found.** `version` is named in the usage and absent from the
dispatch table, because `run` answers it before the split and before the usage check, so
that it works in a directory holding nothing. The test exempts it by name and says why,
rather than relaxing to let any name miss.

Coverage over the module goes from 72.0% to 81.5%, and `cmd/xeno` from 28 of 30
functions at zero to 19 of 31.

<!-- xeno:section:deviations -->
## Deviations from the design

One, and it is the case the intent exists for.

P2 said the provisional exit code would be asserted through `gate verify` against a real
provisional phase, built by the commands. G-Evidence runs from P3 onward, so a
provisional verdict needs a P0, P1 and P2 finished green before it, each with an
`output.md` carrying every field G-Schema requires and two that no writer produces in a
repository with no project file. The first attempt wrote that fixture, did not reach
provisional, and skipped — a test that asserts nothing while reporting success.

So the branch is extracted instead. `verifyCode` takes a `VerifyResult` and returns the
code, and seven constructed results cover provisional alone, provisional beside green,
provisional beside red, a divergence, a red phase and nothing at all. The whole command
is still exercised over this repository, which has a hundred and thirty-eight verdicts
and no divergence, so the path from `run` to the exit code is covered and the branch is
covered exactly.

That is what XENO-0207 did with `tail` for the same reason, and the skip is what made it
obvious: a fixture four phases deep to reach one `if` is worse than the `if` being
reachable.
