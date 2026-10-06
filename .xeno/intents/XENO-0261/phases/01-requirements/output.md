---
intent: github.com/triplem/xeno#237
phase: 01-requirements
created: "2026-10-06T12:03:04Z"
schema_version: "1.0"
runner_version: dev+7885661.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: adab08c7c2c7972f99b12610f2b61da9cce94ac2089d7f7cced670dfa3d7fde9
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Numbered, and P4's mapping cites the numbers.

1. **The two remedies about a sealed lock name section 7's two routes.** Starting the earlier phase
   over, and a second person approving the finding on the phase being gated as still valid. Those are
   the clause's own two — "a stale phase is not deleted, it is re-run or explicitly approved as still
   valid" — and both are acts the runner carries out.

2. **Each of the two says why `section set` and `phase finish` will not clear it.** That is the route
   the refusal offers first and the one a person reaches for, and the reason it fails is not
   guessable: the lock is written only by `phase start` and keeps what the phase was given.

3. **Each names which phase to act on.** The finding is raised while gating one phase and is about
   the lock of an earlier one, so the remedy names the earlier phase for the re-run and the gated
   phase for the approval. A remedy naming neither leaves a person to work out which of two phases
   `gate approve` wants.

4. **The re-run route says what it costs.** Starting the earlier phase over discards it, and every
   later phase built on its verdict. A remedy offering it without saying so is the same trap one step
   further along, which is the fault being repaired.

5. **The approval is phrased as a second person's act, not as a command handed to the reader.**
   `Runner.red` in `next.go` already names `gate approve` and its comment says why it offers neither
   way out as a command: releasing a finding is a statement by a second person, and a runner that put
   the command in the reader's hand would nudge towards the decision it exists to record. The remedy
   has to be true for a reader of `gate.yaml`, who sees no suggestion, without contradicting that.

6. **The unreadable-file remedy keeps "make it readable" and gains a clause saying why it differs.**
   That fault is this machine's — a permission, a broken link, a filesystem — and the next `gate run`
   clears the finding with nobody approving anything. It is the one of the three whose remedy already
   worked, and a reader meeting three findings should not read the odd one as an oversight.

7. **A test asserts what each of the three remedies names.** Three tests cover the two limits and
   none covers a remedy, which is how three strings came to name an act the runner declines. The new
   tests name the routes rather than matching the whole string, so that rewording does not break them
   and removing a route does.

8. **No condition of the check changes.** The loop bound, the `touched` guard, the three cases and
   their causes are untouched. #236 settled when the check fires and this intent changes only what it
   says to do. `git diff` over `staleReads` shows changes inside the three `finding` calls and
   nowhere else.

9. **`go test ./...` passes, `go vet` is clean, `gofmt -l` prints nothing, `xeno gate verify` exits
   0.** The first is the one that matters here, because this intent adds tests.

10. **No normative document is touched.** Section 7 already names both routes, so this makes a string
    agree with the specification rather than asking the specification to change.

<!-- xeno:section:non-goals -->
## Non goals

**Not teaching `phase start` to re-resolve a lock.** It would make the old remedy literally true, and
#215 refused it for two reasons the code restates at `runner.go:409`: the lock is inside
`artifacts_hash`, so a second start rewrites a sealed artifact, and the lock is the only record of
what the phase was given. Reopening that is a section 11 question and a person's commit before any
code.

**Not dropping the instruction and stating only the fact.** It was the third option and it is the
cheapest. Every other refusal and finding in this runner carries a remedy that works, which #216 and
#218 were careful about, so a finding with no remedy replaces one exception with another and leaves a
person to find the release route unaided — which is the fault the issue reports.

**Not #235.** The other wp8 issue is that `result` turns any finding into a red gate, so a budget
overrun stops a phase where section 5 says it must not. It touches this one only in that a red
staleness finding is why an approval is needed at all. Fixing either does not move the other, and
#235 needs a specification commit first.

**Not changing when the check fires.** #236 settled both of section 5's limits and they are in the
code. No loop bound, no guard, no comparison and no cause string changes here.

**Not a new finding, a new cause or a new field.** Three findings in, three out, with the same causes.
The second standing rule makes an addition a specification change, and nothing here needs one.

**Not the other two remedies in the file.** `digestFindings`, `budget`, `links` and the rest print
their own, and a pass over all of them for the same fault is #202's shape of work and an intent of
its own. This one repairs the three the issue names.

**Not `Runner.red`'s sentence.** It already names both ways out and does it deliberately. Changing it
to avoid saying the release twice would be solving the duplication in the wrong place: a reader of
`gate.yaml` has no suggestion at all, and that reader is the one the remedy exists for.

<!-- xeno:section:constraints -->
## Constraints

**The remedy is read in three shapes and has to work in all of them.** `cmd/xeno/main.go:898` prints
it under the finding in the gate's own output; `next.go:142` composes it into "…: <next>. Then judge
it again. Or a second person releases it, with gate approve or gate override, naming <id> and a
reason"; and `gate.yaml` carries it as text with nothing around it. The third reader is the one that
decides the wording, because it is the only one with no sentence beside it, and the second is why the
approval is phrased as a second person's act rather than as a command.

**The composed sentence adds "Then judge it again" after the remedy.** So a remedy ending in an
instruction has to be one that judging again would follow sensibly. "Start the earlier phase over"
does; a remedy ending in "approve it" would read as judging again after an approval, which is the
wrong order.

**The finding is about one phase and raised while gating another.** `staleReads(c, idx)` loops the
locks before `idx` and the finding's `file` is the earlier phase's lock, while the check belongs to
`c.Phase`. Both phases have to appear in the remedy and the strings have to say which is which, or
the remedy is ambiguous in exactly the way that costs a person two refusals.

**An approval releases rather than removes.** Section 5's derivation makes a phase with every failure
approved `approved`, not green, and A12 lets the next phase start on a decided predecessor. So the
remedy can promise that the work proceeds and must not imply the finding disappears.

**The strings are built with the phase names in them.** Both the earlier phase and `c.Phase` are
available in `staleReads`, so the remedies are formatted rather than fixed, and a test that matched
the whole string would then be asserting the format. The tests name the routes and the phases
instead.

**88 columns and `gofmt`.** The remedies are Go string literals in a file with no width rule beyond
`gofmt`, so the constraint is only that the concatenation stays readable; the prose of these
artifacts wraps at 88.
