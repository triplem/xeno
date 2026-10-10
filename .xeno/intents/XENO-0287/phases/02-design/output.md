---
intent: github.com/triplem/xeno#336
phase: 02-design
created: "2026-10-10T13:30:25Z"
schema_version: "1.0"
runner_version: dev+5f645cb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4d45a900747189d6ffc4bd15c64d0c5cf65e4133baf4dbe9a117cd7ea4eabbcd
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:alternatives -->
## Alternatives

**A final bullet in 4.1's "Relevant properties" list.** The cheapest edit and the wrong
place. That list is nine properties of the Software Agent SDK and the Agent Server, each a
capability the platform has; a demo somebody else built is not a property of the platform.
A tenth bullet would read as one, and the reader who scanned the list for what OpenHands
offers would come away thinking labels-as-triggers is among them. Rejected on that, which
also decides the position: the sentence goes after the list, where the list has stopped
making claims about the SDK.

**A subsection, 4.1.1 or similar.** It would hold the four claims comfortably and it would
contradict the issue's "as one sentence" and criterion 9. A subsection is also the shape
section 9 uses for a thing weighed in earnest, so it would put the demo one rung below
OpenSpec rather than in a different category altogether, which is not what it is.

**A section of its own in 4, as 4.6.** This is the reading the sentence exists to prevent:
section 4 weighs candidates for the platform, and a design running on the chosen platform is
not a candidate. Writing it as 4.6 would make it one by placement, whatever the prose said.

**In section 9 instead, beside OpenSpec.** Section 9 is "A specification framework,
declined", and the demo is not a specification framework; its `openspec/` directory is a use
of the framework section 9 already declined. Putting it there would be the same category
error as 4.6 in the opposite direction, and it would separate the sentence from the 4.1
decision it is about.

**Quoting section 12's paragraph in the sentence.** It would save the reader a lookup and it
would put a normative document's prose inside an evaluation, where it would then have two
homes and could drift. The pointer is the better form: the paragraph is the decision and the
sentence is evidence about it, so the sentence names the paragraph and the reader goes to the
document that owns it.

**Pinning the repository by tag.** There are no releases, so there is no tag to pin. The
Sources entry says what was read and at which commit, with the date, and says why it is a
commit rather than a tag — section 9's OpenSpec entry pins a tag and a registry version, so
the departure is worth one clause rather than being left as an inconsistency a reader has to
resolve.

**Citing the repository without a pin.** A URL alone, as the first nine Sources entries do
for documentation sites. Rejected because those are living documentation whose current state
is the thing cited, and this is one author's demo that was pushed to two days after the
commit read. Section 9's own entry establishes the form for a repository read at a point in
time, and this is that case.

**Mentioning the four status labels, the templates and `openspec/` in the sentence.** The
issue lists them, and the sentence that listed them all would be a list. They are what the
intake records and what the Sources entry's "what was read" names; the sentence names the
mechanism and the receiver, which are the two things the 4.1 decision is about.

<!-- xeno:section:impact -->
## Impact

One file, two places in it, and one commit.

**Section 4.1, after the bulleted list.** A short paragraph — one sentence carrying the four
claims, and the subordinate clause criterion 9 permits. It goes after the list because the
list makes claims about the SDK and this makes none; it goes in 4.1 rather than in a section
of its own because the thing being said is that the demo is an instance of what 4.1 chose,
and placement says that before prose does.

The sentence has to do four things and the order of them is the design. It names the
repository and says it runs on the Automations API, which is the not-a-candidate claim. It
says what the design is, labels as triggers with a hosted automation receiving them. It
points at section 12's paragraph as the decision that refuses exactly that, and says the
demo is evidence for the paragraph rather than against it — what the receiver costs is
visible in somebody else's repository. And it contrasts #330, where the label is a
precondition `xeno intent start` reads off an issue it already has, not an event anything
listens for.

**Sources, a new entry.** In the form the section's second paragraph already uses for
OpenSpec: the repository, the commit, the files read, the date, and the reason the pin is a
commit. It names `.github/labels.json` for the labels, the four
`automations/github/<name>/` presets and `scripts/automations/register_github_automations.py`
for the receiver, and the absence of `.github/workflows` — which is the observation that
makes "something has to receive them" literal rather than a figure of speech, and therefore
belongs in what was read.

**What does not move, and is checked rather than asserted.** `docs/process-definition.md`,
by `git diff`. Every section of the evaluation but 4.1 and Sources, by the same. No Go file,
no workflow, no configuration: `internal/model/identity.go` is in scope to be read and not
to be edited, which is the distinction the context rationale draws between what changes and
what the change is read off.

**What the sentence deliberately leaves out.** The status labels, the issue and pull request
templates, and `openspec/`. They are in the intake and in the Sources entry's list of what
was read, so a reader who wants them has them one click away; in the sentence they would
make it a list, and a list of mechanisms reads as a list of things under consideration.

**The one risk the wording carries.** A sentence that says the demo is the refused shape
running can be read as saying the demo is a mistake. It is not what the issue asks for and
not what 4.1 should say: the design is right for a repository whose receiver somebody else
operates, and wrong for this project while a phase invocation is something somebody can
type. The sentence says the latter by pointing at the paragraph that gives the reason,
rather than by characterising the demo, which is why the pointer is load-bearing and not a
courtesy.
