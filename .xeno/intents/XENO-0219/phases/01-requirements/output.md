---
intent: github.com/triplem/xeno#162
phase: 01-requirements
created: "2026-10-01T15:50:45Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 967407d2d2819f32b34f62cf1c160a4da60ed85d244074e215531025070afa9d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Each criterion is a rule file, a repository state, and a verdict. The four commit types also
need a range, so each of them has a criterion for having one and a criterion for not.

**`section-implies-section` holds the implication and nothing else.** `when` non-empty and
`then` non-empty is green. `when` non-empty and `then` empty or absent is red, naming the
artifact, the rule, and which section is owed. `when` empty or absent is green whatever `then`
says, because an implication with a false antecedent is true and a gate that read it otherwise
would require every section of every template.

**It reads the rendered artifact.** A section is non-empty when the rendered `output.md` carries
text under its anchor, which is what a reader sees and what `template.Parse` returns. A section
the template does not have is a configuration error in the rule, red, naming the section and the
template.

**`commit-message` reads every subject in the range against a named pattern.** One
non-conforming subject is red, naming the commit and the pattern. All conforming is green. An
unknown pattern name is red, and says what is shipped, which is the same message
`xeno check commit-message` gives. `exempt: [merge-commits]` leaves a merge commit out; without
it a merge commit is judged like any other.

**`commit-trailer` requires a trailer on every commit in the range.** A commit without it is
red, naming the commit and the trailer. The comparison is on the trailer's key, case
insensitively, as git itself reads trailers, and the shipped case is `Xeno-Intent:`.

**`commit-signature` requires a valid signature on every commit.** An unsigned commit is red. A
commit whose signature does not verify is red. What counts as valid is one decision, recorded,
because git answers in five letters and not a boolean.

**`approver-not-author` compares the `by` of every decision against the authors of the range.**
A finding approved or overridden by somebody who authored a commit in the range is red, naming
the finding, the person and the commit. A decision by anybody else is green. No decisions is
green, because there is nothing to compare and a rule about approvals is not a rule requiring
one.

**A rule that needs a range and did not get one is red**, for all four commit types, naming the
rule and that the range is an input of the run. It is its own criterion because treating a
missing range as an empty one would make every commit rule silently green in exactly the
situation WP4's done-when names.

**A failure from git is a finding and not a panic.** A range that does not resolve, a repository
that is not a git repository, an unreadable object: each is red, naming what was asked and what
came back.

**Both patterns are shipped and tested.** `conventional-commits` as it is, and the second one
for a subject carrying an issue reference, with its own tests and a definition recorded, since
the hosts spell a squashed reference differently.

**The registry answers for all five.** G-Policy evaluates them where it reported them
unimplemented, and a type outside the five is still reported as having no implementation, which
is what keeps the registry a budget rather than an extension point.

**Nothing else moves.** `xeno check commit-message` behaves as before for the pattern it already
had. No set changes, so no `rules_hash` changes. `./xeno gate verify` stays at exit 0 over the
trail, and a repository with no rule tree stays green at every phase.

<!-- xeno:section:non-goals -->
## Non goals

**No rule file ships.** `given/builtin/` stays empty and `examples/rules/` is not created. The
first rule files anybody reads should be written against an independent reading of whether these
parameters make sense, and the person who wrote the types is not that reader.

**No sixth type and no extension point.** The registry's keys are the five names from section 9.
A project naming its own type still gets the finding #160 added, because project-defined
predicates are outside v1 and an external gate covers that ground.

**No expression in a rule file.** A rule names a pattern; it never carries a regex. This is the
line section 9 draws with four consequences behind it, and nothing here reads an expression out
of a project's file.

**No range inferred, ever.** Not from the default branch, not from the merge base, not from a
single commit. A missing range is a finding.

**No judgement of a signature's trust chain.** Whether a key belongs to who it claims is not a
question this gate can answer, and the host that authenticates pushes already answers most of
it. What is checked is what git reports about the signature on the commit.

**No author identity resolution.** `approver-not-author` compares strings, and section 9 says so:
a git author line is free text unless the commit is signed, so the rule enforces the discipline
of a team that means it and does not defeat somebody who does not.

**No change to how `xeno check commit-message` behaves** for the pattern it already serves, and
no second implementation of pattern matching anywhere.

**No new gate.** This fills a registry G-Policy already consults; the gate table does not change,
and the three rows still reporting `not-implemented` stay as they are.

<!-- xeno:section:constraints -->
## Constraints

**A gate is deterministic and network free**, and that now has to hold while running a
subprocess. `git log` reads the clone that is already there; the arguments are fixed by the
code and never taken from a rule file; and the package comment says what the prohibition
protects — no build, no test, no model — rather than forbidding every subprocess.

**One dependency.** `go.yaml.in/yaml/v3`, vendored. Nothing here needs a git library, and adding
one would be a decision rather than a step.

**The range is an input of the run and is not recorded.** The comment on `Base` and `Head` says
why: after a squash a recorded range would point at commits that no longer exist. So a verdict
depends on an input that the artifact does not carry, which is a property to state rather than
to design around.

**Two patterns, one implementation.** The map in `patterns.go` answers both at the gate and in
`xeno check commit-message`, which is the property section 9 asks for by name. The second
pattern goes in the same map and gets the same error message when it is misnamed.

**The second pattern's definition has to be written down.** Section 9 says what it is for and
not what it matches, and the two hosts spell a squashed reference differently, so the choice is
an assumption with the hosts named.

**What counts as a valid signature is a choice to record**, because git's answer is one of five
letters and the rule needs a boolean.

**A finding names a file, a cause and a next step**, including a finding about a commit, where
the file is the rule that asked.

**88 columns, SPDX, `gofmt`, `go vet`, the suite, and `./xeno gate verify` at exit 0.**

**One intent, one branch, one issue.** `162-the-predicate-types`, #162, labelled wp4.
