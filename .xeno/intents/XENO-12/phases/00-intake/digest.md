---
intent: github.com/triplem/xeno#12
phase: 00-intake
created: 2026-09-24T20:34:38Z
schema_version: "1.0"
runner_version: 0.1.0-dev+c7abd5d
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: ce52dcd78a2b436b95932f446aa56a356fd13038320048e89795eb847a6b3814
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Four rule sets were run against this repository and all reported zero, which is worth
nothing until a scan that should find something does. A canary file was written for that
and is found. The vendored pack was then checked three ways against the registry version:
same canary findings, same zero here, and it runs with a proxy pointed at a dead port.

The exit code was measured rather than assumed: without `--error` semgrep leaves 0 even
with findings, so a job built on the default would have reported success on everything.

No secret filter exists, so nothing filtered this text.
