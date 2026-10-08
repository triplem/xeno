---
intent: github.com/triplem/xeno#94
phase: 02-design
created: "2026-10-08T20:15:48Z"
schema_version: "1.0"
runner_version: dev+30b1dea
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0f0aca2a2a02af89a00b250da7db6f9f900aca0a45ff80e175749d222ca0f5b7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The policy is one file and the threshold is zero.** `.golangci.yml` is the home of
what is linted and what is accepted, in the repository as a commit, which is the half of
`.semgrep/golang.yaml`'s own argument that transfers: facts want currency, policy wants
a commit. It enables exactly the four the analysis named, `govet`, `staticcheck`,
`errcheck`, `gosec`, with `default: none` so that the set is the one written and not
the tool's. The threshold is the tool's exit code with nothing excluded by severity or
confidence: measured, zero is reachable, and the maintainer asked for it where it was.
gosec's severity and confidence filters are left at their defaults rather than raised,
because a filter chosen to make the count zero is the threshold-before-the-result the
analysis warned against; the count is zero by the suppressions and the exclusion below.

**Four findings in shipped code are accepted where they stand.** Each is a design this
project has already taken: `git.go` launches `git` with the runner's own arguments,
which is what a commit range reader is; `gitlab.go` sends a request to the host
`project.yaml` names, which is what an adapter is, twice; `rules.go` reads a rule file
at a path the resolver computed, which is what a rule reader is. A `//nolint:gosec //
reason` on each line, naming the rule and the reason, in the position the finding is
reported at, so that a reader of the line sees the acceptance beside the construction.
Not a global exclusion of G204, G304 and G704: that would accept the next occurrence
unread, and the next occurrence is the one that matters.

**Tests are outside gosec, by one rule with its reason.** `exclusions.rules` with
`path: _test\.go` and `linters: [gosec]`, and a comment saying why: a test writes its
fixtures with `0o644` and `0o755` because that is what the code under test expects to
find, launches `sh`, `git` and `go run` because that is the thing being tested, and
carries Amazon's documented example access key because that is the secret filter's
own test input, which is also why its value is not written here: G-Secret reads this
artifact too, and reported it when the first draft of this paragraph quoted it. Twenty-three findings, every one a property of being a
test. errcheck and staticcheck stay on for tests: their eleven findings there are
unchecked writes and closes in fixtures and one expression, and each is fixed by
checking the error, which makes the fixture fail where it would have continued on a
tree it did not set up. The `std-error-handling` preset covers the rest of errcheck's
shipped-code findings, a `fmt.Fprintln` to standard error and a `resp.Body.Close` in a
`defer`, which no project checks and the preset exists for.

**One workflow per tool, in the scanners' shape, and the shape is cited not repeated.**
`lint.yml` and `govulncheck.yml` each have the eight steps `trivy.yml` has: checkout,
setup, make room, scan with `continue-on-error`, print, declare, upload, fail last. The
header of each says what the tool answers and points at `trivy.yml` for the reasoning
the three share, so the paragraph about the run worth reading lives once; `xeno.yml`'s
and `gitleaks.yml`'s comments that cited `semgrep.yml` for it now cite `trivy.yml`.
Job ids `lint` and `govulncheck`, which are the check names GitHub reports and the
strings the branch protection requires.

**golangci-lint arrives through its action, govulncheck through `go install`.** The
action is pinned by sha with `version: v2.14.0` beside it, which is the pair the trivy
row already has and the shape the test and renovate both read; it caches the analysis
and verifies the configuration against its schema, which a bare install would not.
govulncheck has no such action worth a row: `golang/govulncheck-action` is a thin wrapper
around the same `go install`, so the install line is written out, pinned to `v1.8.0`,
checksummed by the Go module sum database on the way in, and the pin is a row like
any other.

**The vulnerability report is judged by a step, not by the exit code.** `govulncheck
-format json` exits 0 whatever it finds, measured on a module requiring
`golang.org/x/text v0.3.0`: text mode exits 3 on the one reachable vulnerability, JSON
mode exits 0 and lists eight findings of which one has a trace reaching a function.
So the step reads the report and fails on a finding whose first trace entry names a
function, which is the text mode's own rule, and prints every finding with its OSV id,
its module, its fixed version and whether it is reachable. A reachable finding with no
fixed version is printed and does not fail the job, for the reason `trivy.yaml` gives
for `ignore-unfixed`. The judgement is in the workflow rather than in a configuration
file because the tool has no configuration file, and the step's comment says so.

**The database's date is already in the report.** govulncheck's JSON carries
`db_last_modified`, so the step prints it beside the result and no metadata file
travels along. That is the field trivy's report lacks and `trivy.yml` works around
with a second manifest item; here the report answers for its own coverage, and the pin
table's row for the database says it cannot be pinned, as trivy's does.

**govulncheck runs daily, lint does not run on a schedule.** A vulnerability arrives
without anything in the repository moving, so the scan runs on the cron trivy has, an
hour apart from it. A lint finding cannot arrive that way: the tool is pinned, the
policy is a commit, and a push to `main` runs it against the default branch. The weekly
schedule `semgrep.yml` had was for the same reason and the same reason is answered by
the push, so it is not carried over.

**semgrep goes in this change.** The comment's fourth step, remove it last once the
three report on the same tree, was for the case where nothing could be compared. The
comparison is in P0: semgrep reports nothing against this tree and has for its whole
life here, the new set reports 51 and is brought to zero in the same change, so a
period of overlap would compare nothing against zero. The required check moves with it
in `scripts/github-settings.sh`, and the script is run against the host before the
merge: a protection naming `semgrep` would wait forever for a check that no longer
exists, on this pull request first.

**The pin table moves with the pins, and the test holds it.** Two rows go, four come:
the action with sha and tag, the linter's version, govulncheck's version, and
govulncheck's database as a row that cannot be pinned. The test's `exempt` map loses
`semgrep's rules` and gains the database row with its reason; its `versionPin` shape
finds `v2.14.0` and `v1.8.0` in the workflows on its own. The page's paragraph that
contrasted semgrep's vendored rules with trivy's fetched database is replaced by one
about a policy file and a scanner's database, which is the same argument with the
names changed.

**Two register rows.** A104: the linter's policy fails on any finding, tests are
outside gosec, and four findings in shipped code are accepted inline with their reason.
A105: the maintainer's "trivy for the container scan" is read as `trivy.yml` unchanged,
scanning the built binary, with the image a separate question.

<!-- xeno:section:alternatives -->
## Alternatives

**depguard for A42, as the comment proposed.** Declined, and this is the one departure
from the approved plan. depguard checks the import lines of the files its glob matches;
`go list -deps ./internal/gates` lists the closure. A gate importing `internal/host`
would be caught by both; a gate importing a new package that imports `net/http` would
be caught by the step alone, and that is the case A42 exists for: "one import in one
gate and every verdict starts depending on whether a service answered". Moving the
rule would weaken it under the name of giving it one home. Cost of declining: the
argument for golangci-lint over staticcheck alone loses depguard and keeps gosec, which
is enough; and the rule's home stays a shell step, which A42 chose on purpose.

**golangci-lint by `go install` too.** One shape for both tools and one action sha
fewer on the table. Declined because the action is the publisher's own, verifies the
configuration against its schema, and caches the analysis across runs; `go install` of
golangci-lint builds it from source on every run, which the publisher documents as
unsupported for exactly the reproducibility reason this table exists for.

**`golang/govulncheck-action`.** One more sha on the table for a wrapper around the
line it would replace; its inputs are the Go version and the package pattern, both of
which `setup-go` and the line already say. Declined.

**Excluding G204, G304 and G704 globally.** Four findings become none without a
comment. Declined, because the next `exec.Command` with a variable is the one a reader
should meet with a finding and not with silence; the inline form accepts the four that
exist and no others.

**gosec on tests, with the fixtures changed to pass it.** `0o600` on fixtures the code
under test reads as `0o644`, subprocess calls restructured, the example key replaced.
Declined: each change makes a test say something false about the code it tests, and
the example key is the secret filter's own documented input.

**Keeping the weekly schedule for lint.** Carried over from semgrep for symmetry.
Declined above; a pinned tool and a committed policy cannot change between pushes.

**Running semgrep and the new set side by side for a while.** The comment's step four.
Declined above; the comparison it would make is already made and is nothing against
zero.

**SARIF to the host's code scanning.** The action can upload one. Declined: the report
is the file the manifest names, and the host's alerts would be a second record of the
same finding, outside the trail.

<!-- xeno:section:impact -->
## Impact

**For a pull request.** Two checks replace one: `lint` and `govulncheck` beside
`verify`, `audit`, `gitleaks` and `trivy`, all required. A finding from any of the four
linters fails `lint`; a reachable vulnerability with a fix fails `govulncheck`. The
branch protection on the host is changed before this merges, and `xeno enforcement
check` reads the enforcement block and not the context list, so it reports nothing
about the rename.

**For the tree.** `.golangci.yml` at the root; four comment lines in shipped code; about
a dozen test lines that check an error they ignored and one expression simplified;
`.github/workflows/semgrep.yml`, `.semgrep/` gone, `lint.yml` and `govulncheck.yml`
there. No package moves, no function changes, `go.mod` and `vendor/` are untouched.

**For the pin table.** Four rows in, two out, and the test that holds the page to the
tree passes both ways; renovate gains two managers and proposes the next golangci-lint
and govulncheck as it proposes the next trivy.

**For the trail.** Nothing sealed changes. The manifests the two workflows write have
the shape every scan manifest here has, so a later phase can declare `golangci-lint.json`
or `govulncheck.json` as `kind: scan` evidence the way the process allows and no
intent has yet done.

**For a reader.** `CONTRIBUTING.md` names the six checks; `docs/supply-chain.md` says
why the policy is a file and the database is not; the one argument that used to be
made about semgrep's rules is made about `.golangci.yml` in the same words, since it
was never about semgrep.

**For an adopter.** Nothing. What `xeno init` generates is unchanged, and the
maintainer's first answer is the reason: this is a project decision with no effect on
projects using Xeno.
