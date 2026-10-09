---
intent: github.com/triplem/xeno#95
phase: 03-implementation
created: "2026-10-09T16:50:26Z"
schema_version: "1.0"
runner_version: dev+d2bc411.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 17b22c9203d0812a477803e0cd11936b90b5ab63212851621544d0dc0191c47a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

**`renovate.json`.** Three `customManagers` entries appended, taking the list from six to
nine; `commitBody` removed; ten lines added to the `description` array.

Each new entry carries the `description` this file gives every entry, because a regex
without one is unreadable to the next person and this file is read by people far more often
than by renovate:

- **`gitleaks/gitleaks`**, from `GITLEAKS_VERSION:` in `.github/workflows/gitleaks.yml`,
  `github-releases`. `extractVersionTemplate: "^v(?<version>.*)$"` because the variable
  holds `8.30.1` and the tags hold `v8.30.1`; without it every comparison would be against
  a version that does not exist. The entry's description says what the next person needs to
  know, which is that `GITLEAKS_SHA256` beside it is not renovate's to move and the install
  step's `sha256sum -c` fails until a person recomputes it.
- **`zensical`**, from the `pip install zensical==` line in `.github/workflows/docs.yml`,
  `pypi`.
- **`docker.io/library/debian`**, from `ARG BASE=` in the `Dockerfile`, `docker`. It
  captures three groups from the one line — `depName`, `currentValue` and `currentDigest` —
  so the tag and the digest cannot be proposed apart. This is the only entry that captures
  `depName` rather than templating it, which is why it has no `depNameTemplate`.

`commitBody` carried `Refs #95` and is gone. The ten lines added to `description` say why,
because a key that is absent cannot explain itself: `CONTRIBUTING.md` puts the issue of the
work in the footer and a dependency bump is the work of no issue; the squash takes the pull
request description, so a branch commit's body never reached main anyway; and the value it
held pointed at the issue this intent closes. The same lines say why the three new managers
exist and which one's proposal cannot be complete.

**`Dockerfile`, line 43.** `ARG BASE=docker.io/library/debian@sha256:a29215f6…` became
`ARG BASE=docker.io/library/debian:13-slim@sha256:a29215f6…`. The digest is byte for byte
the one that was there, so `FROM $BASE` resolves the same manifest: a digest wins over a
tag, and the tag is there to name the stream for a reader and for a datasource.

**Nothing else.** `docs/supply-chain.md` is untouched, because all three pins already have
rows and none of their text became wrong — the base image's row already called that pin
`debian` 13-slim, and the reference now agrees with it. No Go source changed, no workflow
changed, and none of the six existing managers was edited, which is what lets them serve as
the control in P4.
