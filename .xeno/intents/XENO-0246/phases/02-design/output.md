---
intent: github.com/triplem/xeno#239
phase: 02-design
created: "2026-10-05T08:52:37Z"
schema_version: "1.0"
runner_version: dev+3f5ad34
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e83430118de967bf1f163bfdc8475e81070d59607630324db5c6a6016788fa21
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

Three of the four keep their names, lowercased: `docs/assumptions.md`,
`docs/supply-chain.md`, `docs/clause-readers.md`. Every file already under `docs/` is
lowercase, and an uppercase name there would be the inconsistency the move is meant to
remove.

`M0.md` becomes `docs/m0-gate-job.md` rather than `docs/m0.md`. Its siblings are
`implementation-plan`, `orchestrator-evaluation`, `process-definition` and `v2-delta` —
descriptive hyphenated phrases, not identifiers — and the document's own title is "The M0
gate job: Xeno verifies its own repository". #222 offered `docs/m0.md` or `docs/M0.md` and
left the choice open; this is the third, taken knowingly, and it is the one rename whose
new name a reader cannot guess from the old, which is why it is a decision and not a
detail.

The stale references are recorded in one assumption row, not four. They are one fact — the
trail is sealed and the plan is not the agent's — and splitting it across a row per
document would make four records that go stale together and are read apart.

The row goes in `ASSUMPTIONS.md` rather than in a phase's `learning.yaml`. A learning is
narrow by A26: it has to be expressible as a rule or template change, and "these references
are stale and will stay stale" proposes neither. It is a decision about this repository's
construction, which is the register's category.

The moved M0 document's new line goes at the top, under the title, before the first
paragraph. A reader who is going to misread a walkthrough as instructions does it in the
first screen, so a note further down is a note that arrives after the mistake.

<!-- xeno:section:alternatives -->
## Alternatives

`.xeno/` for the register was considered and rejected, and #239 argues it at length.
Section 4 enumerates what an intent directory may hold and G-Complete reports anything
else, so a hand-maintained register among gate-judged artifacts would be a file that ages
with no verdict noticing. That is the shape of defect this project keeps finding, not one
to introduce.

Rewriting the sealed references was considered and is foreclosed rather than rejected: 112
mentions of `ASSUMPTIONS.md` sit inside `artifacts_hash`, and correcting them would stale
every verdict over them. #217 met the same arithmetic over its own artifact and resolved it
the same way.

Four commits, one per document, was the obvious alternative and is what #239 names as the
reason to prefer one: the reference sweep is a single pass over the same files, and four
passes would do it four times and leave three intermediate trees in which some references
resolve and some do not.

`docs/m0.md` was the issue's own first suggestion and is rejected on the convention it was
meant to satisfy. Lowercase was the point, but `m0` alone is an identifier where its four
siblings are phrases, so it would satisfy the case rule and miss the naming one.

Deferring the whole move to WP16 was considered. It is the package that owns link checking
and would otherwise fix this later, which is the argument for doing it now instead: the
checker is easier to write and easier to trust over a tree that already has the rule.

Adding the stale-reference row to the register was weighed against leaving it to this
phase. The register's own banner says the record is closed and nothing is added, which
would decide it — except that A78 through A93 were all added after that banner, by sixteen
intents including the four most recent documentation ones. The banner is stale rather than
binding, so the row follows practice and the banner's contradiction is recorded as a
finding rather than repaired here, which would be a second change wearing this one's
commit.

<!-- xeno:section:impact -->
## Impact

Counted on this tree on 2026-10-05, as occurrences rather than files:

| Document | In the trail | Outside it |
|---|---|---|
| `ASSUMPTIONS.md` | 157 | 10 |
| `SUPPLY-CHAIN.md` | 52 | 8 |
| `CLAUSE-READERS.md` | 36 | 4 |
| `M0.md` | 29 | 5 |

274 sealed references in total. These exceed #239's table, which was counted on 2026-10-04
over a tree five intents behind this one, and the direction is the point: the sealed count
grows with every intent that cites a document, so the arithmetic for this move is strictly
worse the longer it waits.

The 27 outside the trail are this intent's work and all but one are correctable. They are
higher than #239's count of 17 because that count was of references to each document from
elsewhere, and these include the cross-references the four make to each other — `M0.md`
names `ASSUMPTIONS.md` twice, and the register names the other three in five rows.

Nothing executable is affected. No Go file, workflow step, rule or template opens any of
the four by path; the nine references in code and CI are comments. So the build, the test
suite and all five workflows are indifferent to the rename, and the 351 verdicts stand
because no artifact is touched.

What a reader loses is the two paths `README.md` offered from the root and the habit of
finding the register there. What they gain is one place to look for prose. `CLAUDE.md`
carries the register's path and moves with it, which matters more than the count suggests:
that file is sent with every request of every session, so a stale path in it is believed
until something contradicts it.

The one lasting cost is `docs/implementation-plan.md` line 2009, which will name a file
that no longer exists under that name until a person corrects it. It is inline code in a
sentence about the core's other choices, not a link, so nothing resolves it and nothing
fails; it reads as slightly out of date to someone who then cannot find the file. A94
records it beside the sealed ones so the correction is available to whoever makes the next
specification commit rather than rediscovered.
