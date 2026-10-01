---
intent: github.com/triplem/xeno#158
phase: 05-review
created: "2026-10-01T15:16:17Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 03a7821821bd722f378c447f45c6c9a1d87f12d54a65fae42cab3c5a6a6805c3
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

There is still no rule set to render a checklist from — that is the piece this one makes
possible, and G-Policy stays `not-implemented` — so these entries carry no rule id and are the
questions this change raises, answered here.

**Does the resolver do what section 9 says, and only that?** Read back clause by clause against
the section: the two axes, the path deciding the level, `scope` checked against it, `binding`
allowed only under `given/` and rejected under `given/project/`, `learned/builtin/` rejected,
the two kind errors, `abstract` above the project level, the three precedence rules, and the
refusal to resolve a binding collision. Answered yes. Three additions beyond the section are
recorded in A67 with the line they are drawn on, and one candidate addition is refused.

**Is anything invented?** No field, no gate, no rule and no configuration key. The paths are
constants taken from section 9. `rules_hash` is the one place where something had to be decided
rather than read, and Appendix B delegates exactly that, which A62 established and A66 follows.

**Could the hash be wrong without anybody noticing?** Yes, after it is written, and that is in
the gaps rather than hidden: G-Schema does not recompute it, deliberately. What is checked is
that the writer and the definition agree, by a pipeline rather than by the implementation's own
rendering, and one real recorded value was reproduced by hand.

**Does this piece constrain the next ones more than it must?** It fixes two things for them: a
`check` arrives as a type and parameters as read, and a misdeclared rule is absent from the
effective set. The first was deliberately left untyped so the predicates can shape it; the
second is a decision the checklist inherits and would otherwise have to take itself. Answered
yes for the first, and the second is named in the impact so that the next piece knows it was
chosen rather than assumed.

**Does the trail stay honest?** `gate verify` is at exit 0 over 175 verdicts, no sealed artifact
moved, and the field's two meanings are recorded in A66 rather than reconciled by rewriting
anything. The division falling inside this intent is written into its own verification results.

**Is anything here a decision somebody else should take?** Two. Whether `applies_to` against the
known phases becomes an enumerated check, which is a specification change and belongs with the
shipped set. And whether the `by-hand` honesty rule is ever tightened, which A62 already
measured as rewriting sealed history and which now holds for two fields rather than one.

<!-- xeno:section:release-notes -->
## Release notes

**G-Rules is implemented.** A phase now reports a verdict on its rule tree from P0: rule
collisions resolved, `binding` respected, `scope` matching the path, which is the row section 7
has carried since the table was written. Four gates remain `not-implemented` where there were
five.

**`internal/rules` is new.** It reads the rule tree of section 9 across both axes, reports every
file it cannot use with a cause and a next step, resolves the tree to one rule per id by
specificity and origin with `binding` above both, and refuses to resolve a collision between two
binding rules on different levels. It evaluates no predicate: a check is read as a named type
with its parameters and handed on.

**`rules_hash` has a writer.** Every artifact written from now on carries a sha256 over the
effective rule set, defined to the byte in A66 and recomputable by anyone holding the tree.
Artifacts written before this keep `rules_hash: by-hand`, which stays a readable and accepted
value: the field sits inside `artifacts_hash`, so correcting them would rewrite sealed verdicts.

**A project with no rules is unaffected.** No rule ships, `given/builtin/` stays empty, and a
repository with no rule tree resolves to an empty set and passes. Adopting a rule means putting
a file in the tree; there is no switch and no `enabled` field.

**What this does not do yet.** No predicate is evaluated, so `kind: checked` is read and not
run. G-Policy and the P5 checklist are the next piece, the shipped set and the examples the one
after, and external gates are independent of all of them. #158 names them.

Closes #158. Refs #1.

<!-- xeno:section:residual-risk -->
## Residual risk

**The first reader of a finding from this gate will be somebody writing a rule for the first
time.** Every tree exercised here was built by the test suite or by me on a copy, and the
messages were written by the person who knew what they meant. That is the ordinary risk of a
gate that ships before its content, and the shipped set is where it gets tested: if those
findings read badly, the finding is about the wording here and not about the set.

**`rules_hash` is written and unchecked.** G-Schema does not recompute it, for the reason in
A66, so a wrong value is caught by nobody. The value is reconstructable by anyone holding the
tree of that commit, which is the claim actually being made; a reader who takes it for a
verified field is taking more than it offers, and section 16 is where that distinction belongs
if it turns out to matter.

**The honesty rule is one release behind the writer, for the second time.** `by-hand` stays
accepted in `rules_hash` as it does in `secrets_hash`, so a hand written artifact can assert
the placeholder where a writer exists. A62 measured what tightening costs; what nobody has
measured is whether a rule could distinguish an artifact written before the writer from one
written after, and that is the shape a fix would have to take.

**`kind: checked` is currently indistinguishable from `review` in effect.** A checked rule is
read, resolved, hashed and never evaluated, so a project adopting one today gets nothing from
it but a line in the hash. Until the predicates land, the two kinds differ only in what
G-Policy will eventually do with them, and a project that adopted rules in this window would
reasonably believe its checks were running.

**The weakest of the three additions in A67 is the `version` requirement.** Section 9 calls it
a counter and says nothing acts on it, so demanding at least one is a judgment about what a
format field means. If a reviewer disagrees, that is one `if` and one test, and the row is
written to be argued with.

**No fixture has all eight axis combinations at once.** Precedence is tested pairwise and by
level. `more` is two comparisons and the pairwise cases cover both branches, but the test that
would catch a resolution bug under a full tree does not exist, and the full tree is what the
first multi-level project will have.

**The fixtures sit where WP17 does not want them.** They build trees in `t.TempDir()` rather
than under `corpus/`, so the corpus that package specifies absorbs this work later rather than
receiving it now.
