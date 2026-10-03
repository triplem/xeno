---
intent: github.com/triplem/xeno#186
phase: 05-review
created: "2026-10-03T13:28:45Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6c4a1aa.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 64ccbc0f0f44c01357348c8187012014bf4a4eaf1a2f19e6122e1d1f9bd96454
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
  - rule: deviations-are-traceable
    result: met
    note: >-
      Three, each naming what it departs from: the baseline measured under node 26 rather than the
      pinned 24, recorded as best-available in P3 and closed in P4 by this pull request's own run;
      `gofmt` and `go vet` added to a CONTRIBUTING.md section the design did not ask about; and
      #186's `required_pipeline` clause deliberately left open.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      Three check contexts are renamed, which is a migration: the protection setting must be applied
      after the merge, because a required context that never reports blocks the merge that would
      create it, and two open pull requests report the old names and need a rebase. In the release
      notes, the commit message and the pull request description, because this change is correct in
      the diff and wrong in the sequence.
  - rule: new-dependency-needs-a-rationale
    result: met
    note: >-
      `actions/setup-node` is new here. semantic-release@25 declares ^22.14.0 || >= 24.10.0 and both
      workflows ran under whatever node the runner image shipped. Pinned to a commit sha like every
      other action, and to the node major rather than the patch, per D-5.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Three, each naming what it departs from: the baseline measured
under node 26 rather than the pinned 24, recorded as best-available and closed in P4 by the pull
request's own run; `gofmt` and `go vet` added to a `CONTRIBUTING.md` section the design did not ask
about; and #186's `required_pipeline` clause deliberately left open.

**`interface-change-needs-a-migration-note` — met, and it is the one rule that earns its keep here.**
Three check contexts are renamed, and a renamed required check is a migration: the setting must be
applied after the merge, not before, because a required context that never reports blocks the merge
that would create it. Pull requests #182 and #185 report the old names and need a rebase. All of that
is in the release notes, the commit message and the pull request description, because this is the
kind of change that is correct in the diff and wrong in the sequence.

**`new-dependency-needs-a-rationale` — met rather than not-applicable.** `actions/setup-node` is new
to this repository. The rationale is that `semantic-release@25` declares `^22.14.0 || >= 24.10.0` and
both workflows previously ran under whatever node the runner image shipped. It is pinned to a commit
sha like every other action here, and to the node major rather than the patch, per D-5.

**Beyond the three rules.**

*Is the thing actually enforced?* Not yet, and the review has to say so plainly. This intent makes
enforcement expressible and does not apply it, because it cannot: the contexts it names do not exist
on `main` until it merges. One command afterwards, by somebody with repository administration. Until
then the rule in `CONTRIBUTING.md` is a rule and not a mechanism, and P4's first gap says so.

*Was the baseline raised honestly?* It was read first, which is the rule the file sets about itself,
and the reading changed the answer: a newer release of the same tool reports 30 where the old one
reports 35, and drops the one advisory with supply-chain consequence. So the number moved because the
toolchain moved, and the file records both. A margin was refused. The commit message says what was
read, which is what that file asks of the commit that raises it.

*Was a hook considered?* Yes, because it was asked for, and refused on both of section 7's
constraints — advisory by construction, and no exclusive logic. Written into the design's alternatives
rather than answered in conversation only, which is the thing #188 is about.

*What did the investigation find that the issue did not?* Two defects. That only one of five contexts
was required, and that three of them shared a name and could not have been required individually. The
issue described the trigger. Neither of the other two is visible from the workflow that failed; both
came from reading the protection settings and the job ids.

*Could this change a verdict behind it?* No. No Go file, no gate, no rule, no existing artifact;
`gate verify` is 247 at exit 0.

<!-- xeno:section:release-notes -->
## Release notes

**Every check now gates the merge, where one of five did.** `audit` runs on every pull request with
no path filter, so the job that judges the release toolchain reports before a merge rather than after
it.

**Three checks are renamed**, because the job id is the context GitHub reports and `gitleaks.yml`,
`semgrep.yml` and `trivy.yml` all called theirs `scan`. A required check is matched by that name, so
three identical ones could not be required individually. They are now `gitleaks`, `semgrep` and
`trivy`.

**A host setting completes it, and has to come after this merges:**

    gh api -X PATCH repos/triplem/xeno/branches/main/protection/required_status_checks \
      -F strict=false -f 'contexts[]=verify' -f 'contexts[]=audit' \
      -f 'contexts[]=gitleaks' -f 'contexts[]=semgrep' -f 'contexts[]=trivy'

A required context that never reports blocks every merge, and the three new names do not exist on
`main` until this lands. Open pull requests report the old names and need a rebase afterwards.

**The release toolchain moves to `semantic-release@25.0.9`**, which takes 20 distinct advisories to
12: `sigstore`'s `certificateOIDs` constraints silently dropped, `pacote`'s `addGitSha` DoS,
`picomatch`, `postcss-selector-parser` and three of five `brace-expansion` go, and `undici`'s three
arrive. The npm audit baseline is measured against the new pin — high 30, moderate 1, low 0 — and
records what it was measured against.

**Node is pinned in both workflows**, because 25 declares an engine and neither said which node it
ran under. **And the version is written once**, in `release.yml`, which `audit.yml` reads rather than
copies.

**`CONTRIBUTING.md` has a section on it:** a red check is fixed before the next merge, a scheduled red
is the normal way to learn there is work, and raising a threshold is not what a red check is for.

<!-- xeno:section:residual-risk -->
## Residual risk

**Nothing is enforced until one command runs.** This is the main residual risk and it is a sequencing
risk rather than a defect: the intent makes the setting expressible and the setting is what binds.
Between this merge and that command, `main` still requires one check of five, and the window is
exactly as long as somebody takes to run it.

**A renamed required check is the kind of change that breaks what it does not touch.** Pull requests
open across the rename report the old context names and cannot satisfy the new rule until rebased.
Two are open. Named in the release notes because a contributor discovering it from a stuck merge
button learns the wrong lesson about the gate.

**The required set and the workflow set are two lists with nothing comparing them.** After the command
runs, a sixth workflow arrives unrequired and silently, which is this intent's own defect one level
up. What would close it is `required_pipeline` naming the set, which is an Appendix A change and is
left as a finding on #186.

**The release path is unproven.** A major bump of the tool that cuts every release, and a newly
pinned node under it, neither exercised until the next release runs — and the release workflow only
runs on a push to `main`. The audit job proves the tree resolves clean; it does not prove a tag gets
cut. The failure mode is a release that does not happen rather than a bad one, and
`cycjimmy/semantic-release-action@v6`'s own dependency is `^25.0.2`, so 25 is the version that action
expects.

**Thirty-one findings remain.** DoS and parsing issues in `brace-expansion`, `braces`,
`http-cache-semantics`, `ip-address` and `undici`, reached through npm-as-a-library and the GitHub
plugin's fetch stack, in a job that runs on a push to `main` with `contents: write` and no untrusted
input. That is a reading rather than a proof. The baseline is what makes the thirty-second somebody's
problem.

**What is not a risk.** Any verdict behind this intent — no Go file, gate, rule or existing artifact
moves — and the three scanners, which pass unchanged under their new names.
