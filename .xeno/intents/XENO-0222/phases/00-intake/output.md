---
intent: github.com/triplem/xeno#169
phase: 00-intake
created: "2026-10-03T09:00:04Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+93322c5.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 16d0dc70171243e574699e02917b14c99561eb51c01f7dd9f69af1861789538a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

M0 has two readings and one of them is met. The milestone table's M0 is the gate job, "from
which point Xeno checks its own repository": `verify` is a required check on the default branch
with administrators included, `gate verify` is green over 207 verdicts, and twenty-plus intents
have run all six phases with artifacts and recorded verdicts. The sequence in section 6 puts WP11
at step 5 and M0 at step 6, and WP11 is unbuilt.

**Nothing of the agent layer exists.** No `plugin.json`, no `skills/`, no `mcp.json`, no hook
wiring, no lenses. What drives a phase today is an agent reading `CLAUDE.md` and the two documents
and calling the runner's commands, which is exactly the fallback WP11's own done-when requires —
"the same phase can be carried out with commands alone, with no MCP server and no hooks, producing
artifacts a gate cannot tell apart from the ones an agent produced" — and it is the fallback
rather than the package.

**What that costs is not visible in this repository and would be fatal in another.** Twenty intents
of phase artifacts exist because the agent driving them had the specification in context and a
session's worth of accumulated reading. A project adopting Xeno gets the runner, the templates, the
rule set and no description anywhere of what a phase is for, which command carries which step, or
what its gate will refuse. The skills are that description, and section 13 fixes their names and
their count.

**The one thing the plan said to verify rather than assume is the one thing that was still open.**
Section 9's list carries "whether the client accepts the instance URL for the marketplace -->
unsure right now". It is now answered on this machine: the client's own help says
`marketplace add` takes a URL, a path or a GitHub repository; adding a directory as a marketplace
succeeded; and `claude plugin validate` passes both manifests. The marketplace wrapper works, which
is what the plan wanted known before anything was built on it.

**Answering it produced a disagreement with the specification.** Section 13 draws `plugin.json` at
the plugin root with `.claude-plugin/marketplace.json` beside it. The client loads a plugin from
`<source>/.claude-plugin/plugin.json` and a marketplace from `.claude-plugin/marketplace.json`, and
validates both only in that layout. The specification describes a third party's format and the
third party disagrees, which is the one case where the first standing rule's "the specification
wins" cannot be followed literally: a tree that followed section 13 would not load at all.

**And the hooks are in the wrong place today.** `.claude/settings.json` in this repository wires
`xeno cost turn` to the Stop event by hand. That is this project's own configuration, so a project
adopting Xeno records no cost at all unless somebody tells it to write the same file. Section 13
lists hook wiring as part of what the plugin carries.

<!-- xeno:section:scope -->
## Scope

**In scope.** `plugin.json` and the marketplace wrapper, in the layout the client validates. The
six phase skills and `xeno-learning`, with the names section 13 fixes, each saying what its phase
is for, which commands carry it, and what its gate refuses. Hook wiring carried by the plugin
rather than by a project's own settings. `xeno init --vendor` carrying the plugin's artifacts
beside the templates and the rules. The divergence from section 13 recorded as an assumption and
raised as a finding for the specification.

**Out of scope, and each for its own reason.**

The MCP server and `mcp.json`. Five operations plus the index query need a server, and WP11's
done-when says the same phase has to be carryable with commands alone — so the skills describe the
command path first, and the server optimises a path that already works. A `mcp.json` declaring a
server that does not exist would be a declaration a client fails on.

The four lenses. Section 12 fixes what they are, and they are enablement rather than the walking
skeleton. The clause nothing can test until they exist — that a disabled lens changes nothing but
the findings — stays untested either way.

The second agent. Section 6 wants it immediately after, to expose harness leakage early. There is
nothing to leak until one harness has something.

Closing M0. This piece removes the blocker that section 6's reading names; whether that reading is
satisfied by a plugin without an MCP server is a judgement about the plan, and the plan is not mine
to read on your behalf.

**One boundary worth naming.** A skill is prose that an agent loads, and nothing in this repository
can check that it produces the artifacts it describes. What can be checked is that the manifests
validate, that the commands the skills name exist, and that what they claim about a gate matches
what the gate does. The rest is read by a person, which is the same limit the shipped rule set has.

<!-- xeno:section:context-rationale -->
## Why this context

Section 13 is the specification for this piece: the three artifacts, the distribution tree with
every skill named, what `--vendor` copies and what stays behind, and the table of channels and
recipients. Read entire, because the tree is the part that turned out to disagree with the client
and the `--vendor` list is the part nothing implements yet.

WP11 of the plan is read for what the package is and for its done-when, and the sentence that
decided this piece's scope is the last one: a phase has to be carryable with commands alone. That
is what makes the skills the content and the MCP server an optimisation rather than the other way
round. Its paragraph on the marketplace is read for what was to be verified.

Section 12 is read for the lens model, to establish that lenses are out of this slice rather than
forgotten, and for the environment the harness exports — which is already built, and is the one
part of WP11 that landed early.

Section 9's open list in the plan is read for the verification point, and the client itself is
read for the answer: `claude plugin marketplace --help`, `claude plugin validate` against a
candidate layout, and `claude plugin marketplace add` against a directory. The experiment ran in a
scratch directory and the marketplace was removed again.

`cmd/xeno/main.go` is read for the command surface the skills have to name, command by command,
because a skill that names a command that does not exist is worse than no skill. `internal/gates`
is read for what each gate refuses, since that is what a skill has to tell an agent to expect.

The six shipped templates are read for the sections each phase renders, because a skill's job is
to say which sections a phase owes and the template is where that is fixed.

`internal/runner/init.go` is read for `vendorPlugin` as #165 left it, which is the one function
this piece extends.

`.claude/settings.json` is read as the thing to replace: it is this project's own hook wiring, and
a project adopting Xeno has no equivalent.

`.xeno/plugin/` is read for what it already holds — templates, rules, the secret filter — because
section 13's distribution tree and this directory are the same set of names, which is what decides
where the new artifacts go.
