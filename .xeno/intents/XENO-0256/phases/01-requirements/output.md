---
intent: github.com/triplem/xeno#260
phase: 01-requirements
created: "2026-10-06T07:22:51Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d11d8502c5dae48ecae2a4f9401a314acddac936570120d05caf51c541862b81
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

1. `.github/npm-audit-baseline.json` records `critical: 0`, `high: 10`, `moderate: 2`, `low: 0`,
   and `measured: 2026-10-06`.

2. The `advisories` prose describes the twelve advisories that exist now, not the thirty-one that
   did, and no longer says `postcss-selector-parser` is gone — it is back, and the prose says so.

3. The prose records that bumping and fixing were tested and neither is available: 25.0.9 is the
   latest, an install of `@latest` gives an identical tree and identical counts, and `npm audit
   fix` resolves none of the twelve.

4. The prose records why the file drifts: the tree is resolved at install time with no lockfile, so
   the counts are a property of npm's registry at the moment of the run.

5. The `toolchain` field names the version the figures were measured against, and it matches what
   `release.yml` pins.

6. The commit message carries the deliberate act, which the file asks for by name, and says both
   directions: high falls by twenty, moderate rises by one.

7. The audit job passes. That is the point of the intent and it cannot be checked here — nothing in
   this repository runs `audit.yml` — so the criterion is met on the pull request and the
   verification phase says so rather than claiming a local run.

8. `Refs #260`, not `Closes`. The lockfile question is the open half and this intent is the common
   prefix of both remaining routes.

9. `./xeno gate verify` exits 0 with the verdicts that exist now intact, plus this intent's own
   judged phases. The baseline is not covered by `artifacts_hash` and nothing reads it from the
   trail, so the expectation is no effect and the command confirms it.

10. `go build`, `go test ./...`, `go vet ./...` pass and `gofmt -l` outside `vendor/` prints
    nothing. None can be affected by a JSON file the Go code never reads.

11. The JSON parses, and the file is one object with the same five keys it had.

<!-- xeno:section:non-goals -->
## Non goals

Not a lockfile. Whether the counts should become a function of this repository is the open half of
#260, separable, and the counts it would pin are the same 10 and 2 — so this forecloses nothing.

Not a change to what the check asserts. Failing only on critical and high, or only on an advisory
with an applicable fix, is the third route on #260 and gives up signal A44 wanted.

Not a downgrade of `semantic-release`. A44 declined it, A28 pins deliberately, and npm's
suggestion of 15.14.0 is a major downgrade of the release machinery for seven of the twelve.

Not a change to `audit.yml`. The job's comparison is right and its install is what the release
installs, which is the property that makes the audit meaningful; what is wrong is that the numbers
it compares against are three days and twenty advisories out of date.

Not a resolution of any advisory. None has a fix this repository can apply, which is now measured
rather than asserted, and the baseline's purpose is to notice a change rather than to fix one.

Not a refresh cadence. The learning proposes that a baseline measured against an unpinned tree
needs one; proposing it is section 10's route and adopting it is not this intent's.

Not an `ASSUMPTIONS.md` row. A44 already records the arrangement and nothing about it changes; what
changes is the numbers it was measured at, which live in the file A44 points to.

<!-- xeno:section:constraints -->
## Constraints

The file states the rule this intent has to follow. "Raising a number here is a deliberate act and
belongs in the commit message that raises it", so the commit message is part of the work rather
than a description of it, and a message that mentioned only the twenty-advisory improvement would
be using a tightening to smuggle a loosening.

The numbers are not this repository's to choose. They are what `npm audit` reports for the pinned
toolchain today, and the only honest value is the measured one — which is why both the CI run's
output and a local reproduction are recorded, and why they agree.

Nothing here can be verified locally except the file. The audit job runs in CI against a tree
built at run time; a local install reproduces the counts but not the job, so criterion 7 is met on
the pull request and nowhere else. Claiming otherwise would be the defect #201's verification phase
recorded about the release workflow.

The counts will drift again. This intent records a measurement with a date and does not stop the
mechanism; the next upstream publication can fail an unrelated pull request the same way, and the
prose has to say that rather than implying the numbers are a bound.

`Refs #260` and not `Closes`, by `CONTRIBUTING.md`'s rule: `Closes` goes on the commit that
finishes the work, and the lockfile question is unfinished.

One intent, one branch, and the issue carries `wp0`.
