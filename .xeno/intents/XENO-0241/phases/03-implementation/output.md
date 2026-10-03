---
intent: github.com/triplem/xeno#206
phase: 03-implementation
created: "2026-10-03T20:52:22Z"
schema_version: "1.0"
runner_version: dev+8574810.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3c6954234d6ace9479022d9b50217c960ef164addbcec49df8350e499aa52e42
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

`internal/git/git.go` gains `Paths`, which lists the files differing between the trees at
two refs, under a directory or over the whole tree. It is a tree comparison and not a log
over the range, for the reason the trail guard already gives: a directory created and
dropped again inside one branch existed at neither end. Rename detection is off, so a move
reports both of its paths. An empty ref is a caller error, as in `Commits`. The package
comment said it starts exactly one subprocess; it now says two, rewritten rather than left
to be contradicted by the code under it.

`internal/runner/runner.go` gains `CompletenessResult` and `Completeness(base, head)`. It
turns the touched paths into intent keys, calls `summarise` for each, and collects the ones
whose state is neither `complete` nor `abandoned`. An intent whose record cannot be read is
collected too, with the reason, and where neither a state nor a reason comes back it is
given one rather than printed blank.

`cmd/xeno/main.go` gains `intent verify` in the dispatch table, a line in the usage block,
and `cmdIntentVerify`. Exit 2 for a range that could not be read, 1 for an intent that
stopped, 0 otherwise; a range touching no intent says so in words. The refusal names both
endings, because a refusal that only says no is one somebody works around.

`.github/workflows/xeno.yml` gains one step, after `gate verify`, on pull requests only and
with the same base as the trail guard.

`internal/scaffold/files/ci-github.yml` and `ci-gitlab.yml` each gain the same call, using
the `{{.BaseRef}}` and `{{.HeadRef}}` they already pass to `gate run`.

Twelve tests. In `internal/git`: the paths under a directory and over the whole tree, a path
added and dropped inside the range, a move reporting both ends, an absent ref, a ref that
does not resolve. In `internal/runner`, over a real repository built on the fixture: a
complete intent passes, an intent stopped at P3 is named with the phase it reached, an
abandoned intent passes, a provisional P5 does not, a range touching no intent has nothing
to check, an unreadable record is unfinished with a reason, and a missing or unresolvable
ref is an error. In `cmd/xeno`: exit 1 naming the intent and the other ending, the empty
range in words, and exit 2 without a range.

`go build`, `go vet`, `gofmt`, `go test ./...` and `xeno gate verify` over 324 verdicts all
pass.

<!-- xeno:section:deviations -->
## Deviations from the design

One deviation from the acceptance criteria, in the test rather than in the behaviour. The
criteria asked for a test that a range touching no intent passes, and said nothing about
where the key comes from when a path under `.xeno/intents/` is not an intent directory at
all. The reader takes the first path segment after the prefix and skips an empty one, which
is the only reading that cannot invent a key; it is not separately tested, because a path
of that shape does not occur in a trail and a test over it would assert the absence of a
case rather than a behaviour.

Nothing else departs. No field, gate, tool or rule, no document, and no specification
change: section 7's G-Complete is untouched and runs where it ran, and what was added is a
second reader of section 8's sentence about how an intent ends.

`CLAUSE-READERS.md` is unchanged, as the design said. Stated again here because the
temptation at this point in the work is to go and tick the finding off, and the document
declines to be an index somebody maintains.

The local run of this check cannot see this intent until it is committed, since the
comparison is between two trees and the trail is in neither until then. That is correct and
is worth writing down: the first real exercise of the step is the pull request this intent
opens, and if it is wrong there it will be wrong loudly.
