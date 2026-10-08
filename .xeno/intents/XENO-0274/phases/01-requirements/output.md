---
intent: github.com/triplem/xeno#95
phase: 01-requirements
created: "2026-10-08T13:37:19Z"
schema_version: "1.0"
runner_version: dev+042c9bc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a7f8ca46b844f6762872e83ea8990b210ea16d081a20ca1fd4278772d0a17d77
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. **Renovate is configured in the repository.** A `renovate.json` exists at the root, it is
   valid against renovate's own schema reference, and it names what this repository actually
   holds: GitHub Actions pinned by sha, the npm tools the workflows install, and the one Go
   module.

2. **The dashboard is the default and a pull request is not.** The configuration opens no
   pull request of its own accord; what it produces is the dependency dashboard the issue
   asks for, which is a list a person works from. A reader of the file can tell which of the
   two it does without running it.

3. **A bump that would break the pin tests is not proposed silently.** Every sha this
   repository pins is also a row of `docs/supply-chain.md`, and #317's two tests fail when
   the two disagree. The configuration either keeps such an update off the automatic path or
   says in its own comments that a person moves the row with it, and it says which where
   somebody reading it will meet it.

4. **Nothing is bumped by this intent.** The merge changes no pinned sha, version or digest,
   and `go test ./internal/model/` still passes, which is what proves the pins and the page
   still agree.

5. **What this does not reach is written down.** Switching dependabot off is a host setting,
   and the record says so rather than leaving a reader to assume this commit did it.
