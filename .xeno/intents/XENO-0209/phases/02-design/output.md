---
intent: github.com/triplem/xeno#109
phase: 02-design
created: "2026-09-29T18:36:48Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+3ec2429.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 77580eddb091f76d63047840d9ebeffa731ea1c122ae20ad387f344286780cd2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: by-hand
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**`model.KnownIntentFiles`, beside `KnownPhaseFiles`.** Two maps in one place, so a
reader comparing the levels finds both and section 4's two lists are transcribed once
each.

**`intentDirectoryFindings(root, key)` in `internal/gates`, mirroring
`directoryFindings`.** Same shape: read the directory, allow the one legitimate
subdirectory, allow the enumerated files, report the rest. Same two wordings, one for a
file and one for a directory, so that the levels read alike and a finding says which
level it came from by naming its path.

**Called from `CompleteOnClose`, first.** Before the reason and the learning record, so
that a verdict lists what the directory contained before what the intent asserted. Order
is presentation here and the choice is deliberate rather than accidental.

**`phases` is the one allowed directory**, as `evidence` is at the phase level. Named
against `model.PhasesDir` if one existed and against the literal if not, which is what
the phase level does with `evidence`.

**An unreadable directory reports nothing.** `os.ReadDir`'s error is dropped, exactly as
`directoryFindings` drops it: a directory that cannot be read is a finding somebody else
already has, and `CompleteOnClose` reading `intent.yaml` from inside it will say so.

**The hash is untouched.** `hashing.DirHash` and `IntentExcluded` do not change, because
the gap was never in the hash: it was that nothing said what the hash was allowed to
cover.

<!-- xeno:section:alternatives -->
## Alternatives

**Give G-Schema an intent mode.** The other option #109 names. `schema` takes a `Ctx`
carrying a phase and every check it calls reads a phase directory, so the mode would be
a second function sharing a name rather than a gate doing two jobs, and the finding
would then arrive from a gate that reports on phases in a verdict that has none.
G-Complete already has two modes for the two ways an intent ends, and this is one of
them.

**Check it in `IntentClose` rather than in a gate.** The runner computes the hash there,
so the check would sit beside the thing it guards. It would also make the runner report
findings, which is the gates' job, and a command that refused would prevent an
abandonment — which section 6 says is written rather than refused, because the intent is
dropped either way and the record should say what is missing.

**Refuse to hash a directory carrying something unknown.** It would make Appendix B's
property true by construction instead of by report. It would also mean an intent could
not be closed until somebody tidied it, and section 6's argument about abandonment is
that the record is written and the gate says what is wrong. A specification change, not
a code one.

**Skip directories, since the hash does not descend.** One line shorter and it leaves a
directory nobody wrote inside a judged one, where the next file will land unjudged. The
phase level check reports them and this follows it.

**Extend `KnownPhaseFiles` with the intent level names and use one map.** Fewer maps. It
would accept `output.md` in an intent directory and `assumptions.yaml` in a phase, so
both checks would stop distinguishing the two levels section 4 distinguishes.

<!-- xeno:section:impact -->
## Impact

`internal/model`: `KnownIntentFiles`, four names, beside the phase map.

`internal/gates`: `intentDirectoryFindings`, and one call at the top of
`CompleteOnClose`.

`internal/gates` tests: a stray file, a stray directory, a clean intent directory, and
that `phases/` and the four known files are silent.

Nothing else. The hash, its exclusions, the phase level check and every other gate are
unchanged.

What a reader gains: Appendix B's sentence about a checked property is now true of both
hashes it covers.

What they do not gain: prevention. A stray file is reported by the run whose hash
already covers it, which is what the phase level has always done.

What this costs: nothing measurable. One map, one function of a dozen lines, one call,
and a guard that does nothing until an intent is abandoned — which has not happened here
yet.
