---
intent: github.com/triplem/xeno#231
phase: 03-implementation
created: "2026-10-07T09:45:42Z"
schema_version: "1.0"
runner_version: dev+0768c44.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a3ae95c6123c085152e9edf90f51f9f202611c0195021d91d6f8d8b05d43ab8a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

`README.md`, rewritten. It opens with section 1's statement of the problem — artifacts
produced separately, at different times, partly by people and partly by models, held together
only by their binding to an intent — then what Xeno does about it in one sentence, then that
what it produces is evidence and not a management system. The name is a paragraph, with
Appendix C named for the rest. Then the index, with the process definition named as the
normative one to read first. Then four commands in the order somebody runs them, the two
sentences about the next step and `--root`, `xeno --help` and the reference. Then telemetry.
Then one paragraph on this repository's own trail, pointing at the page that carries it, and
one on the register.

`docs/README.md`, new. Every file in `docs/` with one line on what it is for, grouped by the
errand a reader arrives with: what the process is, running it, what is being built, why the
tree is the way it is, and where v2 goes. The two normative documents are marked in their own
entries. `CONTRIBUTING.md`, `SECURITY.md` and `CLAUDE.md` are named at the end as being beside
rather than inside.

`docs/commands.md`, new. A title, three sentences saying the block is what `--help` prints and
that a test holds it, then the `usage` constant in a fenced block, then the per command notes
that moved out of the README.

`docs/the-trail-in-this-repository.md`, new. A heading and three sentences of framing, then the
material that moved: what the three version fields say in this trail and why, the two shapes of
the intents directory, the coverage table against the plan's acceptance criteria, and the two
deliberate mutations.

`cmd/xeno/main_test.go`, one test. `TestThePublishedReferenceIsTheUsage` reads
`docs/commands.md`, takes its first fenced block and compares it against `usage` with the
trailing newline trimmed. It sits beside `TestTheDispatchTableAndTheUsageAgree` and its comment
says why the comparison is over the block and not the file.

`go build`, `go test ./...`, `gofmt -l .` outside `vendor/`, `go vet ./...` and
`xeno gate verify` over 498 verdicts all pass. Every relative link in the four files resolves
to a file that exists, checked by resolving each against the tree rather than by reading them.

<!-- xeno:section:deviations -->
## Deviations from the design

Four, and the first two are corrections to P2 found by checking criterion 10 rather than
trusting it.

**The reference page carries per command prose after all.** P2 decided it would not: "No per
command prose, no examples. The page's whole value is that it is the same text `--help` prints
and that a test says so." Then the check for criterion 10 — that no paragraph of the old README
is lost — found five that had nowhere to go: what `--for` accepts and how the key continues the
sequence, what `learning record` writes and what `--no-finding` means, what `intent status`
without an intent lists and what `cost turn` reads, that every changing command names the next
step and so does `intent status`, and the network paragraph below. They are about behaviour,
not about this repository, so the trail page is the wrong home, and dropping them would have
failed criterion 10, which is the stronger requirement. They sit under their own heading after
the block, and the test still compares only the first fenced block, which is what P2's other
decision was for.

**The no-network-call paragraph is kept, not dropped.** P2 dropped it on the grounds that
section 12 carries it under "Two surfaces, two promises". Section 12 carries the first half.
The second — "`xeno enforcement check` is the one command that does, because asking a host what
it enforces is the one question the repository cannot answer about itself" — is in no normative
document, so dropping it would have lost a fact. It moved with the other notes. The same
reading corrected the index: a first draft of `docs/README.md` said starting a phase is the one
thing that may use the network, and section 12 says "everything else in the CLI may use the
network", so the entry now says that instead.

**One heading was renamed and two paths rewritten on the way.** The migrating material was
moved verbatim, as P2 required, with two exceptions the move itself forces. `## The trail in
this repository` would have repeated the new page's own title, so it is now `## The two shapes
of the intents directory`. And `docs/m0-gate-job.md` and `docs/assumptions.md` were written
from the repository root and no longer resolve from inside `docs/`, so they read
`m0-gate-job.md` and `assumptions.md`, with "beside this file" added to the first. A94 is the
register row about exactly this class of stale reference, and the check that found these is the
same one.

**The old README's `learning record` paragraph was two paragraphs.** It ran into the
development-build paragraph with no blank line between them, so the first extraction carried
the second into `docs/commands.md`, where it does not belong. Split, and the development-build
paragraph appears once, on the trail page. Found by counting its occurrences across all four
files rather than by reading them.

Nothing else departs from P2. Four commands and not six, `gate run` deliberately not among
them, the index at `docs/README.md` rather than `index.md`, one page for the moved material,
and `xeno --help` named before the file.
