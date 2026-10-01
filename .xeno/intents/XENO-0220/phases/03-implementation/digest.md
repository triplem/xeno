---
intent: github.com/triplem/xeno#165
phase: 03-implementation
created: "2026-10-01T16:34:40Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4f94129.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: aa2023e14ab6ca395f7a8b8de1884ffe94fcb9012a6de14f9b87315f53a4d092
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Four rules under `given/builtin/`, three review and one checked, each carrying a comment that it is
shipped, pinned and overridden rather than deleted. `section-non-empty` beside the implication type,
with `unknownSection` and `renderedSections` factored out of both. Four examples with `scope: org`, each
saying in its own text what it costs and what it does not check. Two hook templates carrying their
limits: comfort and bypassable, enforcement and needing the binary. `xeno init` vendors the rule tree
through the same `create`, so a second run still changes nothing. And the mechanism the design did not
have: writing the set turned all twenty-four sealed P5 phases red on `gate verify`, because the review
rules require entries those artifacts cannot carry, so G-Policy now judges a phase only against the set
its own `rules_hash` records (A74) — measured rather than reasoned, since the reasoning alone would have
sounded like caution. The hole it leaves is recorded rather than closed: a rule added after an artifact
was written leaves it unjudged until something renders it again, and closing that means deciding whether
a rule set change makes a phase stale, which is a specification question. Six deviations, including that
the set omits the one rule section 9 writes out in full, because the templates have neither section it
names.
