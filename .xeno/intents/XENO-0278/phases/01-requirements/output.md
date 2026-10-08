---
intent: github.com/triplem/xeno#330
phase: 01-requirements
created: "2026-10-08T17:39:14Z"
schema_version: "1.0"
runner_version: dev+e22a533.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c2fbb2b030f3739956592e2c0315a7a71cd61d6b1b6fa0a7c7192b57203b3a7f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Each criterion is a state of the tree and a verdict over it. "A configured project" below
is one whose `project.yaml` carries a complete tracker block and whose environment holds the
token the block names.

1. **An unlabelled issue does not become an intent.** In a configured project,
   `xeno intent start --for N` on an issue that does not carry the label `approved` exits 1
   with a refusal that names the label, and leaves no directory behind.

2. **A label without the word is not enough.** The same, on an issue that carries the label
   and has no comment whose first line is the word `approved` (case ignored, surrounding
   space ignored): exit 1, a refusal naming the comment, nothing written.

3. **Both present, the intent starts as before.** With the label and such a comment the
   command writes `intent.yaml` exactly as it does today and prints the key and the id.

4. **A milestone holds the issue until its turn.** An issue whose milestone is open and is
   not the earliest open milestone of the project — ordered by due date, undated ones last,
   then by number — is refused with both milestones named. With `--now` it starts. An issue
   whose milestone is the earliest open one, or is closed, starts without the flag.

5. **A read that cannot happen is a refusal.** With a tracker block and no token, with an
   issue the host answers 404 or 403 to, and with an issue whose qualified id names a host
   the block does not reach, the command refuses and says which; it does not start the
   intent on the strength of nothing. A block that is there and incomplete still fails as
   Appendix A says, naming the field.

6. **Without a tracker block nothing changes.** A project with no block takes the whole
   qualified id by hand and starts the intent without a read, as `TestWithoutATracker
   BlockTheWholeIdIsGivenByHand` asserts today.

7. **The intake says by what it was authorised.** `phase start` at P0 writes, above the
   quoted issue, a sentence naming who wrote the approving comment, when, and the rest of
   the comment as the reason; where the issue's milestone is still not the earliest open
   one, a second sentence says the intent was started ahead of it. Where no approval is
   found at that moment, the sentence says that instead of saying nothing.

8. **Both adapters read the same five things in their own words.** GitHub: `labels[].name`,
   `milestone`, the issue's comments by `user.login`, `created_at` and `body`, and the
   repository's open milestones. GitLab: `labels[]`, `milestone`, the issue's notes by
   `author.username`, `created_at` and `body` with system notes left out, and the project's
   active milestones. Comments beyond the first page are followed. A test on each host
   asserts the paths asked and the fields read.

9. **The flag is in both copies.** `[--now]` appears on the `intent start` line of the
   `usage` constant and of `docs/commands.md`, which `TestTheDispatchTableAndTheUsageAgree`
   and the doc test hold together, and the notes on `intent start` in `docs/commands.md`
   say what is refused and what the flag does.

10. **The suite and the gates stay green.** `go test ./...`, `gofmt -l`, `go vet ./...`
    and `./xeno gate verify` all exit 0 on the branch, and the tests that start an intent
    against a tracker block today are given a host that answers approved rather than
    weakened.

11. **The register carries the two implementation decisions.** One row in
    `docs/assumptions.md`, A103, for the refusal on a read that cannot happen and the
    sentence recomputed at P0.

<!-- xeno:section:non-goals -->
## Non goals

- **Comments in the intake.** The comments are read from this intent on, and quoting them
  beside the body is XENO-0277's P0 learning and this intent's own; it is a change to what
  the problem section quotes and gets its own issue.
- **A configurable label or word.** The clause fixes both. A field would be an addition to
  Appendix A, which is a specification change nobody asked for.
- **Judging the comment's author.** Who may approve is the host's business through the
  label's right; the comment is quoted, not vetted.
- **Jira, assignees, due dates.** Weighed on the issue and set aside; a third host is 1.1.
- **Anything in a gate.** The clause says why, and the import graph test says it
  mechanically.

<!-- xeno:section:constraints -->
## Constraints

- Section 12's *Starting an intent* paragraph is the specification; nothing here may read
  it more loosely, and nothing may add to what the artifacts carry beyond the sentence it
  names.
- WP12's contract stays four operations, and the adapter still knows no relationship
  between issue key and repository: project and key are parameters of every call.
- `internal/gates` imports nothing under `internal/host`, which `host_test` asserts.
- One vendored dependency; this needs none.
- A refusal writes nothing: a refused start leaves no directory, as
  `TestStartingAnIntentNeedsTheIssue` asserts for the existing refusal.
- The exit code staircase: a refusal is 1 with the reason on standard error, a credential
  that does not work or an unreadable answer is 2.
- No sealed artifact changes; the trail's 78 finished intents are untouched.
