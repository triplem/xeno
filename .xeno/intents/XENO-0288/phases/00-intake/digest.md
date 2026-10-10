---
intent: github.com/triplem/xeno#334
phase: 00-intake
created: "2026-10-10T15:36:09Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 098667565f56ca8a3f905fe75d55831918384adb2f4cf4827ec95dae149bbcfa
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
The intake of #334: one new section of `docs/orchestrator-evaluation.md`, in the form
section 9 has for OpenSpec, for the AI-DLC methodology pinned at one tag. #323 did the
reading at `main` and weighed the three shapes; #334 is the work of writing the decision
down so that the next reader meets a decision rather than an open comparison.

The decision this phase had to carry was already a person's. The one thing in #323's
reading that was a decision rather than a finding — whether a Xeno verdict should bind the
source it judged by identity, the way AI-DLC's reviewed-source fingerprint does — was put
to the maintainer on the issue at 2026-10-10T13:41:17Z with three options, each with its
consequence and its cost, and a recommendation. He answered "Use the recommendation" at
14:59:46Z. So it is recorded here as D-1 and not as an assumption, and no question of this
intent is open. The reason that decides it is `docs/process-definition.md:480`, "A verdict
says what it judged, by content and not by commit", which is the normative document and
therefore a decision taken rather than a thing overlooked.

The pin is `awslabs/aidlc-workflows` `v2.11.0`, an annotated tag: the ref points at tag
object `4079edbe`, which points at commit `6a378b53c0a4fe0641ed7d8de8dfff94264d5b6a`. Both
shas answer to "v2.11.0" and a reader resolving the tag himself gets the other one, so the
section cites the commit and says which it is. The three preview tags cut around it are
numbered `2.11.1-preview`, which is ahead of `2.11.0`, so "2.11" alone is ambiguous and the
sha is what makes the reading repeatable. The hosted sample, `aws-samples/sample-collaborative-ai-dlc`,
is pinned the same way at `v2.2.0`, also annotated, commit `bc988d0eaaca752d9add2d67b2c861841a363227`.

Three files and 37,939 bytes of context, enumerated rather than globbed. The budget is 8
files and 140,000 bytes, deliberately above both: the work is to make one of those three
files substantially longer, the budget is judged from P1 against each phase's own lock, and
XENO-0286 paid for the version of this mistake where the figure was measured before the
writing that enlarged the files it measured.

Two large documents are excluded and quoted in the context rationale instead, so that no
later phase needs them: `docs/process-definition.md` for line 480, the gate table, section
10 and the override, and `docs/implementation-plan.md` for WP18 and WP21.

One finding about the plan, recorded as this phase's learning rather than absorbed: #334
carries no work-package label and neither did #336, because no package in the plan owns this
repository's own documents. Three intents have now resolved that by leaving the label off and
writing the reason into the intake.

Three sections and one decision, of the five sections this template defines.
