---
intent: github.com/triplem/xeno#336
phase: 05-review
created: "2026-10-10T13:44:34Z"
schema_version: "1.0"
runner_version: dev+5f645cb
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: acdb72f140a1113e695ee6924416e275d0386fea87f8d959266756d85885d21f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'One deviation, recorded in P3 with its measurement and carried into P4 as an unmet criterion rather than re-read as satisfied. Criterion 9 said one sentence and the paragraph is two: written as one, the four claims came to about ninety words with three nested subordinate clauses and the #330 contrast last, after a reader had been asked to hold two others. The deviation is against my own criterion and not against the issue, whose "one sentence" draws a line against a section — section 9 treatment is what a candidate weighed in earnest gets, and #334 is to give AI-DLC the same — and a two-sentence paragraph is not a section, a table or a list, which is criterion 9 own second half. Nothing else departs: the placement, the Sources form and the decision not to edit internal/model/identity.go are each what the design named.'
      result: met
      rule: deviations-are-traceable
    - note: No interface change. Twenty-three lines of prose in one document, in two places in it. No command, flag, artifact field, template, gate or rule moves; no Go file is edited; internal/model/identity.go is in the scope to be read and was read and not touched. docs/process-definition.md is unchanged at 0 lines of diff, so nothing normative moves and there is nothing for anybody to migrate from.
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: 'No new dependency. go.mod and vendor/ are untouched, nothing is installed and no action is added. What is added is a citation: a repository named at a pinned commit in Sources, which this project neither fetches nor builds against. The entry is explicit that the repository has no licence file, so the sentence names a design and not a source of code, and nothing in the change makes any part of that tree a thing this pipeline reaches.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
    - note: 'Citation lens: this change puts a claim about somebody else repository into a document of this one, in the present tense, with nothing able to hold the two together. The Sources entry is dated and pinned, so it stays true as a record of a reading whatever happens upstream; the paragraph says the demo "starts each run from a label on an issue", which is a claim about a design that one author can change in an afternoon — the repository was pushed to on 2026-10-09, two days after the commit cited. Section 9 has the same exposure for OpenSpec and answers it the same way, by pinning and dating, so this is the established trade rather than a new one. What is genuinely new is only that this pin is a commit rather than a tag, because the repository publishes no releases, and a commit on a one-author demo is a weaker anchor than a published tag. Named rather than resolved: no test in this repository can check a claim about a tree it does not fetch, which is the first gap P4 records.'
      result: deviation
      source: lens
---

# Review

<!-- xeno:section:release-notes -->
## Release notes

**The orchestrator evaluation names the labels-as-triggers design, and says what it is.**
Section 4.1 chose OpenHands, and `rajistics-demo/sdlc-automation-github-demo` is that choice
running rather than an alternative to it: it drives the same Agent Server through the
OpenHands Automations API and starts each run from a label on an issue. A reader who met it
without the sentence could reasonably have taken it for a sixth candidate.

**It is the design the process definition declines, working.** Section 12 says issue
commands are deliberately absent, because something has to receive them and every way of
doing that is a component to build and operate. The demo has that component — a hosted
automation — and the paragraph points at the section 12 decision and says the demo is
evidence for it rather than against it. What a receiver costs is now visible in somebody
else's repository rather than argued about in this one.

**The contrast with #330 is drawn, because the two are easily confused.** A label gates work
in both places and the mechanisms are not the same. `xeno-approved` is a precondition
`xeno intent start` reads off an issue it has already fetched; nothing subscribes to it and
nothing is delivered. The sentence says that in the terms `Issue.Approval()` actually
supports.

**Sources carries the repository at a pinned commit.** `2ddf6c91791d95b8de8ff38b683241743ec12049`,
read 2026-10-10, with the reason the pin is a commit rather than a tag — the repository
publishes no releases — and with each path named by what it was read for. The reading's one
observation that is not in the issue is recorded there too: the tree holds no
`.github/workflows` directory at all, which is what makes "something has to receive them"
literal for that repository instead of a figure of speech.

**Nothing in section 12 moves, and nothing else in the evaluation moves.** A trigger label
would be a specification change before it was anything; #338 holds that question and now has
a use case from the maintainer. This intent writes an evaluation sentence and does not
anticipate how that is answered. The four status labels, the issue and pull request templates
and the `openspec/` directory are named in Sources as the rest of the design the sentence
does not take.

<!-- xeno:section:residual-risk -->
## Residual risk

**The paragraph's present tense can go stale and nothing will say so.** The Sources entry is
a dated record of a reading and stays true; the sentence describes a design, and one author
can change it. The repository was pushed to on 2026-10-09, two days after the commit cited.
No test here can check a claim about a tree this repository does not fetch, and
`internal/model/supply_chain_test.go` is the wrong instrument because this is a citation and
not a pin the pipeline reaches. What limits the damage is that a reader who finds the
repository changed can see exactly what was read and when.

**Six of the eleven criteria were judged by the agent that wrote the prose.** The pull
request is the second reading. Each of the six names the clause carrying it, so a reviewer
can disagree with a specific clause rather than with a verdict, which is what the naming is
for.

**Criterion 9 is not met and the record says so in three places** — the deviation, the
mapping table and the review answer. The claim it was written to protect is intact: the demo
does not get a section, which is what would have miscategorised it beside OpenSpec and the
AI-DLC section #334 is to write.

**The context budget is exceeded again, by 663 bytes.** 43,663 against 43,000. This intent
applied XENO-0286's learning and still overshot: a margin of 889 bytes was added for the
growth and the change was 1,552. So the method was right and the estimate was wrong, which
is a sharper finding than XENO-0286's — the growth has to be estimated from what will be
written rather than guessed at as "nine or so lines", and nine lines was in fact
twenty-three. Two intents have now carried this finding to their end, which is what makes it
worth a rule rather than a note.

**What the sentence deliberately does not settle.** Whether a trigger label is wanted here is
#338's question, and that issue now has a use case from the maintainer and an open question
about whether other labels or a title keyword should carry other triggers. This paragraph
records that the design exists and that section 12 declines it; it does not prejudge what
happens if section 12 changes, and if it does change this paragraph is one of the places that
would have to be re-read.

**The build is CI's answer.** `zensical` is installed by `docs.yml` and not here, so that the
page renders is checked on the pull request. The paragraph uses inline code, quotation marks
and one em dash, which is the smallest Markdown surface available.
