---
intent: github.com/triplem/xeno#181
phase: 04-verification
created: "2026-10-03T12:41:54Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6cbeac4.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d7876d0dbc97d1d539856999d9157b37f587893eb714f7e861d5a812bd16a155
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

**The flag writes the field into `output.md`** — `TestTheReportedToolVersionReachesBothArtifacts`
reads it back off the artifact, and `TestTheToolVersionFlagReachesBothArtifacts` does the same
through the command line, which is where somebody types it.

**`phase finish` writes a digest carrying the same value from `output.md`** — the second half of
both tests. The surface test deliberately omits `--tool-version` on the `phase finish` call, so what
it asserts is the copy and not the flag.

**A second `phase finish` that reports nothing keeps it** —
`TestASecondFinishKeepsTheToolVersionWithoutBeingToldAgain`, which finishes once with the value set,
clears `ToolVersion`, finishes again and reads the digest. This is the criterion the intent turns on
and the one that fails by design under every alternative P2 refused.

**A reported version beats the recorded one** — `TestAReportedVersionBeatsTheRecordedOne`, which
records one version in `output.md`, reports a different one at finish, and expects the reported one.

**A later section write that reports nothing does not erase it** —
`TestALaterSectionWriteDoesNotEraseTheVersion`, two section writes with the value cleared between
them.

**No report is the field absent in both files** — `TestWithoutAReportTheFieldStaysAbsent`, which
checks `output.md` for an empty field and the digest's frontmatter map for the key being absent
rather than empty. Those are different assertions on purpose: a key present and empty would pass a
value check and fail A35.

**G-Schema passes on both artifacts with nothing hand edited** — this intent's own trail. All six
phases were written with the flag; twelve artifacts carry the field; no file was edited after a
command wrote it.

**The skills tell the agent to report it** — read rather than tested: `grep -c tool-version` over
`.xeno/plugin/skills/` is six, one per phase skill, and `xeno-learning` has no `section set` to
carry it.

**Nothing else changes** — the existing suite, unchanged and passing, and `./xeno gate verify` over
the whole trail.

<!-- xeno:section:results -->
## Results

**`go build ./cmd/xeno`** — builds.

**`go test ./...`** — eighteen packages, all `ok`, none skipped. Six cases are new: five in
`internal/runner`, one in `cmd/xeno`.

**`gofmt -l .` outside `vendor/`** — nothing. **`go vet ./...`** — nothing.

**`./xeno gate verify`** — `verified 247 verdicts`, exit 0, no divergence, nothing red, nothing
provisional.

**The trail is the result that matters.** Twelve artifacts across this intent's six phases, every
one carrying `tool_version: 2.1.276`, written by `section set` and `phase finish` and edited by
nothing. **P0 went green on its first `phase finish`** — the first time in this repository, where
every previous phase of every previous intent went red once on this field and was corrected by hand.

**Measured on a throwaway tree before the trail was written**, because a claim about absence cannot
be made from an intent that always passes the flag:

- `section set --tool-version 2.1.276` then `phase finish` with no flag: both files carry the value,
  G-Schema passes.
- `phase finish` a second time with no flag: the value is still there.
- A phase whose sections were written with no flag at all: `grep -c '^tool_version:'` on its
  `output.md` is 0, the field absent exactly as before.

**Evidence is self-reported and local.** The gate path makes no network call; the pipeline runs the
same commands on a push, which is where the merge request's evidence comes from.

<!-- xeno:section:gaps -->
## Gaps

**The mechanism section 7 designed is still unbuilt, and this intent names it rather than closing
it.** `XENO_PLUGIN_ROOT`, `XENO_PLUGIN_DATA` and `XENO_HARNESS` are specified and read by nothing.
A flag is a channel, not the channel: it has to be passed on every phase by whoever drives the
runner, where an environment variable set once by an entry point would be passed by nobody. #181
says so and A80 records it; it wants its own issue and a specification commit in front of it.

**Nothing makes the agent pass the flag.** The skills say to, and a skill is instruction rather than
enforcement. A phase written without it is exactly as red as it was before this change, which is the
honest outcome but means the hand edit is avoided by cooperation and not by construction. The
environment variable above is what would make it unnecessary to remember.

**The value is unverifiable.** `--tool-version 9.9.9` is recorded as readily as the truth, and
nothing in the trail could tell. That is true of `model` and `tool` as well, which come from
`project.yaml`, and it is the general shape of A35's territory: the process records what a session
reported about itself.

**No test asserts the skills' text**, only that the flag appears once in each of the six. A skill
that named the flag in a sentence that said not to use it would pass that count.

**The artifacts already written keep their hand-typed value.** Correct and sealed, so the
repository holds two provenances for one field — typed, and reported — with nothing marking which is
which. The commit that introduced the writer is the only boundary, and `runner_version` is what a
reader would have to date it against.
