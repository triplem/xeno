---
intent: github.com/triplem/xeno#190
phase: 04-verification
created: "2026-10-03T14:05:33Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+74553ad.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c0719b7b288bb990a4f900d8b293d2f89b1bed15ed15645b017a717475c97a86
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Both checks are workflow steps, so the evidence is runs and repeatable commands rather than tests.
Before either was written, in a throwaway worktree off main: a committed deletion of `XENO-0225`
reports 31 paths and fails; one new intent directory reports 0 and passes; a directory added and then
removed across two commits reports 0 and passes, which is the case the two-dot form exists for and
the one a commit range would have failed. Six titles were checked against both shipped patterns, and
the plain one accepts `feat(runner): a thing`, the same with `(#191)` and `chore(deps)!: breaking`
while refusing `The audit gate binds`; the other accepts only the form carrying a reference, which is
A71's definition and why the plain pattern was chosen. Eighteen packages ok, `gofmt` and `go vet`
clean with no Go file touched, `gate verify` 265 at exit 0. This pull request is the first live case
of both steps and is expected to pass, which is the weaker half of the evidence; the worktree table
is the strong half, because there each step was shown failing on the case it exists for. Five gaps:
neither step is shown failing in CI, and closing that means a pull request whose purpose is to be
refused; the re-seal hole is now the only quiet way to rewrite the trail, deliberately left because
refusing modification would refuse a re-judgement; G-Complete still runs only at P5; the fifteen
releases stay undescribed; and a title edited after the last push is never re-checked, because
`pull_request` does not fire on `edited` by default.
