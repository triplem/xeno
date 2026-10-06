---
intent: github.com/triplem/xeno#260
phase: 01-requirements
created: "2026-10-06T09:13:57Z"
schema_version: "1.0"
runner_version: dev+081da51
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7d2cb0fe71da5e434491782a09d886ace597be79beb39799a490cd358ea97423
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

Numbered, because P4 maps tests onto these and a mapping that names a sentence cannot be checked
against the sentence it names. Nine of 55 P1 artifacts in this trail number their criteria and 46
do not; this is one of the nine deliberately, under #258's open third decision.

1. **`audit.yml` installs the tree the release installs.** A checkout of
   `cycjimmy/semantic-release-action` at the sha `release.yml` pins, then `npm ci --omit=dev` in
   that checkout, then `npm install semantic-release@<pin>` on top of it. Those are the action's
   own three steps in the action's own order, with `--omit=dev` standing in for the deprecated
   `--only=prod` the action passes.

2. **The sha is read out of `release.yml` and written nowhere else.** The step that resolves it
   fails when it finds no sha, in the same shape as the existing step that resolves the version.
   #187's rule is that two copies of one identifier is how an audit comes to examine a tree nothing
   ships; this change would otherwise introduce exactly that for the sha.

3. **The baseline records the figures of that tree, with its date and the tree it describes.**
   Critical 2, high 23, moderate 2, low 1. The comment says which install produced them, that 531
   of the 532 packages come from the action's committed lockfile, and therefore that the figures
   move when a sha is bumped rather than when npm publishes.

4. **The commit message says the raise is a correction and not a regression.** The baseline file's
   own rule is that raising a number is a deliberate act belonging in the commit that raises it. A
   reader who sees 12 become 28 with two criticals appearing has to be told in one sentence that
   nothing in the pipeline changed and the previous number was of the wrong tree.

5. **The audit still fails only on a count rising above the baseline.** A44's shape is unchanged:
   not on a count being non-zero, and not on a new severity class. The twenty-eight are no more
   this project's to fix than the twelve were, and the check that fires with no action available is
   what A44 exists to avoid.

6. **Every false claim about which tree is audited is corrected where it is made.** The header of
   `audit.yml`, the comment in `.github/npm-audit-baseline.json` and the prose of
   `docs/supply-chain.md` each say "exactly as the release installs it" or a variant; each is
   replaced rather than edited into, per the convention about a paragraph changed a second time.

7. **`docs/supply-chain.md`'s semantic-release row says 25.0.9.** It says 24.2.9 and the pin is
   25.0.9. It is found while correcting the claim beside it and is left wrong only by choosing to.

8. **A44 records what it now covers.** The assumption's reasoning was written about ten high
   advisories under `node_modules/npm/node_modules/`; it now has to hold for two criticals in the
   action's lockfile. The register says so in the row, with the measurement and the date, because
   an assumption whose evidence has changed underneath it is no longer the assumption that was
   approved.

9. **`gate verify` exits 0, `go test ./...` passes, `go vet` is clean and `gofmt -l` prints
   nothing.** No Go source is touched by this intent, so these assert that nothing was touched by
   accident rather than that something works.

10. **The audit job passes on this intent's own pull request.** The measurement was made locally
    under node 26.9.0 against CI's node 24, and 531 of 532 packages come from a lockfile that node
    cannot change, but the one floating install can. The figures are therefore confirmed by the run
    rather than by the local measurement, and a mismatch is a finding for P4 and not a surprise.

<!-- xeno:section:non-goals -->
## Non goals

**Not fixing an advisory.** All twenty-eight are outside this repository's reach. The sixteen the
audit has been missing sit in the action's committed lockfile, which only a bump of the action's sha
can move; the twelve it already saw were measured on 2026-10-06 as unfixable, with `npm audit fix`
changing neither count and `semantic-release@15.14.0` npm's only other suggestion. This intent
changes what is measured, not what is installed.

**Not bumping `cycjimmy/semantic-release-action`.** It is the obvious next question and it is a
different one. A newer sha carries a newer lockfile that may well drop `handlebars` and `tar`, and
it is also a change to the thing that cuts every release, which A28 judges by what demonstrably
works rather than by what is newest. It needs a release to demonstrate it and a number to be judged
against, and this intent produces the number.

**Not committing a lockfile here.** #260's second route, answered rather than deferred: the shipped
tree is already pinned at the action's sha, so a lockfile in this repository would pin a third tree
that nothing installs, and the audit would go on measuring something other than the release.

**Not narrowing what the check asserts.** Failing only on critical and high, or only on an advisory
with a fix this repository could apply, was the third route. It does not touch the fault — it
narrows a check that is measuring the wrong tree — and on the corrected tree it would fire at once
on two criticals with nothing to do about them, which is what A44 was written to prevent.

**Not watching staleness.** Nothing here says whether the pinned sha or the pinned version is still
the one to be on. #44 keeps that apart deliberately, and a correct count that does not move is no
more evidence of currency than a wrong one was.

**Not a gate.** Nothing in section 7 reads an audit baseline and nothing is proposed to. The job is
CI's and the second standing rule makes a new gate a specification change first.

**Not a change to any normative document.** The first standing rule binds
`docs/process-definition.md` and `docs/implementation-plan.md`. Neither says which tree a CI job
installs, so nothing precedes this and no section is amended by it.

**Not auditing the release's Go dependencies or the shipped binary.** Trivy reads what the release
ships and this job reads what produces it; `docs/supply-chain.md` draws that line and this intent
stays on one side of it.

<!-- xeno:section:constraints -->
## Constraints

**The install is somebody else's and has to be copied, not described.** The action's three steps are
`npm ci --only=prod` in its own directory, then `npm install semantic-release@<version> --no-audit`,
and the audit has to perform them in that order or it resolves a different tree again. `--only=prod`
is deprecated in npm 12 and `--omit=dev` is its replacement; the audit uses the replacement, which
is a deliberate divergence of one flag and is the only one. It is recorded here so that a later
reader comparing the job against the action finds the difference explained rather than apparently
accidental.

**The action's behaviour is pinned but not stable across shas.** `index.js` and
`installSpecifyingVersionSemantic.task.js` are read at `b12c8f6`, and a bump of the sha could change
either. So the audit's claim is only ever "the install the action at this sha performs", and the job
has to be read together with the sha it resolves rather than as a standing description. This is the
same limit `docs/supply-chain.md` already states about its own table, where the release is the record
and the table the convenience.

**Two identifiers now come out of `release.yml` instead of one.** The version already does; the sha
joins it. Both resolution steps must fail loudly on finding nothing, because a `sed` that matches no
line yields an empty string and an empty ref would check out a default branch — which is a moving
target and the precise failure A28 exists to prevent. The existing step's `test -n` is the pattern.

**Checking out a second repository costs a permission and a step.** The job has `contents: read`
over this repository; a public checkout of another needs no more than that, and `actions/checkout`
is already pinned here. Nothing else about the job's permissions changes.

**The baseline is severity counts and nothing finer.** It cannot record which advisories it covered,
so it cannot tell a critical replacing a critical from the same critical persisting. A44 accepted
that when the counts were ten and two; it is accepted again at twenty-eight, and the comment carries
the advisory names in prose so that a reader has something to compare against even though the check
does not.

**The figures were measured under a different node than CI runs.** Locally node 26.9.0 and npm
12.0.2; CI pins node 24, and the action declares `using: node24`. 531 of the 532 packages come from
the lockfile, which no node version can re-resolve, so only the single `npm install
semantic-release@25.0.9` is exposed to the difference — and that install added no path the lockfile
did not already hold. The figures are nonetheless to be confirmed by the first CI run, and
criterion 10 is where that is judged.

**The convention about changing a paragraph twice applies to three files here.** The audit header,
the baseline comment and the supply-chain prose have each been written and then amended once
already, so each is replaced as a paragraph and read back as prose rather than edited into.
