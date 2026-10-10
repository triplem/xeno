---
intent: github.com/triplem/xeno#334
phase: 05-review
created: "2026-10-10T16:03:49Z"
schema_version: "1.0"
runner_version: dev+5044a7a
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 50105e941fdae0c61a07eccba06dc82ae3d218650f8a5aa526f69ecc739dbb51
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
Review and merge of #334. One document, one new section, 334 added and 9 removed lines; the
three shipped review rules answered and two lenses recorded.

What shipped is the decision #323 reached, written where the decision for OpenSpec is.
Section 10 pins AI-DLC at `v2.11.0` and says which sha is the annotated tag object and which
is the commit, carries five rows where 9.2 has four, settles the one contradiction #334 named
by going to the code rather than to either document, decides not an extension with the reason
on the page, weighs the three shapes with what each costs, lists five conditions for
revisiting, and names four mechanisms without taking any of them.

The fifth row says this project is behind, which is the reason to have it. The admission
conflict check is named and declined, which is the reason to name a mechanism one does not
want. And the reviewed-source fingerprint is recorded as a gap on the maintainer's decision
of 2026-10-10, with the weaker alternative written into 10.5 rather than into the tree, so a
different answer later changes one bullet and one paragraph and nothing else.

Three rules answered, two of them not-applicable and honestly so: no interface moves and no
dependency is added, with `git diff main` over the Go tree empty. The traceability rule is
met and carries the two deviations P3 recorded and the one gap P4 did, which is the pointer
to #339 that 10.6 does not have.

Two lenses, both of them about the thing this change actually risks. The negative-results
lens records four probed absences with the control that was run beside each, and names the
fifth claim whose control is weaker — which the section handles by not making the claim. The
sealed-finding lens records that the reading was taken from a tarball of the exact pinned
commit rather than from `main`, and that every claim has an address in P4's table, fifty-odd
of them, so a later reader checks rather than trusts.

Eight residual risks. The section is already not a description of `main`; three of its
figures will be stale within days as three of #323's were; 10.6 is missing the reviewer-agent
paragraph; the hosted sample's own rows are unread by anybody; one negative claim has a
weaker control than the other four; nothing re-checks the width of this page's prose; #334
carries no work-package label and no issue holds that; and a green verdict here proves the
record was sealed and the gates were green, not that the reading was the right reading.

One learning, which is the one thing this intent wanted and did not have: the per-claim table
of addresses lives in a sealed phase, the page carries a Sources entry, and the next person to
revisit this section under 10.5's conditions has the entry and has to rebuild the table.

Three sections, three rules, two lenses and one learning. No open question anywhere in this
intent; the one decision is P0's D-1 and it is the maintainer's.
