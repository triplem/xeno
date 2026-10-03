---
intent: github.com/triplem/xeno#169
phase: 05-review
created: "2026-10-03T09:12:01Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+93322c5.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1e1f1961edcb9f201b63b965185442abdded9b0d79b03f954bde04649f14a489
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
  - rule: deviations-are-traceable
    result: met
    note: >-
      The six deviations each name what they depart from: the criterion the verification skill
      broke and the test shaped around it, the column convention and the byte-counting that hid
      it, the client's refusal of a bare dot, a listing that disagreed with itself, section 13's
      vendor list, and the shape of a package that exists for its tests.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      `--vendor` writes three more paths into a project, which is additive: a project that
      vendored before this release gets them on its next init. The manifests are a contract with
      the client, and A77 plus the release notes say which layout to use and why section 13's
      drawn one does not load.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      No dependency was added. The plugin is Markdown and JSON, and the test package uses the
      standard library and this project's own packages.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter, under `review_checklist`. What
follows is the reasoning and the questions this change raises beyond them.

**`deviations-are-traceable` — met.** The six deviations each name what they depart from: the
criterion the verification skill broke and the test that had been shaped around it, the 88 column
convention and the byte-counting measurement that hid it, the client's own refusal of a bare dot,
the listing that disagreed with itself after a removal, section 13's `--vendor` list and the one
entry still missing from it, and the shape of a package that exists for its tests.

**`interface-change-needs-a-migration-note` — met.** Two things outside this intent depend on what
changed. `--vendor` now writes three more paths into a project, which is additive and needs no
migration: a project that vendored before this release gets them on its next `init`. The manifests
are a contract with the client, and the note is A77 plus this phase's release notes — a reader who
had followed section 13's tree would find a plugin that does not load, and the row says which
layout to use and why.

**`new-dependency-needs-a-rationale` — not-applicable.** No dependency was added. The plugin is
Markdown and JSON and the test package uses the standard library and this project's own packages.

**Beyond the three rules.**

*Did the slice stay a slice?* Seven skills, two manifests, one hook, one vendoring loop. No MCP
server, no `mcp.json`, no lenses, no second harness, no slash commands. Each of those is named in
the non-goals with the reason, and the one that was tempting — a stub `mcp.json` to make the next
piece smaller — is refused because a client reading it fails at startup.

*Was the specification followed?* Everywhere except two paths, where it describes a third party's
format and the third party disagrees. A tree following section 13 literally does not load. A77
records what the client requires, the document is untouched, and the correction is raised rather
than taken — which is the only reading of the first standing rule that leaves a working plugin.

*Is anything in the skills untrue?* Every name in them is checked against the tree: commands,
sections, gates. The one place a skill described a verdict nobody will see — `G-Test` — now says
the gate reports `not-implemented` and why the mapping is owed anyway. What remains unchecked is
whether the prose is any good, which no test here can ask.

*What did this piece learn that the plan should know?* That the harness prices the surface. WP11
argues about the tool surface being a standing cost on all work a project ever does and has no
figure; the client prints ~606 tokens always-on and ~620–830 per invocation. That belongs in
whatever decides whether a sixth or seventh operation is worth it.

*Is anything here a decision somebody else should take?* Three, in the residual risk: the
correction to section 13, whether M0's second reading is satisfied by a plugin without an MCP
server, and whether six hundred tokens a session is a cost worth the instruction.

<!-- xeno:section:release-notes -->
## Release notes

**The process ships as a plugin.** A project adds this repository as a marketplace, installs the
plugin, and gets seven skills: one per phase of the sequence, plus the learning record each phase
and each closing intent owes. Each skill says what its phase is for, which sections it owes, the
commands that carry it in order, what its gate refuses, and what it hands to the next phase.

**Installation.** `claude plugin marketplace add ./path-or-url` then
`claude plugin install xeno@xeno`. The plugin needs the `xeno` runner on the path: it is the binary
that writes the artifacts and judges them, and the skills name its commands.

**The hook comes with it.** One Stop hook calling `xeno cost turn`, so a project records a turn's
cost without writing a client settings file of its own.

**`xeno init --vendor` carries the plugin** — the manifest, the skills and the hook wiring — beside
the templates, the shipped rules and the secret filter. A second run changes nothing.

**What the surface costs, measured.** The client prices the seven skills at about 606 tokens added
to every session and 620 to 830 each time a skill fires. Nothing in Xeno's own figures had a number
for that before.

**What is not here.** No MCP server and no `mcp.json`: every step is a command, which is what the
specification requires to keep working whether a server exists or not. No lenses. One harness; the
second is next, because what leaks between harnesses only shows with two.

**The manifests do not sit where section 13 draws them.** The client loads a plugin from
`<source>/.claude-plugin/plugin.json` and a marketplace from `.claude-plugin/marketplace.json` in
the root it was given, and validates them only there. A77 records the divergence; the document is a
person's to correct.

**M0.** Section 6's reading of M0 wants one agent end to end, and that was the blocker. It is gone.
Whether a plugin without an MCP server satisfies it is a reading of the plan and not this piece's
to take.

Closes #169. Refs #1.

<!-- xeno:section:residual-risk -->
## Residual risk

**Nothing shows that the skills help, and they were written by the agent they instruct.** I wrote
down what I already do. That is why it took an afternoon and why it is the weakest thing here: an
instruction that matches its author's habits reads as obvious to the author and can read as
nothing to anybody else. The next intents are the measurement — driven with the plugin installed,
producing artifacts that can be compared with the twenty that came before — and it is the first
comparison of its kind this project could make.

**Six hundred tokens on every session, against nothing.** The cost is measured and the benefit is
not. WP13's comparison would put both in the same currency, and until it exists the honest
statement is that the plugin costs a known amount and buys an unknown one.

**Section 13 now has a known error in it.** Two drawn paths do not match the client it describes,
and the code follows the client. Somebody reading the document to build a second distribution
would get it wrong, and the register row is not where they would look. The correction is a
person's commit.

**The skills are pinned once vendored.** A project that runs `init --vendor` gets a copy that
nothing updates afterwards except another `init` with a newer runner. The templates and the rules
already have that property; the skills make it more visible, because an instruction ages faster
than a template does.

**One harness means no leakage is visible.** The hook event names, the frontmatter shape, the way a
description selects a skill: all of it is this client's. Section 6 wants the second agent
immediately after for exactly this reason, and everything in the skills that is secretly about one
client stays invisible until then.

**The hook wiring is the easy half.** One Stop event that costs nothing in context. The hook the
specification actually wants is `G-Secret` on every write, and that gate is unimplemented — so
what ships is the wiring nobody would argue about and not the wiring that protects anything.

**`mcp.json` is in section 13's `--vendor` list and not in the tree.** The function says so in a
comment. A reader comparing the list with the behaviour finds one entry missing and has to trust
the comment that it is missing on purpose.
