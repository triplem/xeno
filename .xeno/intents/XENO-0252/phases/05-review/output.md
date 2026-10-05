---
intent: github.com/triplem/xeno#205
phase: 05-review
created: "2026-10-05T16:53:03Z"
schema_version: "1.0"
runner_version: dev+9eec4ed.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f6e1917fc60243e475f2fe47872e66768cd68a19853d81ec2ad6862b3dfbec83
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Three, each naming what it departs from. Against P2 design: printEnforcement took a new parameter, because a constant needs no context and a function does, and the design listed the files holding the literal rather than the places reading the constant — which was the only compile error. Against P2 wording: Ledger and ReportPath each gained a resolver function beside the file name rather than every caller calling model.LocalPath, since two callers each would have repeated the name. And one defect in this intents own record, in P4 results: P3 learning lost a word to a shell command substitution, left rather than repaired because the phase was judged and the choice was rewriting a sealed artifact over a typo.'
      result: deviation
      rule: deviations-are-traceable
    - note: 'Two exported identifiers change shape and the migration note is that nothing needs migrating. cost.Ledger and runner.ReportPath were exported path constants; they are now LedgerFile and ReportFile, file names, with LedgerPath(root) and ReportPath(root) beside them. Anything outside this repository reading the old constants as paths breaks at compile time, which is the good failure, and nothing does — the only caller outside their own packages was the command that prints one. For a person rather than a caller: setting XENO_PLUGIN_DATA now moves the run marker, phase.env, the enforcement report and the ledger, and leaving it unset keeps every path exactly where it was, asserted twice.'
      result: deviation
      rule: interface-change-needs-a-migration-note
    - note: 'No external dependency: go.mod is untouched and the resolver uses os, path/filepath and strings. But one internal dependency is added and is worth the note: internal/cost now imports internal/model, where before it imported only the standard library and yaml. It was weighed in P2 — internal/model imports nothing but the standard library, so there is no cycle, and the alternative was putting the resolver in internal/runner, which would have made cost depend on the runner and run the dependency the wrong way, since cost is called by the runner and by the hook.'
      result: deviation
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No normative document is touched: section 7 already specifies
`XENO_PLUGIN_DATA` and this intent makes the runner read it, so the code moves towards the
specification and no specification commit precedes it. Nothing is invented — the variable, its
meaning and the default are section 7's, and the one thing that would have needed a document
change was removing the clause, which was the option not taken. The branch carries one intent,
the commit references #205, and the issue carries `wp7`.

The acceptance criteria. Ten met, one met by the commit this phase precedes. Criterion 7 is the
one that protects every existing repository and it is asserted twice, over the resolved strings
and over the disk after a real `phase start`.

The non-goals held. No `XENO_PLUGIN_ROOT`. No flag. No migration of existing state. No change to
`xeno-env.sh`, to section 7, or to what the variable means. No validation of the value. The
`.gitignore` entry stays at the default, which criterion 5 required and which is the one site
that deliberately does not resolve.

What a reviewer should weigh is the reading of section 7, because everything rests on it. The
clause says "local data location, always `.xeno/local/`", which on its own argues against reading
the variable at all. What settles it the other way is the section around it: it is headed
"Environment normalisation" and says the runner "only ever sees `XENO_*`", so the value is what
the entry point normalises to rather than a value nothing may change. A97 records the reasoning
and P0's learning records that the clause and its context pointed in opposite directions.

The second thing to weigh is why this variable and not the other. A89 removed
`XENO_PLUGIN_ROOT` from section 7 rather than give it a reader, because `rules_hash` and a
rendered artifact resolve from that tree. Nothing under the local data location is hashed, so no
verdict can be made to depend on it. A97 carries both halves.

What this intent got wrong is in P3's deviations and in P4's results. The design listed the files
holding the literal and not the places reading the constant, which was the only compile error;
and P3's own learning lost a word to the shell, which is stated rather than repaired because the
phase was judged and the alternative was rewriting a sealed artifact over a typo.

G-Test judges this intent. P4 declares `test-report/go-test` bound by hash and its sealed verdict
carries `G-Test: pass` — the gate implemented in #212 reading the evidence of the intent after it.

<!-- xeno:section:release-notes -->
## Release notes

`XENO_PLUGIN_DATA` is read. It was the last entry in section 7's normalised environment that the
entry point exported and nothing read.

```
XENO_PLUGIN_DATA      local data location, always .xeno/local/
```

Setting it now moves the run marker, `phase.env`, the cost ledger and the enforcement report. An
absolute value is used as it stands, which is what `xeno-env.sh` exports; a relative one is
relative to the repository root, as every other path the runner handles is; empty or whitespace
reads as unset.

**Leaving it unset changes nothing.** Every path resolves to exactly where it has always been,
which is asserted twice — over the four resolved strings, and over the disk after a real
`phase start`. Nothing under that directory is covered by `artifacts_hash`, so the trail is
untouched and `gate verify` is at exit 0 over 390 verdicts.

Six places wrote `.xeno/local/` as a literal and one of them wrote it twice: `cmd/xeno/main.go`
repeated the `phase.env` path instead of calling the runner's accessor, so two files already had
to agree by hand about a path neither defined. That is gone; `model.LocalDir` and
`model.LocalPath` resolve it, and `PhaseEnvFile` is defined once.

Two exported identifiers changed shape, and nothing outside this repository reads them:
`cost.Ledger` and `runner.ReportPath` were path constants and are now `cost.LedgerFile` and
`runner.ReportFile`, file names, with `cost.LedgerPath(root)` and `runner.ReportPath(root)`
beside them. `xeno enforcement check` now prints where the report actually went rather than a
constant.

**Why this variable and not `XENO_PLUGIN_ROOT`.** A89 removed that one from section 7 rather than
give it a reader, because `internal/gates` resolves rules and templates from the plugin tree, so
an override would make `rules_hash` and a rendered artifact depend on the environment — its
measurement is that a differing but valid rule tree leaves `gate verify` at exit 0 while G-Policy
silently stops judging. Nothing under the local data location is hashed, so that argument does
not reach it. A97 records both halves.

The one cost a person can pay: the value is read at each use, so changing it between
`phase start` and `phase finish` leaves the marker in one directory and the lookup in another,
and the second command reports the phase as not running. A97 says so.

<!-- xeno:section:residual-risk -->
## Residual risk

Everything rests on one reading of section 7, and the clause taken alone argues the other way.
"Local data location, always `.xeno/local/`" can be read as a prohibition: if the value is always
one thing, a reader can only ever produce that thing, and making it configurable contradicts the
word "always". What settles it is the section's heading and its sentence that the runner "only
ever sees `XENO_*`" — normalisation, so the runner need not know a harness's conventions. If that
reading is wrong then this intent has made a specified constant configurable, and the repair is a
specification commit removing the clause, which was the maintainer's other option. A97 carries
the reasoning so the next reader checks the inference rather than inheriting it.

A person can split their own state and nothing can tell them. Changing the variable between
`phase start` and `phase finish` leaves the marker in one directory and the lookup in another,
and the second command reports the phase as not running — indistinguishable from a phase that is
not. Detecting it would mean recording the location inside the phase, which puts an environment
value into the trail and is what A89 refused for the other variable. Low likelihood, confusing
when it happens, and written down rather than guarded.

The entry point is read and not exercised. `xeno-env.sh` exports the value this intent consumes,
and nothing here runs the script, so the absolute case is tested against a value shaped like what
it produces rather than against it. That is #201's gap one size smaller, and it is the only link
in the chain from harness to resolved path that no test touches.

A redirected directory inside the tree is untracked and unignored. The `.gitignore` entry stays at
the default by criterion 5, which is right for the two obvious cases and wrong for the third: a
path inside the repository but outside `.xeno/local/` produces untracked files and no warning. The
runner cannot know which case a person is in.

This intent's own record carries a dropped word. P3's learning lost the name of a tool to a shell
command substitution and was noticed after the phase was judged, so repairing it meant rewriting
a sealed artifact or re-judging a verdict over a typo. It was left, stated in P4's results, and
its cause recorded as P4's learning — which is the right trade and also means the trail now holds
a learning about losing words in learnings.

What is not a risk: the 387 pre-existing verdicts, confirmed at exit 0; every existing repository,
whose paths are asserted unchanged twice over; and the gate path, which this change does not
touch.
