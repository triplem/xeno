---
intent: github.com/triplem/xeno#336
phase: 00-intake
created: "2026-10-10T13:24:38Z"
schema_version: "1.0"
runner_version: dev+5f645cb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0e4d5df3d2510da70264484fa7fef8f5b571ad722364aff061be12d4c7ad5d61
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

Approved by @triplem on 2026-10-09T11:36:18Z: The mentioned project is just as n example implementation and not the foundation for this project.

> **github.com/triplem/xeno#336** — Section 4.1 of the evaluation does not name the labels-as-triggers design running on OpenHands
>
> Asked for on #323. `rajistics-demo/sdlc-automation-github-demo` is not an alternative to
> OpenHands, it runs on the OpenHands Automations API; what it adds is the GitHub side,
> labels as triggers (`openhands-context`, `openhands-build`, `openhands-review`,
> `openhands-qa`), status labels, issue and pull request templates as the approval
> boundary, evidence as comments and files on the pull request, and an `openspec/`
> directory for the change artifacts. No licence file, nine stars, one author, read
> 2026-10-08.
>
> It belongs in section 4.1 as one sentence, because it is the design section 12's
> "Issue commands are deliberately absent" paragraph refuses, running: something has to
> receive the label, and here it is a hosted automation. It strengthens the 4.1 decision
> rather than competing with it, and the sentence should say so and point at the
> paragraph. The nearer relation is to #330, where a label gates work before it starts
> as a precondition `xeno intent start` reads and not as a trigger anything listens for;
> that contrast is the sentence's point.
>
> ## Done when
>
> The sentence is in 4.1 with a Sources entry at a pinned commit. Nothing in section 12
> moves.
>
> Refs #323
>

The issue as it stood at 2026-10-10T13:24:38Z, read by `xeno phase start` and quoted rather than summarised. What this phase concludes about it belongs below.

<!-- xeno:section:scope -->
## Scope

One sentence in section 4.1 of `docs/orchestrator-evaluation.md`, and an entry in Sources
naming the repository it is about at a pinned commit. The issue fixes both the content and
the bound: the sentence says that `rajistics-demo/sdlc-automation-github-demo` is not an
alternative to OpenHands but a design running on the OpenHands Automations API, and it
points at section 12's "Issue commands are deliberately absent" paragraph, which that
design is the refused shape of. Nothing in section 12 moves.

The point the sentence has to make, in the issue's own terms, is that the demo strengthens
the 4.1 decision rather than competing with it. Something has to receive a trigger label,
and there it is a hosted automation — which is exactly the component section 12 declines to
build and operate. The nearer relation is #330: a label gates work there too, but as a
precondition `xeno intent start` reads off an issue it has already fetched, not as an event
anything listens for. That contrast is what the sentence exists to draw.

The claims were confirmed at the pinned commit rather than carried over from the issue, and
what they are is in the context rationale below.

What this does not do.

**It does not move section 12.** The paragraph the sentence points at stays as it is. A
trigger label would be a specification change before it was anything, and this intent writes
an evaluation sentence rather than opening that question. #338 holds the mechanism and the
condition it would have to meet.

**It does not make the demo a candidate.** Section 4 weighs candidates for the platform and
this is not one of them; it runs on the platform 4.1 chose. The sentence belongs in 4.1
beside the chosen candidate's properties and not in a section of its own, which is also why
this is one sentence rather than the treatment section 9 gives OpenSpec and #334 is to give
AI-DLC.

**It does not take any of the mechanisms.** Status labels, issue and pull request templates
as the approval boundary, evidence as comments and files on the pull request, and an
`openspec/` directory for change artifacts are each a design this repository has either
decided against or already has in another form. Naming them is the sentence's business;
adopting one would be an issue of its own.

**It does not revisit the OpenHands choice.** 4.1 is the decision and this adds a sentence
to it. The demo's existence is evidence that the platform is used this way by somebody else,
which is a reason to record it, not a reason to reopen anything.

<!-- xeno:section:context-rationale -->
## Why this context

Two files, 42,111 bytes, against a budget of 43,000 — the resolved figure plus the nine or
so lines this intent adds, because a budget measured before the work that enlarges the
files it measured is a budget the work exceeds by doing itself. XENO-0286 learned that the
expensive way and carried an advisory finding from P4 to its end for it.

**What changes.** `docs/orchestrator-evaluation.md`. The sentence goes in 4.1 and the
pinned entry in Sources, which is the form section 9's own Sources paragraph already uses
for OpenSpec: a repository, a tag or commit, the files read, and the date.

**What the contrast is read off rather than argued from.** `internal/model/identity.go`
holds `ApprovedLabel`, `ApprovedWord` and `Issue.Approval()`, which is what #330's half of
the sentence actually is. The distinction the sentence draws — a precondition read versus
an event listened for — is a property of that function: it is handed an issue the runner
has already fetched and answers whether both halves are present, naming what is missing.
Nothing subscribes, nothing is delivered, and there is no receiver. Reading it is how the
sentence can say that without overstating it, and the comment there makes the same point
about the label carrying the tool's name (#332) that the demo's `openhands-` prefix makes
about its own.

What is deliberately out, and why.

**`docs/process-definition.md`, at 138,216 bytes.** The sentence points at one paragraph of
section 12 and the file is thirty-three times the size of the document being edited.
Including it would be the whole of the budget for five sentences. The paragraph is quoted
verbatim in the problem section through the issue, and it was read at
`docs/process-definition.md:1698` and compared word for word with the issue's quotation
before this scope was fixed: they agree. "Nothing in section 12 moves" is then checked by
`git diff main -- docs/process-definition.md` being empty, which needs no context at all.

**The demo repository's own tree.** It is not in this repository and cannot be in a scope
of it. What was read from it is recorded here, which is also what makes the pinned commit
in Sources mean something.

**`.xeno/intents/**`**, because the trail is the record and not input.

**What was confirmed at `2ddf6c91791d95b8de8ff38b683241743ec12049`, dated
2026-10-07T03:58:00Z, read 2026-10-10.** Nine stars and no licence file, which the API
reports as `license: null`. The four trigger labels are in `.github/labels.json` with
descriptions — `openhands-context` "build a cost-aware context reuse report",
`openhands-build` "clarify a sparse issue and create a PR", `openhands-review` "automated PR
review", `openhands-qa` "QA and test generation" — beside four status labels,
`openhands:ready`, `openhands:in-progress`, `openhands:needs-human` and `openhands:done`.
The receiver is the Automations API and not a workflow: each of the four has an
`automations/github/<name>/automation.prompt-preset.json` with a `prompt.md`, registered by
`scripts/automations/register_github_automations.py`, and the tree holds no
`.github/workflows` directory at all. The templates are
`.github/ISSUE_TEMPLATE/story-openhands-build.md` with a `config.yml` and a
`pull_request_template.md`. `openspec/` holds `project.md` and
`changes/<name>/{proposal,design,tasks}.md` with a `specs/` beneath it.

That the tree has no workflows is the detail that makes the sentence's point precise rather
than approximate: there is nothing in the repository that could receive a label, so the
receiver is somewhere the repository cannot show you.
