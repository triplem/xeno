---
intent: github.com/triplem/xeno#167
phase: 02-design
created: "2026-10-01T16:48:23Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15693cf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 29387d7f9ab77f267b52216c43fa30c81256709e45ac9ac2916ae37ce72fa43b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`internal/external` runs the commands and `internal/gates` keeps the rules, with the runner injecting an
`External` function into `gates.Ctx` so the external checks pass through the same `carryForward`,
`Invariants` and `Status` as every other check — one assembly point, and a verdict path testable
without a subprocess. The hash is read and compared before anything starts, every time, because
Appendix A's "checked before every run" is a frequency and not an optimisation hint; a mismatch is a
failed check naming both hashes rather than a skipped run, since reporting `not-implemented` would let
a modified tool make a red verdict quieter. The contract is an object in — intent, phase, phase_dir,
artifacts_hash, root — and an object out carrying findings with a required cause, written in the
register because neither document specifies it and a project's tool has to agree with it. The exit
status decides and the findings do not: zero with findings is a pass carrying them, non-zero with none
gets a synthesised finding, because `Status` refuses a fail without one and a foreign tool must not be
able to make a verdict unrepresentable. Sixty seconds, stated in the register and in the configuration
comment. Rejected: a shell line instead of a path, which cannot be hashed, and no timeout at all,
which is a property nobody would have chosen.
