---
intent: github.com/triplem/xeno#120
phase: 02-design
created: "2026-09-29T05:59:42Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ad0a764.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 2c0bf120da9cb563376a6beca388eb69a2fa8ace389f50ea01725bda9f04d0e4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: by-hand
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**A package of its own, `internal/secrets`.** It assembles the effective filter, reduces
it to a hash and redacts a text. Section 4 says G-Secret and the digest writer read the
same file, so the loading and the matching belong beside neither of them: a gate that
will be written later reads this package rather than copying it, which is what makes the
shared rules verifiable rather than stated.

**The effective filter is the union, patterns sorted by id then regex.** A project adds
and removes nothing, which section 4 requires, and matching is a disjunction so a
duplicate id can only widen what is caught. Sorting is what makes the set a set rather
than two concatenated files.

**`secrets_hash` covers the set's content, in a canonical rendering.** One line per
pattern, `pattern\t<id>\t<regex>\n`, then one per path, `path\t<glob>\n`, both sorted,
sha256 over the whole. Not the bytes of the files: a comment edited in the shipped
filter would otherwise change every digest's hash without changing what any digest
passed through, and two projects with the same effective filter written in different
orders would disagree. Appendix B delegates this and it is recorded as an assumption.

**Redaction replaces the match and leaves the line.** The replacement names the pattern
that fired, `[redacted: <id>]`, so a reader of a digest sees that something was removed
and by which rule. A digest is read by a person, and a silent deletion would make the
sentence around it a lie.

**No filter means no hash and no filtering.** A repository before its plugin is vendored
has no shipped file. The field is absent, the summary is copied, and nothing is refused.
The runner does not write `by-hand` in its place: that value says a person stood behind
it, and here nobody did.

**The filtering happens in `writeDigest`, between the summary and the file.** Section 16
puts it there and gives the reason this intent exists: the runner holds the text, so the
filtering is outside the model's reach.

**`output.md` is not filtered.** Section 16 puts the filtering on the digest, which is
the file that leaves the session. Filtering the agent's own prose would redact a
discussion of a pattern by the pattern, and this record is the proof: it quotes
`secrets_hash: by-hand` and would have been mangled.

**The `by-hand` honesty rule is not tightened, and the number is in the row.**
Eighty-seven verdicts diverge, measured by removing `secrets_hash` from `WriterlessHash`
and running `gate verify`.

<!-- xeno:section:alternatives -->
## Alternatives

**Hash the two files' bytes in a fixed order.** Three lines instead of twenty and it
fails AC3 and AC4: a comment moves the value, and a project writing the same patterns in
another order disagrees with one that wrote them sorted. The field would then say which
files existed rather than which filter a digest passed through, and Appendix B's word
for what it covers is "effective set".

**Put the loader in `internal/runner` and the matching in the digest writer.** Fewer
files. It also means G-Secret, when it is written, either imports the runner or copies
the patterns, and section 4's claim that the two read the same file becomes something a
reader has to verify by comparing two implementations.

**Redact the whole line, or the whole file.** Safer against a pattern that matches only
part of a secret. It also destroys the summary: a digest is the only thing that leaves a
session, and one whose lines vanish silently is worse than no digest. Rejected because
the digest's readability is the reason it exists, and a named replacement makes the
removal visible instead.

**Tighten `by-hand` and re-run the gates on all sixty intents.** What Appendix B's
wording actually asks for once a writer exists. It turns eighty-seven verdicts red, and
the artifacts cannot be corrected because `secrets_hash` sits inside `artifacts_hash`:
fixing them would rewrite sealed history and every merge commit that names it. Rejected
on the measurement rather than on the principle, and the principle is recorded as
unfinished.

**Ship a large pattern set from a public collection.** More secrets caught on day one. A
regex that matches too much redacts prose and teaches a reader to distrust the
redaction, and nobody in this repository has read those collections line by line.
Rejected in favour of a small set and a documented route for a project to add to it.

**Write `by-hand` where no filter exists.** It would make the field present everywhere
and every gate green. `by-hand` says a person stood behind the value; where no filter
exists nobody did. Rejected as the same lie the previous intent refused about a hash
over nothing.

<!-- xeno:section:impact -->
## Impact

`.xeno/plugin/secrets.yaml`: the shipped filter, small, with a comment naming the
project file as the way to add to it.

`internal/secrets`: `Filter`, `Load`, `Hash`, `Redact`. About sixty lines and its own
tests.

`internal/runner`: `writeDigest` redacts and writes `secrets_hash`; `SectionSet` writes
the field.

`ASSUMPTIONS.md`: A35 amended to two fields, and a new row for the computation Appendix
B delegates, with the `WriterlessHash` measurement in it.

What a reader gains: a digest that passed through a filter, and a field that says which
filter. What they still do not gain: a gate that checks artifacts for secrets, a
`tool_version`, or a `rules_hash`.

What this costs: every digest written from here carries a hash that changes when the
shipped filter changes, which is the point of the field and also means a filter edit
makes new digests disagree with old ones. That is what the field is for, and section 4's
word for it is reconstructable.

Sixty intents keep `by-hand`, and `gate verify` is the check that nothing moved.

One risk worth naming now: the shipped pattern set is the security surface of this
change, and it is small. A secret whose shape nobody anticipated passes through the
filter into a digest exactly as it does today. The field will say which filter it
passed, which is how somebody later finds out what was not caught.
