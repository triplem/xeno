---
intent: github.com/triplem/xeno#359
phase: 05-review
created: "2026-10-10T16:20:38Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8a60c9716fe5ae60679f3908cb5bc9464985150fd4fe26636314cf66f2b7000d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
The change is one new document and two one-line additions to the files that index it, with
no code and nothing normative moved. The three standing rules hold: both normative
documents are at zero lines of diff, nothing is added to any artifact's schema, and the
half of #359 that would have been an addition is a question on the issue rather than a
sentence in a document. The page answers "who acts on a label" in two halves: the tables,
and the fact that the host cannot tell the agent's acts from the maintainer's, because
there is no bot identity and `gh api user --jq .login` returns `triplem` — which is why the
written convention "the decision is the maintainer's; the transcription is not" carries a
distinction the tracker's own authorship field cannot. The three review rules of the
effective set are answered — one met, two not applicable — and a documentation lens entry
records the deviation that matters: a third of the page is a dated observation about a
tracker that no commit can hold it to. The residual risks are that staleness, the page's
central claim depending on code it does not sit beside, a question open on #359 that no
gate can see is open, two sealed figures that are wrong and corrected elsewhere, the line
and byte counts in `03-implementation` left low by the last changes to the page with the
reason for not chasing them, a reflow script that merged both of the page's code blocks
into prose while three checks passed over the damage, `04-verification`'s lock holding an earlier hash of that page
with the reason nothing will raise it, and the work package label still absent from the
issue. What was documented and what was drafted are kept apart in the checklist.
