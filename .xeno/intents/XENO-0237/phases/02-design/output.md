---
intent: github.com/triplem/xeno#202
phase: 02-design
created: "2026-10-03T19:25:06Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e12ea06c99772542d0453a615454d618cb6acd241acf33fcb553943bc995f995
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

**A file in the repository, not an issue comment.** `CLAUSE-READERS.md` beside
`ASSUMPTIONS.md`. The two have the same job — a measurement of the project kept where
the conventions are read — and an issue comment is found only by somebody who already
knows it exists.

**Four kinds, not read against unread.** A clause addressed to a person has no code
reader and is not a gap; a clause whose reader is a constant has one and is. Two
categories cannot hold both facts, so the file carries tool requirement, architectural
property, addressed to a person, and explanation.

**The two interesting kinds are enumerated; the other two are counted.** Enumerating 126
explanatory sentences would bury the eleven that matter, and leaving the counts out would
make the list unfalsifiable: a reader could not tell a pass over everything from a pass
over what caught the eye.

**Readers are named by what fails, and looked up.** "G-Schema" where a gate judges it,
"test" where only a fixture does, "nothing" where the property is a habit. The
distinction between a gate and a test is kept because a gate acts on the repository and a
test acts on a fixture.

**The findings are listed, not fixed.** Three of them — `XENO_PLUGIN_DATA`, G-Complete at
P5 only, `tool_version` self-reported — become issues after the merge.

<!-- xeno:section:alternatives -->
## Alternatives

**A gate that reads the documents and checks clause coverage.** Rejected. It would need a
machine-readable marker in prose the agent cannot edit, the gate list is a budget by
section 7, and a gate that parses English would fail for reasons that are about the
parser. The thing missing was knowledge, and this intent produces knowledge.

**A table in the issue.** Rejected for the reason the first decision gives: #202 would
close and the list would be unfindable six weeks later, which is the failure mode it
describes.

**Rows in `ASSUMPTIONS.md`.** Rejected. That file records decisions taken while building,
each a statement the project stands behind. This is a survey with a date on it, and
mixing the two makes the assumptions look provisional.

**Fixing as the pass went.** Rejected explicitly by #202. The first finding would have
taken the rest of the intent and the other ten would be unknown, which is exactly the
eight-times pattern repeating.

<!-- xeno:section:impact -->
## Impact

One new file, no code. No field, gate, tool or rule, so the second standing rule is not
engaged and no specification change precedes this.

`CLAUSE-READERS.md` is not read by anything, by design, and nothing will notice when it
goes stale. That is the same shape as the problem it documents, and the file says so
rather than pretending otherwise: it dates its pass and names the commit.

The three findings it surfaces become separate issues, each its own intent, which is
where the cost of this pass is actually paid.

`ASSUMPTIONS.md` gains one row for the choice of a dated survey over a gate, because a
later reader will otherwise ask why clause coverage is not enforced.
