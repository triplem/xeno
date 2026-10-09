---
intent: github.com/triplem/xeno#94
phase: 00-intake
created: "2026-10-08T20:12:42Z"
schema_version: "1.0"
runner_version: dev+30b1dea
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f7942c793409a86798848db26f58f4758c59e7d84267d0c9fe50af4ae692cb99
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake for #94, the first intent started through #330's approval clause against the
real host. The issue's body names three tools; its comments hold the analysis and the
maintainer's three answers, restated here: replace semgrep with golangci-lint running
gosec, staticcheck and errcheck, add govulncheck for the source, keep trivy for what the
release ships, keep the report plumbing. Measured first: 51 findings today, four gosec
ones in shipped code, none from govulncheck, so zero is reachable. One step of the plan
does not survive measurement: depguard reads direct imports and A42 is about the
closure, so the `go list -deps` step keeps the rule and depguard is not enabled.
