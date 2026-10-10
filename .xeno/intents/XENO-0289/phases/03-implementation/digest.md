---
intent: github.com/triplem/xeno#359
phase: 03-implementation
created: "2026-10-10T15:59:23Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9584993f198902de4c50c852a8b54d0aa16e2ac7beb9edb8ddfd5b8508e010da
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
`docs/labels.md`, 193 lines and 12,146 bytes, with one entry in `docs/README.md` under
"Running it" and one line in `zensical.toml`'s nav after "Commands". No code changes and
neither normative document is touched. The page opens by saying that no label holds an
intent's state, that the state is in `.xeno/intents/` and is what `xeno intent status`
prints, that nothing in Xeno ever writes a label because the adapter contract is three
operations, and that one label is read anywhere; then the 33 labels the host carries, in
three groups ordered by how binding they are, each group a five-column table answering
#359's questions and prose carrying what a table cannot. The name of
`xeno-needs-decision` is settled by the maintainer as the plural, D-2 of this phase, and
the paragraph recording it is the page's worked example of its own purpose: no code read
either spelling, so the cost of a near-miss is a person filtering on the wrong one. Three
open points remain in the page with their issues rather than left as silence: the
lifecycle of `xeno-approved`, whether a fresh approval is owed, and triggering, which is
#338's. The operational paragraph prescribes the read-back and reports that the
carried-in defect did not reproduce, rather than recording it as a defect. Two figures
stated from the session rather than read from the host turned out wrong — the label count
in a sealed criterion, and the list of issues in the intake's decision record — and both
are reported rather than quietly corrected. This phase was judged once before the
maintainer's answer arrived and was redone with it, which is why the verdict and the page
match. The gates are green, the strict site build finds no issues, and the page is in the
built navigation.
