---
intent: github.com/triplem/xeno#186
phase: 03-implementation
created: "2026-10-03T13:23:43Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 85205c2614fb99e6ec66f9901c811b95dae6b9dfbbb0b6734eebb010cd2a02fe
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

**`.github/workflows/audit.yml`.** `pull_request:` with no filter, carrying the comment that a job
running only on a push to `main` can report only after the merge and that six commits went past a red
run of it. `actions/setup-node@8207627` with `node-version: '24'`. A `pin` step reading
`semantic_version` out of `release.yml` with `sed`, `test -n "$v"` refusing an empty result, and the
version going to `$GITHUB_OUTPUT`; the install step consumes it. Two header paragraphs: that
"installed exactly as the release installs it" is now a check, and that node is pinned to the major
for D-5's reason.

**`.github/workflows/release.yml`.** `semantic_version: 25.0.9`. `setup-node` at the same pin before
the release step, with the engine as its stated reason. A header paragraph saying the version lives
here and nowhere else.

**`.github/workflows/gitleaks.yml`, `semgrep.yml`, `trivy.yml`.** Job id `scan` to `gitleaks`,
`semgrep`, `trivy`, each with the comment that the id is the reported context and three identical
ones cannot be required individually. Nothing else in those files moves — their manifests already
wrote distinct `job:` values and their artifacts distinct names, which is what made the rename safe.

**`.github/npm-audit-baseline.json`.** Counts to the measurement: high 30, moderate 1, low 0,
critical 0. `measured` to 2026-10-03. Two new keys: `toolchain`, naming
`semantic-release@25.0.9, read from .github/workflows/release.yml`, and `advisories`, recording 20
against 12, which advisories the upgrade moved, which remain, and that npm's suggestion for thirty of
them is a nine-major downgrade.

**`CONTRIBUTING.md`.** `gofmt -l .` and `go vet ./...` added to "Before sending a change", which the
project's build instructions list and that section did not. A new section, "A red check is fixed
before the next merge": every check gates the merge rather than one of them, a scheduled red is the
normal way to learn there is work, the six commits and why each was honestly green, and that raising
a threshold is not what a red check is for.

**Measured before any of it was written.** `semantic-release@24.2.9` installed and audited locally:
35 high, 2 moderate, 1 low, 20 distinct advisories, matching CI's run exactly. Then `25.0.9`: 30
high, 1 moderate, 0 low, 12 distinct advisories. The fix availability was read too — `fixAvailable`
names `semantic-release@15.14.0` for thirty of them.

**No Go code, and no document.** `internal/` and `cmd/` are untouched.

<!-- xeno:section:deviations -->
## Deviations from the design

**One from the design, and it is a measurement that could not be taken.** The baseline's numbers come
from an audit run under node 26 locally, not under the node 24 the workflows now pin. A container run
to measure it under 24 was attempted and timed out. npm resolves a tree, and a different npm can
resolve a different one, so the recorded counts are the best available measurement and not a verified
one. This is what the `pull_request` trigger is for and is the first thing it is used for: CI measures
it under the pinned node on this very pull request, and the number is corrected there if it differs.
Recorded as a deviation because a baseline that disagrees with its own job is the defect this intent
is about.

**One widening, in `CONTRIBUTING.md`.** `gofmt` and `go vet` were added to "Before sending a change",
which the design did not ask for. That section listed two of the four commands `CLAUDE.md` and the
README give as the build, and a section telling a contributor what to run before sending a change
while omitting the two that gate it is the same class of gap as the rest of this intent. Small, and
outside what the issue asked for, so it is written down.

**One thing the issue asked for that this does not do.** #186's "done when" includes
`xeno enforcement check` reporting unmet where a check exists and is not required, which means
`required_pipeline` naming the set. That is an Appendix A addition and therefore a person's commit
first. The issue says so itself — "this issue is the finding, not the change" — and it stays open on
that clause after this merges.
