---
intent: github.com/triplem/xeno#343
phase: 02-design
created: "2026-10-09T13:19:47Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bd5cac852dea3bd30df8014a92338ef4a983da8f5e82ad46de9e66af6652f6c5
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

**A pseudo-version, because it is a pin the table already knows how to read.** `go
install` accepts `v2.29.1-0.20261009120814-7b1b5cebe007` and resolves it through the
module proxy and the checksum database like a tag; the pin test's `versionPin` shape
reads `v2.29.1` out of the workflow and out of the row and holds them together, so the
row is written with the same string and nothing in the test moves. A bare commit sha on
the line would be a pin the test cannot read without a new shape.

**The reason lives on the line.** The workflow's install step gets a comment naming Go
1.27.2, `x/tools` 0.51.0, the commit and the release that ends the commit pin, because a
pseudo-version with no sentence beside it reads as somebody's accident.

**A104's status cell is replaced, not appended to.** It says the scanner is pinned to a
commit since #343, why, and that the row returns to a release at 2.30.0.

<!-- xeno:section:alternatives -->
## Alternatives

**Pin Go to 1.27.1 in this one job.** Green today, and a toolchain pinned by patch in one workflow against the major everywhere else; the next patch does the same thing again. Declined.

**A tools module.** `tools/go.mod` requiring gosec 2.29.0 and `x/tools` 0.51.0. Reproducible, and a second module file in a repository whose supply chain story is one vendored dependency, for a line a pseudo-version answers. Declined.

**`securego/gosec` as an action.** Its binary loads packages with the runner Go and would fail the same way until the action ships the same commit. Declined.

<!-- xeno:section:impact -->
## Impact

**For a pull request.** `gosec` is green again under 1.27.2, with the same policy.

**For the pin table.** One row changes its version string; renovate proposes 2.30.0 when it exists and the pseudo-version goes with it.

**For the trail.** Nothing sealed changes.
