---
intent: github.com/triplem/xeno#228
phase: 02-design
created: "2026-10-06T20:28:48Z"
schema_version: "1.0"
runner_version: dev+3f1fab3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fec25d92f5563bdedd73c107a1312efed0de0812ecfc6d8968d1acd269899d1a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The sentence hangs on the state, not on the command that asked.** `Next` answers from the
phase states and knows nothing about who called it, so a sentence that appeared only after
`phase finish` would need the caller to pass that in, and `Suggestion` would stop being a
reading of section 6's table. The `not-started` answer is the one a person sees after a finish
that went green, and it is also what they see after `intent start`, where starting P0 in a
session of its own is the same advice for the same reason. So the sentence goes on the state
and reaches both.

**It says why in the lock's own terms.** "so that its context is what `context.lock.yaml` says
it was given" is the sentence the issue drafted, and it is the right one because the reason is
checkable against a file the person already has. A hint that said only "start a fresh session"
would be followed once and then skipped, since nothing enforces it and the cost of ignoring it
is invisible.

**The command stays.** `xeno phase start` is still what the person types; the sentence is about
what they do first. Dropping the command to make the sentence land would trade a step a person
needs every time against a reminder they need until it is habit.

**The end-of-intent sentence names what is not read again.** The last phase's suggestion
already carries no command and already says the next step is somebody else's — commit, push,
let the review and the pipeline run. The addition says that nothing of this intent's context is
read again, because the next intent starts from the tree and the scope its own intake declares.
That is the reason a person can check, and it is the one #228's comment asks for as a hint
rather than a requirement.

**Neither sentence reads `XENO_HARNESS`.** Section 7 is not a style rule here: the variable is
not consulted, there is no condition to add, and the CI check that asserts the constant appears
once and in no condition is untouched. The sentences are the same under Claude Code, under
Codex, and for a person working with commands alone, which is the third reader the plugin
cannot reach at all.

**The skill text gains nothing.** The decision is recorded rather than left, because #228 puts
it as a live option. A skill is read by the model and clearing a session is the person's act,
so the sentence would be addressed to a reader who cannot carry it out. And the plugin is one
vendored tree, hashed whole by G-Supply and shipped for both clients, so a harness's command in
it is wrong under the other one. The runner's sentence reaches every reader the plugin does and
the one it does not.

**The cost figure is WP20's and is recorded as scheduled rather than refused.** The gap #228
describes is already visible: `internal/cost/cost.go` records a turn with no phase open as
`none` rather than attributing it, and its package comment carries the seven per cent a phase's
own window captures. What is missing is a baseline to read a figure against, and the plan puts
both the baseline and session discipline in WP20, where the warning on a resumed session is
already written down. Building the figure here would be the lever before the measurement, which
is the one order WP20 fixes.

**One row in `docs/assumptions.md` carries both.** They are decisions of this intent that
outlive it: the next person who reaches for the skill text or for the cost record finds the
reason rather than the absence. A decision sealed only in this intent's P2 would be read by
somebody who already knew to look here.

<!-- xeno:section:alternatives -->
## Alternatives

**The harness's own command in the plugin's skill text.** #228's table offers it, and it is
the version a person would most readily act on: `/clear` is one keystroke sequence and "a fresh
session" is a thing to work out. It costs the interchangeability the skills exist for. Section
13 ships one plugin, vendored under `.xeno/plugin/` and hashed whole, and gives as the reason
lenses are skills rather than subagents that they "work in both clients"; a `/clear` in that
tree is wrong under Codex and reaches nobody working with commands alone. It also addresses the
model, which cannot clear the session it is in. Refused, and the reason is recorded in the
register because the option will be reached for again.

**A hook that ends the session at `phase finish`.** The only form that would actually hold,
since nothing else here does. #228 prices it in its own table: a hook that closes somebody's
session is the kind of help nobody asks for twice. It also breaks the suggestion's rule that it
is never an action, and it would end the session before the person has read what the gate said.

**A figure in the cost record now, rather than in WP20.** The ledger has what it needs: lines
naming `none` are turns spent with no phase open, and a phase's total is already the sum of
differences over the lines naming it. A figure comparing that against what `context.lock.yaml`
declared could be printed without a new field. What is missing is the thing WP20 exists to
build first — a baseline from a fixed reference intent — and a number with nothing to compare
it against invites the reading that whichever way it came out was expected. Scheduled rather
than refused, and the row says where.

**A gate.** G-Freshness already compares a phase's context hash against its lock, so the shape
exists. But what it would have to judge is the session, which leaves no trace in the
repository, and the only signal is the cost ledger, which is gitignored local data and outside
`artifacts_hash` by design. A gate reading it would make a verdict depend on one machine's
file, which is the opposite of what makes a verdict reproducible.

**Say nothing and leave it to the person.** Free, and the state #228 found. It was chosen
against for the reason the issue gives: the whole context economy of the process assumes
somebody clears, and the one place that was built to say a thing nobody has to act on said
nothing about the one thing nothing else can enforce.

<!-- xeno:section:impact -->
## Impact

`internal/runner/next.go`: two strings. The `not-started` case of `next` gains a clause and
keeps its command; the last phase's case gains a sentence and keeps having none. No new
function, no new field on `Suggestion`, no branch.

`internal/runner/runner_test.go`: two assertions, in the tests that already cover those two
suggestions. They hold the reason as well as the phrase, so a rewording that drops "what it
was given" fails.

`docs/assumptions.md`: one row, with both decisions #228 asks for beyond the sentence and the
reason for each.

**For a person.** Two more sentences on paths they already read: once per phase, and once at
the end of an intent. Nothing new to run, nothing refused, and `--no-next` silences both as it
silences the rest.

**For the trail.** Nothing. No artifact gains a field, no gate changes, no verdict anywhere
moves. Every sealed phase in this repository verifies exactly as before, which is what makes
this change safe to make at 118 intents.

**For a project that is not this one.** The sentence arrives with the runner and names no
harness, so it is the same advice for Claude Code, for Codex and for somebody typing commands.
A project that already clears between phases reads a confirmation; one that does not learns why
its digests are not doing the work they were built for.

**What is not measured.** Whether anybody follows it. That is the figure WP20 owns, and the row
in the register is what connects this sentence to it, so the next reader finds the hint and the
scheduled measurement together rather than one of them.
