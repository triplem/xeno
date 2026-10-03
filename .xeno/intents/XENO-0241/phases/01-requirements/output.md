---
intent: github.com/triplem/xeno#206
phase: 01-requirements
created: "2026-10-03T20:44:22Z"
schema_version: "1.0"
runner_version: dev+8574810
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ceb3f25885139b68b7eb77825a75163b65ed2d62921e3902b18c422feda20525
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

`xeno intent verify --base REF --head REF` exits 1 where an intent the range touches is
neither `complete` nor `abandoned`, naming each one and the state it is in, and exits 0
where every touched intent is one of the two.

The state is the one `xeno intent status` prints. It is read from the same function, not
derived a second time: two derivations of one answer can disagree, which is worse than
the gap either was written to close.

A range that touches no intent exits 0 and says so in those words, so that a reader of
the log can tell "nothing to check" from "checked and clean".

Both ends of the range are required. An absent, empty or unresolvable ref exits 2, which
is "could not run" and not "clean", because a check that passes on no evidence is worse
than one that is missing.

An intent whose directory the range touches but whose `intent.yaml` cannot be read exits
1 with the reason. The listing's habit of reporting such a row rather than leaving it out
becomes a refusal here, because this command's answer is a verdict.

The `verify` job calls it, next to the trail-rewrite step and with the same base. A pull
request whose intent has not reached a decided P5 is therefore red until it has, which is
what the other gates already do.

Both shipped CI wrapper templates carry the call, so the check is what an adopter gets
and not a property of this repository.

Tests cover: a complete intent passes; an intent stopped at P3 fails and is named; an
abandoned intent passes; a range touching no intent passes; a missing base exits 2; a
provisional P5 fails.

<!-- xeno:section:non-goals -->
## Non goals

No change to G-Complete and no new gate. The gate list is a budget and the clause it
carries is already correct; what this adds is a second place that reads the clause, at the
one moment the process definition says a merge is decided.

No specification change. Section 7 says G-Complete runs at P5 and at intent close, and
both still do. The check asserts that a change arriving at a merge has reached one of
those two endings, which is section 8's sentence about how an intent ends rather than a
new rule about gates.

Nothing that writes. The command reads the trail and the range and reports, which keeps
section 12's rule that CI writes nothing intact rather than carving at it.

No completeness check for a branch that touches no intent. The rule that every change
belongs to an intent is real and unenforced, and enforcing it here would hide it inside a
check about something else. #120 owns it.

No migration. The trail is complete today, measured before this was designed.

No edit to `CLAUSE-READERS.md`. It is a dated measurement that says so about itself.

No new field, gate, tool or rule.

<!-- xeno:section:constraints -->
## Constraints

The gate path reaches no network and no subprocess, and the `verify` job greps for both.
This command is not a gate and `internal/git` already starts a subprocess for the commit
predicates, so the range may be read there; what it may not do is reach `internal/gates`,
which is why the reader sits in `internal/runner` and not beside G-Complete.

Section 9's rule that the commit range is always passed in and never inferred. `Commits`
refuses an empty base or head for that reason, and the new path refuses the same way
rather than falling back to a default, so there is no ref the tool chose for itself.

The check becomes red on a pull request whose trail is not yet at P5, which is a cost
rather than a flaw: it is what every other gate does, and a required check that is green
until the last push would report nothing at the moment it matters. Said here so that it
is a stated consequence and not a surprise.

A rename inside `.xeno/intents/` is already refused by the trail-rewrite step, so the
range reader asks git not to detect renames and treats both paths as touched. That is the
conservative reading, and it cannot disagree with the step above it.

Prose at 88, tables and code blocks exempt. Go at whatever `gofmt` produces.
