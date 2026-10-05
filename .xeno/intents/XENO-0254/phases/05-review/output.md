---
intent: github.com/triplem/xeno#254
phase: 05-review
created: "2026-10-05T19:04:03Z"
schema_version: "1.0"
runner_version: dev+e471bbb.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 1ab725da20c3454e39db13fbbc43ee87badc5a7f2b694677c24b8883d375245b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'One, and it is a correction inside the phase rather than a departure from the plan: the new paragraph came out at 89 columns on two lines, one over the Markdown prose rule, and was replaced whole and read back rather than rewrapped at the break — which is what the conventions ask of a paragraph changed a second time, and which is how "naming what each is read by" became "naming what reads each". Nothing departed from P2 design: the two rows, their order, the second reader column wording, the correction paragraph and the count all landed as specified. Recorded beside it in P3: this intent is thirteen lines of one document behind about 1,400 lines of record, which is a finding about the plan rather than about the work.'
      result: deviation
      rule: deviations-are-traceable
    - note: No interface changes. Two rows of a document that nothing reads — no gate, rule, template or Go file opens docs/clause-readers.md, which A90 records as deliberate — plus a count and a paragraph. No code, no command, no field, no artifact shape, and git diff --stat names one file. Nothing outside this intent could depend on it, because the only dependant a document of this kind has is a person reading it, and what changes for them is that the second column is now true.
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: 'None added, and none could be: the change touches no code and go.mod is untouched. The one tool this would have benefited from is something that checks a row against the clause it describes, and P4 gaps records why there is none — a tool could verify that a named symbol exists and is called, which is not where this error was; the row named a real gate that does not read the clause. A90 finding about its own document applies: a reader that cannot fail in the interesting case is worse than none.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No normative document is touched: this corrects a description of what
reads section 8 and moves the description towards the code rather than the other way. Nothing
is invented — both rows are section 8's own words, no clause is added or reclassified, and the
only count that moves is the one row becoming two. The branch carries one intent, the commit
references #254, and the issue carries `wp5`, because #229 changed the reader and #212's
learning puts the row in the changing intent's package rather than the document's.

The acceptance criteria. Ten met, one met by the commit this phase precedes.

The non-goals held. No re-audit of the other thirty-six rows. No recommendation's reason, which
is #247. No sequence rule, which is #248. No code. No gate over the document. No change to the
four kinds or to the other three counts.

What a reviewer should check is one chain, and it is three greps: `QuestionShape` is called
from `phaseResult`, `phaseResult` from `schema`, and `schema` is the table's G-Schema entry from
phase 0. If that holds, the row's reader column is right and the old one was wrong; if it does
not, this intent has replaced one inaccuracy with another.

**Half of what prompted this needed no change, and that is the part worth reading.**
`internal/plugin/embedded/plugin` is gitignored at `.gitignore:10`, added by #200. #201's
finding that it was not came from running `git check-ignore` against a path moments after
deleting it, and git answers nothing for a path that is not there. The claim reached a
verification phase, a review phase, a pull request and two reports before anything tested it
with the directory present. XENO-0249's phases are sealed and carry it; the correction is in
#254, in a comment on PR #246, and in P0's learning here.

What this intent cost is itself a finding. Thirteen lines of one document behind seventeen
sections and about 1,400 lines of record — the clearest instance yet of the figure #117 has
collected eight times, and the case the plan's proportionality worry is actually about.

<!-- xeno:section:release-notes -->
## Release notes

The clause audit's one row about questions was wrong about what is checked and about what
checks it. It is now two rows.

| § | clause | reader |
|---|---|---|
| 6 | a question carries two to four options, each with its consequence, and one free entry | G-Schema's shape check, G-Questions |
| 8 | a question recommends exactly one of its options | `QuestionAsked`, the writer only; no gate reads it |

Before, one row read "a question carries two to four options and one free entry | G-Questions".
That was accurate when the pass was made on 2026-10-03 and #229 overtook it three days later by
giving the consequence a reader in the gate's shape check and the recommendation a reader in the
writer, without touching the row.

**The reader was wrong even for what the row did describe.** `QuestionShape` is reached through
`phaseResult`, so **G-Schema** calls it from P0, not G-Questions from P5. That is not a detail:
it is what #229 turned on, because a check added to the shape function reaches every artifact
ever written and would have made `gate verify` report XENO-3's sealed P0 as `DIVERGENT`.

**The second row records an asymmetry nothing else can.** The recommendation is enforced when
`xeno question record` writes a question and not when a gate reads one, so a green verdict says
nothing about it. That fact lived in two function comments and a call site; this document is the
one place in the repository whose purpose is holding it.

The tool-requirement count moves 37 to 38, and a paragraph above the table says these two are a
correction rather than a late addition — the document already explains three rows that arrived
after the pass, and a correction is a different claim about the pass's accuracy: it "was not
incomplete here; it was overtaken, and then wrong".

Nothing else changes. No code, no clause added or reclassified, no other row touched, and
`git diff --stat` names one file.

**A finding from #201 is withdrawn.** `internal/plugin/embedded/plugin` *is* gitignored, at
`.gitignore:10`, added by #200. The check that said otherwise ran `git check-ignore` against the
path moments after deleting the directory, and git answers nothing for a path that is not there.
There is nothing to fix; the correction is in #254 and in a comment on PR #246.

<!-- xeno:section:residual-risk -->
## Residual risk

Thirty-six rows are unexamined and this corrects one. It was found because #212 happened to be
adding rows beside it, not because anything looked; four intents since the pass have changed
what reads a clause, two updated their rows, one did not, and whether any other aged the same
way is unknown. The corrected count may now read as a warrant for the whole table, which it is
not — the paragraph above it says the pass was overtaken, and that is the only guard.

The convention that would prevent the next one is a learning and nothing more. #212 proposed
that an intent changing what reads a clause updates its row in the same commit; section 10
routes a learning through review, so until something adopts it the document's accuracy depends
on each intent remembering it exists, which is precisely what failed here and what will fail
again.

The second row describes an absence, and absences go wrong silently. "No gate reads it" holds
because `QuestionShape` does not look at `Recommended`; a later intent adding that check — which
#229 showed would turn a sealed verdict red — makes the row wrong in the same direction as the
one it replaces, by claiming less than is true. Nothing would notice.

Nothing can check this correction. A tool could confirm `QuestionAsked` exists and is called; it
could not confirm that the row describes the clause section 8 states, which is where the error
was. A90's finding about its own document holds, and this intent is an instance of it rather
than a fix for it.

The sealed artifacts of XENO-0249 carry a false finding and always will. The gitignore claim
runs through its P3, P4 and P5, and what is sealed is never rewritten. A reader of that intent
meets the claim with nothing beside it; only #254, this intent and a comment on PR #246 carry
the correction, and none of the three is where somebody reading XENO-0249 is looking.

What is not a risk: the 399 pre-existing verdicts, confirmed at exit 0, and anything executable.
No gate, rule, template or Go file reads this document, and the suite would pass identically if
both new rows were wrong — which P4 says rather than letting a green suite stand as evidence.
