---
intent: github.com/triplem/xeno#186
phase: 01-requirements
created: "2026-10-03T13:22:11Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a67d39b98e6fb315eb6f265f3c78bdd21b44e158bdfcd26f0cc4d653af27b04b
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

**`audit` runs on every pull request.** No path filter. The job that judges the release toolchain
reports before a merge, which is the one property it did not have.

**Each check reports under a name that identifies it.** `verify`, `audit`, `gitleaks`, `semgrep`,
`trivy` — five distinct contexts, so each can be named in a protection rule. Three of them were
`scan`.

**The version is written once.** `release.yml` holds `semantic_version` and `audit.yml` reads it out
of that file. Two copies of one number is how an audit comes to examine a tree nothing ships, and
the audit's header claims it installs what the release installs.

**The audit refuses where it cannot read the version**, rather than falling back to a literal. A
guard that silently audits the wrong thing is the defect it exists to prevent.

**The toolchain moves to `semantic-release@25.0.9`** in both workflows together, and node is pinned
to the same major in both, because 25 declares an engine and neither workflow said which node it ran
under.

**The baseline is measured against the new pin, not raised to meet the old one.** 20 distinct
advisories on 24.2.9 against 12 on 25.0.9; `sigstore`'s dropped `certificateOIDs` constraints,
`pacote`'s `addGitSha` DoS, `picomatch`, `postcss-selector-parser` and three of the five
`brace-expansion` advisories go with the upgrade, and `undici`'s three arrive. The counts recorded
are what the tool reports, exactly, with no margin: a margin is a new finding nobody is told about.

**The file says what was read.** The baseline carries the toolchain it was measured against and what
remains after the upgrade, because its own rule is that raising a number is a deliberate act.

**`audit` is green on `main`.**

**`CONTRIBUTING.md` says that a red check is fixed before the next merge**, that every check gates
the merge rather than one of them, and that a scheduled run going red is the normal way to learn
there is work — with the six commits named, because the rule exists for a reason that happened.

**Nothing that was right is changed.** The baseline's drift-not-presence rule, the daily schedules,
`enforce_admins`, and the three scanners' own logic and artifacts.

<!-- xeno:section:non-goals -->
## Non goals

**The protection settings are not changed by this intent.** They are a host setting, the category
`xeno init` prints as what Xeno cannot do, and they cannot be applied before this merges: the three
renamed contexts do not exist on `main` yet, and a required context that never reports blocks every
merge including the one that would create it. The command and the ordering are in the release notes.

**`required_pipeline` does not learn to name the set here.** It is the part that was supposed to
notice, and it is an Appendix A addition, so a specification change and its own commit first. #186
carries the finding.

**No hook.** Section 7 answers it twice: a hook's result is advisory because its configuration lives
on a developer machine, and a hook invokes only checks the runner already implements. Querying a
host for the status of its check runs is neither. A hook is the fast half and this is about the
binding half.

**The thirty-one remaining findings are not fixed.** They are read, named and recorded. npm's
suggestion for thirty of them is to downgrade `semantic-release` to 15, which is a worse answer than
the problem.

**No new scanner, no new gate, no change to what the three scanners examine.**

**No retroactive judgement of the six commits.** Each was green on what ran, the trail is sealed, and
`gate verify` has nothing to say about them. What changed is what the next one has to pass.

<!-- xeno:section:constraints -->
## Constraints

**A required context that never reports blocks every merge.** That single fact fixes the order of
this work: rename first, merge, then require. Doing it the other way round locks the repository out
of the commit that would unlock it.

**The baseline's own rule.** "Raising a number here is a deliberate act and belongs in the commit
message that raises it." So the number could not be raised without reading the findings first, and
the reading is what produced the upgrade.

**A margin is a lie.** The counts recorded are exactly what the tool reports. A baseline set one
above the measurement tolerates the next finding silently, which is the failure mode the whole
arrangement exists to avoid.

**The audit's premise has to hold.** Its header says the tree is installed exactly as the release
installs it. With the version in two files that was a claim; it has to become a check or the header
is prose about something that may not be true.

**Section 7 on hooks is binding on the design.** Advisory by construction, no exclusive logic. The
question was asked and this is the answer.

**The specification does not change.** Nothing here needs section 5, section 7 or Appendix A to say
anything new; the one thing that would — `required_pipeline` naming the set — is deliberately left
as a finding.

**The node pin follows D-5.** The major and not the patch, for the reason the Go directive has:
a patch reaches the release without anybody editing a file, and the exact form freezes CI on one.

**88 columns, SPDX where it applies, `gofmt`, `go vet`, the suite, `./xeno gate verify` at exit 0.**

**One intent, one branch, two issues** — `186-the-audit-gate-binds`, closing #186 and #187, labelled
wp10 and wp0. Two issues because they are one act: the gate that did not bind, and the red it did
not bind on.
