---
intent: github.com/triplem/xeno#336
phase: 04-verification
created: "2026-10-10T13:44:02Z"
schema_version: "1.0"
runner_version: dev+5f645cb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 063fafecabb56f805283544ba7d8269ceb50c371da8fe31c73efc691cca1a8fd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
Ten of eleven criteria met, one not met and recorded as not met.

The division that matters is that six of the eleven are answered by reading rather than by a
command, because the deliverable is prose and no command can answer whether a sentence says a
thing. What keeps that from being an opinion is that each of the six names the clause carrying
it — "is this choice running rather than an alternative to it" for criterion 1, "evidence for
that paragraph rather than against it" for criterion 3, "reads off an issue it has already
fetched" for criterion 4. A reviewer reading the same clause can disagree, which is the most
verification of prose admits of, and the learning records it as a template matter.

Criterion 4 was checked against the code rather than the issue. `Issue.Approval()` iterates
`i.Labels` and `i.Comments` on a value the runner already holds and returns what is missing;
it fetches nothing and nothing subscribes to it. So the paragraph's wording is exact and the
two wordings the criterion forbids are both absent.

Criterion 9 is not met: the paragraph is two sentences. It is in the mapping table as not met,
because a mapping that re-reads its own failure as success is worth less than none.

Three gaps. Nothing holds the Sources entry to the commit it pins and nothing can, since no
test here can check a claim about somebody else's tree; the six judged criteria get no second
reader before the pull request; and the absence of `.github/workflows` is a negative result
about another repository, controlled by a positive probe in the same call — 89 blobs returned
including three under `.github/` — which is the strongest form available and not the rule's
own remedy of recreating the thing.

Three sections, no open question, no decision.
