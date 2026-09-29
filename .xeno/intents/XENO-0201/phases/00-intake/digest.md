---
intent: github.com/triplem/xeno#120
phase: 00-intake
created: "2026-09-28T20:29:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6e72fed.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 27f9bd71ff159236dbda07d988c3f750316110ceae0d6a82533228b1d6821738
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The scoping of #120 said the writers had to come first and this phase found that two of
the three missing fields were never missing: project.yaml carries the agent block and
the runner reads that file for the language alone. The restraint worth recording is
secrets_hash, which stays absent because section 5 says the runner filters and there is
no filter, so the digest the runner writes must not claim to have passed through one.
