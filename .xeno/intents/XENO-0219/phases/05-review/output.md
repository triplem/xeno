---
intent: github.com/triplem/xeno#162
phase: 05-review
created: "2026-10-01T16:05:07Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+94d4c6c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 465204f3e03e323b9c438cf6f3595e5591d8b4a70367559342c2d28e999bfaeb
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

`given/builtin/` is still empty, so no review rule renders an entry and these carry no rule id.
The piece that writes the shipped set is the first one whose own checklist will have entries, and
it is now also the first outside reader of five predicates' findings.

**Do the five types read what section 9's table says they read?** Read row by row against the
implementation: a subject against a named pattern, a required trailer on every commit, a valid
signature on every commit, the person who decided a finding against the authors of the range, and
the section implication for the fifth. Answered yes, with two choices recorded because the table
names an input and not a boundary: A70 for what a valid signature is, A71 for what the second
pattern matches.

**Is the line about expressions held?** No rule file carries a regex, both patterns are named and
shipped, a check parameterises a type, and a type outside the five is reported rather than
evaluated. Section 9's four consequences — the failure message, comparability, cost and
behaviour, provenance — all rest on that line and none of them is undercut here.

**Is the subprocess honest?** One place starts it, the arguments are fixed in code, the only
outside values are the two refs, nothing reaches the network, a failure is a finding, and the
package comment now says what the prohibition protects rather than overstating it. The one thing
a reader has to take on trust is that `git log` over a local clone is a read, which is what the
comment says and what the arguments show.

**Does anything evaluate where it should refuse?** The two cases that could: a missing range,
which refuses per rule and reads nothing, and a type nobody implements, which #160 already
refused and this piece leaves refusing. A third case refuses more quietly than it should —
`section-implies-section` where no template loads — and it is in the gaps.

**Is the trail untouched?** No set changed, so no `rules_hash` changed. Nothing sealed was
rewritten; the demonstration that judged sealed phases ran on a copy that was deleted.
`gate verify` is at exit 0 over 193 verdicts.

**Is anything here a decision somebody else should take?** Two, both in the residual risk:
whether `commit-trailer` should read its value, which needs a format and is therefore a
specification question, and whether a range needs a bound.

<!-- xeno:section:release-notes -->
## Release notes

**`kind: checked` works.** The five predicate types section 9 names are implemented, so a rule
naming one is evaluated rather than reported as having no implementation. A project can now write
a checked rule and have it enforced.

**`section-implies-section`** holds an implication over two sections of the phase's artifact:
where the first is filled, the second has to be. An empty antecedent leaves the rule green, and a
section the template does not have is reported as an error in the rule.

**`commit-message`, `commit-trailer`, `commit-signature` and `approver-not-author`** read the
commit range the run was given. A non-conforming subject, a missing trailer, an unverified
signature and an approval given by an author of the change are each red, naming the commit by its
short hash and subject. Merge commits are exempt where a rule says `exempt: [merge-commits]`.

**A rule that needs a range and did not get one is red.** The range comes from `--base` and
`--head` and is never inferred, because a guessed range means different verdicts locally and in
CI from the same repository state.

**A second shipped pattern.** `conventional-commits-with-issue` requires a Conventional Commits
subject ending in `(#123)` or `(!123)`, which is what survives a squash merge on either host.
Both patterns answer at a gate and in `xeno check commit-message`, so a hook cannot start
rejecting what a gate accepts.

**`internal/git` is new**, and it is the only subprocess this runner starts: one `git log` over
the range, reading the clone that is already there. The package comment in `internal/gates` now
says what being network free and deterministic forbids — no build, no test suite, no scanner, no
model — rather than forbidding every subprocess.

**A project with no rules is unaffected**, at every phase.

**What is still to come in WP4.** The shipped set under `given/builtin/`, the examples under
`examples/rules/` and the two hook templates under `examples/hooks/`, and external gates.

Closes #162. Refs #1.

<!-- xeno:section:residual-risk -->
## Residual risk

**The signature predicate has never seen a valid signature.** Every case is an unsigned commit
or a letter set by hand, because signing in a test needs a key and A21 says this project does not
sign yet. A70's boundary is tested from both sides as a function and the path where git reports
`G` over a real commit has never run. The first project to adopt the signature rule is the test.

**A green commit rule is evidence from the run, not from the record.** The range is an input and
the artifact does not carry it, so nobody can re-derive later which commits a verdict was taken
over. `gate verify` recomputes hashes and does not re-evaluate predicates, which is why it still
passes; a reader who takes a green `commit-message` as a statement about a known set of commits
is taking more than the record offers. This is the deliberate consequence of not recording a
range that a squash would invalidate, and it belongs in section 16's list of what a green gate
does not mean — which is a person's commit.

**`commit-trailer` proves a habit and not a reference.** The value is never read, so
`Xeno-Intent: anything` passes. Reading it needs a format, which is an expression in all but
name, so the proposal belongs with the shipped set where the example lives. Until then a project
adopting the trailer rule gets a check that every commit has the field and no check that the
field says anything.

**The configuration error for a section name is unreachable where there is no template.** A rule
naming a section that does not exist is caught only if the phase's template loads, so the
repositories most likely to have mistyped a section — ones that have not vendored a plugin — are
the ones where nothing checks. Cheap to improve and not improved: the alternative is reporting
every rule's sections as unverifiable, which would be noise in every repository before `xeno
init`.

**Nothing bounds the range.** One `git log` reads every commit between two refs into memory. No
limit, no measurement, and the only range this project has ever passed is a branch. A rule
adopted on a repository whose `--base` is an old tag would read everything since that tag.

**The tests skip where git is absent**, which is the price of one dependency. CI has git because
it clones, so nothing is hidden in the pipeline, and what a developer without git sees is a pass
over skipped tests.

**Five predicates' findings have been read only by their author.** Six wordings for the commit
types and three for the section implication, all written and reviewed by me. The shipped set is
where that stops, which is the second piece in a row to end with that sentence — and a reason to
treat the shipped set as the piece where wording is part of the acceptance rather than a
by-product.
