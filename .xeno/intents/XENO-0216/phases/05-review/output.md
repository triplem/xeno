---
intent: github.com/triplem/xeno#156
phase: 05-review
created: "2026-10-01T14:42:05Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 55778d600ac477ef7ffe43dd7bf00c64976b5653e1238508f970fbe0504fbabc
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

The rule set is not built, so there is no rule to render an entry from and no lens enabled.
Every entry below carries `source: lens` in spirit and no rule id in fact: they are the
questions this change actually raises, answered here, and G-Policy stands as
`not-implemented` rather than counting them.

**Does any corrected sentence state something the tree does not say?** Read back one by one
against the thing each describes. The register's two paragraphs against the gate table, the
command surface and A62 to A65; the README's opening against the package list; the command
block against the dispatch table; the trail paragraphs against the phase directories; each
coverage row against the test it names. Answered yes, with the one difference recorded in
04-verification.

**Do the three replaced paragraphs read as paragraphs?** Read whole rather than as a diff,
which is the convention's own instruction. The register's "Not built" now runs blocked work
first, then unstarted, then WP12's half, which is an order a reader can follow. The README's
trail reads as two shapes with a boundary. Answered yes; this is a judgment and is recorded
as one.

**Is the ownership finding stated once?** #156 carries it in its own section. Neither
corrected file mentions it, because a file remarking on who maintains it is the absorption
the third standing rule forbids. Answered yes.

**Was anything written that the documents forbid?** No file under `docs/` is touched, no
field, gate, tool or rule is invented, and the figures go to #117 as a comment rather than
into an artifact. Answered yes.

**Does the change belong to a package and an intent?** It belongs to no package, which is the
finding; the issue carries `wp16` as the nearest with the label disclaimed in its body, the
branch carries this intent alone, and the commit will reference and close #156. Answered as
far as the plan allows.

**Is anything here a decision somebody else should have taken?** Two. Naming an owner for the
two files, left to a person as a plan change. Building the check that would catch recurrence,
left to WP16 where the generated reference lives. Both are in the residual risk.

<!-- xeno:section:release-notes -->
## Release notes

Documentation only. No command, flag, gate, artifact field or verdict changes, and no
released binary differs.

**`README.md`** now describes the runner as it is: the digest and its secret filter, the cost
record, the branch rules port and the symbol index reader beside the four capabilities it
already named. The network claim is narrowed to the one that holds, the gate path, with
`enforcement check` named as the one command that reaches a host. The command block carries
the three `assumption` commands, `cost turn`, `--export`, `--summary` and
`intent status --all`. The section on this repository's own trail no longer says the intents
stop after intake: from `XENO-0107` they run all six phases and finish green. The coverage
table gains WP8, WP11, WP12, WP13, WP15 and WP17.

**`ASSUMPTIONS.md`** no longer lists the secret filter or the digest writer as unbuilt, names
WP12's open part as its tracker half, says what G-Supply and G-Secret wait on, and leaves the
marketplace URL as the one open verification point. A paragraph is added for what was built
after its "Built" list was written.

**Nothing is deprecated and nothing moves.** A reader who had learned either file will find
every heading where it was, except that the README's title is now `# xeno` rather than
`# xeno, core`.

Closes #156.

<!-- xeno:section:residual-risk -->
## Residual risk

**The cause is untouched and the class is open.** Both files are accurate today and still
owned by nobody, so nothing stops the next twenty intents from repeating this exactly. The
finding in #156 is a record, not a guard, and the only reason this intent does not fix it is
that the fix is a plan change and a person's commit.

**The check is cheap and was not written.** A diff of the README's command block against the
dispatch table is three lines of shell and found the one expected difference in this intent's
verification. It was left out because WP16 specifies the generated reference and a second
generator gets deleted by the package that should have owned it. That is a defensible reason
and it is still a decision to carry a known-cheap check as a risk: if WP16 slips, this slips
with it.

**Two files were checked and the rest were not.** `CONTRIBUTING.md`, `SECURITY.md`,
`SUPPLY-CHAIN.md`, `M0.md` and the `.xeno/config/project.yaml` comments make claims about the
tree too, and none of them was read for staleness here. The survey that found this looked at
the two most-read files, so the probability that the others are clean is the probability that
drift stopped at the front door.

**The register's own history was deliberately not re-examined.** Entries closed months ago
were taken at their state column rather than re-verified against the tree. A closed entry
that was answered wrongly would survive this intent untouched, and the only evidence against
that is that A6's closure is what made the README's paragraph false, not the register's.

**The proportionality figures are still partial.** Cost per intent has no writer in a form
these artifacts carry (#65), so #117 receives sections, words, lines and timings again and
not the token figure the plan's question is about. Every intent that posts a partial figure
makes the eventual answer later rather than sooner.

**The `gate run` figure is one phase on a copy.** Three milliseconds says the gates are not
the lever, which is consistent with the earlier measurement of nine milliseconds over a
smaller trail, and it is a single phase of a documentation intent. Nothing here supports a
claim about a phase with evidence, findings or a large information base.
