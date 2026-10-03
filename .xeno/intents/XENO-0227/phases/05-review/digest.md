---
intent: github.com/triplem/xeno#179
phase: 05-review
created: "2026-10-03T12:19:32Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4d472b6.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6a7702e7d5515a764c7dc5a928a6bb28343da58e73467655b20b484a3212a44b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three shipped review rules are answered. `deviations-are-traceable` is met with two: `--intent`
optional where the issue wrote it required, and the tracker block consolidated where the design did
not ask for it. `interface-change-needs-a-migration-note` is met because the surface gains a command
and takes nothing away — a hand-written `intent.yaml` is read exactly as before, which is a criterion
of this intent rather than a note about it. `new-dependency-needs-a-rationale` is not applicable;
nothing was added. Beyond the rules: the departure from the issue was raised before any code existed,
answered, and recorded in four phases, so it is a decision in the trail and not a discovery at
review. Nothing is invented — no field in `intent.yaml`, none in Appendix A, no gate, no rule, no
status — and the two derivations read values already in the repository and refuse where those are
silent, which is what A78 and A79 are rows for. No verdict behind this intent can change, because
the gate list and the rule set are untouched: `gate verify` is 241 at exit 0 against 237 before, the
four being this intent's own phases. The residual risks are the host derivation nobody can point at
a self-managed deployment, a wrong issue number still sealed as it was before, a two-prefix
repository nobody has met, and `tool_version`, which this change neither caused nor worsened and
which is now written down as a learning that wants an issue.
