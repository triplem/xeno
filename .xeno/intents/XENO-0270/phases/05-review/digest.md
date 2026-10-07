---
intent: github.com/triplem/xeno#231
phase: 05-review
created: "2026-10-07T09:51:25Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ec88b96575e78b36934a55c398da078fc1a245e0478084e7a97b0f1f52b5ce0b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three standing rules hold: no normative document is touched, nothing is invented, and the
change belongs to WP16 by subject and to intent XENO-0270 for issue #231.

All four of #231's done-when clauses are met, with one correction from the maintainer's
comment: the issue said the existing material stays further down and named the no-network-call
property among it, and the comment said that is deep information and should move. It moved,
and the half section 12 does not carry moved with it rather than being dropped.

What a reviewer should look at first is the per command notes at the bottom of the reference.
They are prose about behaviour beside a block a test holds, and nothing holds them, which is a
smaller form of the defect this intent was filed about. The choice was between an unheld note
and losing five paragraphs, and criterion 10 forbids the loss.

Two questions no rule asks are answered: whether four command lines are the right number, when
a phase does not finish without `section set` and the block leaves it to the prose underneath;
and whether the index earns its place before WP16 exists, which it does on the grouping and
the normative marking that a directory listing cannot carry.

Three rules answered, all not-applicable. The residual risk is the unheld prose, three files
no test holds, a coverage table of unknown age, a four-line block that will not run a phase on
its own, a guard WP16 may read as the design, a filename a generator will want changed, and
two wrong claims written inside this intent that reading caught and nothing else would have.
