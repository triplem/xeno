---
intent: github.com/triplem/xeno#277
phase: 04-verification
created: "2026-10-07T13:35:44Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 87103ee9eda9d4f77bc202e4fb4e10aeb0a21d7c7ceadca8c4c916846cb207bb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - kind: test-report
      result: pass
      produced_by: go test ./...
      sha256: 2b32de1c136473a1e83588c0668ef489df02b9b926e189b4ed29c59678a83f1b
      path: evidence/go-test.txt
      job: test
    - kind: build-log
      result: pass
      produced_by: the greps with their control, honest before and after, the inverted test case, go build, gofmt, go vet, xeno gate verify
      sha256: 8c5afa328484ae4e168e5ed5acaba83b0fe253b1aaa00739ef6239b2499f787d
      path: evidence/checks.txt
      job: checks
    - kind: other
      result: pass
      produced_by: 'the four artifacts that declare tool: manual, their placeholder fields and template refs, and the count of phases with no lock'
      sha256: 36c72a8fda814289f5a23c7fb81176549871759fa76d9b9bac3aadf5cc3714f2
      path: evidence/measurement.txt
      job: measurement
    - kind: other
      result: pass
      produced_by: the two gate files after the four approvals, and gate verify's exit code
      sha256: e0461c96d4bc5111cd7fdc974eeda1445b0c770e1638705f0888025b5e4f4384
      path: evidence/releases.txt
      job: releases
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Thirteen criteria, by number, all passing. Nine are checks and four are read in the files.

| # | What it asserts | What proves it |
|---|---|---|
| 1 | `honest` no longer reads the `tool` field | `evidence/checks.txt`: the expression before and after, and `grep 'raw["tool"]'` returning the line on `main` and nothing here |
| 2 | no gate reads any of the triple to decide anything | `evidence/checks.txt`: the grep over `internal/gates/*.go` returns one line, `sessionFields`, which asserts presence and reads no value |
| 3 | the comment says what was removed and why | read in `internal/gates/gates.go` above `honest` |
| 4 | `hashShape`'s doc comment describes what exists | read: it names what the two remaining cases have in common and no longer mentions `tool: manual` |
| 5 | the inverted test case passes and the other five are unchanged | `evidence/go-test.txt`, and the failure the old assertion produced before it was inverted |
| 6 | exactly two phases change verdict on exactly four findings | `evidence/measurement.txt` for which four artifacts could be affected and why no others, `evidence/releases.txt` for the two gate files |
| 7 | both phases read `approved` | `evidence/releases.txt`: `status: approved` in each, with each finding's decision, type and the name |
| 8 | the releases are approvals, not overrides | `evidence/releases.txt`: four `type: approved` and zero occurrences of `obligation` across both files |
| 9 | `xeno gate verify` exits zero | `evidence/checks.txt` and `evidence/releases.txt`, which also records that it exited 1 between the code change and the releases |
| 10 | no artifact's content is edited | `git status`: the two modified files under `.xeno/intents/XENO-1` and `XENO-2` are `gate.yaml` and nothing else |
| 11 | no normative document changes | `git status`: `docs/assumptions.md` is the only file under `docs/` |
| 12 | the register carries one row | A101, five cells, state naming the maintainer and 2026-10-07 |
| 13 | the suite and the static checks pass | `evidence/go-test.txt`, `evidence/checks.txt` |

The measurement the decision rests on is in `evidence/measurement.txt`: the four artifacts that
declare `tool: manual` with the placeholder fields each carries, their `template: intake@0.1.0`
against the shipped `intake@1.0.0`, and the count of phases with no `context.lock.yaml`, which
is zero and is what ruled out the replacement term.

<!-- xeno:section:results -->
## Results

**`honest` no longer reads the field, shown both ways.** `grep 'raw["tool"]'` over
`internal/gates/gates.go` returns nothing here and returns `byHand := raw["tool"] == "manual"`
on `main`, which is the control: the pattern finds the line where it exists. The expression
itself is printed before and after in `evidence/checks.txt`.

**No gate reads any of the triple to decide anything.** The grep for `"tool"`, `"model"` and
`"tool_version"` across `internal/gates/*.go` outside the tests returns one line,
`sessionFields`, which is the required-field list: it asserts presence and reads no value. That
is the claim the whole change rests on and it is an absence, so it was established by the grep
with its control rather than by reading.

**What the term did, measured before it went.** Four artifacts in a trail of 499 phases declare
`tool: manual`: the `output.md` and `digest.md` of XENO-1's and XENO-2's intakes. The two
`output.md` files carry the placeholder in all four hash fields and the two `digest.md` files in
two. `secrets_hash` and `rules_hash` are in `model.WriterlessHash`, so they never needed the
`tool` term; both `output.md` files record `template: intake@0.1.0` against a shipped
`intake@1.0.0`, so `goneBundle` covers their `strings_hash`. That leaves `context_hash`, and
four findings is what removing the term produced.

**The replacement was measured and found empty.** Keying the exemption on a phase having no
`context.lock.yaml` to hash would have been a fact about the tree rather than a declaration.
Every phase in the trail carries a lock — 495 when the term was measured, 499 once this intent's
own phases were sealed — so the term would never have fired. That count is the load bearing
fact behind the decision and P3's first deviation records that it was taken while writing the
design rather than before it.

**The four releases, as the gate files record them.** `evidence/releases.txt` prints each:
`status: approved` on both phases, four findings, four decisions of `type: approved`, each with
`by: Markus M. May`, and zero occurrences of `obligation` across both files, which is what an
approval leaves. `xeno gate verify` exits zero again; it exited 1 between the code change and
the releases, which is the state the releases exist to resolve.

**Both test tables were checked the other way round.** Leaving the gates case as it was, after
the term was removed, fails naming the case and printing both causes; the runner case fails
with "got red, wanted green". `evidence/checks.txt` carries the first. The second is the
deviation in P3: `go test ./internal/gates/` passed with only the first inverted, so the
package under change was green while `go test ./...` was not.

**Nothing in the four artifacts is edited.** `git status` shows `gate.yaml` as the only modified
file under each of `XENO-1` and `XENO-2`. Each artifact still reads `context_hash: by-hand`.

**One thing about how this phase was run.** `gate run` was used on P4 while it was still running,
to read the staleness finding P3's re-judging produced. That wrote a verdict, which made
`phase start` refuse, so the phase was removed and started over; its `context.lock.yaml` is
therefore newer than its siblings' and records the predecessor as re-judged. G-Freshness passes.

**The checks.** `go test ./...` passes across 20 packages. `go build`, `gofmt -l .` outside
`vendor/` and `go vet ./...` are clean. `xeno gate verify` matches 499 verdicts and exits zero.

<!-- xeno:section:gaps -->
## Gaps

**Two sealed phases now read `approved` rather than `green`, permanently.** XENO-1's and
XENO-2's intakes carry a finding each person released with a reason, and nothing will ever
close them, because the artifacts cannot be corrected. Anything that counts green phases in
this trail counts two fewer from now on, and the reason is a change made after the fact rather
than anything those two intents did.

**The gate is stricter and the trail is the only evidence of how often that matters.** An
artifact written by hand with `context_hash: by-hand` is a finding now. Twice in 499 phases is
the entire sample, and both are from the same two days before M0. A project that writes
artifacts by hand more often meets this immediately and nothing here tells them how often that
is.

**The four approvals are a person's statement typed by the agent.** The maintainer chose the
shape and approved the wording, and the commands were run with `--by` naming them. That is the
same arrangement as the specification passages written on instruction earlier today, and it is
further from the thing the process protects: `next.go` refuses to offer `gate approve` as a
command precisely so that a runner does not nudge towards the decision it exists to record, and
here the agent typed it. The record says who decided; it does not say who typed.

**Nothing stops the term coming back.** No test asserts that `hashes` reads no field of the
triple. The two inverted cases assert the consequence for `context_hash`, which is the right
place for them, but somebody adding a new term from a declared field would pass the suite. The
rule is in a comment and in A101.

**`goneBundle` is still a term that can be wrong.** It is a fact about the tree, which is the
property the change establishes, and it is still unable to tell a bundle that is gone from one
that was renamed — #273 measured that and left it standing. So the rule "only facts about the
tree" is met while one of the two facts is coarser than it looks.

**The measurement is of this repository.** Four artifacts, two phases, one project. The claim
that removing the term moves nothing else is a statement about this trail, and a project with
its own trail would have to run `gate verify` before and after to know its own number.

**P4 was started over and its lock is newer than its siblings'.** Caused by running `gate run`
on a phase that was still running, in order to read a staleness finding. G-Freshness passes and
the phase is honest, but the sequence of events is not recoverable from the artifacts: what the
trail shows is a verification phase whose lock postdates the implementation it verifies, with
no record of why beyond this paragraph.
