---
intent: github.com/triplem/xeno#228
phase: 01-requirements
created: "2026-10-06T20:27:34Z"
schema_version: "1.0"
runner_version: dev+3f1fab3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1f573db0af2bf07b90d09b3305f207aa78d2f5dd97fa98f4f6ba06b6afa3673c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Numbered, and P4's mapping cites the numbers.

1. **The `not-started` suggestion says the phase is started in a fresh session.** A person who
   has just finished a phase and reads the next step is told, in the same sentence, what to do
   before typing the command it offers.

2. **It says why**, in the terms the lock already uses: so that the phase's context is what
   `context.lock.yaml` says it was given. A hint whose reason is left out is followed once.

3. **It keeps `xeno phase start` as its command**, unchanged. The sentence changes what a
   person does before the command, not the command, and a suggestion that dropped the command
   would cost more than the sentence gains.

4. **The last phase's suggestion says the same for the session that is ending**, and keeps
   naming the merge as the step that is not a xeno command. The end of an intent is the second
   place #228 asks for, and its comment puts it at "just a hint, no requirement".

5. **Neither sentence names a harness or a harness's command.** No `/clear`, no Codex
   equivalent, and no branch on `XENO_HARNESS`. Section 7's constant stays in one place and in
   no condition, which the CI check in `xeno.yml` already asserts and which stays green.

6. **Neither sentence is an action and nothing enforces it.** A phase started in a reused
   session is judged exactly as before: no gate reads this, no field records it, and
   `--no-next` silences it as it silences every other suggestion.

7. **No new field, gate, rule or tool.** The change is prose inside two existing `Suggestion`
   values.

8. **A test asserts the fresh-session sentence on the `not-started` suggestion**, beside the
   existing assertions on the same suggestion, so that a later rewording cannot drop it
   silently.

9. **A test asserts it on the last phase's suggestion**, for the same reason and in the same
   place.

10. **`docs/assumptions.md` carries one row** with both decisions #228 asks for beyond the
    sentence: that the plugin's skill text stays free of a harness command, with the reason
    from section 13, and that the figure comparing declared context against spend is WP20's,
    with the reason from the plan's session discipline.

11. **`go test ./...`, `gofmt -l .`, `go vet ./...` and `xeno gate verify` all pass**, which
    is the project's own gate and the one evidence P4 attaches.

<!-- xeno:section:non-goals -->
## Non goals

**No document changes.** The suggestion machinery reports section 6's sequence and adds
nothing to it, so there is nothing the specification has to gain for the sentence to be true.
The first standing rule bars the agent from the documents in any case, and an intent that
needed one would have drafted the wording and put it to a person before starting.

**No harness command anywhere in this tree.** Not in the runner, which section 7 forbids, and
not in the plugin either. #228's table offers the skill text on the ground that the plugin is
per harness; section 13 ships one vendored plugin, hashed whole by G-Supply, and gives the
reason skills exist at all as working in both clients. A `/clear` there would be wrong under
Codex, and a skill is read by the model while clearing a session is the person's act.

**No hook.** #228 lists one and prices it in the same line: a hook that closes somebody's
session is the kind of help nobody asks for twice. Nothing here runs on a turn.

**No figure in the cost record.** The comparison between what `context.lock.yaml` declares and
what a session spent is WP20's, under session discipline, where the plan already puts "one
phase, one session" and a warning on a resumed session that has grown past a configured size.
WP20's own first rule is that a baseline comes before a lever, and this intent has no baseline
to read the figure against.

**No enforcement and no gate.** A phase started in a reused session goes green. That is what
makes the sentence a suggestion rather than a rule, and #228's comment asks for exactly that.

**No change to the suggestion for a running, red, provisional or stale phase.** Those name a
step inside a phase that is already open, and a session cannot usefully be cleared in the
middle of one: the artifact is half written and the context that would be cleared is the work.

**No sentence on `intent close`.** That command abandons an intent with a reason rather than
finishing one, and the end #228's comment means is the end of the work, which is the last
phase's suggestion. A hint about the next session on an abandonment would reach a person who
has just decided not to have one.

<!-- xeno:section:constraints -->
## Constraints

**Section 7's rule on the harness.** `XENO_HARNESS` is recorded and never branched on,
because "the moment the runner behaves differently per harness, the tools stop being
interchangeable". The CI check in `xeno.yml` asserts the name appears in one constant and in
no condition, and it stays green: nothing here reads the variable at all.

**What a suggestion may be.** `Suggestion`'s own comment binds this change: it is never an
action, it adds no field, no gate and no rule, and where the next step is not the runner's it
carries no command. Both sentences are inside those limits, and the second is in a suggestion
that already has no command.

**The first standing rule.** The documents are not the agent's to edit. `docs/assumptions.md`
is not one of them — it is the register of decisions taken while building, open by its own
first paragraph — and the row that lands there is a decision of this intent that outlives it,
which is the test that file states for a row.

**No invented fields, gates, tools or rules.** Nothing is added to an artifact, to the gate
list or to the MCP surface. The change is two strings and two assertions.

**One dependency.** Untouched. Nothing here compiles differently.

**Prose wraps at 88 characters**, tables and code blocks do not, and Go source has no width
rule beyond `gofmt`. The two suggestion strings are concatenations inside the existing style
of the file, which already wraps its text across lines that way.

**A paragraph changed a second time is replaced rather than edited into.** Both suggestion
texts are being changed, and both are read back whole afterwards as sentences rather than as
diffs, because each has to still read as one sentence a person acts on.

**Every change belongs to a work package and to an intent.** This is WP11, the agent layer,
and intent XENO-0267 for issue #228. The commit references the issue and closes it.
