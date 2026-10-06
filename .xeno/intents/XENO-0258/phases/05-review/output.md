---
intent: github.com/triplem/xeno#256
phase: 05-review
created: "2026-10-06T08:40:32Z"
schema_version: "1.0"
runner_version: dev+5163d1b
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 06f57d9fc1e8078b3baa47f96e7d48f601756d438fa8d7b9e8689ab028c80df4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: 'Two, both departing from this intent own earlier phases and both caught by reading rather than by a tool. Against P1 non-goals: the first draft listed test -f, grep -l and find alongside git check-ignore, where P1 had said the paragraph names the class and one instance because #256 holds the table — the list went, and P3 learning is that in a change this small a non-goal is the specification of the output and its only reader is whoever writes that output a phase later. Against the conventions on prose: the shortened paragraph was rewrapped by hand and came out with three lines at 89 columns, then reflowed with a wrapper over the whole paragraph, which is what replacing a paragraph whole means in practice. No criterion was falsified; criterion 6 is a judgement about length and P4 reports the figure rather than claiming the judgement was verified.'
      result: deviation
      rule: deviations-are-traceable
    - note: No interface changes. One paragraph in a file no gate, rule or Go file reads — it is prose sent to a model, which is the only mechanism a convention of this kind has. No code, no command, no field, no artifact shape, no rule set change, and git diff --stat names CLAUDE.md alone. Nothing outside this repository depends on it. What changes for a session is four sentences read before it investigates, and nothing can verify that it read them.
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: 'None added and no code changed; go.mod is untouched and the whole diff is Markdown. Nothing was weighed either: the failure this addresses is a tool answering the same way for absent and non-matching, and no dependency fixes that — what would is re-running the check, which is what the paragraph asks for and costs three seconds.'
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three standing rules. No normative document is touched: `CLAUDE.md` is the conventions file, #210
changed it and #259 added to it, and neither the process definition nor the plan says anything about
how a finding is arrived at. Nothing is invented — no field, gate, tool or rule, and the paragraph says
nothing checks it rather than implying something does. The branch carries one intent, the commit
references #256, and the issue carries `wp0`.

The acceptance criteria. Eleven met, one met by the commit this phase precedes.

The non-goals held, with one that did not and was caught. No mechanism, no row in
`docs/clause-readers.md`, no repair of XENO-0249, no second instance in the file, no general advice,
no reorganisation of `CLAUDE.md`. The one that failed first: the draft listed four tools where P1 had
said the paragraph names the class and one instance, and the list went when the draft was read against
the section. P3's deviations records it and its learning is the general form — in a change this small
a non-goal is the specification of the output, and its only reader is whoever writes that output a
phase later.

What a reviewer should weigh is the framing rather than the content. The generic rule — verify a
negative before trusting it — would be the first line in `CLAUDE.md` about nothing specific to this
project. What makes it belong is the sealing consequence, and if that reading is wrong then this
paragraph is misplaced rather than false, and #256's other branch was available: record the decision
not to write it down, with the reason. P2's alternatives argues both.

The second thing to weigh is length. Ten lines on 87, in a file whose last line asks it to stay short,
and the second such addition this session. P4's gaps says plainly that `CLAUDE.md` is now accumulating
rules nothing can check and that its own last line is the only thing arguing against the next one.

The instance was re-verified from the tree rather than recalled, which is what the paragraph asks of
anyone writing a finding: the entry at `.gitignore:10`, the commit `154bb09` that added it, and the
two sealed phases of XENO-0249 that carry the false claim.

<!-- xeno:section:release-notes -->
## Release notes

`CLAUDE.md` gains one convention: **a negative result is evidence only when the thing checked was
there to be found.**

A tool asked about something absent answers the way it answers about something that does not match.
`git check-ignore` exits 1 with no output for both. So recreate the thing and run the check again
before writing the finding down.

**Why this is in a file about this project rather than general advice.** A finding goes into an
artifact and is sealed with it, so a wrong one is permanent rather than corrected, with the correction
somewhere a reader of that phase will not be. That consequence is a property of this process, and it
is the whole argument for the paragraph's place.

The instance it names is this session's own. #201 recorded that a directory was not gitignored; the
entry was at `.gitignore:10` all along, added by #200, and `git check-ignore` had run moments after
the directory was deleted. The claim reached two sealed phases of XENO-0249 and a pull request before
anything re-ran it. All three facts were re-verified from the tree while writing, because a paragraph
about verifying a negative resting on a remembered one would be its own counter-example.

It closes with "Nothing checks this (#256)", the form the sequencing convention beside it uses. There
is no gate and none is proposed: a gate would have to know which of a command's two silences it
received, and A90's standing answer is that a reader which cannot fail is worse than none.

**What it does not do.** It does not repair XENO-0249, whose phases are sealed and carry the false
claim — the correction lives in #254, in a comment on PR #246, and in this intent. It does not name
the second instance, #212's intake reading the wrong field and concluding the opposite of the truth;
that is the same shape by a different mechanism and #256 keeps both with the table of tools that share
the ambiguity.

For anyone reading `CLAUDE.md`: it is now 97 lines and carries two conventions whose only reader is a
model and whose compliance nothing can see. Both arrived from this session's mistakes rather than from
the plan. The file's last line asks it to stay short, and that is the only thing arguing against the
next addition.

<!-- xeno:section:residual-risk -->
## Residual risk

Nothing checks it, and the paragraph's own last clause says so. A session that writes an unverified
negative into an artifact breaks the convention with no finding, no refusal and no verdict; the only
reader is a model reading four sentences before it investigates, and whether it did is invisible from
here. That is the ceiling for this kind of rule, not a shortfall in this one — but it means the
convention's effect is unmeasurable and its failure will look exactly like its absence.

`CLAUDE.md` is now 97 lines carrying two conventions of that kind, both from this session's own
mistakes. The direction matters more than either addition: a file sent with every request, asking to
stay short, accumulating rules nothing can check. Its last line is the only argument against the next
one, nothing enforces it, and I am the author of both additions — which is the least reliable position
from which to judge whether a third is warranted.

The paragraph teaches one example and the class is wider. A reader will generalise from
`git check-ignore` to paths, which is narrower than the failure: #212's instance was a script reading
the wrong field on an artifact and concluding the opposite of the truth, with no path involved. The
class is named in one clause and carried by a path instance, so the field case is likelier to recur
than the path case it was written for. #256 has both and nothing points a reader of `CLAUDE.md` at it.

Two of the three learnings this joins are still unread. #229's and #212's sit in their phases'
`learning.yaml` with section 10's route unused, which is the observation #256 made about its own
origin. One learning became a convention because somebody asked; the mechanism that was supposed to
route them did not fire for any of the three.

The instance could age into reading wrong. It is written in the past tense — "recorded ... when the
entry was at `.gitignore:10` all along" — which should survive the entry being removed, and if it ever
is, nothing will report that the example now describes a state rather than a corrected claim.

What is not a risk: the verdicts, confirmed intact at exit 0; anything executable, since no gate, rule
or Go file reads `CLAUDE.md`; and the sealed artifacts of XENO-0249, which this does not touch and
which #254's residual risk already records as carrying the false claim permanently.
