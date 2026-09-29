---
intent: github.com/triplem/xeno#136
phase: 03-implementation
created: "2026-09-29T18:08:34Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4b87535.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3223d98fcb3a42e9df8ef74be19432b8753d2827cdcaf98bf57129ca3c8a5c44
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: by-hand
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`cmd/xeno`. Two calls, three words each: `CREATED  INTENT  STATE  PHASE` and `PHASE
STATE  VERDICT`.

The comment above `listRow` now says which convention applies and which one XENO-0207
borrowed. It names `ps` rather than arguing from taste, and it says that this project's
rule about a heading naming its section in words is about prose, in files and in issues.
Without that, the next reader has the same citation available for the same wrong
conclusion.

`cmd/xeno/list_test.go`. Four expected strings. The tests are the same two, asserting
the same property against the same widest rows; only the literals they look for moved.

Nothing else. No format string, no width, no order, no default, no notice, no data, and
no package outside `cmd/`.

**What it prints.** `CREATED     INTENT       STATE             PHASE` above the
listing, and `PHASE STATE                  VERDICT` above a phase table. The widths are
XENO-0207's, because the columns were sized for the data and upper case is no wider.

<!-- xeno:section:deviations -->
## Deviations from the design

None. Two literals, a comment and four strings in a test, as P2 decided.

The thing worth recording is not a deviation but the shape of the intent: the diff is
two lines of behaviour and the record around it is seventeen sections. That is not a
complaint and it is not an accident either — section 11's rule about sealed artifacts is
what made a second intent the honest route, and #117 is where the figure belongs rather
than here.

The order of the work was right. P0 named the borrowed reason, P1 said what must not
move, P2 chose to reverse rather than rewrite, and the code followed.
