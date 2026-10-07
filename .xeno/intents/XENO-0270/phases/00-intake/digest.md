---
intent: github.com/triplem/xeno#231
phase: 00-intake
created: "2026-10-07T09:38:08Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 441bd135d84cb648244b895bdd0ebea3a0641a8d0af7d03bef5a61eaa623d606
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`README.md` opens with a list of the runner's internal parts, then twenty-four command lines,
then nine paragraphs about field values in this repository's own artifacts, then a coverage
table. What a first reader needs is absent: what the tool is for, which is the first page of
the process definition; where the name comes from, which is Appendix C of the same document
and which the implementation plan goes out of its way to decline; and where to read further,
which has no index at all.

The command list is a second hand maintained copy of `usage`, and it has drifted: measured at
six commands in `usage` and absent from the README, none the other way, 27 against 21. The
repository already refuses this shape, in a test that reads the names out of `usage` because
"a third copy would be a third place to be wrong".

The maintainer asked on the issue for an index in `docs/` highlighting the process definition,
the most used commands named briefly with `xeno --help` and the reference linked, and the rest
moved. That last part corrects the issue, which said the existing material stays further down.

In scope: the README rewritten, `docs/README.md` as the index, `docs/commands.md` whose
command block is held against `usage` by a test, and
`docs/the-trail-in-this-repository.md` taking what moves. Out of scope: the normative
documents, generating the reference from the flag declarations, which is WP16's, the
documentation site, the skills' own short lists, and shortening `usage`.
