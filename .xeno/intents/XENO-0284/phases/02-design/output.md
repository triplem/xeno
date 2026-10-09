---
intent: github.com/triplem/xeno#337
phase: 02-design
created: "2026-10-09T15:54:20Z"
schema_version: "1.0"
runner_version: dev+9590797.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2990c1f20bfdb9f8d6ffda27ba829e18c1842991a8e43860b2bfd0da6965f3e5
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

**Four places in the plan and no new section.** Section 1's deferred list, a package
in section 2, the sentence and the table row in section 6 that already carry WP18, and
section 8's deferred list. Every one of them is where WP18 already is, because WP18 is
the one precedent for a package specified and not built and a reader who finds the
dashboard finds the record beside it. A section of its own, or a new "raised by"
paragraph in section 8, would have put the reasoning in two places; WP21's own text
is where the reasoning lives and the other four are pointers.

**WP21 is written in WP18's shape.** An opening paragraph that says what the stage is
and why it is specified now and built later, bold-led paragraphs for the shape, the
contents and what building it touches, and a done-when. The shape paragraph cites #224
in one sentence rather than repeating its argument, because the issue's done-when asks
for exactly that and because a plan that restates an issue's analysis ages separately
from it; what it adds is the reason section 12 contributes, which #224 predates.

**The done-when is three properties and not a feature list.** A repository with no
issue reaches its first approved intent through a verifiable record; every issue the
record produced names it and every P0 started from one quotes it; a record run here
costs a figure #117 can place beside an intent's. The third is the plan's own rule
that proportionality is measured and not argued, applied to the package before it
exists.

**The names are left to the package.** "A key sequence of its own", "a directory
beside `.xeno/intents/`", "a seventh template", "the commands that write and judge the
record": each is the shape the decision fixed and none is a name, because a name in
the plan is a specification the process definition does not yet carry, and the second
standing rule puts the specification first.

**The plan was edited in place and read as a diff before it was offered.** Three for
three, every clause drafted in prose had needed changing on contact with the file;
this one was written into the file, wrapped, read back, and the one change after the
read was trimming the #224 paragraph from a restatement to a citation.

<!-- xeno:section:alternatives -->
## Alternatives

**A "raised by the greenfield case" paragraph in section 8 instead of a package.** The shape section 8 has for requirements raised outside the document. Declined: the decision made it a package, and a package has a done-when, which a paragraph does not.

**A clause in the process definition now.** Would fix the record type and its path while the decision is fresh. Declined: it designs a 1.1 package in a v1 intent, and the plan's own text says the process definition moves when the package is built.

**Naming WP21 in the Medium row.** D-2's cost line said Medium at least. Declined: the plan sizes what it sequences, and the figure "twenty packages built for v1" that follows the table would need a caveat.

**Restating #224's three reasons in WP21.** More self-contained for a reader. Declined: the done-when asks for a citation, and an argument copied from an issue ages apart from it.

<!-- xeno:section:impact -->
## Impact

**For the plan.** Twenty-one packages specified, twenty built for v1, two deferred; the sentence after the size table stays true.

**For 1.1.** One package more in the scope of record, with a done-when to be met and a process definition change to be made first.

**For this repository.** Nothing changes in how work starts: issues are written by hand and approved one at a time until WP21 is built.

**For the trail.** Nothing sealed moves. The three decisions live in XENO-0284's intake, where `intent status` lists them.
