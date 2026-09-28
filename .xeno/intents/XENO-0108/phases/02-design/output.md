---
intent: github.com/triplem/xeno#108
phase: 02-design
created: "2026-09-28T16:58:24Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+023193e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: b544e2ccc83490b91cce9a1b9c0a90161c68b6c92dd72c6ff5999086bd9b7781
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: by-hand
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The writer is `SectionSet`, and it hashes with `hashing.FileHash`.** The field covers
a file, the file lies beside the artifact, and the function that computes the value is
the one G-Schema recomputes with. Writing the hash any other way would let the writer
and the checker drift apart while both looked correct, which is the failure a shared
Appendix B implementation exists to prevent.

**Written on every render, not once at creation.** `SectionSet` carries the frontmatter
of an existing file over whole, so a field set only when the file is created would be
whatever the first writer left behind, including a value a hand had edited. Writing it
on each render makes the rendered file agree with the lock it was rendered beside, which
is the only claim the field makes.

**An absent lock leaves the field absent.** `FileHash` returns an error and the field is
skipped. The file that is missing is `context.lock.yaml`, and G-Freshness is the gate
that speaks about that file; a hash of nothing, or a refusal from the section writer,
would put the finding in the wrong place and blame the wrong writer.

**A35 and A51 are amended in place.** Two sentences that were right and unfinished, not
two sentences that were wrong. The project's rule about a paragraph changed a second
time is what decides the form: the row is read back as a row, and the amendment names
the issue that drew the consequence.

<!-- xeno:section:alternatives -->
## Alternatives

**Write the field once, when `output.md` is created.** Cheaper by one hash per render,
and wrong in the case that matters: the frontmatter of an existing file is carried over
from whatever wrote it, so a file whose frontmatter a hand had touched would keep the
hand's value and the gate would compare the lock against something nobody could account
for. The cost of the rejected variant is one file read per section write, which is
smaller than the cost of the case it fails.

**Have `phase start` write the field into `output.md`.** It knows the value at the
moment it writes the lock. It would have to create `output.md` to put the field in it,
which makes `phase start` a writer of the rendered artifact, and `SectionSet`'s
guarantee that the file on disk is always in rendered form would then have a second
author. Rejected for the ownership, not for the mechanics.

**Move the check to G-Freshness and let the field stay absent.** This is A51 read
backwards. It would make the gate agree with the runner by lowering the gate, and
section 5 requires the field either way, so the artifact would still be incomplete
against the specification and the finding would simply have moved.

**Refuse a `section set` where no lock lies beside the artifact.** It would make the
missing lock loud at the moment it is cheapest to fix. It also makes the section writer
the gate for a file it does not own, and it turns writing prose into an operation that
fails on the state of a neighbouring file. G-Freshness already reports it, one layer
later and in the right voice.

<!-- xeno:section:impact -->
## Impact

`internal/runner/runner.go`, one writer and its comment. `SectionSet` gains four lines
and the paragraph that says why the value cannot drift inside a phase.

`internal/runner/runner_test.go`, two tests. One reads the field back and writes a
second section to show it does not move; one walks the supported path and asserts that
no finding names `context_hash`.

`ASSUMPTIONS.md`, two rows amended.

No gate changes. `recomputed` already hashes the lock to check the field and has been
correct the whole time; what it lacked was a value produced by a writer rather than by a
hand.

Every artifact produced from here on carries a `context_hash` the runner wrote.
Artifacts already sealed keep the hand-written value they were judged with, which is the
same value the same lock produces, so nothing that exists is invalidated: `gate verify`
recomputes 52 verdicts unchanged.

The fixtures that write the field by hand stay as they are. They test G-Schema against
values they control, including wrong ones, which is what they are for; the new tests are
what covers the produced value.
