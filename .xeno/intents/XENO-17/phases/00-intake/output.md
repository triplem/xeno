---
intent: github.com/triplem/xeno#17
phase: 00-intake
created: 2026-09-24T17:53:28Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 0e37f25fc2aefc50994370627bd2d25438075f1f5bcdfd149ba5131c1e05da06
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
open_questions:
  - key: Q-1
    text: Does the go directive carry the patch version or only the minor?
    options:
      - text: "go 1.27, the minor only"
        consequence: setup-go resolves the spec to the newest 1.27.x, so a toolchain
          patch reaches the release without anybody editing a file
        recommended: true
      - text: "go 1.27.1, the exact patch"
        consequence: makes the patch a requirement of the module, and freezes CI on
          that one, since setup-go installs exactly what the spec names
      - text: Something else
        free: true
---

# Intake

Issue #17. `go.mod` declares `go 1.22`, the machine this is developed on runs 1.27.1,
and CI installs whatever the file asks for through `go-version-file`.

## Scope

The declared version, and whatever follows from raising it: the vendoring, which carries
an annotation in `vendor/modules.txt` that does not match what a module at this level
normally has, and `SUPPLY-CHAIN.md`, which lists the toolchain among the things the
pipeline fetches.

The reason to move at all is the support window: 1.22 is out of it, so a vulnerability
found in that toolchain will never be patched there.

## Non goals

Using anything the new version offers. Raising the declared version and writing code
against it are separate changes, and a build that fails after both cannot say which of
them did it.

Both workflows, which read `go-version-file: go.mod` and therefore need no edit. That
they need none is worth checking rather than assuming.
