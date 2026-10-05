---
intent: github.com/triplem/xeno#205
phase: 00-intake
created: "2026-10-05T16:22:24Z"
schema_version: "1.0"
runner_version: dev+97d07da.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 703eded3d3a399750676b4a085fb3d290ec17617c83883f93e60b96624e3a1a5
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

Section 7 lists `XENO_PLUGIN_DATA` in the normalised environment:

```
XENO_PLUGIN_DATA      local data location, always .xeno/local/
```

`.xeno/plugin/bin/xeno-env.sh` sets and exports it. Nothing reads it. Six places in the runner
write the path as a literal instead:

| where | what |
|---|---|
| `internal/runner/runner.go:89` | the run marker, `.xeno/local/runs/KEY/PHASE.lock` |
| `internal/runner/runner.go:95` | `.xeno/local/phase.env` |
| `internal/runner/enforcement.go:18` | `ReportPath`, `.xeno/local/enforcement.yaml` |
| `internal/cost/cost.go:30` | `Ledger`, `.xeno/local/cost-ledger.yaml` |
| `internal/runner/init.go:150` | the `.gitignore` entry |
| `cmd/xeno/main.go:852` | `phase.env` again, written out a second time |

It is the last entry in that list with no reader. `XENO_PLUGIN_ROOT` left the list under A89,
and `XENO_HARNESS` and `XENO_HARNESS_VERSION` are read by the runner under A84, so this is the
one remaining case of a variable that is specified, exported and inert.

A84 records it, so it is a known gap rather than a surprise. The clause audit's point is sharper
than that: a reader cannot tell a variable with no reader from one whose reader is a constant,
and here the constant is written out six times, once inconsistently — `cmd/xeno/main.go` repeats
the `phase.env` literal rather than calling the runner's own accessor, so two places already have
to agree about a path that nothing defines in one place.

So the export is a promise the runner does not keep. A person who sets `XENO_PLUGIN_DATA`
expecting local state to move gets no error and no effect.

<!-- xeno:section:scope -->
## Scope

In scope is one resolver and six call sites. `model.LocalDir(root)` returns the local data
location: `XENO_PLUGIN_DATA` where it is set, `.xeno/local` under the repository root where it is
not, and an absolute value is taken as given rather than joined to the root.
`model.LocalPath(root, parts...)` builds a path inside it, and the six literals become calls.

In scope is `internal/model` as the home, because it is the only package all six callers already
import and because it imports nothing but the standard library, so `internal/cost` can reach it
without a cycle.

In scope is removing the duplicated literal. `cmd/xeno/main.go` writes `.xeno/local/phase.env`
and so does `internal/runner/runner.go`, and two places agreeing by hand about a path neither
defines is the smaller half of this defect and the half most likely to drift.

In scope is the `.gitignore` entry staying `.xeno/local/`. It is about this repository, not about
where a person redirected their own state: a relocated directory outside the tree needs no entry,
and one inside it is the default the entry already covers. Worth stating because it is the one
site that deliberately does not change.

In scope is a test over the resolution: unset gives the default under the root, a relative value
is joined to the root, an absolute value is used as it stands, and the run marker, `phase.env`,
the enforcement report and the ledger all land inside whatever the resolver returned.

Out of scope is anything hashed. Nothing under this directory is covered by `artifacts_hash` —
it holds the run marker, `phase.env`, the cost ledger and the enforcement report — which is why
reading the variable cannot make a verdict depend on the environment. That is the objection A89
raised against the plugin root and the reason it does not apply here, and the issue asks for it
to be stated rather than assumed.

Out of scope is `XENO_PLUGIN_ROOT`, which A89 removed from section 7 and which this intent does
not reopen.

Out of scope is a flag. Section 7 normalises through the environment and the entry point already
exports this one; adding `--local-dir` would be a second way to say the same thing and an
addition to the command surface.

Out of scope is migrating anything. A person who sets the variable gets a fresh directory; the
runner does not move or look for existing state elsewhere, and nothing under there survives a
`retention` window anyway.

No normative document is touched. Section 7 already specifies the variable and this intent makes
the runner read it, so the code moves towards the specification and no specification commit
precedes this.

<!-- xeno:section:context-rationale -->
## Why this context

The input is the clause, the six literals, the entry point that exports the variable, and the two
register rows about the environment.

`docs/process-definition.md` is read for section 7's own line rather than for the issue's
quotation, because the words "always `.xeno/local/`" are the whole of the ambiguity. They can be
read as a prohibition — the value is always that, so reading the variable could only ever produce
one answer — or as a description of what the entry point normalises to, so that the runner need
not know a harness's conventions. The second is what the section is for: it is headed
"Environment normalisation" and says the runner "only ever sees `XENO_*`". Reading the heading
settles it, and a decision taken off the table alone would have gone the other way.

`docs/assumptions.md` is read for A89 and A84. A89 is the precedent that matters and the one this
intent has to distinguish itself from: the plugin root was removed from section 7 rather than
read, because `rules_hash` and a rendered artifact resolve from that tree and an override would
make a verdict depend on the environment. The measurement in that row is specific — a valid rule
tree that differs leaves `gate verify` at exit 0 while G-Policy silently stops judging. Nothing
under `.xeno/local/` is hashed, so the same argument does not reach this variable, which is the
fact this intent rests on.

`.xeno/plugin/bin/xeno-env.sh` is read for what is actually exported: `: "${XENO_PLUGIN_DATA:=$root/.xeno/local}"`,
an absolute path built from the git root. So the value the runner will see in practice is
absolute, which is what decides that the resolver must take an absolute value as given rather
than join it to the root — a detail invisible from the issue and from section 7.

The four Go files are read for the literals and for one surprise: `cmd/xeno/main.go` repeats
`.xeno/local/phase.env` instead of calling `phaseEnv`, so the duplication is already there and
this intent removes it rather than adding a resolver on top of it.

`internal/cost/cost.go` is read for its imports, because the resolver's home depends on them. It
imports only the standard library and `yaml`, so it can take `internal/model` without a cycle;
had it imported `internal/runner`, the resolver would have had to live somewhere else.
