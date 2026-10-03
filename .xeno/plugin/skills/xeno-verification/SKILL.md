---
name: xeno-verification
description: Run P4, the verification phase of a Xeno intent — map each acceptance criterion to the check that proves it, record the results and the gaps. Use after 03-implementation is green, or when `xeno intent status` shows 04-verification as the next phase.
---

# Verification

The fifth phase. It maps the criteria to the checks, reports what the checks said, and
says what is still not covered. The phase that catches what a green suite would have
shipped.

## What the phase owes

Three sections, from the `verification` template:

- `test-mapping` — one row per acceptance criterion, naming the check that proves it. A
  criterion with no check is named as such rather than left out of the table.
- `results` — what the checks actually said, with figures. Run the thing against the
  real repository on a throwaway copy where that is possible: fixtures carry the current
  shape of everything, and the tree carries every earlier shape.
- `gaps` — what is not covered, what a check cannot say, and what was proved only for
  the case it was written for. This section is the honest half of the phase and the one
  a reader trusts.

## The commands

    xeno phase start   --intent KEY --phase 04
    xeno section set   test-mapping --intent KEY --phase 04 --file PATH
    xeno section set   results      --intent KEY --phase 04 --file PATH
    xeno section set   gaps         --intent KEY --phase 04 --file PATH
    xeno phase finish  --intent KEY --phase 04 --summary PATH

Where a verdict needs the commit range, pass both ends; the runner never works it out,
because a guessed range means different verdicts locally and in CI:

    xeno gate run --intent KEY --phase 04 --base REF --head REF

The one field the runner cannot know is `tool_version`, because section 7 forbids it
branching on the harness. G-Schema requires it, so the harness says it: section 7's
`XENO_HARNESS_VERSION` for a whole session, or `--tool-version` on the phase's first
`section set` for one run. The digest takes it from the artifact rather than asking
again, and where neither says it the field is absent and the gate reports it missing.

    export XENO_HARNESS_VERSION=2.1.276     # once, for the session
    --tool-version 2.1.276                  # or per run, over the top of it

## What the gate refuses

Everything the earlier gates refuse, plus:

- **G-Test**, from P4 — a declared test result that failed, or a mapping of the
  acceptance criteria that is incomplete. It reports `not-implemented` in this release,
  so a phase does not go red on it yet. What the phase owes is the mapping either way:
  the gate list is a declaration the runner grows into, and a mapping written only when
  something checks it is a mapping written to pass.

## What it hands on

The results and the gaps. P5 reviews against them, answers the review rules, and writes
the residual risk — which is largely this phase's gaps, read again by somebody deciding
whether to release.
