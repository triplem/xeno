---
intent: github.com/triplem/xeno#210
phase: 02-design
created: "2026-10-03T19:40:35Z"
schema_version: "1.0"
runner_version: dev+d19a1ca.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 76ba9b4e2d9bbf3d53dcc085a77cbb8c7fc9db358f6a85673b1d032523707c0f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**Three rules, stated separately.** The sentence is split because the three cases have
three answers, and one sentence covering all of them had to be wrong about two.

**Markdown prose at 88, tables and code blocks exempt.** The exemption is written down
rather than left to judgement, because it was already being taken silently: the clause
table in `CLAUSE-READERS.md` is wider than 88 and wrapping a row would break it.

**Go: nothing beyond `gofmt`.** Not 88, not 100, not 120. A width rule for Go would need a
tool to be real, `gofmt` has none, and a linter that reflows Go is a dependency and a
decision. The honest statement is that the formatter is the rule.

**Commit messages and descriptions at 72, with the reason left in `CONTRIBUTING.md`.** The
number is in `CLAUDE.md` because that is what is read before writing one; the reason is
one sentence away in the file that explains the host settings it depends on. Restating it
would be the second copy that goes stale.

**No checker.** The decision is deferred and A91 records it as deferred rather than
rejected, which is the distinction `CLAUSE-READERS.md` was written to keep.

<!-- xeno:section:alternatives -->
## Alternatives

**Keep one width and apply it everywhere.** Rejected: it is the current state. 88 for a
commit message produces something GitHub reflows into lines nobody wrote, and 88 for Go
needs a tool that does not exist in this project.

**Keep 88 for Go and write a width check for it.** Rejected for now, not forever. It is
the only one of the three that could be checked cheaply — a line-length pass over `.go`
files is ten lines and has no ambiguity. What it would need is a decision about which
existing lines get rewrapped and whether a long string literal or a struct tag is a
violation, and that is a conversation rather than a step. #210 says so and this intent
does not pre-empt it.

**Say nothing about width at all.** Rejected. The Markdown rule is real, it is followed
throughout this repository, and it is why the prose here reads as prose. Dropping it
because two thirds of the sentence was wrong would throw away the part that works.

**Put all three in `CONTRIBUTING.md` and drop them from `CLAUDE.md`.** Rejected. The
widths are what an agent needs before writing a line, and `CLAUDE.md` is the file that
arrives with every request. The reason belongs in the longer document; the number belongs
in the short one.

<!-- xeno:section:impact -->
## Impact

One paragraph in `CLAUDE.md`, one row in `ASSUMPTIONS.md`. No code, no field, no gate, no
rule, no artifact shape.

Nothing in the tree becomes wrong. Every Markdown file here is already at 88 with tables
exempt, and no Go file is currently over 88 — the one that was, `internal/plugin/plugin.go`,
was rewrapped in the intent that found it. So the change is to the statement, not to the
tree, which is the opposite of the usual direction and worth saying because it means
nothing has to be reflowed.

What changes in practice is what a future session believes. Go lines will drift past 88
where an identifier makes them, and that is now the stated rule rather than an unnoticed
violation of one. A commit message written at 72 will survive the squash intact.

The deferred checker is a real cost: three widths are stated and none is enforced, so the
Markdown rule remains a convention held by whoever is reading. `CLAUSE-READERS.md` already
lists that class of rule, and A91 adds this one to it knowingly.
