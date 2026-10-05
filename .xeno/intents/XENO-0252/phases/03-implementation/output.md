---
intent: github.com/triplem/xeno#205
phase: 03-implementation
created: "2026-10-05T16:48:59Z"
schema_version: "1.0"
runner_version: dev+9eec4ed.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 41443ef3f4a9a6d642bfca421843e7e3e79a0065c90cfe90f5ef26ddae0cc79d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`internal/model/model.go` gains `LocalDataEnv`, `LocalDefault`, `LocalDir(root)` and
`LocalPath(root, parts...)`. `LocalDir` returns the variable's value where it is set and
non-empty, the default under the root where it is not; an absolute value as it stands, a relative
one joined to the root. It creates nothing.

The six literals are gone. `internal/runner/runner.go` resolves the run marker and `phase.env`
and exports `PhaseEnvFile`, because the command resolves the same file.
`internal/runner/enforcement.go` keeps `ReportFile` and gains `ReportPath(root)`;
`internal/cost/cost.go` keeps `LedgerFile` and gains `LedgerPath(root)`. Both constants were
paths and are now file names, which is the directory moving to the resolver.

`cmd/xeno/main.go` no longer repeats `.xeno/local/phase.env`; it calls
`model.LocalPath(o.root, runner.PhaseEnvFile)`. `printEnforcement` takes the root so it can print
where the report actually went rather than a constant that no longer is a path.

`internal/runner/init.go` is the one site that does not resolve, by criterion 5. It writes
`model.LocalDefault + "/"` into `.gitignore` with a comment saying why: the entry describes this
repository rather than where a person redirected their own state.

`internal/model/local_test.go`, new, five tests. Unset resolves to what the literals meant, per
path; an absolute value is used as it stands; a relative one is relative to the root; empty and
whitespace read as unset; resolving creates nothing.

`internal/runner/runner_test.go` gains three. Setting the variable moves the marker and
`phase.env` and leaves nothing under the default; unsetting it puts both exactly where every
existing repository has them; and the report and the ledger resolve into the same directory, so
one setting moves all four files rather than two.

`docs/assumptions.md` gains A97, which is the row a reader needs when they meet
`XENO_PLUGIN_DATA` read and `XENO_PLUGIN_ROOT` removed: the difference is that nothing under the
local data location is hashed, with A89's measurement as the contrast.

It was also run rather than only tested. In a scratch repository with the variable set,
`intent start` and `phase start` wrote `phase.env` and `runs/DEMO-0001/00-intake.lock` under the
redirected directory and nothing under `.xeno/local`, which is the behaviour the export has
promised since it was written.

<!-- xeno:section:deviations -->
## Deviations from the design

One addition beyond the design, which named four files and touched five. `printEnforcement` in
`cmd/xeno/main.go` printed `runner.ReportPath` as a constant and now takes the root, because once
the path is resolved a function cannot print it without knowing which repository it is for. The
design said `ReportPath` becomes a function and did not notice that its one caller outside the
package was a printer with no root in scope.

One departure from P2's wording rather than its substance. The design said `Ledger` and
`ReportPath` "stop being path constants and become base names", and they did — but each also
gained a resolver function beside it, `LedgerPath` and `ReportPath`, rather than every caller
calling `model.LocalPath` with a file name. Two callers each would otherwise have repeated the
file name, which is the duplication this intent exists to remove.

One thing the design got right for a reason it did not state. `LocalDir` trims whitespace before
testing for empty, which P2 justified as "an exported-but-empty variable is what
`export XENO_PLUGIN_DATA=` produces". The test covers a space and a tab as well, because a value
built from a failed shell substitution is as likely to be whitespace as nothing, and the
resolver's behaviour should not depend on which.

No deviation on compatibility, and it is asserted twice rather than argued. `TestUnsetResolves­ToWhatTheLiteralsMeant`
checks all four paths against the strings the literals produced, and
`TestWithTheVariableUnsetTheLocalPathsDoNotMove` runs `phase start` and looks at the disk. The
second is the one that would catch a resolver that was right about strings and wrong about where
a writer puts things.

Nothing else departs. The gitignore site is unchanged by criterion 5, no normative document is
touched, the trail is untouched, and `gate verify` is at exit 0 over 390 verdicts.
