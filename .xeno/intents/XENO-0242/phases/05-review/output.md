---
intent: github.com/triplem/xeno#188
phase: 05-review
created: "2026-10-03T21:47:05Z"
schema_version: "1.0"
runner_version: dev+efc44a9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 17d9cefa5fee9a04a8b103c429fcf7096e6b8e839adcec942d899b6af3b12c18
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
  - rule: deviations-are-traceable
    result: met
    note: >-
      Four, each naming what it departs from: the --resolves existence check names the design,
      which did not call for it; the rename of summarise names the design sentence it completes;
      the help text change names no criterion and says so; and the missing test for the printed
      mark names the P1 criterion and gives the judgement behind it.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      Two commands and three flags are added and nothing is taken away or behaves differently.
      A repository that runs neither command sees one difference, a mark in intent status, which
      changes no verdict. The one thing worth knowing is that --resolves now belongs to three
      commands and its help text no longer says assumption.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added. go.yaml.in/yaml/v3 was already the one dependency and already imported
      by the runner; errors in the new test file is from the standard library.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Four, each naming what it departs from. The
`--resolves` existence check names the design, which did not call for it, and says what
it buys: a typo is otherwise silent until P5 and leaves something that looks like an
answer in the trail. The rename of `summarise` names the design sentence it completes,
which said both forms read the predicate without saying how the one-intent form reaches
it. The help text of `--resolves` names no criterion and says so, which is what makes it
a deviation rather than a quiet line in the diff. The missing test for the printed mark
names the P1 criterion it departs from and gives the judgement behind it.

**`interface-change-needs-a-migration-note` — met.** Two commands are added, three flags
are added, one flag's help text changes, and nothing is taken away or behaves
differently than before. A repository that never runs either command sees one
difference: a mark in `intent status` on an intent that reached P5 having asked nothing,
which is a figure and changes no verdict. There is nothing to migrate, and the one thing
worth knowing is that `--resolves` now belongs to three commands and its help text no
longer says assumption.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added.
`go.yaml.in/yaml/v3` was already the one dependency and was already imported by the
runner; `errors` in the test file is from the standard library.

**Beyond the three rules.**

*Did this intent do what it says the trail should do?* It is the test the issue sets,
and the answer is in the trail rather than in this paragraph: two questions raised with
options and consequences, two decisions with a person in `decided_by` and the agent in
`proposed_by`, each question settled in a later phase than the one that raised it. Q-1
and D-1 are hand written because the commands did not exist when they were needed, which
P0's learning records as the bootstrap case; Q-2 and D-2 went through the commands. A
reader can tell which is which from the indent, which P4's gaps explain.

*Was the first standing rule kept?* No document changed, and the one option that would
have needed one — strictness between `proposed_by` and `decided_by` — was recorded as
out of scope with the rule as its reason rather than taken because it was tempting.
Section 5 defines both blocks and section 8 defines the exchange, so everything built
here was already specified.

*Was anything invented?* No field, no gate, no rule. The two commands spend from the
surface, which the plan calls a budget rather than a list, and they write only what
section 5 enumerates. The unknown-field decode is the standing rule enforced at the
door: a key nothing enumerates is refused rather than dropped.

*What would a stranger find hardest?* That `question record` reads a document while
`decision record` takes flags. The design says why in one sentence and the alternatives
section records the three shapes that were weighed, so the answer is in the trail; but
it is the first thing somebody will ask, and if the surface grows a third writer of this
kind the asymmetry will need restating rather than repeating.

*What did this intent not do that the issue asked for?* Nothing in its scope. What #188
asks that no single intent can give is the evidence that the habit changed, and P4's
gaps say so plainly: the figure reports on the intents that follow.

<!-- xeno:section:release-notes -->
## Release notes

**Two commands write the exchange section 8 defines.** `xeno question record` reads an
`open_questions` entry — its text, two to four options with a consequence each, the
recommendation, and the free entry — and writes it into the frontmatter of the phase
named, with the next `Q-n` of the intent as its key. `xeno decision record` writes the
`decisions` entry that settles one: `--chosen`, `--reason`, `--by`, and optionally
`--resolves` and `--proposed-by`, with `--withdraw` for section 8's third exit. Both
refuse before writing what the gates would report afterwards, and neither invents the
person.

**`xeno intent status` says when an intent asked nothing.** An intent that has reached
the review phase with no question and no decision in any phase is marked in the listing
and carries a line in the one-intent form. It is a figure and not a finding: no gate
changed and nothing goes red.

**For an existing repository.** Nothing to migrate. `--resolves` now belongs to three
commands and its help text no longer says assumption. Every other command, flag and
artifact behaves exactly as before.

<!-- xeno:section:residual-risk -->
## Residual risk

**The habit is the risk, not the code.** The commands exist and the figure reports;
neither makes anybody put a question to a person. What changes the trail is the agent
raising the question where a decision is being taken in conversation, and the only thing
watching that is a mark in a listing somebody has to read. D-2 accepted exactly that
bargain, with the reason that a figure tuned to be quiet in its first week is not a
figure.

**A question recorded too early is sealed too early.** P3's learning and P4's gaps both
record it: the asking is iterative and the entry is not. The habit that avoids it —
settle the options with the person, then record — is cheap and unenforced, and the
alternative is a `--replace` path that would have to refuse a judged phase.

**The refusal that names a key nobody wrote.** The third refusal in P4's results says
`question Q-3` when no Q-3 exists. Harmless, since the number is not consumed and the
line begins with "refused", and left alone rather than fixed during verification.

**Two blocks are still written by hand.** `evidence` is #208 and `review_checklist` is
WP7's, including the one in this very phase's frontmatter. This intent built the writer
for the two blocks whose gap it was filed about and deliberately not a generic one,
which the design argues for and which leaves the other two exactly as they were.

**Nobody compares the proposer with the decider.** Q-1's third option, not taken, and a
document change before any code if it is ever wanted. Until then a trail could carry the
agent's name in both fields and nothing would notice; that it does not here is
discipline and not enforcement.
