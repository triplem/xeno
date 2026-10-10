---
intent: github.com/triplem/xeno#336
phase: 01-requirements
created: "2026-10-10T13:29:16Z"
schema_version: "1.0"
runner_version: dev+5f645cb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 27ebaf271c8c7b4cd1d79b2ea88aa47e31c306220fbc67b4c50e199e79369ec6
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. **Section 4.1 carries a sentence naming the design, and says it is not a candidate.**
   `rajistics-demo/sdlc-automation-github-demo` runs on the OpenHands Automations API, so
   it is an instance of the platform 4.1 chose rather than an alternative to it. A reader
   who takes it for a sixth candidate has been misled by the sentence rather than informed.

2. **The sentence points at section 12's paragraph and says the design is the shape it
   refuses.** "Issue commands are deliberately absent. Something has to receive them, and
   every way of doing that is a component to build and operate." The demo has such a
   component: a hosted automation. The sentence has to make the paragraph the thing the
   reader is sent to, because the paragraph is the decision and the sentence is the
   evidence about it.

3. **It strengthens the 4.1 decision rather than competing with it.** That is the issue's
   wording and it is a claim about what the sentence asserts: the design working elsewhere
   is a reason to keep the paragraph, since what it costs somebody else is visible. A
   sentence that read as a reservation about 4.1 would fail this.

4. **The contrast with #330 is drawn and is accurate.** A label gates work in both places
   and the mechanisms are not the same. `Issue.Approval()` in
   `internal/model/identity.go` is handed an issue the runner already fetched and answers
   whether both halves are present; nothing subscribes, nothing is delivered, there is no
   receiver. The sentence may not describe #330 as a trigger, and may not describe the
   demo's labels as preconditions.

5. **Sources carries the repository at a pinned commit.** In the form that section's own
   paragraph already uses for OpenSpec — the repository, the commit, what was read, and
   the date. `2ddf6c91791d95b8de8ff38b683241743ec12049`, read 2026-10-10.

6. **Every claim the sentence makes was read at that commit.** Not taken from the issue.
   The four trigger labels, the status labels beside them, the Automations API as the
   receiver, the templates, and `openspec/` are each in the tree at that sha, and the
   absence of `.github/workflows` is too. A claim this phase cannot point at is a claim
   the sentence does not make.

7. **Nothing in section 12 moves.** `git diff main -- docs/process-definition.md` is
   empty. A trigger label is a specification change before it is anything, #338 holds that
   question, and this intent is not entitled to answer it by writing an evaluation
   sentence.

8. **Nothing else in the evaluation moves.** `git diff main --
   docs/orchestrator-evaluation.md` touches 4.1 and Sources and no other section. 4.2 to
   4.5, section 5's decision, and section 9's treatment of OpenSpec are not this intent's
   business.

9. **It is one sentence, not a section.** The issue says "as one sentence" and the
   distinction is load-bearing: section 9 is what a candidate weighed in earnest gets, and
   #334 is to give AI-DLC the same. A paragraph here would make the demo look like a
   third such treatment. A sentence may carry a subordinate clause and the pointer; it may
   not become a table or a list.

10. **The page still builds and wraps.** Markdown prose at 88 characters, and the docs
    workflow green on the change.

11. **The whole gate suite passes on the tree as it stands.** `go build`,
    `go test ./...`, `gofmt -l .` empty outside `vendor/`, `go vet ./...`, and
    `./xeno gate verify` at exit 0.

<!-- xeno:section:non-goals -->
## Non goals

**Changing section 12.** The paragraph stands. A trigger label is a specification change
first, the use case it would need is #338's business, and that issue now has one from the
maintainer. Nothing here anticipates how it will be answered.

**Taking any of the mechanisms.** Status labels, the templates as an approval boundary,
evidence as comments and files on the pull request, and `openspec/` for change artifacts
are each named by the sentence and adopted by none of it. Each is either already answered
here in another form or an issue of its own.

**Giving the demo a section.** Section 9's treatment — what was evaluated, pinned; the rows;
the decision; the conditions for revisiting; the mechanisms worth taking — is for a thing
weighed as a candidate. This is not one, and #334 is the issue that gives AI-DLC that
treatment. One sentence is the whole of the ask.

**Reopening the OpenHands choice.** 4.1 is decided. A design running on the chosen platform
is a reason to record it and not a reason to weigh it again.

**Evaluating zenflow, which #323 also asks about.** #323 is the parent and carries three
readings; this is one of them, scoped to one sentence about one repository. The others are
their own issues.

**Reconciling the demo's own design decisions with this project's.** The demo uses a status
label where this project uses a phase directory and a verdict, and an `openspec/` directory
where this project has six phases. Comparing them properly is the shape of section 9, which
is the previous non-goal.

**Measuring anything about the demo.** Nine stars and one author are facts about adoption
and they are already in the issue; the sentence is about the design and not about its
popularity. No benchmark, no run, nothing installed.

**Updating `docs/process-definition.md` or `docs/implementation-plan.md` in any way.** Both
are normative and neither is this intent's subject.

<!-- xeno:section:constraints -->
## Constraints

**The standing rule about the documents.** `docs/orchestrator-evaluation.md` is not one of
the two normative documents — the process definition and the implementation plan are — and
its history shows it edited alongside the work it records. #336 is approved and names the
sentence, the section and the done-when, so the wording is delegated. What is not delegated
is section 12, which stays untouched.

**The sentence has to survive the paragraph it points at being read next.** A reader who
follows the pointer gets the whole of section 12's refusal, including the half the sentence
does not quote: "Neither earns its keep while the result is a phase invocation somebody can
type." So the sentence may not claim the design is wrong, only that it is the shape with the
receiver, and the receiver is what the paragraph declines to build and operate.

**Nine stars, one author, no licence.** The absence of a licence file is a fact about reuse
and it constrains what the sentence may suggest: nothing in that repository is available to
copy under a licence, so the sentence names a design rather than a source of code. The
Sources entry is a citation and not an invitation.

**A pinned commit and not a tag.** The repository publishes no releases, so there is nothing
to pin but a commit. That is a departure from section 9's OpenSpec entry, which pins a tag
and a registry version, and the reason belongs in the entry rather than in a reader's head.

**The repository may move under the pin.** It is one author's demo and was pushed to on
2026-10-09, two days after the commit being cited. The entry says what was read and when,
which is the only thing a citation of a moving repository can promise.

**88 characters for prose.** The document is Markdown prose and wraps there; tables and code
blocks do not.

**One sentence has to carry four claims.** Not a candidate, the refused shape running, a
strengthening of 4.1 rather than a reservation, and the #330 contrast. That is the real
constraint on the writing: four claims in a sentence is a long sentence, and the alternative
is a short paragraph, which criterion 9 forbids. Where the two collide, the criterion wins
and the sentence takes a subordinate clause.
