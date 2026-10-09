---
intent: github.com/triplem/xeno#95
phase: 02-design
created: "2026-10-09T16:49:13Z"
schema_version: "1.0"
runner_version: dev+d2bc411.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f6d8680bc355e51dbd4bbe037c45ec6e673f547e62e0518ab48017c1f3e75426
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:alternatives -->
## Alternatives

**Enable the `dockerfile` manager instead of a regex.** The native manager is the obvious
answer and it does not fit: the base image is the value of `ARG BASE`, and `FROM $BASE`
names a variable. Renovate does read some `ARG`/`FROM` pairs, but relying on which ones
makes the coverage depend on a behaviour this repository cannot test and would not notice
losing. A regex that names the line is checkable here, by the same script that checks the
other eight, and it fails visibly if the line is rewritten. Rejected for that, not because
the manager is wrong in general.

**Enable `pip_requirements` for `zensical`.** There is no requirements file. `zensical` is
an argument to `pip install` inside `docs.yml`, so the manager would resolve nothing while
its presence in `enabledManagers` read as coverage. That is the failure mode this intent
exists to remove, so adding an instance of it would be perverse.

**Leave the `Dockerfile` reference digest-only and match the digest alone.** This keeps the
diff to `renovate.json` and nothing else, which is attractive. It was rejected because a
digest names no stream: renovate needs a tag to know which image's newer digest to offer,
and with none it falls back to `latest`, which would propose the digest of `debian:latest`
for a pin the page calls 13-slim. The tag also closes a smaller gap that exists regardless
of any bot — `docs/supply-chain.md` names a tag the tree does not carry, so a reader
checking the page against the `Dockerfile` cannot confirm the one from the other.

**Leave `gitleaks` uncovered, because its proposal cannot be merged as it arrives.** The
honest version of this argument is that a proposal which always fails its own CI is noise.
It was rejected because the alternative is worse in the way this project cares about: a pin
that goes stale with nothing to notice. The configuration already made this trade for
`github-actions`, where a bumped sha needs a `docs/supply-chain.md` row moved by hand, and
its `packageRules` entry says so in as many words. A failing checksum is loud, nothing
automerges, and a person finishes the proposal — which is the same sentence.

**Keep `commitBody` and point it at #350 instead of #95.** Rejected: #350 is the issue for
adding a secret, not the issue a dependency bump belongs to, so the footer would be wrong
in a new way rather than stale in the old one. Renovate has no template for the dependency
dashboard's own issue number, which is the only reference that would have been defensible,
and the squash discards the branch commit's body in any case.

**Write a Go test over the regexes instead of a script.** Rejected: the dialect is
JavaScript's, so a Go test would check a translation of the regexes rather than the
regexes. `(?<name>…)` and the `/…/` delimiters of `managerFilePatterns` are renovate's, and
a translation that passed while the original failed is precisely the undetected gap being
closed. The check runs the real strings instead, and its output is the verification phase's
evidence.

<!-- xeno:section:impact -->
## Impact

Two files change and nothing in the runner does.

**`renovate.json`.** Three entries appended to `customManagers`, bringing it from six to
nine, each with the `description` this file already gives every entry. One key removed,
`commitBody`. Ten lines added to the `description` array at the top, which is where this
file explains itself to a reader who has not read #95.

The three entries:

| pin | file | datasource | the detail that matters |
|---|---|---|---|
| `gitleaks/gitleaks` | `gitleaks.yml` | `github-releases` | `extractVersionTemplate` drops the tag's `v`, because `GITLEAKS_VERSION` carries `8.30.1` and the tags carry `v8.30.1` |
| `zensical` | `docs.yml` | `pypi` | matched on the `pip install` line; its fourteen transitive dependencies stay unreachable, which the page's row already says |
| `docker.io/library/debian` | `Dockerfile` | `docker` | captures `depName`, `currentValue` and `currentDigest` from one line, so the tag and the digest move together |

**`Dockerfile`, one line.** `ARG BASE=docker.io/library/debian@sha256:…` becomes
`ARG BASE=docker.io/library/debian:13-slim@sha256:…`. Nothing else in the file, and nothing
about how the image is built: `FROM $BASE` resolves the same manifest it resolved before,
because the digest is unchanged and a digest wins over a tag.

**What this does to `verify`.** `internal/model/supply_chain_test.go` extracts the
`Dockerfile`'s digest with `([A-Za-z0-9][A-Za-z0-9._/-]*)@sha256:([0-9a-f]{64})`, and the
character class excludes `:`. So on the new line the first group captures `13-slim` instead
of `docker.io/library/debian`, and the second group — the digest, which is the only one
stored — is unchanged. The test keeps passing for the same reason it passed before, and the
reason is in the pattern rather than in luck. That was read before the edit was made, which
is why `internal/model/supply_chain_test.go` is in the context scope.

**What this does not touch.** `enabledManagers`, `packageRules`, `postUpgradeTasks`,
`vulnerabilityAlerts`, `automerge`, `dependencyDashboard`, `schedule`,
`prConcurrentLimit`, `semanticCommits`, `pinDigests`, `postUpdateOptions`, and every one of
the six existing `customManagers`. The decided shape is untouched and the six entries that
already resolve correctly are not edited, which is what makes them usable as the control in
the verification.

**What it does to `docs/supply-chain.md`.** Nothing. Every one of the three pins already has
a row, and no row's text becomes wrong: the base image's row already calls the pin `debian`
13-slim, which is what the reference now says too.
