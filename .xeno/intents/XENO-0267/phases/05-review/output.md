---
intent: github.com/triplem/xeno#228
phase: 05-review
created: "2026-10-07T06:40:56Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b9e8695e59946aa922c6d82bb2740a86b6fd729aa052f9c7999b98ae1c85611b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'One, in P3. The register row A98 is open rather than approved: P1''s criterion 10 named what the row carries and not which state, and the state column of docs/assumptions.md reports what a person has said about a row. Nobody has said anything about either half, so approved would have put a yes in a person''s mouth. The row says what it asks for, a yes to a reading rather than a choice between paths, and the cost half carries accepted until WP20 inside it because that package can carry the figure whichever way the first half is answered. The deviation names criterion 10, which is what the rule asks.'
      result: deviation
      rule: deviations-are-traceable
    - note: 'No interface changes. No artifact gains a field, no gate gains a check, no command changes its behaviour or its refusals, and no exit code moves. Two suggestion strings get longer, which is output a person reads and not a surface anything calls: Suggestion keeps its four fields and --no-next silences the sentences as it silenced the text before them. gate verify reports 481 verdicts verified after the change, which is every sealed phase in the trail. There is nothing to migrate.'
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: None added. go.mod and go.sum are untouched and nothing compiles differently. The diff is two strings and two comments in internal/runner/next.go, two assertions in internal/runner/runner_test.go, and one row in docs/assumptions.md.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** No normative document is touched: not the process definition, not
the implementation plan. `docs/assumptions.md` is neither — its own opening paragraph declares
it open and states the test a row meets — and the row that went in is a decision of this intent
that outlives it. Nothing is invented: no field, no gate, no rule, no tool, and the MCP surface
is not reached at all. The change belongs to WP11 and to intent XENO-0267 for issue #228, and
the commit references it.

**What the issue asked for, and what it got.** Three things. The sentence after `phase finish`,
which is the `not-started` suggestion and now says it. The end-of-intent hint, which #228's
comment asked for as "just a hint, no requirement", and which the last phase's suggestion now
carries. And two decisions the issue left live, both taken and both recorded in A98 with the
clause behind each rather than as a preference. Nothing in the issue is left unanswered, which
is the test for closing it.

**Where this intent decided something the issue offered the other way.** #228's table lists the
plugin's skill text as the place for the harness specific command, "because the plugin is per
harness and may name one". In this tree it is not: section 13 ships one plugin, vendored and
hashed whole, and names working in both clients as the reason the skills exist rather than
subagents. That is a reading of a document the agent may not change and did not, and the row
says plainly that what it wants from a person is a yes to the reading. If the maintainer reads
section 13 the other way, the row is where the no goes and the skill text is a line's work.

**The honest state of the hint.** Nothing enforces it and the change says so three times: in
the suggestion's own comment, in the gaps section, and in A98. A reader who thinks a sentence
in `next.go` makes sessions get cleared has been told otherwise everywhere it could be said.
What would make it more is the figure, and the figure needs a baseline that WP20 owns.

**The question no rule asks: is the first sentence in the right place?** It hangs on the
`not-started` state, so it also appears after `intent start`, where the phase being started is
P0 and the session is usually the one the person just opened. There it reads as advice they have
already taken. The alternative was to make `Next` know which command called it, which would turn
a reading of section 6's table into something that reports its caller. A sentence that is
occasionally redundant is the cheaper of the two, and the redundant case is also the one where
it does least harm.

**The question after that: two sentences, or one.** They say the same thing for different
reasons — the first so the phase's context matches its lock, the second because nothing of a
finished intent is read again. One sentence covering both would have had to give the weaker
reason or neither, and a hint with no reason is followed once. Two is the right count and each
carries its own.

**What a reviewer should look at first.** The two strings, read as sentences rather than as a
diff, including the case where the staleness suffix is appended to the first one. That was
checked and reads as two sentences. After that, A98's second half, because scheduling something
to WP20 is the kind of answer that is right until nobody comes back to it.

<!-- xeno:section:release-notes -->
## Release notes

The next step a phase prints now says to start it in a fresh session, and says why: so that the
phase's context is what `context.lock.yaml` says it was given. The same suggestion still offers
`xeno phase start`, so nothing about what you type changes.

At the end of an intent, after P5 is decided, the suggestion adds that nothing of the intent's
context is read again — the next intent starts from the tree and the scope its own intake
declares — so the next one begins best in a fresh session.

Both are hints. Nothing enforces either, no gate reads them, no artifact records whether they
were followed, and `--no-next` silences them with the rest of the suggestion. Neither names a
harness's command, because the runner records which harness it is in and never behaves
differently for one: your client's own way of starting a session is yours to know.

`docs/assumptions.md` gains A98, which records the two decisions behind the shape of this: the
vendored plugin's skill text stays free of a harness command, because one plugin ships for both
clients and a skill is read by the model while clearing a session is the person's act; and the
figure that would measure whether anybody follows the hint belongs to WP20, which owns session
discipline and the baseline such a figure is read against.

<!-- xeno:section:residual-risk -->
## Residual risk

**The hint is unenforced and unmeasured, and that is the whole of the risk.** A phase started in
a session carrying four earlier phases goes green. The digest it was given stops being load
bearing without anything saying so, which is the state #228 described and which this change
makes visible rather than fixes. What would close it is the figure, and the figure is WP20's.
A98 is the only thing connecting the hint to that package, so the real risk is that nobody
reads the row.

**A sentence added to a suggestion is read until it is not.** Two suggestions now carry a
clause that nothing checks the usefulness of. If the advice turns out to be noise — if every
reader skips to the command — there is no signal that would say so, for the same reason there
is no signal that it works.

**The reading of section 13 could be wrong.** A98's first half rests on one plugin shipping for
both clients and on the lenses' stated reason for being skills. If the maintainer intends the
plugin to be built per harness — which the implementation plan leaves open, since it says which
of content, packaging and wiring WP11 builds per harness is "settled in dogfooding rather than
in advance" — then the skill text is the right place after all and the row's first half is a no.
The cost of being wrong is small and reversible: a line in a SKILL.md.

**Scheduling to WP20 is an answer with a shelf life.** The figure is correct to defer and
nothing makes it arrive. WP20 is late in the order — eleventh, after WP13 provides the token
baseline — so the gap stands for as long as that takes, and the row is the only record that
somebody decided it should.

**The absence of a harness command is checked over three trees and not over the documents.**
`internal/`, `cmd/` and `.xeno/plugin/`, which is the same reach as the `xeno.yml` check. A
`/clear` in a document would pass, and documents are where onboarding text will eventually go.
Not widened here, because widening a CI check is its own change with its own argument.

**For a person, not for the code.** Whether "fresh session" is the right phrase for a reader
who is not holding the specification. It names a thing every client has and no client calls
that, and the alternative — naming each client's command — is exactly what section 7 refuses.
