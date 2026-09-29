---
intent: github.com/triplem/xeno#120
phase: 05-review
created: "2026-09-29T06:05:05Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ad0a764.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: ebf8354b5fd84dd1b92b18350f4d0e41ca4ab06eb64a67df8813b96955eb281c
context_hash: e7018e573be1deed9893015cc2f771810399d8148358d9b374a4a9467957e0b4
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

**Nothing under `docs/` changed.** Section 4 defines the two files and the additive
rule, section 16 says the runner filters and why, and Appendix B delegates the
computation. All three described a runner that did not do it.

**Nothing was invented.** The filter is the file section 4 specifies, `secrets_hash` is
a field section 5 enumerates, and the one decision the specification declined to take is
recorded as A62 rather than made silently.

**The one thing the specification asked for is what was built.** Section 16 says the
filtering is deterministic and outside the model's reach so that a model cannot route
around it. The probe in the verification is that sentence on disk: a summary went in
with two fake secrets and the digest came out with two named redactions.

**The runner writes nothing it cannot source.** An empty filter yields no `secrets_hash`
and no redaction, and `by-hand` is not written in its place, because that value says a
person stood behind it.

**A project can only widen the filter.** The union is sorted and matching is a
disjunction, so a project reusing a shipped id leaves the shipped pattern in force.
Section 4's sentence, tested.

**Nothing fails on bad configuration.** A missing filter finishes a phase and an
uncompilable regex drops one pattern and stays in the hash, so what was in force is
still reconstructable.

**The honesty rule was not tightened, and the number is in the row.** 88 of 90 verdicts
diverge, measured twice, and A62 carries it with the reason: `secrets_hash` sits inside
`artifacts_hash`, so correcting sixty sealed artifacts would rewrite history and the
merge commits naming it.

**`output.md` is not redacted, and this record is why.** Its acceptance criteria quote a
field name with its value, and its shipped filter is tested against a sentence from its
own prose. A filter over the agent's writing would have mangled the explanation of the
filter.

**Every verdict still matches**, 91 of them, and the phases were written in their order.

<!-- xeno:section:release-notes -->
## Release notes

Digests are filtered. `xeno phase finish --summary` now redacts the summary against the
effective secret filter before writing `digest.md`, and each match is replaced by the id
of the pattern that caught it, so a reader sees what was removed and by which rule while
the sentence around it stays.

The filter ships as `.xeno/plugin/secrets.yaml`: five patterns, for AWS access keys,
bearer tokens, GitHub tokens, private key blocks and Slack tokens, and five never
digested paths. A project adds to it in `.xeno/config/secrets.yaml` and can remove
nothing from it, because a filter a project can switch off is not a filter.

`secrets_hash` is written into `digest.md` and `output.md`: a sha256 over the effective
set in a canonical rendering, so it says which filter an artifact was written under. It
is independent of comments, of the order patterns appear in and of which of the two
files they came from. A62 records the computation, which Appendix B leaves to whichever
package writes the field first.

`output.md` carries the field and is not itself redacted. The filtering belongs to the
digest, which is the file that leaves a session.

A repository without a vendored plugin has no filter: nothing is redacted, no
`secrets_hash` is written, and a phase finishes. A project pattern that will not compile
is dropped and still counted in the hash.

Of the five fields A35 left to a hand, two remain: `tool_version` and `rules_hash`.

<!-- xeno:section:residual-risk -->
## Residual risk

**Five patterns are the security surface.** A secret whose shape nobody anticipated
reaches a digest as it always did. What the change buys is not that digests are clean
but that each one says which filter it passed, so somebody can later tell which digests
predate a pattern. Saying more than that would be the overstatement this intent has
spent three phases avoiding.

**Nothing checks the other artifacts.** G-Secret is still `not-implemented`, so a secret
pasted into a requirements section sits in the repository with nothing to notice it. The
digest is the file that leaves a session and it is now filtered; the files that stay are
not. That gate can now read the same package, which is the whole reason the package is
separate, and it is the next piece of work.

**`by-hand` is still accepted where a writer exists**, by decision and with the number
recorded. It matters only where somebody writes it deliberately, which is precisely the
case #120 is about, so the loophole and the issue point at each other.

**`paths_never_digested` is inert.** Shipped, hashed, and consulted by nothing, because
no writer reads a file into a digest. Honest and doing nothing, and a reader of the
filter will reasonably assume otherwise.

**The redaction is visible and therefore informative.** `[redacted: aws-access-key]`
tells a reader which rule fired, and it tells anybody reading the digest that something
matching that shape was there. That is the right trade for a file people read, and it is
a trade.

**`tool_version` still has no source**, so #120's enforcement question is unchanged: an
agent cannot produce a complete artifact with the tool alone, and a rule forbidding the
hand would still be unkeepable.

**Accepted with the six named.** None blocks the change, and the state it replaces was a
digest the runner's name was on that had passed through nothing.
