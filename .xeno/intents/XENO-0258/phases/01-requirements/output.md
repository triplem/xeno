---
intent: github.com/triplem/xeno#256
phase: 01-requirements
created: "2026-10-06T08:35:09Z"
schema_version: "1.0"
runner_version: dev+5163d1b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 651e5597c812abd78f7455adcbc1d5167f532954fe316021cb1e0b2e55bd28fe
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. `CLAUDE.md` gains one paragraph under Conventions, with no new heading, beside the one about
   putting a decision one at a time.

2. It says the rule: a negative result from a check over a path or a tree is evidence only when the
   thing checked exists, so the thing is recreated and the check re-run before the finding is
   written down.

3. It says why here rather than in general: a finding goes into an artifact that is sealed, so an
   unverified negative is permanent and the correction has to live somewhere else.

4. It names the instance — the `git check-ignore` run against a just-deleted path, and that the
   claim reached two sealed phases and a pull request — because a convention with an instance is
   checkable and one without is a preference.

5. It says nothing checks it, as the sequencing paragraph does, so the paragraph does not read as
   something a gate enforces.

6. No new heading, no list, and the file stays short enough that the addition is defensible against
   its own last line. 87 lines before; the paragraph is counted in the verification phase.

7. No row in `docs/clause-readers.md`. That document is a pass over the two normative documents and
   this is a convention in neither, which is the line #255 drew for the sequencing convention.

8. No normative document is touched, and `git diff --stat` shows one file.

9. `./xeno gate verify` exits 0 with the 423 verdicts that exist now intact, plus this intent's own
   judged phases.

10. `go build`, `go test ./...`, `go vet ./...` pass and `gofmt -l` outside `vendor/` prints nothing.
    None can be affected by a Markdown file.

11. Markdown prose stays within 88 columns.

12. One commit, `Closes #256`, and the issue carries `wp0`.

<!-- xeno:section:non-goals -->
## Non goals

Not a mechanism. Nothing can check this: a gate would have to know which of a command's two silences
it received, and A90's finding is that a reader which cannot fail is worse than none. The paragraph
says nothing checks it rather than proposing something that would.

Not a row in `docs/clause-readers.md`. The audit is a pass over the two normative documents and a
convention in `CLAUDE.md` is a clause in neither. #255 declined a row for the sequencing convention
on that reasoning and declining it again is what keeps the reasoning a rule.

Not a repair of XENO-0249. What is sealed is never rewritten; the correction is in #254, in a comment
on PR #246, and in this intent's own artifacts.

Not a list of the tools that share the ambiguity. `test -f`, `grep -l` and `find` all answer the same
way for absent and non-matching, and naming each would turn a paragraph into a table. The paragraph
names the class and one instance; #256 holds the table.

Not the second instance. #212's intake read the wrong field and drew the opposite conclusion, which
is the same shape by a different mechanism; one example earns the lines and the issue carries both.

Not general advice about investigation. The paragraph is framed on sealing, because that is what
makes this a property of this process rather than of carefulness, and a file carrying project context
should not acquire a line that would be true of any repository.

Not a change to how `CLAUDE.md` is organised. One paragraph, under the heading that already covers
how the work is done, no new section.

<!-- xeno:section:constraints -->
## Constraints

`CLAUDE.md` asks not to be added to, in its own last line, and this session has already added nine
lines to it in #259. That is the binding constraint on length: the paragraph has to be worth being
read before every request of every session, and the test is behavioural — would a session reading it
have redone the three-second check that would have caught the gitignore claim.

The framing is not free. A paragraph saying "verify your checks" would be true of any repository and
would be the first line in that file not about this project; one saying "a finding is sealed, so an
unverified negative is permanent" is specific and is the reason the convention belongs there at all.
Getting that wrong does not make the paragraph false, it makes it misplaced.

Nothing can enforce it, and the paragraph has to admit that. A90 is the row, the sequencing paragraph
is the precedent for saying so in the same breath, and a convention that implied a reader would be
the defect `docs/clause-readers.md` catalogues.

The instance has to be the true one. The claim was that the tree is not gitignored; the entry is at
`.gitignore:10` from #200; the cause was a check against a path deleted moments earlier. All three
were re-checked while writing the intake rather than recalled, because a paragraph about verifying a
negative resting on a remembered negative would be its own counter-example.

Markdown prose wraps at 88 and `CLAUDE.md` has no tables, so the whole paragraph is bound by it.

One intent, one branch, `Closes #256`, and the issue carries `wp0`.
