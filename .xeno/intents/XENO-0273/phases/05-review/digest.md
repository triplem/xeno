---
intent: github.com/triplem/xeno#153
phase: 05-review
created: "2026-10-07T14:30:55Z"
schema_version: "1.0"
runner_version: dev+6b48c17.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e15e2eabf0fc5e9dd9a0236e672dfbefdd3951b3f08a9d88a67a4f67d6a2f8ec
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three standing rules hold: section 5 and Appendix A are cited and not touched, the page's
own test was that no sentence may state a field or a default, and the change belongs to WP16
and to intent XENO-0273 for issue #153.

The format is now a page in `docs/`, linked from the index, carrying the annotated file with a
test that fails if the two part. The issue's warning about a second copy is honoured rather
than worked around, which is why there is no field table.

What a reviewer should look at first is the three paragraphs, because they are the only thing
that can be wrong and nothing holds them afterwards. The sentence to re-read is the one sending
a reader to section 5 for the format and Appendix A for the two keys: the intake's first draft
had it wrong, saying "section 5's `index` block", which does not exist.

The second thing is that #153 closes with half of its own analysis declined. Its "what it needs
to say" section reads like a specification for the page that was not written and its last
paragraph argues against writing it, so the issue disagrees with itself and the maintainer
resolved it.

Two questions no rule asks are answered: whether a published YAML comment block is
documentation, which rests on #151's annotation being good enough and on a thin description
beating two; and that two committed copies with tests now stand where WP16 wants a generator,
with each intent's alternatives the only thing saying which is the stopgap.

Three rules answered: the traceability rule met, since P3 recorded no deviation and says so,
and two not-applicable.
