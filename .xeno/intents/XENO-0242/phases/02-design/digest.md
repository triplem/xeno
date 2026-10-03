---
intent: github.com/triplem/xeno#188
phase: 02-design
created: "2026-10-03T21:20:53Z"
schema_version: "1.0"
runner_version: dev+efc44a9.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: bfd2d13a3817d1fad98faea469c685df66186e9d525a8341d580fea9c3061a5e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Two commands, each taking the channel its shape asks for: `question record` reads an
`open_questions` entry as YAML on stdin or `--file`, because up to five options with a
consequence each do not fit flags without an invented separator, and `decision record`
takes flags, because its entry is six scalars. The input is the artifact's own shape,
so the gate's checks are the only authority on what a question is; `questionShape`
and `decisionShape` become exported and the runner calls them before it writes, with
uniqueness widened from the file to the intent because `resolves` is read from any
later phase. The key and the id are assigned and not given, continuing the intent's own
highest rather than counting entries, which is `nextAssumptionID`'s reason. `--reason`
carries the rationale and `--by` the person, both words this surface already uses;
`--chosen`, `--proposed-by` and `--withdraw` are new, the last being section 8's third
exit. A new unexported `amendFront` replaces one frontmatter key and writes the body
back byte for byte, rather than going through `SectionSet`, because re-rendering the
body is the template's business; it refuses a phase whose `output.md` does not exist
yet and names `section set`. Neither command looks at `gate.yaml`: a write after a
verdict invalidates it as a section write does, which is #216's state. The figure
is two counts filled on the walk `summarise` already does and one `AskedNothing`
predicate, true where the intent reached 05-review, is not abandoned, and both
counts are zero; the listing marks the row and the one-intent form prints a line,
neither changing a column width. The alternatives are recorded with their reasons,
the generic frontmatter writer being the one worth stating: a writer that serves
three blocks can refuse nothing, and the refusals are what each block means. No
document changes, no gate is added, no existing artifact moves, and this intent's
own P0 and P1 keep the hand written blocks the commands could not yet write.
