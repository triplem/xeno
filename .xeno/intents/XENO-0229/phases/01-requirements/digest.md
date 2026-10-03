---
intent: github.com/triplem/xeno#184
phase: 01-requirements
created: "2026-10-03T12:54:42Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5b221a1f0bd97eb1a152932f25c7bb73149cde8cc0329df694eab03f3610ce02
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The specification change is its own commit made before any code, and `export
XENO_HARNESS_VERSION=2.1.276` then carries `tool_version` into every artifact of every phase of a
session with no flag anywhere. `--tool-version` over the top of it wins, because the variable is how
a session was set up and a flag is somebody saying it about this run. Neither said leaves the field
absent in both files, which is unchanged from #181 and is still A35's rule. The variable is read
once, where the runner is constructed, because a value that holds for a session is read at the start
of one and not re-read per command — which also decides that `parse` must stop assigning the field
unconditionally. The tests clear it rather than inheriting what the machine exports, since a
developer with it in their profile would otherwise watch the absence tests pass a value. The runner
reads this variable and no other: not the remaining three in section 7's list, and no client
specific one, because normalising those is the entry point's work by the section's own account. The
flag is not removed; it becomes the override A80 said it would become. Out of scope and each for a
reason: the entry point and the other three variables, which stay with #183 and leave the runner
reading exactly one of four; any default, which is the one line that must not move; per-command
re-reading, which would make one phase disagree with another for a change nobody made; and
backfilling XENO-0228, whose values are correct.
