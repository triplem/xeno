---
intent: github.com/triplem/xeno#145
phase: 02-design
created: "2026-09-30T05:45:26Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4e41c6b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c9c2f1629f6b4237b0927638ade8126dea77f9a250d76efe2d642794ec9bd89e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: by-hand
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**`WrapperHost` gains the two tracker values, so the table is the one place a host is described.**

```go
type WrapperHost struct {
	ID      string
	Path    string
	BaseRef string
	HeadRef string
	Adapter string // what project.yaml's tracker.adapter is written with
	APIBase string // and its base_url
}
```

The alternative was a second table, or the two values reaching the project scaffold from somewhere
else. One row per host is what #98's design says, and the reason it is right here is that the
wrapper's range expressions and the tracker's address are the same fact about a host arriving at two
files: splitting them is how they drift.

**`scaffold.Project` gains `Adapter` and `APIBase`,** and the scaffold's two literals become
`{{.Adapter}}` and `{{.APIBase}}`. Since `Project` is already the shape of what `project.yaml` needs,
this is the same change as adding the host to it.

**`--host` loses its default in `cmd/xeno` and `Init` stops filling it in.** Both did it, which is why
the flag's default was invisible: `init.go` has `if host == "" { host = "github" }` as well. The
refusal comes from `Wrapper`, which already refuses an unknown host with the list, so an empty host
reaches the same sentence rather than a second one written for the occasion.

**`WrapperHosts()` derives from the table, sorted.** It returned a literal `[]string{"github"}`,
which is the duplicate #98's design was meant to avoid and which nobody noticed because it was
correct.

**The GitLab wrapper is `ci-gitlab.yml`, scoped to merge request pipelines.** `rules:` with
`$CI_PIPELINE_SOURCE == "merge_request_event"`, because that is where the base SHA exists. The body
is GitHub's wrapper's body in GitLab's syntax, with the same two steps in the same order and the
same comments about why both ends of the range are passed, since those reasons are the host's
neither time.

**The squash template goes in `Manual`, not in the wrapper.** It is somebody's act on the host, and
the list exists for exactly that. It is worded per host, because GitHub's setting and GitLab's are
different settings with the same consequence.

<!-- xeno:section:alternatives -->
## Alternatives

**A `tracker` block that names no host, with the adapter selected by a flag at check time.** Rejected.
It would remove the need for the scaffold to know the host at all, and it is the wrong direction:
section 12 makes the tracker part of the repository's configuration, and #143's whole point is that
the host is a recorded decision rather than an invocation's argument.

**Deriving the API address from the adapter name inside `internal/host`.** Rejected, and this is the
closer call. `internal/host/github` could export its canonical address and `BranchRulesFor` could
fall back to it, which would make the scaffold's `base_url` a convenience rather than a requirement.
That is a host default in code, which is what #104's first done-when is about and what #143 removed;
the reason it stays removed is that a self managed deployment has no canonical address, so a default
is right for the hosted case and silently wrong for the other. The scaffold writing an address a
project can change is the same value in a place where changing it is expected.

**Keeping `--host`'s default and reading it as a documented convenience.** Rejected on #104's
criterion, and on what the default actually does: it decides what an adopter who does not say gets,
which is the one case where the tool should ask rather than assume. It also made the flag's default
value the third place the string `github` appeared for the same reason.

**Generating a GitLab wrapper that runs on every pipeline and computes the base itself,** from
`git merge-base` against the target branch. Rejected. It would work where the merge request pipeline
is not configured, and it puts range arithmetic into the wrapper, which the plan's own reason for
passing both ends forbids: a CI system may check out a merge commit it produced, and a wrapper that
derived the range would be the component making that mistake.

**Adding a `--tracker-adapter` flag so the two can differ.** Rejected as a feature nobody asked for.
The wrapper's host and the tracker's host are the same host in every case anybody has, and where they
genuinely differ the generated file is a starting point a project edits.

<!-- xeno:section:impact -->
## Impact

**`xeno init` gains a required flag.** Anyone scripting `xeno init` without `--host` now gets a
refusal. That is criterion 4 and it is a deliberate break: the population is scripts written against a
version where the flag defaulted, and the refusal names the flag and lists the hosts.

**The generated `project.yaml` changes for `--host github` in the two tracker values only,** and the
strings are identical. Criterion 6 checks that by comparison rather than by reading, because a
template edit can move whitespace.

**A GitLab repository can be initialised and cannot yet be checked.** `tracker.adapter: gitlab` with
no adapter behind it means `xeno enforcement check` refuses and lists what there is. This is the
state P1's non-goals name, and it is the piece of #143's work that makes shipping this one safe:
before it, the same configuration would have asked GitHub about a GitLab project.

**#98's design is tested.** A second host is one scaffold file and one table row, and if that holds
the claim is no longer just a claim. It is also the first evidence about #143's port question, from a
different direction: the wrapper table taking two more fields per host without a second table is the
same kind of boundary holding.

**`internal/scaffold` and `internal/runner` move together.** `Project` gains two fields and `Init`
passes them from the table, so the host is resolved once in `Init` and used twice. Nothing else calls
`RenderProject`.

**The manual list grows to five items.** It is printed by `xeno init` and read by whoever adopts the
tool, and the squash template is the item most likely to be skipped and most expensive to discover
later, since the symptom is a footer that vanished from history.

**Nothing in the gate path, the artifacts or the hashes.** No gate, no template, no field, no
dependency.
