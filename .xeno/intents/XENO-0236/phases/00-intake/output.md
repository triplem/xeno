---
intent: github.com/triplem/xeno#183
phase: 00-intake
created: "2026-10-03T18:37:22Z"
schema_version: "1.0"
runner_version: dev+0462eb7.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 5f7b8cd6560c75bb4a1bd9c254bd3d3fb25963f1d28ffcc24f57241c6575ee8a
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

Section 7 described a resolution order for the plugin root and nothing read it. The order was the
last open clause of #183, and it had been open for three intents while each of them worked around
it.

**What it said.** "Resolution order, first match wins: the `--plugin-root` argument, an already set
`XENO_PLUGIN_ROOT`, the vendored `.xeno/plugin/` found from the git root, and only then client
specific variables such as `CLAUDE_PLUGIN_ROOT`."

**Why it could not be built as written.** `internal/gates` reads `internal/rules` and
`internal/template`, both of which resolve from the plugin directory, so a root taken from the
environment makes `rules_hash`, `strings_hash` and a rendered artifact depend on it. And A74 judges a
phase only against the rule set its own artifact records, so a changed set does not disagree with
the trail — it stops judging it. Measured on a copy of this repository: a valid rule tree that
differs leaves `gate verify` at exit 0 over 273 verdicts while G-Policy silently reports nothing
about the 75 phases that recorded the previous hash.

**The thing the section relied on arrived and did not settle it.** Section 7's own safety sentence is
"A mismatch between the two is a finding, not a reason to fall back: G-Supply fails red." That gate
was built two intents ago and made deliverable one intent ago. It does catch an override — for a
released binary. A development build carries no digest, reports `not-implemented`, and is what wrote
every artifact in this trail. So the safety the section claims holds for an adopter and not for the
work done here.

**And the bottom of the order could not work at all.** A project with no vendored plugin resolves no
template — `no template "intake" under .xeno/config/templates or .xeno/plugin/templates` — so no
phase of it renders, and `plugin_version` is absent, so G-Schema reports it. Falling back to a plugin
the client installed lets such a project fail differently rather than work. That was measured in this
intent's intake and it is what turned the question from how to guard the order into whether to keep
it.

**So every position in it was unreachable or unsafe**, and the clause had no reader in either
direction: nothing implemented it, and nothing could have implemented it safely under the build that
does the work.

<!-- xeno:section:scope -->
## Scope

**In scope.** The specification change, which is its own commit and already made: section 7 says the
plugin is the vendored one, describes no resolution order, and `XENO_PLUGIN_ROOT` leaves the list of
variables the normalised environment carries. The plan's WP7 sentence named both and now says what
replaced them. One sentence is kept about the condition an override would have to meet the day
somebody needs one.

Then what follows from it, which is comments and rows rather than code: `internal/plugin`'s package
comment and the entry point's comment described the order as a decision held open, and both now
describe it as removed. A84 said the resolution order was its open clause. The test asserting the
entry point exports no plugin root keeps its assertion and changes its reason.

**Out of scope, and each for its own reason.**

Implementing an override under a condition. It was the alternative — honour it only where the binary
carries an anchor — and it was declined in favour of removing the clause, because a clause with no
reader is cheaper to delete than a mechanism with a guard is to maintain. The sentence section 7
keeps is what a later intent would start from.

Changing anything about how the plugin is found. The code already read the vendored directory and
still does; this is the document catching up with it, which is the opposite of the usual direction
and is why there is no code in the first commit.

`XENO_PLUGIN_DATA`. Still in the list, still set by the entry point, still read by nothing because
section 7 pins it to one value. Recorded in A84 and unchanged here.

G-Supply's behaviour. It compares the vendored tree, which is now the only tree, so nothing about it
moves. That it is inert under a development build is A86's and is unchanged.

The three other things #183 named: the entry point, `XENO_HARNESS` and `XENO_PLUGIN_DATA`. All
settled in the intent that built the entry point. This closes the issue's last clause.

<!-- xeno:section:context-rationale -->
## Why this context

Section 7's environment normalisation is read whole, for the fourth time in six intents, and this
time for what to remove rather than what to implement: the variable list, the resolution order
paragraph, the safety sentence that rests on G-Supply, and the two constraints on hooks.

Section 5's field list is read for `rules_hash` and `strings_hash`, because they are what an
override would make environment-dependent and the reason the order could not be honoured in the gate
path.

A74 is read for the mechanism that makes the failure silence rather than disagreement, which is the
measurement's explanation and the sentence that decided this.

A86 is read for what G-Supply does and does not give: it compares the vendored tree against a digest
the binary carries, and a development build carries none. That is why section 7's own safety sentence
does not cover the build that writes this trail.

`internal/gates`' dependency graph is read rather than assumed — `go list -deps` shows
`internal/rules` and `internal/template`, which is the fact the whole argument rests on.

`internal/template`'s resolution is read for the message it produces with no plugin, which is what
established that the bottom of the order is unreachable rather than merely weak.

`internal/plugin/plugin.go` and `.xeno/plugin/bin/xeno-env.sh` are read for the comments that
described the order as a decision pending, because they are what follows from the document change.

The three shapes put to the maintainer are read back out of this session: the order under a
condition, the order as written with A42 amended, and the split by purpose. The third was retracted
in the exchange that produced this decision, because a gate path pinned to the vendored tree would
break the one case the client fallback existed for — and then the fallback turned out to be
unreachable anyway.

Nothing outside the repository is needed. Both measurements were taken on copies.
