---
intent: github.com/triplem/xeno#273
phase: 01-requirements
created: "2026-10-07T09:54:13Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: e5640c8e075d25b1bc120f83b4e96b28e08c2b993a96884afd1b4aeed3337d40
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

1. **A33's assumption and reason are unchanged.** The row still says the template id is the
   phase without its ordering prefix, still gives the prefix ordering the phases and saying
   nothing a template needs as the reason, and still names the mapping table it was chosen
   over.

2. **Its state column carries what the first variant would cost**: renaming the directory alone
   makes `Load` resolve by phase, so resolution stops being per template id and section 5's
   sentence becomes false, and a project override moves to `.xeno/config/templates/00-intake/`.

3. **And what the second would cost**: renaming the `id:` with it sets `goneBundle` for every
   sealed artifact, so `strings_hash` stops being recomputed across the whole trail.

4. **The figure is in it.** 495 artifacts record a template ref, which is the number of
   verdicts the second variant retires the check over.

5. **It says that `gate verify` does not notice either variant.** That is the part worth the
   row: a reader who tried it would see 495 verified and conclude nothing had broken.

6. **It names the issue.** `#273`, so the two variants and their measured outputs are findable
   where they were written out.

7. **Nothing is renamed.** `.xeno/plugin/templates/` is untouched, `model.TemplateID` keeps
   stripping the prefix, and every `template.yaml` keeps its `id:`.

8. **No normative document changes.** Section 5 is cited as the clause one variant would
   contradict, not amended.

9. **The row stays one line with five cells**, as every row in that table is, and `approved`
   stays in the state: the amendment records what was learned, not a new question.

10. **`go test ./...`, `gofmt -l .`, `go vet ./...` and `xeno gate verify` pass**, and the
    verdict count is the trail's as it stands. Nothing in this intent can change behaviour,
    which is the claim the suite checks rather than demonstrates.

<!-- xeno:section:non-goals -->
## Non goals

**No rename, in either form.** That is the decision and not a thing postponed. Both variants
are written out on the issue with what each costs, so either stays available if the `ls`
ordering is later judged worth the price.

**No guard against the second variant.** A check that `template.yaml`'s `id:` matches its
directory name would make the destructive form impossible to do by halves. It would also
forbid the harmless one, and it is a decision about what a template id is rather than a guard
on a mistake nobody has made. Not asked for, and it would need its own argument.

**No change to `goneBundle`.** It cannot tell a bundle that is gone from one that was renamed,
and that is why the second variant is quiet. The branch exists because a bundle the repository
no longer carries cannot be hashed by anybody, which is right; distinguishing the two cases is
a change to what the gate can know and is written on the issue instead.

**Nothing about the all-digit hash scalar.** A `strings_hash` of 64 digits with no letter is
parsed as a number and skipped by both the shape check and the recompute. It is how the first
attempt at this measurement went wrong and it is on the issue. A real sha256 being all digits
is a one-in-ten-trillion accident, so it is a curiosity rather than a hole, and fixing it means
deciding what the gate should say about a field whose type is wrong — which is a finding with
its own shape and not a line in this row.

**No second row.** The facts belong to A33 and amending it is what keeps them findable from the
assumption they qualify. A new row would say "A33 was measured", which is a row about a row.

**No change to section 5.** It is correct. One of the variants would have made it false, which
is a statement about the variant.

**No re-measurement of the figure.** 495 artifacts recording a template ref is counted from the
tree as it is today and will be wrong tomorrow in the harmless direction, since the trail only
grows. The row says what it is a count of rather than pretending it is a constant.

<!-- xeno:section:constraints -->
## Constraints

**The first standing rule.** `docs/process-definition.md` is not the agent's to edit and is not
edited. `docs/assumptions.md` is the open register, declared so by its own first paragraph, and
A33 is a row in it.

**A row stays, with its state column.** The register says a row deleted for having been
superseded would be the rewriting this process refuses everywhere else, and that every row keeps
its state. So A33 is amended in its state cell and its assumption and reason are left alone.

**A negative result is evidence only when the thing checked was there to be found.** This is the
rule the measurement broke on its first attempt and the one it finally met: the corrupted value
has to be one the field's type accepts, and the control has to be run against the unmodified
tree. Both variants were measured with a hex value containing a letter, after an all-digit one
produced the same false green on the untouched tree. P0's `learning.yaml` carries the general
form.

**What is sealed is never rewritten.** No artifact and no verdict moves, which is also why the
measurement ran on throwaway copies: `gate run` rewrites the phase's `gate.yaml`, and timing or
corrupting a sealed phase in place would dirty a committed artifact.

**Prose wraps at 88 characters**, tables and code blocks do not. The register's rows are table
rows and keep their width.

**No invented fields, gates, tools or rules.** Nothing is added anywhere; one cell of one row
gains sentences.

**One dependency.** Untouched. No Go file is in the diff.

**Every change belongs to a work package and to an intent.** WP2 by subject, since the template
engine is where the id lives, and intent XENO-0271 for issue #273. The key was passed
explicitly, because two intents are open on other branches and the sequence counts from this
one.
