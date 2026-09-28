---
intent: github.com/triplem/xeno#118
phase: 04-verification
created: "2026-09-28T20:16:38Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+57e9207.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: a7a52416bb4401f127e8905056e3b59e521f77214e78b677c2046e93a1981c66
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

| Criterion | What proves it | Fails without the change |
|---|---|---|
| AC1 the listing exists and is ordered by `created` | `TestIntentsAreListedInTheOrderTheyWereCreated`, and the transcript's run over this repository | yes |
| AC2 the single intent form is unchanged | the transcript runs `intent status --intent XENO-0111` and prints six phase lines as before | no, and must not |
| AC3 XENO-0107 after XENO-0111 | the same test, with the four keys of this session, and the live run showing 0108, 0111, 0107, 0121 | yes, and it failed once during the work |
| AC4 an undated intent is listed with the reason | `TestAnIntentWithoutACreatedIsListedWithTheReason`, asserting it is present, carries a reason, and sorts before a dated one | yes |
| AC5 identical times are stable | `TestTheOrderIsStableForIdenticalTimes`, two runs compared, and the tie break asserted | yes |
| AC6 no intent directory is not an error | `TestNoIntentDirectoryListsNothingAndDoesNotFail` | yes |
| AC7 `CLAUDE.md` states the new rule | read: the paragraph gives the sequence, the startpoint, where the issue lives and why the old keys stay | no gate reaches it |
| AC8 nothing renamed, no verdict changed | `gate verify` over 79 verdicts, and this intent's key is `XENO-0200` | no, and must not |
| AC9 the usage names the optional argument | the transcript prints the usage line | no gate reaches it |
| the row names the furthest phase | `TestTheRowNamesTheFurthestPhaseWithAVerdict`, which is a design decision rather than an AC | yes |

Seven rows fail without the change, two must not, and two are read. AC3 is the one that earned its
place: it failed against the first implementation.

<!-- xeno:section:results -->
## Results

`go test ./...` passes on every package. `gofmt -l .` outside `vendor/` prints nothing,
`go vet ./...` is silent, and `./xeno gate verify` recomputes 79 verdicts and matches,
which is the evidence that nothing was renamed and no hash moved.

The listing over this repository puts the four intents of this session in the order they
happened: XENO-0108, XENO-0111, XENO-0107, XENO-0121. Sorted by key they read 0107,
0108, 0111, 0121. That difference is the issue, and the transcript shows both the
command and the tests agreeing on it.

`intent status --intent XENO-0111` prints its six phase lines and its next step exactly
as before, which is AC2, and the usage line names the argument as optional.

The result worth recording is a failure. AC3 failed against the first implementation,
which truncated `created` to a date inside the summary and then sorted the truncated
value, so two intents of one day fell through to the key tie break and XENO-0107 came
out before XENO-0108 — the order this listing exists to correct, reproduced by the
listing. The criterion named that pair, so it was caught in the phase rather than after
it.

Read rather than executed: the `CLAUDE.md` paragraph and the usage text. No gate in this
repository reads either.

<!-- xeno:section:gaps -->
## Gaps

**Two key schemes now live in one directory, and nothing enforces which is which.** The
gap at two hundred makes them legible to a reader and means nothing to the tool: the key
is opaque, so an intent created tomorrow as `XENO-0119` would be accepted. The rule is
in `CLAUDE.md` and in nothing that checks.

**Nothing allocates the next number.** Whoever creates an intent reads the listing and
adds one. A duplicate key would collide with an existing directory and be noticed; a gap
in the sequence would not be noticed at all, and neither would two intents created from
two branches taking the same number before either merged.

**AC7 and AC9 are proved by reading.** The convention paragraph and the usage line have
no gate, and the usage line in particular is the kind of text that goes stale first
because nothing fails when it does.

**The listing reports its own row's problem and no gate does.** An intent whose
`intent.yaml` cannot be read shows a reason in the listing and is otherwise invisible:
no gate looks at an intent directory outside a phase, which is #109, still open. This
change makes that state visible to a person and does not judge it.

**The order is the order of creation, not of work.** An intent created in the morning
and started a week later sorts by the morning. P2 rejected reading the first phase's
lock instead, because an intent with no phases would then have no place in the list, and
the cost is this imprecision.

**The evidence is a local run.** Fifth intent in a row, same reason.
