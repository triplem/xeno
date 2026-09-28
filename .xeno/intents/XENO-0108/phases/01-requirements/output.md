---
intent: github.com/triplem/xeno#108
phase: 01-requirements
created: "2026-09-28T16:57:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+023193e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 6ea202ada72d56477f14abe6de27214bda1c8532d30444966effc808e11eb69c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**AC1.** A phase started with `phase start` and written with `section set` alone carries
a `context_hash` in the frontmatter of `output.md`, and the value is the sha256 Appendix
B computes over the `context.lock.yaml` lying in the same phase directory.

**AC2.** The walk section 5 supports — `phase start`, then a `section set` per required
section — produces no finding naming `context_hash`, from G-Schema or from any other
gate. The other five fields A35 leaves out are still reported missing, unchanged.

**AC3.** A second `section set` in the same phase leaves the value it wrote first alone.
The field does not move while a phase is open.

**AC4.** A `section set` where no lock lies beside the artifact writes no `context_hash`
rather than a hash of nothing, and refuses nothing on that account.

**AC5.** `ASSUMPTIONS.md` records what changed about A35 and A51 without deleting what
either row said.

<!-- xeno:section:non-goals -->
## Non goals

The other five session fields. `model`, `tool` and `tool_version` come from a harness
the runner cannot see, and `secrets_hash` and `rules_hash` have no writer anywhere in
this tree. A35 stands for all five, and a phase produced by the supported path stays red
on them; that is a different issue with a different work package.

`digest.md`. Section 5 lists `context_hash` on both files a session produces, and
nothing in the runner writes `digest.md` at all — `phase finish` judges it and the agent
writes it. A writer for it is not this change.

The shape of `context.lock.yaml`. A10 still names the blocks it does not carry, and this
change hashes the file as it stands rather than completing it.

Resolving what a hash covers over a network. Unchanged and out of reach, as A8 and
limitation 11 already say.

A field written once per phase rather than once per render. Rejected in design rather
than left out here, and the reasoning belongs to P2.

<!-- xeno:section:constraints -->
## Constraints

The documents are not editable here. Section 5 already requires the field and Appendix B
already defines its value, so nothing in `docs/` needs to change and nothing in `docs/`
may. This is the code catching up with a specification that was always ahead of it.

A35 and A51 are amended, not replaced. Both rows said something true; only their
consequences were unfinished. The project's rule about changing a paragraph a second
time applies to the row as a paragraph, and the amendment is written as one.

No new dependency, no new field, no new gate. The field exists in section 5's
enumeration and the check exists in G-Schema. Adding either would be a specification
change first.

The value must be recomputable by a stranger. Appendix B fixes it, `hashing.FileHash`
implements it, and the writer must use that function rather than hashing the file its
own way, or the writer and the checker could drift apart while both looked right.

One intent, one work package. The issue carries `wp7`, the branch carries this intent,
and the commits reference #108.
