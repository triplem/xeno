---
intent: github.com/triplem/xeno#336
phase: 00-intake
created: "2026-10-10T13:27:06Z"
schema_version: "1.0"
runner_version: dev+5f645cb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0e4d5df3d2510da70264484fa7fef8f5b571ad722364aff061be12d4c7ad5d61
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
The intake of #336, which asks for one sentence in section 4.1 of
`docs/orchestrator-evaluation.md` and a Sources entry naming the repository it is about at a
pinned commit. The issue fixes the content, the bound and the done-when, so there is nothing
to decide and no question to ask.

What this phase added to the issue is confirmation. Everything #336 claims about
`rajistics-demo/sdlc-automation-github-demo` was read at
`2ddf6c91791d95b8de8ff38b683241743ec12049`, dated 2026-10-07, rather than carried over: nine
stars, `license: null`, the four trigger labels in `.github/labels.json` with their
descriptions, four status labels beside them, the four
`automations/github/<name>/automation.prompt-preset.json` presets with their `prompt.md`
files and the script that registers them, the issue and pull request templates, and
`openspec/` with `project.md` and a change directory under it.

One detail was found that the issue does not have, and it is the one that makes the
sentence precise rather than approximate: **the tree holds no `.github/workflows` directory
at all.** So there is nothing in that repository which could receive a label, and the
receiver is somewhere the repository cannot show you. That is section 12's paragraph
exactly — "something has to receive them, and every way of doing that is a component to
build and operate" — observed rather than inferred.

The section 12 paragraph was read at `docs/process-definition.md:1698` and compared word for
word with the issue's quotation of it; they agree. The file is then deliberately out of
scope at 138,216 bytes, thirty-three times the document being edited, for one paragraph that
is already quoted in the problem section. "Nothing in section 12 moves" is checked by a
`git diff` and needs no context.

`internal/model/identity.go` is in scope for the other half of the sentence. #330's label is
a precondition and not a trigger, and that is a property of `Issue.Approval()`: it is handed
an issue the runner already fetched and answers whether both halves are present, naming what
is missing. Nothing subscribes and there is no receiver. Reading it is how the sentence can
draw the contrast without overstating it.

Two files and 42,111 bytes against a budget of 43,000, which is the resolved figure plus the
nine lines this intent adds. That is XENO-0286's learning applied: it set its budget from the
resolved figure alone and carried an advisory finding to the end with no repair available.

Three sections, no open question, no decision.
