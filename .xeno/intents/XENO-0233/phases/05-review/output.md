---
intent: github.com/triplem/xeno#177
phase: 05-review
created: "2026-10-03T17:07:16Z"
schema_version: "1.0"
runner_version: dev+a18c1f3.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b3618a008e8dbde7cf2d88d214c759931c131bf39bc1bf056d8c3e2506c41fb3
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
      Five, and the first is evidence rather than untidiness: this intent's P0 and P1 record
      plugin_version 0.29.2 and its later phases 0.30.0, because #196 merged and cut a release while
      the branch was open. Then a test replaced rather than amended, the upstream learning test
      adjusted, A82 skipped, and two clauses of #177 left open.
  - rule: interface-change-needs-a-migration-note
    result: met
    note: >-
      Two fields change what they say in every artifact written from here: runner_version changes
      shape and plugin_version changes source and value, and a repository with no vendored plugin
      now produces a G-Schema finding where it produced a plausible number. In the README, in A83
      and A85, and in the release notes.
  - rule: new-dependency-needs-a-rationale
    result: not-applicable
    note: >-
      Nothing was added; io/fs is the standard library.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

The three shipped review rules are answered in the frontmatter.

**`deviations-are-traceable` — met.** Five, each naming what it departs from, and the first is
evidence rather than untidiness: this intent's P0 and P1 record `plugin_version: 0.29.2` and its
later phases `0.30.0`, because #196 merged and cut a release while the branch was open. Then a test
replaced rather than amended, because its stated premise was the opposite of the decision; the
upstream learning test adjusted, because it compared against a constant that no longer exists; A82
skipped; and two clauses of #177 left open.

**`interface-change-needs-a-migration-note` — met, and it is the substance.** Two fields change what
they say in every artifact written from here. `runner_version` changes shape, `dev+<commit>` where it
read `0.1.0-dev+<commit>`, so the trail carries two forms with the commit as the boundary.
`plugin_version` changes source and value. A repository with no vendored plugin now produces a
G-Schema finding where it produced a plausible number, which is a behaviour change a reader will meet
as a red verdict — and is what section 5 requires. All of it is in the README, in A83 and A85, and in
the release notes.

**`new-dependency-needs-a-rationale` — not-applicable.** Nothing was added; `io/fs` is the standard
library.

**Beyond the three rules.**

*Was the decision the maintainer's to make actually put to them?* Yes, twice, and the second time
because the first answer was not enough. Three shapes for #183's scope with the measurements, and
then the two version literals separately when the fields were sourced correctly and still read
`0.1.0`. The second question was the useful one: sourcing a field from a wrong literal still reports
a wrong number, and nothing in the first exchange would have caught that.

*Does anything here make a verdict depend on the environment?* No, and that is the property the whole
scope was chosen to protect. `plugin.Version` and `plugin.Hash` read the vendored tree and no
variable; the lock records what they found; `XENO_HARNESS` reaches a frontmatter field and no
judgement. The check that nothing branches on it is in the required job.

*Is the trail safe?* `gate verify` is 275 at exit 0. Every sealed artifact keeps the header it was
written with, because a header lives in the file and verify recomputes verdicts. One sealed lock was
modified during this work by a mistyped `phase start`, caught by verify as a divergence and restored
from git — which is the contrast #193 was about: a modification is loud.

*What is the state of the two fields now?* Four header fields, four sources, none a constant. Before
this intent `plugin_version` was a constant the release overwrote with the runner's number, and
`runner_version` carried a literal twenty-nine minors stale. What remains is that neither is checked
against anything, because G-Supply does not exist.

*Could this change a verdict behind it?* No. No gate, no rule, no document, no existing artifact.

<!-- xeno:section:release-notes -->
## Release notes

**`plugin_version` names the plugin an artifact was rendered from**, read from
`.xeno/plugin/.claude-plugin/plugin.json`, and absent where no plugin is vendored. It was a constant
that the release overwrote with the runner's own version, so a released 0.28.0 runner in a project
whose plugin was 0.26.0 recorded 0.28.0 — a number that cannot disagree with the runner can never be
proved wrong, which is the one thing section 5 wants it for.

**A repository with no vendored plugin now produces a G-Schema finding** rather than a plausible
number. Section 5 requires the field in every process file; `xeno init --vendor` is what puts a
plugin there.

**`context.lock.yaml` carries `plugin: { version, sha256 }`**, which section 5 already enumerated.
The hash is defined to the byte in `internal/plugin`, following Appendix B's own practice for the two
hashes it defers to their writers.

**A development build's `runner_version` is `dev+<commit>[.dirty]`** and claims no version. It read
`0.1.0-dev` through twenty-nine minor releases, because a literal cannot learn the newest tag. A
release still sets the real number through ldflags.

**The plugin ships `bin/xeno-env.sh`**, the entry point section 7 describes, called by the plugin's
hook. It normalises `XENO_PLUGIN_DATA`, `XENO_HARNESS` and `XENO_HARNESS_VERSION` and reads the
client's own variables, which is the one place allowed to know which harness it is.

**`XENO_HARNESS` reaches section 5's `tool` field** and takes precedence over `agent.tool`. Nothing
branches on it, which the `verify` job checks.

**What this does not do.** `--plugin-root` and `XENO_PLUGIN_ROOT` are still read by nothing, and the
reason is in the entry point and in A84: `internal/gates` reads the vendored rules and templates, so
an override would make a verdict depend on the environment — and with G-Supply unimplemented the
failure would be silence rather than disagreement.

<!-- xeno:section:residual-risk -->
## Residual risk

**The manifest is already going to be stale.** It was bumped twice during this intent because a
release landed while the branch was open, and nothing keeps it in step with the next one. Until that
is decided, `plugin_version` is accurate about the manifest and the manifest is behind the project —
which is strictly better than a constant that was wrong about both, and is not right.

**`plugin_version` is now falsifiable and nothing falsifies it.** A project can vendor a plugin whose
manifest declares any number, and the lock records the hash that would catch it with no gate to make
the comparison. The arrangement is one step short of working, in the direction that leaves it
finishable.

**A repository without a plugin becomes red where it was green.** Correct by section 5 and a change
somebody will meet without expecting it. `xeno init --vendor` is the answer and the finding says the
field's name rather than that.

**The trail carries two `runner_version` shapes**, and two `plugin_version` sources, with commits as
the boundaries. Third time this has been true today; the answer is the same, that the hashes are
sealed and backfilling would describe writes that did not happen.

**The entry point is asserted in shape, not in behaviour.** Nothing in CI runs it, because running
it means a harness. Its client detection — `CLAUDE_PLUGIN_ROOT` or `CLAUDECODE` for claude-code,
`CODEX_HOME` for codex — is a guess about two clients' environments that no test here can check.

**The plugin hash is over the tree, so whitespace moves it.** Two vendored copies of one released
plugin that differ in a comment hash differently. Whether G-Supply compares strictly or reads the
manifest's version first is that gate's decision, and this writes the stricter value.

**What is not a risk.** Any verdict behind this intent, and any repository that vendors a plugin and
exports nothing: `plugin_version` comes from a file that was already there and `tool` from the
declaration that was already read.
