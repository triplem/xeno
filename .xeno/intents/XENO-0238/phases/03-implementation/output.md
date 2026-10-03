---
intent: github.com/triplem/xeno#210
phase: 03-implementation
created: "2026-10-03T19:41:43Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4ac3d6f590093f49ae32fc5d266d17552e56588d2a6c6db827c0d1da7b676575
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`CLAUDE.md`'s Conventions paragraph is replaced. Where it said "Lines wrap at 88
characters, in code and in prose" it now says that Markdown prose wraps at 88 with tables
and code blocks exempt, that Go source has no width rule beyond `gofmt`, and that commit
messages and pull request descriptions wrap at 72 because GitHub reflows them to that
width, with `CONTRIBUTING.md` named as where the reason is written. The SPDX and dependency
sentences are unchanged.

The paragraph was replaced rather than edited into and read back as prose, which is this
project's rule for a second change to a passage.

`ASSUMPTIONS.md` gains A91, carrying the measurement and the deferral of a checker.

Nothing else is touched. No code, no field, no gate, no rule, and nothing in the tree is
rewrapped.

<!-- xeno:section:deviations -->
## Deviations from the design

**The design was wrong about the tree and the measurement corrects it.** Its impact section
says "no Go file is currently over 88 — the one that was, `internal/plugin/plugin.go`, was
rewrapped in the intent that found it". That is false, and the phase is sealed, so the
correction is here.

Measured at `9fcc639`: **1 772 Go lines exceed 88 characters, in 58 of 64 files**, the
longest at 217. `internal/plugin/plugin.go` was not an outlier that got fixed; it was one
of fifty-eight, and rewrapping it changed nothing about the rest. The Go half of the
sentence was never a rule that drifted — it was a sentence nothing had ever followed.

This strengthens the change rather than altering it. The approved wording already says Go
has no width rule beyond `gofmt`, and the measurement is why that is a description rather
than a concession.

**The Markdown half is in much better shape than expected.** 62 prose lines over 88 across
both normative documents, every one of them at 89 or 90 — a rule followed to within two
columns by hand, with nothing checking it. That asymmetry is the finding of this intent:
the same sentence produced near-perfect compliance in prose and none at all in code,
because prose is wrapped by whoever writes it and code is wrapped by a tool that does not
wrap.

**A91 therefore carries numbers the design did not have**, and says plainly that three
stated conventions now join `CLAUSE-READERS.md`'s list of rules whose reader is a person.

Nothing else deviates.
