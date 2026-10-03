---
intent: github.com/triplem/xeno#171
phase: 00-intake
created: "2026-10-03T09:42:55Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e9bd2f93f760341fa7dc3e3041d0e9c62c6f10bbde0d79a20aff4a7dcb3e2388
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

WP8's done-when has two clauses and one of them has no mechanism. "Every read outside the profile
appears in `context.lock.yaml`" holds, in the only sense section 5 allows: the lock states what the
phase was given, nothing stops an agent reading further, and the section says treating the file as
a measurement of what was read would be wrong. "A repeated phase reads only what changed" has
nothing behind it at all. Nothing computes what changed, nothing reports it, and a phase that runs
again is told the same thing a phase that runs for the first time is told.

**And the lock does not carry what section 5 says it carries.** The section writes the file out
with `repo_commit`, `files`, `rules_applied`, `template_source`, `plugin` and `tools`. The tree
writes `files`, `template_source`, the common header, and two fields section 7 needs —
`predecessor_artifacts_hash` and `evidence_source`. Four of the enumerated fields are absent.

Two of those four could not have been written until this week. `rules_applied` needs a rule set,
and `given/builtin/` was empty until #165; `repo_commit` needs git, which came into the tree with
#162. So the gap is not an old oversight in both cases — it is a field that became writable and
nobody went back for it. `rules_hash` says that a set was in force; `rules_applied` says which
rules, with the version counter each carried, which is the difference between a hash somebody can
recompute if they hold the tree and a list somebody can read.

**Three things the profile specifies and nothing reads.**

*The budget.* Section 5 says G-Schema "reports a finding where the recorded context exceeded the
budget", deliberately a finding rather than a red gate, and then says what the check is for: "what
it prevents is the budget quietly becoming decoration". `Profile.Budget` is in the model, no gate
reads it, and the sentence has come true.

*The links.* "Links between code and documentation are declared in the context profile, not
inferred. An inferred mapping is an assumption, and assumptions in this process are either
registered or absent." The field is parsed and contributes nothing: a profile declaring that a
component is documented by a particular file does not put that file in the phase's information
base unless `include` happens to match it, so the declaration is ornamental.

*The assembly order.* "Context is assembled in order of volatility, not in the order it was
thought of... which is why `context.lock.yaml` records the assembly order and not only the set."
The runner sorts `files` alphabetically. That is deterministic, which the hash needs, and it is not
an order of volatility, which the cache does.

**None of it is exercised here, because this repository has never written a profile.** Twenty-three
intents, no `context-profile.yaml`, so `files` is empty in every lock and the second half of
G-Freshness has been told nothing in every run. That is the position the rule engine was in before
#165 and the external gate invariants were in before #167: a mechanism whose verdicts nobody has
seen.

<!-- xeno:section:scope -->
## Scope

**In scope.** `repo_commit` and `rules_applied` in the lock, with each rule's own version counter.
The budget as a G-Schema finding, reported and not blocking. A declared link's document in the
information base whether or not `include` matches it. The `files` order taken from the order of the
profile's `include` patterns, so the project decides volatility and the lock records what it
decided. And the first clause: `xeno phase start` naming the files of the base that changed since
the preceding phase's lock, derived from the two locks rather than recorded in either, because
section 5 enumerates the lock's fields and a "what changed" field is not among them.

**Out of scope, and each for its own reason.**

`plugin: { version, sha256 }`. It is G-Supply's input — section 5 says the frontmatter names what
was used and the lock proves it with a hash, and that where the two disagree G-Supply fails. The
gate is unimplemented, and a field written for a check that does not read it is a field nobody
maintains.

`tools`. WP15's remaining item, settled by #155 and waiting for the writer.

Blocking on the budget. The section says in as many words that it is a finding and not a red gate
in the sense of stopping work, because "blocking against a number nobody has experience with yet
would be the wrong way round".

Inferring a link. The specification forbids it in the same paragraph that declares the field.

**And deliberately not: writing a profile for this repository.** With one in force, the second half
of G-Freshness becomes live, and a file an earlier phase was given that a later phase changed is a
finding against the earlier phase. That is the mechanism working, it is what "a repeated phase
reads only what changed" is for, and releasing such a finding takes a person's decision. An agent
adopting it here would be an agent writing the findings and approving them. The mechanism is
demonstrated against a copy of this repository with a profile written for it, the cost is measured,
and the adoption is a decision with figures attached.

<!-- xeno:section:context-rationale -->
## Why this context

Section 5 is the specification for this piece twice over: the context economy subsection for the
profile, the budget, the links, the volatility order and the sentence that re-reading follows
change; and the paragraph that writes `context.lock.yaml` out field by field, which is how the four
missing ones were found. Read entire rather than by subsection, because the two halves are the same
mechanism seen from the file and from the phase.

The three sentences that decided the design are all in that section. "It records the context that
was declared, not everything that was read" is why the first clause is a report rather than a
record. "`context.lock.yaml` already carries paths with hashes, so a repeated phase knows which
files changed and reads only those" is the mechanism, and it says the inputs are the locks. And the
budget's "deliberately a finding and not a red gate" fixes the verdict before any code is written.

WP8 of the plan is read for the done-when and for the sentence that this half needs no baseline,
only hashes that already exist — which is what makes it implementable now and measurable only by
WP15.

`internal/runner/runner.go` is read for `informationBase`, which resolves the profile, and for
`Start`, which writes the lock. Both are short and both are where every change here lands.

`internal/model/model.go` is read for `ContextLock`, `ContextFile` and `Profile`, and the reading
established what is parsed and never used: `Links` and `Budget`.

`internal/gates/gates.go` is read for G-Schema's handling of a phase's files and for
`staleReads`, which is the second half of G-Freshness and the thing a profile would make live. Its
comment already says it has been told nothing in every intent so far.

`internal/git/git.go` is read for the one thing it does, because `repo_commit` needs a commit and
this project starts one subprocess for git and has a package for it.

`internal/rules` is read for `Effective`, since `rules_applied` is that set with each rule's path
and version rather than its hash.

The locks of earlier intents are read as data: what the field order looks like today, and what a
reader of one actually learns from it, which is less than section 5 intends.
