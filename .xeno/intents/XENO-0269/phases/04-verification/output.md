---
intent: github.com/triplem/xeno#207
phase: 04-verification
created: "2026-10-07T09:09:14Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 72a703cf70aade44db4a52257938af36f5074f7a810042446b34d4ab35329d97
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - kind: test-report
      result: pass
      produced_by: go test ./...
      sha256: bd831099f0331da8b01c5b137e911eb0dabf34036cb971b0bf960b04f72bb9c2
      path: evidence/go-test.txt
      job: test
    - kind: build-log
      result: pass
      produced_by: go build, gofmt -l ., go vet ./..., xeno gate verify, git diff against main
      sha256: 984c5428be2171d65185b870b700f5cd85b39662447ef3ce590ae425640f8af7
      path: evidence/checks.txt
      job: checks
    - kind: other
      result: pass
      produced_by: the three writers and G-Schema's involvement, read in the code, with a positive control
      sha256: 62470584c1f9756a9e39f75f392a6109182923bfa5da80839fb8d19aee175abc
      path: evidence/writers.txt
      job: writers
    - kind: other
      result: pass
      produced_by: the paragraph, the clause count and the register row, read in the files
      sha256: 10913719551ae2da46083acc29db8bb74b7d95328ed520b7293025a28424f0fd
      path: evidence/passages.txt
      job: passages
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Twelve criteria, by number, all passing. Eight are read by a person, because the deliverable is
prose. Four are checks.

| # | What it asserts | What proves it |
|---|---|---|
| 1 | section 12 says the triple is declared and not measured, after the provider-register paragraph | read: `evidence/passages.txt` prints both, in order |
| 2 | it says where each of the three comes from, as the code has it | read against `evidence/writers.txt`, which prints `ToolVersion`, `recordedToolVersion` and `agent()` |
| 3 | it says G-Schema does presence and shape and nothing more | `evidence/writers.txt`: `sessionFields` is the list, and the grep for the three fields in the gates returns that line and one other, which is criterion 3's one qualification and is recorded in the results |
| 4 | it cites section 7 rather than asserting an impossibility | read: the paragraph's fourth sentence |
| 5 | it says why a check against the second copy is refused | read: the fifth sentence, against #207's own last sentence |
| 6 | it names the gateway, what it is worth and why it is out of reach | read: the sixth and seventh sentences, against the plan's section 9 answer |
| 7 | it ends on how the register is read | read: the last sentence |
| 8 | nothing in it asks a tool for anything | read: no normative word in it addresses the runner, a gate or CI, which is what makes it an explanation; `evidence/passages.txt` carries the clause-readers paragraph that states it |
| 9 | the explanation count moves 129 to 130 with the reason | `evidence/passages.txt`: the table row and both paragraphs |
| 10 | the register carries one row | `evidence/passages.txt`: A100, five cells, state naming the decider and 2026-10-07 |
| 11 | nothing else changes and the suite passes | `evidence/checks.txt`: the diff against main is one document and no Go file; `evidence/go-test.txt`; `gofmt`, `go vet`, `xeno gate verify` over 481 verdicts |
| 12 | the specification commit is its own and first | `git log`: one commit carrying the paragraph, its message naming the exception and the correction |

<!-- xeno:section:results -->
## Results

**The three writers, read in the code rather than recalled.** `evidence/writers.txt` prints
them. `r.ToolVersion` is set from `XENO_HARNESS_VERSION` at construction and from
`--tool-version` over the top; `SectionSet` writes it into the frontmatter and
`recordedToolVersion` reads it back out of `output.md` for the digest, so the second write
copies the first rather than asking again. `agent()` reads `tool` from `project.yaml`'s
`agent.tool` and `model` from `agent.model.default`, and `XENO_HARNESS` replaces `tool` and
nothing else. The function's own comment says both are project level and "not among the fields
that come from the harness", which is what the paragraph now says and the approved draft did
not.

**One qualification to criterion 3, found while collecting this evidence.** The claim was that
G-Schema does presence and shape with the triple and nothing else. The grep says otherwise in
one place: `hashes` reads `raw["tool"] == "manual"` and uses it to decide that a `by-hand`
placeholder is honest in every hash field. That is not corroboration — nothing checks the value
— but it is a check relaxed on a declaration, which is a sharper instance of #207's concern
than the register is: an artifact that says `tool: manual` has every placeholder accepted, and
nothing can tell a manual artifact from one that claimed to be. The paragraph as written stays
true, because what it says is that nothing corroborates the three. The extra fact is outside
#207's "done when" and outside the approved wording, so it is raised rather than absorbed and
not written into the specification here.

**The grep for the three fields in the gates was run with a positive control.** The same
pattern over `internal/model` and `internal/runner` returns the writers, so an absence in the
gates is an absence and not a pattern looking in the wrong place. That is also how the
qualification above was found rather than missed.

**The paragraph and the one it follows, read in the file.** `evidence/passages.txt` prints both
so that a reviewer reads the pair in the order section 12 has them. Eight of the twelve criteria
are that reading, because a check that a paragraph exists says nothing about whether it says
the thing.

**The clause-readers change.** The explanation count reads 130. Two paragraphs sit beside
#267's three: the first says which clause moved the count and why it has no row, the second
says the clause is reader-shaped in one direction only and that the direction is already taken,
since building the refused check would leave the clause true and the gap invisible.

**The register row.** A100, five cells, state `approved (#207)` naming the maintainer and
2026-10-07, with the two options not chosen and the correction to the draft. A98 and A99 are on
the unmerged branches for #228 and #238; P3's deviations say why the number skipped to A100.

**The checks.** `go build`, `gofmt -l .` outside `vendor/` and `go vet ./...` clean.
`xeno gate verify` recomputes and matches 481 verdicts. `go test ./...` passes across 20
packages. The diff against `main` is one file, `docs/process-definition.md`, and no Go file —
which is the fastest way to see criterion 11 hold.

<!-- xeno:section:gaps -->
## Gaps

**The gap is stated and not closed, which is the whole of what this intent does.** A register
built on the triple still says what an agent declared. Nothing in the tree is more trustworthy
than it was yesterday, and a reader who wanted the fields corroborated gets a reason instead.
That was #207's own preferred branch and it is still the weaker of the two it offered.

**`tool: manual` relaxes a check, and this intent does not fix it.** `hashes` accepts the
`by-hand` placeholder in every hash field when the artifact says `tool: manual`. The value is
self-reported, so an artifact that claimed to be manual would have its placeholders accepted,
and nothing distinguishes the two cases. It is outside #207's "done when" and outside the
wording that was approved, so it is written here and raised rather than slipped into the
specification. Until somebody decides, the honest statement is that one gate acts on a
declaration and section 12's new paragraph does not mention it.

**Nothing checks that the paragraph stays true.** It is an explanation, counted and not
enumerated in `docs/clause-readers.md`, and the second paragraph there says why no reader is
possible: the thing that would falsify it is a gateway and a figure read against it, both
outside v1, and the thing somebody is likeliest to build — the comparison against
`project.yaml` — would leave the clause true while making the gap invisible. So the clause is
unguarded by design and the document says so.

**The gateway's record was read from the plan, not measured here.** The claim that it names the
deployment a request was routed to rather than what a provider attests comes from section 9's
answer, measured against LiteLLM 1.102.1 for #56. Nothing in this intent re-measured it, and
the paragraph inherits whatever that measurement is worth. A different gateway may record
something else.

**Eight of twelve criteria are a person reading prose.** `evidence/passages.txt` prints each
passage so that the reading is of what was written rather than of a claim about it. That is the
most this phase can offer for a prose deliverable and it is not a check.

**The register is discontinuous on this branch.** A97 is followed by A100, because A98 and A99
are on the unmerged branches for #228 and #238. If either is abandoned the gap stays.

**The approved wording was wrong once and nothing caught it but a reading.** The error survived
drafting and approval, and what found it was opening `agent()` while writing the intake. P0's
`learning.yaml` proposes the rule; nothing enforces it, and the next drafted sentence about what
the code does is as exposed as this one was.
