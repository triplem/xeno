---
intent: github.com/triplem/xeno#208
phase: 02-design
created: "2026-10-04T08:29:11Z"
schema_version: "1.0"
runner_version: dev+49f2794
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: cdf4164ab7f4a154712181361be4dfe7ddb99d8879ba480dcd046e2003257874
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The design settles the verb, the forms and the one property the issue actually asks for.
The command is `xeno evidence declare`, in section 4's own verb rather than the `record`
the four other writers share, because the section's method is keeping declaring apart
from running and from judging and `attach` is already the family it belongs to. One
command carries all three forms, which are the three states section 4 gives an item and
differ only in which fields are filled, and the refusals keep them apart: `--file` with
`--uri`, `--sha256` with `--file`, `--result` on a pending item, `--uri` without a hash
in the words `evidence.unbindable` already uses, a kind or a result or a format outside
the closed sets, a second declaration of a kind and job already declared, and a phase
with no artifact yet. There is no flag that supplies a hash for a file: the `--file`
form copies the report into `evidence/` and computes the `sha256` of what it copied,
which is the `CLAUSE-READERS.md` row the issue cites — a command that changes a hashed
file recomputes what it invalidates — satisfied for the second time in this tree after
`SectionSet` and `context_hash`. The copy and the declaration are one act so that
G-Schema's clause about an undeclared file in `evidence/` cannot be broken by using the
tool correctly. Shape is judged by the gate's own check, exported as `EvidenceShape` the
way `QuestionShape` was, and `format` joins the closed sets G-Schema compares against,
which is safe in exactly this commit and no later one, since nothing in the trail
carries a `format` yet. The writer amends one frontmatter field through `exchange.go`'s
two helpers and says nothing about `gate.yaml`, so a declaration after a verdict leaves
#216's state rather than a new rule. Ten alternatives are recorded against, the closest
being judging the attachment's `result` in the same commit, which would turn fifteen
sealed phases divergent through `Verify` and is filed as its own issue, and declaring
the test report from a local run, which would have saved the CI round trip by taking the
shortcut the fifteen already in the trail took. Read in this phase: `exchange.go` and
`runner.go` again for the two helpers and for `frontmatterOrder`, `gates.go` for the
three shape checks and `BuildKind`'s comment about a closed set nothing compared
against, `attach.go` for `unbindable`'s wording, `hashing.go` for `FileHash`, and
`semgrep.yml` for the publishing pattern `xeno.yml` is about to copy.
