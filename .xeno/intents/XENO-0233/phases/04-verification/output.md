---
intent: github.com/triplem/xeno#177
phase: 04-verification
created: "2026-10-03T17:06:24Z"
schema_version: "1.0"
runner_version: dev+a18c1f3.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9200ab5300824cae228c25689a9a8a0fd07021b5b7dc9d06f9087a0d4bd9a08c
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

Each acceptance criterion of P1 against what holds it.

**`plugin_version` is the vendored plugin's own** —
`TestThePluginVersionComesFromTheVendoredPlugin`, which writes a manifest declaring `0.26.0` and
reads the field back, then asserts it is not `model.RunnerVersion`, which is the confusion the
constant and the ldflags made together.

**Absent where no plugin is vendored, and reported** —
`TestWithoutAVendoredPluginTheVersionIsAbsentAndReported`, which removes the plugin, checks the
field is empty on a created intent, runs a phase and looks for a G-Schema finding naming
`plugin_version`. Both halves, because absence without the finding would be a quiet gap.

**The release sets the runner's version and not the plugin's** — read in the diff: one `-X` where
there were two. Held by the file, which is the only place it can be held; a test would be testing
the workflow runner.

**The lock carries the plugin block** — `TestTheLockCarriesThePluginsVersionAndHash`, which checks
the version against the manifest, the hash for sixty-four hex characters, and that the lock's header
agrees with its own block, which is the comparison G-Supply will make.

**Absent together where there is no plugin** — `TestWithoutAPluginTheLockCarriesNoBlock`.

**The tree hash is defined to the byte** — cross-checked against a shell pipeline that implements
the doc comment's definition independently, and the two agree. Recorded in the results; a Go test
asserting a literal hash would pin the tree rather than the definition.

**A development build claims no version** —
`TestTheStampClaimsNoVersionAndItsMetadataIsWellFormed`, which asserts the metadata after `+` is
well formed and that the part before it does *not* match a semver prefix. The second half is the
decision; the first is what was worth keeping from the test it replaced.

**The manifest declares the newest tag** — read, and `git describe --tags` is how it was chosen. It
has already moved once during this intent, which is P3's first deviation.

**The entry point normalises what section 7 lists and not `XENO_PLUGIN_ROOT`** —
`TestTheEntryPointStartsTheRunnerFromThePath`, which checks the script is there and executable,
that it `exec`s `xeno` from the path, that it names both variables, and that it does *not* export
`XENO_PLUGIN_ROOT`. The last assertion is the one that would catch somebody adding it later without
reading why it is absent.

**The hook calls it plugin-relative** — `TestThePluginCarriesTheHookWiring`, rewritten: the command
is `${CLAUDE_PLUGIN_ROOT}/bin/xeno-env.sh cost turn` and must not contain `.xeno/plugin`.

**`XENO_HARNESS` reaches `tool` and beats the declaration** —
`TestTheHarnessIsRecordedOverTheProjectsDeclaration`, both directions in one test: `codex` over the
project's `claude-code`, then nothing exported and the project's declaration standing.

**Nothing branches on the harness** — the `verify` job's step, run locally as CI runs it: one
mention, no comparison.

**The trail is unaffected** — `gate verify` at 275 and exit 0, with the code captured rather than
piped.

<!-- xeno:section:results -->
## Results

**`go test ./...`** — eighteen packages ok. Thirteen cases are new or retargeted.

**`gofmt -l .` outside `vendor/`** and **`go vet ./...`** — nothing.

**`./xeno gate verify`** — `verified 275 verdicts`, exit 0, no divergence. Exit code captured into a
variable: three times today a `| head -1` reported `head`'s status and hid a failure, once hiding a
sealed artifact this branch had modified.

**Measured before the tests were written.**

    plugin.Version(".")  ->  0.30.0
    plugin.Hash(".")     ->  the same value a sha256sum pipeline produces over the same tree
    ./xeno version       ->  dev+<commit>.dirty

The pipeline is `find . -type f | LC_ALL=C sort`, each file's normalised content hashed, one line of
`<sum>  <path>` each, the stream hashed. It implements the doc comment's definition independently
and agrees byte for byte, which is how `artifacts_hash` was checked when it was written.

**With the plugin removed:** the created intent's `plugin_version` is absent, the lock carries no
`plugin` block, and the phase's G-Schema reports `required field missing: plugin_version`. All three,
which is the arrangement A35 and section 5 produce between them.

**The harness, end to end.** With `XENO_HARNESS=claude-code` exported, this intent's own artifacts
record `tool: claude-code` — and `project.yaml` declares the same, so the precedence is held by the
test rather than by the trail. The trail holds that the variable reaches the field at all.

**This intent's own four header fields**, which is what it is about:

    runner_version: dev+<commit>.dirty
    plugin_version: 0.30.0
    tool: claude-code
    tool_version: 2.1.276

Four fields, four sources, none of them a constant. Before this intent two of the four were.

<!-- xeno:section:gaps -->
## Gaps

**The resolution order is not built, and that is the larger half of #183.** `--plugin-root` and
`XENO_PLUGIN_ROOT` are still read by nothing, so section 7's normalised environment is three
variables specified and one of them, `XENO_PLUGIN_DATA`, set by the entry point and read by nobody.
The decision is section 7's and the measurements are in A84 and in the script.

**Nothing keeps the manifest in step with the next release**, and this intent is the proof: it was
bumped twice in one session because a release landed while the branch was open. #177's open clause.

**No statement in section 16.** #177 asks for the trail's version shapes to be explained where that
section lists what the record does not mean. The README carries it instead, which a reader of the
process definition will not find.

**G-Supply still does not exist**, so the lock's new block is written and read back by nothing. The
comparison section 5 describes — the frontmatter names what was used, the lock proves it with a
hash — is now possible and not performed. That is the way round that leaves the question answerable,
but it is not an answer.

**The plugin hash pins the tree and not the plugin's identity.** Any edit under `.xeno/plugin/`
changes it, including a comment in a skill, so two vendored copies of one released plugin that differ
in whitespace hash differently. That is what a hash over a tree is; whether G-Supply should compare
it strictly or compare the manifest's version first is that gate's decision.

**`plugin_version` is now falsifiable and nothing falsifies it.** A project can vendor a plugin whose
manifest declares any number. The field says what the manifest says, which is strictly better than a
constant that said what the runner said, and it is still a self-report.

**The entry point is untested as a whole.** Its shape is asserted — present, executable, `exec`s
`xeno`, names two variables, does not export a third — and its behaviour under a real client is not.
Nothing in CI runs it, because running it means a harness.
