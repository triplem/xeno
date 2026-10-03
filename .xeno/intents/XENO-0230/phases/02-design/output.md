---
intent: github.com/triplem/xeno#186
phase: 02-design
created: "2026-10-03T13:23:03Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fe2286b13b0dc85d68bdf13302a1ef15d9ff7fe59859965eaf6cf91e9cc6aa7f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**`audit`'s `pull_request:` takes no filter at all.** Not a wider path list: a list is a thing to
maintain and the next file it forgets is the next merge past a red gate. The job installs one npm
tree and takes about seven seconds to do it, so the cost is not the argument.

**The three scanners are renamed by their job id, not by adding a `name:`.** The id is what GitHub
reports when there is no name, so one change does it, and an id that says `scan` in a file called
`trivy.yml` was already telling the reader less than the filename.

**`audit.yml` reads the version out of `release.yml` with `sed`, and `test -n` refuses an empty
result.** The release is where the number belongs — it is the release's pin, and the audit exists to
examine it — so the audit is the reader. The refusal matters more than the read: a guard that falls
back to a literal when the expression stops matching audits the wrong tree and says nothing.

**The pin moves to 25.0.9 and node is pinned with it, in both workflows.** They are one change. The
engine requirement is the reason the node pin is not separable: a toolchain that declares one, run
under whatever the runner image ships, is a release that breaks on an image update and nowhere else.

**Node is pinned to the major, `'24'`.** D-5's reasoning for the Go directive, applied: a patch
reaches the release without anybody editing a file, where the exact form freezes CI on one.

**The baseline records the toolchain it was measured against and what the upgrade moved.** Two new
keys beside the counts. The file's own rule is that raising a number is deliberate, and a deliberate
act that does not say what was deliberated about is indistinguishable from the other kind.

**The counts are exact.** `high: 30, moderate: 1, low: 0, critical: 0`, with no margin anywhere.

**The rule goes in `CONTRIBUTING.md`, in a section of its own, and names the six commits.** That file
already holds the merge procedure and the host settings it depends on, which is the same subject. A
rule whose reason is omitted gets read as caution; this one has a morning behind it.

**The protection change is release notes and not code.** It is the actual enforcement, it needs
repository administration, and it has to happen after this merges.

<!-- xeno:section:alternatives -->
## Alternatives

**Requiring `scan` as it stood.** Refused, and the reason is the finding: three workflows report that
context, so the rule names no particular one and its behaviour with three identical contexts is not
something to depend on. Renaming is what makes the setting expressible at all.

**Widening `audit`'s path filter instead of removing it.** Refused. The filter's job was to save a
job run; what it cost was six merges. A list of paths that must trigger the audit is a list that is
correct until somebody adds a file to the release path, and nothing would report that it had gone
stale.

**Leaving the version in two files and checking them with a third thing.** Refused as the same defect
one level up. One writer, one reader; the release owns the number because the release is what runs
it.

**Keeping `24.2.9` and raising the baseline to 35.** Refused after reading the findings, which is the
order the baseline's own comment requires. A newer release of the same tool reports 30 and drops the
one advisory with supply-chain consequence — `sigstore`'s `certificateOIDs` constraints silently
dropped and never enforced — so recording 35 as acceptable would have been recording a number nobody
had to accept.

**Taking npm's suggested fix.** Refused: for thirty of the thirty-one findings npm proposes
`semantic-release@15.14.0`, a nine-major downgrade of the tool that cuts every release. It is a
resolver's answer to a constraint problem and not a security answer.

**A pre-push hook that asks the host whether `main` is green.** Refused on both of section 7's
constraints. Its result would be advisory because hook configuration lives on a developer machine,
and querying a host for check-run status is exclusive logic that the runner does not implement — so
the same push would succeed or fail depending on whose machine it left from. The enforcement is the
protection setting; a hook would be a second, weaker answer that invites somebody to think the first
one is in place.

**A margin on the baseline, 32 for a measurement of 30.** Refused. It converts the next two findings
into silence, which is what a drift check is for.

<!-- xeno:section:impact -->
## Impact

**`.github/workflows/audit.yml`.** `pull_request:` with no filter, and the comment that says why,
naming the six commits. `actions/setup-node` pinned to a sha with `node-version: '24'`. A new `pin`
step that reads `semantic_version` out of `release.yml` and refuses an empty result. The install step
takes the version from that step's output. The header gains two paragraphs: that "exactly as the
release installs it" is now a check rather than a claim, and why node is pinned to the major.

**`.github/workflows/release.yml`.** `semantic_version: 25.0.9`, `setup-node` before the release
step with the same pin and the engine as its reason, and a header paragraph saying the version is
written here and nowhere else.

**`.github/workflows/gitleaks.yml`, `semgrep.yml`, `trivy.yml`.** The job id, `scan` to the tool's
name, with the comment explaining that the id is the reported context and that three identical ones
cannot be required individually.

**`.github/npm-audit-baseline.json`.** `high` 15 to 30, `moderate` 3 to 1, `low` 1 to 0, `critical`
unchanged at 0; `measured` to today; two new keys, `toolchain` and `advisories`, recording what it
was measured against and what the upgrade moved.

**`CONTRIBUTING.md`.** `gofmt` and `go vet` added to "Before sending a change", which the project's
own build instructions already list and that section did not. Then a new section: every check gates
the merge, a scheduled red is the normal way to learn there is work, the six commits and why they
were honest, and that raising a threshold is not what a red check is for.

**No Go code changes.** Nothing in `internal/` or `cmd/` moves, so the suite and every verdict are
untouched; `gate verify` is unaffected by this intent except that its own trail grows by six phases.

**No document changes.** The one that would be needed is left as a finding on #186.

**What this does not change and somebody must.** `main`'s protection, after this merges:

    gh api -X PATCH repos/triplem/xeno/branches/main/protection/required_status_checks \
      -F strict=false -f 'contexts[]=verify' -f 'contexts[]=audit' \
      -f 'contexts[]=gitleaks' -f 'contexts[]=semgrep' -f 'contexts[]=trivy'
