---
intent: github.com/triplem/xeno#151
phase: 05-review
created: "2026-10-01T14:02:25Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c54db93.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0eadf91e941d548586f4facf27d6cd0f886d1256eadc39a966bfe89d928d5f22
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

| item | state |
|---|---|
| The three standing rules | held. No document edited; this is the first code written against the specification `c54db93` changed, and it adds nothing to it. No invented field: five symbol fields and three provenance fields, all from section 5 and WP15. The change belongs to #151, labelled `wp15`, on `151-the-index-format` |
| #149 is implemented and not reversed | Xeno ships no indexer. The producer is `scripts/go-symbols.go` behind `//go:build ignore`, beside two shell scripts, and the module never builds it. Criterion 6 and `deviations` both say so, because a Go indexer in the tree is the thing a reader would misread |
| The degradation property | five causes, one outcome, no error return. The signature says absence is ordinary, which P1 recorded as a constraint on the type rather than the behaviour |
| A42's property, extended | `internal/gates` reaches no `internal/index`, checked in CI beside the network check rather than by a test, because the property is about the import graph |
| One dependency | `go.mod` and `go.sum` unchanged. The producer is standard library |
| Nothing enters the trail | the index is under `.xeno/local/`, gitignored; `gate verify` recomputes the same verdicts |
| Eighty-eight columns, SPDX, no copyright line | held on all three new files |
| The check suite | build, vet, test, gofmt clean; `gate verify` 168 verdicts, exit 0 |
| Commit and reference convention | Conventional Commits, `Refs #151` on the commit and `Closes #151` on the one that finishes it |

**What a reviewer should push back on, hardest first.**

**`kind` and `container` are free strings.** This is the decision I would contest if I were reviewing. An
enumeration would let Xeno validate an index and give a project a list to write against. The argument
against is in P2 and at the type: once #149 put the choice of indexer on the project, Xeno cannot
enumerate what that tool thinks a symbol is, and nothing in Xeno compares either value. A reviewer who
thinks a format with unconstrained fields is not a format is making a fair point, and the counter is that
the constraint would be Xeno's taxonomy imposed on four languages it does not parse.

**Criterion 2 is half met and I have claimed it as half.** The format is documented in a `testdata` file.
A reviewer may reasonably say that is not documentation and the criterion is unmet rather than partly
met. I would not argue: the published half is WP16's and there is nowhere else it could have gone that
standing rule 1 permits.

**A package with no caller.** `internal/index` and `runner.symbolIndex` are read by nothing. P2 predicted
it and the reason holds — a format has to exist before anything reads one — but a reviewer is entitled to
ask whether the reader should have waited for the record and arrived with it.

<!-- xeno:section:release-notes -->
## Release notes

**Xeno can read a symbol index a project produced.** Configure it in `.xeno/config/project.yaml`:

```yaml
index:
  path: .xeno/local/index/symbols.yaml
  max_age_hours: 24
```

Both fields are optional. An absent block, an absent file, an unreadable one, one without provenance, or
one older than `max_age_hours` all mean no index, and a phase runs without one. Nothing fails.

**Xeno ships no indexer and names none.** Produce the index with tree-sitter, ctags, your build system or
your language server. The format is eight fields and is shown, annotated, in
`internal/index/testdata/example-symbols.yaml`: three of provenance — the tool, its version, when it was
produced — and five per symbol — name, kind, file, line, enclosing container. `kind` is whatever your
tool calls things; Xeno compares it against nothing.

An index that cannot say when it was produced is treated as absent, because a stale index is worse than
none.

**Nothing queries it yet.** The record in `context.lock.yaml` and the MCP operation are the next pieces,
so this release reads an index and does nothing with it. A project can produce one now and have the
format checked against its tool before anything depends on it.

**For this repository:** `go run scripts/go-symbols.go > .xeno/local/index/symbols.yaml` writes one for
its own Go source. That script is this project's own tooling, not part of Xeno.

<!-- xeno:section:residual-risk -->
## Residual risk

**The format is unproven against a real question, which is the risk that outranks the others.** Nothing
consumes the index. Eight of nine tests read a fixture written in the same intent as the reader, so they
assert the reader accepts what its author intended. The first real consumer — a phase, an agent with a
question — is what judges whether five fields are the right five, and if they are not, the cost is a
format change after a project has written a producer against it. The mitigation is that the next piece is
the record and the one after is the query, so the distance to a real consumer is short.

**`produced_at` is trusted and unverifiable.** Staleness is the one judgement Xeno makes about the index,
and it is made on a number the producer wrote. A tool with a wrong clock, or one stamping the time it
started a long run, yields an index reported fresher than it is — the direction that matters, since the
stale case is the one the check exists for. A file mtime would be a different claim and forgeable too, and
the documents permit the index to be wrong, so nothing in Xeno can close this. What it means in practice:
a project whose index is silently stale gets confidently wrong locations, and the agent pays one wasted
read per wrong answer rather than producing a wrong verdict.

**`Lookup` is linear.** Fine at 997 symbols, wrong at a brownfield Java repository's hundreds of
thousands, on every query of every phase. A map by name at load time is the answer and is deliberately not
built, because nothing queries it and the shape of the caller is a guess until the MCP operation exists.
The risk is that the first real use is also the first performance surprise.

**An open `kind` means two projects' indexes are not comparable.** Nothing in Xeno compares them, so
nothing breaks, but an organisation reading two projects' `tools` records will find `func` in one and
`function` in the other. That is the price of #149's reasoning and it is paid knowingly.

**A producer this project maintains for itself will drift from the format.** `scripts/go-symbols.go` is
tested against the reader, so a format change breaks the test rather than the producer silently. That is
the only thing holding them together and it is enough while they are in one repository.

**Not a risk.** That the index is gitignored and not reproducible from the trail. Derived data does not
belong in the trail, the documents say so, and the `tools` entry is what makes a phase's answer
attributable without carrying the index itself.
