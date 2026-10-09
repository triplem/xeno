---
intent: github.com/triplem/xeno#95
phase: 04-verification
created: "2026-10-09T16:55:14Z"
schema_version: "1.0"
runner_version: dev+d2bc411.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0396db2a4b80d486cf5aea37057af4f55e49d852f0ecd97528d678b4e3c57709
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
Seven criteria, all met, and the figures taken from the tree with the change in it.

The result that matters is nine managers resolving to nine values the tree carries, with the six
that already existed run beside the three that are new. Before the change the same script
printed `NO MATCH` for those three while printing these six. That ordering is the whole of why
the finding is a finding: six working entries are what distinguish a regex that matches nothing
from a script that never ran, and this project has a rule about that because #201 recorded an
absence as evidence when the thing checked had just been deleted.

Criterion 6 was checked by comparing every key against `HEAD`'s copy rather than by reading the
diff. Thirteen keys identical, the six existing managers byte-identical as a list, one key
moved, and it is `commitBody`. A diff would have shown the same thing and would not have proved
the absence of anything.

Criterion 4's answer was known before the edit, from the pattern rather than from the run: the
digest extractor's first group excludes `:`, so a tag inserted before `@sha256:` moves what
that group captures and leaves the stored digest alone. Running it confirmed a reading instead
of discovering an outcome.

Two gaps are recorded rather than closed. Nothing re-runs the manager check, so the rot this
intent fixed by hand can recur the same way; a committed script beside `scripts/plugin-digest.go`
is the likely answer and was not taken, because it is a tool the design did not propose and
adding one unannounced is how a scope stops meaning anything. And nothing has run renovate: the
datasources are unexercised, so `pypi` answering for `zensical` and `extractVersionTemplate`
yielding `8.30.1` from `v8.30.1` are read off documentation, not observed. #350 carries that
half.

Three sections of the five, no question, no decision.
