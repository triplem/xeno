---
intent: github.com/triplem/xeno#231
phase: 02-design
created: "2026-10-07T09:39:32Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 92f2ab415eaf7e0eeb82de72e8f2c1f7a29150f3fcca730f282ba5ee7529c4b8
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: Rewrite README.md around what Xeno is for, the name and an index; keep four commands and point at xeno --help and a reference page; move the field notes, the two trail shapes, the coverage table and the mutations to docs/the-trail-in-this-repository.md; add docs/README.md as the index and docs/commands.md held against the usage constant by a test.
      rationale: 'Asked for on the issue on 2026-10-07: point to an index file in docs highlighting the process definition, name the most used commands briefly and then link xeno --help and the reference, and move the rest, naming the gate path''s absence of a network call as the example of what is too deep for the README. That corrects the issue''s own done-when, which said the existing material stays further down and listed the no-network-call property among the things a reader needs. It is a thing a reader needs and not a first reader, and section 12 carries it in more detail under Two surfaces, two promises, so it is dropped rather than moved and the index entry names the section.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The README opens with the problem, not with the tool.** "Requirements, design, code and
evidence are produced separately, at different times, partly by people and partly by models"
is section 1's first sentence, and a reader who recognises it knows in one line whether this
applies to them. A reader who does not recognise it is not the audience, and the two sentences
after it — the binding to an intent, and evidence rather than a management system — tell them
so without a list of packages.

**Four commands in the README, not six and not twenty-seven.** `xeno init`, `xeno intent
start`, `xeno phase start`, `xeno phase finish`. They are the sequence somebody runs on their
first intent in the order they run it, which is what makes the block worth reading rather than
scanning. `gate run` is deliberately not among them: the gates run from the phase commands and
from CI, and a reader who types it first learns the wrong thing about when a verdict happens.
Everything else is `--help` and the reference.

**The reference page carries `usage` and nothing else of its own.** No per command prose, no
examples. The page's whole value is that it is the same text `--help` prints and that a test
says so; prose around it would be prose nothing holds, which is the problem being fixed one
level up.

**The test compares the fenced block, not the file.** The page needs a title and a sentence
saying what it is and where it comes from, so the assertion is over the first fenced block.
Comparing whole files would mean the page could carry no explanation, and a page with no
explanation is a file somebody deletes as redundant.

**The index is `docs/README.md` and not `docs/index.md`.** GitHub renders `README.md` when
somebody opens the `docs/` directory, which is where a reader following a link from the top
level arrives. `index.md` is what a static site generator wants, and WP16 can rename it when
there is a generator; today the reader is on GitHub.

**The index marks the normative two and groups the rest by errand.** Not alphabetically, and
not by age. A reader comes to `docs/` with one of four errands — what the process is, what is
being built and in what order, why the tree is the way it is, or how something outside the
process works — and the grouping answers that. The two normative documents are named as such
in their own entries rather than in a preamble, because a reader who follows a link from the
README to one of them never reads the preamble.

**The moved material goes to one page, not four.** `runner_version`, `plugin_version`,
`tool_version`, the two trail shapes, the coverage table and the mutations are all answers to
one question: what does this repository's own trail show and why does it read oddly. Splitting
them by subject would make four pages nobody has a reason to open.

**The moved paragraphs are moved verbatim.** Not improved on the way. A commit that both moves
and rewrites leaves no way to see which happened, and the convention about replacing rather
than editing into is about a paragraph being changed, which these are not.

**The no-network-call paragraph is dropped rather than moved.** It is section 12's, under "Two
surfaces, two promises", in more detail and in the normative place. The maintainer named it as
the example of what should move; moving it would make a second copy of a normative sentence,
so the index entry for the process definition names the section instead.

**`xeno --help` is named before the reference.** A reader at a terminal has it already, and the
page is for somebody reading on the web. Naming the file first would send the person who can
answer their own question to a browser.

<!-- xeno:section:alternatives -->
## Alternatives

**Keep the full command list in the README and hold it with the same test.** The smallest
change that fixes the drift, and it fails the other half of the issue: a first reader meets
twenty-seven lines before they learn what the tool is for. The list is not wrong because it
drifted, it is wrong because it is first.

**Generate the reference from the flag declarations, as WP16 describes.** The end state, and
it would make `usage` derived rather than authoritative. It is a change to how `cmd/xeno`
declares its commands, with a generator in the build and a check that the committed page
matches; the drift this issue is about would be gone by construction rather than by assertion.
Out of scope because it is WP16's package and because the cheap version buys the same
guarantee for this page today. Worth saying plainly: the test here is a guard, and WP16's
generator is the fix.

**Drop the coverage table rather than move it.** It is twelve rows of test names and the
honest question is who reads it. It moves because it is the only statement anywhere of what
this repository's own suite covers against the plan's acceptance criteria, and because
deleting it in the same commit as a restructuring would be a deletion nobody noticed. If it
turns out to have no reader, that is a separate finding with its own argument.

**Put the moved material in `docs/assumptions.md`**, since most of it cites an A number. It
would be wrong in both directions: those paragraphs explain what the trail reads like, which
is not an assumption taken while building, and the register's rows already carry the
assumptions themselves. A reader who wants A85 should find A85, not a paragraph about A85.

**One page under `docs/` for everything, index included.** Fewer files, and it makes the index
the thing a reader scrolls past to reach the content, which is the README's current problem
moved one directory down.

**Leave the README alone and write only the new pages.** Tempting because the README is
accurate about what it describes. It fails the issue entirely: the issue is about what the
README is, not about what is missing beside it.

<!-- xeno:section:impact -->
## Impact

`README.md`: rewritten. From 24 command lines and nine paragraphs of field notes to four
commands and the facts a first reader acts on.

`docs/README.md`: new, the index. One line per file under `docs/`, the two normative ones
marked in their own entries.

`docs/commands.md`: new, the reference. A title, a sentence saying where the text comes from,
and the `usage` constant in a fenced block.

`docs/the-trail-in-this-repository.md`: new, taking the material that moves, verbatim.

`cmd/xeno/main_test.go`: one test, beside `TestEveryCommandInTheUsageResolves`.

**For a first reader.** The three things #231 names are answerable in the first screen: what
it is for, what the name means, where to read on. The command block is four lines in the order
somebody runs them.

**For a reader who was using the README as a reference.** They lose the long list and gain a
page that cannot be wrong. The six commands missing from the old list are in it.

**For the trail.** Nothing. No artifact, gate or verdict is touched, and `xeno gate verify`
recomputes as before. The README is in no context scope that matters: it is read by people.

**For WP16.** Three pages it would otherwise write, and a landing page to rename to
`index.md` when there is a generator. The reference's test is the guard WP16's generator
replaces, which the alternatives say plainly so that WP16 does not inherit it as a decision.

**What gets worse.** Four files where there was one, and a reader who knew the README now has
to follow a link. That is the trade the issue asks for, and the index exists so the link is
one hop.
