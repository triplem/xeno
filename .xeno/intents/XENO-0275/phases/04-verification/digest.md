---
intent: github.com/triplem/xeno#324
phase: 04-verification
created: "2026-10-08T13:49:07Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 94d31936a06e7ada2eabd5194359e0eb252cb40db55365476e9bc8c0e2a1b7f3
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Five criteria, five met, and the checks are over the rendered wrapper rather than the template:
`xeno init --host github` in a throwaway repository carries the commented line under the
image with its six lines of reason, the file parses as YAML, and `xeno init --host gitlab`
mentions no user at all.

The check that matters about this phase is the one that failed first, and the failure was
mine: I read a wrapper rendered by a binary built before the template changed, because the
templates are embedded. The test over the template passed at the same moment, which is the
whole gap between a test and a check over what an adopter gets. That is this phase learning
record.

Not checked: no container job ran on any runner, so uid 1001 matching a hosted work directory
is still read from the runner sources rather than observed. #305 left that open and this
intent does not close it.

Two sections, which is what this template requires.
