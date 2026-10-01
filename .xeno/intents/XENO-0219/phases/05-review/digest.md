---
intent: github.com/triplem/xeno#162
phase: 05-review
created: "2026-10-01T16:05:37Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+94d4c6c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 465204f3e03e323b9c438cf6f3595e5591d8b4a70367559342c2d28e999bfaeb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Reviewed row by row against section 9's table of five types and what each reads, with two boundaries
recorded because the table names an input and not a limit: A70 for what a valid signature is, A71 for
what the second pattern matches. The line about expressions holds — no rule carries a regex, both
patterns are named and shipped, and a type outside the five is reported rather than evaluated. The
subprocess is honest: one place starts it, the arguments are fixed in code, the only outside values are
the two refs, a failure is a finding, and the package comment now says what the prohibition protects
instead of overstating it. The release: `kind: checked` works, the four commit types read the range the
run was given, a rule needing a range it did not get is red, and `conventional-commits-with-issue`
ships. Residual risk: the signature predicate has never seen a valid signature, because signing needs a
key A21 says this project does not have; a green commit rule is evidence from the run and not from the
record, since the range is deliberately unrecorded, which belongs in section 16 and is a person's
commit; `commit-trailer` proves a habit and not a reference; the section-name error is unreachable
where no template loads; nothing bounds a range; the tests skip without git; and five predicates'
findings have been read only by their author, which is the second review in a row to say so.
