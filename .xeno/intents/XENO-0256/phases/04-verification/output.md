---
intent: github.com/triplem/xeno#260
phase: 04-verification
created: "2026-10-06T07:25:22Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 95cd72d5326b34e68abeb782e8b9ca404762db27b357d3fdfa02d75a7cac7a17
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - format: other
      job: go-test
      kind: test-report
      path: evidence/go-test.txt
      produced_by: go test ./...
      result: pass
      sha256: 31560457a210f970908b35fdfb868ac0f5753096a09fc6c3d5ace672b788efd2
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

No test is added and none could be. The change is four numbers, a date and a paragraph in a JSON
file that no Go code reads and that lies outside `artifacts_hash`. A test would compare this
commit's numbers against a copy of themselves.

What stands in for it is one reproduction, two checks and a reading:

| criterion | what answers it |
|---|---|
| 1, the four counts and the date | reading the file; `python3 -c json.load` printing the values back |
| 2, the prose describes the twelve | reading it, and that it names `postcss-selector-parser` as returned rather than gone |
| 3, bumping and fixing recorded | the prose, carrying the three figures the experiment produced |
| 4, why it drifts recorded | the prose, naming `npm init -y` and the absent lockfile |
| 5, `toolchain` matches the pin | `grep` for the version in `release.yml` against the field |
| 6, the commit message | the message itself, which criterion 6 makes part of the work |
| 7, the audit job passes | **the pull request, and nothing here.** The job builds its tree at run time; what is reproducible is its comparison, done below |
| 8, `Refs` not `Closes` | the message footer |
| 9, the trail | `./xeno gate verify` |
| 10, the suite | `go build`, `go test ./...`, `go vet ./...`, `gofmt -l` |
| 11, the JSON parses with the same five keys | `json.load` and the key list |

**The job's comparison was reproduced rather than assumed.** The failing run uploads its
`npm-audit.json` as an artifact; it was downloaded and the baseline's own rule — each count against
the recorded one, failing on a rise — was run over it with the new numbers. Nothing rises, so the
job passes. That is as close as this repository gets to verifying `audit.yml` without running it,
and it is a stronger statement than criterion 7 asked for.

What it still does not verify is the job: the install, the pin resolution and the report generation
all happen in CI, so a change there could fail in ways this cannot see. #201's verification phase
recorded the same limit about the release workflow, and the same sentence applies.

**The figures agree three ways**, which is why they went into the file without hedging: the CI
run's own output, a local install of `semantic-release@25.0.9`, and a local install of `@latest` all
report `high 10, moderate 2`. The last pair is also what established there is nothing to upgrade to.

The suite is run and proves nothing about this change, which this phase says rather than letting a
green run stand as evidence. No Go code reads this file.

<!-- xeno:section:results -->
## Results

Ten criteria met, one met on the pull request and nowhere else.

1. Met. `critical: 0, high: 10, moderate: 2, low: 0`, `measured: 2026-10-06`.

2. Met. The prose names the twelve and says `postcss-selector-parser` returned, where the old
   paragraph said it had gone with the upgrade.

3. Met. The prose carries all three figures: 25.0.9 is the latest, `@latest` gives an identical
   tree, `npm audit fix` leaves high 10 and moderate 2 unchanged with `semantic-release` still at
   25.0.9, and the five marked `fixAvailable` are constrained by their parents' ranges.

4. Met. The prose names `npm init -y`, `npm install semantic-release@<pinned>`, the absent
   lockfile, and says the numbers are a dated measurement rather than a bound.

5. Met. `toolchain` says `semantic-release@25.0.9, read from .github/workflows/release.yml, and
   also the latest published version`, and the version matches the pin.

6. Met by the commit, which criterion 6 makes part of the work.

7. **Met on the pull request, and the comparison is reproduced here.** The job builds its tree at
   run time so nothing local can run it; what was run is the baseline's own rule over the
   `npm-audit.json` the failing run uploaded, with the new numbers. Nothing rises, so the job
   passes. The install, the pin resolution and the report generation remain unverified locally.

8. Met. `Refs #260`.

9. Met. `./xeno gate verify` exits 0 with the verdicts that existed before this intent intact and
   the rest this intent's own judged phases.

10. Met. `go build` succeeds, `go vet ./...` is silent, `gofmt -l .` outside `vendor/` prints
    nothing, and `go test ./...` is 18 packages `ok` with 0 failures — and this time the suite was
    allowed to finish, its process confirmed gone and the file's hash taken twice before anything
    declared it. The declared hash and the file's hash are the same `31560457`.

11. Met. The JSON parses and carries the same five keys.

**The figures agree three ways**, which is why they are in the file without hedging: CI's own
report, a local install of the pinned version, and a local install of `@latest`. The second and
third are what established there is nothing to upgrade to, which is the question the old prose
could not answer.

<!-- xeno:section:gaps -->
## Gaps

The drift is untouched and will recur. This records a measurement with a date; the mechanism that
made it go stale — a tree resolved at install time against npm's registry — is exactly as it was,
so the next upstream publication can fail an unrelated pull request the same way. The only defence
in the repository is that the prose now says so. The lockfile is the route that would end it and it
is the open half of #260.

The job itself is still unverified locally, and that is a real limit rather than a formality. What
was reproduced is the comparison, over an artifact CI produced. The install, the pin resolution from
`release.yml`, and `npm audit`'s own report generation all happen in the job; a change to any of
them fails in a way nothing here would see until a pull request went red. #201 recorded the same
about the release workflow and the sentence has not improved since.

Nobody has decided whether this file should need a cadence. It has now taken two deliberate acts in
three days — #189 raised it, this re-measures it — and a file that needs re-measuring on somebody
else's publishing schedule is either a file with a refresh rule or the wrong instrument. Two data
points is not a trend, and it is enough that the next one should not arrive unremarked.

The ten high advisories are counted and not examined. None has a fix this repository can apply,
which is measured; whether any of the ten is actually reachable in the way the release uses
`semantic-release` is a different question that nothing here asks. A44's position is that the count
is the signal, and that position is unexamined rather than wrong.

Nothing records what the counts were between the two measurements. High fell from 30 to 10 at some
point in three days and the file will now say it fell, without saying when or in how many steps —
so a later reader cannot tell a steady improvement from a spike and a recovery. A log of the runs
exists in the Actions history and in the uploaded artifacts, for as long as those are retained.
