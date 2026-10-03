---
intent: github.com/triplem/xeno#183
phase: 00-intake
created: "2026-10-03T18:37:22Z"
schema_version: "1.0"
runner_version: dev+0462eb7.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5f7b8cd6560c75bb4a1bd9c254bd3d3fb25963f1d28ffcc24f57241c6575ee8a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Section 7 described a resolution order for the plugin root and nothing read it — the last open clause
of #183, open for three intents while each worked around it. It could not be built as written:
`internal/gates` reads `internal/rules` and `internal/template`, both resolving from the plugin
directory, so a root from the environment makes `rules_hash`, `strings_hash` and a rendered artifact
depend on it, and A74 judges a phase only against the set its own artifact records — so a changed set
does not disagree with the trail, it stops judging it. Measured: a valid rule tree that differs
leaves `gate verify` at exit 0 over 273 verdicts while G-Policy silently reports nothing about 75
phases. The thing the section relied on arrived and did not settle it: its own safety sentence is "a
mismatch... G-Supply fails red", and that gate catches an override for a released binary while a
development build carries no digest and is what wrote every artifact here. And the bottom of the
order could not work at all — a project with no vendored plugin resolves no template and has no
`plugin_version`, so the client fallback lets it fail differently rather than work, which is the
measurement that turned the question from how to guard the order into whether to keep it. Every
position was unreachable or unsafe. In scope: the specification change as its own commit, and the
comments and rows that described the order as a decision pending. Out of scope: implementing an
override under a condition, which was the declined alternative and is what the one sentence section 7
keeps would start from; and anything about how the plugin is found, since the code already read the
vendored directory — this is the document catching up with it.
