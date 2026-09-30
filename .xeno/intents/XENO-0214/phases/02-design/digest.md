---
intent: github.com/triplem/xeno#147
phase: 02-design
created: "2026-09-30T21:15:20Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0a51653.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3dbf42d0e686b9199939fa59ecdf6b67957fe45955a44283466b94414bfa7802
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Three requests, each failing locally: the protected branch for allow_bypass, the project for
required_pipeline, the two approvals endpoints for the approvals rows. That is the structural difference from
the GitHub adapter, which answers everything from one request and one status code, and it is why no code is
shared — the interface hides it because both satisfy it identically. allow_bypass reads four things and names
the first it finds, unprotect first because it is the largest bypass, with GitLab's own words rather than
"administrators may bypass". A tier limit is not-available whatever status carries it: 401, 403, 402 and 404
are all plausible and #104 guesses none, so the adapter distinguishes only a readable answer from an
unreadable one and puts the status in the reason. A rejected token stays an error, because five not-availables
would hide a misconfiguration behind the tier.
