---
intent: github.com/triplem/xeno#179
phase: 02-design
created: "2026-10-03T12:15:34Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4d472b6.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b1e3edaa171f608ccf4ff3637c40cc4a6486be32deed051931e2e8996a135d5a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The command is `xeno intent start --for ISSUE [--intent KEY]`, with `needsKey` false for it alone
among the commands that name an intent, because here the intent does not exist yet. The id is built
in `internal/model`, where section 3's identity belongs, as `Tracker.Qualified(issue)` — a method on
the configuration block, so an error can name the field that is missing. One rule covers `--for`:
split at the last `#`, take the project from the left where it is there and from `tracker.project`
where it is not, and prefix the host where the project's first segment carries no dot, so `179`,
`triplem/xeno#179` and the whole id all arrive at the same place and a repository with no tracker
block still has a way through. The host is `base_url`'s host with a leading `api.` removed (A78),
because that is the only field in the block that names a machine and a name mapped to an address is
a table that is right until the first self-managed deployment — which is what `BranchRulesFor`
already refuses to do, in those words. The key is the shared prefix and the number after the highest,
padded as that key is (A79), which is D-7 read rather than reimplemented; where there is no sequence
or two prefixes it refuses and asks for `--intent`. What is checked about a key is that it holds no
space, because section 3 says the key need not be a number and YAML reads " #" as a comment. Seven
alternatives were refused, three of them by quoting a refusal this repository already makes.
