---
intent: github.com/triplem/xeno#156
phase: 03-implementation
created: "2026-10-01T14:38:57Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8e3f7ab4846b70e1f0ca6c66b9149a378b2ea401c98dc142b3e2728a5315f4df
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: by-hand
---

# Implementation

<!-- xeno:section:changes -->
## Changes

Two files, 58 lines added and 19 removed, no file under `cmd/` or `internal/` touched.

**`ASSUMPTIONS.md`, "Not built", replaced.** The paragraph now opens on the harness — the
MCP server, the hooks and the plugin — and says what two gates wait on: G-Supply verifies a
plugin against the digest the runner carries, and G-Secret belongs to the hook the plan puts
on every write. Rules follow, with the two consequences that make them visible from
elsewhere, G-Rules and G-Policy standing as `not-implemented` and a learning having no rule
set to be merged into. Then the documentation site, then WP12's tracker half, with the port
and both adapters named as built and the open part named as carrying issue content into P0
and reporting a verdict onto a merge request. The secret filter and the digest writer are
gone from the list. A second paragraph puts the verification points in their current state:
the marketplace URL open, the harness headers answered per harness with `phase start
--export` sending them (#58), the gateway's two answered against LiteLLM 1.102.1 (#56).

**`ASSUMPTIONS.md`, "Built", one paragraph added** after the WP9 and WP10 entry, opening
"Built since that list was written, which reads as complete and on its own is not" and naming
`xeno section set`, the three `assumption` commands, the digest and the filter it passes
through with A62, the cost record and `xeno cost turn` with A63 and A64, the branch rules port
with both adapters and A65, and the symbol index reader with its degradation.

**`README.md`, heading and opening, replaced.** `# xeno, core` becomes `# xeno`. The opening
names what the binary carries: hashing, gates, the phase sequence, evidence attachment, the
digest and its secret filter, the cost record, the branch rules port and the symbol index
reader, and it points at the register's last two sections for what is built and what is not.
A second paragraph narrows the network claim to the one that holds — the gate path makes no
network call, `enforcement check` is the one command that does, because asking a host what it
enforces is the one question the repository cannot answer about itself.

**`README.md`, command block, brought level with the dispatch table.** Added: the three
`assumption` commands, `cost turn`, `--export` on `phase start`, `--summary` on
`phase finish` with the note that it writes `digest.md`, and `[--all]` on `intent status`. A
paragraph below it says what `intent status` without an intent lists, what `cost turn` reads
and attributes, and that every command takes `--root`.

**`README.md`, trail section, replaced.** Two paragraphs. The first keeps A20's shape for the
intents up to `XENO-0106` and points at `M0.md`. The second says that from `XENO-0107` the
intents run all six phases and finish green, names the second half of G-Freshness and #63 as
what made a later phase judgeable, and sends the reader to `xeno intent status --all` rather
than carrying a count.

**`README.md`, coverage table, six rows added and one extended.** WP8, WP11, WP12, WP13, WP15
and WP17 after the WP7 row, each naming tests that exist: `TestNoProfileRecordsNoInformationBase`
and `TestSectionSetWritesTheContextHashOfTheLockBesideIt`;
`TestExportPrintsTheEnvironmentAndNothingElse`; `host_test` with the two adapter packages;
`cost/cost_test`; `index/index_test` with the two runner tests for an absent index and a
relative path; and `cmd/xeno/main_test` for the exit code staircase and the dispatch table
checked against the usage. The WP7 row gains the digest written from the agent's summary and
filtered before it is hashed, with `TestFinishWritesTheDigestFromTheSummary`,
`TestTheDigestIsFilteredAndSaysWhichFilter` and `secrets/secrets_test`.

<!-- xeno:section:deviations -->
## Deviations from the design

**The first draft of the trail paragraph reproduced the error it was correcting.** It was
written from the README's own sentence, kept the claim that the intents stop after P0, and
added that carrying one through all six phases waits on the rule engine and the agent layer.
Both halves were false: twenty intents already run all six phases. It was caught one step
later, when `xeno intent status` was run for the next intent key and printed ten intents
complete at `05-review` green. The paragraph was then replaced from the phase directories,
counted rather than read. Recorded because it is the sharpest evidence for the finding: a
correction written from the stale record inherits its claim, and the only defence is going to
the tree for every sentence.

**The edits were made before the intent was started.** Both files were corrected in the
working tree, and `phase start` for 00-intake came afterwards, so the artifacts here describe
work that already existed rather than guiding work that followed. This is the shape every
intent in this repository has had while the agent layer is unbuilt — the plan's working
sequence assumes a harness driving the phases — but it is a deviation from that sequence and
not a feature of it, and the artifacts are weaker for it: the design section records
alternatives that were weighed against a diff rather than before one.

**The heading change is beyond what the issue enumerates.** #156 lists the stale passages and
does not mention `# xeno, core`. The heading contradicted the paragraph under it after that
paragraph was rewritten, so leaving it would have created a new, smaller version of the same
drift inside the same commit. It is one line, it is named in the design decisions, and it is
flagged here rather than counted as part of what was asked.

**Four over-length lines are left in the register.** Lines 105, 126, 130 and 133 of
`ASSUMPTIONS.md` run past 88 columns and predate this change. Rewrapping them would mean
editing paragraphs that are not stale, and under the replace-rather-than-edit convention the
honest form of that is a rewrite of correct prose, which is a worse trade than a long line.

**No check was built.** The acceptance deliberately contains no guard against recurrence, and
the residual risk says so plainly. That is a decision recorded in the design rather than a
shortfall discovered here, but it is the reason this intent does not close the problem it
describes.
