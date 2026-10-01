---
intent: github.com/triplem/xeno#160
phase: 02-design
created: "2026-10-01T15:33:54Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e32006ffc4ee12de0147fd7cc798bf337f7324399be4527fbbd148a3c78946d3
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Completeness runs from the rules to the entries, because section 12's clause about a lens only holds if
the counted set is the rules and section 9's two sentences only hold if a missing entry is a finding — a
gate walking the entries would pass an empty checklist. Every review rule in the effective set is
answered at P5 whatever its `applies_to`, since the checklist exists once and section 9 renders it from
the set; the alternative reading would leave a review rule about design answered nowhere, and the choice
is recorded in the register because the section supports both. A checked rule is reported and not
evaluated: `internal/gates` holds a predicate registry, empty here, and a type not in it is a finding
naming rule and type rather than a pass — the silent pass is the claim section 16 catalogues. The
registry sits with the gate because a predicate reads an artifact or a commit range and the resolver
reads neither. The entry is a `review_checklist` frontmatter list with `rule`, `result`, `note` and
`source`, keyed on the absence of `rule` rather than on `source`, so a lens that forgot the field cannot
answer a rule. An entry naming a rule outside the set is a finding, which is the one finding section 9
does not write and is recorded as such.
