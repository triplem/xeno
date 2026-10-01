---
intent: github.com/triplem/xeno#158
phase: 03-implementation
created: "2026-10-01T15:12:32Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9f896d467eee3262f207fb2941c527c83f78911370636f2e0bf86506edec6bbf
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`internal/rules` is 412 lines plus 320 of test: the two axes as constants, `Rule` with origin, level and
path supplied by the path so `scope` stays the file's own claim, `Check` keeping its type and parameters
as read, `Load` walking both trees and reporting what it could not use, `Effective` resolving by
specificity then origin with binding above both, and `Render` and `Hash` for A66's rendering. The
binding collision leaves the id out of the set and names every file that declared it. `rulesGate`
replaces `notImplemented` in the table and words both steps' problems into findings with file, cause and
next step; `SectionSet` now writes the hash where the harness wrote `by-hand`. A66 records the hash
definition, the empty case and why neither the placeholder rule nor G-Schema's recomputation is
tightened with it; A67 records the three findings added beyond the enumerated ones and the one
deliberately not added. Five deviations, the first worth reading: the design argued from first
principles for what A62 had already measured, and the register's state column is where that measurement
sits.
