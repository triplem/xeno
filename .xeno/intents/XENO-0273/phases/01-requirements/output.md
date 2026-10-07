---
intent: github.com/triplem/xeno#153
phase: 01-requirements
created: "2026-10-07T14:23:05Z"
schema_version: "1.0"
runner_version: dev+6b48c17.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 43bea369197e1537629b226d47cb453062c32d080bb1a8b8efd4754d42dcc774
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

Numbered, and P4's mapping cites the numbers.

1. **`docs/symbol-index.md` exists and its first fenced block is byte-identical to
   `internal/index/testdata/example-symbols.yaml`.** Not a summary of it, not a reflow, and not
   the comments stripped.

2. **A test fails when the two part**, and it was checked by making them part. It sits in
   `internal/index/index_test.go` beside the test that reads the example.

3. **The page's prose says nothing the example does not.** Every claim in it is either in the
   example's comments or in section 5 or Appendix A, and the page says which.

4. **It says where the format is fixed**, which is section 5 under "A symbol index, not a
   graph", and that the two configuration keys are in Appendix A's `project.yaml` rather than
   in section 5.

5. **It says a project produces the index and Xeno ships no indexer**, which is the first thing
   somebody arriving at the page needs in order to know what is being asked of them.

6. **It says an absent, stale, unreadable or malformed index is not an error**, in the reader's
   own terms: four causes with one outcome, and a phase runs without one.

7. **`docs/README.md` carries an entry for it**, in the group where `commands.md` already sits.

8. **Nothing in `example-symbols.yaml` changes.** It is read by `index_test` and was the
   deliverable of #151.

9. **No normative document changes.**

10. **`go test ./...`, `gofmt -l .`, `go vet ./...` and `xeno gate verify` pass**, and every
    relative link in the new page and in the index entry resolves to a file that exists.

<!-- xeno:section:non-goals -->
## Non goals

**No prose reference page.** Put as the second option and declined. #153's own "what it needs
to say" section reads like a specification for one, and the issue's last paragraph is the
argument against it: a field table cannot be held against a YAML file, so it would be the
second copy, and the first copy has a test.

**No change to `example-symbols.yaml`.** It is read by `index_test`, its annotation was #151's
deliverable, and it already says what the page needs. Editing it to read better as a published
page would be editing a fixture for a reader it was not written for.

**No move out of `testdata`.** It would make the published copy primary and the fixture a
reference to it, which is a change to how `internal/index` is tested. #153 asks for the format
to be published, not relocated.

**No documentation site.** WP16 is unstarted and its generator is under question in #226. This
is a third file a site can later publish, after `commands.md` and
`the-trail-in-this-repository.md`.

**No link checker, and the gap it would close is named.** Nothing notices when a file added
under `docs/` is missing from the index; this page makes that marginally worse and does not
change the argument, which XENO-0270 already recorded as its own change.

**No new facts about the index.** Not about the reader's degradation, not about the
configuration schema, not about what `kind` may hold. Everything is described already and the
page's job is to put the description where a project will find it.

**No generated page.** WP16 says the reference is generated from the code and that remains the
end state. What is here is a committed page with a test, which is a guard rather than the fix,
and the same stopgap `docs/commands.md` carries. Saying so is the point of this entry.

<!-- xeno:section:constraints -->
## Constraints

**The first standing rule.** Neither normative document is edited. Section 5 and Appendix A are
cited by the page, and the implementation plan's WP16 is cited as the owner of the end state.

**One place a fact is written down.** The constraint the whole change exists to honour. The
page carries the example and does not restate it, the prose adds no claim the example or a
normative document does not already make, and a test holds the block against the file.

**A negative result is evidence only when the thing checked was there to be found.** The claim
that the page's prose says nothing new is an absence, so it is checked sentence by sentence
against the example's comments, section 5 and Appendix A rather than asserted.

**Prose wraps at 88 characters**, tables and code blocks do not. The fenced block is the
example's own width and is not reflowed, which is also what the test requires.

**Headings name their section in words.**

**No invented fields, gates, tools or rules.** The change is one page, one test and one index
entry.

**One dependency.** Untouched.

**Every change belongs to a work package and to an intent.** WP16 by subject, which is #153's
own argument for where it belongs rather than WP15, and intent XENO-0273 for issue #153. The
key was passed explicitly, because the sequence counts from one branch.
