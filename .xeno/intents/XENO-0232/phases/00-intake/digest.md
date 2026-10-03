---
intent: github.com/triplem/xeno#195
phase: 00-intake
created: "2026-10-03T14:19:32Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+b54626e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: dd6466b301c625dece60e0a2a0f287bcd31c658ce077a7e74e5f711bab22c72d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`learning.yaml` was the last artifact of this process that no command wrote, and one phase directory
shows the cost: `output.md`, `digest.md` and `gate.yaml` record `0.1.0-dev+74553ad.dirty` and
`learning.yaml` records `0.1.0-dev`, because a person typed it. Four artifacts, one phase, one build,
two strings — which is #179's sentence about `intent.yaml` needing no amendment to apply here, and
`learning.yaml` was not in that issue's scope. The file is inside `artifacts_hash`, since Appendix B
excludes only `gate.yaml` and `cost.yaml`, so the typed value is sealed with the verdict and thirty
intents' worth of records are in that state. It is also the artifact the process cares most about
keeping honest, being the input to section 10's route through review. In scope: a command taking the
four keys as arguments with the header from the runner, the phase optional so the intent level record
has the same writer, the empty case said rather than left out, entries accumulating, and the category
and keys refused before anything is read. Out of scope: the content, because section 10 routes a
learning through review so that it does not take effect on being noticed and a runner that drafted it
would be proposing rules; `plugin_version`, which is a constant that does not even match the
manifest and is #177; backfilling the sealed records; and a sealed-phase guard, which would refuse
the first half of a re-judgement. The maintainer found this by asking why the exported harness
version was not in the frontmatter — it correctly is not, since section 5 scopes `tool_version` to
the two files produced in a session, and what was wrong was the field beside it.
