---
intent: github.com/triplem/xeno#183
phase: 04-verification
created: "2026-10-03T19:08:44Z"
schema_version: "1.0"
runner_version: dev+0462eb7.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cf91720178eae3b115b4fddd1933d4db0a5d1c2cb88e3728d98b21a3c659874f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Each acceptance criterion of P1 against what holds it. Most are held by reading a file, which is
what a change to a document means, and two are held by a command.

**Section 7 says the plugin is the vendored one and describes no order** — read in
`docs/process-definition.md`. Four paragraphs where there was one, and the order appears only as
something removed.

**`XENO_PLUGIN_ROOT` is not in the list** — read: three entries remain.

**The plan's WP7 sentence names the same three** — read, and `grep -c 'XENO_PLUGIN_ROOT\|plugin-root'
docs/implementation-plan.md` is 0.

**One sentence states the condition an override would have to meet** — read, the last paragraph of
the four.

**The specification change is its own commit with no code** — `git show --stat` on that commit: two
files, both under `docs/`.

**No code changes how the plugin is found** — `git diff` over `internal/` in the second commit
touches comments and one test comment, and the suite is unchanged. `internal/rules`,
`internal/template`, `internal/secrets` and `internal/plugin` resolve from the vendored directory
before and after, which is also what `gate verify` at 285 says.

**The comments call it removed rather than pending** — read in `internal/plugin/plugin.go` and
`.xeno/plugin/bin/xeno-env.sh`.

**A84's open clause is closed** — read in `ASSUMPTIONS.md`.

**The test keeps its assertion with a changed reason** —
`TestTheEntryPointStartsTheRunnerFromThePath`, which still fails if the entry point exports a plugin
root. Held by the test rather than by reading, and it is the only part of this change with a
mechanism behind it.

**Nothing still claims the mechanism exists** — the grep in the results below, five files, each
explaining the removal.

**The suite, `gofmt`, `go vet` and `gate verify` are unchanged** — which is the criterion a document
change with no mechanism behind it should be judged by, and the one that would have caught an
accidental behaviour change.

<!-- xeno:section:results -->
## Results

**`go test ./...`** — eighteen packages ok, no case added or changed except one comment.

**`gofmt -l .` outside `vendor/`** and **`go vet ./...`** — nothing.

**`./xeno gate verify`** — `verified 285 verdicts`, exit 0, exit code captured.

**The specification commit, alone** — two files under `docs/`, 32 insertions, 11 deletions, no `.go`
file and no `.yml`.

**The straggler grep**, over `docs/`, `ASSUMPTIONS.md`, `README.md`, `internal/` and the entry
point:

| file | what it says |
|---|---|
| `docs/process-definition.md` | the removal and why |
| `ASSUMPTIONS.md` | A84's closure and A89 |
| `internal/plugin/plugin.go` | the order is gone rather than unimplemented |
| `.xeno/plugin/bin/xeno-env.sh` | there is none to set |
| `internal/plugin/plugin_test.go` | why the assertion stays |

`docs/implementation-plan.md`: none. No file claims the mechanism exists.

**The two measurements this intent rests on were taken before it**, in the exchange that produced
the decision, and are recorded in the document rather than only here: a valid rule tree that differs
leaves `gate verify` at exit 0 over 273 verdicts with G-Policy silent about 75 phases; and a project
with no vendored plugin answers `no template "intake" under .xeno/config/templates or
.xeno/plugin/templates`.

**What this intent did not need to measure.** There is no behaviour to exercise. That is the result:
a document that described a mechanism now describes its absence, and the code that always read the
vendored directory still does.

<!-- xeno:section:gaps -->
## Gaps

**P3's justification for the rewrap is about to stop being true.** That phase recorded the comment
rewrap as done because the file "breaks a convention this project states in `CLAUDE.md`", and the
maintainer has since decided that Go source has no width rule beyond `gofmt`. The phase is sealed and
keeps the reason it was written with, which is correct — a phase is judged against what was in force
— and a reader of that deviation will find a convention the repository no longer has. The rewrap
itself is harmless and stays.

**Nothing prevents the order being reinstated by somebody who does not read the paragraph.** The
section argues for the removal and the code has no hook for a root, so reinstating it means writing
the resolution as well as the clause. That is the protection: the clause and the mechanism would have
to arrive together, which is the state everything else in section 7 is already in.

**`XENO_PLUGIN_DATA` is still listed and read by nothing.** The entry point exports it and five
packages use the constant. Section 7 pins it to one value, so a reader cannot tell whether that is a
variable with no reader or a variable whose reader is a constant. A84 records it and this intent
deliberately did not touch it, which means the list still contains one entry in the state the
removed one was in.

**G-Complete still runs only at P5.** Not this intent's subject and still nobody's issue, which is
the entry #202 names as known and unfiled.

**The condition kept in section 7 has no test and cannot have one.** It describes what a future
override would have to do. It is a sentence for a reader, which is the category #202 warns against
counting as a gap — and it is worth saying that this intent added one of those deliberately.

**The removal was argued from two measurements taken on this repository.** A project with a different
rule set, or one where the gate path did not read the plugin, would have different numbers. The
argument is sound here and is not a general claim about the design, which is the kind of thing a
later reader could over-read from a document that states it without the measurement beside it — which
is why both measurements are in the document.
