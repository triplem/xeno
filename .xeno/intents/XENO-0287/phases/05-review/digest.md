---
intent: github.com/triplem/xeno#336
phase: 05-review
created: "2026-10-10T13:45:24Z"
schema_version: "1.0"
runner_version: dev+5f645cb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: acdb72f140a1113e695ee6924416e275d0386fea87f8d959266756d85885d21f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
The review of twenty-three lines of prose in one document. Three rules answered, one met and
two not applicable, and one lens.

The lens is a citation lens, and it names what this change actually exposes: a claim about
somebody else's repository, in the present tense, in a document of this one, with nothing
able to hold the two together. Section 9 has the same exposure for OpenSpec and answers it
the same way, by pinning and dating, so the trade is established rather than new. What is new
is only that this pin is a commit and not a tag, because the repository publishes no releases,
and a commit on a one-author demo is the weaker anchor.

Criterion 9 is not met and the record says so in three places — the deviation, the mapping
table and the review answer. The claim it was written to protect is intact: the demo does not
get a section, which is what would have put it beside OpenSpec and the AI-DLC section #334 is
to write.

The context budget is exceeded again, by 663 bytes, and this time the finding is sharper than
XENO-0286's. That intent learned to add the growth; this one added it and still overshot,
because the margin was estimated as "nine or so lines" and the change was twenty-three — the
sentence was nine and the Sources entry was another fourteen, and the same intake had planned
both. The method was right and the estimate was not. Two intents have now carried this
finding to their end, which is what makes it a rule rather than a note, and the rule is to
count what will be written rather than guess, and to round up, because too large a budget
costs nothing and too small a one cannot be repaired after the intake is judged.

Two sections, three rule answers and one lens, no open question, no decision.
