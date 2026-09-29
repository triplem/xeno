---
intent: github.com/triplem/xeno#136
phase: 01-requirements
created: "2026-09-29T18:06:43Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c05acae.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 727b63c8e4abb7784299d1767ee5fe6abf2fa78703da4e4ee91bb590c0323666
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

**AC1.** The listing's heading reads `CREATED  INTENT  STATE  PHASE`.

**AC2.** The one-intent form's heading reads `PHASE  STATE  VERDICT`.

**AC3.** Nothing else about either output changes: the same words, the same widths, the
same order, the same default of ten, the same notice, and the same values in every row.

**AC4.** The alignment still holds at the widest data each column can take. Upper case
changes no width, so XENO-0207's two tests cover it and are updated to the new strings
rather than replaced.

**AC5.** No other output changes case. Verdicts, states and findings stay as they are,
being values and sentences rather than labels.

**AC6.** XENO-0207's design decision is named as reversed, in this intent's record,
rather than edited where it was judged.

**AC7.** Everything green stays green and `gate verify` matches every verdict.

<!-- xeno:section:non-goals -->
## Non goals

The words. `created`, `intent`, `state`, `phase`, `verdict`, in upper case and otherwise
unchanged.

The widths, the order, the default, the notice, and the data behind any of them.

XENO-0207's phases, which carry verdicts. The reversal is recorded here.

This project's convention about prose headings, which is unchanged and was never in
question.

The case of anything else the tool prints.

A general rule about output case. One table is the subject, and `gate run` prints
findings rather than a table.

<!-- xeno:section:constraints -->
## Constraints

Nothing under `docs/` changes.

The heading keeps sharing its format string with the rows, which is what makes the
alignment hold, so the change is to the arguments and not to the format.

Upper case must not change a width. It does not, since the columns were sized for the
data rather than for the labels, and AC4 is the check.

XENO-0207's sealed phases are not touched. Section 11's rule is the reason and the
intake says so.

The tests are updated, not added to. The property they assert is unchanged and only the
expected strings move; a second pair would test the same thing twice.

One intent, one issue. The commits reference #136.
