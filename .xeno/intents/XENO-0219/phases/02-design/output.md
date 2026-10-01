---
intent: github.com/triplem/xeno#162
phase: 02-design
created: "2026-10-01T15:51:59Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e1a37e0184327a264b7a6c84f47b57e3bc4673d9df7c8c7b336c04bd5f807f74
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**`internal/git` is a package of its own, and it is the only place a subprocess is started.**
One function, `Commits(root, base, head)`, returning a hash, a signature status, a parent count,
an author and an email, a subject and the trailers, per commit. The arguments are fixed by the
code and the only values that come from outside are the two refs, which the run supplies. A gate
that shelled out inline would put process handling in the middle of a verdict and make the
prohibition in the package comment untestable.

**The package comment in `internal/gates` is corrected rather than bent.** It said gates "read,
they never run anything and never ask a model". It now says what that protects: no build, no
test suite, no scanner and no model, and the one subprocess is `git log` over the range the run
was given, which is reading the clone that is already there. The sentence was written when the
only runnable thing was a build.

**One `git log` per gate run, not one per rule.** The evaluation context loads the range lazily
on the first predicate that needs it and keeps the result, including the error. Five rules naming
`commit-message` read one subprocess. This is the shape #160's gap asked for and did not get for
the rule tree, so it is taken here where it is cheap.

**A missing range is a finding before anything else is read.** Where a rule needs commits and the
run carries no `--base` or `--head`, the finding names the rule and says the range is an input of
the run. Nothing is read, nothing is assumed empty, and the four commit types share one check so
they cannot disagree about it.

**A failure from git is a finding on the rule that asked.** Not a panic, not a gate that refuses
to run: the cause names what was asked — the range — and what came back, and the next step is to
pass a range that resolves in this repository. A gate whose verdict depended on the exit status
of a subprocess failing quietly would be the worst of both.

**A valid signature is `G` or `U`.** Git answers with one of five letters, where `G` is a good
signature and `U` is a good signature from a key the local configuration does not trust. Trust is
explicitly not this gate's question — the non-goals say so and section 9 says the host that
authenticates pushes answers most of it — so what is checked is that the commit carries a
signature that verifies. `B`, `X`, `Y`, `R`, `E` and `N` are red, which covers bad, expired,
revoked, missing key and unsigned.

**`section-implies-section` reads the rendered artifact and treats an empty antecedent as true.**
`when` empty or absent makes the rule green whatever `then` says, because an implication with a
false antecedent is true; the alternative would make every such rule a requirement that every
section of every template be filled. A section that the phase's template does not have is a
configuration error in the rule rather than an empty section, red, naming both.

**A trailer is matched on its key, case insensitively, and the value is not examined.** That is
how git reads trailers and it is all section 9 asks: a required trailer on every commit. What the
value says is the rule's business and not the predicate's, and a predicate that validated it
would need an expression, which is the line the whole section draws.

**`approver-not-author` reads the previous run's `gate.yaml` and compares strings.** Section 9
says both: it evaluates on the run after the decision, and it compares two self-asserted strings.
No decisions is green, because a rule about who may approve is not a rule requiring an approval.
The comparison is on the author's email where the decision carries one and on the name otherwise,
lowercased, which is the most a string comparison can honestly do.

**The second pattern is `conventional-commits-with-issue`**, matching a Conventional Commits
subject that ends in a parenthesised reference, `(#123)` or `(!123)`. GitHub builds a squashed
subject ending in `(#123)` and GitLab in `(!123)`, which is exactly the case section 9 names, and
accepting both is what keeps the pattern shipped rather than per host. Recorded as an assumption,
since the section says what the pattern is for and not what it matches.

<!-- xeno:section:alternatives -->
## Alternatives

**A git library instead of a subprocess.** `go-git` reads a repository without starting a
process, which would keep the package comment true as written and make the tests independent of
whoever's git is on the path. Rejected on the dependency: one vendored dependency is the
project's standing position, a second is a decision rather than a step, and the thing being read
here — subject, trailers, signature status, author — is what `git log --format` hands over in one
call. The cost of the rejection is that the tests need a real git, which is in the gaps.

**Read `.git` directly.** No dependency and no subprocess. Rejected as absurd for packfiles and
signatures, and it would make this repository the owner of a format it does not control.

**Treat a missing range as an empty range.** Every commit rule would pass where no range was
given, which is convenient locally and is the precise failure WP4's done-when names. Rejected,
and the criterion for it was written per type rather than once so that the rejection cannot be
undone by accident.

**Accept only `G` for a signature.** Strictly stronger: a signature from an untrusted key would
go red. Rejected because trust is a property of the verifier's keyring and not of the commit, so
the verdict would depend on whose machine ran the gate — the same class of problem as a guessed
range. A project that wants the stricter reading wants a key policy, which is an external gate.

**Make an empty antecedent red.** It would catch a rule whose `when` section was mistyped, since
a section that never fills makes the rule vacuous. Rejected because it inverts an implication:
the rule says "a change to a published interface must come with a migration note", not "every
phase must change a published interface". The mistyped-section case is caught instead by the
configuration-error finding when the section is not in the template at all.

**One `git log` per predicate.** Simpler: each predicate asks for what it needs. Rejected for the
cost it hides — five rules naming a commit type would start five processes per phase — and because
the error would then be reported five times for one failure.

**Validate a trailer's value.** `Xeno-Intent: XENO-0219` could be checked against the intent the
phase belongs to, which would make the trailer prove the thing G-Trace proves from the other
side. Rejected: it needs a format for the value, which is an expression in all but name, and
section 9's `commit-trailer` reads "a required trailer on every commit in the range" and nothing
more. The proposal belongs in the shipped set's discussion, where the `Xeno-Intent:` example
lives.

**Resolve author identity.** Mapping a git author to the person who approved, through the host's
user list, would make `approver-not-author` mean what its name suggests. Rejected because it
needs the network, which a gate may not have, and because section 9 states the limit itself: the
rule enforces the discipline of a team that means it.

<!-- xeno:section:impact -->
## Impact

**`kind: checked` becomes usable.** A rule naming one of the five types is evaluated instead of
reported, so the choice between `checked` and `review` becomes the choice section 9 describes
rather than a choice between a red gate and a person's answer. This is the piece that makes the
shipped set writable.

**A gate starts a subprocess for the first time.** The prohibition in `internal/gates` moves from
"never runs anything" to a named list, and the one exception is a read of the local clone. Anybody
auditing the network-free claim now has a second thing to check, and the package comment is where
they will look.

**Four of the five types make a verdict depend on an input the artifact does not carry.** The
range is passed in and deliberately not recorded, so a phase judged against `commit-message` can
be re-judged later only by somebody who knows the range. That is already true of `gate verify` in
CI, where the wrapper passes both ends, and it becomes visible here because a rule can now turn a
phase red over it.

**`approver-not-author` makes a verdict depend on the previous run.** The decision it reads was
written by `gate approve` or `gate override`, so the rule is evaluated one run later, which the
decision's own commit triggers. A reader of a red verdict from it is looking at a comparison
between two files written at different times by different writers.

**The registry stays a budget.** Five keys, named in the specification, and a sixth name still
gets #160's finding. Nothing about this piece makes a project-defined type easier to add, which is
the point: section 9 puts them outside v1 and names the external gate as the way out.

**The shipped patterns become two**, and both answer at the gate and in `xeno check
commit-message`. A project adopting the second one gets the same message from a hook and from a
gate, which is the property section 9 wanted a named pattern for.

**What the next piece inherits.** Five types with their parameters settled by tests rather than by
guess, two patterns, and one open question it has to answer rather than inherit: whether the
documentation consistency rules ship enabled or as examples, which the plan leaves to this work
package and which only makes sense once there is a set to put them in.
