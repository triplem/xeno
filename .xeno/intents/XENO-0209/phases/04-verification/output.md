---
intent: github.com/triplem/xeno#109
phase: 04-verification
created: "2026-09-29T18:44:46Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+3ec2429.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2f7660d423fa7a4707b7ff9b79c5501a304507fb4c5d540899837a8cadda4adb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
evidence:
  - kind: test-report
    job: go-test
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | What proves it |
|---|---|
| AC1 `KnownIntentFiles` is section 4's four | `TestTheFilesSectionFourNamesAreSilent`, which asserts the count as well as the silence |
| AC2 a stray file is a finding | `TestAStrayFileInTheIntentDirectoryIsAFinding`, and the two-binary run: the same intent is green on `main` and red here |
| AC3 a stray directory has its own wording | `TestAStrayDirectoryIsAFindingOfItsOwn`, which also asserts it is not reported as a file |
| AC4 the known files and `phases/` are silent | the same test as AC1, including the `gate.yaml` a second close finds |
| AC5 the finding lands in the close's verdict and turns it red | the transcript: `status: red` in the `gate.yaml` the new binary wrote, with both findings in it |
| AC6 the phase check, the hash and its exclusions are unchanged | the transcript prints the same `artifacts_hash` from both binaries, `bf49ef35…`; the diff touches neither `hashing` nor `directoryFindings` |
| AC7 an unreadable directory is no worse than before | read: `os.ReadDir`'s error is dropped as `directoryFindings` drops it, and `CompleteOnClose` reads `intent.yaml` from inside it next |
| AC8 everything green stays green | `go test ./...`, `gofmt`, `go vet`, `gate verify` over 133; no intent here is abandoned so none gains a finding |

Eight rows. AC6's evidence is the one worth reading twice: the hash is byte for byte the same on both
sides, and the verdict over it went from green to red. The hash was never wrong; nothing judged what it
covered.

<!-- xeno:section:results -->
## Results

Five tests pass, the suite passes, `gofmt` and `go vet` are clean, and `gate verify`
matches 133 verdicts.

The before and after is two binaries rather than two test runs, because the new tests
name `model.KnownIntentFiles` and `model.PhasesDir` and cannot compile against the old
tree. XENO-0201 met the same wall with a changed signature and had to substitute
evidence from the tree; a binary has no compile time dependency on a symbol, so here the
comparison is real.

What it shows: the same scratch intent, carrying a stray file and a stray directory,
closes green on `main` and red here. Both verdicts record `artifacts_hash: bf49ef35…`,
identical, which is the whole argument of #109 in one line — the hash covered the stray
file either way, and until now nothing said so.

**A second defect surfaced and had to be fixed for AC5 to hold.** Two findings on a
close made `Status` refuse with `finding id collision`, because `CompleteOnClose`'s
findings had never passed through `carryForward` and both ids were empty. No verdict was
written at all. A25 promises every finding reaches a verdict through that one function
and `Invariants` exists to check the promise; `IntentClose` was written afterwards and
met neither, which is exactly the second path into a verdict that `Invariants`' own
comment says would break the rule. It is #138, it is fixed here, and
`TestFindingsOnACloseCarryTheirIds` asserts both ids and the red status.

Read rather than executed: that the hash is untouched, and that an unreadable directory
behaves as it did.

<!-- xeno:section:gaps -->
## Gaps

**Reported, not prevented.** A stray file still enters the hash of the run that reports
it. Section 4's arrangement is that the verdict records both, and the transcript shows
it: the red verdict carries the same hash the green one did. Appendix B's property is
now checked at both levels and checked means said aloud.

**No intent here is abandoned, so this has judged nothing real.** It is a guard placed
before its data, like the refusal XENO-0107 put in `Status`, and the first abandonment
is when it earns its keep. The two-binary run is a scratch intent rather than one of
this repository's.

**#138 rode along and is a separate issue.** It had to be fixed for AC5, since without
it two findings produce an error and no verdict, and it is not #109's subject. It is
filed, its commit is its own, and the deviations say why it is here.

**Nothing checks that the two levels stay in step with section 4.** Both lists are
transcribed by hand from a document no gate reads. If section 4 gains a file at either
level, the check will report it as unknown, which is the failure in the safe direction,
and nobody will notice until it happens.

**The wordings are duplicated across the two levels.** Four strings, two per level,
differing only in the words "phase" and "intent". A reader comparing them sees the
parallel; a change to one will not change the other.

**The evidence is a local run**, fourteenth in a row.
