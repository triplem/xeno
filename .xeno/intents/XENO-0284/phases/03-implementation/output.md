---
intent: github.com/triplem/xeno#337
phase: 03-implementation
created: "2026-10-09T15:59:03Z"
schema_version: "1.0"
runner_version: dev+9590797.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f64c5b21370d14d6db3cfe941fdd1073cf7d7bdd0cc9d04c53e2e9e0f495d241
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

One file outside the trail, `docs/implementation-plan.md`, 53 lines in and 6 out, in
`ec372ea`, which is the first commit after `main` on this branch and the only one the
plan is in.

**Section 1.** The paragraph listing what waits for 1.1 gains a last sentence naming
the entry point before the first issue as WP21, the greenfield case, and the one item
there that adds a record rather than reach.

**Section 2, `### WP21 Ideation, deferred to 1.1`.** Forty-four lines after WP20 and
before the architecture section, in WP18's shape: what the stage is and why it is
specified now and built later; the record as a sealed directory beside the intents
holding what a phase holds, not a phase for #224's reasons in one sentence and for
section 12's reason in another, and a thing a gate can seal because the approval asks
for gated documents; what it holds, the vision and the plan with decisions put one at
a time, and what it produces, issues through the adapter that carry its key and that
the approval act takes one by one; what building it touches, each a specification
change first, with the MCP budget argued then and not now; and a done-when of three
properties.

**Section 6 and 6.1.** "WP18 is specified but not built" becomes the plural with
WP21, and the size table's unsized row names both.

**Section 8.** The deferred list's sentence on the dashboard names the ideation
record beside it, each specified in section 2 and not built.

**The edit was made in the file, not drafted in prose.** Written with the line wrap
checked, read back as a diff, and trimmed once: the #224 paragraph had restated the
issue's three reasons and was cut to a citation with one sentence of it, which is what
the done-when asks. The maintainer read the diff and approved it as written, and the
commit message records the exception to the first standing rule.

**Measured on the branch.** `git log main..HEAD` is the plan commit alone before the
trail; every added line of the plan is within 88 columns; `go test ./...` is
unaffected by a document change and stays green; the `gate verify` figure is P4's.

<!-- xeno:section:deviations -->
## Deviations from the design

None. The design named four places in the plan, each where WP18 is, with WP21 in WP18 shape, #224 cited in one sentence, a three-property done-when and no names, and the file carries each as written. The trim of the #224 paragraph happened before the diff was offered and is inside the design, not a departure from it.
