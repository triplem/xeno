---
intent: github.com/triplem/xeno#231
phase: 01-requirements
created: "2026-10-07T09:38:22Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7bdda5fa86f0286f5e20c8ad0e133dd87040477b99a644f6e2b99285af6e30bd
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

1. **`README.md` opens with what Xeno is for**, in two or three sentences somebody outside
   this repository can act on, and not with a list of the runner's parts.

2. **It says what the problem is before it says what the tool does.** Artifacts produced
   separately, at different times, partly by people and partly by models, bound to an intent;
   that is section 1's statement and it is what makes the rest legible.

3. **It says where the name comes from**, in a sentence, and points at Appendix C rather than
   reproducing it.

4. **It links `docs/README.md`**, and names the process definition as the normative one at the
   point of linking rather than leaving a reader to find out.

5. **`docs/README.md` exists and has one line per document under `docs/`**, saying what each is
   for, with the process definition and the implementation plan marked normative and the others
   grouped by what a reader would come to them for. Every file in `docs/` appears.

6. **The README names the handful of commands a reader starts with** and no more, and points at
   `xeno --help` and `docs/commands.md` for the rest.

7. **`docs/commands.md` exists and its command block is byte-identical to `usage`.** Not a
   summary of it and not a reflow.

8. **A test fails if the two diverge.** It sits beside the one that reads the command names out
   of `usage`, and it was checked by making them diverge.

9. **`docs/the-trail-in-this-repository.md` carries everything that moved**, unchanged in
   substance: the three version fields and what they say here, the two shapes of the intents
   directory, the coverage table, and the two mutations.

10. **Nothing is lost.** Every paragraph of the old README is either in the new one or in one
    of the new pages, and the no-network-call property is the one deliberate exception, because
    section 12 carries it and the README now points there.

11. **No normative document changes.** The diff touches `README.md`, three new files under
    `docs/`, and one test.

12. **`go test ./...`, `gofmt -l .`, `go vet ./...` and `xeno gate verify` pass**, and every
    relative link in the new files resolves to a file that exists.

<!-- xeno:section:non-goals -->
## Non goals

**No change to the process definition or the implementation plan.** The purpose and the name
are pointed at and quoted in the README's own words, not moved. The plan's refusal to record
the name stays: it is a reasonable thing for a plan to say, and the README is where the answer
belongs.

**The reference is not generated from the flag declarations.** WP16 describes that and it is
the right end state. `usage` is a hand written constant; what this intent buys is that the
page cannot drift from it, which is the half that was costing something. Generating the
constant itself is a larger change and WP16's.

**No documentation site.** No `mkdocs.yml`, no navigation, no theme. WP16 is unstarted and its
generator is under question in #226, so these are files a site can publish and not a site.

**No shortening of `usage`.** It is the reference's source and a reader of `--help` wants the
whole of it. The README carries fewer commands; `--help` carries the same number it did.

**The skills' command lists stay as they are.** Seven skills, four to seven lines each, naming
the commands of their own phase. Each is correct because it names only what its phase runs,
and holding them against `usage` would assert a sameness that is not wanted.

**No new facts.** Everything in the new pages is already in the README or in a normative
document. A page that explained something for the first time would be documentation nobody
reviewed, which is what the first standing rule is about.

**The coverage table is not re-measured.** It moves as it stands. Checking twelve work
packages against their tests is worth doing and is not this issue; a table that moved and
changed in one commit would make both halves hard to read.

**No link checker in CI.** The links are checked once, here, by resolving each against the
tree. A check that ran on every push would be worth having and is a change to `xeno.yml` with
its own argument, not a step in a README rewrite.

<!-- xeno:section:constraints -->
## Constraints

**The first standing rule.** `docs/process-definition.md` and `docs/implementation-plan.md` are
not the agent's to edit and neither is edited. The three new files under `docs/` are not
normative and the maintainer asked for two of them by name on the issue; the third is where the
material the same comment said to move has to land.

**A paragraph changed a second time is replaced rather than edited into.** The README is being
restructured, so it is written rather than patched, and read back as a document rather than as a
diff. The paragraphs that move are moved verbatim, which is the one case where not rewriting is
the point: changing them on the way would hide what moved.

**Headings name their section in words.** No numbers in the README's or the new pages'
headings, and the index's entries are the file names with what each is for.

**Markdown prose wraps at 88 characters**, tables and code blocks do not. The coverage table
and the command block are the exceptions and keep their width.

**No invented fields, gates, tools or rules.** Nothing is added to an artifact or to the gate
list. The one new test asserts an equality between a file and a constant.

**One dependency.** Untouched.

**The `usage` constant is the single source for the command reference.** The page is held
against it and the existing test reads names out of it. Nothing may introduce a third place
where a command name is written down.

**Every change belongs to a work package and to an intent.** WP16 by subject, intent XENO-0270
for issue #231. The key was passed explicitly, because two intents are open on other branches
and the sequence counts from this one.
