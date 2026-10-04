---
intent: github.com/triplem/xeno#217
phase: 02-design
created: "2026-10-04T20:05:58Z"
schema_version: "1.0"
runner_version: dev+fbf8a72.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 55d12062332bcb3a6bdd4151b20f9d59091413375abeac7e56a095ed0a59671e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Design settles five things and names one gap. The rename goes first and is mechanical:
model.ContextProfile and model.Profile become ContextScope and Scope with five readers
following, this intent's own artifact is renamed and P0 re-judged, and docs/v2-delta.md
is deliberately left alone as a different document. The writer is xeno scope set, with
the grammar of section set, which it most resembles because a person supplies content
and the runner supplies provenance. It takes no --phase, since informationBase reads the
scope from P0 and nowhere else, and a flag would offer a choice the format does not
have. It reads its entry whole rather than from flags, decoded into model.Scope with
KnownFields(true), which is RecordQuestion's shape and its argument: a nested structure
does not fit flags without inventing a separator, and the shape read is then the shape
written. The header fields come from r.common and never from a person, by A35. include
is required, because a scope that resolves to no files is indistinguishable from an
absent one in every reader, which is the state this intent exists to end. The resolution
is extracted so the writer and phase start call one function and the figure reported is
the figure recorded; the writer prints the count and the byte total, which is the
criterion P1 took from this intent's own P0, where the figure was established by hand
and the first budget was wrong. The refusal is the first statement in Finish, before the
digest and the cost are written, so a refused finish writes nothing, and it is not in
phase start of P1 because a P0 can only be left scopeless if it was allowed to finish.
Section 5's second limit is the loop bound in staleReads, i < idx instead of i <= idx,
one character and the whole of what unblocks this intent. Its first limit narrows to the
commit range through git.Paths, using the Base and Head the gate context already
carries. The gap is named rather than hidden: with no commit range there is no change
under review, git.Paths treats an empty range as a caller error because section 12
forbids inferring one, so the staleness half reports nothing locally and becomes a CI
check in practice. That is a silent pass of exactly the kind this intent was opened
about, and it cannot be reported as anything else until a check that did not run has a
result, which is #235. Nine alternatives are recorded with their reasons, among them the
two that a later reader will reach for first: putting the requirement in G-Schema and
releasing the 100, which is Q-2's rejected third option and whose arithmetic the design
did not change, and inferring the commit range so the staleness half would work locally,
which section 12 forbids and which would make a verdict depend on the branch the runner
happened to be on. Impact: no new dependency, no document change, 347 verdicts at exit 0
before and after, the 100 sealed P0 phases untouched because the requirement lives in a
command, and one verdict affected, this intent's own P1, where limit 2 removes the
reason the two releases were needed. D-3 records the widening and the cost it was paid
with.
