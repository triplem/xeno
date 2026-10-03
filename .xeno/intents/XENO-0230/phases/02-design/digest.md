---
intent: github.com/triplem/xeno#186
phase: 02-design
created: "2026-10-03T13:23:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fe2286b13b0dc85d68bdf13302a1ef15d9ff7fe59859965eaf6cf91e9cc6aa7f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`audit`'s `pull_request:` takes no filter at all, because a wider list is a thing to maintain and the
next file it forgets is the next merge past a red gate — and the job costs seven seconds. The three
scanners are renamed by their job id, since the id is what GitHub reports and `scan` in a file called
`trivy.yml` already told the reader less than the filename. `audit.yml` reads `semantic_version` out
of `release.yml` and refuses an empty result; the refusal matters more than the read, because a guard
that falls back to a literal audits the wrong tree and says nothing. The pin moves to 25.0.9 and node
is pinned with it in both workflows, as one change: a toolchain that declares an engine, run under
whatever image the runner ships, breaks in the release job and nowhere else. Node is pinned to the
major by D-5's reasoning. The baseline records the toolchain and what the upgrade moved, with exact
counts and no margin. The rule goes in `CONTRIBUTING.md` beside the merge procedure it belongs to,
naming the six commits, because a rule whose reason is omitted reads as caution. Seven alternatives
refused, three of them the same defect at different scales — a stale path list, a number written
twice, a margin on a drift check — each trading a silent wrong answer for a small convenience, and
each already taken once here. npm's own fix, a nine-major downgrade of the release tool, is a
resolver's answer to a constraint problem and not a security answer. A hook is refused on both of
section 7's constraints.
