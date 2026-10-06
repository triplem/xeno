---
intent: github.com/triplem/xeno#260
phase: 02-design
created: "2026-10-06T09:15:12Z"
schema_version: "1.0"
runner_version: dev+081da51
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 38e8e8503515d95456b8f0cb2098b662a2a339e2198b2be69c39e7c6c40cbdcf
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
decisions:
    - id: D-1
      chosen: The audit installs the tree cycjimmy/semantic-release-action installs — a checkout of the action at the sha read out of release.yml, npm ci --omit=dev in it, then npm install semantic-release@<pin> on top — and the baseline records that tree's figures, 2 critical, 23 high, 2 moderate, 1 low.
      rationale: 'It is the only one of the three routes that makes the number true, and it closes both halves of #260 with one change. The audit reported 0 critical for a pipeline that runs 2, because it resolved a fresh tree where the release resolves the action''s committed lockfile; 531 of the release tree''s 532 packages come from that lockfile, so the shipped tree is already pinned under A28 and the tree that drifted was the audit''s own. Re-measuring the wrong tree keeps a number false in the optimistic direction, and narrowing the check would narrow a check that is measuring the wrong thing. The cost accepted is a baseline rising from 12 to 28 with two criticals in it, which A44 must now say are not this project''s to fix.'
      decided_by: Markus M. May
      proposed_by: claude-opus-5
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The audit performs the action's install rather than describing it.** Three steps, in the action's
order: check out `cycjimmy/semantic-release-action` at the pinned sha, `npm ci --omit=dev` in that
checkout, `npm install semantic-release@<pin>` on top. The alternative was to reconstruct the tree
from the action's `package.json` — a fourth tree, resolved by this repository's reading of somebody
else's manifest, and wrong the first time the action changes how it installs. Copying the steps is
the only shape where the audit's claim stays true without anybody re-reading the action.

**The sha is resolved out of `release.yml` by the step that already resolves the version.** One
step, two outputs, one failure mode. `sed` yields an empty string where it matches nothing and an
empty ref checks out a default branch, so both outputs carry `test -n` — the pattern #187 left,
extended rather than copied beside.

**`--omit=dev` stands in for the action's `--only=prod`.** The flag is deprecated in npm 12 and the
two select the same set. This is the one place the audit deliberately differs from the action, and
it is named in the job's comment so that a reader comparing the two finds a decision rather than a
discrepancy.

**The baseline keeps A44's shape and changes only its subject.** Still severity counts, still
failing on a rise and not on a non-zero, still no per-advisory record. What changes is which tree
the counts are of, and the comment now says so first rather than last, with the 531-of-532 figure
that explains why the numbers move on a sha bump and not on npm's schedule.

**The advisory names go in the comment's prose and not in the check.** The check cannot tell a
critical replacing a critical from the same critical persisting, which was true at twelve and is
true at twenty-eight. Recording the names gives a reader something to diff by hand without
pretending the job does it.

**A44 is amended in place rather than superseded.** The assumption's shape is unchanged and its
evidence is not: it was approved on ten high advisories under `node_modules/npm/node_modules/`, and
it now has to hold for `handlebars` and `tar` at critical. The row records the new measurement and
its date, because an assumption whose evidence moved under it has stopped being the one that was
approved, and the register is where that is visible rather than in a workflow comment.

**The claim is corrected in all three places at once.** `audit.yml`'s header, the baseline's comment
and `docs/supply-chain.md`'s prose each assert that the audit mirrors the release. Leaving any one
of them would leave the false sentence in the file a reader is most likely to reach for, and which
that is depends on the reader.

<!-- xeno:section:alternatives -->
## Alternatives

**Re-measure the current tree and leave the install alone.** The cheapest route and the one #260
proposed, already performed once in #261. It keeps a baseline of 12 with 0 critical for a pipeline
that runs 28 with 2. A44's objection is to a check that fires with nothing to do; A90's is to a
reader that cannot fail. This is worse than either, because the number is not merely uninformative,
it is wrong in the direction that makes somebody stop looking. Rejected on that, not on cost.

**Narrow what the check asserts** — fail only on critical and high, or only where a fix is
available. It was the third route in #260 and it answers a question nobody has: the fault is not
that the check is too broad, it is that the tree is the wrong one. On the corrected tree it would
fire immediately on `handlebars` and `tar` with no action available, which is the failure A44 was
written to prevent. Rejected as aimed at the wrong thing.

**Commit a `package-lock.json` here.** #260's own preferred route, and the measurement is what
removes it: the release tree is 531 of 532 packages verbatim from the action's lockfile, so it is
already pinned and already does not drift. A lockfile in this repository would pin a third tree,
installed by nothing, and the audit would still not be reading the release. Rejected because the
problem it solves is not present.

**Bump the action's sha.** The tempting one, because a newer lockfile might simply drop the two
criticals, and it is the only route that would change what runs rather than what is reported. It is
out of scope for the reason A28 gives: what is wanted is the set that demonstrably works, which
means a release that demonstrates it, and judging the bump needs a correct number to judge against.
Deferred rather than rejected, and this intent produces the number it needs.

**Reconstruct the action's tree from its `package.json` instead of checking the action out.** No
second checkout, no second pinned reference. It resolves ranges this repository would be
interpreting on the action's behalf, which is a fourth tree and a new way to be subtly wrong; and
it goes stale silently the first time the action changes its install. Rejected: the whole fault
being repaired is a described install standing in for a performed one.

**Audit both trees and report both.** Honest about the fact that the two differ, and it keeps the
historical series intact. It also doubles the install time, gives two baselines to raise
deliberately, and asks a reader to know which of two numbers answers their question. The tree
nothing installs has no reader once it is named as such. Rejected as a cost with no consumer.

**Drop the audit job and rely on trivy.** Trivy reads the shipped binary and this reads the
pipeline; `docs/supply-chain.md` separates them deliberately, and the twenty-eight advisories are
exactly the ones trivy cannot see. Rejected: it would make the blind spot total rather than
partial.

<!-- xeno:section:impact -->
## Impact

**`.github/workflows/audit.yml`.** The header's claim is replaced, the resolve step gains a second
output, and the install step becomes a checkout plus the action's two commands. No change to the
job's triggers, its permissions, the comparison script, the manifest it declares or the artifact it
uploads. The comparison reads the same baseline file and fails the same way.

**`.github/npm-audit-baseline.json`.** The counts go from 0/10/2/0 to 2/23/2/1 and the comment is
rewritten rather than appended to. It names the tree, the three steps that build it, the 531-of-532
figure, the advisory names at critical and high, and the date. The sentence carrying #260's open
half goes, because this closes it.

**`docs/supply-chain.md`.** Two corrections in the same paragraph — which tree the audit reads, and
that it is one package since #121 rather than three — and one in the table, where the
semantic-release row says 24.2.9 against a pin of 25.0.9. The trivy-and-audit boundary above it is
unchanged and is what makes the two criticals an audit finding rather than a release one.

**`docs/assumptions.md`.** A44's row records the measurement that moved under it. No new assumption:
the decision not to commit a lockfile is a finding of this intent recorded in its artifacts and in
#260, and A28 already carries the pinning it rests on.

**What gets louder.** The audit now reports two criticals on every run, every pull request and every
scheduled pass. That is the intended effect and it is also a standing number somebody has to not
become numb to — the same hazard A44 names from the other side. The mitigation is that the number
can only change on a sha bump or a version bump, both of which are commits, so an unexplained move
is now genuinely unexplained rather than ordinary.

**What stops happening.** A pull request touching three Markdown files can no longer fail this job
because npm published overnight. The counts are a function of two pinned identifiers in this
repository, so the job's failures become attributable to commits in it, which is what #260 asked for
and what the lockfile route was reaching after.

**What is now visible and was not.** That A28's pinning is the reason `handlebars@4.7.8` and the old
`npm` tree are still installed. The pin buys reproducibility and pays in advisory age, and until now
no artifact in this repository showed the price. That is a real input to the deferred bump and it
only exists because the audit stopped resolving a fresh tree.

**No Go code, no gate, no artifact schema.** Nothing in `cmd/`, `internal/` or `.xeno/plugin/` is
touched, so `gate verify`, the test suite and `gofmt` assert absence of accident rather than
presence of behaviour. Section 7's gate list is unchanged and the second standing rule is not
engaged.

**Adopters are unaffected.** The audit job is this repository's own CI and is not part of the
plugin, the templates or the rule set; nothing shipped by `xeno init --vendor` changes.
