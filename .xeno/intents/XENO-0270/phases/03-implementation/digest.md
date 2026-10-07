---
intent: github.com/triplem/xeno#231
phase: 03-implementation
created: "2026-10-07T09:46:19Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a3ae95c6123c085152e9edf90f51f9f202611c0195021d91d6f8d8b05d43ab8a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`README.md` rewritten: section 1's problem first, then what Xeno does about it, then evidence
rather than a management system, then the name with Appendix C named for the rest, then the
index with the process definition marked normative, then four commands in the order they are
run, then telemetry, then one paragraph each on this trail and on the register.

Three new files. `docs/README.md` is the index, every file in `docs/` with one line on what it
is for, grouped by the errand a reader arrives with and the two normative documents marked in
their own entries. `docs/commands.md` carries the `usage` constant in a fenced block, with the
per command notes that moved out of the README after it. `docs/the-trail-in-this-repository.md`
takes the version fields, the two shapes of the intents directory, the coverage table and the
two mutations.

One test, `TestThePublishedReferenceIsTheUsage`, beside the one that reads the command names
out of the same constant.

Four deviations, two of them corrections to P2 found by checking that nothing was lost rather
than trusting it. The reference does carry per command prose, because five paragraphs had
nowhere else to go and criterion 10 is the stronger requirement. The no-network-call paragraph
is kept rather than dropped, because section 12 carries its first half and not the second, and
the same reading corrected a wrong sentence in the index. One heading was renamed and two
paths rewritten, which the move forces. And the old `learning record` paragraph turned out to
be two paragraphs with no blank line between them, so the first extraction carried one into
the wrong file.

Build, suite, `gofmt`, `go vet` and `gate verify` over 498 verdicts all pass, and every
relative link resolves.
