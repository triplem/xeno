---
intent: github.com/triplem/xeno#186
phase: 01-requirements
created: "2026-10-03T13:22:11Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a67d39b98e6fb315eb6f265f3c78bdd21b44e158bdfcd26f0cc4d653af27b04b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`audit` runs on every pull request with no path filter, and each check reports under a name that
identifies it: `verify`, `audit`, `gitleaks`, `semgrep`, `trivy`, where three of them were `scan` and
could therefore not be required individually. The version is written once, in `release.yml`, and the
audit reads it out of that file and refuses where it cannot — the header's claim that it installs
what the release installs was a claim and becomes a check. The toolchain moves to
`semantic-release@25.0.9` in both workflows together, with node pinned to the same major in both
because 25 declares an engine and neither workflow said which node it ran under. The baseline is
measured against the new pin rather than raised to meet the old: 20 distinct advisories against 12,
with `sigstore`'s dropped certificate constraints, `pacote`, `picomatch`, `postcss-selector-parser`
and three of five `brace-expansion` gone, and `undici`'s three arriving. The counts are exactly what
the tool reports, with no margin, because a margin tolerates the next finding silently. And
`CONTRIBUTING.md` says that every check gates the merge rather than one of them, that a red one is
fixed before the next merge, and why — naming the six commits. Out of scope: the protection settings,
which cannot be applied before the renamed contexts exist on `main`, since a required context that
never reports blocks the very merge that would create it; `required_pipeline` naming the set, an
Appendix A change left as a finding; and a hook, which section 7 rules out twice.
