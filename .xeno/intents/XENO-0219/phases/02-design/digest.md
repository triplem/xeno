---
intent: github.com/triplem/xeno#162
phase: 02-design
created: "2026-10-01T15:52:45Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e1a37e0184327a264b7a6c84f47b57e3bc4673d9df7c8c7b336c04bd5f807f74
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`internal/git` is the only place a subprocess starts: one function returning hash, signature status,
parent count, author, subject and trailers per commit, with fixed arguments and the two refs the run
supplies. `internal/gates`'s package comment is corrected rather than bent — no build, no test suite, no
scanner, no model, and one `git log` over the given range. One subprocess per gate run, loaded lazily
and kept with its error, which is the shape #160's gap asked for and did not get for the rule tree. A
missing range is a finding before anything is read, shared by all four commit types so they cannot
disagree. A valid signature is `G` or `U`: trust lives in the verifier's keyring, so requiring it would
make the verdict depend on whose machine ran the gate, which is the same argument as refusing to infer
a range. An empty antecedent makes `section-implies-section` green, because the rule says a change to an
interface needs a note and not that every phase changes an interface; a section the template lacks is a
configuration error instead. A trailer is matched on its key and its value is not examined, because
checking the value needs an expression. `approver-not-author` compares strings, email where there is
one, and no decisions is green. The second pattern is `conventional-commits-with-issue`, accepting
`(#123)` and `(!123)` because GitHub and GitLab build squashed subjects differently.
