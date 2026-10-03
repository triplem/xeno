---
intent: github.com/triplem/xeno#186
phase: 00-intake
created: "2026-10-03T13:21:27Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 55df07b8aec14a356f3bf9328891f3dbdd5ad219268b3fb95eaddbb0021c4d81
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Six commits reached `main` in one morning while `audit` was red, from six pull requests that were
honestly green. Three defects, where the issue described one. `audit` triggered on `push` to `main`
and on `pull_request` filtered to three paths, so an ordinary pull request never ran it and the job
reported after the merge — a check that cannot run before a merge cannot gate one. Only one context,
`verify`, was required, so the other four bound nothing. And three of those four could not have been
required even deliberately: the job id is the check's name and `gitleaks.yml`, `semgrep.yml` and
`trivy.yml` all named their job `scan`, which names no particular workflow. The check that exists to
catch this reported met, because the adapter answers met when any status check is required and
Appendix A's `enforcement` block has no row for which. What was red is advisory drift found by the
daily schedule: 35 high against the 15 the baseline recorded, nothing in the repository having
changed. In scope: `audit` on every pull request, distinct names for the three scanners, the drift
read and answered by moving the pin and measuring rather than raising the number, the version read
out of the release's workflow instead of copied, node pinned in both, and the rule written in
`CONTRIBUTING.md` because it failed. Out of scope: the protection settings, which are a host setting
and cannot be applied before the renamed contexts exist on `main`; `required_pipeline` naming the
set, which is an Appendix A change and a person's commit; and a hook, which section 7 rules out
twice — advisory by construction, and no exclusive logic.
