---
intent: github.com/triplem/xeno#228
phase: 04-verification
created: "2026-10-07T06:38:35Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2086af031f4c57a41fa056ebc850343c8a7ef9a71e7402709d1c38891639c7ec
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - job: test
      kind: test-report
      path: evidence/go-test.txt
      produced_by: go test ./...
      result: pass
      sha256: 216e22a66c2324aaca7bdb774313442cbbc4cf522d06452505f82347698514cb
    - job: harness
      kind: other
      path: evidence/harness.txt
      produced_by: the two greps the xeno.yml check runs, with a positive control
      result: pass
      sha256: 4192cb41874420616b5b05464c953cc22eefb0d6719974a3da77cc2d0e8411dc
    - job: suggestions
      kind: other
      path: evidence/suggestions.txt
      produced_by: the two suggestions as the binary prints them
      result: pass
      sha256: 15bd707b61af0dd554f57209f1f0f09c6b0a9a3da2cde2767291c957171db68c
    - job: timing
      kind: other
      path: evidence/timing.txt
      produced_by: gate run on a throwaway copy, gate verify over 481 verdicts
      sha256: 00ba94599e0a6d388c601a4f5c46fe2454e63762437d64d7e990ad742d78f680
    - job: checks
      kind: build-log
      path: evidence/build.txt
      produced_by: go build, gofmt -l ., go vet ./..., xeno gate verify
      result: pass
      sha256: b0726e3c4189405e99ebdce771dc5f788230cb6006e1e4ef6e30c2e107b4e7af
    - job: mutation
      kind: test-report
      path: evidence/mutation.txt
      produced_by: each new string reverted, the two tests run, then restored
      result: pass
      sha256: f2a2ef59b1d7b476977e97ca97ccdf8cde31f25695bda077601b3dd46b977943
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Eleven criteria, by number, all passing. Two are held by a test, two by the binary's own
output, four by a grep with a positive control, and three by the project's check suite.

| # | What it asserts | What proves it |
|---|---|---|
| 1 | the `not-started` suggestion says the phase is started in a fresh session | `TestTheSuggestionMovesOnAndBackWithTheVerdict`, which now asserts "fresh session" in the text; and `evidence/suggestions.txt`, where the live P3 finish prints it |
| 2 | it says why, in the lock's terms | the same test, asserting "context.lock.yaml" and "was given" |
| 3 | it keeps `xeno phase start` as its command | the same test, whose first assertion on `s.Command` is unchanged and still passes |
| 4 | the last phase's suggestion says the same for the ending session | `TestAfterP5TheNextStepIsNotACommand`, asserting "fresh session" and "read again", and still asserting `s.Command == ""` and the merge; and `evidence/suggestions.txt`, where a decided P5 prints it |
| 5 | neither sentence names a harness or branches on one | `evidence/harness.txt`: the two greps `xeno.yml` runs, both clean, plus the test's own `/clear` assertion |
| 6 | neither is an action and nothing enforces it | the gate list of this phase is the same list as P3's, and `--no-next` is untouched; read in `cmd/xeno/main.go`'s `suggest`, which returns the code it was given |
| 7 | no new field, gate, rule or tool | `evidence/build.txt`: `xeno gate verify` recomputes 481 verdicts and all of them still match |
| 8 | a test holds the fresh-session sentence on the first suggestion | `go test ./internal/runner/` passes, and fails on a reverted string |
| 9 | a test holds it on the last phase's suggestion | the same |
| 10 | `docs/assumptions.md` carries one row with both decisions | A98 is in the file, with the five cells the table has, and `git diff` shows one line added |
| 11 | the project's own checks pass | `evidence/go-test.txt` and `evidence/build.txt` |

<!-- xeno:section:results -->
## Results

**The suite.** `go test ./...` passes: 20 packages, no failure, `internal/secrets` the slow one
at 132.7 s. `go build`, `gofmt -l .` outside `vendor/` and `go vet ./...` are clean.
`xeno gate verify` recomputes and matches 481 verdicts, which is every sealed phase in this
repository, so the change moves nothing that was already judged.

**The two sentences, as the binary prints them.** Not asserted from the strings but read off a
run. This intent's own P3 finish printed:

    next: start 04-verification in a fresh session, so that its context is what
    context.lock.yaml says it was given
      xeno phase start --intent XENO-0267 --phase 04

and a decided P5, run on a throwaway copy of the tree so that no sealed `gate.yaml` was
rewritten, printed:

    next: P5 is decided, so the merge is next. That is not a xeno command: commit, push, and
    let the review and the pipeline run. Nothing of this intent's context is read again, so
    the next one begins best in a fresh session

**The two tests.** Both pass, and both were checked the other way round rather than trusted:
each new string was reverted in `next.go`, the two tests run again, and both failed naming the
phrase that had gone — `evidence/mutation.txt` carries the output and the restored run. An
assertion that would pass either way is worth less than the line it costs.

**No harness is named, checked rather than asserted.** `grep -rn '/clear' internal/ cmd/
.xeno/plugin/` returns one line, and it is the test asserting the absence — no suggestion, no
skill and no other source carries it. The grep was run with a positive control in the same
file, `grep -rn 'fresh session'` over the same three trees, which returns six lines, so the
empty result above is an absence and not a grep that was looking in the wrong place. Both
checks `xeno.yml` runs were then run locally: `XENO_HARNESS"` outside the tests appears once,
in `internal/runner/runner.go:83`, and the comparison grep returns nothing.

**The register row.** A98 is in `docs/assumptions.md`, one line, with the five cells the table
has, immediately after A97 and before the `Decisions` heading. It went in on the wrong line
first — appended to A97's row, which is one very long line with no blank after it — and was
split onto its own line before anything was judged.

**Timings, for the figures #117 collects.** `gate run` for one phase: 23, 8 and 5 ms over three
runs on the throwaway copy. `gate verify` over the whole trail: 1832, 1654 and 1863 ms at 481
verdicts, which is 3.6 ms per verdict and in line with the 3.5 ms measured at 477.

<!-- xeno:section:gaps -->
## Gaps

**Nothing checks that anybody follows the hint.** This is the gap the change knowingly leaves,
and it is #228's own third question. A phase started in a session carrying four earlier phases
goes green exactly as before. The only signal is the cost ledger, which is gitignored local
data outside `artifacts_hash`, so no gate may read it; the figure that would make the gap
visible is WP20's and is recorded as such in A98.

**The sentence is not proved to be the right sentence.** It is proved to be printed, to carry
its reason and to name no harness. Whether a person reading "in a fresh session" does something
about it is an observation nobody has yet, here or anywhere: this repository is one developer
and the figure that would show it is the one not built.

**`/clear` appears once in the tree, and the absence is therefore narrower than it looks.** The
grep's one hit is the test asserting it is not in the suggestion text. So what is checked is
that no suggestion and no skill carries a harness command; what is not checked is that nobody
adds one to a document, since the greps cover `internal/`, `cmd/` and `.xeno/plugin/` and not
`docs/`. The `xeno.yml` check has the same reach and this intent did not widen it.

**The suggestions for a phase that is already open are untouched and unproved.** P1's non goals
say why nothing was added there, and no test asserts that nothing was: a later change that put
a fresh-session hint on the `running` suggestion would pass the suite. The reason it should not
is written in the design and in no assertion.

**The end-of-intent sentence was read on a copy, not on this intent.** A decided P5 is what
prints it, and this intent's own P5 is not decided while P4 is being written. The copy is this
tree with the new binary, so what was read is the real sentence from the real code path, but
the run that produced it rewrote a `gate.yaml` on the copy and that is why it was a copy.

**`intent close` says nothing about a session.** Deliberate, with the reason in P1: it abandons
an intent rather than finishing one. Nobody has checked whether a person abandoning an intent
would want the hint anyway.

**One figure is a measurement of this machine.** The `gate verify` total, 1654 to 1863 ms over
481 verdicts, is a timing on one laptop with a warm cache. It is reported for the series #117
keeps and means nothing on its own.
