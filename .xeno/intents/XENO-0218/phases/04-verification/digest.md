---
intent: github.com/triplem/xeno#160
phase: 04-verification
created: "2026-10-01T15:42:46Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e9fb6d2693ca8c24c8ad60edc56f51de5fa6f5164f40ed713f333e0b0b9f21ed
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
94 cases pass in `internal/gates`, 17 packages `ok`, `gofmt` and `go vet` clean, `gate verify` at exit 0
over 187 verdicts. The gate was also run against the real binary on a copy, using XENO-0217's sealed
phases: a review rule with no answer turned P5 red naming the rule; an entry with `result: deviation`
and a note turned it green, which is section 9's mechanism working end to end; removing the note turned
it red again; and the section 9 example checked rule turned 02-design red naming the rule and the
unimplemented type while G-Rules stayed green on the same tree, which is two statements coming out as
two gates. The copy also showed what the fixtures cannot: that phase carries `rules_hash: by-hand`,
because it predates #158's writer, so for the whole existing trail the gates that read rules can
disagree with the field meant to record which rules applied. Gaps: entries before P5 are ignored
silently; nothing checks that an answer is true and section 16 does not say so yet; a forged rule id is
indistinguishable from an answer; no checklist has been written by anything but a test; the registry's
evaluate branch is written and unexercised; and both gates resolve the tree separately, which honours
section 7's "the same way" and not its "once".
