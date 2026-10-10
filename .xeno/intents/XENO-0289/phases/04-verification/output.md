---
intent: github.com/triplem/xeno#359
phase: 04-verification
created: "2026-10-10T16:00:11Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0b4b5e12d3b9c8cd93279d8fa488f3239b78aaaa19af5ced1c438a4916770fae
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

Twelve criteria from `01-requirements`, and no test in the Go suite covers any of them:
nothing in this intent is code, and the design argued against a test that would hold the
page against the host because it would put a network call in a repository whose gate path
makes none. So each criterion maps to a command run in this phase, and the commands are
given with their output in the results section rather than named here.

| Criterion | How it is answered | Result |
|---|---|---|
| 1. Every host label on the page, and none that is not | the host's label list read with `gh api`, compared against the page in both directions | pass; the count the criterion states is wrong and is reported below |
| 2. Who sets, who clears, what reads, for every label | the three tables read column by column | pass |
| 3. What the runner reads, stated exactly | a search for a label comparison over `internal` and `cmd`, with `ApprovedLabel` as the positive control | pass |
| 4. Nothing in Xeno writes a label | the three operations in `internal/host/host.go`, and every HTTP method both adapters use, enumerated | pass, and on stronger evidence than the criterion asked for |
| 5. The page defines nothing | `git diff main --stat` against both normative documents, with a file this intent did change as the control | pass |
| 6. Which half of #359, and no label that does not exist | a search for `xeno-start` on the page, with a present string as the control | pass |
| 7. Reachable from the index and from the site's navigation | `zensical build --clean --strict`, the built `index.html` read for the link, and a deliberate broken link as the positive control | pass |
| 8. Prose at 88, tables and code blocks exempt | a character-count pass over both changed markdown files | pass |
| 9. The question is on the issue, labelled, and read back | the issue's labels and comments read from the host after the write | pass |
| 10. The operational note records the probe, not the report | the probe on #359 and what the page says about it | pass; the carried-in claim did not reproduce |
| 11. One intent, one branch, one issue | the branch name, the commits and the issue's labels | pass, with the work package label absent as `00-intake` recorded |
| 12. Gates green and the tree clean | `gofmt`, `go vet`, `go test ./...`, `xeno gate verify` | pass |

Criterion 2 and criterion 11 are the two that no command can answer, and they are read
rather than measured. Said here so that twelve rows of "pass" are not read as twelve
mechanical checks.

<!-- xeno:section:results -->
## Results

Every check below was run in this phase against the tree as it stands. Each negative
result carries a positive control in the same command, because a tool asked about
something absent answers as it does about something that does not match, and this page
makes four claims of the form "nothing reads this".

**Criterion 1, the label set, and a wrong figure in a sealed artifact.** The host was read
with `gh api repos/triplem/xeno/labels --paginate --jq '.[].name'` and answered 33 names.
Criterion 1 says 34. The figure was written into `01-requirements` from this session's own
earlier reading rather than from a command, and that artifact is sealed, so it stays
wrong; this is where a reader of it finds the correction, and the learning record of
`03-implementation` is where the cause is. What the criterion actually asks for is the set
comparison, and it passes in both directions: no host label is absent from the page, the
twenty-one `wp` labels are contiguous from 0 to 20 so the page's `wp0` to `wp20` covers
them exactly, and the only hyphenated names in backticks on the page that are not host
labels are `open-questions`, which is a section id, and `xeno-need-decision`, which
occurs once, in the sentence that says it is not the label's name.

**The second wrong figure, in `00-intake`.** Its decision record names #329, #334, #338 and
#352 as carrying `xeno-needs-decision`. The host says seven issues carry it — #120, #224,
#234, #329, #334, #338 and #359 — and #352 carries `xeno-approved` only. Same cause: a
list written from the session rather than read. The page carries the host's seven.

**Criterion 3, what the runner reads.** A search for a label comparison across `internal`
and `cmd` returns one line, `internal/model/identity.go:124`, which compares each entry of
`Issue.Labels` against `ApprovedLabel` with case ignored and whitespace trimmed. The
positive control is the same search for `ApprovedLabel` itself, which returns four sites
in that file and one in `internal/runner/init.go`, where `xeno init` names the label among
the settings it does not make. A search for a `wp` label literal across `internal` and
`cmd` returns three test fixtures and nothing else: `wp12` sits in a label list that an
adapter carries into `Issue.Labels` and that the comparison passes over, which is the
positive demonstration that a work package label reaches the code and is ignored.

**Criterion 4, nothing writes a label, on better evidence than a search.** Every HTTP
method both adapters use was enumerated: `internal/host/github/tracker.go` makes three GET
requests and one POST, `internal/host/gitlab/tracker.go` the same, and the one POST in
each goes to the comments endpoint — `repos/%s/issues/%s/comments` on GitHub. So the claim
does not rest on a grep for the word label returning nothing; it rests on there being
exactly one write in the adapter layer and on its address. `internal/host/host.go` is why
that is a property of the contract rather than of one host: `Host` is `BranchRules`,
`Issues` and `Comments`, and nothing else.

**Criterion 5, the page defines nothing.** `git diff main --stat -- docs/process-definition.md
docs/implementation-plan.md` prints nothing. The control is the same command on
`docs/README.md`, which prints one file and five insertions, so the empty output is the
absence of a change and not a command that failed to look.

**Criterion 6, no label that does not exist.** `xeno-start` does not occur on the page.
The control is `xeno-approved`, which occurs five times.

**Criterion 7, reachable, with a positive control for the strict build.** `zensical
0.0.69` was installed into a throwaway virtual environment, since it is not part of this
repository's toolchain, and `zensical build --clean --strict` reported "No issues found".
The built `index.html` carries `href="./labels/"` in its navigation and `href="labels/"`
from the index page's own list, and `.xeno/local/site/labels/index.html` exists. The
control: a link to `no-such-page.md` was appended to the page, the build aborted with
`RuntimeError: Aborted because --strict flag is set` and exit 1, and the page was restored
and rebuilt clean. So the green build is evidence that the index link resolves, and not
merely that the build ran.

**Criterion 8, widths.** No line of `docs/labels.md` or `docs/README.md` outside a table
row or an indented code block exceeds 88 characters, and no line leaves a code span open.
Two lines were over by one or two characters after the first reflow and one markdown link
and one code span had been split across a line break; all four were repaired by rewording
rather than by leaving the break, since a code span broken across lines renders as a space
and reads as a different string in the source.

**Criterion 9, the question and the label.** `repos/triplem/xeno/issues/359/labels` reads
back `xeno-approved,xeno-needs-decision`. The issue carries two comments: the approval at
14:36:07Z and this intent's question at 15:36:24Z. The question carries three options,
each with its consequence and its cost, a recommendation with its reason, and a free
entry, and it says that the second of #359's questions follows from it and is asked after.

**Criterion 10, the probe that did not reproduce.** `gh issue edit 359 --add-label
xeno-needs-decision` was run with the label present on the repository and absent from the
issue, the two conditions that make a failure visible rather than ambiguous. It printed
the issue URL, exited 0, and the read-back showed the label applied. `gh` is 2.101.0. The
carried-in observation was that this form exits 0 without applying; it did not reproduce,
and the page therefore prescribes the read-back on its own merits rather than recording a
defect. The `gh api` form was then run against the same issue and is idempotent: it
returned the full label list and the read-back agreed.

**Criterion 12, the gates.** `gofmt -l .` outside `vendor/` prints nothing. `go vet ./...`
exits 0. `go test ./...` exits 0 across twenty packages, run twice, the second time in the
foreground so that the exit code was read rather than inferred from the absence of the
word FAIL. `xeno gate verify --intent XENO-0289` verified four verdicts and exited 0 at
that point in the run, and the per-phase verdicts are green for 00 to 03 with this phase
judged at its finish.

**What this phase does not evidence.** Criterion 2 is a reading of three tables and
criterion 11 is a reading of the branch and the commits; neither has a command behind it,
which is said in the mapping so that twelve passes are not taken for twelve measurements.

<!-- xeno:section:gaps -->
## Gaps

**Two figures in sealed artifacts are wrong and stay wrong.** The label count in
`01-requirements` criterion 1 and the list of labelled issues in `00-intake`'s decision
record. Both are corrected in the results section above and the cause is the learning
record of `03-implementation`; neither artifact is rewritten, because a verdict covers the
content it was computed over. A reader who arrives at either phase and not at this one
reads the wrong figure, which is the cost the project's own rule about sealed findings
names, and it is why the rule the learning record proposes is about reading a figure before
sealing it rather than about correcting one afterwards.

**Nothing keeps the page and the tracker in agreement.** The page records the set as read
on 2026-10-10 and names the command that read it. A label created tomorrow will not appear
on it, and nothing will say so. The design argued against the two mechanisms that would
fix it — a test calling the host, which the gate path forbids, and a generator, which makes
the page a build product — and the plan's own treatment of the `wp` labels is the
precedent. This is a known and accepted gap, not an oversight.

**The claim that nineteen issues carry `xeno-approved` is a count from one command.** It
was read with `gh issue list --state all --limit 500 --json labels`, which covers issues
and not pull requests, and the repository's highest issue number is 359, so the limit does
not truncate. The inference drawn from it on the page — that the issues approved before
#332 carry a label that no longer exists — rests on A107 and on `approved` being absent
from the host's 33 labels, not on the count alone. What is not evidenced is that every one
of those older issues was approved at all; the page says the figure shows what the label is
and is not, and does not say more.

**The strict build was verified with a generator this repository does not pin locally.**
`zensical 0.0.69` matches the version `.github/workflows/docs.yml` installs, and its
fourteen dependencies resolved unpinned at install, which `docs/supply-chain.md` says of
the CI install too. So the local build is the same check as CI's and not the same
environment, and the `docs` job on the pull request is the run that counts.

**The question on #359 is unanswered, and this intent ends with it open.** That is the
arrangement the scope section chose and the reason it chose it: the question is about a
section 12 clause only the maintainer can write, it is on the issue with the
`xeno-needs-decision` label beside it, and it is deliberately not an `open-questions`
entry of any phase. `G-Questions` therefore has nothing to resolve and reports green,
which is correct about this intent and silent about the question. Nothing in the artifacts
distinguishes a question deferred to an issue from no question at all, which is the
learning record of `01-requirements`.

**A redo sits between the first verdict of `03-implementation` and this phase.** The
maintainer's answer on the label's name arrived while this phase was first running, so the
page was written again, D-2 was recorded, `03-implementation` was finished again, and this
phase was started afresh after its unsealed and uncommitted directory was removed. The
trail shows one verdict for `03-implementation`, the one that covers the page as it now
stands; what it does not show is the earlier green, because the finish that replaced it
replaced the artifact with it. The deviations section of that phase is where that is
written down, and this is the second place, because a reader of this phase is who would
otherwise wonder why the page and the first verdict could not both be true.
