---
intent: github.com/triplem/xeno#128
phase: 02-design
created: "2026-09-29T09:30:41Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+38430d4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7a491f4b36247b02e42c027c7571efbb29b5c9c677e07ebfd6ad484d99851f9d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: by-hand
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**gitleaks, pinned by version and verified by checksum, not taken from an action.** The
tool whose rules #127 translated, so the two instruments share a vocabulary and a
redaction marker here names a rule a reader can look up there. An action would be a
second thing to pin and a third party running in the job; a tarball with a published
sha256 is the smaller surface.

**The configuration extends the binary's rules rather than pointing at the project's
copy.** `.xeno/plugin/secrets.yaml` is that set with the entropy thresholds, keyword
prefilters and allowlists removed, which section 4 has no room for. Running it here
would repeat the runner's work with the precision devices missing.

**The allowlist names shapes, never rules.** One path and four regexes: the pattern
file, which contains the literals it matches on; the artifact hash fields; the two
planted samples; and the Go identifier `key` followed by a string. Each says what this
repository contains. Disabling `generic-api-key` is one line and removes a rule that
catches real keys.

**The subject is a checkout.** What is tracked, which is what was committed. Ignored
files are not there, and the comment says that a local run reading a real `.env` is the
scan working.

**The report's existence is asserted before the failure is re-raised.** gitleaks exits 1
both when it finds a leak and when it cannot run, and those must not read alike. semgrep
needs the opposite guard, `--error`, because it exits 0 with findings present, and the
two workflows each say which way their tool is wrong.

**Nothing prints the match.** Rule, file, line, fingerprint. The report carries the
matched text and a build log is a poor place for it.

**Every finding fails the job.** No threshold and no baseline, which a repository with
no findings can afford and is the reason AC3 insists on landing clean.

<!-- xeno:section:alternatives -->
## Alternatives

**Use the official gitleaks action.** One line instead of five. It is another artifact
to pin by sha, it runs third party code in the job, and its licensing terms differ for
organisations, which is a question this repository should not have to answer to run a
scan. The tarball and a checksum are smaller and say more.

**Point the scan at `.xeno/plugin/secrets.yaml`.** Attractive because one file would
then define what a secret is everywhere. It is the wrong direction: that file is
upstream's rules with the precision devices stripped out, so the scan would inherit the
widening that a redactor can afford and a scanner cannot, and it would find nothing the
runner had not already filtered.

**Disable `generic-api-key`.** One line, and it removes the rule that reported
thirty-nine of forty findings. Those thirty-nine are one shape, the artifact hash
fields, and an allowlist naming that shape keeps the rule for everything else. Rejected
because the count made it tempting and the count is one pattern.

**Exclude `.xeno/` from the scan.** It is where the noise was. It is also where a leaked
secret in a digest would land, which is the case the whole of #120 is about. Rejected
outright.

**Scan the git history rather than the working tree.** It finds a secret that was
committed and later removed, which is the case that most needs finding. It also fails
forever once anything is in the history, since history does not change, and it makes
every future run pay for the whole log. A repository with a clean tree today should
start with the tree; the history is a separate decision with a separate remedy.

**Add a baseline file.** The usual way to introduce a scanner to a repository with
findings. This repository has none once the shapes are named, so a baseline would record
noise as accepted and hide the next real finding behind it.

<!-- xeno:section:impact -->
## Impact

`.github/workflows/gitleaks.yml`: the workflow, in the shape `semgrep.yml` established —
scan with `continue-on-error`, report, assert the report exists, raise the failure
again.

`.gitleaks.toml`: `useDefault = true` and one allowlist.

`SUPPLY-CHAIN.md`: a row for the binary beside the one for the rules.

Nothing else. No Go, no dependency, no change to `internal/secrets`.

What a reader gains: every file that stays in the repository is read once per change by
a maintained rule set with its precision devices intact, which is the coverage the
digest filter was never going to provide.

What they still do not gain: G-Secret, a pre-commit catch, and any reading of the
history.

What this costs: a job of about a minute per change, a file whose allowlist has to be
kept true as the repository grows, and the certainty that the first real finding will
arrive as a red pull request from somebody who did not write the secret.
