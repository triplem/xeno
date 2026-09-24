---
intent: github.com/triplem/xeno#28
phase: 00-intake
created: 2026-09-24T19:51:40Z
schema_version: "1.0"
runner_version: 0.1.0-dev+fce3d0f.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 3e1e03503c799ea8ce4ac52b7e63e339378745b398e28441fa4f7c5875cbe91d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
---

# Intake

Issue #28. The merge of #27 ran no workflow, because its description carried a marker
that switches CI off, quoted as an example of a message the new checker accepts.

## Scope

A check that fails a pull request whose title or description carries one of those
markers, before the merge rather than after it, and a sentence in `CONTRIBUTING.md` and
the pull request template saying that a description is a commit message.

## Non goals

Recovering the release that did not happen. semantic-release derives from the history
since the last tag, so the `feat:` in that merge is picked up by the next push to `main`.
Nothing is lost; it is deferred.

## A constraint this intent has to obey itself

Nothing written here may quote a marker in prose, since quoting one is what caused the
problem. The authoritative list is the check, and everything else points at it. This
intake, the commit that carries it and the pull request that merges it all have to hold
to that, or the work switches off the run that would have verified it.

## The first artifact with a stamped version

Not scope, but worth noting where it happened: this phase's `context.lock.yaml` is the
first to carry a runner version that identifies a build, `0.1.0-dev` followed by the
commit and the state of the tree. Before #16 every artifact said the same thing.
