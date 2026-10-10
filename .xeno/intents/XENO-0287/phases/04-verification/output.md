---
intent: github.com/triplem/xeno#336
phase: 04-verification
created: "2026-10-10T13:42:55Z"
schema_version: "1.0"
runner_version: dev+5f645cb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 063fafecabb56f805283544ba7d8269ceb50c371da8fe31c73efc691cca1a8fd
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

Eleven criteria. Six are answered by reading the paragraph that was written, four by a
command, and one by CI. The third column says which, because the six read ones are the
substance and a reader of a sealed phase should know they were judged rather than measured.

| criterion | checked by | answered |
|---|---|---|
| 1, not a candidate | the paragraph's first clause, read | here, by reading |
| 2, points at section 12 and says it is the refused shape | the paragraph's middle, read against `docs/process-definition.md:1698` | here, by reading |
| 3, strengthens 4.1 rather than reserving against it | the paragraph's "evidence for that paragraph rather than against it", read | here, by reading |
| 4, the #330 contrast is accurate | the second sentence read against `Issue.Approval()` in `internal/model/identity.go` | here, by reading |
| 5, Sources carries a pinned commit in the section's own form | the new entry read beside the OpenSpec entry above it | here, by reading |
| 6, every claim was read at that commit | the intake's record of what was read at `2ddf6c9`, against what the paragraph and the entry assert | here, by reading |
| 7, section 12 does not move | `git diff main -- docs/process-definition.md` | command |
| 8, no other section of the evaluation moves | `git diff -U0 main -- docs/orchestrator-evaluation.md`, hunk headers | command |
| 9, one sentence not a section | the paragraph, read — **not met, see the deviation** | here, by reading |
| 10, prose wraps and the page builds | `awk 'length>88'`; the build is the docs workflow | command, then CI |
| 11, the gate suite | `gofmt -l`, `go vet`, `go test -count=1 ./...`, `xeno gate verify` | here, then CI |

**Why six criteria are judged and not measured, and what that is worth.** This intent's
deliverable is prose, and no command can answer whether a sentence says a thing. What makes
the six more than an opinion is that each names the clause that carries it: criterion 1 is
"is this choice running rather than an alternative to it", criterion 3 is "evidence for that
paragraph rather than against it", and criterion 4 is "reads off an issue it has already
fetched". A criterion pointing at a clause can be disagreed with by a reviewer reading the
same clause, which is the most a prose criterion can offer, and it is why they were written
as claims about what the sentence must assert rather than as impressions of it.

**Criterion 4 is the one with a wrong answer available, and it was checked against the code
rather than against the issue.** `Issue.Approval()` is a method on a value the runner already
holds: it iterates `i.Labels` and `i.Comments` and returns what is missing. It performs no
fetch and nothing subscribes to it. So "reads off an issue it has already fetched" is exact,
and the two wordings the criterion forbids — calling #330 a trigger, or calling the demo's
labels preconditions — are both avoided.

**Criterion 9 is not met and the deviation says so.** It is listed in the table as not met
rather than quietly re-read as satisfied, because a mapping that marks its own unmet
criterion as met is worth less than no mapping.

<!-- xeno:section:results -->
## Results

Ten of eleven met. One not met, criterion 9, by the deviation P3 records.

**The four commands.**

| check | result |
|---|---|
| `git diff main -- docs/process-definition.md` | **0 lines** |
| `git diff -U0 main -- docs/orchestrator-evaluation.md`, hunk headers | **two**, `@@ -96,0 +97,10 @@` in 4.1 and `@@ -478,0 +489,13 @@` in Sources |
| `awk 'length>88'` over the file, excluding tables | **nothing** |
| `gofmt -l .` outside `vendor/`, `go vet ./...`, `go test -count=1 ./...`, `xeno gate verify` | clean, clean, **exit 0 over 20 packages**, **exit 0 over 600 verdicts** |

The second row is criterion 8 answered exactly: two hunks, one in each of the two places the
design named, and no third hunk anywhere. The fourth row's test suite was run to its own exit
code rather than read off a filtered pipeline, because a `grep` that finds no failures exits
1 and that is indistinguishable from a suite that failed if only the pipeline's status is
read.

**The six judged criteria, each against the clause that carries it.**

Criterion 1 is met by "is this choice running rather than an alternative to it", which is the
paragraph's first clause and therefore the frame everything after it is read in. Criterion 2
is met by naming section 12 and quoting the paragraph's title, and the quotation was compared
word for word with `docs/process-definition.md:1698`. Criterion 3 is met by "and it is
evidence for that paragraph rather than against it", which is the issue's "strengthens the
4.1 decision" in the form the sentence can actually assert.

Criterion 4 is met and was checked against the code. `Issue.Approval()` iterates `i.Labels`
and `i.Comments` on a value the runner already holds and returns what is missing; it fetches
nothing and nothing subscribes. "Reads off an issue it has already fetched" is therefore
exact rather than approximate, and neither forbidden wording appears.

Criterion 5 is met: the entry opens on what it was read for and the date, as the OpenSpec
entry above it does, gives the repository and the reference, names each path by what it was
read for, and states the facts about the repository. It also answers the one question its form
raises, why a commit and not a tag.

Criterion 6 is met. Every claim the paragraph and the entry make is in the intake's record of
the reading at `2ddf6c91791d95b8de8ff38b683241743ec12049`, and the claims the reading does not
support are absent: nothing is said about the demo's quality, its adoption beyond the three
facts in the entry, or how well the design works for the people using it.

**Criterion 10's second half is CI's.** `zensical` is installed by `docs.yml` and not on this
machine, so that the page builds is the docs workflow's answer. The wrap is checked here.

**Criterion 9, not met.** The paragraph is two sentences where the criterion says one. P3's
deviations section carries the reason and the measurement behind it — one sentence came to
about ninety words with three nested subordinate clauses, with the #330 contrast last. What
the issue's "one sentence" was for is intact: it is not a section, not a table and not a
list, which is the criterion's own second half.

<!-- xeno:section:gaps -->
## Gaps

**Nothing holds the Sources entry to the commit it pins.** The entry says what was read at
`2ddf6c91791d95b8de8ff38b683241743ec12049` and the repository was pushed to on 2026-10-09,
two days after that commit, by one author. If the four labels are renamed or the automations
restructured tomorrow, the entry stays true — it is a record of a reading — but the
paragraph's present tense drifts: "starts each run from a label on an issue" is a claim about
a design, and a design can change. No test in this repository can check a claim about
somebody else's tree, and `internal/model/supply_chain_test.go` is the wrong instrument,
since this is a citation and not a pin the pipeline fetches. What makes it tolerable is that
the entry is dated and names the commit, so a reader who finds the repository changed knows
exactly what was read and when. Written down rather than absorbed.

**The six judged criteria have no second reader until the pull request.** A criterion about
what a sentence asserts is verified by somebody reading the sentence, and in this phase that
somebody is the same agent that wrote it. Naming the clause for each is what makes a
reviewer's disagreement possible rather than a matter of taste, and the review in P5 and the
pull request are where that second reading happens. This is not a defect of this intent so
much as the shape of verifying prose, and it is the reason criterion 9's failure is recorded
in the table rather than reasoned away: a self-verified mapping that never reports an unmet
criterion is not a mapping.

**The absence of `.github/workflows` is a negative result about another repository, and the
control is weaker than the rule wants.** The project rule says to recreate the thing and
check again, which cannot be done for somebody else's tree. What was done instead: the full
recursive tree listing at the pinned sha was read, and it returned 89 blobs including
`.github/ISSUE_TEMPLATE/config.yml`, `.github/labels.json` and
`.github/pull_request_template.md`, so the probe demonstrably reaches inside `.github/`. That
is a positive control in the same call, which is the strongest form available here, and it is
the same arrangement XENO-0286's learning proposed for a registry. It is not the rule's own
remedy, and saying so is the point of this paragraph.

**Criterion 10's build half is unverified locally and will stay that way.** `zensical` is
installed by `docs.yml` and pinned there, and installing it on this machine to check one
paragraph would be a dependency nobody asked for. The docs workflow answers it on the pull
request, which is where the answer belongs; the risk in between is a Markdown construction
that wraps correctly and renders wrongly, and the paragraph uses nothing but inline code,
quotation marks and an em dash.
