---
intent: github.com/triplem/xeno#215
phase: 02-design
created: "2026-10-03T20:23:09Z"
schema_version: "1.0"
runner_version: dev+b626f1a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7abd887b3a02d4e2eec1a9d7e1a0be5a824952a49223e5fe5efbb02764faa9aa
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**A refusal, not a mechanism.** The act is forbidden by section 11 and was detected
afterwards. What was missing is a reader at the point of the act, and that is four lines.

**It keys on `gate.yaml`, not on the lock.** A phase with a lock and no verdict is the
dead-run recovery `Start`'s existing refusal already points at — "if that run died, remove
the marker" — and keying on the lock would have closed it. A verdict is also the right
test on the merits: it is what makes the directory sealed.

**The refusal names both ways forward.** `section set` with `phase finish` to redo the
work, and removing the verdict to start over. The second is the deliberate escape, and it
is named rather than left to be discovered, because a refusal that only says no is worked
around by deleting whatever is in the way — which is how a sealed artifact gets rewritten
in the first place.

**Nothing is written before the refusal.** It sits with the other preconditions, above
`common()` and the lock write, and the test asserts the lock's hash is unchanged after the
refused call.

**The duplicate derivation is deleted, not finished.** `ChangedSince` already answers what
changed and its comment says why it is derived rather than recorded: "A derivation that is
printed cannot drift from its inputs; one that is recorded can." Two derivations of one
answer can disagree, which is worse than the gap the second was written for.

<!-- xeno:section:alternatives -->
## Alternatives

**Let `phase start` overwrite and rely on `gate verify`.** The status quo. Rejected: the
guard is the last line and it reported the divergence after a sealed artifact had already
been rewritten, twice in one day, with `git checkout --` as the recovery. Section 11's clause
deserves a reader where the act happens.

**Make `phase start` idempotent — detect the existing lock and leave it.** Rejected. It
would make a forbidden act look supported, and the phase would then be "started" with a lock
describing a different tree, which is the one thing the lock must not do.

**Refuse on the existence of the lock rather than the verdict.** Rejected, and this is the
one that would have shipped a regression: it closes the dead-run recovery `Start` itself
documents.

**Write the changed set into the lock so a restart does not lose it.** Rejected twice over.
It needs a field and so a specification change, and the answer is derivable from the
predecessor's lock already, which is what `ChangedSince` does.

**Finish the second derivation that was started here.** Rejected as soon as the first was
found. Section 5's field list has no entry for a changed set and the existing derivation is
printed rather than recorded for that reason; a second one in the suggestion path would have
printed the same answer in two places and could drift from it.

<!-- xeno:section:impact -->
## Impact

Four lines of refusal in `internal/runner/runner.go` and two tests. No field, no gate, no
rule, no document.

One behaviour changes: `xeno phase start` on a phase that has a verdict now exits non-zero
with a message instead of rewriting the lock. Nothing in this repository depends on the old
behaviour, and what did almost depend on it — the dead-run recovery — is why the refusal
keys on the verdict.

Section 11 gains a reader at the point of the act. `CLAUSE-READERS.md` lists the
modification half of "what is sealed is never rewritten" with `gate verify` as its reader;
it now has two, one before and one after, which is the shape the deletion half got in #193.

What does not change: `gate verify` still reports a divergence for every other way a sealed
artifact can be rewritten, and there are others — a hand edit is the common one, and this
session made two.
