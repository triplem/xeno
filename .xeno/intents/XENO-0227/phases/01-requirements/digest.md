---
intent: github.com/triplem/xeno#179
phase: 01-requirements
created: "2026-10-03T12:14:39Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4d472b6.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: fb0dee9dfb504fd16659ab4ece640d32f23d88d149d27f6dff84f482aad36a73
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`xeno intent start --for ISSUE` writes `intent.yaml` with section 5's seven fields and no eighth.
The qualified id is the host from `tracker.base_url`, the repository from `tracker.project` and the
key from `--for`, which also accepts any longer prefix of the id up to the whole of it, so a
repository with no tracker block — the block is optional — gives the id whole. `created` is the
runner's clock and the two version fields are the runner's own, so an intent reads the same version
string as the `gate.yaml` of its first phase, which is the half of #177 this closes. The key
continues the sequence on disk, `XENO-0227` after `XENO-0226`, padded as that key is, and `--intent`
names it instead where there is no sequence to continue. A second run refuses rather than
overwriting, because the file is inside `artifacts_hash`. Without `--for` it refuses and names the
flag. The seventy-nine hand-written intents are untouched and go on working, the listing shows the
new one like any other, and the next-step suggestion stops saying that no command creates an
intent. Nothing is added to Appendix A's tracker block and nothing to section 5: the host comes
from the only field that names a machine and the key prefix from the directory, because no field
names either, and where there is nothing to derive from the command refuses rather than guessing.
