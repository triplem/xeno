---
intent: github.com/triplem/xeno#183
phase: 03-implementation
created: "2026-10-03T18:39:48Z"
schema_version: "1.0"
runner_version: dev+0462eb7.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 33771f44c24e934b919254b45a12555ceeb5c9c098e49b5e01230669c3d5407c
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

**The specification, committed alone and first.** Section 7's list drops `XENO_PLUGIN_ROOT` and keeps
three. The resolution order paragraph becomes four: the plugin is the vendored one found from the git
root and nothing else; the order is removed rather than built, with the gate path's dependency on the
rule set and the templates and the silent-unjudging measurement; the client fallback is unreachable,
with the no-template measurement; and the condition an override would have to meet. The plan's WP7
sentence names three variables and says what was on it.

**`internal/plugin/plugin.go`.** The package comment said the order was "not implemented here on
purpose" and called it "a decision for section 7, not something to build quietly". It now says the
order is gone rather than unimplemented and keeps the measurement, because the measurement is why.

**`.xeno/plugin/bin/xeno-env.sh`.** "XENO_PLUGIN_ROOT is deliberately not set here, and that is the
one thing this script does not do" becomes "There is no XENO_PLUGIN_ROOT to set", with the same
reasoning behind it.

**`internal/plugin/plugin_test.go`.** `TestTheEntryPointStartsTheRunnerFromThePath` keeps its
assertion that the script exports no plugin root and gains a comment saying why it keeps it: the
thing worth preventing is a root arriving from the environment, and that does not stop being worth
preventing when the document stops naming a way to do it.

**`ASSUMPTIONS.md`.** A84's status closed — it said the resolution order was its open clause and the
decision was section 7's. A89 added, with both measurements, the three shapes and which was chosen.

**Both files' comments rewrapped to 88 columns.** They were written at 95 in an earlier intent of
mine; nothing checks the width, so it went unnoticed until these files were open for another reason.

**No change to any Go behaviour.** `internal/rules`, `internal/template`, `internal/secrets` and
`internal/plugin` resolve from the vendored directory exactly as before, and that is the whole reason
the first commit carries no code.

**Checked for stragglers.** `grep -rn 'XENO_PLUGIN_ROOT\|plugin-root'` over `docs/`,
`ASSUMPTIONS.md`, `README.md`, `internal/` and the entry point: five files, every one of them
explaining the removal, and none claiming the mechanism. The plan has none left.

<!-- xeno:section:deviations -->
## Deviations from the design

**One the design named and one it did not.** The rewrap to 88 columns was recorded in the design, as
something done because the files were open. What the design did not anticipate is how much of
`internal/plugin/plugin.go` it touched: the file was written at about 95 columns throughout, in an
intent of mine three changes ago, so the rewrap reflows comments that have nothing to do with this
change. It is in the diff and it is noise around the two paragraphs that matter.

The alternative was leaving a file that breaks a convention this project states in `CLAUDE.md`,
while editing it for something else. Nothing enforces the width — there is no check for it, which is
why it went unnoticed for three changes — so the choice was between a wider diff now and a file that
stays wrong until somebody notices again.

**One thing worth saying about the shape of this intent.** The specification change is the whole of
the work and the code change is comments. That is unusual here and it is what the first standing rule
produces when the thing being corrected is a document rather than an implementation: the ordering
rule says the document moves first, and when nothing implemented the clause, the document moving is
all that happens.
