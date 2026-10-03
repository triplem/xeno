---
intent: github.com/triplem/xeno#169
phase: 02-design
created: "2026-10-03T09:02:42Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+93322c5.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 506e0d791a8d9de24c73df72315ae8858cbc343c3732c04ba4e2c0d7eec48f0b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The distribution root in this repository is `.xeno/plugin/`, so that is where the plugin's
artifacts go.** Section 13's tree lists `templates/`, `rules/given/builtin/` and `secrets.yaml` at
the distribution root, and in this repository those three are under `.xeno/plugin/`. The tree and
that directory are the same set of names, so `skills/` and the plugin manifest join them there
rather than at the repository root. `--plugin-from` already points `xeno init` at that directory,
which is the existing evidence for the reading.

**The plugin manifest sits at `.xeno/plugin/.claude-plugin/plugin.json`, which section 13 draws
elsewhere.** The client loads a plugin from `<source>/.claude-plugin/plugin.json` and validates it
only there. Section 13 draws `plugin.json` at the plugin root. The document describes a third
party's format and the third party disagrees with it, so following the document literally produces
a tree that does not load — which is the one case where "the specification wins until a person
changes it" cannot be followed as written. The layout that loads is used and A77 records it.

**The marketplace wrapper sits at the repository root, outside the distribution tree.** A client
adding this repository as a marketplace reads `<repository>/.claude-plugin/marketplace.json`, and
the manifest names `./.xeno/plugin` as the plugin's source. Section 13 draws the wrapper inside the
tree, which would work only if the tree were the repository root. Recorded with the same row,
because it is the same disagreement seen from the other end.

**Seven skills, named as section 13 names them, each one the same five things.** What the phase is
for, in two sentences. The sections it owes, taken from that phase's template. The commands that
carry it, in order. What its gate refuses, named gate by gate. What it hands to the next phase.
That shape is the same in all seven so an agent that has read one knows where to look in the
others, and it is the shape that made the mechanical checks possible.

**A skill names commands and never operations.** There is no MCP server, so every step is a
command from the dispatch table. When the server lands, a skill gains the operation beside the
command rather than instead of it, because WP11's done-when requires the command path to keep
working.

**The hook comes from the plugin and calls `xeno` on the path.** `hooks/hooks.json` in the plugin
wires the Stop event to `xeno cost turn`, which is how a project that installs the plugin records a
turn's cost without writing a settings file. It says `xeno` rather than `./xeno`, because section
13's channel table delivers the runner to a developer machine as a platform binary; this
repository keeps its own `.claude/settings.json` entry, which runs the binary it just built.

**`--vendor` copies what exists of section 13's list and invents nothing.** The manifest, the
skills and the hook wiring, beside the templates, the rules and the filter. No `mcp.json`, because
a declaration of a server that does not exist is a client failing at startup rather than a
placeholder.

**The skills say what a phase owes, not how to think.** Each one points at the two documents for
the process and at the templates for the sections, and none of them restates a rule, a gate's
internals or a convention of this repository. A skill is loaded into a context that already has
the project's own instructions, and a skill that argued with those would be the harness leakage
section 6 wants exposed rather than shipped.

<!-- xeno:section:alternatives -->
## Alternatives

**Follow section 13's tree literally and ship a plugin that does not load.** The first standing
rule says the specification wins until a person changes it, and the literal reading here produces
a manifest the client ignores. Rejected, with the divergence recorded and raised rather than
hidden: a rule that protects the documents from an agent's convenience is not a rule that the
documents describe other people's file formats correctly.

**Move `templates/`, `rules/` and `secrets.yaml` to the repository root** so that the distribution
tree and the repository root coincide, which would put both manifests where section 13 draws them.
Rejected for its size and its blast radius: `template.Load`, `rules.PluginDir`, `secrets.Shipped`,
the scaffold and `init` all name `.xeno/plugin`, every vendored project would need the same move,
and the gain is cosmetic agreement with a tree that is wrong about the manifests anyway.

**Put the marketplace wrapper inside `.xeno/plugin/.claude-plugin/`** as section 13 draws it.
Rejected because a client adding the repository would not find it, and asking an adopter to add
`<repository>/.xeno/plugin` as a marketplace source would make the wrapper's location part of
every project's setup instructions.

**Fold `xeno-learning` into the six phase skills.** Section 10 has learning happen at the end of
every phase and once more at close, so the instruction would be repeated six times. Rejected
because section 13 names seven skills and because a repeated instruction is the thing that drifts:
six copies of it are six places for one correction to miss.

**Write the skills as slash commands instead.** The client supports them and they would be
invoked explicitly rather than selected by description. Rejected: section 13 enumerates skills,
and a surface the specification does not name is an invented one — the same line the predicate
registry draws.

**Carry the hook in the scaffold's `project.yaml` or in `xeno init`'s output.** It would let the
runner own the wiring rather than the plugin. Rejected because section 13 lists hook wiring as
part of what the plugin carries and because `xeno init` deliberately does not touch a developer's
client configuration, which is the same reason it does not touch their git configuration.

**Write the skills against the MCP operations and ship `mcp.json` as a stub.** It would make the
later piece a smaller change. Rejected: a client reading a declaration of a server that is not
there fails at startup, and a skill describing an operation that does not exist is worse than a
skill describing a command that does.

<!-- xeno:section:impact -->
## Impact

**A project can install the process rather than be told about it.** After this, adopting Xeno is
the runner, `xeno init`, and a plugin whose seven skills say what each phase owes and which command
carries it. The twenty intents in this repository were driven by an agent that had read the
specification; an adopting project gets the same instructions in the form a client loads.

**M0's remaining blocker goes.** Section 6 puts WP11 before M0, and the part of WP11 that section
6's M0 needs — one agent, end to end — is the skills and the manifest. Whether that is enough
without the MCP server is a reading of the plan, which this piece deliberately leaves open and
makes answerable.

**The hook moves from this project's settings to the plugin.** A project that installs the plugin
records a turn's cost; today only this repository does, through a file it wrote by hand. That also
makes WP13's figures collectible outside this repository for the first time, which is the thing
#117 has been missing.

**Section 13 acquires a known error.** Two of its drawn paths do not match the client it describes.
The code follows the client, the register says so, and the document is a person's to correct — the
first time in this project that the specification is wrong about an external fact rather than about
itself.

**`--vendor` becomes three trees and a file.** Templates, rules, skills, and the manifest plus the
hook wiring. Every addition to section 13's list from here is one more loop in one function, which
is the right shape, and the function is now the single place that knows what a vendored project
receives.

**The skills are the first thing in this repository written to be loaded by a model rather than
read by a person.** The templates are rendered, the rules are answered, the documents are read; a
skill is context. Nothing here can measure whether it helps, and the next intents are the
measurement: they will be driven with the skills loaded, and the artifacts they produce are the
evidence.
