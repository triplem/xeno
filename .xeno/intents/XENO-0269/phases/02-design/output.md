---
intent: github.com/triplem/xeno#207
phase: 02-design
created: "2026-10-07T09:06:25Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8e9aba80b8f2d8ab3b512d07a577b6d9372efb9facaa266cd5b45e58bd894951
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - chosen: 'Record the triple as a declaration rather than corroborating one of the three: a paragraph in section 12 saying where each value comes from, what G-Schema does with them, why the runner cannot corroborate one, why a check against project.yaml is refused, what the gateway''s record is and why it is out of reach, and that the register is read as what an agent declared.'
      decided_by: Markus M. May
      id: D-1
      proposed_by: claude-opus-5
      rationale: 'Corroboration is not available inside the boundary this project drew. Section 7 keeps the runner from branching on the harness, so it cannot ask the harness anything beyond what the harness volunteered, and a CI check asserts the name appears in one constant and in no condition. The only other copy of tool and model in the repository is project.yaml, which is the same declaration in a second place; #207''s own last sentence names a check against it as what should not happen, because it would make the gap invisible instead of closing it. The one record from another hand is a gateway''s, and the plan''s section 9 says both what it is worth, a routing record rather than an attestation from the provider, and why it is out of reach: the runner makes no model request in v1 and xeno gate never touches the network, which is what makes a verdict reproducible on any machine. The gateway comparison was put as the second option with that cost, that it needs v2''s gateway and could only be a figure outside the gates, and was not chosen. The third option, leaving the issue open, was put with the cost that a reader of section 12 gets no warning and the register keeps being read as if the triple were measured.'
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**Where the paragraph goes.** Immediately after the one that says the triple is the raw
material a project needs for a provider register. That paragraph is the only place in the
document that tells a reader what the three fields are for, so it is where the question of what
they are worth arises. The alternative was the **Tools** paragraph above it, which says the
tool and its version are recorded "as a record rather than a variable" — close, but that
sentence is about them not changing while a project runs, and a qualification there would read
as being about stability rather than about provenance.

**It names each of the three separately rather than saying "all three are self-reported".** A
reader who has to act on this needs to know which file to look in, and the three are not alike:
two are project level declarations and one is passed per run or per session. The sentence is
longer for it and the alternative tells somebody nothing they can use.

**The reason the runner cannot corroborate is cited, not asserted.** "Section 7 keeps it from
branching on the harness" rather than "nothing can be done". An asserted impossibility invites
somebody to try; a cited rule tells them what they would have to change first, and that is the
honest shape, because the rule is a choice this project made and not a law.

**The refusal of the second-copy check is in the paragraph and not only in the issue.** `tool`
and `model` are also in `project.yaml`, and the obvious thing to build is a comparison. #207's
last sentence warns against it. Putting the warning in the document rather than leaving it in
the issue is the difference between a reader finding the reason and a reader having the idea.

**The gateway is named, with what it is worth and why it is out of reach.** The weaker version
of this paragraph says nothing can corroborate the triple, which is false in a way that matters:
a record from another hand exists, it is just not one the gate path may read. Naming it and
pricing it is what stops the paragraph being read as "corroboration is impossible" when what is
true is "corroboration is outside this boundary, deliberately". The plan's own phrase, a routing
record rather than an attestation, is used so that the two documents do not drift.

**It ends on how the register is read, in the issue's own terms.** #207 asks for the register to
be "read in that light", and that is the sentence, because it is the consequence a reader of a
provider register needs and the rest of the paragraph is how it is arrived at.

**No em dash.** The document uses one in 132 kilobytes. The first draft had two, and both were
replaced with a colon and a full stop when the count came out.

**The clause is counted in `docs/clause-readers.md` and gets no row.** Nothing can violate it by
writing the wrong thing, because what it states is an absence, so it is an explanation in that
document's four kinds. The count moves and a paragraph says which clause moved it, which is
exactly the shape #267's three clauses were given there. A row would have claimed something
fails if the clause is broken, and nothing would.

**One row in the register, naming `plugin_version` as the contrast.** The decision outlives the
intent, and the thing that makes it legible to a later reader is the comparison #207 drew
itself: `plugin_version` had this shape and now has an anchor compiled into the runner that the
checked side cannot influence. A row that only said "the triple is self-reported" would repeat
the document.

<!-- xeno:section:alternatives -->
## Alternatives

**The gateway comparison.** The one corroboration that would be worth something, because the
gateway's record comes from another hand: it names the deployment a request was routed to, in
`x-litellm-model-name` and in the spend log row, which is not a second copy of what the agent
said. Three things cost it. It needs a gateway, and section 12 puts storage behind one in v2
while v1 keeps full text local. It could only ever be a figure and never a gate, because
`xeno gate ...` never touches the network and that promise is what the governance rests on. And
the plan already says what it would prove: a routing record rather than an attestation from the
provider, so a comparison would show which deployment answered and not what served the tokens.
Put as the second option with those costs and not chosen.

**A check comparing the artifact against `project.yaml`.** The cheapest thing that looks like a
check, and the one a later reader will reach for first, which is why it is refused in the
document rather than merely not built. `tool` and `model` are in both places because the runner
copied one into the other, so the comparison is a copy against its source. #207's last sentence
is the argument: it "would make the gap invisible instead of closing it", and a green G-Schema
is worth less afterwards than it was before.

**A gate result or a finding saying the field is unverified.** It would put the fact where a
reader of a verdict sees it rather than where a reader of the specification does. It needs a
result outside the four A4 and A42 fix, which is a specification change of its own, and it
would fire on every artifact of every phase forever, which is the shape #235 was about and
XENO-0265 answered. A permanent finding is a thing people learn to scroll past.

**Weakening the fields instead: dropping `tool_version`, or making the triple optional.** The
honest reading of "enforced and meaningless" taken all the way. It fails because the fields are
not meaningless: a project that keeps a provider register needs them, section 12 says so, and
the one thing wrong was what a reader took them for. Removing the raw material to fix the label
is the wrong half.

**Saying nothing until corroboration exists.** The third option put. It keeps the issue open as
an honest record that nobody has answered, and it costs a reader of section 12 any warning at
all, which is the state #207 was filed about. Not chosen.

**Asserting that nothing can corroborate the triple.** Not an option so much as the weaker
paragraph that nearly got written. It is false in the way that matters: a record from another
hand exists and is outside a boundary this project drew on purpose. The paragraph says that
instead, which is longer and true.

<!-- xeno:section:impact -->
## Impact

`docs/process-definition.md`, section 12: one paragraph, 12 lines. Nothing is deleted and no
existing sentence is reworded.

`docs/clause-readers.md`: the explanation count moves from 129 to 130, and a paragraph says
which clause moved it and why it has no row.

`docs/assumptions.md`: one row, with the decision, who took it and when, the option not chosen
with its cost, and `plugin_version` named as the contrast.

Nothing else. No Go file, no gate, no field, no artifact schema, no skill.

**For a reader of a verdict.** Nothing. A green G-Schema means what it meant yesterday. That
is the point of putting this in prose rather than in a check: the gap was never that a gate was
wrong, it was that a reader of the register could not tell what the fields were worth.

**For a project with a provider register.** The one thing it needed to know and could not get
from this document: the triple is what an agent declared. A project that links its register
from its README, which section 12 describes as the whole extent of the connection, now links
something whose provenance is stated.

**For the trail.** Nothing moves. `xeno gate verify` recomputes every sealed verdict and
matches, because the documents are read by context scopes and each lock is compared against
what it recorded rather than against today's tree.

**For #207's own finding.** It closes with the branch the issue itself preferred, and the
reason corroboration was unavailable is in the document rather than only in the issue, so the
next person to reach for a check finds the argument instead of having the idea.

**What is left for later, named and not assigned.** The gateway comparison. It has a reason to
exist, a cost, and no package in v1, and this intent does not give it one because assigning
work is the plan's job. The register row carries it so that the figure is findable when a
gateway exists.
