---
intent: github.com/triplem/xeno#172
phase: 01-requirements
created: "2026-10-03T10:34:33Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2d52dbc0404c2707477d8d29760ffa12c765e485efeb5c900c32b0264feb0605
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**A link naming a document that is not in the tree is a G-Schema finding.** The cause names the
component the link declares, the path it gave, and that the file is not there. The next step says to
correct the path or take the link out. The file the finding names is the profile, because the
profile is what is wrong.

**A link whose document exists is unchanged.** It joins the information base with its hash, last,
after the files the `include` patterns claimed — which is what #171 settled and what this intent
must not disturb.

**The check reads the profile where the profile is.** From P0, whatever phase is being judged,
which is what the budget check established: a profile mistake is the intent's mistake and not one
phase's.

**It does not block.** A missing link's document is a finding like the budget's, decidable and
recorded, and `phase start` still resolves the base and still starts the phase.

**A profile with no links produces nothing**, and a repository with no profile produces nothing,
which is every project today.

**The byte count is not touched.** The finding section 5's enumeration would need is named in the
review with the wording a person would add, and nothing in this intent writes a field.

**Nothing else moves.** The base's contents and order are unchanged, no rule set changes, no sealed
artifact is rewritten, `./xeno gate verify` stays at exit 0, and the suite stays green.

<!-- xeno:section:non-goals -->
## Non goals

**No field in the lock.** Recording the byte count is a change to section 5's enumeration and
therefore a person's commit before any code. Named, worded, not taken.

**No refusal at `phase start`.** A profile mistake is reported at the gate like every other
profile mistake. A second mechanism for one error would make the cheap check the loud one.

**No validation of `include` or `exclude`.** A pattern matching nothing is a legitimate state: the
files may not be written yet, and a project excluding a directory it does not have yet is being
careful rather than wrong.

**No check that a link's component exists.** The component is a path prefix in the project's own
vocabulary, which may name a directory, a module or a thing the repository does not model as a
directory at all. The document is the part that is unambiguously a file.

**No profile for this repository.** That is the experiment this unblocks, and it comes after, on
one intent, with the releases counted.

<!-- xeno:section:constraints -->
## Constraints

**The criterion is #171's and it is quoted rather than reinterpreted**: "a link naming a document
that does not exist is a finding against the profile rather than a silently missing file".

**The finding belongs to G-Schema**, beside the budget check that reads the same two files. The
gate list is a budget and nothing here needs a new gate.

**The base must not change.** Which files reach it, and in what order, is #171's design and this
intent is a check rather than a change to the resolution.

**No register row.** `ASSUMPTIONS.md` closed with M0, so the one decision this intent takes — where
the check lives — is recorded in its design phase, which is the loop the plan switched to.

**Short artifacts, deliberately.** One condition and one test. XENO-0224 measured the floor of the
sequence at about sixty-two thousand bytes with shortness as a constraint; this intent has less to
say than that one did.

**No new dependency, 88 columns, SPDX, `gofmt`, `go vet`, the suite, and `./xeno gate verify` at
exit 0.**

**One intent, one branch, one issue.** `172-a-link-naming-a-document`, #172, labelled wp8, off a
`main` that carries WP8 and a closed M0.
