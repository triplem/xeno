---
intent: github.com/triplem/xeno#267
phase: 00-intake
created: "2026-10-06T19:56:29Z"
schema_version: "1.0"
runner_version: dev+1d61fa2
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d9dab5eb586badbf5d808ef5b68c697c1d992fd2902db37e9482ec069bc306f0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`phase start` writes `context.lock.yaml` from the context scope, and at P0 the scope does not
exist yet, because `scope set` refuses to run before the phase has started. So no P0 lock in this
trail records a `files` list, 0 of 117, and `staleReads`, `budget` and `ChangedSince` each read
nothing at the intake. #267 put three candidates and the maintainer chose the third: section 5 and
section 7 say why the intake stands outside the two checks, and two code comments that explain the
empty list as a phase older than #217's rule are corrected, because that is true of P1 to P5 and
false of every P0.
