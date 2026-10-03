---
intent: github.com/triplem/xeno#179
phase: 00-intake
created: "2026-10-03T12:13:38Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4d472b6.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1f18dba6bf50164a47bf524f262249a8e2f5a8c84a5271e4e61901c303ea3a21
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`intent.yaml` is the one artifact of this process that no command writes, so three of its seven
fields are guessed where the runner knows them: `created` is a typed timestamp, and the two version
fields are copied from memory, which is why every `intent.yaml` here records `0.1.0-dev` while the
artifacts beside it record the commit as well. `status: in-progress` is the only value a person
ever types that the runner would have defaulted to, since `intent close` writes the other one and
A20 says there is no third. The file sits inside `artifacts_hash`, so a typo in a hand-written
field is sealed with the phase and a correction reads as a divergence. What is left once the
derivable fields are taken away is the issue: `tracker.project` holds the repository and
`tracker.base_url` names the host, so the key of the issue is the one input. The command is
therefore `xeno intent start --for ISSUE`, and `--intent` is optional rather than required as the
issue wrote it, because the intent's own key is the next number of the sequence the intents
directory already holds. Out of scope and each for a reason: `assumptions.yaml`, which nothing
refuses for; the seventy-nine hand-written intents, whose keys sit inside hashes and merge commits;
reading the issue from the host, which belongs to the open half of WP12; and a `merged` status,
which A20 decided against.
