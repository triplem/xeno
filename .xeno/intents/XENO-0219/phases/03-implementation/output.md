---
intent: github.com/triplem/xeno#162
phase: 03-implementation
created: "2026-10-01T16:01:09Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+94d4c6c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d05be10bba319e4b359fafffb96cdfdc61bd02d72fb18fc3ee0e807a67a5d9ef
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
---

# Implementation

<!-- xeno:section:changes -->
## Changes

One new package, 177 lines plus 230 of test. 283 lines added in `internal/gates`, 7 in its
patterns, 61 across the two existing test files, 3 in the runner, two register rows, and 408
lines of new test file for the predicates.

**`internal/git`.** `Commits(root, base, head)` runs one `git log --reverse` over `base..head`
and returns a `Commit` per record: hash, signature letter, parent count, author name and email,
subject and trailers. `Signed` turns git's letter into the boolean A70 defines, `IsMerge` reads
the parent count, and `HasTrailer` compares a key case insensitively as git reads trailers and
never looks at the value. An empty base or head is an error here rather than a default, because
section 12 forbids inferring one. A git failure comes back carrying what git wrote to standard
error, so the finding can name the ref that did not resolve.

**The separators are the ASCII unit and record separators**, not NUL. NUL is what git offers for
machine reading and it cannot be used: the format string is an argument to the process, and an
argument containing NUL is rejected by exec before git runs, with a message that names nothing.

**`internal/gates`, the package comment.** Replaced. It said gates "read, they never run
anything and never ask a model"; it now names what that protects — no build, no test suite, no
scanner, no model — and says that one subprocess exists, `git log` over the range the run was
given, because four of section 9's predicates read a history and a history is not a file a gate
could parse instead.

**`Ctx` carries `Base` and `Head`,** with the comment saying an absent range is a finding and
never a default, and `runner.go` passes them where it builds the context. That is the whole
plumbing change, and it is what the comment on those two fields has been waiting for since they
were added.

**The registry has five entries**, and an `eval` carries the run: the `Ctx`, the commits, the
error, and whether the range has been loaded. `commitsOf` loads on the first predicate that
asks, so five rules naming a commit type read one `git log`. `inRange` returns the commits a
rule judges — the range, less merge commits where the rule exempts them — or the one finding
that says the range could not be read. `exempts` and `param` read a check's parameters;
`shortHash` is how a finding names a commit, eight characters and the subject, so a reader
recognises it without looking it up.

**`section-implies-section`.** Reads `when` and `then` as section references, refuses a check
that names neither, checks both against the phase's template and reports a section the template
does not have as a configuration error in the rule, then parses the rendered artifact and
reports the one case the implication forbids: antecedent filled, consequent empty. An empty or
whitespace-only antecedent is green whatever the consequent says. `non_empty` is read and not
acted on, because it is the only form the specification gives and behaving differently without
it would be inventing a second.

**`commit-message`.** Refuses a check with no pattern or an unshipped one, naming what is
shipped — the same list `xeno check commit-message` prints — then judges every subject in the
range through `CheckMessage`, which is the one implementation section 9 asks for.

**`commit-trailer`.** Refuses a check naming no trailer, then requires the key on every commit.

**`commit-signature`.** Requires `Signed` on every commit, with the next step saying what is
checked is git's own verification.

**`approver-not-author`.** Builds the authors of the range by email and by name, lowercased,
reads this phase's `gate.yaml`, and reports every decision whose `by` is one of them, naming the
finding, the person and the commit. An unreadable `gate.yaml` is no decision to judge, and no
decision is green.

**`patterns.go`.** `conventional-commits-with-issue` beside the pattern that was there, matching
a Conventional Commits subject ending in `(#123)` or `(!123)`, with A71's reasoning in the
comment.

**Tests.** `internal/git`: a real repository per test, the range read oldest first with author
and parent count, a subject carrying pipes, tabs and quotes surviving the format, trailers by
key in four spellings, an unsigned commit against every one of git's letters, a merge commit
recognised, an absent end of the range an error in all three combinations, an unresolvable ref
naming itself, a directory that is no repository, and an empty range that is no commits and no
error. `internal/gates`: the section implication in four states, an unknown section, a check
with no sections, every subject judged, the second pattern wanting a reference, an unknown
pattern listing what is shipped, merge commits exempt and not exempt, the trailer on every
commit, the unsigned commit, `approver-not-author` by email and by name and by somebody else and
with no decisions at all, the missing range for each of the four types, and five rules producing
five findings from one `git log`. `patterns_test.go`: ten subjects against the new pattern and
the listing carrying both names.

<!-- xeno:section:deviations -->
## Deviations from the design

**The NUL separator could not be used, and the error said nothing about why.** The first version
read records with NUL between fields, which is the separator git itself offers for machine
reading. Every test that ran a commit failed with `fork/exec /usr/bin/git: invalid argument`,
which names neither the argument nor the byte. An argument containing NUL is rejected by exec
before the process starts, so the format string can never carry one. The separators are now the
ASCII unit and record separators and the comment says why, because the next person to reach for
NUL will reach for it for the same reason I did.

**A test's own fixture produced a merge conflict.** The git helper named each commit's file after
the first eight characters of its message, so "feat: on the side" and "feat: on main" wrote one
file and the merge the test needed failed to be a merge. The helper now names the file after the
whole message. The test was asserting something real and failing for a reason entirely of its
own making, which cost more time than the predicate it was testing.

**#160's test for an unimplemented type had to be rewritten.** It named `section-implies-section`
as the type nothing implements, which was true when it was written and is now false. It now names
a type no specification lists, which is the case that actually matters — the registry is a budget
and a project naming its own type still gets the finding — and a second test asserts the five
names are registered and that there are five. A test that would have gone green as the registry
filled was replaced by two that cannot.

**`section-implies-section` needed a template in its fixture, which the design did not
anticipate.** The predicate checks a rule's section names against the phase's template, so a
fixture without a vendored plugin cannot answer the question at all and the check is skipped. The
first version of the test therefore asserted a configuration error that never came, and passed a
different finding. The fixture now writes a minimal design template. Worth recording because the
skip is deliberate — a repository with no plugin should not have a rule's sections judged — and
it means the configuration error is unreachable in exactly the repositories that have no
templates.

**`causes()` does not print the next step, and two criteria are about it.** The existing helper
prints file and cause, which is what the earlier gates tests needed. Two criteria here — that an
unknown pattern says what is shipped, and that a missing range says the range is an input of the
run — are about the next step, so this file adds `nexts()` beside it rather than changing a
helper four test files read.

**The duplicated git helpers are not shared.** `internal/git`'s tests and the gates' predicate
tests each build a repository, and the two helpers are near-identical. They are in different
packages, so sharing them would mean an exported test helper or a third package, and neither is
worth it for twenty lines. Recorded because it will look like an oversight to the next reader.
