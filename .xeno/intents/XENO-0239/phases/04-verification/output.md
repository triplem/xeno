---
intent: github.com/triplem/xeno#213
phase: 04-verification
created: "2026-10-03T20:07:36Z"
schema_version: "1.0"
runner_version: dev+b2b1f0f.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f041867672e6ec8881b2f256e1d2aea445703c85d9b7611365b9006e38228a77
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

The deliverable is a paragraph, so what can be checked is whether the claims it makes
about the trail are true. Four properties carry the resolution and each maps to one
measurement.

"Written when the phase closes" maps to `xeno phase finish`'s own usage line and to the
runner's single digest writer.

"Inside `artifacts_hash`" maps to `hashing.PhaseExcluded`, which names the two files the
hash leaves out.

"Outside `context_hash`" maps to section 5's sentence on what each hash covers.

"Read by the phases after it" maps to the process definition's "Digests instead of
re-derivation" paragraph, which is the claim the first draft got wrong.

The contradiction itself maps to a search for the clause in both documents: it should now
appear in the plan only.

The project's suite is run because the commit touches the repository.

<!-- xeno:section:results -->
## Results

**Written at finish.** `xeno phase finish --intent KEY --phase NN [--summary PATH|-]
writes digest.md`, from the command's own usage, and the runner has one writer for it.

**Inside `artifacts_hash`.** `hashing.PhaseExcluded` is `{"gate.yaml": true, "cost.yaml":
true}`, so `digest.md` is covered and cannot be altered after the verdict without a
divergence that `gate verify` reports.

**Outside `context_hash`.** Section 5 at line 474: "`context_hash` in an artifact's
frontmatter covers `context.lock.yaml` and says what the phase was produced from." The
digest is not in it, so recording in the digest leaves the lock's seal undisturbed.

**Read by the phases after it.** Section 6's "Digests instead of re-derivation": "Later
phases read the earlier phase's output and digest rather than searching the codebase
again. P3 works from the design, not from a fresh scan." This is the one the first draft
got wrong, which said P5.

**The contradiction is gone.** "outside the profile" now matches at two lines of
`docs/implementation-plan.md` and none of `docs/process-definition.md`. Section 5 is
byte-identical to what it was.

The suite: build succeeds, `go test ./...` ok, `go vet` silent, `gofmt -l` outside
`vendor/` empty, `./xeno gate verify` 313 verdicts at exit 0.

<!-- xeno:section:gaps -->
## Gaps

Nothing enforces the new clause, deliberately, and A92 records it. An agent that reads
outside the profile and says nothing in its digest leaves no trace, the budget is still
judged against `files`, and no gate can close that because no gate knows what was read.
What changed is that the plan now says so rather than implying a measurement.

So this replaces an unsatisfiable requirement with a satisfiable weak one. That is an
improvement in honesty rather than in capability, and `CLAUSE-READERS.md` gains another
clause whose reader is a person.

WP8's first half is untouched and still unmet: "a repeated phase reads only what changed".
The hashes are in `context.lock.yaml` and G-Freshness compares them, but nothing presents
the difference to a phase being redone. It gets its own issue.

The audit's blind spot is a gap in a document rather than in the code. `CLAUSE-READERS.md`
sorted clauses by their reader and never compared two clauses about the same artifact,
which is how a direct contradiction survived a pass designed to find exactly this class of
problem. The learning record carries it; whether the file gains a column is a question for
whoever reads it, and not a change this intent makes to a document it has already
amended once.
