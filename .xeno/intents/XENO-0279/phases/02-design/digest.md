---
intent: github.com/triplem/xeno#94
phase: 02-design
created: "2026-10-08T20:16:08Z"
schema_version: "1.0"
runner_version: dev+30b1dea
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0f0aca2a2a02af89a00b250da7db6f9f900aca0a45ff80e175749d222ca0f5b7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The design for #94. The policy is one file that fails on any finding; four gosec
findings in shipped code are accepted on their lines with the reason, tests are outside
gosec by one rule with its reason and inside errcheck and staticcheck with their
findings fixed. Two workflows in the scanners' shape, the shape cited from trivy.yml
rather than repeated: golangci-lint through its action pinned by sha and version,
govulncheck by a pinned go install, its JSON report judged by a step because the JSON
mode exits 0 whatever it finds, measured. govulncheck runs daily, lint does not need
a schedule. semgrep goes in this change and the required check moves on the host before
the merge. depguard is declined against the approved plan because it reads direct
imports and A42 is about the closure.
