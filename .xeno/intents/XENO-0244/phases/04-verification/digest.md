---
intent: github.com/triplem/xeno#221
phase: 04-verification
created: "2026-10-04T09:41:10Z"
schema_version: "1.0"
runner_version: dev+24becc3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 312534f21640d4a3d373eb4d7dd4f6cbc12f49e126442156f4dadb57823993b8
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Verification maps fourteen criteria to six new tests, to the diff and to one figure
taken three times. `xeno gate verify` is at exit 0 over 342 verdicts before the change,
after the fifteen corrections and after the gate, with the verdict count identical and
every one of the fifteen phases keeping its `artifacts_hash`. That is the reading of
Appendix B confirmed rather than assumed — `artifacts_hash` covers the files lying
directly in a phase directory, so `evidence/attached.yaml` is outside it — and the
middle reading is the one that proves it, because a wrong reading would have reported
fifteen divergences at once. The correction is fifteen files, fifteen insertions,
fifteen deletions, each a `result:` line, with `pipeline: local` and every `sha256`
untouched and G-Evidence resolving each report afterwards. Eighteen packages ok, build
clean, `gofmt` and `go vet` silent. The refusal was run against this phase's own
declarations with a manifest that published one entry with `result: success` and one
with a uri and no hash: both were declined with their own remedy in their own sentence,
exit 1, and nothing was written — the phase had no `evidence/` directory afterwards and
all four declarations stayed pending. Those four were written by the command, which
makes this the second intent to declare evidence and the first to do it as a habit
rather than as the thing being built, so the phase finishes `provisional` and the
pipeline resolves it. Six gaps are recorded rather than absorbed. An attachment with no
result at all on a kind that needs one is still unjudged, which is D-1's named exclusion
and is filed; G-Test is still `not-implemented`; the fifteen still point at transcripts
somebody typed, with a provenance that is accurate and stays; a manifest corrected in
the artifact store still needs the attach run again by hand; `gate run` is the third
caller of `Attach` and says nothing about a declined entry, which is a line of output
and not this intent's scope; and the correction was a hand edit, which leaves a
precedent whose only safeguard is the three readings above. Read in this phase: the diff
of the fifteen, both new test files, and the four workflow manifests whose entries this
phase declares.
