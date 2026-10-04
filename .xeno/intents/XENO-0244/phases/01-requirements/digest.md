---
intent: github.com/triplem/xeno#221
phase: 01-requirements
created: "2026-10-04T09:28:45Z"
schema_version: "1.0"
runner_version: dev+24becc3
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6784c959221bf4600bd84653f30aee81a6ad0262663b878af724d215c9c76a15
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The requirements phase records D-1, which is Q-1's first option, and turns it into ten
criteria. The load bearing one is not about the gate: `gate verify` at exit 0 over the
whole trail with an unchanged verdict count and an unchanged `artifacts_hash` per phase,
measured three times — before the change, after the correction, after the gate — because
that figure is the entire argument for doing the correction first. Two criteria keep the
check narrow: a value outside the set is a finding, an absent result is not, since
section 4 has the field absent where a producer reports nothing and this repository's
pipeline publishes exactly such an entry in `other/trivy-db`. Two keep the refusal in
the shape `Attach` already has: the entry is declined rather than recorded, counted as
pending, with the reason returned in `Unbindable` so three callers each say it in their
own voice, and nothing is left behind for a republished manifest to trip over. Seven
things are out of scope, the first being the one worth filing: section 4 requires a
result on `test-report` and `build-log` because G-Test and G-Build read it, and an
attachment arriving without one leaves the same hole one level down. It is not in D-1,
no attachment in the trail lacks a result, and widening the scope while correcting
fifteen sealed files is how a correction becomes a redesign. The constraints name the
order as a constraint rather than as a plan, because reordering the two steps is the one
way this intent can go wrong in a way CI reports against history. Read in this phase:
section 4 once more against each criterion, `attach.go` for where `Unbindable` is read
by each of the three callers, and the trivy manifest for the entry that has no result.
