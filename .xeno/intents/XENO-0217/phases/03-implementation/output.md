---
intent: github.com/triplem/xeno#158
phase: 03-implementation
created: "2026-10-01T15:11:56Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9f896d467eee3262f207fb2941c527c83f78911370636f2e0bf86506edec6bbf
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

Three files new, four changed: 412 lines of `internal/rules/rules.go`, 320 of its test, 101 of
`internal/gates/rules_test.go`, and 72 added across the gate table, the runner, the runner's
test and the register.

**`internal/rules`, the package.** `PluginDir` and `ConfigDir` as constants, `Levels` from
least to most specific with a `rank` over them, and `treeDir` mapping an origin and a level to
the one directory section 9 puts it in. `Rule` carries the file's fields plus `Origin`, `Level`
and `Path`, which the path supplies and the file does not, so that `Scope` stays the file's own
claim and remains comparable. `Check` keeps a `Type` and its `Params` as read, split by an
`UnmarshalYAML` because section 9 writes a check as one mapping whose `type` names the
predicate and whose other keys parameterise it.

**`Load` walks both axes and returns rules and problems.** An absent directory is a level
nobody maintains and is skipped; a present `learned/builtin/` is a problem on the directory,
because its existence is the error. A file that cannot be read, cannot be parsed, or fails any
check below is reported and left out of the set.

**What `read` checks, in the order the findings are written.** The format's own completeness
first — `id`, `version` as a counter of at least one, `statement`, a non-empty `applies_to` —
then `scope` against the level its path gives it, then `kind` with `checked` requiring a
`check` and `review` forbidding one and anything else named as itself, then `binding` outside
`given/` and `binding` under `given/project/` as two findings rather than one, then `abstract`
for a learned rule at the provider or org level. `duplicates` then reports one id claimed twice
at the same level by the same origin, which is the collision precedence cannot resolve.

**`Effective` resolves and refuses.** Rules are grouped by id, ids walked in sorted order for
a stable result. Where one candidate is `binding`, it wins regardless of specificity. Where
more than one is, the id is left out of the set and the problem names every file that declared
it, sorted by level so the message reads from the least specific up. Otherwise `more` orders
the candidates by level and, at the same level, `given` over `learned`, and the first wins.

**`Render` and `Hash`.** One line per effective rule, sorted by id, tab separated, nine fields,
newline terminated, with the statement's whitespace normalised by `oneLine` and the check
flattened by `renderCheck` to `type=` plus its parameters as sorted dotted keys. `flatten`
walks maps and sequences itself rather than marshalling them, so the ordering belongs to this
package. `Hash` is `hashing.Hex` over that rendering. A66 records the definition.

**`internal/gates`, G-Rules.** `rulesGate` replaces `notImplemented` in the table, reads the
tree, resolves it, and words every problem from both steps into a finding with its file, cause
and next step. It judges nothing of its own, so G-Policy resolving the same tree later cannot
disagree with it about what the set is.

**`internal/runner`, the writer.** `SectionSet` now computes `rules_hash` from the resolved set
and writes it, where the harness wrote `by-hand` before. A tree with a problem still produces a
hash over the rules that were usable: G-Rules reports the problem and the phase is red, and the
field says which set actually resolved rather than going absent and saying nothing.

**Tests.** `internal/rules`: an absent tree, the empty rendering, precedence by specificity and
by origin, binding beating specificity, the binding collision staying unresolved with both
files named, a table of the ten configuration errors each asserting its cause and that it
carries a next step, `learned/builtin/` on the directory, the duplicate id, four hash
invariances and one change, only the effective set reaching the hash, the hash against a
`sha256sum` pipeline, the rendering's shape, and a check's key order not reaching the hash.
`internal/gates`: an empty tree passing, a well formed tree passing, a misplaced scope red with
file, cause and next step, a binding collision reaching the verdict, and the table no longer
reporting `not-implemented`. `internal/runner`: a new artifact carrying the hash of the
effective set rather than the placeholder, and a rule entering the tree changing it.

**`ASSUMPTIONS.md`, A66 and A67.** A66 is the hash definition with its reasoning, the empty
case, and why neither the placeholder rule nor G-Schema's recomputation is tightened with it.
A67 is the three findings this piece adds beyond the enumerated ones and the one it
deliberately does not add.

<!-- xeno:section:deviations -->
## Deviations from the design

**The register had already settled what the design argued out.** 02-design reasons from first
principles that the placeholder rule must not be tightened, because `rules_hash` sits inside
`artifacts_hash` and seventy-one sealed intents carry `by-hand`. A62 records exactly that for
`secrets_hash`, and it is not an argument there but a measurement: removing the field from
`WriterlessHash` was tried and diverged all 87 verdicts this repository carried at the time.
The design should have cited it and did not. A66 now does, and the learning says what to read
first next time.

**G-Schema not recomputing the field is a decision taken in implementation.** The design says
the runner writes the hash and says nothing about whether a gate checks it. Writing the gate
made the question unavoidable: `context_hash` and `strings_hash` are recomputed because they
name something in the tree now, while `rules_hash` names the rule set as it was when the phase
was written, so recomputing it would turn every earlier artifact red the first time a rule
changed. It is recorded in A66 rather than in the design, which is the wrong file for a
decision of that weight, and it is flagged here for that reason.

**`version` is required to be at least one, which section 9 does not say.** The section calls
it a change counter and says nothing acts on the number. A file with no `version` is reported
as a missing field, which puts a judgment inside A67's reasoning rather than inside the
enumeration: an absent counter is a format field missing, not a rule about rules. It is the
weakest of the three additions and the one most likely to come back as a finding against this
intent.

**The harness's own patcher broke on this intent's prose.** The script that writes the fields
a harness supplies looked for `rules_hash:` anywhere in the file, and 00-intake's problem
section quotes `rules_hash: by-hand` in prose, so the field was never written to the
frontmatter and the phase came out red on a missing field. An artifact about `rules_hash` broke
the tool that writes `rules_hash`. The script now reads the frontmatter block alone. Recorded
because it is the second time in two intents that a tool read prose as structure, and because
the gate caught it, which is the part that worked.

**The corpus is still `testdata`-shaped.** WP17 puts the per-gate corpus under `corpus/` at the
repository root, and these tests build their trees in `t.TempDir()` instead. That is the
convention every other package here follows today, so this piece matches the tree rather than
the plan, and the fixtures will have to move when WP17 lands.
