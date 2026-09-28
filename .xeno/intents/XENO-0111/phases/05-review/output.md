---
intent: github.com/triplem/xeno#111
phase: 05-review
created: "2026-09-28T17:39:07Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0e77cc2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 1db87047416e79f3a963434590d1a971324672a2f78f2b6bdffcb013922661ba
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

**Nothing under `docs/` changed.** Section 4 already permitted the nesting the code
failed to survive, and one rejected alternative would have removed that permission in a
code change. It was rejected on that ground and the design says so.

**Nothing was invented.** No field, no gate, no finding, no dependency. `filepath` was
already imported; the walk replaces a read.

**The fix does not widen what is reported.** A directory of declared files is no
finding, which is a criterion rather than an observation, and the test for it fails
against the tree before the change. Widening what a check reads is where a new report
slips in.

**Every verdict in the repository still matches.** `gate verify` recomputes 61 of them.
`evidence/` as `Attach` writes it, flat, produces what it produced before.

**Both halves are one reading, and the record keeps them together.** The intake says why
they are one intent, which is the only thing about this pair worth remembering.

**The phases were written in their order.** P0 to P2 before any code, the code inside
P3, the tests run against the reverted tree inside P4. The deviation XENO-0108 had to
record does not recur, and this is the check that says so rather than a claim in a
commit message.

**The comment carries the reason and not its history.** It says what the paragraph
documents; that it once sat eighteen lines lower belongs to this intent's record and to
the commit, which is where it is.

<!-- xeno:section:release-notes -->
## Release notes

The doc comment about what `xeno init` did, what it left alone and what a person still
has to do now documents `printInit`. It previously ran into the comment of the next
function, so godoc published it as documentation of the enforcement report.

G-Schema's check for undeclared files in `evidence/` compares paths relative to
`evidence/` rather than basenames, and walks the directory rather than reading its top
level. A file declared under a subdirectory is recognised as declared, two files of the
same name in different directories no longer stand in for each other, and a directory
whose files are all declared is not a finding.

A declaration with no path, which is what a pending item carries, adds nothing to the
set. It previously added an entry keyed `.` that matched nothing.

A declaration whose path names something outside `evidence/` declares nothing inside it.

Nothing changes for a phase whose `evidence/` is flat, which is every phase this
repository carries and everything `xeno evidence attach` produces.

<!-- xeno:section:residual-risk -->
## Residual risk

**The class of defect behind the comment is untouched.** One comment moved. Nothing
renders godoc in CI, nothing checks that a doc comment reaches the function it names,
and the next paragraph written one blank line short will read as documentation of
whatever follows it. This fix is one instance of something a check could cover, and
naming that is all this phase can do about it.

**The case the second fix is for has no producer.** `Attach` writes one flat level under
`evidence/`, so a nested layout exists only in the four new tests. The fix is correct
against the paths a declaration can carry and unproven against a directory somebody
actually nested by hand. That is the ordinary state of a latent defect closed before its
data arrives, and it is worth saying rather than implying the opposite with a green
gate.

**A walk in a gate path.** One `Stat` per file per run, on directories holding a handful
of files, unmeasured. A phase with a large `evidence/` would be the first to find out,
and the gate path is the wrong place to discover a cost.

**The evidence is a local run.** As in the predecessor intent, and for the same reason:
the pipeline uploads nothing. `attached.yaml` cannot distinguish it from a pipeline
artifact and only prose says which it is.

**Accepted with the four named.** None blocks the change; each is smaller than the state
it replaces, which was a comment documenting the wrong function and a check that agreed
with its directory by accident.
