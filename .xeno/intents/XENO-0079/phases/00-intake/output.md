---
intent: github.com/triplem/xeno#79
phase: 00-intake
created: "2026-09-26T14:56:18Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15b30d0
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 307ad27a37305be76aa7fb1da110c5dcdf06746c7b4c812df5302f08757ff67f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

The release builds five binaries and generates one bill of materials, from
`linux-amd64`, then ships that document for all five. A24 recorded why that is accurate
in this tree, a single pure Go dependency and no platform dependent code, and that
nothing checks the condition. The day a build constrained dependency enters, the release
would ship a document describing one target while claiming five, and no step would say
so.

<!-- xeno:section:scope -->
## Scope

A check in the release that compares the module set of every target it builds and fails
naming the ones that disagree, `scripts/module-set.sh`, run before the step that
generates the document. The five targets move into one place that the build loop and the
check both read. `SUPPLY-CHAIN.md` says the single document rests on a checked
condition, and A24 closes.

Not the document itself. One per target is the repair the check points at, not the
change made here, because there is nothing yet to repair.

<!-- xeno:section:context-rationale -->
## Why this context

**Modules, not packages, and that was measured before it was written.** Package sets
differ across the five targets today: `internal/runtime/syscall/linux` against its
Windows counterpart, `internal/runtime/cgroup` on Linux only, all of it standard
library. A check at that level would fail on the first run and prove nothing about a
document that describes modules. At module level the five agree exactly,
`github.com/triplem/xeno` and `go.yaml.in/yaml/v3 v3.0.5`, which is what makes one
document honest today and a divergence meaningful tomorrow.

**The check belongs where it can stop something.** In the release, beside the step it
guards, rather than in the daily scan or a later audit: a document that describes one
target while claiming five is worse than a release that stops, and a release is the only
place the document exists.

**Two lists of targets would drift.** The build loop had them inline and the check needs
the same five, so they are one job level entry now. A check reading a different list
from the build it guards is the failure this would otherwise invite.

**It needs no network.** The dependency is vendored, so `go list` answers from the tree,
which keeps the gate path's rule intact even though this runs in the release rather than
in the gates.