---
intent: github.com/triplem/xeno#169
phase: 00-intake
created: "2026-10-03T09:00:28Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+93322c5.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 16d0dc70171243e574699e02917b14c99561eb51c01f7dd9f69af1861789538a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
M0 has two readings and one is met: the milestone table's gate job is required on the default branch
with administrators included, green over 207 verdicts, with twenty-plus six-phase intents behind it.
Section 6 puts WP11 before M0 and WP11 is unbuilt — no `plugin.json`, no skills, no `mcp.json`, no hook
wiring. What drives a phase today is an agent with the specification in context calling the runner's
commands, which is the fallback WP11's done-when requires rather than the package; what it costs is
invisible here and fatal elsewhere, because a project adopting Xeno gets the runner, the templates and
the rules with no description of what a phase is for or which command carries which step. The plan's
one open verification point is answered on this machine: the client's help says `marketplace add` takes
a URL, a path or a GitHub repository, adding a directory succeeded, and `claude plugin validate` passes
both manifests — after which the test marketplace was removed. Answering it produced a disagreement:
section 13 draws `plugin.json` at the plugin root and the client loads it from `.claude-plugin/`, which
is the one case where "the specification wins" cannot be followed literally, because a tree that
followed it would not load. And the hooks sit in this project's own `.claude/settings.json`, so an
adopting project records no cost at all.
