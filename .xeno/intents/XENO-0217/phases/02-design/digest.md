---
intent: github.com/triplem/xeno#158
phase: 02-design
created: "2026-10-01T15:04:29Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1834952b8c2b2dc20d54a586d7c7d7c5acc83231bf75dedb8f5321ef482dd5c5
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`internal/rules` returns the set and its problems and lets the gate word them, which is A61's shape and
matters because three callers are coming: G-Rules, G-Policy and the hash writer. The effective set holds
only well formed rules, so a rule whose `scope` contradicts its path is reported and left out rather
than entered into a precedence contest it should not win. A binding collision removes the id entirely:
section 9 forbids resolving it, and that refusal is the one behaviour the code would get wrong if nobody
had read the sentence. `rules_hash` is a sha256 over a canonical rendering of the effective set, one
tab-separated line per rule sorted by id, semantic fields only, exported so the test can pipe it through
`sha256sum` — A62's shape for `secrets_hash`. An empty set hashes the empty rendering, because the field
is required and G-Schema accepts only a sha256 or the placeholder, so the digest's option of leaving it
out is not available. The placeholder stays accepted and the runner stops writing it, since tightening
the gate would turn seventy-one sealed intents red to tidy a field only the runner writes now. Validated
is what sections 7 and 9 name, plus parse failures, missing format fields and a duplicate id at one
level; `applies_to` against the known phases is deliberately not checked, because an unenumerated check
is a spec change first.
