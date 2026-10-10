---
intent: github.com/triplem/xeno#336
phase: 02-design
created: "2026-10-10T13:31:13Z"
schema_version: "1.0"
runner_version: dev+5f645cb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4d45a900747189d6ffc4bd15c64d0c5cf65e4133baf4dbe9a117cd7ea4eabbcd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
Eight alternatives for one sentence, and most of them are about where it goes rather than
what it says, because placement makes the claim before the prose does.

The one that nearly won is a tenth bullet in 4.1's "Relevant properties" list — the
cheapest edit available. It is wrong because that list is nine capabilities of the Software
Agent SDK, and a demo somebody else built is not one of them; a reader scanning the list
for what OpenHands offers would come away thinking labels-as-triggers is among them.
Rejecting it also settled the position: the sentence goes after the list, where the list has
stopped making claims about the SDK. The other three placements rejected — a 4.1.1, a 4.6,
and a home in section 9 — each make the demo a candidate or a specification framework by
where they put it.

The design is a paragraph in 4.1 and an entry in Sources. The Sources entry pins a commit
rather than a tag because the repository publishes no releases, and the entry says so,
since section 9's OpenSpec entry pins a tag and a registry version and an unexplained
departure is an inconsistency a reader has to resolve.

One risk is named in the impact section: a sentence saying the demo is the refused shape
running can be read as saying the demo is a mistake. It is not, and 4.1 should not say so —
the design is right for a repository whose receiver somebody else operates. The sentence
carries that by pointing at the paragraph which gives the reason rather than by
characterising the demo, which is why the pointer is load-bearing and not a courtesy.

Two sections, no open question, no decision, no learning this phase.
