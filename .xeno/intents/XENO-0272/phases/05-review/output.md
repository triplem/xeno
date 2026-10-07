---
intent: github.com/triplem/xeno#277
phase: 05-review
created: "2026-10-07T13:37:15Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1aa23a71d6c8988dc5c7a225f5a9d93d45d22d63b0932a1ee7040ace45294d41
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Two, both in P3, and both name what they depart from. The first is about P2''s reason rather than what was built: the replacement term keyed on a phase having no context.lock.yaml was measured and found to exempt nobody, every phase in the trail carrying one, and that count was taken while writing the design rather than before it, so P2 reads as though the conclusion came first. It names P2''s decision to remove rather than replace. The second names P1''s criterion 5, which said the five other cases of TestHashFieldShape are unchanged, in the singular: the removed behaviour was specified by two tables of that name, one in internal/gates/schema_test.go asserting the finding and one in internal/runner/runner_test.go asserting the colour a phase comes out as. Only the first was inverted when P3 was first finished and go test ./internal/gates/ passed, so the package under change was green while the suite was not. P3 was rewritten and judged again and its no_finding record was withdrawn by hand to carry the learning. The criterion is now met twice over.'
      result: deviation
      rule: deviations-are-traceable
    - note: 'There is no interface in the usual sense — no field, command, flag or exit code moves — and a gate got stricter, which section 13''s table calls a major step: a check that was green yesterday is red today. An artifact carrying context_hash: by-hand is now a finding whatever its tool field says. What an adopter has to do about it is in the release notes rather than in a migration note, because the only migration available is the one section 11 forbids: the artifact cannot be corrected, so the finding is released by a second person with a reason. This repository did that twice, for the two pre-M0 intakes, and those two phases now read approved instead of green. Answered as a deviation rather than not-applicable because the rule asks about an interface change and the honest answer is that this is a behaviour change with the same consequence for somebody who has artifacts already.'
      result: deviation
      rule: interface-change-needs-a-migration-note
    - note: None added. go.mod and go.sum are untouched and nothing compiles differently. The diff removes a variable and a term, replaces two comments, inverts one case in each of two test tables, moves two gate.yaml files from green to approved, and adds one row to docs/assumptions.md.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** No normative document is touched: section 12's paragraph on the
triple is the clause this change acts on and it stays as it is. Nothing is invented — a term is
removed. The change belongs to WP1 by subject, since the artifact schema and the core gates are
where `hashes` lives, and to intent XENO-0272 for issue #277.

**What the issue asked, and what it got.** #277's done-when is "either the exemption no longer
rests on a field nothing checks, or section 12 says that it does and why that is acceptable".
The first. The maintainer was put all three shapes with the measurement and chose the narrower
exemption, and what remains in `honest` is two terms that are facts about the tree.

**What a reviewer should look at first: two sealed phases changed verdict.** XENO-1's and
XENO-2's intakes were green and are now approved, permanently, because a check got stricter
after they were sealed. That is a real cost and it is the whole price of the change. Both
alternatives that would have avoided it were worse: leaving the exemption keeps a verdict
resting on a declaration, and writing the real hash into the four artifacts is the rewriting of
a sealed trail section 11 forbids. Anything in this repository that counts green phases counts
two fewer from now on.

**The second thing, and it is about me rather than the code.** The four approvals are a
person's statement, and the agent typed them. The maintainer chose the shape and approved the
wording before the commands were run, and `--by` names them, so the record is accurate about
who decided. It is still further from what the process protects than the specification passages
written on instruction earlier today: `next.go` refuses to offer `gate approve` as a command
specifically so that nothing nudges towards the decision it exists to record, and an agent
running it four times is past nudging. P4's gaps say the record names who decided and not who
typed. If that is the wrong trade, the four decisions can be withdrawn and retyped by hand and
nothing else in the change depends on it.

**The question no rule asks: is a comment and a register row enough to keep the rule?** The
change establishes that what makes a placeholder honest is only a fact about the tree. Nothing
asserts it. A new term reading a declared field would pass the suite, because the two inverted
cases assert the consequence for `context_hash` and not the rule. A test that asserted the rule
would have to inspect the expression or enumerate the fields the gate may read, and both are
tests of how the code is written rather than of what it does — which is why this is a question
rather than a non-goal with a reason.

**And: the rule is met while one of its two terms is coarse.** `goneBundle` is a fact about the
tree and it cannot tell a bundle that is gone from one that was renamed, which #273 measured
earlier today and left standing. So "only facts about the tree" is true and one of the facts is
blunter than the sentence suggests. Both are now written down in the register, A101 and A33,
which is the most this intent can do about it.

**What CI forced after the work was judged, and where it is recorded.** Approving the four
findings rewrote XENO-1's and XENO-2's `gate.yaml`, which put both intents in the commit range,
and `Completeness` requires every touched intent to be `complete` or `abandoned`. Both are
pre-M0 intents whose later phases are `not-started` on purpose, which `Summarise` calls
`in-progress`, so `xeno intent verify` refused the branch. There is no way to resolve a finding
on such an intent without writing to its files, so the problem is not avoidable by doing the
release differently.

The maintainer was put three options — teach `Completeness` the M0 shape, which needs section 8
to gain a third ending and so a specification commit; close the two intents, which is the
ending section 8 already offers; or revert to #277's clause — and chose the close. Each now
carries `status: abandoned` with a reason saying that nothing was dropped, that
`m0-gate-job.md` and A20 describe the shape, that `abandoned` is the ending available rather
than the one that fits, and why it happened today. Each also gained the intent level
`learning.yaml` section 10 owes at a close, recorded as `--no-finding`: the close is
administrative and the work it ends was judged weeks ago.

Two things about it are worth a reviewer's attention. It closes two of roughly 106 intents in
that shape, because this branch happened to touch them, so a reader comparing XENO-1 with
XENO-3 finds no reason for the difference beyond the reason field. And the first `intent close`
ran half way — it wrote `status: abandoned` and then went red on the missing learning record —
so `intent.yaml` was reset from `origin/main` and the close run once more, which is why both
records are consistent now and why the intermediate state is named here rather than left for
somebody to infer.

It is recorded in this phase and not in P3, where the implementation is described, because it
happened after P3 was judged and re-judging P3 would have made P4 stale and P5 with it. The
phase that records a thing is the phase that was open when it happened; what P3 describes is
still accurate about what it describes.

**One process note.** This phase's predecessor was judged twice. The removed behaviour was
specified in two test tables with the same name, `go test ./internal/gates/` passed with only
the first inverted, and the second surfaced on `go test ./...` during verification. P3 was
rewritten and judged again, its `no_finding` withdrawn by hand to carry the learning, and P4
was started over because reading the staleness finding with `gate run` had written it a verdict.
The trail shows all of that; what it does not show without this paragraph is the order.

<!-- xeno:section:release-notes -->
## Release notes

**A gate no longer relaxes itself on a field the artifact declares about itself.** `hashes`
accepted the `by-hand` hash placeholder in every field when the artifact said `tool: manual`.
That term is gone. What makes a placeholder honest is now only a fact about the tree: a field
nothing writes yet, which is `secrets_hash` and `rules_hash`, or a strings bundle the
repository no longer carries.

**This is a stricter gate, which by section 13's table is a major step.** An artifact carrying
`context_hash: by-hand` is a finding from now on, whatever its `tool` field says. If you have
hand written artifacts in a trail, they will raise one, and the artifact cannot be corrected:
section 11 forbids rewriting a sealed artifact to satisfy a check. The way out is the one the
process already has — a second person releases the finding with `xeno gate approve`, naming a
reason.

**In this repository that happened twice.** Four artifacts in 499 phases declared
`tool: manual`: the intake output and digest of XENO-1 and XENO-2, written by hand on
2026-09-21 and 2026-09-22 before M0. Removing the term raised four `context_hash says by-hand
where a writer exists` findings and nothing else, because the other two terms already covered
every other placeholder those artifacts carry. All four were approved, with a reason naming
each artifact's date, that the lock beside it is hashable today so the finding is correct
rather than spurious, and that the artifact cannot be corrected. Those two phases now read
`approved` where they read `green`.

**Why this was worth a change rather than a sentence.** Section 12 says the `model`, `tool` and
`tool_version` triple is a declaration and that nothing corroborates it. `tool` was the only one
of the three any gate read, so one verdict in the process rested on a self-report. It no longer
does, and `docs/assumptions.md` carries A101 with the measurement and the decision.

**Two pre-M0 intents are now closed.** XENO-1 and XENO-2 carry `status: abandoned`
with a reason saying that nothing was dropped: `m0-gate-job.md` runs one intake per work
package and goes no further, A20 records the shape, and `abandoned` is the only ending
section 8 offers for an intent that will not reach a decided P5. They were closed because
resolving a finding on an intent puts it in a commit range where `xeno intent verify`
requires one of those two endings, and about 106 intents are in the same shape without
being closed.

<!-- xeno:section:residual-risk -->
## Residual risk

**Two sealed phases read `approved` for ever, for a check that changed after they were
sealed.** XENO-1's and XENO-2's intakes. Nothing can close the four findings, because the
artifacts cannot be corrected. Anything that counts green phases in this trail counts two
fewer, and a reader who finds them has to follow the reason to learn the cause is a later
change rather than anything those intents did.

**The four approvals were typed by the agent.** The maintainer chose the shape and approved the
wording and `--by` names them, so the record is accurate about who decided. `next.go` refuses
to offer `gate approve` as a command so that nothing nudges towards the decision it exists to
record, and an agent running it four times is past nudging. The decisions can be withdrawn and
retyped by hand; nothing else in the change depends on them.

**Nothing asserts the rule the change establishes.** A term reading a declared field would pass
the suite. The two inverted cases assert the consequence for `context_hash`, which is the right
thing for them to assert, and the rule — only facts about the tree — lives in a comment and in
A101. A test for it would have to inspect the expression or enumerate the fields the gate may
read, which tests how the code is written.

**One of the two remaining terms is coarser than the rule sounds.** `goneBundle` cannot tell a
bundle that is gone from one that was renamed, measured in #273 earlier today and left
standing. So "only facts about the tree" is true and one of the facts is blunt.

**The sample is two.** Twice in 499 phases is all the evidence there is about how often this
finding fires. A project that writes artifacts by hand meets it more often and nothing here
tells them how much more.

**The verification phase's lock postdates the implementation it verifies.** P3 was judged
twice, because the removed behaviour was specified in two test tables and only one was found
before the first judging; P4 was then started over, because reading the resulting staleness
finding with `gate run` had written it a verdict. The artifacts are each honest and the order of
events is only in P5 and in P3's deviations.

**For a person, not for the code.** Whether a check that turns two historical phases amber is
worth closing a gap nobody had exploited. The threat model this project does not have is a
harness that lies about itself, and the argument for the change is not about threat: it is that
a verdict should not rest on a declaration the documents call a declaration. Somebody could
reasonably weigh those the other way, and the register row is where that disagreement would go.

**Two intents are closed and about 104 others in the same shape are not.** XENO-1 and
XENO-2 were closed because this branch touched them, and nothing distinguishes them from
XENO-3 onwards except that. The reason field carries the truth and the `status` field
carries `abandoned`, which is the wrong word for what they are; `Summarise` will now call
them `abandoned` and `xeno intent status` will show it. The alternative was a third ending
in section 8, which is a specification change the maintainer declined for now, so the
inconsistency is the chosen cost and it grows each time a change touches another pre-M0
intent.

**A half-run `intent close` left a record nothing recomputes.** The first attempt wrote
`status: abandoned` and then failed its own gate on the missing learning record, leaving a
red close gate that `gate verify` does not recompute — it checks the 513 phase verdicts and
not the intent level one. The state was reset and the close run again, so the committed
record is consistent, but the general fact stands: an intent level gate can be stale and the
command that recomputes everything else will not say so.
