---
intent: github.com/triplem/xeno#120
phase: 04-verification
created: "2026-09-28T20:35:34Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6e72fed.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 9f98e596b5914c1541794cfd924dbcc3b90fda4594cfa045bd4bf12ba4e5a57c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The mapping lost the column the last four intents relied on. Finish gained a parameter, so the new
tests do not compile against the old tree and there is no row of failures to report; the substitute is
stronger than a test run, since no digest in this repository was written by a runner before this change
because no writer existed. The result worth recording is that the first one the runner wrote turned its
phase red on the two fields with no source, which is the writer working rather than failing.
