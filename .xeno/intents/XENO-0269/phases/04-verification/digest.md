---
intent: github.com/triplem/xeno#207
phase: 04-verification
created: "2026-10-07T09:11:42Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 72a703cf70aade44db4a52257938af36f5074f7a810042446b34d4ab35329d97
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Twelve criteria, all passing: eight read by a person, because the deliverable is prose, and
four checks.

The three writers were read in the code rather than recalled. `r.ToolVersion` comes from
`XENO_HARNESS_VERSION` and `--tool-version`, and `recordedToolVersion` copies it back out of
the artifact for the digest. `agent()` reads `tool` from `project.yaml` with `XENO_HARNESS` over
the top, and `model` from `project.yaml` alone, which is what the paragraph now says and the
approved draft did not.

One qualification to criterion 3, found while collecting the evidence and with the positive
control that made it visible: `hashes` reads `tool == "manual"` and accepts the `by-hand`
placeholder in every hash field on the strength of it. Not corroboration, but a check relaxed
on a declaration, and a sharper instance of #207's concern than the register is. The paragraph
stays true; the fact is outside #207's done-when and outside the approved wording, so it is
raised rather than absorbed, with the learning record carrying the general form.

`go test ./...` passes across 20 packages, build, `gofmt` and `vet` clean, `xeno gate verify`
matches 481 verdicts, and the diff against main is one document and no Go file.

The gaps say the rest. The gap is stated and not closed. `tool: manual` relaxes a check and
this intent does not fix it. Nothing checks that the paragraph stays true, by design and said
so. The gateway claim is inherited from the plan's measurement rather than re-measured. And the
approved wording was wrong once, caught by a reading and by nothing else.
