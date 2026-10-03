---
intent: github.com/triplem/xeno#183
phase: 01-requirements
created: "2026-10-03T18:38:05Z"
schema_version: "1.0"
runner_version: dev+0462eb7.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1678610fe109d5cfa4c7f576262ef253ed19a644e259f21b514d2bc07cdca4d8
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Section 7 says the plugin is the vendored one and describes no resolution order; the paragraph is
replaced rather than annotated and carries why each position was unreachable or unsafe, with both
measurements, because a removal is a heavier claim than an addition and somebody reinstating it
should have to argue with evidence rather than with an absence. `XENO_PLUGIN_ROOT` leaves the list,
three variables remain, and the plan's WP7 sentence says the root and the order were on it and are
not. One sentence states the condition an override would have to meet, so a later intent starts from
it: a runner that cannot verify a plugin does not accept one from outside the repository. The
specification change is its own commit with no code, which the first standing rule requires, and
nothing follows but comments and rows — because the order was never implemented, which is what makes
the removal safe and means there is no behaviour to migrate and no artifact that recorded a root. The
comments that called it a decision pending now call it removed, A84's open clause is closed, and the
test asserting the entry point exports no plugin root keeps its assertion with its reason changed:
the thing worth preventing is a root arriving from the environment, and that does not stop being
worth preventing when the document stops naming a way to do it. Out of scope: an override under a
condition, which was declined; `XENO_PLUGIN_DATA`, still listed and read by nothing; G-Supply, which
compares the only tree there is; and amending A42, which the order was the one thing that would have
required — the argument for removing it rather than for weakening A42.
