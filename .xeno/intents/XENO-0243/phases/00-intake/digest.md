---
intent: github.com/triplem/xeno#208
phase: 00-intake
created: "2026-10-04T07:36:40Z"
schema_version: "1.0"
runner_version: dev+49f2794
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c84fb91310fc1166c95dfb7283857d31876d2b7e1496a127edb430a0f04d1cb3
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake reads #208, recomputes its figures against the tree, and replaces the one
sentence of it that no longer holds. What the issue gets right is the whole of its
diagnosis: section 4 makes declaring the only act the runner knows, the declaration is a
frontmatter block of `output.md` inside `artifacts_hash`, and no command writes it, so
`evidence attach` has nothing to bind and G-Evidence passes over an empty set. What it
gets wrong is the count. Fifteen of the forty-seven verification phases in the trail do
carry a declaration, all `test-report`/`go-test`, all attached, every attachment reading
`pipeline: local` and `result: success` — a word section 4 does not define and no gate
checks, because the shape check reads the frontmatter and `attached.yaml` is another
file. The newest is XENO-0210 on 2026-09-29 and nothing since declares anything. What
those fifteen point at is a hand curated transcript, which is the case section 4 gives
as its own reason for sourcing evidence from CI. Meanwhile three workflows publish a
manifest in exactly the shape the attach reads, on every pull request, and none has ever
been attached. Scope, after Q-1, Q-2 and Q-3: `xeno evidence declare` in section 4's own
verb, with `produced_by` and `format` added to the model because WP6 names them and
nothing has ever written them; four pending items declared in this intent's own P4 and
pulled in at P5's start; and a report, a manifest and an upload added to `xeno.yml` so
that a test report exists to declare. Out of scope: judging the fifteen attachments
against the closed set, which would turn fifteen sealed phases red through `gate verify`
and is filed as its own issue; inferring a declaration from a manifest, which section 4
forbids and D-3 leaves forbidden; and any gate for the two provenance fields. Read:
#208, section 4 whole, sections 5, 6, 7 and 9 in part, nine source files under
`internal/` and `cmd/`, and four workflows. Nothing outside the repository.
