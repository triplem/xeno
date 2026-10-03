---
intent: github.com/triplem/xeno#172
phase: 04-verification
created: "2026-10-03T10:39:22Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8d484db13807a7920c9c83f16c11d5a057989aaa9e36f72a11c76e282569ccc0
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

| Criterion | Test |
|---|---|
| A missing link's document is a G-Schema finding naming component, path, profile and a next step | `TestALinkNamingADocumentThatIsNotThereIsAFinding` |
| A link whose document exists is no finding | `TestALinkWhoseDocumentExistsIsNoFinding` |
| A profile with no links, and a repository with no profile, produce nothing | the same test's second half |
| The finding comes through G-Schema at any phase | `TestTheLinkFindingComesFromGSchemaAtAnyPhase` |
| A link whose document exists still joins the base, last, with its hash | `TestADeclaredLinksDocumentIsInTheBase`, from #171, still passing |
| It does not block | the demonstration: the phase was red and the next phase was startable after a decision, as every finding is |
| The byte count is not touched | this intent's diff: `internal/gates/gates.go` and its test, nothing in the model or the runner |
| Nothing else moves | the suite, `go vet`, `./xeno gate verify` |

Nothing here needed a new fixture. #171's budget tests already wrote a profile, a lock and the files
the base names, so each of these is that fixture with one more field.

<!-- xeno:section:results -->
## Results

**The suite is green.** 431 cases pass, nothing fails, `gofmt -l` outside `vendor/` lists nothing,
`go vet ./...` is silent, `./xeno gate verify` is at exit 0 over 228 verdicts.

**The real mistake was made on a copy and reported.** A profile for this repository with one link,
whose document was typed `docs/process-defintion.md` — the transposition somebody actually makes —
produced:

    .xeno/intents/XENO-9998/phases/00-intake/context-profile.yaml:
      the link for internal/gates names "docs/process-defintion.md", which is not in the tree

That is the finding naming the file that is wrong, the component the link claims, and the path as
written, so a reader sees the typo rather than being told to go looking for it. Before this intent
the same profile produced nothing at all.

**The base is unchanged.** #171's test for a declared link's document joining the base, last, with
its hash still passes, which is the criterion this intent was most at risk of breaking: the check
and the resolution read the same field for different purposes.

**It does not block.** The demonstration's phase was red on the finding and decidable like any
other, which is the budget's precedent and what #171's criterion asked for.

**One thing the implementation phase predicted and the demonstration confirmed.** A link with a
document and no component reads badly: "the link for a link names …". It is in the gaps, it is not a
profile anybody writes on purpose, and the alternative was to skip such a link and leave its
document unchecked.

**What was not done is not done quietly.** The byte count is untouched: this intent's diff is one
file in `internal/gates` and its test. Section 5's `files` is a path and a hash, so recording a size
is a field the specification does not have.

<!-- xeno:section:gaps -->
## Gaps

**The byte count still moves after a phase is sealed**, which is the other half of what this intent
was asked to fix and which it cannot: `gate verify` can report a budget finding against a phase that
was inside its budget when it ran, because the size is measured from the tree and the lock carries
no sizes. The specification change it needs is one clause and it is named in the review.

**A link with no component produces an awkward sentence.** "The link for a link names …". Not a
profile anybody writes deliberately, and the better sentence needs either a second message or a
decision to skip such a link, which would leave its document unchecked.

**Nothing checks that a link's component means anything.** It is a prefix in the project's own
vocabulary and may not be a directory, so a link can name a component that exists nowhere and a
document that exists, and be reported by nothing. That is deliberate — the non-goals say why — and
it means half of a link is validated and half is not.

**A link's document is checked for existence and nothing else.** It is not read, not hashed by this
check, and nothing says whether it documents what the component contains. Section 5 forbids
inferring that mapping, so what a declared link proves is that somebody declared it.

**Still no profile in this repository**, so this check, like everything else in WP8, is exercised by
tests and one demonstration on a copy. The experiment it unblocks is next and it is the first thing
that will read these findings in anger.
