---
intent: github.com/triplem/xeno#151
phase: 04-verification
created: "2026-10-01T14:01:29Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c54db93.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cf2e58665419a3d8807a64acbecdee845b5a719da7308383c08335de623fcdc4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Seven criteria in full and one in half. Build, vet, test, gofmt clean; gate verify 168 verdicts, exit 0; go.mod
unchanged; the gate path reaches no net, no net/http and no internal/index; coverage 96.3%. The producer emits 997
symbols for this tree and four line numbers were checked by hand against their definitions. The demonstration is
Requirements, which has five definitions across four files — a struct field, three methods on two types, an
interface method — distinguishable only by the container field, which is why the producer emits contained symbols
although the acceptance did not ask. Criterion 2 is half met: the shape is annotated and machine checked in
testdata and published nowhere, which is WP16's. The gap that matters is that nothing consumes the index, so the
format is unproven where it will be judged; the second is that provenance is trusted entirely, since produced_at
is whatever the producer wrote and staleness is the one judgement Xeno makes about the index.
