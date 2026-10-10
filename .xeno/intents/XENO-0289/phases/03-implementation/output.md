---
intent: github.com/triplem/xeno#359
phase: 03-implementation
created: "2026-10-10T15:48:49Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9584993f198902de4c50c852a8b54d0aa16e2ac7beb9edb8ddfd5b8508e010da
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - chosen: The label is named xeno-needs-decision, with the s, and nothing is renamed or created
      decided_by: triplem
      id: D-2
      rationale: 'The comment that approved #359 asked for the label and wrote it xeno-need-decision; the label created on 2026-10-10 and applied to seven issues is the plural. The maintainer settled on the plural in session on 2026-10-10, in answer to the agent raising the mismatch, and the answer was relayed to this intent and written here by the agent on instruction: the decision is the maintainers, the transcription is not. It costs nothing to carry out, because the tracker already holds the plural and the tree names no label at all outside internal/model/identity.go, which names only xeno-approved. The alternative, renaming to the singular, would have re-associated seven issues on the host and broken any saved filter naming the plural silently, in exchange for matching a spelling written once in a comment. What the near-miss cost is recorded on the page rather than here, because it is the page the next reader will have: no code read either spelling, so nothing failed, and the cost a second spelling can have is a person filtering on the wrong one and reading an empty list as there being no blocked issues.'
---

# Implementation

<!-- xeno:section:changes -->
## Changes

Three files, no code, and the labels on the host untouched but for the one this intent's
own question needed.

**`docs/labels.md`**, new, 193 lines and 12,146 bytes. It opens by saying what a label is
here — that no label holds an intent's state, that the state is in `.xeno/intents/` and
is what `xeno intent status` prints, that nothing in Xeno ever writes a label because the
adapter contract is three operations, and that exactly one label is read anywhere. Then
the set as read from the host on 2026-10-10, with the command that read it, in three
groups ordered by how binding they are.

*The one label Xeno reads.* `xeno-approved`, in a five-column table — who sets it, who
clears it, what reads it, what it blocks — followed by the prose the table cannot hold:
that approval is the label and the comment and neither alone, with section 12 cited
rather than paraphrased; that `Issue.Approval()` compares case-insensitively and takes
the last qualifying comment; that the label is read twice, at `intent start` and at P0,
and never again, with A103 for why the second read exists; and that nineteen issues
carried it against a trail of more than a hundred and forty intents, which A107 explains
and which is the page's clearest demonstration that the durable record of an approval is
the sealed intake sentence and not the label.

*The labels this repository has given itself.* `xeno-needs-decision` and `wp0` to `wp20`,
with the sentence that neither is read by any code and an adopter needs neither. The
first says who created it and when and on whose instruction, what the description says,
how it sits beside section 8's `open-questions` rather than replacing it, which seven
issues carried it, and that its name is the plural by the maintainer's decision of
2026-10-10, D-2 of this phase. The paragraph on the name is the page's one worked example
of what it is for: no code read either spelling, so nothing failed and no verdict moved,
and the cost a second spelling can have is a person filtering on the wrong one and
reading an empty list as there being no blocked issues. The second group quotes the plan
on what a work package label is for and reports the search that found nothing reading
one.

*The host's own labels.* All ten named, with the three in occasional use and their
counts, and the reason they are listed at all: a page that enumerates a tracker's labels
and silently omits ten of them cannot be checked against the host.

*Applying a label, and reading it back.* Both forms, the read-back as the practice, and
the probe that did not reproduce reported as a probe that did not reproduce.

*What this page does not say.* Three open points, each with its issue: the lifecycle of
`xeno-approved`, whether a fresh approval is owed, and triggering, which is #338's and is
the half of #359's title this intent is not doing. The section closes by saying that
nothing mechanically checks any of it.

**`docs/README.md`**, five lines under "Running it", between `commands.md` and
`symbol-index.md`. The entry says what the page is and, in the same breath, what it is
not: a record of what is in use rather than a definition of what a tool reads.

**`zensical.toml`**, one line in `nav`, under "Working with it" after "Commands".

**On the host.** `xeno-needs-decision` applied to #359, and a comment carrying one
question with three options, each option's consequence and its cost, a recommendation
with its reason, and a free entry. No label created, renamed or removed; no work package
label applied, which stays the maintainer's act.

**What was corrected while writing.** Two figures taken from the session rather than from
the host were wrong, and the page carries the host's. The label set is 33 and not 34, as
`01-requirements` criterion 1 says; the count was written into a sealed artifact before
being read, and the check that criterion actually asks for is the set comparison in both
directions, which the page passes. And `xeno-needs-decision` is on seven issues — #120,
#224, #234, #329, #334, #338 and #359 — not on the four that `00-intake`'s decision
record names, and #352, which that record names, does not carry it. Both are reported in
`04-verification` as well, because a sealed artifact is not rewritten and a correction a
reader of that phase cannot find is not a correction.

<!-- xeno:section:deviations -->
## Deviations from the design

**This phase was redone.** It was judged green once, with the page carrying the spelling
of `xeno-needs-decision` as an open point and four open points in its closing section.
The maintainer then settled the name — the plural, nothing renamed, nothing created — and
the answer reached this intent while `04-verification` was running. Rather than leave a
sealed artifact describing a page that no longer matches it, the page was written again,
the decision was recorded as D-2 of this phase, and this phase was finished again, which
is section 6's redo: the writes landed under `.xeno/local/staged/` and the finish applied
them and judged in one step, so the artifact and the verdict moved together. The
`04-verification` directory was removed before the redo and the phase started again
afterwards. It held no verdict and was in no commit, so nothing sealed was rewritten; had
it held one, its lock would have recorded a predecessor that the redo moved and the
honest repair would have been the same.

**The design said "on the order of 18,000 bytes" and the page is 12,146 in 193 lines.**
The three tables carry what the design put in them and the prose came out shorter than
estimated, which is a difference in the estimate and not in the design. The context
budget is 14 files and 130,000 bytes against eleven files and some 76,000, so nothing is
affected either way; this is why the budget was set generously rather than tightly.

**The built navigation was read back locally, which the design did not plan for.**
Criterion 7 asks for the nav to be read from the built site rather than assumed from the
toml, and `zensical` is not part of this repository's toolchain — the `docs` job installs
it. It was installed into a throwaway virtual environment, the site was built with
`--clean --strict`, and the page is in `index.html`'s navigation as `./labels/` with its
own directory under the site root. The same build gave the positive control the next
phase reports. A step taken rather than a deviation, recorded because the design named CI
as where this would be checked and it was checked here too.

**No deviation on where the page went, on how it is grouped, or on the plugin script.**
`.xeno/plugin/bin/xeno-labels.sh` is untouched and its claim to create "the one label
Xeno asks a project's tracker for" is still true, which the page states in its own words
at the end of the `xeno-approved` section.
