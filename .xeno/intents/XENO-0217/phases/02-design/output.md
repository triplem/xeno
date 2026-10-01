---
intent: github.com/triplem/xeno#158
phase: 02-design
created: "2026-10-01T15:03:36Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1834952b8c2b2dc20d54a586d7c7d7c5acc83231bf75dedb8f5321ef482dd5c5
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

**`internal/rules` returns a set and problems; the gate words them.** `Load` walks the tree
and returns the rules it read and a list of problems, each with a path, a cause and a next
step. `Effective` takes the rules and returns the resolved set and the problems resolution
itself produces. G-Rules turns problems into findings and decides the verdict. That is A61's
shape, where `Attach` returns what it declined and each caller words it: three callers are
coming — G-Rules, G-Policy and the hash writer — and a package that printed or failed would
decide for all three.

**The effective set holds only well formed rules.** A rule with a problem of its own is
reported and left out of resolution. Resolving a rule whose `scope` disagrees with its path
means choosing which of the two to believe, and a rule that is red for being misdeclared
should not also win a precedence contest. The verdict is red either way, so the only question
is whether the set handed to the next piece is coherent, and it is.

**A binding collision removes the id from the set.** Section 9 says the gate is red and stays
red, and that the matter is between two organisations. So resolution does not pick the more
specific rule, and it does not pick either: the id resolves to nothing, the problem names both
files, and the set is reported as unresolved rather than silently narrowed.

**`rules_hash` is a hash over a canonical rendering of the effective set, one line per rule.**
Following A62 for `secrets_hash`: tab separated fields, sorted by id in byte order, newline
terminated, and the rendering is exported so the test can pipe it through `sha256sum` rather
than compare the implementation with itself. Per line: id, version, origin, level, kind,
binding, `applies_to` sorted and comma joined, the statement with its whitespace normalised,
and the check flattened to `type` plus its parameters as sorted dotted keys. Semantic fields
only, so a comment, a file's name and the order the walk happened to take do not reach the
hash, and an edit to a statement does.

**An empty effective set hashes as the rendering of nothing.** `rules_hash` is a required
field, so the option the digest took for `secrets_hash` — leave it out rather than assert a
hash over nothing — is not available here, and G-Schema accepts only a sha256 or the
placeholder. So an empty tree produces the hash of the empty rendering, which is a statement
about a resolved set that turned out to be empty rather than about a tree nobody looked at.
That distinction is recorded rather than left in the value.

**The placeholder stays honest and the runner stops writing it.** `rules_hash` remains in
`WriterlessHash`, so the seventy-one intents that carry `by-hand` stay green, and the runner
computes the field for every artifact written from now on. Tightening the gate to reject the
placeholder would turn every sealed artifact in this repository red on `gate verify`, which
trades the trail for a check on a field only the runner writes now.

**What is validated is what the documents name, plus the format's own completeness.** Section
7 gives G-Rules collisions, `binding` and `scope` against the path; section 9 adds the two
kind-versus-check errors, the rejection of `learned/builtin/`, `binding` under
`given/project/`, and `abstract` for a learned rule above the project level. To those this
piece adds only a parse failure, a missing required field of the format, and a duplicate id at
one level, each of which is a file that cannot be read as a rule at all rather than a new rule
about rules. Nothing else is checked, and `applies_to` naming an unknown phase is deliberately
not a finding, because the second standing rule makes an unenumerated check an invented rule.

**The paths are fixed, not configured.** `.xeno/plugin/rules/given/builtin/` and
`.xeno/config/rules/{given,learned}/{provider,org,project}/`, exactly as section 9 writes
them, read from constants. A path that configuration could move would make `scope` against the
path uncheckable, which is the one thing the gate is for.

<!-- xeno:section:alternatives -->
## Alternatives

**Hash the rule files rather than the effective set.** One sha256 per file, sorted by path,
which is simpler and needs no rendering. Rejected on two counts. It would make a comment or a
reformatting change `rules_hash`, so an artifact would record a different rule set from the
one that was in force. And it would hash files that lost their precedence contest, which are
not the set that applied; `rules_hash` is what proves which rule set was in force, and the
set that was in force is the effective one.

**Resolve the binding collision by specificity.** The natural implementation: binding wins
over non-binding, and among binding rules the more specific one wins. Rejected because section
9 forbids it in a sentence written for this case — a matter between two organisations and not
something a gate may resolve. This is the alternative worth recording precisely because it is
what the code would do if nobody had read that sentence.

**Keep misdeclared rules in the effective set.** Report the problem and resolve anyway, so the
set is complete and the verdict red. Rejected because a rule whose `scope` contradicts its path
has two reaches and resolution would have to choose one, and because the next piece renders a
checklist from this set: an entry for a rule nobody can place would be worse than a missing
entry.

**Return a verdict from the package.** `rules.Check(root) []model.Finding`, with the gate a
thin call. Rejected: the hash writer needs the set and not findings, G-Policy needs the set
too, and section 7 says the tree should be resolved once for both gates. A package that
returned a verdict would be re-walked by whoever needed the set.

**Type the check parameters now.** Give `section-implies-section` and the four commit types
their Go shapes in this piece, so that the next one only writes evaluation. Rejected as fixing
a shape before anything uses it: the types' own piece will learn something from the first
predicate it evaluates, and a shape written a week early is a shape that gets changed by the
first test rather than by a reason. The check keeps its `type` and its parameters as read.

**Accept the placeholder nowhere, and migrate the trail.** Rewrite the seventy-one sealed
artifacts to carry a computed hash so the field has one meaning. Rejected outright: the field
sits inside `artifacts_hash`, so every verdict in every intent would change and the trail
would be rewritten to make a check tidier. Two meanings divided by a commit is the honest
record.

**Make `applies_to` against the known phases a finding.** It is cheap and it would catch a
typo that silently disables a rule. Rejected under the second standing rule, and recorded in
the residual risk instead: an unenumerated check is a spec change first, and this is a good
candidate for one.

<!-- xeno:section:impact -->
## Impact

**Two gates stop being promises, one of them in this piece.** G-Rules moves from
`not-implemented` to a verdict, which changes what a green phase means in every project: nine
gates were running here, ten after this. G-Policy stays `not-implemented` and now has a
resolver to call.

**`rules_hash` changes meaning for every artifact written after this commit.** Before, the
field was a person's word that a rule set was in force. After, it is a hash over a resolved
set, recomputable by anyone holding the tree. The trail will carry both, divided by this
commit, and `ASSUMPTIONS.md` is where that division is written down.

**Nothing sealed moves.** No `artifacts_hash` changes, because the field is read as the text
each artifact was written with. `./xeno gate verify` has to stay at exit 0 over the whole
trail, and that is an acceptance criterion rather than an expectation.

**This repository gains a gate it does not use yet.** Xeno keeps no rule tree, so G-Rules will
pass over an empty set here until the shipped set lands. The first real exercise of the
resolver is therefore the corpus, which makes the fixtures the test rather than the
demonstration — and makes the shipped set the piece where this gets its first use against
something a person wrote.

**The next two pieces get their foundation and their constraint.** The predicates receive a
`check` with its type and parameters as read, and the checklist receives an effective set
with one rule per id. Both inherit the decision that a misdeclared rule is absent from the
set, which is what keeps either from having to decide what a half-placed rule means.

**The cost is a new package and a new concept in the tree**, four levels and two origins, which
is more surface than the secret filter's two layers. Section 9 fixes all of it, so the surface
is the specification's rather than this intent's, but a reader of `internal/` now meets the
rule model before anything evaluates a rule, and the package comment has to say so.
