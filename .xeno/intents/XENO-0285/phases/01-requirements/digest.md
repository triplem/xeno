---
intent: github.com/triplem/xeno#95
phase: 01-requirements
created: "2026-10-09T16:48:52Z"
schema_version: "1.0"
runner_version: dev+d2bc411.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 96b20eddde2429059cf00105ca3f91d1228819255f06e8371bbf7b4986d6a730
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
Seven acceptance criteria, and the first is an accounting rather than a count. The page's
thirty pipeline rows and one module row were classified row by row: thirteen actions by sha,
one module, the `go` directive, six tool versions inside workflow text, the three this intent
reaches, two toolchain majors left alone deliberately, gitleaks' translated rules, and four
that cannot be pinned. A first draft of this section said "twenty-four rows" from memory and
was wrong by six; counting them is what the criterion is for, so it was counted before it was
written down.

The second criterion is the one that decides whether this intent did anything: nine managers,
nine resolved values, with the six that already exist run beside the three new ones. That
ordering is deliberate. Six resolving correctly is what makes a tenth resolving to nothing a
finding rather than a script that never worked.

The constraint worth naming is that renovate's regex dialect is not Go's, so these regexes
cannot be checked by the test suite. There is no renovate binary here and adding one is a
second dependency, which this project treats as a decision rather than a step. So the check is
a script, run in the verification phase and recorded as evidence.

A gitleaks proposal can never be complete, and that is accepted rather than solved: renovate
moves the version and cannot recompute the checksum beside it. The configuration already
struck that bargain for `github-actions`, whose bumped sha needs a Markdown row moved by hand,
and its `packageRules` entry says so. Loud and waiting for somebody beats a pin going stale
with nothing to notice.

Three sections of the five, no question, no decision.
