---
intent: github.com/triplem/xeno#183
phase: 03-implementation
created: "2026-10-03T18:39:48Z"
schema_version: "1.0"
runner_version: dev+0462eb7.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 33771f44c24e934b919254b45a12555ceeb5c9c098e49b5e01230669c3d5407c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The specification is committed alone and first: section 7's list drops `XENO_PLUGIN_ROOT` and keeps
three, and the resolution order paragraph becomes four — the plugin is the vendored one and nothing
else; the order is removed rather than built, with the gate path's dependency and the
silent-unjudging measurement; the client fallback is unreachable, with the no-template measurement;
and the condition an override would have to meet. The plan's WP7 sentence names three variables and
says what was on it. Then the comments: `internal/plugin`'s package comment said the order was "not
implemented here on purpose" and "a decision for section 7", and now says it is gone rather than
unimplemented; the entry point's "deliberately not set here" becomes "there is none to set"; and the
test keeps its assertion with a comment saying why — the thing worth preventing is a root arriving
from the environment, which does not stop being worth preventing when the document stops naming a way
to do it. A84's clause is closed and A89 records the decision with both measurements and the three
shapes. No Go behaviour changes, which is why the first commit carries no code. Stragglers checked by
grep: five files mention the mechanism, every one explaining its removal, and the plan none. Two
deviations: the rewrap to 88 touched more of `plugin.go` than the design anticipated, because the
file was written at 95 throughout in an earlier intent of mine and nothing checks the width; and the
shape of this intent is unusual — the specification change is the work and the code change is
comments, which is what the first standing rule produces when the thing corrected is a document.
