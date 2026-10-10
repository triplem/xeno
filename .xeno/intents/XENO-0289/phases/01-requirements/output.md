---
intent: github.com/triplem/xeno#359
phase: 01-requirements
created: "2026-10-10T15:37:53Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a86b6eff5f1d5a3919b922fb0af90d1dc6e363ff3edc8eb9c48d721d9047cbd1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. **The page names every label the tracker carries, and none it does not.** The set is
   read from the host rather than remembered: `gh api repos/triplem/xeno/labels
   --paginate --jq '.[].name'` lists 34 names today, and each one appears on the page,
   grouped by what it is for. Checked in both directions, because a page that invents a
   label is worse than one that omits it: a reader can look for a missing label, and
   cannot know that a listed one is fiction.

2. **Every label on the page says who sets it, who clears it, and what reads it.** Those
   three are what #359 asks for in the words "who acts in a certain label and when a
   label is removed", and no entry is complete without all three. Where the answer is
   nobody — a label no code reads, a label nothing removes — the page says nobody rather
   than leaving the cell out, because an empty cell reads as unknown and the point of the
   page is that these are known.

3. **What the runner reads is stated exactly and is right.** One label, `xeno-approved`,
   compared in one place: `Issue.Approval()` in `internal/model/identity.go`, against the
   constant `ApprovedLabel`, case-insensitively and trimmed, and only alongside a comment
   whose first line is `/xeno approved`. Checked by a grep for a label comparison over
   the whole of `internal` and `cmd` returning that one site, run with a positive control
   in the same command so that an empty result is evidence rather than a tool answering
   about something absent.

4. **The page's claim that nothing in Xeno ever writes a label is grounded in the
   contract, not in a search.** `internal/host/host.go` fixes the adapter at three
   operations — read branch rules, read an issue, write a comment — so a label write has
   nowhere to live. Checked by the interface, and corroborated by a search for a label
   write across `internal/host` returning nothing while the same search shape finds the
   one POST both adapters do.

5. **The page defines nothing.** It adds no label the runner reads, no field, no gate and
   no rule, and it does not restate section 12's clause in its own words where it can
   cite it. Checked by `git diff main -- docs/process-definition.md
   docs/implementation-plan.md` being empty at the end of the intent, and by the page
   containing no sentence of the form that a tool shall read some label.

6. **The page says which half of #359 it is, and names no label that does not exist.**
   The title asks about labels and about trigger labels; the page does the first, says in
   its own words that triggering is #338's open question and why, and does not list
   `xeno-start` or any other proposed label among the labels in use. Checked by the name
   `xeno-start` appearing on the page only inside the sentence that says it does not
   exist, if at all.

7. **The page is reachable from both places a document here is found.** An entry in
   `docs/README.md`, which is the index a reader of the repository follows, and an entry
   in `zensical.toml`'s `nav`, which is the published site's navigation. Checked by the
   `docs` job: it builds with `--strict`, so the index link is verified mechanically, and
   by reading the built navigation rather than assuming the toml was parsed as intended.

8. **The page follows the repository's own conventions.** Prose wraps at 88 characters,
   tables and code blocks do not, headings name their section in words with no leading
   number. Checked by a line-length pass that excludes fenced blocks and table rows.

9. **The question is on the issue and labelled, and the state is read back from the
   host.** One question, with its options, each option's consequence and its cost, a
   recommendation with its reason, and a free entry; `xeno-needs-decision` on #359.
   Checked by `gh api repos/triplem/xeno/issues/359/labels` after the write and not by
   the exit code of the write, which is the practice the page itself records.

10. **The operational note on applying a label records the probe and not the report.**
    The claim that `gh issue edit --add-label` can exit 0 without applying the label was
    carried into this intent from an earlier session. It was probed under the conditions
    that make a negative result evidence — the label present on the repository, absent
    from the issue beforehand — and it did not reproduce. The page says what the probe
    showed and prescribes the read-back on the ground that an exit code is not evidence
    of a label, rather than recording a defect that one session saw and the next could
    not.

11. **One intent, one branch, one issue.** `359-labels-and-the-states-they-stand-for`,
    `Closes #359` on the commit that finishes the work and in the pull request
    description. The issue carries no work package label, WP16 is the nearest fit and
    nothing in the plan owns this repository's own tracker conventions; that gap is
    P0's learning record and is not absorbed into the page.

12. **The gates are green and the tree is clean.** `gofmt -l .` outside `vendor/` prints
    nothing, `go vet ./...` and `go test ./...` pass, `./xeno gate verify` exits 0, and
    the six phases each hold a verdict that covers their own `artifacts_hash`.

<!-- xeno:section:non-goals -->
## Non goals

**It does not answer what becomes of `xeno-approved` after a decision against the work.**
That is the first of #359's two questions and it is on the issue, with its options, their
consequences and their costs, and a recommendation with its reason. The page records what
the label means today — read once by `xeno intent start` as a precondition, never read
again, never removed by anything — and says that what it should mean afterwards is open
and where the question is. The wording for section 12 is drafted after the answer, for
the maintainer's own commit, and not before it.

**It does not answer whether a fresh approval is owed after a decision.** #359's second
question, which follows from the first: every option on the first changes what the second
is asking. It is asked separately once the first is answered.

**It does not rename `xeno-needs-decision`.** The approving comment spells it
`xeno-need-decision`; the label created and applied is `xeno-needs-decision` and five
issues carry it. The page records the applied spelling, names the discrepancy, and leaves
the rename as a question to be asked rather than taken — asked after the one already on
the issue, because two at once is the fault `CLAUDE.md` records XENO-0243 for.

**It does not answer #338.** Whether a label per trigger or a keyword in the title starts
work is that issue's question; the maintainer has chosen one trigger label and no analysis
trigger, and that choice is a section 12 change nobody has made. No such label exists on
the tracker, and the page lists what exists.

**It does not create, rename or remove any label on the host.** One label is applied,
`xeno-needs-decision` on #359, which is the device for the question this intent put. The
work package label #359 does not carry is the maintainer's act and stays theirs; the gap
is recorded instead.

**It does not change `.xeno/plugin/bin/xeno-labels.sh`.** The script's comment claims it
creates "the one label Xeno asks a project's tracker for", and that claim is still true:
`xeno-approved` is the one label the runner reads, and `xeno-needs-decision` is a
convention of this repository that no code reads and no adopter needs. The page says so
in as many words, which is what keeps the script's claim honest without editing it. Had
the second label been something the runner read, the script would have been the place for
it and this would have been a specification change first.

**It does not become a guide for adopters.** `wp0` to `wp20` are this project's own
construction plan and mean nothing in a repository that merely uses Xeno, and
`xeno-needs-decision` is a habit of this tracker. The page separates the one label Xeno
asks for from the ones this repository has given itself, so that a reader can tell which
of the two they are reading.

**It does not add a row to `docs/assumptions.md`.** That page holds a fact that outlives
the intent that found it; the labels in use are a state of the tracker, and the page that
records them is the place they belong.

**It adds no gate, no rule and no test over the labels.** Nothing mechanically checks
that an issue carries the label its state implies, and nothing checks that a label is
cleared when it should be. The page says that plainly rather than implying a check: the
labels are a convention people keep, and one of them — `xeno-approved` — is a
precondition the runner refuses without, which is a different thing from being enforced.

<!-- xeno:section:constraints -->
## Constraints

**The two normative documents are not editable by the agent, and this intent needs
nothing from them that it cannot cite.** `docs/process-definition.md` section 12 fixes
what approval is; `docs/implementation-plan.md` fixes what a work package label is. The
page quotes neither at length and restates neither: it says what the label does and names
the section, so that a clause changing does not leave a second copy of it saying the old
thing. This is the binding constraint on the whole intent and it is what decides the shape
of the page.

**No invented fields, gates, tools or rules.** A page that said a label meant something to
the runner would be adding to the specification by describing it. The page is written in
the past and present tense about what is in use — this label exists, this person sets it,
this code reads it — and in the conditional only where it names an open question and
points at the issue it is on.

**A label is a string compared by whoever compares it.** The one comparison in the code is
case-insensitive and trims whitespace, and that is a property of one function rather than
of labels in general: a filter in a saved query, a person's eye, and a shell script
grepping a list all compare it differently. So the page gives each label's exact spelling
and says where the one comparison happens, and the spelling discrepancy on
`xeno-needs-decision` is named rather than smoothed over.

**The site's navigation is hand-written and the build is strict.** `zensical.toml` carries
an explicit `nav`, so a page not named there is unreachable from the published site even
though it builds. `.github/workflows/docs.yml` builds with `--strict`, so a link to a page
or anchor that does not exist is a red check before a merge. The two together mean the
index entry is verified mechanically and the nav entry is not, which is why the nav is
read back from the built site rather than assumed.

**A negative result is evidence only when the thing checked was there to be found.** The
page makes four claims of the form "nothing reads this" or "nothing writes that", and each
one is probed with a positive control in the same command, both results reported. This is
the constraint this project cares about most and it bears directly here, because a page
about labels is mostly a page about what does not happen to them.

**Prose wraps at 88 characters; tables and code blocks do not.** The page is largely a
table, which is the shape #359 asks for — a row per label, a column per question — and the
table is the one part that may run long.

**One question at a time.** Three questions are open at the start of this intent: #359's
two, and the spelling of `xeno-needs-decision`. One is asked. The other two are named in
the record with the reason they are not asked yet, which is what keeps them from being
forgotten without putting them on somebody as a batch.
