---
intent: github.com/triplem/xeno#177
phase: 02-design
created: "2026-10-03T17:04:55Z"
schema_version: "1.0"
runner_version: dev+a18c1f3.dirty
plugin_version: 0.30.0
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 596bd4101e3eeddc62a7257123f25d602347f3cf6db75010a4d0290e5946cbf7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`internal/plugin` stops being a test-only package and holds the three things the tree can be asked:
where it is, what version it declares, and the hash over it. `Dir` is the vendored directory and the
resolution order is absent, with the measurement in the package comment so a reader finds the reason
rather than an omission. The tree hash is defined in the package that writes it, which Appendix B's
own words make the precedent for the two hashes it defers, and it descends where `hashing.DirHash`
does not — a phase directory's subdirectory is evidence and a plugin's subdirectories are the
plugin — cross-checked against a `sha256sum` pipeline as `artifacts_hash` was. `model.PluginVersion`
is removed rather than left unused, so there is nothing for the next ldflags line to set, and both
headers gain `omitempty` because absent is not a default. The fixture gains a vendored plugin, since
one without it was not a repository, and the absence is asserted in its own test rather than being
the condition of every other. `devVersion` is `dev`, because a literal cannot learn the tag and the
release is the only thing that knows the number. `XENO_HARNESS` beats `agent.tool`, since a project
declares what it uses and the variable is what is running. Eight alternatives refused, four of them
sharing one shape: the cheap version returns something indistinguishable from the right answer — a
hash over one file called the tree, a constant with the ldflags fixed, a declaration that cannot be
falsified, and an override whose failure is silence.
