---
intent: github.com/triplem/xeno#106
phase: 00-intake
created: "2026-09-27T14:31:59Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0b3a547.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: a618da7e7909a912dd9b300251b714d0ae402318dab6113756a1144a6a07f198
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

`Attach` copies a manifest entry's `uri` and `sha256` unconditionally, so an entry
published with a uri and no hash becomes an attached record carrying an empty one.
G-Evidence then passes it: its check takes no action where the path is empty and the uri is
not, which is right for an item bound by a hash and wrong for one that has none.

A8 says an item with a `uri` is bound by its declared hash alone, because resolving a uri
needs a network the gate path never has. An attachment with no hash is bound by nothing.
Section 4 names the cost of attaching outside the `artifacts_hash` and names what pays it:
`gate.yaml` records the hashes it judged, so altering an attachment afterwards makes it
disagree with the verdict. With no hash there is nothing to disagree with, and the verdict
records that evidence arrived without being able to say what arrived.

An entry with no file, no uri and no hash is the same defect with less to look at, and it
produces an attached record naming a result and nothing it belongs to.

<!-- xeno:section:scope -->
## Scope

Two halves, because two different writers can produce the state.

`Attach` refuses an entry it cannot bind. The item stays pending, which is what it was, and
the reason is carried out rather than swallowed: `Attach` returns a result naming each
entry it declined, `evidence attach` prints them, and the refusal a `phase start` gives on
a still provisional predecessor says which entry the pipeline published wrong.

G-Evidence reports an attached record with no hash. `evidence/attached.yaml` lies outside
the `artifacts_hash`, so it is the one file in a judged phase that a person can edit
without invalidating anything, and a gate that trusts its writer checks nothing.

Not the network. A uri is still resolved by nobody, and this changes what binds an item
rather than what checks the bytes behind it.

<!-- xeno:section:context-rationale -->
## Why this context

**Refused at the attach, because that is where the entry is still identifiable.** An item
recorded and then found unjudgeable has lost what would explain it: the gate sees a record
with no hash and cannot say whether the pipeline published it wrong or somebody edited the
file. Declining to record it keeps `attached.yaml` a list of what was attached rather than a
list a gate is expected to distrust, and leaves the phase in the state it was already in,
provisional and waiting, which section 6 calls honest when nobody has carried on.

**And reported at the gate anyway, because the file is outside the hash.** That is not
redundancy. Every other file a verdict judges is covered by `artifacts_hash`, so editing it
makes the verdict stale and CI finds it. `evidence/attached.yaml` is deliberately not, which
is the whole arrangement that lets a later writer touch it without breaking a seal, and the
price is that this one file can change under a verdict without anything noticing. The gate
check is what the missing seal is replaced by. `CompleteOnClose` already makes this argument
for the abandonment reason: the command requires one, and the gate checks it anyway, because
the command is not the only way a file gets written.

**The reason is carried, not printed where it happens.** `Attach` writes no output; it is
called from `phase start`, from `gate run` and from its own command, and a package that
printed would be deciding for three callers. So it returns what it declined and each caller
says it in its own voice. Without that the item stays pending with no explanation anywhere,
which is the failure mode the issue is really about: not a wrong verdict, an invisible one.

**What stays broken on purpose.** Nothing checks that the hash matches the bytes behind a
uri, and nothing can without a network. Limitation 11 of section 16 already says the trail
is as trustworthy as the repository, and A8 is unchanged: the declaration binds, and whether
the content behind it still matches is outside what a local runner can answer.
