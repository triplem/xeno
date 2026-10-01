---
intent: github.com/triplem/xeno#158
phase: 01-requirements
created: "2026-10-01T15:00:53Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0c8276d54456994c605672f9d24129b76ba26ae05f013520ec7934483b064748
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Each criterion is a tree in the corpus and a verdict over it. Section 9 names every one of
them, so the list below is its sentences turned into fixtures rather than a design of my own.

**A repository with no rule tree is green.** No `.xeno/config/rules/`, no
`.xeno/plugin/rules/`, and G-Rules passes with an empty effective set. This is the state of
every repository today, including this one, so it is the first case rather than the edge: a
project that maintains no rules does not have a broken rule set.

**A well formed tree resolves in precedence order.** `project` over `org` over `provider` over
`builtin`, `given` over the matching `learned` at the same level, and the effective set
carries one rule per id with the level it won from. Resolution is a function that returns the
set, not a verdict, because G-Policy resolves the same tree and section 7 says it should
happen once.

**`scope` disagreeing with its path is red**, naming the file, the scope it claims and the
level its path gives it.

**`binding: true` outside `given/` is red**, and **`binding: true` under `given/project/` is
red**, which are two findings and not one: the first is a learned rule claiming what only a
given rule may claim, the second is a given rule claiming it where there is nothing left to
bind.

**Anything under `learned/builtin/` is red**, on the directory rather than on the file, because
section 9 says the directory does not exist.

**`kind: checked` without a `check` is red, and `kind: review` with one is red.** Both are
named configuration errors in section 9 and neither is a warning.

**A duplicate id at the same level is red.** Two files claiming one id on one level is a
collision nothing can resolve, as against the same id on two levels, which is the ordinary
case precedence exists for.

**Two colliding `binding` rules on different levels are red and stay red.** Red is not the
notable part; staying red is. The resolver must not pick the more specific one, because
section 9 says that is a matter between two organisations. The fixture has a binding rule at
`given/provider/` and another at `given/org/` with the same id, and the finding says both
files and refuses to choose.

**`rules_hash` is byte exact and reproducible.** Two runs over the same tree produce the same
hash; adding, removing or editing a rule that reaches the effective set changes it; a comment,
a file's mtime and the order the files are walked in do not. Its definition is written in
`ASSUMPTIONS.md` with the rendering spelled out, as A62 did for `secrets_hash`, and the test
checks the hash against that rendering piped through `sha256sum` rather than against itself.

**The runner writes the hash it computed.** A phase started after this piece carries a
`rules_hash` over the resolved set, and `by-hand` no longer appears in a new artifact. Where
there is no tree, the field says so in one fixed way rather than being absent, and which way
is a decision this intent records.

**G-Rules is no longer `not-implemented`**, applies from P0, and every finding it reports names
a file, a cause and a next step, which is the one requirement section 7 puts on every red
verdict.

**Nothing else moves.** `./xeno gate verify` stays green over the whole trail: no sealed
artifact's `artifacts_hash` changes, because the field they carry is the text they were written
with and this piece does not rewrite them.

<!-- xeno:section:non-goals -->
## Non goals

**No predicate is evaluated.** A `check` is read, its `type` is kept as a string and its
parameters as a map, and nothing looks at either. A rule whose `check` names a type that does
not exist resolves here without complaint, because the type registry belongs to the piece that
implements the types and inventing one now would fix its shape before anything uses it.

**No checklist and no G-Policy.** G-Policy stays `not-implemented`. The resolver returns an
effective set so that G-Policy can use it unchanged, which is the only accommodation this
piece makes for the next one.

**No rule ships.** `given/builtin/` stays empty and this repository keeps no rule tree of its
own. The shipped set is content for a later piece, and a set written now would be written
against a reader that had never read anything.

**No configuration field is added.** Section 9 fixes the paths, so there is nothing for
`project.yaml` to declare, and a path field would be an invented field under the second
standing rule.

**No `enabled`, no default-off, no switch.** A rule that is not in the tree does not apply,
which section 9 states twice. Whatever the shipped set turns out to be, it is adopted by being
there.

**Sealed artifacts are not rewritten.** Seventy-one intents carry `rules_hash: by-hand` and
keep it. The field is inside `artifacts_hash`, so rewriting one would change every verdict in
its intent, and the two meanings of the field are recorded rather than reconciled.

**No learned rule is harvested, promoted or abstracted.** Section 10 routes a learning through
a merge request against the rule set; this piece reads `learned/` as a level with its
precedence and its `abstract` requirement, and nothing writes into it.

**Rule versions are not compared.** `version` is a change counter that claims nothing, and
nothing in the runner acts on it. It is read, checked for being a number, and otherwise left
alone.

<!-- xeno:section:constraints -->
## Constraints

**Section 9 is the authority and it is not negotiable here.** Where this piece would prefer a
different shape, the specification wins until a person changes it. Two places where that bites:
the binding collision must not be resolved however tempting the more-specific reading is, and
`learned/builtin/` must be rejected rather than merely ignored.

**One dependency.** `go.yaml.in/yaml/v3`, already vendored. A rule file is YAML and nothing
else is needed; a second dependency is a decision and not a step.

**Appendix B governs the hash, and what it delegates has to be written down.** The rendering
goes into `ASSUMPTIONS.md` as an assumption with its reasoning, following A62, and the test
checks it against an external `sha256sum` so that the definition and the implementation are
compared rather than the implementation with itself.

**Deterministic and network free.** `internal/gates` says it in its package comment: gates
read, they never run anything and never ask a model. Resolution walks files and sorts; it does
not shell out, and the walk order must not reach the hash.

**The gate has to be cheap at the front of the table.** Section 7 puts G-Rules late precisely
because it resolves the whole tree, and says G-Rules and G-Policy should resolve it once. So
the resolver is a function with no side effects that both can call, and this piece does not
cache anything.

**A red verdict names a file, a cause and a next step.** That is the runner's own promise in
WP7's done-when, and it applies to every finding added here.

**88 columns, SPDX on every Go file, no per-file copyright line.** `gofmt`, `go vet` and the
whole suite stay clean, and `./xeno gate verify` stays at exit 0 over the trail.

**One intent, one branch, one issue.** `158-the-rule-tree-has-no-reader`, #158, labelled wp4,
with the remaining pieces of the package named in that issue rather than carried here.
