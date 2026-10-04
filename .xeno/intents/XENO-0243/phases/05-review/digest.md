---
intent: github.com/triplem/xeno#208
phase: 05-review
created: "2026-10-04T08:48:17Z"
schema_version: "1.0"
runner_version: dev+49f2794.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bc732b8a221d9cfb93ad730806787b3edf48cb78107ea7c2819355be8579429d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The review answers the three shipped rules, two met and one not applicable, and turns on
the one check a checklist cannot make: the single change to a gate was recomputed
against the whole trail, 336 verdicts at exit 0, because a gate change is the one kind
whose cost is paid by phases nobody is working on. The release notes say what a reader
who did not follow the work needs: the command, its three forms, the hash no flag
supplies, the two provenance fields written for the first time, the test report CI now
publishes, and the pull path of section 6 having run end to end — four items declared in
P4 by the command, P4 finished provisional, the branch pushed, and all four attached at
this phase's start from the pipeline's own artifacts with their hashes, run ids and
commit, P4's `artifacts_hash` unchanged. That is the first attachment from a real
pipeline in this repository and the first verification phase whose evidence is a file no
person wrote. Seven residual risks are recorded rather than absorbed. The first is that
nothing makes a verification phase declare anything, which is #188's risk in a second
field and has the same answer: a figure, not a gate, because a gate would be met with a
declaration of something nobody ran. The second is #221, the attachment's unjudged
`result`, left to a decision about fifteen verdicts. The third is that `produced_by` is
a self-report with no reader, which is #207 arriving in a third field. The checklist
itself is now the only frontmatter block of section 5 with no writer, which this phase
had to type by hand and which its learning record proposes a command for. Read in this
phase: the four rule files for what each one asks, the attached reports and
`attached.yaml` as they arrived, P4's verdict before and after the attach, and
XENO-0242's P5 for how a checklist entry is written.
