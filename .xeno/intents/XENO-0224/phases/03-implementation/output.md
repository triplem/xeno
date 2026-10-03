---
intent: github.com/triplem/xeno#1
phase: 03-implementation
created: "2026-10-03T10:22:59Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4e9938bff7762cfb5d4b96a710555c5c83b30e8ddf824d101b15fee8708f82d5
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

One file, 19 lines, and an issue closed. There is no code in this intent.

**`ASSUMPTIONS.md`, the header.** Five paragraphs where there was one. The first says the record
is closed, cites the plan's section 4 for why M0 is where it closes, names A77 as the last row and
this intent as what closed it, with the date. The second says what replaces it — a decision in the
design phase of the intent that took it, a learning through section 10's merge request — and states
the cost of the switch in the same breath: a row in a table was findable and a paragraph in one
intent's P2 is not. The third says every row stays with its state column, and names four of them
with what a reader learns from each, so that "closed" cannot be read as "finished with". The
fourth and fifth are the original header, moved behind the closure and put in the past tense.

The rows are untouched. The table is untouched. Nothing in the file was deleted.

**The issue.** #1, M0's own issue, closed with what was checked rather than with a sentence: the
branch protection as the host reports it, the job and what its verdict covers, the counts, and
which of the two readings of M0 each piece of evidence answers.

**What was measured for the closure**, on this branch:

- `verify` is a required status check on the default branch, with `enforce_admins: true`. No merge
  reaches `main` without a recomputed verdict, and no administrator can bypass it.
- that check runs `./xeno gate verify`, which is at exit 0 over 216 verdicts.
- twenty-four intents have run all six phases; the trail holds seventy-six that ran the intake
  alone, which `M0.md` and A20 explain.
- eleven of fourteen gates are implemented. The plan's M0 paragraph expects six to be missing;
  three are.
- the agent layer exists for one harness: seven skills, a manifest the client validates, hook
  wiring, and a measured cost of about 606 tokens a session.

**What this intent does not touch, and says so.** `docs/` — the plan's two dated paragraphs are
named in the closure and owed to a person. The `Xeno-Intent:` trailer, which is the next intent.

<!-- xeno:section:deviations -->
## Deviations from the design

**Two project decisions arrived mid-intent and this intent is not where they go.** While P2 was
deciding where decisions live after the register closes, the maintainer settled two: that the
shipped rule set is closed at four rules for v1 with everything else routed to `examples/rules/`,
and that the `by-hand` placeholder rule stays as A62 and A66 left it, with the fact recorded in
section 16 that two hash fields accept it for artifacts written before their writers existed.

Neither belongs to this intent's subject, and the file that would have taken them as rows is the
one being closed. So the first is carried to the next intent that touches the rule set, which is
the trailer adoption, and the second is a `docs/` edit and therefore a person's. Both are named
here and in the review, because a decision held in a conversation is a decision nobody can find.

This is the cost of the switch, arriving within the hour of the switch being made, and it is
recorded in P2's learning rather than argued with: the register could take anything, a phase
artifact can only take what its phase is about.

**The header needed rewrapping by a script and the first version was over the column limit.** I
wrote five paragraphs by eye at about ninety-four characters against a convention of eighty-eight,
which is the same slip as #169's skills. The register's table lines are far longer and are left
alone, so the check had to be restricted to the header rather than run over the file — which is
also why the slip survived a casual look at `awk`'s output.

**The artifacts of this intent are the shortest of the twenty-four and still about forty thousand
bytes.** P1 made shortness a constraint and the sections are visibly shorter, and the floor is the
sequence itself: six phases, five files each, a digest and a learning record apiece. What an author
controls is the length of what they write, which is what was tested here; what they do not control
is the shape, which is most of the figure on #117.
