---
intent: github.com/triplem/xeno#151
phase: 00-intake
created: "2026-10-01T13:52:27Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c54db93.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8ea5fe383f0ff29112adcff6c0877726977491462547ab619636ce789ca65f17
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
#149 made the format the deliverable: Xeno ships no indexer, so what it fixes is the shape of the answer. Nothing
in the tree reads an index or knows what one looks like, and index.path is a field the specification carries and
the runner has never looked at — the same state index.grammars_dir was in before #149 replaced it. Three
requirements the tree cannot express: the content, five fields per symbol and explicitly no call relationships;
the provenance, which #149 made a requirement because an index that cannot say how old it is cannot be judged;
and the degradation, which both documents call the property everything else follows from, and whose first code
path is the one where there is nothing to read. The acceptance also needs this repository to produce an index for
its own Go source, which is this project being a project and not Xeno shipping an indexer.
