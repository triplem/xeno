---
intent: github.com/triplem/xeno#228
phase: 01-requirements
created: "2026-10-06T20:28:09Z"
schema_version: "1.0"
runner_version: dev+3f1fab3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1f573db0af2bf07b90d09b3305f207aa78d2f5dd97fa98f4f6ba06b6afa3673c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Eleven criteria. Two of them are the sentences themselves: the `not-started` suggestion says
the phase is started in a fresh session and why, in the lock's own terms, and the last phase's
suggestion says the same for the session that is ending. Three constrain their form: the
command stays, no harness is named, nothing is enforced. Two are the tests that hold the
wording. One is the row in `docs/assumptions.md` carrying the two decisions #228 asks for
beyond the sentence. The last is the project's own check suite.

The non goals are where the issue's other candidates are refused with their reasons: no
document change, because the suggestion reports the sequence rather than adding to it; no
harness command anywhere, including the plugin, because section 13 ships one vendored tree for
both clients; no hook; no cost figure, because WP20 owns session discipline and the baseline;
no gate; and no sentence on a phase that is already open or on `intent close`, which abandons
an intent rather than finishing one.

The binding constraints are section 7's rule against branching on the harness, `Suggestion`'s
own comment on what a suggestion may be, and the first standing rule, which is why the
specification is untouched and why the register row is the only prose outside the runner.
