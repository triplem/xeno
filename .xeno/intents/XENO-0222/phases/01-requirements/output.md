---
intent: github.com/triplem/xeno#169
phase: 01-requirements
created: "2026-10-03T09:01:28Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+93322c5.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 65e8fec5334d47a343f367ee8db23932901eaf4524a8662ad752704144e97ece
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**Both manifests validate.** `claude plugin validate` passes the marketplace manifest and the
plugin manifest, with no warning left unaddressed — the description warning the experiment
produced is one of them.

**The marketplace points at this repository and loads from it.** Adding the repository as a
marketplace and reading the plugin back shows the skills the plugin declares. Checked against the
client rather than asserted, and the state it leaves on the machine is cleaned up afterwards.

**Seven skills exist, with the names section 13 fixes.** `xeno-intake`, `xeno-requirements`,
`xeno-design`, `xeno-implementation`, `xeno-verification`, `xeno-review`, `xeno-learning`. Not six
with learning folded in, and not eight.

**Every command a skill names exists in the runner.** Checked mechanically against the dispatch
table in `cmd/xeno/main.go`, because a skill that names a command the binary does not have is
worse than no skill at all.

**Every section a skill tells a phase to write is a section of that phase's template.** Checked
against the shipped template set, for the same reason.

**What a skill says a gate refuses matches what the gate does.** The gate ids a skill names exist
in the table, and a skill does not promise a check that reports `not-implemented`.

**Nothing in a skill only makes sense inside this repository.** No Go, no `gofmt`, no reference to
this project's own layout or conventions: the skills travel to every project that installs the
plugin, which is the same bar `given/builtin/` has.

**A skill tells an agent the command path, not the MCP path.** There is no server, so a skill that
said "call the read-intent operation" would describe something absent. Each one names the command
that does the step, which is also what WP11's done-when requires of the fallback.

**The hooks come from the plugin.** The wiring that records a turn's cost is carried by the plugin
rather than by a project's own settings file, so a project that installs the plugin records cost
without writing anything. This repository's `.claude/settings.json` stops being the only place
that knows.

**`xeno init --vendor` carries the plugin's own artifacts.** Section 13 says `--vendor` copies
`plugin.json`, `skills/` and `mcp.json` beside the templates, the rules and the secret filter. What
exists of that list is copied, a second run changes nothing, and what does not exist yet is not
invented.

**The divergence is recorded and raised.** An assumption says what the client requires and why the
tree follows it, and a finding for the specification is written where a person will act on it —
not absorbed into a comment.

**Nothing else moves.** No gate changes, no rule set changes, so no `rules_hash` changes. No sealed
artifact is rewritten. `./xeno gate verify` stays at exit 0 and the suite stays green.

<!-- xeno:section:non-goals -->
## Non goals

**No MCP server and no `mcp.json`.** A declaration without a server is a client failing on
startup, and the five operations plus the index query are their own piece.

**No lenses.** Four skills section 12 fixes, out of this slice by scope rather than by oversight.

**No second harness.** The plan wants it immediately after; immediately after is not now.

**No change to what the runner does.** This piece adds content and one vendoring loop. If a skill
needs the runner to behave differently, that is a finding about the runner and not a licence to
change it here.

**No slash commands, no subagents.** The client supports both; section 13 names skills, and a
surface the specification does not enumerate is an invented one.

**No claim that the skills work.** They are prose an agent loads. What is checked is that they
name things that exist; whether they produce better artifacts than `CLAUDE.md` and two documents
is measurable only by running intents with them, which is the next piece's evidence and not this
one's.

**No edit to `docs/`.** The divergence with section 13 is raised, not fixed.

**No closing of M0.** The blocker is removed; the reading is yours.

<!-- xeno:section:constraints -->
## Constraints

**The client decides the layout, because the specification describes its format and is wrong about
it.** Section 13's tree would not load. The layout that validates is used, the divergence is
recorded as an assumption, and the document stays for a person to change — which is the only
reading of the first standing rule that leaves a working plugin.

**A skill is read by an agent in another project, with none of this session's context.** No Go, no
`gofmt`, nothing about `internal/`, no reference to a rule this repository happens to have adopted.
The bar is the one `given/builtin/` has: a project could not reasonably refuse it.

**Every name a skill uses is checked against the tree.** Commands against the dispatch table,
sections against the templates, gates against the table. A skill is prose and these three are the
parts of it that can be wrong mechanically.

**Seven skills, with section 13's names.** Not a different split, not a shorter set: the names
appear in the specification and a client resolves them by name.

**The plugin version is the shared one.** Section 13 says three artifacts and one shared version
number, and the artifacts already record `plugin_version: 0.1.0-dev`.

**`xeno init` stays idempotent**, which is WP9's criterion and now has a third tree to not break.

**No new dependency, 88 columns in prose, and the suite, `gofmt`, `go vet` and
`./xeno gate verify` all green.** The skills are Markdown with frontmatter and carry no SPDX line,
as the templates do not.

**One intent, one branch, one issue.** `169-the-agent-layer`, #169, labelled wp11, off a `main`
that carries all of WP4.
