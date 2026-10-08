---
intent: github.com/triplem/xeno#95
phase: 01-requirements
created: "2026-10-08T13:38:10Z"
schema_version: "1.0"
runner_version: dev+042c9bc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a7f8ca46b844f6762872e83ea8990b210ea16d081a20ca1fd4278772d0a17d77
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Five acceptance criteria for a configuration file, and the one that matters is AC-3: this
repository pins every action by sha and holds those pins against a page with a test, so an
update proposed automatically is an update that fails CI on the page rather than on the
bump. The criteria therefore ask the configuration to say what it does about that where a
reader meets it, rather than asking it to be clever.

AC-4 is the criterion that makes the intent's own claim checkable: nothing is bumped here,
and `go test ./internal/model/` passing is what says the pins and the page still agree.

One section written of the four the requirements template defines. `constraints` and
`non-goals` were left out: the constraint is AC-3 and the non-goals are the intake's
scope, and restating either would have been a second copy of a sentence already in the
trail.
