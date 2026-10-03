---
intent: github.com/triplem/xeno#195
phase: 04-verification
created: "2026-10-03T14:23:34Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+b54626e.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: ea5acde82962207262a2c3f87b322b0aa72a98003f070796a460fd128782e892
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

Each acceptance criterion of P1 against the test that holds it.

**The header is the runner's** — `TestTheLearningRecordCarriesTheRunnersHeader`, which checks the
returned record against `model.RunnerVersion`, reads the file back off disk for the header fields,
then finishes the phase and compares `learning.yaml`'s `runner_version` with the `gate.yaml` the same
binary wrote beside it. The agreement and not a literal, because a pinned string would pass on a
build that stamped neither file.

**The four keys are arguments** — the same test supplies all four and the surface test passes them
through the command line.

**`--no-finding` writes the empty record with the same header** —
`TestNoFindingIsTheEmptyRecord`, which checks `no_finding` is set, no entries exist, and the header
still carries the runner's version.

**Entries accumulate in order** — `TestLearningEntriesAccumulate`, three calls, then the count and
the first and last categories.

**Without `--phase` the record is the intent's own** — `TestWithoutAPhaseTheRecordIsTheIntentsOwn`,
which asserts the record carries no phase, exists at intent level, and was *not* also written into a
phase. The last of those is the assertion that would have caught a `learningPath` that defaulted.

**The contradictions are refused in both directions** —
`TestNoFindingAndAnObservationContradict`, which builds each order on its own fixture.

**Every argument refusal, with no file left behind** —
`TestALearningIsRefusedBeforeItReachesTheFile`, seven cases: a category outside the four, each of
the three missing keys, whitespace as a target, nothing at all, and `--no-finding` beside an entry.
Each asserts the refusal's wording and that `learning.yaml` does not exist afterwards.

**G-Learning passes on what the command writes** — measured on the throwaway tree, and held by this
intent's own five finished phases, each of whose records was written by the command and each of whose
verdicts carries `G-Learning pass`.

**A hand-written record still works** — the whole existing suite, whose fixture writes none, and
`gate verify` over thirty intents' worth of hand-written records at exit 0.

**The staircase** — the surface test: 0 with the record named, and 1 with `refused:` on standard
error and nothing on standard output for a category the set does not have.

<!-- xeno:section:results -->
## Results

**`go test ./...`** — eighteen packages ok. Fourteen cases are new, counting the refusal table's
seven as the cases they are: thirteen in `internal/runner` and one in `cmd/xeno`.

**`gofmt -l .` outside `vendor/`** and **`go vet ./...`** — nothing.

**`./xeno gate verify`** — `verified 271 verdicts`, exit 0.

**On a throwaway tree, before the tests were written.** A record written by the command:

    runner_version: 0.1.0-dev+b54626e.dirty

against `0.1.0-dev` in every hand-written record in this repository. The intent level form writes to
`.xeno/intents/T-1/learning.yaml` with no `phase` field. All seven refusals fire with the wording the
tests now assert. A second call reports `records 2 learning(s), the last of category prompt`.
`phase finish` then reports `G-Learning pass`.

**This intent's own trail is the demonstration.** Every learning record of XENO-0232 was written by
`xeno learning record` and none by hand. Each of the five finished phases carries `G-Learning pass`,
and each record's `runner_version` is the one its `gate.yaml` carries — which is the criterion,
checked five times by the trail rather than once by a test.

**What the trail now shows that it could not this morning.** A phase directory where all four
artifacts agree:

    output.md      0.1.0-dev+b54626e.dirty
    digest.md      0.1.0-dev+b54626e.dirty
    gate.yaml      0.1.0-dev+b54626e.dirty
    learning.yaml  0.1.0-dev+b54626e.dirty

**Evidence is self-reported and local**, with the pipeline running the same commands on the push.

<!-- xeno:section:gaps -->
## Gaps

**`plugin_version` is still `0.1.0-dev` in all four.** The four artifacts now agree about the runner
and all four are wrong together about the plugin: `model.PluginVersion` is a constant whose comment
says there is no plugin yet to have a commit, and `.xeno/plugin/.claude-plugin/plugin.json` declares
`"version": "0.1.0"`, which it does not match either. #177. This intent makes the disagreement
between artifacts go away and leaves the one between the field and the thing it names.

**Thirty intents' worth of records keep the typed header.** Sealed, correctly, and the repository
now holds two provenances for one field with nothing marking which is which — the third time that
has been true today. The commit is the boundary and `runner_version` is what dates it.

**Nothing makes the agent use the command.** The skill says to and a skill instructs. A record
written by hand still passes every gate, carrying `0.1.0-dev`, because G-Learning checks the header
fields are present and not that they are true. The same shape as `--tool-version` before the
variable, and the same answer: it is cooperation until something reads the value.

**No test asserts the skill's text**, only that the YAML block is gone from it. A skill that named
the command in a sentence telling the reader to ignore it would pass.

**The content is unverifiable and deliberately so.** An observation can say anything; the gate checks
that four keys are present and that the category is one of four. That is section 10's arrangement —
the record is an input to a review — and nothing here changes it. What the command removes is the
chance to get the header wrong, not the chance to write a poor observation.

**The empty record is indistinguishable from a phase nobody thought about.** `--no-finding` is one
flag, and section 10 calls it honest. It is also the cheapest way to satisfy G-Learning, and nothing
can tell a phase that genuinely produced nothing from one whose author did not look.
