---
intent: github.com/triplem/xeno#237
phase: 05-review
created: "2026-10-06T12:11:39Z"
schema_version: "1.0"
runner_version: dev+7885661.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a3a5bee7aeafa0fe86b3ea0d922aeb38dc251976b503c84454f44ce836aa9525
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: deviation
      note: 'Four, in P3. (1) the composed suggestion in next.go now names the release twice, accepted so that a reader of gate.yaml, who has no sentence beside the field, gets both of section 7''s routes. (2) ''Then judge it again'' sits between the two routes in that sentence, which reads right for the re-run and wrong for the approval; it is next.go''s sentence and a non-goal. (3) the end-to-end check had to be rebuilt: the first attempt moved the ground under P0 and the gate passed, because no P0 lock in this trail records a files list, so a green gate was about to be written down as proof. (4) the remedy names who approves rather than which command, which costs the gate.yaml reader one hop and respects next.go''s stated reason for not handing the command over. No criterion was falsified.'
    - rule: interface-change-needs-a-migration-note
      result: not-applicable
      note: No interface changes. Three remedy strings and three tests, in internal/gates. model.Finding is unchanged, so gate.yaml's shape is unchanged and nothing in the trail re-hashes; 0 of 443 sealed gate.yaml files carry a staleness finding, so nothing already written is made historical, and a finding that was would correctly keep its text because next is sealed with its verdict. No command, flag, field, gate, template or rule changes. The new wording reaches an adopter with the next release, because internal/gates ships in the binary rather than in the plugin, and needs nothing of them.
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: None added, and go.mod is untouched. The fault was a string naming an act the runner refuses, with no test over it; what catches that is a test that can fail, which three mutations confirmed, and following the remedy rather than reading it. Neither costs a dependency.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** No normative document is touched, and this is the unusual case where
that is the whole point: section 7 already says "a stale phase is not deleted, it is re-run or
explicitly approved as still valid", so the repair makes a string agree with the specification rather
than asking the specification to change. Nothing is invented — no field, no finding, no cause, no
gate; `model.Finding` is unchanged and so is `gate.yaml`'s shape. The work belongs to #237 under
`wp8`, on a branch carrying one intent, and the commit references it.

**The decision was put one at a time, with its options, their consequences and a recommendation.**
Three options: name section 7's routes, teach `phase start` to re-resolve a lock, or state the fact
and drop the instruction. Each carried its cost — the second contradicts #215 and needs a section 11
commit, the third replaces one exception to #216 and #218 with another — and the recommendation was
named with its reason.

**The remedy was followed, not read.** The issue's "Done when" is that following the remedy clears
the finding, and that is a claim about behaviour, so it was run: a scratch intent, a file moved under
P1, P2 gated, the finding produced, `gate approve` on P2, the verdict `approved`, P3 started. Reading
the new string and judging it plausible would have satisfied every other criterion and not this one.

**The first proof was green and wrong, and that is the session's second instance of the same shape.**
Moving the ground under P0 and gating P1 produced a passing gate, because P0's lock records no files
at all. A green gate read as "the remedy works" would have been #263's convention broken from the
positive side — the thing checked was absent rather than passing. Earlier in this session XENO-0259
hit the same shape from the other direction, where a tree that had not been built would have been
grepped for a package that was not in it. The convention is a month old and has now earned its place
twice in one day.

**The tests were confirmed able to fail.** Three mutations, three failures, and the file restored.
A90's objection is about readers that cannot fail and it applies to a test over a string more than
anywhere: an assertion that the remedy "contains the phase name" passes on the cause as readily as on
the route.

**What was found and deliberately not fixed.** 0 of 117 P0 locks and 17 of 66 P1 locks record a
`files` list, so the check this repairs has had almost nothing to read for most of the trail. That is
why a remedy naming a refused act survived four intents. It is a change to when `phase start` writes
the lock, which is #215's and section 5's, and the third standing rule says a finding that belongs to
no package gets written down rather than absorbed. It is in P3's changes, P4's gaps and the issue.

**One cost is accepted and visible rather than smoothed over.** `Runner.red` now names the release
twice, because the remedy carries it for the reader of `gate.yaml` who has no suggestion beside the
field. P3's deviations records it with the output that shows it, and the alternative — leaving the
release out of `next` — would have given that reader one route where section 7 gives two.

**The unreadable case was defended rather than made symmetrical.** It is the one of the three whose
remedy already worked, and its test asserts the *absence* of the other two routes, so that a later
pass tidying three strings into one fails instead of succeeding.

<!-- xeno:section:release-notes -->
## Release notes

**A staleness finding now names a remedy the runner will carry out.** It named one the runner
refuses. The three findings `staleReads` raises all said "read the phase again", and `phase start`
declines that for a judged phase: the lock is inside `artifacts_hash` and is the only record of what
the phase was given, which are #215's two reasons. Worse, the first route the refusal offers —
`section set` and `phase finish` — re-renders the artifact and leaves the finding exactly where it
was, because only `phase start` writes a lock. A person following the printed remedy was turned away
twice before reaching the one act that works, and that act was not in the remedy.

**The two findings about a sealed lock now carry section 7's two routes**, in the clause's own order:

    start 01-requirements over, which discards it and every phase after it, or a second
    person approves this finding on 02-design as still valid; section set and phase
    finish will not clear it, because the lock keeps what the phase was given and only
    phase start writes one

The re-run is named first because it is the act that makes the staleness untrue; leading with the
approval would tell a reader the staleness is acceptable before they had looked. Both phases are
named, because the finding is about the lock of an earlier phase and is raised while gating a later
one, and the two routes apply to different ones.

**The approval is named as a second person's act and not as a command.** `Runner.red` in `next.go`
already prints `gate approve` and its comment says why it hands over neither way out: releasing a
finding is a statement by a second person. A reader of `gate.yaml` has no suggestion beside the
field, so the route has to be in the string, and naming who does it rather than what to type is what
both readers can have.

**The unreadable-file remedy stays.** "make it readable and run the gate again; the tree is wrong
here and not the lock, so nothing has to be approved or started over." That fault is this machine's
— a permission, a broken link, a filesystem — and the next `gate run` clears the finding with nobody
approving anything. Its test asserts that it does *not* ask for the other two routes, so a later pass
making the three alike fails.

**It was proved by following it.** On a scratch copy: the finding appeared, `gate approve` on the
gated phase marked it `approved`, the verdict became `approved`, and the next phase started. That is
the issue's "Done when" and it is measured rather than argued.

**Found while proving it, filed rather than fixed.** No `context.lock.yaml` of any P0 in this trail
records a `files` list — 0 of 117 — because `phase start` writes the lock and `scope set` writes P0's
scope afterwards. Of the P1 locks, 17 of 66 carry files. So this check has had almost nothing to read
for most of the trail, which is why a broken remedy survived four intents. Changing it is #215's and
section 5's territory.

**Not here:** a `phase start` that re-resolves a lock, which contradicts #215 and needs a section 11
commit first; anything about when the check fires, which #236 settled; and #235, the other wp8 issue,
which needs a specification commit of its own.

<!-- xeno:section:residual-risk -->
## Residual risk

**The check has almost nothing to read.** 0 of 117 P0 locks and 17 of 66 P1 locks record a `files`
list. A repaired remedy on a check that cannot fire is a better sentence nobody meets, and the
repair is worth exactly as much as the check becomes worth. Filed, in P3's changes and P4's gaps, and
not fixed: it is a change to when `phase start` writes the lock, against #215's reasons for writing
it once.

**The re-run route is quoted from the refusal and not followed.** Only the approval was run end to
end. Starting the earlier phase over means removing its directory and every phase after it, which
would have meant rebuilding the scratch copy; the refusal naming the act was reproduced instead and
the wording written against it. So one of the two routes the remedy offers has been performed and the
other has been read.

**Nothing asserts that the remedy and the refusal stay in agreement.** The remedy says only `phase
start` writes a lock; `Start` refuses a second start for #215's reasons. Two statements in two
packages that have to remain true together, with no test over the pair and a person as the only
connection. A new instance of the shape `docs/clause-readers.md` catalogues.

**The composed suggestion names the release twice and says "judge it again" between the routes.**
Accepted and recorded. The duplication is for the reader who has the suggestion; the alternative was
an omission for the reader who does not. The ordering reads right for the re-run and wrong for the
approval, and it is `next.go`'s sentence.

**The unreadable case's test assumes it is not run as root.** It chmods to 0o000 and expects the read
to fail; as root it would succeed, the check would find nothing, and the test would fail for a reason
unrelated to the remedy. Unguarded, and the kind of assumption that holds until the one environment
where it does not.

**No sealed artifact carries the old remedy, which was checked rather than assumed.** 0 of 443
`gate.yaml` files in the trail hold a staleness finding; the only occurrences of the old string are
XENO-0245's intake, which quotes it, and this intent's own. So nothing in the trail is made
historical, and a finding already written would correctly keep its text, because `next` is sealed
with the verdict it belongs to.
