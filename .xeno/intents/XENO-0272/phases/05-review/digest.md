---
intent: github.com/triplem/xeno#277
phase: 05-review
created: "2026-10-07T14:18:45Z"
schema_version: "1.0"
runner_version: dev+27fc33e.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1aa23a71d6c8988dc5c7a225f5a9d93d45d22d63b0932a1ee7040ace45294d41
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The three standing rules hold: section 12 is acted on and not amended, a term is removed rather
than anything invented, and the change belongs to WP1 and to intent XENO-0272 for issue #277.

#277's done-when offered two branches and this is the first: the exemption no longer rests on a
field nothing checks, and what remains in `honest` is two facts about the tree.

What a reviewer should look at first is that two sealed phases changed verdict. XENO-1's and
XENO-2's intakes were green and are now approved, permanently, because a check got stricter
after they were sealed. Both alternatives that would have avoided it were worse: keeping the
exemption leaves a verdict resting on a declaration, and writing the real hash into the four
artifacts is the rewriting section 11 forbids.

The second thing is about the agent rather than the code. The four approvals are a person's
statement and the agent typed them, after the maintainer chose the shape and approved the
wording. `next.go` refuses to offer `gate approve` as a command so that nothing nudges towards
that decision, and running it four times is past nudging. The decisions can be withdrawn and
retyped; nothing else depends on them.

Two questions no rule asks are answered: whether a comment and a register row are enough to
keep the rule the change establishes, when a new term from a declared field would pass the
suite; and that the rule is met while `goneBundle` is coarser than it sounds, which #273
measured and left standing.

Three rules answered, two of them deviations: the traceability rule, because P3 recorded two
and both name what they depart from, and the migration rule, because there is no interface
change and a gate got stricter, which has the same consequence for somebody with artifacts
already.

After this phase was first judged, CI refused the branch: approving the four findings put
XENO-1 and XENO-2 in the commit range, and `Completeness` requires every touched intent to be
complete or abandoned, which a pre-M0 intent whose later phases never started is not. There is
no way to resolve a finding on such an intent without writing to its files. The maintainer was
put three options and chose closing the two, which is the ending section 8 already offers; each
now carries a reason saying nothing was dropped and an intent level `learning.yaml` recorded as
`--no-finding`. It closes two of about 106 intents in that shape, and the first close ran half
way and was reset and run again. All of it is recorded here rather than in P3, because it
happened after P3 was judged.
