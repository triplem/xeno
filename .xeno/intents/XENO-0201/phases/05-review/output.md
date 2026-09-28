---
intent: github.com/triplem/xeno#120
phase: 05-review
created: "2026-09-28T20:36:05Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6e72fed.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: fd462d32d66c827da28473e5197ac8fd7e44676eeb70443e05bb587300b89570
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**Nothing under `docs/` changed, and that is the point of the intent.** Section 5 says
the agent supplies the summary text and the runner filters, hashes and writes the file,
and the working sequence of section 6 lists the digest under `phase finish`. Both
sentences were normative before the runner existed. This is the code catching up.

**Nothing was invented.** Every field written is one section 5 enumerates, the two new
values come from the block section 12 defines, and the command surface gained a flag
rather than an operation.

**The runner writes nothing it cannot source.** `secrets_hash` is absent because there
is no filter, `tool_version` because the config that carries `tool` does not carry it,
and a project with no `agent` block gets neither field rather than a plausible one.
A35's argument is kept by the change that amends it.

**The summary is optional.** Section 5 says the supported path is not an enforced one,
so a phase finished without one is judged exactly as before, and every test written
before this change passes with an empty summary.

**The digest is written before the gates run.** It is inside `artifacts_hash`, so the
order is not a detail: written afterwards it would seal a hash over a tree that lacked
the file.

**Nothing existing was rewritten.** Sixty intents keep their hand written digests,
inside their hashes, and `gate verify` over 85 verdicts is the check.

**A35 is amended rather than superseded**, because its claim still holds for the three
fields that remain and only its arithmetic was wrong. The row has now been read twice as
a list of members, and it says so.

**The phases were written in their order**, and this intent's own digests from P3 onward
were written by the runner, which is the first time that has been true of any digest
here.

**What the checklist cannot cover is named.** AC5 is argued from this intent's own phase
going red and then green, not from a test, because a test would have to fabricate the
two fields the writer exists not to fabricate.

<!-- xeno:section:release-notes -->
## Release notes

`xeno phase finish --summary PATH` writes `digest.md`: the frontmatter section 5
requires of a file produced in a session, then the summary text. `--summary -` reads
stdin. Without the flag the command behaves exactly as before and writes no digest.

The digest carries `intent`, `phase`, `created`, `schema_version`, `runner_version`,
`plugin_version`, `language`, `context_hash`, `model` and `tool`, in the order section 5
lists. It carries no `template` or `strings_hash`, because a digest is not rendered, and
no `rules_hash`.

`model` and `tool` are read from the `agent` block of `.xeno/config/project.yaml` and
are written into `output.md` as well. A project that records neither gets neither field,
never a default.

It carries no `secrets_hash`. Section 5 gives the runner filtering as well as writing,
and the effective filter is a shipped pattern file that does not exist yet, so a digest
written now passes through no filter and does not claim to. G-Schema reports the field
missing, which is true, and a phase whose digest the runner wrote is red until
`secrets_hash` and `tool_version` are filled.

Of the five fields A35 left to a hand, three remain: `tool_version`, `secrets_hash` and
`rules_hash`.

Nothing existing changed. Digests written before this keep their content and their
verdicts.

<!-- xeno:section:residual-risk -->
## Residual risk

**The writer produces a file the gates reject.** Two required fields have no source, so
`phase finish --summary` yields a red phase and a hand still finishes the frontmatter.
That is honest and it is not finished: whoever reads the release note will expect a
digest they no longer have to touch, and they will have to touch two lines of it.

**Nothing filters, and the file now looks produced.** A summary carrying a secret is
copied verbatim into the tree, exactly as a hand written digest always was, but under a
writer whose name suggests it was processed. Section 5's other half is the mitigation
and it does not exist. Of everything here this is the one to close first.

**#120 is not answered.** The scoping named this the gap that had to close before
enforcement could be argued, and it is half closed. An agent still cannot produce a
complete digest with the tool alone, so a rule forbidding the hand would still be a rule
nobody could keep.

**The before and after cannot be run.** `Finish` changed shape, so four intents' worth
of method stopped working here, and the substitute is evidence from the tree rather than
a comparison. Any future change to a signature will have the same limit.

**`tool_version` has no home.** Section 12's block carries the tool and the model and
not the version, which is defensible and leaves the field to WP11 with nothing recorded
about where it should come from.

**Accepted with the five named.** None blocks the change, and each is smaller than the
state it replaces, which was a runner that judged a file it was specified to write and
sixty digests typed by hand because nothing else could produce them.
