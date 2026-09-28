---
intent: github.com/triplem/xeno#108
phase: 05-review
created: "2026-09-28T17:01:15Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+023193e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: e8e86631f6bb8c34852d3835844640b48dd3aaf24572415829fa42ea0687dcc3
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The specification was not touched.** Section 5 already required the field and Appendix
B already defined its value. Nothing under `docs/` is in the diff, which is what the
first standing rule requires of a change that makes the code agree with the documents.

**Nothing was invented.** No field, no gate, no tool, no dependency. `context_hash` is
in section 5's enumeration, `recomputed` was already checking it, and `hashing.FileHash`
already implemented Appendix B.

**The writer and the checker compute the same value the same way.** Both call
`hashing.FileHash`, which is the only reason a stranger recomputing the trail gets the
same answer as the runner that wrote it.

**Nothing already sealed was invalidated.** `gate verify` recomputes every verdict the
repository carries and matches. The intents sealed before this keep their hand written
values, which are the same values their locks produce.

**The comment says what the construction is and why.** The paragraph on `SectionSet`
explains the render-time write and the frozen lock, and neither sentence describes what
it replaced.

**One intent, one work package, one branch.** The issue carries `wp7`, the commits
reference
#108, and this is the first intent in the repository carried through all six phases.

**The record admits where it was assembled out of order.** P3's deviations section says
the code preceded the design phase. A checklist that passed over that would be worth
less than the finding.

<!-- xeno:section:release-notes -->
## Release notes

`xeno section set` now writes `context_hash` into the frontmatter it renders, computed
over the `context.lock.yaml` of the phase it is writing.

A phase produced the way the process describes — `xeno phase start`, then one `section
set` per required section — is no longer red on a missing `context_hash`. Filling that
field by hand is no longer necessary, and hand written values in existing artifacts stay
valid because the runner computes the same hash over the same file.

The value does not change while a phase is open. The lock is written once at `phase
start`, so every section write of one phase produces the same hash.

Where no lock lies beside the artifact, the field is left out and nothing is refused.
The missing lock is reported by G-Freshness, as it was before.

Unchanged: `model`, `tool`, `tool_version`, `secrets_hash` and `rules_hash` are still
written by whoever has them, and a phase produced by the tool alone is still reported
incomplete on those five.

<!-- xeno:section:residual-risk -->
## Residual risk

**A hash written on every render.** The value is recomputed each time a section is
written, and it is stable only because `context.lock.yaml` is frozen at `phase start`.
Anything that later refreshes the lock mid-phase would make the field move between two
writes of one phase and would be found, if at all, by a reader noticing that a sealed
artifact disagrees with its lock. The comment is the guard, and a comment is a weak
guard. A test that refreshed the lock and asserted the field followed it would be a real
one, and is not written.

**The field is bound to the lock and to nothing else.** It says which file the phase's
information base was recorded in, not that the base was right. A10 still names the
blocks the lock does not carry, so a correct `context_hash` over an incomplete lock is
exactly as trustworthy as the lock, which is less than a reader may assume of a green
G-Schema.

**`digest.md` is still written by hand.** Half of what section 5 asks for is produced by
the runner and half by whoever writes the digest, which is the asymmetry this fix
reduced rather than removed.

**The evidence of this intent is a local run.** Nothing in `attached.yaml` distinguishes
it from a pipeline artifact, and the only place that says so is prose in two sections
and a comment in a manifest that does not travel with the repository.

**Accepted, with the residual risk named rather than mitigated.** None of the four
blocks the change; each is smaller than the state it replaces, which was a red phase on
every artifact the tool produced.
