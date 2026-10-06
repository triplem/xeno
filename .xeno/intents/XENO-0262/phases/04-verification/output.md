---
intent: github.com/triplem/xeno#258
phase: 04-verification
created: "2026-10-06T15:15:32Z"
schema_version: "1.0"
runner_version: dev+90593d0
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8706f5ccaab1377000cf962753623b88784cd5f0aac74348871924e7f4902c71
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - format: other
      job: go-test
      kind: test-report
      path: evidence/go-test.txt
      produced_by: go test ./...
      result: pass
      sha256: 433d736d3d2f44e6399dc05137840bdac8361d4e861f28c12b37f33b719b69b3
    - format: other
      job: checks
      kind: other
      path: evidence/checks.txt
      produced_by: the checks over the two amendments, run by hand
      sha256: 638f30665c94c3d5a959ef07e2d1e9706c258f246f77e536c99ce8ca35224d15
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Eleven criteria, by number. Ten pass and one does not, which is criterion 5.

| # | what it asserts | how it is checked | result |
|---|---|---|---|
| 1 | section 8 says where the reason lives | a person reads the paragraph | pass |
| 2 | section 8 says why, and names the cost | a person reads it | pass |
| 3 | section 5 gains the subsection, after Rendering | a person reads it; `git diff` places it before `## 6.` | pass |
| 4 | the subsection says the anchor is the template version | a person reads the second paragraph | pass |
| 5 | **two commits, neither carrying code** | `git log` over the branch | **fail — one commit** |
| 6 | the wording is the drafted wording | `diff` against the comment on #258 | **pass with two changes**, both in P3's deviations |
| 7 | nothing already covering this is duplicated | `git diff` shows 20 insertions and 0 deletions, none in section 7 | pass |
| 8 | `docs/clause-readers.md` is not changed | `git diff --stat`, which names one file | pass |
| 9 | `gate verify`, `go test`, `go vet`, `gofmt` | run; `evidence/go-test.txt` | pass |
| 10 | no new line over 88 columns | the measure before and after, as multisets | pass |
| 11 | the agent's edit is recorded where a reader meets it | P0's problem, D-1, the commit message | pass |

Criterion 5 failing is the honest result and not a technicality. P4's job is to run the measurement
again rather than to agree with the phase before it, and the phase before it wrote a criterion that
could not be met in this repository.

<!-- xeno:section:results -->
## Results

## The diff is twenty lines added and nothing removed

    docs/process-definition.md | 20 ++++++++++++++++++++
    1 file changed, 20 insertions(+)

Zero deletions is criterion 7 and it is also the check on the first departure from the draft: the
draft's section 8 heading said "replacing the third sentence", and a replacement would have shown
deletions. `grep -c "An open question without options moves the whole of the"` returns 1, so the
sentence that would have gone is where it was.

## The specification does not cite this repository, and now still does not

    grep -cE '#[0-9]{2,3}\b|XENO-|this trail' docs/process-definition.md   → 0
    grep -cE '#[0-9]{2,3}\b' CLAUDE.md                                      → 4

The first is the check and the second is what makes it evidence rather than a pattern that matches
nothing — the same search finds four references one file away. That is why the draft's "57 of the 64
in this trail" came out: it would have been the only sentence in the document telling an adopter a
count of somebody else's work.

## Criterion 5 fails, measured

`git log` over this branch shows one commit touching `docs/process-definition.md`, not two. Every
merge on `main` is a squash — the five before this intent are squashes of pull requests #262 through
#266 — so two commits on a branch arrive as one, and two on `main` would have meant two pull
requests and two intents. The first standing rule's substance, that a specification change does not
arrive mixed with its code, holds: the commit carries no Go source, no template and no rule.

The criterion is reported failed rather than reinterpreted. It asked for two commits and there is
one.

## The clause binds nothing yet, which is the anchor working

    65 × template: requirements@1.0.0

Every P1 artifact in the trail declares 1.0.0 and the new clause applies from 1.1.0, so no sealed
artifact comes into scope and no verdict moves. That is the second paragraph of the subsection
asserting itself rather than being believed.

## The suite

`go test ./...` exits 0 across every package, in `evidence/go-test.txt`. `go vet` is clean, `gofmt
-l` outside `vendor/` prints nothing, and `xeno gate verify` exits 0. No Go source is touched and
the specification is in no hash, so these say nothing was touched by accident and nothing more.

Over-88 lines outside table rows: 27 before the change and 27 after, as the same multiset of
lengths. None of the twenty new lines is one of them.

<!-- xeno:section:gaps -->
## Gaps

**Criterion 5 is unmet and is not being met later.** One commit, and the shape of the history is
what it is once this merges. The alternative remains available at a price — two pull requests and
two intents for two paragraphs — and is not being taken. A reader bisecting for the day the reason
moved onto the option will land on a commit that also moved the criteria clause.

**Nothing reads either new clause, and nothing will until the code exists.** A specification sentence
has no reader by construction, which is the ordinary state of this document and is exactly what
`docs/clause-readers.md` is a catalogue of. Two clauses were unsayable this morning and are now
unread, which is a step and not an arrival.

**The specification names two template versions that do not exist.** `requirements@1.1.0` and
`verification@1.1.0` are in no `template.yaml`. Deliberate, and it means a reader of the document
today cannot find the thing it names. The window closes when the implementing intent creates them,
and nothing bounds how long it stays open.

**Two sentences of the committed text differ from what was approved.** Named in P3's deviations and
in the report to the person, because a specification committed by an agent on a person's instruction
is only as good as the person's ability to see what actually landed. The diff is twenty lines and
they are being told which two sentences moved and why, rather than being left to read it.

**`docs/clause-readers.md` is now stale by two sentences.** Its decision paragraphs say each clause
"waits on the section 8 commit" and "waits on the section 5 commit". They now wait on the code. The
third replacement of that passage, which XENO-0260 predicted and is why it was written whole; still
a thing somebody has to remember, with nothing to remind them.

**Nothing checks that the specification and the plugin agree about template versions.** The document
now says what `requirements@1.1.0` requires and the plugin ships 1.0.0. When 1.1.0 is created,
nothing will compare its `template.yaml` against the clause describing it, and G-Schema checks that
an artifact's anchors match its declared template version rather than that the template matches the
specification. A new instance of the shape this project keeps finding.
