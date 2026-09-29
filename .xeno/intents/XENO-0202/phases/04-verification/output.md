---
intent: github.com/triplem/xeno#120
phase: 04-verification
created: "2026-09-29T06:04:24Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ad0a764.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: ebf8354b5fd84dd1b92b18350f4d0e41ca4ab06eb64a67df8813b96955eb281c
context_hash: da6d05948a550e6974a85b317d0888121c9487e82b734d31dae8eecf4ee20bf0
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: by-hand
evidence:
  - kind: test-report
    job: go-test
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | What proves it |
|---|---|
| AC1 the shipped filter exists in section 4's shape | `.xeno/plugin/secrets.yaml`, five patterns and five paths, and `TestTheShippedFilterCatchesWhatItNames` compiles and fires every one |
| AC2 a project adds and removes nothing | `TestAProjectAddsAndRemovesNothing`: a project reusing a shipped id leaves it in force, and its own pattern fires too |
| AC3 the hash covers the set, not the bytes | `TestACommentDoesNotChangeTheHash`, `TestAddingAPatternChangesTheHash`, `TestAddingAPathChangesTheHash` |
| AC4 order and file do not matter | `TestTheOrderAndTheFileDoNotChangeTheHash`, over the same set split across both files in the other order |
| AC5 the summary is redacted and the digest carries the hash | `TestTheDigestIsFilteredAndSaysWhichFilter`, and the end to end probe in the transcript: two fake secrets in, two named redactions out, the sentence intact |
| AC6 `section set` writes the field | `TestSectionSetWritesTheHashAndDoesNotRedactTheProse`, which also asserts the prose survives |
| AC7 no filter means no hash and no filtering, and no refusal | `TestNoFilterIsEmptyRatherThanAnError` and `TestWithoutAFilterNothingIsRedactedAndNoHashIsWritten` |
| AC8 A35 names two fields | read: the row lists `tool_version` and `rules_hash` and points at A62 |
| AC9 the honesty rule unchanged, with the number | measured in the transcript: removing `secrets_hash` from `WriterlessHash` diverges 88 of 90 verdicts. A62 carries it |
| AC10 everything green stays green | `go test ./...`, `gate verify` over 91 verdicts, and sixty intents' `by-hand` untouched |

Twelve tests, nine on the package and three through the runner. The before and after could be run this
time, since no signature changed, and the transcript carries it.

<!-- xeno:section:results -->
## Results

`go test ./...` passes on every package, including the new one. `gofmt -l .` outside
`vendor/` prints nothing, `go vet ./...` is silent, and `./xeno gate verify` recomputes
91 verdicts and matches.

The result to read is the probe. A summary carrying `AKIAIOSFODNN7EXAMPLE` and a `ghp_`
token went through `phase finish` and the digest on disk carries `[redacted:
aws-access-key]` and `[redacted: github-token]`, with "The run used" and "against the
bucket" intact around them, and a second line of ordinary prose about `secrets_hash`
untouched. The frontmatter carries the hash of the effective filter.

That is section 16's sentence working: the agent supplied the text, the runner held it,
and what reached the file is not what was handed in. Nothing the agent could have done
would have changed it, which is the one place in #120 where prevention is real rather
than refused.

The measurement that settled the honesty rule is in the transcript too: removing
`secrets_hash` from `WriterlessHash` turns 88 of the 90 verdicts divergent, green
recomputing to red. The two that survive are this intent's own, which carry a hash the
runner wrote.

The shipped filter catches each shape it names and leaves prose alone, which is a test
rather than a claim, and the prose it is tested against is a sentence from this intent's
own record.

<!-- xeno:section:gaps -->
## Gaps

**The pattern set is the security surface, and it is five patterns.** A secret whose
shape nobody anticipated reaches a digest exactly as it did before. `secrets_hash` is
what makes that discoverable afterwards: it says which filter a digest passed, so
somebody can tell which digests were written before a pattern existed. That is the
honest claim, and it is smaller than "digests are filtered".

**Nothing checks artifacts for secrets.** G-Secret stays `not-implemented`. The digest
is filtered; `output.md`, `learning.yaml` and every other file a phase carries are not,
and a secret pasted into a requirements section is in the repository with nothing to
notice. Section 4 says the gate reads the same file this writes against, which is now
possible and not done.

**`by-hand` is still accepted where a writer exists.** Appendix B says that is wrong
once a writer exists, and tightening it diverges 88 of 90 verdicts, so the loophole is
open by decision. A new artifact carries a real hash because the writer writes one, so
it matters only to somebody writing `by-hand` deliberately — which is to say, to the
case #120 is about.

**`paths_never_digested` governs nothing.** The globs are shipped, enter the effective
set and enter the hash, and no writer reads a file into a digest, so nothing consults
them. The field is honest and inert.

**`tool_version` still has no source**, so an agent still cannot produce a complete
artifact with the tool alone, and #120's enforcement question stays where its scoping
left it.

**The evidence is a local run.** Seventh intent in a row, same reason.
