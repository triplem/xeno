---
intent: github.com/triplem/xeno#256
phase: 03-implementation
created: "2026-10-06T08:37:26Z"
schema_version: "1.0"
runner_version: dev+5163d1b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1a36a876017b4b75a663c408d0303e9cba1eb3e408286b5e61b3508c144832fb
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

One file, one paragraph, ten lines added. `CLAUDE.md` goes from 87 to 97.

The paragraph sits under Conventions, directly after the one about putting a decision one at a time,
with no new heading. It states the rule as an act — recreate the thing and run the check again before
writing the finding down — then why it matters here rather than in general: a finding goes into an
artifact and is sealed with it, so a wrong one is permanent rather than corrected, with the correction
somewhere a reader of that phase will not be.

It names one instance, #201, with the three facts that make it checkable: the directory was recorded
as not gitignored, the entry was at `.gitignore:10` all along, and `git check-ignore` had run moments
after the directory was deleted. It closes with "Nothing checks this (#256)", which is the form the
sequencing paragraph beside it uses.

**The instance was re-verified rather than recalled**, because a paragraph about verifying a negative
resting on a remembered one would be its own counter-example. `.gitignore:10` carries the entry;
`git log -S` puts it in `154bb09`, which is #200; and the false claim is in two of XENO-0249's sealed
phases, P4 and P5, by grep rather than by memory.

Two corrections inside the phase, both caught by reading the result rather than by a tool.

The first draft listed `test -f`, `grep -l` and `find` alongside `git check-ignore`, which is exactly
what P1's non-goals ruled out — "the paragraph names the class and one instance; #256 holds the
table". The list went; the class is now one clause and the issue keeps the table.

The second draft was ten lines and had three at 89 columns after a hand rewrap. It was reflowed with
a wrapper rather than by eye, which is what the 88-column rule needs when a paragraph is replaced
whole, and the result is nine lines inside the bound.

Nothing else changed: no normative document, no row in `docs/clause-readers.md`, no register row, and
no code.

<!-- xeno:section:deviations -->
## Deviations from the design

One deviation from this intent's own non-goals, caught in the writing. P1 said the paragraph would
name the class and one instance and would not list the tools that share the ambiguity, because that
turns a paragraph into a table and #256 already holds one. The first draft listed four tools anyway.
It is the kind of error the non-goal exists to catch and it was caught by reading the draft against
the section rather than by anything mechanical — which is the honest account of what non-goals are
for in a change this small.

One correction of method rather than of content. The paragraph was rewrapped by hand after being
shortened and came out with three lines at 89 columns, one over. It was then reflowed with a wrapper
over the whole paragraph, which is what replacing a paragraph whole means in practice; rewrapping by
eye is how a paragraph ends up one column over three times in a row.

No deviation on the framing, which was the design's substantive decision. The paragraph is about
sealing rather than about carefulness, and it says so in the sentence beginning "It matters more here
than elsewhere" — the clause that is the whole argument for the addition being in this file at all.

No criterion was falsified. Criterion 6 asked that the file stay short enough for the addition to be
defensible against its own last line; it is ten lines on 87, which is a judgement rather than a
measurement, and P4's results states the figure rather than claiming the judgement was verified.
