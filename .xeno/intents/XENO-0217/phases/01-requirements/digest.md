---
intent: github.com/triplem/xeno#158
phase: 01-requirements
created: "2026-10-01T15:01:36Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0c8276d54456994c605672f9d24129b76ba26ae05f013520ec7934483b064748
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The criteria are section 9's sentences turned into fixtures. A repository with no rule tree is green,
which is the first case and not the edge, since it is the state of every repository today. A well formed
tree resolves in precedence order into an effective set, returned as a set rather than a verdict because
section 7 says G-Rules and G-Policy should resolve the tree once. Red, each with its own fixture: a
`scope` disagreeing with its path; `binding` outside `given/`; `binding` under `given/project/`; anything
under `learned/builtin/`, on the directory; `checked` without a `check`; `review` with one; a duplicate
id at one level; and two binding rules colliding across levels, where staying red rather than resolving
to the more specific is the criterion that matters. `rules_hash` is byte exact, reproducible, unaffected
by walk order, defined in `ASSUMPTIONS.md` and tested against `sha256sum` rather than against itself.
Non-goals: no predicate evaluated, no checklist, no rule shipped, no configuration field, no `enabled`,
and no sealed artifact rewritten — seventy-one carry `by-hand` and keep it, so the trail holds two
meanings of the field divided by this commit.
