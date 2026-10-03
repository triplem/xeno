---
intent: github.com/triplem/xeno#186
phase: 04-verification
created: "2026-10-03T13:27:49Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6c4a1aa.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8e560a8ae9655ac7daca2b1b4ec2ba2675e1dde8b1a96617aa23cc900280627d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Each acceptance criterion of P1 against what holds it. Most of these are held by a pipeline run and
not by a test, which is what a change to the pipeline means.

**`audit` runs on every pull request** — run `37126158635`, on this pull request, which touches none
of the three paths the old filter named. Under the old trigger it would not have run at all. Held by
the run existing.

**Each check reports under a name that identifies it** — the same pull request's check list:
`verify`, `audit`, `gitleaks`, `semgrep`, `trivy`, five distinct contexts where there were three
called `scan`.

**The version is written once and read** — the `pin` step's own output, `auditing
semantic-release@25.0.9`, which is the number taken out of `release.yml` and not a literal in
`audit.yml`. Held by the log line, which exists so that a reader of the run knows what was examined.

**The audit refuses where it cannot read the version** — `test -n "$v"` under `set -eu`. Asserted by
construction rather than by a run: deliberately breaking the expression to watch the job fail would
have meant a commit whose purpose was to be wrong.

**The toolchain moves to 25.0.9 in both workflows, with node pinned in both** — the `pin` step read
`25.0.9` out of `release.yml`, which is the only place it is written, so the two agree by
construction rather than by comparison.

**The baseline is measured, not raised** — and this is the criterion that was unverified when P3 was
sealed. The local measurement was taken under node 26 because a container run under 24 timed out.
CI's run under the pinned node reports `counts: {"critical": 0, "high": 30, "moderate": 1, "low": 0}`,
which is the baseline exactly. The number is now measured under the node the workflows pin, by the
trigger this intent added, on the pull request that adds it.

**The file says what was read** — `.github/npm-audit-baseline.json`'s `toolchain` and `advisories`
keys, naming 20 against 12 and which advisories moved.

**`audit` is green** — on this pull request. On `main` it follows from the merge, since the job is the
same one and the tree it audits is pinned.

**`CONTRIBUTING.md` says the rule** — read rather than tested, which is what a document is.

**Nothing that was right is changed** — the three scanners pass unchanged, `gitleaks` and `semgrep`
and `trivy` reporting as before under new names; the baseline's drift-not-presence comparison is the
same code; `gate verify` is 247 at exit 0 and no Go file moved.

<!-- xeno:section:results -->
## Results

**The pull request's own `audit` run, `37126158635`**, which is the first run of the trigger this
intent adds and the thing that verifies it:

    auditing semantic-release@25.0.9
    added 301 packages in 10s
    counts: {"critical": 0, "high": 30, "moderate": 1, "low": 0}

Three results in three lines. The version came out of `release.yml`. The tree resolved under the
pinned node. And the counts are the baseline exactly, so the one number P3 recorded as
best-available rather than verified is now verified — under node 24 rather than the node 26 it was
measured on, which was the open question.

**The check list on this pull request:** `verify`, `audit`, `gitleaks`, `semgrep`, `trivy`. Five
distinct contexts, each nameable in a protection rule. Before this, three of them were `scan`.

**Locally, before anything was written.** `semantic-release@24.2.9` installed and audited: 35 high, 2
moderate, 1 low, 20 distinct advisories — matching the failing CI run on `main` exactly, which is
what established that the failure was reproducible and not an artefact of the runner. Then `25.0.9`:
30 high, 1 moderate, 0 low, 12 distinct advisories. `fixAvailable` read for all of them: 30 of the 31
name `semantic-release@15.14.0`, a nine-major downgrade, and four name a direct update.

**`go build`, `gofmt -l .` outside `vendor/`, `go vet ./...`** — all clean, and no Go file is touched
by this intent.

**`go test ./...`** — eighteen packages ok.

**`./xeno gate verify`** — `verified 247 verdicts`, exit 0.

**Evidence is the pipeline's own, for once.** Every figure above except the local measurements comes
from a run this repository's CI produced on this pull request, which is the arrangement section 6
describes and the only kind of evidence a change to the pipeline can offer.

<!-- xeno:section:gaps -->
## Gaps

**The protection setting is not applied, so nothing is enforced yet.** This intent makes enforcement
*expressible* — five distinct contexts, and a job that runs before the merge — and the setting that
binds them is one command somebody with repository administration runs after this merges. Until then
`main` still requires one check of five and the rule in `CONTRIBUTING.md` is a rule rather than a
mechanism. The ordering is forced and stated, not an oversight, but the gap is real while it lasts.

**`required_pipeline` still cannot say which checks**, so `xeno enforcement check` will go on
reporting met the moment one context is required, including in the state above. The tool that was
supposed to notice this will not notice the half-finished version of the fix either. It is an
Appendix A addition, #186 keeps it, and it is the one clause of that issue this intent does not close.

**Nothing verifies that the renamed contexts are the ones required.** After the command runs, the set
in the protection rule and the set of jobs in `.github/workflows/` are two lists with nothing
comparing them. Adding a sixth workflow tomorrow adds a check nobody required, silently — which is
the same shape as the defect this intent fixes, one level up, and is what the `required_pipeline`
gap above would close if it were closed.

**The release path is unproven.** `semantic-release` went from 24 to 25, a major, and the release
workflow only runs on a push to `main`. Its node is newly pinned and its tool newly bumped, and
neither is exercised until the next release. The audit job proves the tree resolves and is clean; it
does not prove the release cuts a tag.

**Thirty-one findings remain and are recorded rather than fixed.** `brace-expansion`, `braces`,
`http-cache-semantics`, `ip-address` and `undici`, reached through npm-as-a-library and the GitHub
plugin's fetch stack. They are denial-of-service and parsing issues in a job that runs on a push to
`main` with `contents: write` and nothing untrusted as input. That is a reading, not a proof, and the
baseline exists so that the thirty-second is somebody's problem rather than nobody's.
