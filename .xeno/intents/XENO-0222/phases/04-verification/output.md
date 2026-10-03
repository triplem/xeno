---
intent: github.com/triplem/xeno#169
phase: 04-verification
created: "2026-10-03T09:10:43Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+93322c5.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e10f3232b12687468aa8d8a28cc588884989614c9b2bd653709f154b13eda70d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Most of this piece is prose, so the mapping has two kinds of row: what a test asserts, and what
was run against the installed client.

| Criterion | Checked by |
|---|---|
| Both manifests validate, with no warning left | `claude plugin validate` on each, run against the client |
| The marketplace points at this repository and loads from it | `marketplace add ./`, `plugin install xeno@xeno`, `plugin details xeno` |
| Seven skills with section 13's names | `TestTheSkillsAreTheSevenSectionThirteenNames`, and the client's own inventory |
| Every command a skill names exists | `TestEveryCommandASkillNamesExists`, against the dispatch table in `main.go` |
| Every section a skill asks for is in that phase's template | `TestEverySectionAPhaseSkillAsksForIsInItsTemplate`, through `template.Load` |
| What a skill says a gate refuses matches the gate | `TestEveryGateASkillNamesExistsAndIsImplemented` |
| Nothing in a skill only makes sense here | `TestNoSkillNamesThisProjectsOwnWorld`, seven words |
| A skill names the command path, not the MCP path | `TestNoSkillPromisesTheServerThatDoesNotExist` |
| A skill is selectable: name, description, when to use | `TestEverySkillCarriesANameAndADescription` |
| The hooks come from the plugin | `TestThePluginCarriesTheHookWiring`, and the client's inventory showing one Stop hook |
| `--vendor` carries the plugin's artifacts, idempotently | `TestVendorPutsThePluginsOwnArtifactsInTheRepository` |
| The divergence is recorded and raised | A77, and #169's own text |
| Nothing else moves | the suite, `gofmt`, `go vet`, `./xeno gate verify` |

The criterion with no check is the one that matters most and cannot be checked here: whether the
skills make an agent produce better artifacts than `CLAUDE.md` and two documents do. The next
intents are that measurement, and the artifacts they produce are the evidence.

<!-- xeno:section:results -->
## Results

**The suite is green.** 417 cases pass, nothing fails, `gofmt -l` outside `vendor/` lists
nothing, `go vet ./...` is silent, `./xeno gate verify` is at exit 0 over 210 verdicts.

**Both manifests validate.** `claude plugin validate .` passes the marketplace manifest and
`claude plugin validate .xeno/plugin` passes the plugin manifest, each with no warning — the
missing-description warning the scratch experiment produced is why the marketplace manifest
carries one.

**The plugin installs from this repository and the client reports what it found.** Added with
`claude plugin marketplace add ./`, installed with `claude plugin install xeno@xeno`, and
`claude plugin details xeno` reports:

    Skills (7)  xeno-design, xeno-implementation, xeno-intake, xeno-learning,
                xeno-requirements, xeno-review, xeno-verification
    Agents (0)
    Hooks (1)   Stop  (harness-only — no model context cost)
    MCP servers (0)

Seven skills, one hook, no agents and no server: the slice as scoped, confirmed by the thing that
loads it rather than by a test of mine. The plugin was then uninstalled and the marketplace
removed, and the machine is as it was — four marketplaces, no xeno plugin.

**The client priced the surface, which WP11's budget argument has never had a figure for.**

    Projected token cost
      Always-on:   ~606 tok   added to every session
      per skill:   ~80–100 always-on, ~620–830 on invoke

WP11 says the tool surface is "a standing cost on all work the project ever does" and that a
seventh MCP operation needs an argument. The skills now have the same kind of number: six hundred
tokens on every session of every project that installs the plugin, and seven hundred-odd each time
a phase skill fires. Six phases firing once each is roughly 4,500 tokens of instruction per
intent, against the ten thousand words of artifacts the same intent produces.

**Every mechanical claim in the skills holds.** Seven names against section 13. Every command
against the dispatch table. Every section against `template.Load` for that phase's template. Every
gate against the table, with the one unimplemented gate a skill names saying so in the same file.
No skill names Go, `gofmt`, `internal/`, this repository or a tool of it; none mentions MCP, a tool
call or an operation.

**`xeno init --vendor` carries the plugin.** Into an empty repository: the manifest, the hook
wiring and seven skill directories, beside the templates and the four rules. A second run changes
nothing, which is WP9's criterion and now has a third tree and two single files to not break.

**The hook is where the specification puts it.** One Stop hook in the plugin, calling `xeno cost
turn` on the path. This repository's `.claude/settings.json` still calls `./xeno`, and that
difference is deliberate: it runs the binary it just built.

<!-- xeno:section:gaps -->
## Gaps

**Nothing here shows that a skill helps.** Seven files of prose, checked for naming things that
exist and for not naming this repository. Whether an agent loading them produces better artifacts
than an agent reading `CLAUDE.md` and the two documents is the question the piece exists for and it
is unanswered. The evidence is the next intents, driven with the plugin installed, and it is
evidence of a kind this project has never collected: a comparison between two runs of the same
process with different context.

**The skills were written by the agent they instruct.** I wrote what I already do, which is why
they were quick and why they are suspect: an instruction that matches the author's habits reads as
obvious to the author and may read as nothing to anybody else. The shipped rule set had the same
problem and the same answer — the first outside reader finds out.

**Six hundred tokens on every session, measured and not yet weighed.** The client prices the
always-on cost at ~606 tokens and a skill invocation at ~620–830. Nothing in this repository says
what that buys, and the figure that would settle it is the one #117 has been missing: cost per
intent, which WP13 owes.

**No MCP server, so the five operations are still only commands.** WP11's done-when needs both
paths: the server, and the command path producing artifacts a gate cannot tell apart. The second
half is what twenty intents already did and what the skills now write down; the first half is a
piece of its own, and until it exists the surface the plan calls a budget has one kind of entry in
it rather than two.

**No lenses, so section 12's clause stays untested**: a disabled lens changes nothing but the
findings. Nothing can test that until a lens exists.

**One harness, so no leakage is visible.** Section 6 wants the second agent immediately after, to
expose harness leakage early. Everything in these skills that is secretly about this client — the
hook event names, the frontmatter shape, the way a description selects a skill — will only show
when a second client reads them.

**The hook is the only wiring and it is the easy one.** `xeno cost turn` on Stop costs nothing in
context and needs no decision. The hook WP7 actually wants, `G-Secret` on every write, needs the
gate, which is not implemented — so the plugin's hook wiring is one event where the specification
describes two kinds.

**A vendored project gets skills it cannot update.** `--vendor` copies them into
`.xeno/plugin/skills/`, and nothing updates a vendored tree afterwards except running `init` again
against a newer runner. That is the same property the templates and the rules already have, and it
is worth naming once: pinned means pinned.
