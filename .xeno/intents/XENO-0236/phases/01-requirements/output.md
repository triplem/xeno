---
intent: github.com/triplem/xeno#183
phase: 01-requirements
created: "2026-10-03T18:38:05Z"
schema_version: "1.0"
runner_version: dev+0462eb7.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1678610fe109d5cfa4c7f576262ef253ed19a644e259f21b514d2bc07cdca4d8
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

**Section 7 says the plugin is the vendored one and describes no resolution order.** The paragraph
that described one is replaced, not annotated, and says why each position in it was unreachable or
unsafe, with both measurements.

**`XENO_PLUGIN_ROOT` is not in the list of variables the normalised environment carries.** Three
remain: `XENO_PLUGIN_DATA`, `XENO_HARNESS`, `XENO_HARNESS_VERSION`.

**The plan's WP7 sentence names the same three** and says the root and the order were on that list
and are not.

**One sentence states the condition an override would have to meet**, so a later intent starts from
it rather than rediscovering it: a runner that cannot verify a plugin does not accept one from
outside the repository.

**The specification change is its own commit with no code in it**, which the first standing rule
requires, and nothing follows from it but comments and rows — because the order was never
implemented.

**No code changes how the plugin is found**, because nothing did. `internal/rules`,
`internal/template`, `internal/secrets` and `internal/plugin` read the vendored directory before and
after.

**The comments that called the order a decision pending now call it removed.**
`internal/plugin`'s package comment and the entry point's.

**A84's open clause is closed** — it said the resolution order was what remained and the decision was
section 7's.

**The test that asserts the entry point exports no plugin root keeps asserting it**, with its reason
changed: the thing worth preventing is a root arriving from the environment, and that does not stop
being worth preventing when the document stops naming a way to do it.

**Nothing in the repository still claims the mechanism exists.** Every remaining mention explains its
removal.

**The suite, `gofmt`, `go vet` and `gate verify` are unchanged**, which is what a document change
with no mechanism behind it should look like.

<!-- xeno:section:non-goals -->
## Non goals

**No override, under a condition or otherwise.** It was the alternative and it was declined: a clause
with no reader is cheaper to delete than a mechanism with a guard is to maintain. Section 7 keeps one
sentence so that the day somebody needs one, the condition is already written down.

**No change to how the plugin is found.** The code read the vendored directory and still does.

**No change to `XENO_PLUGIN_DATA`.** Still listed, still set by the entry point, still read by
nothing, because section 7 pins it to one value. A84 records it and this does not revisit it.

**No change to G-Supply.** It compares the vendored tree, which is now the only tree. That it is
inert under a development build is A86's and is unchanged by removing a clause.

**No amendment to A42.** The order was the one thing that would have required weakening it, which is
the argument for removing the order rather than the argument for amending A42.

**No revisiting of the three clauses of #183 already settled** — the entry point, `XENO_HARNESS`
recorded into `tool`, and the check that nothing branches on it.

**No deletion of the client fallback's reasoning.** It is unreachable and the section now says so,
because a reader who finds a mechanism missing is better served by the reason than by the silence.

<!-- xeno:section:constraints -->
## Constraints

**The documents are not editable by the agent, and this change is to a document.** It exists because
the maintainer chose it from three shapes with the measurements in hand, and it is its own commit
made before anything that follows — the first standing rule, applied to a removal rather than an
addition.

**A removal is a heavier claim than an addition.** So the paragraph that replaces it carries the
reason for each position rather than only the conclusion, and both measurements, so that somebody
reinstating it has to argue with evidence rather than with an absence.

**The order was never implemented**, which is what makes the removal safe and is worth stating: there
is no code to delete, no behaviour to migrate, and no artifact that recorded a root.

**Section 7's safety sentence rested on G-Supply**, which now exists and covers a released binary
only. That asymmetry is the reason the clause could not simply be honoured, and it belongs in the
replacement rather than in a commit message.

**One sentence is kept rather than none.** A condition discovered by measurement is worth more
written down than rediscovered, and the day an override is wanted is the day somebody will look here
first.

**Nothing may still claim the mechanism exists.** A document that removes a clause while a comment
elsewhere describes it as pending is the stale-list failure this repository has already met twice.

**88 columns, SPDX, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0, exit code captured
and not piped.**

**One intent, one branch, one issue** — `183-the-plugin-is-the-vendored-one`, #183, labelled wp11, on
main. The issue closes on this: it is its last clause.
