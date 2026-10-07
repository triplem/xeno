---
intent: github.com/triplem/xeno#273
phase: 05-review
created: "2026-10-07T10:00:43Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 30e86fd6121080f301a6280de1cfadcfde35c687003f9b65c6e1addad99ac085
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - note: The rule applies to 03-implementation and that phase recorded one deviation, which names what it departs from and what it does not. A decision in P2 was drafted on a false premise, that TemplateID's comment already cites A33; a grep for A33 across the Go files returns one line, in phaseNumber's comment about the --phase prefix, and TemplateID's comment restates A33's reason without naming the row. The decision itself is unchanged, no comment is touched, and the reason now says what is true; it departs from no criterion, because P1 set none about comments and its constraint that no Go file appear in the diff still holds. Recorded rather than quietly fixed because a decisions section is read as checked. Marked not-applicable rather than met because the rule is scoped to the implementation phase and this is the review answering about it.
      result: not-applicable
      rule: deviations-are-traceable
    - note: 'No interface changes, and the point of the intent is that there are none. Nothing is renamed: git status over internal/, cmd/ and .xeno/plugin/ is empty, the only modified file is docs/assumptions.md, and xeno gate verify reports 499 verdicts verified. No artifact gains or loses a field, no gate changes, no command changes its behaviour or its refusals, and a project override still goes to .xeno/config/templates/<id>/ as section 5 says. What an adopter meets is one longer cell in a register they may never open.'
      result: not-applicable
      rule: interface-change-needs-a-migration-note
    - note: None added. go.mod and go.sum are untouched, no Go file is in the diff, and nothing compiles differently. The diff is one cell of one row in docs/assumptions.md.
      result: not-applicable
      rule: new-dependency-needs-a-rationale
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**The three standing rules.** No normative document is touched; section 5's per-id resolution
sentence is cited as the thing one variant would contradict and is correct as it stands.
Nothing is invented. The change belongs to WP2 by subject, since the template engine is where
the id lives, and to intent XENO-0271 for issue #273.

**What the issue asked, and what it got.** #273 is a question — "Does this make sense?" — so
the deliverable is an answer. It got one with a measurement behind it, on the issue, and the
reason recorded where the convention is stated rather than only in a thread. The maintainer was
put three options and chose the row.

**What a reviewer should look at first.** That the answer is "no" to a question the maintainer
asked, and that the reason rests on a measurement the maintainer did not see taken. Both
variants were run on throwaway copies; what is in the evidence is the output, not the tree it
was run on, which is gone. A reviewer who doubts it has the commands and the code paths and
would have to repeat it. That is the same standing as every other measurement in this trail and
it is worth naming once, because an answer that declines a maintainer's suggestion carries more
weight when the thing behind it is reproducible rather than reported.

**The second thing: the row is long.** It is now the second longest state cell in the register
and it is prose in a table. The alternative was a shorter cell pointing at the issue, which is
the option the maintainer declined for a reason that applies to its length too: a reader finds
the assumption and not the thread. If it reads as a wall rather than an argument, that is a
real cost and P4's gaps say nobody but its author has judged it.

**The question no rule asks: should the harmless variant have been recommended?** It is
genuinely tidier. `TemplateID` collapses to the identity function, the directory listing sorts
in phase order, and nothing in the trail breaks. What stopped it is one sentence in section 5,
and a sentence can be reworded by the person who owns it. The recommendation was to leave
things alone because the ordering buys little and `template.yaml` already names its phase, but
somebody who values the tidiness more could reasonably reword section 5 first and then rename,
and the row is written so that path stays open rather than closed.

**And: is a decisions section the right place for a claim about the tree?** This intent wrote
one that was false — that `TemplateID`'s comment cites A33 — into P2 and corrected it in P3.
XENO-0269 did the same thing one document up, in specification wording that had been approved.
Two intents in one day, the same shape, both caught by running a command while writing the next
sentence. P3's learning record proposes that such a claim cite what was run; nothing enforces
it and the next one is as exposed.

**What is deliberately not here.** No rename, no guard matching `id:` against its directory, no
change to `goneBundle`, nothing about the all-digit scalar, no second row, no change to section
5, and no code comment. P1's non-goals carry a reason for each, and the three findings left
standing are named in P2's impact so that leaving them is visible.

<!-- xeno:section:release-notes -->
## Release notes

The shipped template directories keep their names. `.xeno/plugin/templates/intake/` does not
become `00-intake/`, `model.TemplateID` goes on stripping the ordering prefix, and nothing in
the plugin or the runner changes.

What changes is that A33 now says what the alternative would cost. The row has always said the
template id is the phase without its ordering prefix and that the prefix orders the phases and
says nothing a template needs. It now also says that the directory name and the template id are
separate things — `Resolved.Ref()` comes from `template.yaml`'s `id:` field and the directory is
only the path `Load` resolves — and what follows from that.

Renaming the directory alone is harmless to the trail and costs a normative sentence:
`TemplateID` would collapse to the identity function, resolution would key on the phase, and
section 5's "resolution is per template id: project beats plugin" would become false, with a
project override moving to `.xeno/config/templates/00-intake/`.

Renaming the `id:` with it, which is the consistent form and what the issue asks for, retires a
check. Every sealed artifact records an old ref and the shipped set would answer a new one, so
`goneBundle` is set and `strings_hash` stops being recomputed across every artifact in the trail
that carries a template ref. `xeno gate verify` reports all of them verified either way, which
is the part worth knowing: the one command that recomputes the whole trail would show a reader
that nothing had broken.

Both were measured on throwaway copies rather than reasoned about, and both are written out in
#273 with their outputs, so either stays available to anybody who decides the directory ordering
is worth the price.

<!-- xeno:section:residual-risk -->
## Residual risk

**The destructive variant is still silent and this intent does not change that.** Rename the
`id:` fields and every `strings_hash` in the trail stops being compared, while `gate verify`
reports every verdict verified. What stands between somebody and that is a paragraph in a
register. P1 refused the guard that would catch it — a check that each `template.yaml`'s `id:`
matches its directory — because it would forbid the harmless variant as well, which is a
decision about what a template id is. That decision is now the thing worth taking, and nobody
has been asked for it.

**`goneBundle` cannot tell a rename from a deletion.** It is the root of the above and the
branch is right for what it was written for: a bundle the repository no longer carries cannot be
hashed by anybody. Distinguishing the two needs somewhere to record that a ref was retired,
which is a new field or a new file, so it is a change to what the gate can know. Written on the
issue, out of scope here, and nothing schedules it.

**The all-digit hash scalar is still skipped.** 64 digits in a hash field is parsed as a number,
so neither the shape check nor the recompute sees it. Harmless for a real sha256 and a reliable
trap for the next person who fabricates a value to test a gate, which is what happened here
twice before the cause was found.

**A row is the weakest form of this answer.** It is findable from the convention it qualifies,
which is why it was chosen over leaving the measurement in a closed issue, and it is still prose
nobody has to read. The register's own argument is that a reader of the tree learns why it is
the way it is from these rows; that holds only for readers who open the file.

**Variant A's cost was read, not demonstrated.** That renaming the directories alone contradicts
section 5 follows from the sentence and from `Load`'s signature. Nobody ran a project override
under the renamed layout to watch resolution pick by phase. The reading is plain; the
measurement that exists covers safety and not this.

**The row is long, and only its author has judged it.** Six of ten criteria are a person reading
one table cell, now the second longest in the register. Whether it argues or just occupies space
is a judgement, and the alternative — a short cell pointing at the issue — is the option the
maintainer declined.

**A false claim about the tree reached a sealed decisions section before being caught.** It was
corrected in the next phase, and the same shape appeared in XENO-0269 earlier the same day, in
approved specification wording. Both were found by running a command while writing the next
sentence. P3's learning record proposes that a claim about the tree cite what was run; nothing
enforces it.
