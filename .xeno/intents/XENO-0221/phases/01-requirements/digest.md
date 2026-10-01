---
intent: github.com/triplem/xeno#167
phase: 01-requirements
created: "2026-10-01T16:47:15Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15693cf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 227b859d9a650b2b25f777b7fd4f84096c17990de217a4b8efff2a1e08bb8a81
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
A project that declares nothing is unaffected, which is every project today and Appendix A's default. A
declared gate whose hash matches runs, its exit status is the verdict, and its findings get their ids
through the same `carryForward` as every other finding — indistinguishable in identity, distinguishable
in provenance. A gate whose hash does not match does not run and is red naming both hashes, which is
WP4's done-when in its own words. Four criteria are about a foreign command behaving badly: a path that
cannot be run, an answer that is not the agreed JSON, a non-zero exit with no findings, and a command
that does not finish — the third because `Status` refuses a check that fails without a finding, so a
project's tool must not be able to make a verdict unrepresentable. A decision on an external finding
does not survive the next run, which is the first time that rule has had anything to act on. Non-goals:
no sandbox, since the marking is the point and a half-measure would read as containment; no example
declaration, since that is a command somebody has to delete; no discovery; no project-defined predicate
type; no retry; no parallelism. Execution stays out of `internal/gates`, which reads, and whose comment
was rewritten once already for one `git log`.
