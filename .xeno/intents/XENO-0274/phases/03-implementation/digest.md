---
intent: github.com/triplem/xeno#95
phase: 03-implementation
created: "2026-10-08T13:40:15Z"
schema_version: "1.0"
runner_version: dev+042c9bc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 712f210d6a5ae60d551299ee40074a7ccf24ef86f36bed4dee6da26081df429a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
One file, `renovate.json`, and nothing else touched.

Three managers for the three kinds of pin this repository holds, the dashboard rather than
pull requests as the issue asked, and `dependencyDashboardApproval` on every rule — not as
caution but because a bumped sha and its row in `docs/supply-chain.md` move together and
only a person can move a row. Each rule carries that reason in its own `description`, where
renovate reads it and so does the next person to open the file.

One exception: `vulnerabilityAlerts` does not wait for the dashboard, because an advisory
arrives without anything in the repository moving and CONTRIBUTING.md already puts that work
before the next merge.

Nothing is bumped, which is acceptance criterion 4. One section written, `changes`, of the
three this template defines.
