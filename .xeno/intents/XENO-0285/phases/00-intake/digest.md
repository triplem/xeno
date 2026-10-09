---
intent: github.com/triplem/xeno#95
phase: 00-intake
created: "2026-10-09T16:46:57Z"
schema_version: "1.0"
runner_version: dev+d2bc411.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 529394b65a69b529fb75660837924fcfe7496c59e8ae5cb1dfd4598b927d7989
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
The intake of #95 a third time, to finish the configuration rather than to decide its shape.
XENO-0274 contradicted the issue's answers, XENO-0277 corrected it, and what is left is the
argument renovate was chosen on: it can see versions written inside workflow text, and three
such pins have no manager at all.

Those three were found by running each existing `customManager` regex against the tree and
reading what it resolved to. Six resolve to the version the tree carries; the `Dockerfile`,
`gitleaks.yml` and `docs.yml` are matched by nothing, and all three are rows of
`docs/supply-chain.md`. The six passing beside them is what makes the three a finding rather
than a harness that never worked.

Two decisions are recorded here, and neither was put in this issue. D-1 is that #95 closes on
the configuration being complete, with switching on as #350, so the issue does not stay open as
a reminder of a secret. D-2 is that the agent set the `xeno-approved` label itself, on explicit
instruction, after the label was twice reported as not landing — the approval is the
maintainer's and in words, the label write is not, and an agent that may set its own input makes
`xeno intent start`'s check unfalsifiable. That is the property #330 exists to create and #95 is
the case it was written from, so it belongs in the trail rather than in a session.

No open question. The one trade-off that looked like one — covering `gitleaks`, whose proposal
can never be complete because renovate moves the version and cannot recompute the checksum
beside it — is the bargain this configuration already struck for `github-actions`, where a
bumped sha needs a Markdown row moved by hand. Precedent, not a new question.

The context is thirteen files and 103,892 bytes against a budget set from those figures.
XENO-0277 declared ten, resolved eleven, and carried two advisory findings through every phase
for it.

Three sections and two decisions, of the five this template defines.
