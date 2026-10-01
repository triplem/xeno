---
intent: github.com/triplem/xeno#160
phase: 05-review
created: "2026-10-01T15:43:14Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c19858f8185222b73e3b73dfc2c06ef3214de04a3d4acb16bfeb3591ab32f538
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

The gate this piece adds would count entries here if a rule set existed; `given/builtin/` is
still empty, so these carry no rule id and are the questions the change raises. The first
checklist this gate counts will be in the shipped set piece, which is now also the first outside
reader of its findings.

**Does the gate do what sections 7, 9 and 12 require, and only that?** Read clause by clause:
the row and its `From` column, the three results and the note asymmetry, "every entry carries a
result, and nothing more", the omission that is now visible, and the lens clause that keeps a
lens out of the counted set. Answered yes. One finding beyond the enumeration — an entry naming
a rule outside the set — is recorded in A69 with its reasoning.

**Is the reading of `applies_to` for review rules defensible?** It is the decision most likely to
be overturned, so it is written out against its alternative in 02-design and recorded in A69
rather than left in the code. The alternative leaves a review rule about design answered
nowhere, which is why it was not taken.

**Does anything pass that should not?** Two things pass deliberately and both are in the gaps: a
checklist in a phase before P5, which is ignored rather than reported, and an answer that is
untrue, which no deterministic gate may judge. One thing passes that I would rather report, which
is the first of those, and it needs a spec change to become a finding.

**Does anything fail that should not?** A checked rule turns its phases red today, in every
project, until the predicates land. That is the deliberate consequence of refusing the silent
pass, it is stated in the release notes, and it is the strongest argument for the predicate piece
being next. No shipped rule is checked, so nothing in any repository is affected yet.

**Is the trail untouched?** No set changed, so no `rules_hash` changed; nothing sealed was
rewritten; `gate verify` is at exit 0 over 187 verdicts. The demonstration that did modify sealed
artifacts ran on a copy, which was then deleted.

**Is anything here a decision somebody else should take?** Three, all in the residual risk: making
a checklist outside P5 a finding, adding the unexamined-review case to section 16's list of what a
green gate does not mean, and whether an entry should record who wrote it, which is what would
make the lens limit checkable.

<!-- xeno:section:release-notes -->
## Release notes

**G-Policy is implemented.** A phase reports a verdict on its rule set from P0, and the five rows
a phase used to report as unrun are now three: G-Supply, G-Secret and G-Test.

**A review rule now has to be answered.** The P5 artifact carries a `review_checklist` list in
its frontmatter, one entry per review rule of the effective set, each with a result of `met`,
`deviation` or `not-applicable` and a note where it is not `met`. A rule with no entry is red, an
entry with no result is red, and a `deviation` with no note is red. That is the whole extent of
what a deterministic gate says about a review rule, and it is what turns one from a suggestion
into a recorded step.

**A lens cannot answer a rule.** An entry with no rule id is counted towards nothing, whatever
its `source` says, so a checklist of lens entries leaves a missing answer missing. The gate keys
on the absent rule rather than on the label, so a lens that omits its own label gains nothing.

**A checked rule is reported rather than quietly passed.** Until the predicate types ship, a
`checked` rule whose type has no implementation turns the phases it applies to red, naming the
rule and the type, with the next step of writing it as `review` or taking it out. A project
adopting a checked rule today gets a red gate instead of a green one that checked nothing.

**A project with no rules is unaffected**, at every phase, which is every repository today.

**What this does not do.** No predicate is evaluated. The `review-checklist` section is still
prose the agent writes beside the frontmatter, as all four structured lists are. No command
writes an entry. The shipped set and the examples are still to come, and external gates are
independent of all of it.

Closes #160. Refs #1.

<!-- xeno:section:residual-risk -->
## Residual risk

**A checklist written outside P5 is ignored in silence.** The check belongs at P5 and the silence
does not: a `review_checklist` in a design artifact is read by nobody, so somebody answering
rules in the wrong phase gets a green gate and an unanswered rule set. Making it a finding is a
specification change and therefore a person's commit, and it is the same proposal A67 already
carries for `applies_to`.

**Section 16 does not yet say what this gate's green does not mean.** A `met` on every entry is a
green verdict over a review nobody examined, which is exactly the kind of statement that list
exists for. The gate is honest in what it claims; the document that tells an auditor what the
claim is worth has not caught up. Also a spec change.

**The whole existing trail carries a `rules_hash` that records nothing.** Every artifact sealed
before this week says `by-hand`, and the copy showed a phase being judged against a checked rule
while carrying it. Nothing is wrong in the code and nothing can be fixed in those artifacts, so
what is missing is any marker in an artifact of which era it belongs to. A66 records the
division; nobody reading one artifact can tell which side of it they are on.

**The lens limit is unenforceable by construction.** A lens that writes a rule id is
indistinguishable from an agent answering a rule, because both are self-asserted strings in one
file. What would make it checkable is an entry recording who wrote it, which the artifact does
not carry and which is a schema change. It becomes real with WP11.

**The registry's evaluate branch is unexercised.** The map is empty, so the line that calls a
predicate has never run. The next piece covers it by existing, and until then this is the one
path in the gate with no test behind it.

**Both gates resolve the tree separately.** Section 7 asks for one resolution; this piece
guarantees one resolver and not one resolution, so a `gate run` walks the tree twice. On an empty
tree that is nothing, and on a large one it is twice the walk and twice the parse, with no cache
and no obvious place to put one while each gate is a pure function of its context. Cheap to fix
and not fixed, with no measurement behind the claim that it does not matter yet.

**No finding of this gate has been read by anybody but its author.** Six wordings ship, and the
one case where that was tested — the copy — was read by the person who wrote them. The shipped
set piece is where that stops being true.
