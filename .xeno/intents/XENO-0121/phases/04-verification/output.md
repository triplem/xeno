---
intent: github.com/triplem/xeno#121
phase: 04-verification
created: "2026-09-28T19:44:53Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c2ba6b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 91fb384b30efa143f6b4b66bd2ff645512239cf3c0fbbef49c40139865d10ca6
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
| AC1 neither plugin in `.releaserc.json` | the transcript prints the parsed plugin list: commit analyzer, release notes generator, github |
| AC2 no `extra_plugins`, comments describe the present | `grep -c extra_plugins` is 0, and the header paragraphs are read |
| AC3 `audit.yml` and `SUPPLY-CHAIN.md` name neither | the sweep across every `.md`, `.json`, `.yml` and `.sh` outside `vendor/` returns one line, A23, which records what it superseded |
| AC4 `CHANGELOG.md` gone, nothing references it | the same sweep, and `ls` reports the file absent |
| AC5 A23 superseded, the row still readable | the sweep prints the whole row: the original claim, then what replaced it |
| AC6 the release job's own gates untouched | the step names in order: format, vet, test, verify, release, build, module set, bill of materials, checksums |
| AC7 protection unchanged, no secret added | `scripts/github-settings.sh` differs from `main` by zero lines, and nothing in the diff adds a secret reference |
| AC8 everything green stays green | `go test ./...`, `gofmt`, `go vet`, and `gate verify` over 73 verdicts |
| AC9 the next merge publishes a release | **not proved here.** Nothing in a branch can show what a push to `main` does |

Eight rows are evidence and the ninth is a prediction. Nothing in this change is Go, so the
tests and the gates prove that the artifacts of the intent are well formed rather than that the
release works.

<!-- xeno:section:results -->
## Results

`.releaserc.json` parses and carries three plugins. All three workflows parse.
`extra_plugins` occurs nowhere in `release.yml`. `CHANGELOG.md` is absent.
`scripts/github-settings.sh` is identical to `main`.

The sweep for `CHANGELOG`, `@semantic-release/git` and `@semantic-release/changelog`
across every Markdown, JSON, YAML and shell file outside `vendor/` returns exactly one
line: A23, which names what it superseded. That is the result worth having, because the
failure mode of this change is a pin or a sentence left behind somewhere nobody looked.

`go test ./...` passes, `gofmt -l .` outside `vendor/` prints nothing, `go vet ./...` is
silent, and `./xeno gate verify` recomputes 73 verdicts and matches.

What none of it shows is that a release will be cut. That is AC9, and it is a property
of the next push to `main` rather than of this branch. The evidence for it will be a run
of the release workflow that reaches `publish` and a v0.19.0 that exists.

Read rather than executed: the corrected comments in `release.yml` and `audit.yml`. A
comment is proved by reading, as #111 established, and this change turns one true
sentence from a reassurance into the explanation of a failure.

<!-- xeno:section:gaps -->
## Gaps

**The change is unproved where it matters.** AC9 needs a push to `main`, so the first
evidence that the pipeline works is the merge of this pull request. If it still fails,
it fails on something this branch cannot see, and the diagnosis in #121 was wrong about
the cause rather than incomplete.

**No test covers the release configuration.** Nothing in the repository reads
`.releaserc.json` except semantic-release in CI, so a plugin list that is valid JSON and
wrong is green here. The same holds for the workflow files beyond their being parseable
YAML. This intent adds no check for that and the gap is the same one every CI
configuration in every repository has.

**`SUPPLY-CHAIN.md` is maintained by hand against three other files.** The requirements
phase recorded it and the implementation confirmed it: four places named the plugin set,
and a grep found a fifth mention the design had not listed. Nothing makes them agree.
That is a finding about the document rather than about this change.

**The release history is now one command away rather than one file.** `git log --
CHANGELOG.md` renders what the tree used to carry, and finding it takes knowing the file
existed. The releases page carries the notes for anybody who does not.

**The evidence is a local run.** Fourth intent in a row, same reason: the pipeline
publishes no artifact, so the transcript is bound by hash through the stand-in directory
and `attached.yaml` cannot say what produced the bytes.
