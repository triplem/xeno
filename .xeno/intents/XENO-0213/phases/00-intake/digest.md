---
intent: github.com/triplem/xeno#145
phase: 00-intake
created: "2026-09-30T05:43:53Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4e41c6b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3843c51c73f5b0ff2f21e956d1c0c2e97dd28f4681626fafd4e1084ce71c08cf
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
wrapperHosts has been a one row table since #98 so that a second host is a row and a scaffold file, and the
claim has never been tested. Beyond the missing row, two things are wrong: project.yaml writes adapter and
base_url as literals while the wrapper is generated from --host, so --host gitlab produces a repository whose
pipeline is one host's and whose tracker configuration is another's; and --host itself defaults to github in
cmd/xeno, which is the last default host name in code and the remainder of #104's first done-when. Confirming
the two range expressions against GitLab's reference added a constraint #104 did not carry: the base SHA exists
only in merge request pipelines, so the job must be scoped to merge_request_event or the runner judges a range
nobody wrote.
