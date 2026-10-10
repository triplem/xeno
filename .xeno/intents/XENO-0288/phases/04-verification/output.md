---
intent: github.com/triplem/xeno#334
phase: 04-verification
created: "2026-10-10T15:58:18Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8f88e6c3c30dcd672ba8341f312a4023d3794cf11cee0c3d6d07107f431c1fe5
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Fourteen criteria. There is no test suite for a document, so the third column says what
kind of answer each one got: a command whose output is quoted in the results section, a
reading at the pinned commit, or a reading of the tree this change lands in. A reader of a
sealed phase cannot otherwise tell a criterion that was checked from one that was believed.

| criterion | checked by | kind |
|---|---|---|
| 1, the section exists in section 9's form and is section 10 | the heading list of the file, read against section 9's | command |
| 2, the pin is one tag named with the sha, and the sha is the commit the annotated tag points at | `gh api` on both refs and both tag objects, for both repositories | command |
| 3, five rows and the fifth is rules and learning | the table, read | reading of the change |
| 4, every AI-DLC cell re-read at the pinned sha, and the figures that moved say so | the cell-by-cell table below | reading at the pin |
| 5, the third row settles the contradiction and says which document is current | both sentences at their lines, plus the four places in the code | reading at the pin |
| 6, every absence probed with a positive control in the same command, both answers recorded | the three probes below, each with its control | command |
| 7, the decision is "not an extension" with the reason on the page | 10.3 and 10.4, read | reading of the change |
| 8, the conditions for revisiting carry the alternative not taken | 10.5's first bullet, which names `docs/process-definition.md:480` | reading of the change |
| 9, the mechanisms are named and not taken, each pointing at its issue | 10.6, and `grep` for the three issue numbers | command |
| 10, the hosted sample is a line and compares on the same rows | 10.1's third paragraph and the second row's cell | reading at the pin |
| 11, the page's own index names the new section beside section 9 | section 1's paragraph, read back as a paragraph | reading of the change |
| 12, Sources carries what was read | the new entry, against the form of the two existing ones | reading of the change |
| 13, nothing normative moves | `git diff main` over both documents, with a control | command |
| 14, the gates are green and the conventions hold | the four commands `CLAUDE.md` names, plus the width check | command |

**Criterion 4, cell by cell.** Every AI-DLC claim in the section, with where it was read at
commit `6a378b53c0a4fe0641ed7d8de8dfff94264d5b6a`. The tree was taken as a tarball of that
exact commit and read on disk, so each line number below is a line number in that tree and
not in `main`.

| claim in the section | read at |
|---|---|
| an intent is "a single run of the AI-DLC lifecycle, scoped to one task" | `docs/guide/03-spaces-and-intents.md:100` |
| the registry row is `{uuid, slug, dirName, scope, repos, status}`, with no issue field | `docs/guide/03-spaces-and-intents.md:101-104` |
| the record dir is `intents/<YYMMDD>-<label>/`, identity a UUIDv7 in the registry | `docs/guide/03-spaces-and-intents.md:82-85` |
| an intent is auto-created the first time work is described | `docs/guide/03-spaces-and-intents.md:115` |
| five phases hold 33 stages | `docs/guide/04-phases-and-stages.md:3` |
| a jump reopens a stage and every later one; the files stay | `docs/guide/04-phases-and-stages.md:69` |
| eleven scope profiles | the eleven files under `core/scopes/`, counted |
| fourteen agents | the fourteen files under `core/agents/`, counted |
| seven harnesses | the onboarding table of `docs/guide/harnesses/README.md:14-21`, six rows for seven harnesses because Kiro CLI and Kiro IDE share one |
| every stage outside the three initialisation stages ends with an approval gate | `docs/guide/07-interaction-modes.md:72` |
| keep, modify or redo on a reopened stage | `docs/guide/07-interaction-modes.md:287` |
| reviewed outputs remain frozen within an attempt | `docs/reference/04-stage-protocol.md:1084` |
| the reviewer never blocks; the human has final say at the gate | `docs/reference/04-stage-protocol.md:1269` |
| 115 audit event types | `docs/reference/12-state-machine.md:593` |
| `STAGE_JUMPED` is emitted by `tools/aidlc-jump.ts` on a `--stage`/`--phase` jump | `docs/reference/12-state-machine.md:624` |
| quote A, a sensor result is advisory in this release, unqualified | `docs/guide/09-rules-and-the-learning-loop.md:140` |
| the paragraph two lines above it describes gate-fired sensors | `docs/guide/09-rules-and-the-learning-loop.md:138` |
| quote B, blocking is enforced for `fire_on: gate` | `docs/reference/07-sensor-system.md:101` |
| `fire_on` defaults to `write` | `docs/reference/07-sensor-system.md:104` |
| `fireGateSensors` narrows to `fire_on === "gate"` then `default_severity === "blocking"` | `core/tools/aidlc-state.ts:3465`, `:3480`, `:3509` |
| `enforceBlockingGateSensors` returns on an empty list and otherwise calls `error()` | `core/tools/aidlc-state.ts:3725`, `:3734`, `:3767` |
| `gate-start` calls it | `core/tools/aidlc-state.ts:5928` |
| the override is refused in autonomous mode, on a wrong answer, and without a fresh receipt | `core/tools/aidlc-state.ts:3743`, `:3752`, `:3758` |
| the write path exits 0 always | `core/hooks/aidlc-run-sensors.ts:305` |
| all six shipped sensors declare `advisory` | line 5 of each of the six files under `core/sensors/` |
| `SCOPE_PRIORITY` is org 0, team 1, project 2, phase 3, strictly additive, resolved once | `docs/guide/09-rules-and-the-learning-loop.md:41,45`, `docs/reference/08-rule-system.md:111` |
| the diary is surfaced verbatim, no paraphrase and no filtering | `docs/guide/09-rules-and-the-learning-loop.md:70` |
| "Keep none of these" is the first choice; an empty diary asks nothing | `docs/guide/09-rules-and-the-learning-loop.md:73` |
| there is no widen-to-org path | `docs/guide/09-rules-and-the-learning-loop.md:82` |
| each write leaves `RULE_LEARNED` or `SENSOR_PROPOSED`, so no rule is installed silently | `docs/guide/09-rules-and-the-learning-loop.md:86` |
| a learning does not change the rules for the rest of the current workflow | `docs/guide/09-rules-and-the-learning-loop.md:96` |
| the admission check is an LLM check offering revise, skip or escalate with no override | `docs/guide/09-rules-and-the-learning-loop.md:90`, `docs/reference/08-rule-system.md:54` |
| "an audit aid, not a deterministic enforcement boundary", persist not re-running it | `docs/reference/08-rule-system.md:54,111` |
| a plugin's `memory` declarations are rejected | `docs/reference/18-plugin-mechanism.md:140,541` |
| `when:` is parsed and not evaluated | `docs/reference/18-plugin-mechanism.md:539` |
| contributed artifact names are plugin-prefixed | `docs/reference/18-plugin-mechanism.md:547` |
| the receipt carries a `Unit Source Fingerprint` over claimed paths and manifest bytes | `docs/reference/20-commit-provenance.md:11` |
| the committed evidence path, and that it travels with every clone | `docs/reference/20-commit-provenance.md:71` |
| each row carries repo selector, path, file mode and blob OID; the header binds the manifest | `docs/reference/20-commit-provenance.md:85` |
| the five statuses, with `unverifiable` and `indeterminate` failing closed | `docs/reference/20-commit-provenance.md:100-103` |
| `--record-ref` strips a self-written approval | `docs/reference/20-commit-provenance.md:31,148` |
| attribution "keys on blob content (OIDs), not commit ancestry" | `docs/reference/20-commit-provenance.md:138` |
| the release is checksummed with a Sigstore bundle and build provenance | `docs/reference/19-supply-chain-security.md:146-151,231` |
| MIT-0, created 2025-11-13, 5,123 stars, 924 forks | the GitHub API, 2026-10-10 |
| `v2.11.0` is annotated; tag object `4079edbe`, commit `6a378b53c`, tagged 2026-10-08 | the GitHub API, 2026-10-10 |
| three preview tags around it, numbered `2.11.1-preview` | the releases list, GitHub API, 2026-10-10 |
| the sample is an early-preview AWS sample deployed into the adopter's own account | `README.md:33` at `bc988d0e` |
| it can start an intent from a GitHub, GitLab or Jira issue | `README.md:128,345` at `bc988d0e` |
| its record lives in Neptune and DynamoDB | `README.md:122,247` at `bc988d0e` |
| the sample is MIT-0, and `v2.2.0` is annotated at commit `bc988d0e` | the GitHub API, 2026-10-10 |

**Three figures from #323 moved, and the section carries the current ones.** 117 audit
event types became 115, which is what `docs/reference/12-state-machine.md:593` says at the
pin and what the drift test `tests/integration/t48-audit-event-emitters.test.ts` named there
holds. 5,084 stars became 5,123, read on 2026-10-10 and dated in the section rather than
stated as a fact about the project. And #323's third row — "sensors that may be advisory, a
reviewer that never blocks, a person at every stage who may override" — is replaced by the
settled account, because "may be advisory" was the contradiction rather than a reading of
it.

**One claim from #323 did not reproduce and is not in the section.** "later agents may
revise earlier artifacts if downstream stages detect gaps" is not a sentence at the pin. A
search of the whole tree for `revise` and `revising` returns the learnings admission check,
the gate's `report --result revised` within one attempt, and `STAGE_JUMPED`; nothing says a
later agent revises an earlier artifact on its own initiative. What the tree does say is
`docs/reference/04-stage-protocol.md:1084`, "Reviewed outputs remain frozen", and
`docs/guide/07-interaction-modes.md:287`, where a person jumps back and chooses keep, modify
or redo. So the fourth row is written from those two and says "a person's jump back", which
is a weaker and checkable claim, and the fourth column says "with a person at every turn"
rather than #323's "with a gate on it".

<!-- xeno:section:results -->
## Results

Twelve of the fourteen criteria are met, one is met in its operative clause and not in the
illustrative one beside it, and one cannot be answered by this tree at all. Every command
below was run in this worktree at `295bcc4`, the commit that carries the documents change.

**The four gates `CLAUDE.md` names.**

    gofmt -l . | grep -v '^vendor/'   → 0 lines
    go vet ./...                      → exit 0
    go test ./...                     → exit 0, twice
    ./xeno gate verify                → exit 0, "verified 607 verdicts"

`go test ./...` was run twice because the first run's exit code was lost when its shell
ended before the status line was appended, and a log of twenty `ok` lines with no `FAIL` in
it is not an exit code. Both runs took about ninety seconds and both exited 0. The `gofmt`
line is the one that needs saying carefully: the pipeline's own exit status is 1, because
`grep` found nothing, which is indistinguishable from a failed suite if the status is read
instead of the output. What the criterion rests on is that `gofmt -l .` outside `vendor/`
printed zero lines.

**Criterion 1, the form.** Section 10 has six subsections, as section 9 has six, and the
pairs line up: what was evaluated, the rows, the decision, the measurement, the conditions
for revisiting, the mechanisms. Two headings differ from 9's in words rather than in place —
"The five rows" against "The four rows", and "The three shapes, and what each costs" against
"What the measurement said" — and both differences are the thing the subsection is about.

**Criterion 2, the pins.** `gh api` on `repos/awslabs/aidlc-workflows/git/ref/tags/v2.11.0`
returns an object of type `tag` with sha `4079edbea52cf4d3c80e580825da3a4dd8213271`, and
that tag object dereferences to commit `6a378b53c0a4fe0641ed7d8de8dfff94264d5b6a` with a
tagger date of 2026-10-08T08:19:42Z against a commit date of 07:39:40Z. The same holds for
the sample: ref to tag object `9ecc0d2327c0c7252d7f58e4488f73e4926f6388`, which dereferences
to commit `bc988d0eaaca752d9add2d67b2c861841a363227`. So both pins are annotated, both shas
in the section are the commits, and the section says so.

**Criterion 6, the three negative claims, each with its control in the same command.** This
is the criterion this project cares about most, so all three answers are here rather than
summarised.

*Nothing in AI-DLC's core binds an intent to a tracker issue.* A case-insensitive search of
`core/` at the pinned commit for `jira|issue_url|issueKey|tracker|github.com/.../issues`
returns exactly one file, `core/knowledge/aidlc-shared/worktree-info-schema.md`, and the two
hits in it are prose citations of AI-DLC's own issues 1281 and 1252 in a sentence about
untrusted audit shards. The control, a search of the same tree in the same shape for
`intents.json`, returns five files including `core/tools/aidlc-utility.ts` and
`core/tools/aidlc-lib.ts`. So the search reaches `core/`, and what it does not find there is
a tracker.

*Nothing but the reviewed-source fingerprint hashes AI-DLC's record.* A search of `core/`
for `artifacts_hash|record hash|seals the record` returns nothing. The control, a search of
the same tree for `Unit Source Fingerprint`, returns `aidlc-swarm-checkpoints.ts`,
`aidlc-swarm.ts` and `aidlc-log.ts`. The absence is therefore an absence, and the one thing
that is hashed is the thing 10.6 describes.

*No file under `vendor/` or elsewhere in this repository is misformatted.* `gofmt -l .`
printed nothing; the same binary run against a deliberately misformatted file printed that
file's path. An empty answer from a formatter is worth nothing without that second half.

*Nothing normative moved.* `git diff main -- docs/process-definition.md
docs/implementation-plan.md` produced a zero-line diff; `git diff --numstat main --
docs/orchestrator-evaluation.md` in the same breath produced `334 9`. So the diff command
was pointed at a tree where it can see a change, and it saw none in those two files.

**Criterion 4, the re-reading.** The cell-by-cell table in the mapping section is the
answer, and its figure is that every AI-DLC claim in the section has an address at the
pinned commit. The reading was done against a tarball of that exact commit extracted to
disk, so the line numbers are line numbers in the pinned tree. Three figures from #323 moved
and the section carries the current ones; one claim from #323 did not reproduce and is not in
the section, with what replaced it recorded in the mapping.

**Criterion 9, met in its operative clause and not in the illustrative one.** The operative
requirement is that no mechanism is taken and that a mechanism with an issue points at it:
section 10 names four mechanisms, takes none, and `#340` appears in 10.6 beside the sensor
that fires on write. The criterion's second sentence then said the section points at #338,
#339 and #340, and it does not: a search of section 10 for issue references returns `#323`,
`#334`, `#337` and `#340`. #338 is the trigger-label mechanism, which came from the
OpenHands demo of 4.1 and not from this reading, so naming it in section 10 would be wrong.
#339 is AI-DLC's reviewer agent, which is an AI-DLC mechanism and is genuinely missing from
10.6. That is recorded as a gap rather than repaired, for the reason the gaps section gives.

**Criterion 14's second half, the conventions.** Every non-table line of the file is at most
88 characters, counted as characters and not as bytes. The distinction is not pedantic here:
the first pass of this prose was written at 90 and a byte-counting check reported the same
compliant lines as over-wide, because an em dash is three bytes, so the check cannot tell a
wide line from a typographic one in either direction. The commit message is at most 72
characters by the same count, and `./xeno check commit-message` exits 0.

**What this tree cannot answer.** Nothing in criterion 3, 7, 8, 10, 11 or 12 is a property a
command can decide: whether the fifth row is the right fifth row, whether the reason in 10.3
decides, whether the replaced paragraph reads as a paragraph. Those were read rather than
tested, and the third column of the mapping says so for each. A document's acceptance
criteria are answerable by a reader and not by a suite, and the honest form of that is to
say which is which rather than to invent a check.

<!-- xeno:section:gaps -->
## Gaps

**10.6 does not name AI-DLC's reviewer agent, and #339 holds it.** The reviewer is a
sub-agent the conductor invokes after a stage body produces its artifacts and before the
approval gate, and it "never blocks — the human always has final say at the gate"
(`docs/reference/04-stage-protocol.md:1055` and `:1269`). That is an AI-DLC mechanism of
exactly the kind 10.6 is for, #323 kept it in view, and #339 was opened for it. It is not in
the section.

It is recorded rather than repaired because repairing it now would cost more than it is
worth. The document change is committed at `295bcc4` and P3's verdict seals the phase that
describes it line for line; editing the document after P3 is decided makes P3 stale by
G-Freshness, and the repair for that is a redo of P3, which this phase's own
`context.lock.yaml` already points at. P4 is running, so it cannot be started again to pick
up a new P3. What is lost by leaving it is one pointer: a reader who follows #323 from 10.1
reaches #339 in the comment that opened it, so the mechanism is findable and the section is
incomplete rather than wrong. What would close it is one paragraph in 10.6, which belongs to
whoever next touches that section, and #339 is where to say so.

**The hosted sample's own five rows have not been read by anybody.** Not by #323, not here.
The line 10.1 gives it is held to its `README.md` at `bc988d0e` and makes no claim about its
gates, its sensors or its approvals. Its human gates, its typed traceability graph and its
per-stage cost figures are each named in that README as things it has, and each would be a
cell in a row if somebody read them. Whether it earns a reading of its own is a question for
the next person who wants the hosted comparison; this section claims only what one file
states.

**The section is a claim about one commit and AI-DLC moves faster than that.** Its default
branch moved on the day of this reading, and three preview tags were cut in the three days
around `v2.11.0`. So the section is already not a description of `main`, which is the form
section 9 chose and #323 argued for. What has no repair is the part where AI-DLC fixes its
own contradiction: the day the rules guide's sentence at `:140` gains the gate-fired
qualification, 10.2's third row is still correct about the pin and reads as an argument about
a thing that is no longer there. Nothing in this project notices that, and nothing could.

**Nothing re-checks the width of this page's prose.** The first pass of section 10 was
written at 90 characters and reached a commit-ready state only because the width was
measured again by hand. `gofmt` says nothing about Markdown, no test reads the documents for
width, and `./xeno gate verify` judges the artifacts of the trail rather than the files the
trail is about. The next long document written here will make the same mistake in the same
way, and the learning recorded in P3 is the general form of it.

**#334 carries no work-package label and no issue holds that.** P0's learning records the
finding and names `docs/implementation-plan.md` as what a merge request would change. No
issue exists for it, so it lives in this trail only, which is the state the third standing
rule says to write down rather than absorb — written down here, and in P0, and nowhere a
reader of the plan would find it.
