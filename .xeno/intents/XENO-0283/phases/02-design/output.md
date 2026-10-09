---
intent: github.com/triplem/xeno#346
phase: 02-design
created: "2026-10-09T15:11:56Z"
schema_version: "1.0"
runner_version: dev+9590797
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a2ff9f917f79cf2c442457c46827e3458da43747b185ecf193d9c5f4e0311b27
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

**GitHub reads by name, GitLab by a bounded page, because that is what each host offers.**
`gh label list --search "$label"` returns the labels whose name or description matches,
sorted by best match, and the default limit of thirty applies to that set rather than
to the repository's; the only way it misses is thirty labels all mentioning
`xeno-approved`, which is not a shape a tracker has. `glab label list` has no search,
so its read takes `--per-page 100`, the host API's ceiling for one page, and the line
says so. The two reads differ because the two tools do; one flag spelled for both
would be a `--limit` on `gh` that is a guess at the repository's size, where
`--search` is not.

**The exact match stays.** `grep -qx "$label"` on the names the search returns, because
`--search` matches substrings and descriptions: a project with `xeno-approved-legacy`,
or a label whose description cites this one, would otherwise be reported as having it.
The description the script writes names the label, which makes this label its own
such case on the GitLab branch's JSON too, where the existing `"name":"$label"` match is
already exact and stays.

**The test adds two strings and nothing else.** The list the existing test iterates
gains `--search` and `--per-page`; a test that failed on `main`'s script and passes on
this one is what holds the shape, and a shape test is all a test without a host can
be. It does not assert the number after `--per-page`, since the bound is a judgement
the comment carries and not a contract.

**The proof is a run, recorded in P4.** The second run against this repository, with
thirty-two labels, is what the issue's done-when asks for and what XENO-0282 could not
do before the label existed; the output and exit code go into P4's results.

**No register row.** The reads are inside the clause's room and the choice between
search and limit is a tool fact rather than an assumption; the design paragraph above
is where a reader finds why.

<!-- xeno:section:alternatives -->
## Alternatives

**`--limit` on both, a large number.** One shape for both hosts, and a guess at how many labels a repository may carry on the one host that offers a better read. Declined for GitHub, taken for GitLab as `--per-page`, where there is nothing better.

**Create first and treat the refusal as success.** `gh label create` refuses an existing name, so the script could try and swallow that exit. Declined: the refusal text is the host's and changes; `--force` would update the colour and description of a label a project may have customised; and a script that writes before it reads is not the shape `github-settings.sh` set.

**Paging in a loop.** Correct for any count; a loop with a counter and a JSON parse in POSIX shell for a case no tracker using labels for packages reaches. Declined, with the bound on the line.

**A test that runs the script against a fake `gh` on `PATH`.** Would exercise the branch rather than the text. Declined for this intent: it is a harness the plugin's `bin/` has no precedent for, and the one script already tested by shape is `xeno-env.sh`; a finding for the plan if a third script arrives.

<!-- xeno:section:impact -->
## Impact

**For a project.** A second run of the script exits 0 where it exited 1; a first run is unchanged. A GitLab project past a hundred labels still meets the old fault, and the line says so.

**For this repository.** Nothing to do at the merge: the label exists and the script is not re-run by anything.

**For the trail.** XENO-0282's gap paragraph stands and this intake says what it missed; nothing sealed changes.

**For the plugin.** One shipped file changes, so the vendored tree and the digest a release compiles move with the next release, as any plugin change does.
