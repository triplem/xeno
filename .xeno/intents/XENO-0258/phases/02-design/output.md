---
intent: github.com/triplem/xeno#256
phase: 02-design
created: "2026-10-06T08:35:54Z"
schema_version: "1.0"
runner_version: dev+5163d1b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fd8dd923fc4b5a863b7fc90f140d608d8ed94cc1e5cc45211bc7250f53a9cdda
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

The paragraph is framed on sealing and not on carefulness, and that decision is what makes the
addition defensible. "Verify a negative before you trust it" is true of any repository; "a finding
goes into an artifact that is sealed, so an unverified negative is permanent and the correction has
to live somewhere else" is true of this one. `CLAUDE.md` carries this project's context, so the
second form belongs and the first would be the first line in that file about nothing in particular.

It goes under Conventions beside the sequencing paragraph, with no new heading. Both are about how
the work is done rather than about what the artifacts hold, the heading already covers that, and a
third section for two paragraphs would be organising a file that asks to stay short.

The instance is named and the second one is not. The `git check-ignore` run against a path deleted
moments earlier, and that the claim reached two sealed phases and a pull request, is what makes the
rule checkable — a reader can go and look. Naming #212's field read as well would double the length
to add a second proof of the same point, and #256 carries both with the table.

The paragraph says nothing checks it. The sequencing paragraph ends that way and A90 is why: a
reader which cannot fail is worse than none, and a convention that implied enforcement would be the
shape this repository keeps cataloguing. Saying it also stops somebody later looking for the gate.

The rule is stated as an act rather than as a caution. "Recreate it and re-run" is something a
session does; "be careful about negatives" is something it agrees with. The gitignore check took
three seconds to redo and that is the whole of the ask.

No row in `docs/clause-readers.md`, for the third time in three intents and on the same reasoning:
the audit is a pass over the two normative documents, and a convention in `CLAUDE.md` is a clause in
neither. #255 drew that line for the sequencing convention; redrawing it identically is what makes
it a rule rather than a judgement made twice.

No register row. A97 and the rows around it record decisions about how the repository is built; this
is a convention about how an agent works, which `CLAUDE.md` is the file for, and putting it in both
would be two places to keep true.

<!-- xeno:section:alternatives -->
## Alternatives

Recording the decision not to write it down was the issue's other branch and is the one this was
weighed hardest against. `CLAUDE.md` asks to stay short, this session already added nine lines to it,
and a learning sitting in `learning.yaml` is at least honest about having no reader. It was rejected
because the cost is asymmetric: the paragraph is four lines read before every request, and the
failure it prevents put a false sentence into two sealed phases where it cannot be corrected. A
learning nothing reads would have left the next instance to the same three-second check nobody ran.

Leaving it as general advice outside the repository was considered — the lesson is real and not
specific to Xeno. Rejected on where a finding goes: in most projects an unverified negative is a
mistake somebody corrects, and here it is sealed. That difference is the only reason this earns a
place in a file about this project, and it is also what the paragraph has to say.

A `CONTRIBUTING.md` paragraph was considered. That file addresses a person opening a pull request,
and the failure is an agent writing a finding into a phase; `CLAUDE.md` is the file sent with every
request, which is where the reader is.

A row in `docs/clause-readers.md` was considered and rejected for the third time on the same
reasoning as #255's: the audit is a pass over the two normative documents. A convention in
`CLAUDE.md` would be the audit describing something outside its own subject, and that document has
already needed one correction this session.

A checked rule was considered briefly and is not expressible. A rule would have to know that a
command's silence meant "absent" rather than "no match", which is not in the artifact, and A90 is
the standing answer to a reader that cannot fail in the interesting case.

Naming every tool with the ambiguity — `git check-ignore`, `test -f`, `grep -l`, `find` — was
considered and makes the paragraph a table. The class is one sentence and the table belongs in the
issue, which has it.

Adding it to the shipped skills under `.xeno/plugin/` was considered, since those are prose an agent
loads and this is advice to an agent. Rejected because the plugin is vendored into every adopter and
this is a convention about this repository's trail; A72's reasoning for `examples/` over the shipped
set applies to prose as much as to rules.

<!-- xeno:section:impact -->
## Impact

One file, one paragraph. `CLAUDE.md` goes from 87 lines to about 91.

Nothing executable changes and nothing could. No gate, rule or Go file reads `CLAUDE.md`; it is
prose sent to a model, which is the only mechanism this intent has and the strongest one available
for a convention of this kind.

What changes is what a session reads before it investigates. That is the entire effect and it is not
measurable from here: whether the next false negative is caught depends on a model reading four lines
and acting on them, and no verdict will record either outcome. The paragraph is the second in that
file with that property — the sequencing one from #259 is the first — which is worth noticing because
a file whose instructions cannot be checked is a file that can quietly stop being true.

What the trail gains is nothing, deliberately. The claim in XENO-0249 stays false and sealed; this
does not touch it and #254's residual risk already records that a reader of that intent meets it with
no correction beside it. What the repository gains is that the next intent is likelier not to produce
a second one.

`CLAUDE.md` is now carrying two conventions that arrived from this session's own mistakes rather than
from the plan. That is the honest reading of the file's growth and it cuts both ways: the conventions
are earned by instances rather than invented, and a file accumulating them at this rate will need
somebody to decide what it is for. Its last line says keep it short, and nothing enforces that
either.

The third place this reasoning has now declined a row in `docs/clause-readers.md` is itself a small
result: the boundary between a normative clause and a project convention has been drawn the same way
three times, which is about as close to a rule as an unenforced distinction gets.
