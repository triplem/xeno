---
intent: github.com/triplem/xeno#257
phase: 03-implementation
created: "2026-10-06T07:56:17Z"
schema_version: "1.0"
runner_version: dev+77a35b1.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7e2ae78eccded6a15a6b1e60f05835ec4aa9e7983d762c2d61936ddc5cfed888
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

Thirteen files added under `examples/templates/`, nothing else changed.

Six directories, one per phase id, each with `template.yaml` and `strings.en.yaml`. The bundle is
copied from the shipped one unchanged, because the headings belong to the sections and not to the
flags, and because `Load` treats a missing bundle as an error rather than falling back to another
language.

Six of the seventeen stay required:

| phase | required |
|---|---|
| intake | `problem` |
| requirements | `acceptance-criteria` |
| design | — |
| implementation | `changes` |
| verification | `test-mapping`, `results` |
| review | `release-notes` |

Eleven become `required: false`. Nothing is deleted: every id, `version` and `phase` matches the
shipped template, which was checked by comparing the parsed section lists rather than by reading.

Each `template.yaml` carries a header in `examples/rules/`'s shape — what it is, that it is enabled
nowhere, the measurement behind it, how to adopt it — plus two lists specific to that file: what it
drops with a reason each, and what it keeps required with a reason each.

Two of the keeps say what kind of reason they have, because the kinds differ and conflating them is
how a reader infers a gate that does not exist. `release-notes` says **READ BY A GATE** and names
`release-notes-are-filled` and its `section-non-empty` check. `acceptance-criteria` and
`test-mapping` say "by decision, not by a reader", and name #258.

The review override says that the `review_checklist` frontmatter is read by G-Policy whatever the
`review-checklist` section's flag is, so the prose becomes optional and the answers do not.

`examples/templates/README.md` carries the figures from nine intents, the table of what is kept and
on what grounds, the measurement that one of seventeen sections is read by a gate, how to adopt a
subset, what to collect, and one instruction that is there because of this session's own mistake:
**do not measure a timing**, because four comments on #117 read agent latency as process cost before
anybody noticed.

**The candidate was validated against the shipped set rather than eyeballed.** A script parsed both
and asserted the section id lists are identical, the `id`, `version` and `phase` match, and every id
the template defines has a heading in its bundle. Six required across the six templates, counted
from the parsed files.

<!-- xeno:section:deviations -->
## Deviations from the design

No deviation from the design. The six keeps, the eleven drops, `required: false` rather than
deletion, the per-file reasons, the two kinds of keep distinguished, the README rather than six
copies of the method, and `examples/` over the shipped set all landed as P2 specified.

One addition beyond the design, taken from this session rather than from the issue. The README
carries an instruction not to measure a timing. P2 said what to collect and did not say what to
avoid; four comments on #117 reported per-phase elapsed figures before noticing they measure the
agent's session shape rather than the process, and a fixture whose whole purpose is a comparable
measurement should not let the next person repeat that. It is one paragraph and it is the only part
of the README that exists because of a mistake rather than a finding.

One correction inside the phase. The README came out with ten prose lines at 89 and 90 columns
against the 88 the conventions ask for. They were reflowed by replacing whole paragraphs rather than
breaking at the overflow, which is what the conventions ask of prose changed a second time.

One thing was generated rather than written by hand, and it is worth saying which. The twelve
template and bundle files were produced by a script that read the shipped set, moved the flags and
copied the bundles, so the ids and headers are the shipped ones by construction rather than by
transcription — and then checked by a second script that parses both and compares. The headers and
the reasons are written; the structure is derived. Transcribing seventeen section entries by hand
across six files is exactly where a wrong id would have gone unnoticed.
