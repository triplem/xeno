---
intent: github.com/triplem/xeno#215
phase: 04-verification
created: "2026-10-03T20:24:07Z"
schema_version: "1.0"
runner_version: dev+b626f1a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5f1f5a185b8e28ab580214e12a795419375ce60f809d36fd8a49651e2b36f8a0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Four claims, each mapped to a check.

"A judged phase is refused" maps to `TestStartIsRefusedWhereThePhaseHasAVerdict`, and to
running the real command against this intent's own sealed phase.

"The refused call writes nothing" maps to the same test, which hashes
`context.lock.yaml` before and after and compares.

"The refusal names both ways forward" maps to the same test, which looks for `section set`
and `gate.yaml` in the message rather than asserting the whole string, because the wording
should be free to improve.

"The dead-run recovery stays open" maps to
`TestStartAfterRemovingTheVerdictIsAllowed` and to `TestSecondStartIsRefused`, which was
already there and still passes — the marker refusal and the verdict refusal are different
preconditions and both are asserted.

The project's suite is run because the change is code.

<!-- xeno:section:results -->
## Results

Both new tests pass, and the package's existing tests pass unchanged, which is the part that
matters: a refusal added to a command every other test calls is a regression risk, and
`TestSecondStartIsRefused` and the dead-run path are the two it would have broken.

Against the real trail, `xeno phase start --intent XENO-0240 --phase 03` on a sealed phase:

    rc=1
    refused: 03-implementation has a verdict; redo the work with section set and
    phase finish, or remove .xeno/intents/XENO-0240/phases/03-implementation/
    gate.yaml to start it over

That is the call that rewrote a sealed lock twice during the past day's intents. `gate
verify` reports 319 verdicts at exit 0 afterwards, where before it would have reported a
divergence on that phase.

The suite: `go build` succeeds, `go test ./...` ok for every package, `go vet` silent,
`gofmt -l` outside `vendor/` empty.

<!-- xeno:section:gaps -->
## Gaps

This closes one way of rewriting a sealed artifact, not the class. A hand edit is the common
one and nothing refuses it — this session made two, and `gate verify` caught both. The
deletion half has the trail guard from #193. So section 11's clause now has three readers,
two of which act after the fact, and that is still the honest state.

The refused start is not the only path that overwrites a lock. `phase start` on a phase that
is running after its marker was removed by hand still rewrites the lock of an unfinished
phase, which is correct — nothing is sealed there — but it means the lock of a phase in
progress is not stable, and `ChangedSince` compares against the predecessor's lock rather
than the phase's own partly for that reason.

**`ChangedSince` has never had an input in this repository.** There is no
`context-profile.yaml`, `informationBase` returns nothing, and every lock's `files` list is
empty, so the changed set is silent for every phase of every intent in the trail. The
mechanism is built, tested, and unexercised outside its tests. It is a finding of this
intent's reading rather than of its change, and it gets its own issue; the figures are that
WP8's profile format is fixed in `model.Profile`, validated by G-Schema, and used by no
intent.

The `tool_version` and `review_checklist` hand edits into sealed frontmatter remain the
other side of this. A refusal protects what the runner writes; it cannot protect what the
runner has no command for, which is #188 and #208.
