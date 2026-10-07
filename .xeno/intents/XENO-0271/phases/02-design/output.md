---
intent: github.com/triplem/xeno#273
phase: 02-design
created: "2026-10-07T09:54:44Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7d22b99895c81d313f15e4b10bd37c244af34aaae621ff15c25dd0cf68d34886
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: Leave the shipped template directories as they are, and amend A33's state column with what spelling the phase prefix out would cost, measured both ways.
      rationale: 'The measurement was put on #273 on 2026-10-07 with three options. Renaming the directory alone is harmless to the trail but makes Load resolve by phase, so section 5''s ''resolution is per template id'' becomes false and a project override moves to .xeno/config/templates/00-intake/; it is a specification change with a code change behind it rather than a tidy-up. Renaming the id: with it is the consistent version a reader of the issue would assume, and it sets goneBundle for every sealed artifact so that strings_hash stops being recomputed across the whole trail, while gate verify goes on reporting all 495 verified. The prefix would buy ls ordering and nothing else, since template.yaml already carries phase: 00-intake. The third option, closing the issue with the comment and changing nothing, was declined because A33 would go on reading as a preference about prefixes with a reason that does not mention cost, and the next reader finds the assumption rather than the closed issue.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The measurement goes in A33's state column, not in its reason.** The reason is why the
decision was taken when it was taken, and it was taken without this measurement. Putting the
figure there would make the row claim a basis it did not have. The state column is where the
register already records what a row learned afterwards — A95 carries a correction to the whole
file's opening, A97 carries what settled a reading the clause's own wording argued against —
so it is the established place and not a convenience.

**`approved` stays.** The register says the state column reports what a person has said about a
row: a yes, a no, or a replacement. The yes stands; what is added is what the yes now rests on.
Changing it to anything else would say somebody had reopened the question, and the maintainer
closed it.

**The row names both variants, not just the one chosen against.** #273 asks for a prefix, and
the obvious reading is one change. The whole value of the measurement is that it is two, with
different costs, and a row that recorded only the destructive one would leave the next reader
thinking the harmless variant was available. It is available, and it costs section 5.

**The figure is written as a count taken on a date, not as a constant.** 495 artifacts record a
template ref. The trail only grows, so the number is wrong tomorrow in the direction that makes
the cost larger. Saying what it is a count of is what keeps it honest; A94 is the precedent for
a row carrying a count of references rather than a rule about them.

**`gate verify` noticing neither variant is stated as its own sentence.** It is the part a
reader would not predict and the part that makes the second variant dangerous rather than
merely wrong: somebody who renamed the ids and ran the one command that recomputes the whole
trail would see 495 verified and conclude it was fine. A row that only said the check retires
would be read as a thing a verdict would catch.

**The two incidental findings stay on the issue.** The all-digit scalar being skipped as a
non-string, and `goneBundle` being blind to a rename. Both are real, neither is A33's, and the
register's own test is that a row belongs there when the fact outlives the intent that found
it — these outlive it and belong to the gate rather than to the template id. Written where they
were found, which is the issue, and named in P1's non-goals so that leaving them out is visible
as a choice.

**No code comment, and the drafted reason for that was wrong.** The decision was first written
as "`TemplateID`'s comment already cites A33, so the row is reachable from the function". It
does not. `TemplateID`'s comment restates A33's reason in A33's words — "the prefix orders the
phases and says nothing a template needs" — without naming the row, and the only citation of
A33 in Go is in `phaseNumber`'s comment in `internal/runner/next.go`, which is about the
`--phase` prefix rather than about the template id. So the row is reachable by searching for
A33 and not by reading the function that implements it.

The comment still does not change here. Adding the citation is a one-word improvement and it
would put a Go file in a diff whose claim is that nothing in the runner moves, which is
criterion 7 and P1's constraint; and repeating the measurement in a comment would be a second
copy of a count that grows. What it is instead is a finding about reachability that this intent
leaves standing and names, rather than fixing on the way past.

<!-- xeno:section:alternatives -->
## Alternatives

**Do the rename, variant A: directory only.** The thing the issue asks for, in its harmless
form. `ls .xeno/plugin/templates` would sort in phase order and `model.TemplateID` would
collapse to the identity function, which is simpler than both A33 and the mapping table A33 was
chosen over. Measured as harmless to the trail: all 495 verdicts verify and a corrupted
`strings_hash` is still caught. What it costs is section 5. `Load` would be called with the
phase, so resolution would be per phase rather than per template id, the `id:` field would be
decorative, and `.xeno/config/templates/intake/` would stop being where a project override
goes. The sentence would have to be reworded before the code, which makes this a specification
change with a code change behind it rather than a tidy-up.

**Do the rename, variant B: directory and `id:`.** The consistent version, and the one a reader
of the issue would assume. It removes the mismatch instead of moving it. It also retires the
`strings_hash` recompute across every artifact in the trail, silently, while `gate verify` goes
on reporting all 495 verified. Refused on the measurement rather than on principle.

**Rename and renumber the artifacts so nothing goes stale.** The only version of B that keeps
the check alive: rewrite `template: intake@1.0.0` to `template: 00-intake@1.0.0` in all 495
artifacts and recompute every `artifacts_hash` behind them. It is the rewriting of a sealed
trail that section 11's "what is sealed is never rewritten" exists to forbid, and it would
change every verdict in the repository to buy a directory listing.

**Add a guard so B cannot be done by halves**, a check that each `template.yaml`'s `id:`
matches its directory name. It would catch the destructive variant, and it would also forbid
the harmless one, so it is a decision that the id and the directory are the same thing. That
may be the right decision; it is not this issue's, and nobody has asked for it.

**Teach `goneBundle` to tell a rename from a deletion.** The root of why B is quiet. It is a
change to what the gate can know — today the artifact names a ref and the tree answers one, and
nothing records that a ref was retired — so it needs a place to record retirement, which is a
new field or a new file. Written on the issue, out of scope here.

**Close #273 with the comment and change nothing.** The answer would be on the issue and A33
would go on reading as a preference about prefixes with a reason that does not mention cost.
The maintainer was given this as an option and chose the row instead, on the ground that the
next reader finds the assumption and not the closed issue.

<!-- xeno:section:impact -->
## Impact

`docs/assumptions.md`: one row, one cell. A33's state column gains the measurement of both
variants, the figure it is a count of, and the issue number.

Nothing else. No Go file, no template, no normative document, no gate, no artifact.

**For a reader of A33.** The row stops being a preference with a reason and becomes a decision
with a price. Somebody who has the idea #273 had finds, in the row that states the convention,
what each way of spelling it out costs and that the one command which would appear to check it
does not.

**For a reader who goes ahead anyway.** Variant A is available and the row says what has to be
reworded first. Variant B is available and the row says what it retires. Neither is forbidden
by anything; what changes is that neither can be done believing it is free.

**For the trail.** Nothing. `xeno gate verify` recomputes every sealed verdict and matches,
because the register is read by people and sits in no gate's input.

**For `.xeno/plugin/templates/`.** Nothing, which is the decision.

**What is left standing and named.** Three things, all on the issue and none fixed here:
`goneBundle` cannot tell a renamed bundle from a deleted one; a hash field whose value is 64
digits is skipped as a non-string by both the shape check and the recompute; and A33 is not
cited from `TemplateID`, so the row is reachable by search rather than by reading the function
that implements it.
