---
intent: github.com/triplem/xeno#324
phase: 02-design
created: "2026-10-08T13:46:50Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fe01733c917a299e8973c17db55a1691110f151d0ad7dde15af460f416382187
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:impact -->
## Impact

Two lines in one template, and a test.

`internal/scaffold/files/ci-github.yml` gains, inside the job's `container:` block:

      # options: --user 1001        # the uid your runner's work directory belongs to

with a comment above it saying why: the host mounts its work directory into the container
and chowns nothing, `actions/checkout` runs inside the container, and the image's user has
to be able to write there. A hosted runner needs nothing, because 1001 is what the image
already runs as and what that account is.

Why commented rather than present. An uncommented `--user` would have to carry a number, and
the only number this repository can write is the one that is already right — so the line
would either be redundant on the hosted runner or wrong everywhere else. Commented, it is
the one thing the file can honestly offer: the name of the knob and the reason to reach for
it.

Why not a setting. `project.yaml` is enumerated by Appendix A, and the uid is a property of
the runner the adopter operates rather than of their project. A field would put a machine's
accident into the file that describes the work.

The test goes beside `TestEveryWrapperNamesTheRunnerImage` in `internal/runner/init_test.go`,
which already renders every wrapper and reads it back, and it asserts both halves: the GitHub
wrapper carries the commented line, and the GitLab wrapper does not.
