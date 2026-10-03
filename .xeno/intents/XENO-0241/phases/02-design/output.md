---
intent: github.com/triplem/xeno#206
phase: 02-design
created: "2026-10-03T20:45:07Z"
schema_version: "1.0"
runner_version: dev+8574810
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ac8f826d59c1efb7e9c4b36c5364f0e69e4e4bb1be9cc735e3b660db2c165276
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

**A command, not a gate.** G-Complete is correct and runs where section 7 says. What was
missing is a reader at the merge, and the gate list is a budget. `xeno intent verify` sits
beside `xeno enforcement check` and `xeno check commit-message`: commands that answer one
question about a repository without producing a verdict about a phase.

**It reads `state`, it does not reimplement it.** `summarise` already returns `complete`,
`abandoned` or the phase reached, and the command's whole condition is which of the three
came back. XENO-0240 deleted a second derivation of an answer the tree already had, for
the reason its design gives, and this is the same shape with the second derivation never
written.

**The range decides which intents are in question, not the trail.** A branch is answerable
for the intents it touches and for nothing else. Checking every intent in the repository
would make a pull request red for an intent somebody abandoned without closing a year ago,
which is a finding about that intent and not about this change.

**`--no-renames`, so both paths of a move count as touched.** The trail-rewrite step above
already refuses a rename under `.xeno/intents/`, so the two can never disagree; asking git
not to detect renames means this step cannot be the one that lets a move through.

**Both ends required, no default.** Section 9 says the range is passed in and never
inferred. An absent or unresolvable ref exits 2, "could not run", because a vacuous
comparison reports the same thing as a clean one — which is the sentence the trail-rewrite
step already makes about its own base.

**A provisional P5 is not complete.** `Decided` excludes provisional, so the state reads
`05-review` and the check refuses. Section 6 says a provisional verdict is reported rather
than failed and becomes "a hard condition at the merge request"; this command runs at the
merge request and is that condition. `xeno gate verify` keeps exiting 0 on provisional,
which is the same sentence read at the other end.

**The shipped wrappers carry the call.** Both already pass both ends of the range for
`gate run`, so the check reaches an adopter with no new template field. A check that lived
only in this repository's workflow would be a property of this repository.

**`CLAUSE-READERS.md` is left alone.** It dates itself, calls itself a measurement rather
than a specification, and says ageing is the most a document can do about itself. Editing
it to report that one of its findings is closed would make it an index somebody maintains.

<!-- xeno:section:alternatives -->
## Alternatives

**Accept it and record why**, which the issue offers as a legitimate third answer. Put to
a person and rejected: an intent abandoned in flight is a real thing, but `xeno intent
close` already exists for it and costs one command, so what is being accepted would not be
the cost of closing but the cost of nobody noticing. It also could not have been done by
the agent, since the plan is a document and a document change is a person's commit.

**A step of shell in `.github/workflows/xeno.yml`**, greping the touched directories and
reading `xeno intent status` output. Rejected: the condition would live in YAML, parsed
back out of text meant for a person, with no test over it and no way for an adopter to get
it. The job already has one step whose logic is shell, the trail guard, and that one is a
`git diff` and a string test with nothing to derive.

**Extend `xeno gate verify` with `--base`.** Rejected. `gate verify` recomputes verdicts
and its exit code carries a decided meaning — provisional is 0, by #110 and section 6 —
and this check has to refuse exactly the provisional case. Two conditions with opposite
answers about the same state do not belong behind one exit code.

**Have `xeno intent status` exit non-zero.** Rejected: it is a listing, and a listing that
exits 1 because a row in it is unfinished cannot be used to look at a trail that has
unfinished rows in it. This is the issue's second answer, and it collapses into the first
as soon as you ask what reads the report.

**Check every intent in the repository rather than the touched ones.** Rejected above, and
worth naming as the tempting one: it would have caught XENO-0230 on the next pull request
after the one that merged it, which sounds better and is worse. Every later branch would
then be red for somebody else's intent, and the project would learn to pass the check by
closing intents it had not finished.

**Refuse at `phase start` instead**, the way the sequence is enforced. Rejected: there is
no command to refuse. An intent that stops short is characterised by nothing happening,
and only something that runs anyway — the merge check — can notice an absence.

<!-- xeno:section:impact -->
## Impact

`internal/git` gains one function that lists the paths differing between two trees under a
directory. Its package comment currently says it starts exactly one subprocess, `git log`;
that becomes two, and the comment is rewritten rather than left to be contradicted by the
code below it.

`internal/runner` gains one reader that turns those paths into intent keys and reports the
ones whose state is neither `complete` nor `abandoned`. It uses `summarise`, so the state
has one definition in the tree.

`cmd/xeno` gains `intent verify` in the dispatch table, in the usage block, and one
printer. The command takes no `--intent` and so is the third in its family to check its
own arguments, which the table already expresses as a field rather than a special case.

`.github/workflows/xeno.yml` gains one step, next to the trail guard and with the same
base. One behaviour changes for everyone working in this repository: a pull request is red
until its intent reaches a decided P5. Nothing in the trail is affected, since all seventy
post-M0 intents are already complete.

`internal/scaffold/files/ci-github.yml` and `ci-gitlab.yml` each gain the same call. An
adopter's generated wrapper therefore refuses a merge whose intent stopped short, which is
the clause this issue found unread.

Section 7's G-Complete gains a second reader, as section 11's sealing clause did in #193
and #216: one inside the process and one at the boundary. Nothing in either document
changes, and no gate, field or rule is added.

What does not change: `gate verify` still exits 0 on a provisional verdict, `xeno intent
close` is still the way an intent is abandoned, and nothing about this writes.
