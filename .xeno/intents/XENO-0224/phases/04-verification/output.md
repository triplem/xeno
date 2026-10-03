---
intent: github.com/triplem/xeno#1
phase: 04-verification
created: "2026-10-03T10:23:18Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+ea0cb1c.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 54e6d1fd001031dc87470402a060fafc96003ce6e0db0253224a30a100d7fa3e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Every criterion here is a fact to check rather than a behaviour to test, so the mapping names the
command or the file that answers it.

| Criterion | Checked by |
|---|---|
| The required check, and whether administrators are included | `gh api repos/triplem/xeno/branches/main/protection` |
| The job that runs the verdict, and what it covers | `.github/workflows/xeno.yml`, and `./xeno gate verify` |
| How many intents ran all six phases | the phase directories, counted |
| How many gates report `not-implemented` against the plan's six | the gate table, and a phase's own verdict |
| The register is closed in its header and nothing is deleted | the file, read back; `git diff --stat` |
| Nothing is added to the register | this intent's own diff: one file, no row |
| #1 is closed with what was checked | the comment on #1 |
| What M0 triggers and this intent does not do is named | the closure comment, and the review's release notes |
| No document under `docs/` is touched | `git diff --name-only` |
| Nothing else moves | the suite, `gofmt`, `go vet`, `./xeno gate verify` |

There is no test in this intent and there should not be: nothing in it is code. The one thing that
would have deserved one — a check that the register stops growing — would be a test asserting that
a file does not change, which is what a reviewer is for.

<!-- xeno:section:results -->
## Results

**The host enforces what the closure claims.** `required_status_checks.contexts` is `["verify"]`
and `enforce_admins` is `true`. One required check, no bypass for administrators, on the default
branch. That is the setting the plan's first steps call the thing the whole tool rests on, and it is
read from the host rather than from a document.

**The check is the verdict.** `.github/workflows/xeno.yml`'s `verify` job runs `./xeno gate verify`,
which recomputes every verdict in the trail and writes nothing. On this branch it is at exit 0 over
**216 verdicts**.

**The trail is twenty-seven six-phase intents and fifty-one intakes.** Of seventy-nine intent
directories, 27 carry all six phases, 51 carry the intake alone — the shape `M0.md` and A20 explain
— and one carries four, which is this intent in flight. The plan's M0 paragraph says "M0 proves one
phase"; twenty-seven intents have proved six.

**Eleven of fourteen gates are implemented.** G-Supply, G-Secret and G-Test report
`not-implemented`. The plan's M0 paragraph expects six to be missing and names G-Questions among
them; G-Questions has been implemented since WP1.

**The agent layer answers section 6's reading.** Seven skills with the names section 13 fixes, both
manifests validated by the client, one hook, no MCP server — and a measured cost of about 606 tokens
on every session. WP11's done-when requires a phase to be carryable with commands alone, which is
what every one of the twenty-seven did.

**The register is closed and nothing was removed.** `git diff --stat` on `ASSUMPTIONS.md` is 19
insertions and 1 deletion, the deletion being the header sentence that moved into the past tense.
Seventy-seven rows, every state column, the whole table: unchanged. This intent adds no row, which
its own diff shows.

**No document under `docs/` is touched.** `git diff --name-only` is `ASSUMPTIONS.md` and this
intent's own artifacts.

**The suite is green**, `gofmt -l` outside `vendor/` lists nothing, `go vet ./...` is silent, and
`./xeno gate verify` is at exit 0.

**What the closure names as owed.** The plan's two dated paragraphs, which are a person's. The
`Xeno-Intent:` trailer, which is the next intent. And two decisions that arrived while this intent
ran, which the deviations record and the next intent carries.

<!-- xeno:section:gaps -->
## Gaps

**Nothing watches a milestone.** M0's conditions were met weeks before anything recorded it, and
what found the gap was a question about which work packages are open. M1 is now one clause and one
reading away and the same silence applies to it: no gate, no command and no artifact field has an
opinion about a milestone, so the next one will be noticed the same way, by somebody asking.

**The closure is a claim about a setting that can change.** `enforce_admins` and the required check
are the host's state today. Nothing in the trail records them, `xeno enforcement check` reads them
only when somebody runs it with a token, and a repository that quietly loses the required check
keeps a closed M0 and stops deserving it. The scheduled run WP10 mentions is what would catch that
and it is not set up here.

**"Xeno is developed through Xeno" is true for twenty-seven intents and not for fifty-one.** The
earlier ones ran the intake alone. The README says so and A20 explains it, and the sentence is still
one somebody will quote without the qualification.

**The register's closure is a convention, not a mechanism.** Nothing stops A78 from being added
tomorrow: no gate reads the file, and the only thing between it and another row is the paragraph at
the top and whoever reads it. Given that two decisions arrived within the hour looking for a home,
that is a weaker guard than it looks.

**Where a decision goes now is harder to find than where it went before.** A row in a table was
greppable and indexed by number; a paragraph in one intent's design phase is reachable only by
knowing which intent. This is the plan's switch rather than this intent's choice, and the first
thing it will cost is somebody re-deciding something already settled because they could not find it.

**The two readings of M0 are closed together and could have been closed apart.** The gate job was
met weeks ago and the walking skeleton this week. One closure records both and a reader who wants to
know when each became true has to read the dates in the evidence rather than the milestone.
