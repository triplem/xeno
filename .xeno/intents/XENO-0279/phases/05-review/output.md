---
intent: github.com/triplem/xeno#94
phase: 05-review
created: "2026-10-08T20:41:26Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 288c84294245e3ab24ea84f4f8078a14eb44c8b13297882aa3ba4794086e0d12
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: met
      note: 'Three deviations in P3, each naming the design decision it departs from and the measurement that forced it: gosec as its own binary, the acceptances configured and excluded by class rather than four inline, tests outside gosec by a flag.'
    - rule: interface-change-needs-a-migration-note
      result: met
      note: The required checks change from one to three; the move is scripts/github-settings.sh run against the host before the merge, with the read-back on the pull request, and a rebase for any pull request open across it, which strict protection asks for anyway. Nothing an adopter uses changes.
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: No module enters go.mod; three pipeline tools are rows on the pin table with their fetch named and the action weighed against go install in the design's alternatives.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**deviations-are-traceable — met.** P3 records three deviations and each names the
design decision it departs from: gosec as its own binary against the first decision,
the shape of the acceptances against the second, a flag against the third's rule. Each
carries the measurement that forced it, five runs of one commit with five sets, and the
one that confirmed the replacement, two runs of the binary with one set.

**interface-change-needs-a-migration-note — met.** The interface that moves is the set
of checks a pull request has to pass: `semgrep` goes, `lint`, `gosec` and `govulncheck`
come, and the branch protection on the host names them. A pull request open across the
merge sees three new required checks appear and one disappear; it needs a rebase onto
`main` to carry the workflows, which `strict: true` asks of it anyway. The move is the
script in `scripts/github-settings.sh`, run once against the host before the merge, and
the pull request description says it was run and shows the read-back. No command, flag
or file an adopter uses changes.

**new-dependency-needs-a-rationale — not-applicable.** No module enters `go.mod`. Three
pipeline tools do, each a row on the pin table with its fetch named, and the design's
alternatives weigh the action against `go install` for each; that is the table's
business and not this rule's.

**What no rule asks.** Whether the fifty-eight findings read by rule rather than one by
one hide one that is real: P4's gaps carry the claim they rest on, that no path this
runner reads arrives from the network, and where it was checked. And whether the
intake's wrong figure should have been caught before the design was written against it:
P3's learning proposes the rule, two cold runs or a note that there was one.

<!-- xeno:section:release-notes -->
## Release notes

**semgrep is replaced by three Go tools, and the pull request has three checks where it
had one.** `lint` runs govet, staticcheck and errcheck through golangci-lint 2.14.0, with
the policy in `.golangci.yml`. `gosec` runs gosec 2.29.0 as its own binary, with the file
modes a git checkout has configured in `.gosec.json`, the two path rules excluded for a
reason the workflow states, tests outside its scan, and every accepted finding carrying
its rule and its reason on the line. `govulncheck` scans the source daily at 1.8.0 and
fails on a reachable vulnerability that has a fix, judged by a step that reads the
report, since the tool's JSON mode exits 0 whatever it finds. Each writes a JSON report
and a manifest in the shape the other scans have, so a phase can declare it as evidence.

**trivy is unchanged.** It scans what the release ships, by version; govulncheck scans
what the source calls, by call path. Two answers to one question from two sides, kept
on purpose.

**Why gosec is not inside golangci-lint.** Measured on this tree: under golangci-lint
its findings differed between five runs of one commit and missed forty-seven of
fifty-eight, while the binary reported the same fifty-eight twice. A required check has
to report the same thing twice on one tree, so gosec runs alone until the integration
does.

**Zero findings, reached.** 0 lint, 0 gosec with four accepted, 0 vulnerabilities, on a
tree whose only changes to shipped code are four comments, one selector and one
`Fprintf`. The threshold is the tools' exit codes, with nothing filtered by severity or
confidence.

**For whoever maintains the pipeline.** The pin table has five new rows and two fewer;
renovate proposes the next version of each tool; the required checks are `verify`,
`audit`, `gitleaks`, `lint`, `gosec`, `govulncheck`, `trivy`, set from
`scripts/github-settings.sh`. `.semgrep/` is gone and nothing reads it.

**For an adopter.** Nothing. This is this repository's own pipeline, which the
maintainer's first answer said it would be.

<!-- xeno:section:residual-risk -->
## Residual risk

In order of weight.

**The G304 and G703 exclusion is a claim about the whole tree, checked by rule.** Every
path this runner reads is computed from the repository or given on the command line,
and nothing arrives from the network: true of the adapters, which read bodies and not
paths, and of the twenty-eight findings' files, and asserted rather than proved for the
rest. A future reader that takes a path from a host answer would be excluded from the
rule that exists to catch it. The remedy is a reader, not a rule: the day one exists,
the exclusion narrows to the files that are a runner's and the new one meets the
finding.

**gosec's integration into golangci-lint is unexplained.** Five sets from five runs is
shown; why is not. If the cause is in gosec itself and the binary is stable only on
this tree's size, the second job is stable by luck. Two identical runs of the binary are
the evidence; a third on the runner is the next.

**The host's branch protection is changed by hand.** The script is the record and the
pull request carries the read-back, but nothing in the tree can say the host agrees,
and `xeno enforcement check` reads the enforcement block and not the context list. A
protection still naming `semgrep` blocks every merge, loudly; one that forgot `gosec`
passes a red audit silently. The second case is the one to check after the merge.

**The reports are evidence nobody declares.** Three manifests in section 4's shape, and
no phase of this intent or of the eighty before it attaches a scan report. The plumbing
is kept because the maintainer wants it; the practice does not exist yet, and a shape
nobody uses is the kind that drifts.

**The intake's figure was wrong and the design was written against it.** Caught in
implementation, recorded as deviations, and the register row rewritten once inside the
phase; the record is honest about the sequence, and the lesson, two runs or a note of
one, is a learning and not yet a rule.

**Belongs to the specification, not to this code.** Whether a scan report's judgement
belongs in a workflow step, as govulncheck's now does, or whether section 4's evidence
shape should carry the threshold a `kind: scan` item was judged by; and whether
`trivy.yml` should scan the image #282 publishes, which A105 leaves to a question of
its own.
