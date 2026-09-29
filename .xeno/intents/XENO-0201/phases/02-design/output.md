---
intent: github.com/triplem/xeno#120
phase: 02-design
created: "2026-09-28T20:30:17Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6e72fed.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 5c93b8fbddef9da37397f458c7273664bb9df6ecd66fa7cf7c611fc18fb62208
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

**`model.Project` grows the `agent` block, and one accessor returns both fields.** The
struct modelled the language block alone, which is what made a present source look
absent. It now models what the runner reads, and `r.agent()` returns the tool and the
default model, empty where the file or the block is missing.

**`Finish` takes the summary and writes the digest before it judges.** Order matters and
it is the whole of the decision: the digest is inside `artifacts_hash`, so writing it
after the verdict would seal a hash over a file that was not there. Written first, the
phase is judged with its digest in place, and a phase finished with a summary is green
where the same phase finished without one is red on a missing file.

**An empty summary is the same as none.** `--summary` with a path reads the file,
`--summary` alone reads stdin, and neither writes an empty digest. A digest whose body
is nothing is worse than a missing one, because G-Schema would pass it.

**The frontmatter comes from `common`, `language`, `context_hash` and the two new
fields, in `frontmatterOrder`.** The same builder both artifacts use, so a reader
comparing a digest with its output.md sees one order. A digest carries nothing from the
rendered group: no template, no strings hash, no rules hash.

**`secrets_hash` is not written, and the digest writer says so in a comment.** Section 5
gives the runner two jobs, filtering and writing, and this does one. The comment names
the missing file rather than the missing feature, so whoever ships the filter knows
where the other half goes.

**`section set` writes `model` and `tool` from the same accessor.** One source, two
writers, and the fields appear in `output.md` exactly as the digest carries them.

**Nothing is refused that was accepted before.** A phase with a hand written digest
finishes as it always did, which is section 5's rule about the supported path rather
than a concession.

<!-- xeno:section:alternatives -->
## Alternatives

**A command of its own, `xeno digest set`.** Symmetrical with `section set` and it lets
a digest be written while the phase is still open, which is a digest of a phase that has
not happened. The working sequence puts the digest at `phase finish`, and the reason
holds: the earliest honest moment to summarise a phase is when it is complete. Rejected
on the spec and on the argument behind it.

**Write the digest after the gates run.** Tempting, because then the verdict could be
summarised too. It would seal `artifacts_hash` over a tree without the digest and then
add the file, so the hash a verdict recorded would never match the directory again.
Rejected as a hash that lies, which is the defect three intents this week have been
about.

**Require the summary.** It would close the hand written route in one line and make the
tool mandatory, which is exactly what section 5 refuses and what #120 cannot have
without a specification change. Rejected on authority, and it would also turn every
existing flow red.

**Compose the summary from the phase's own artifacts.** The runner has `output.md` and
could render a digest from it with no agent involved. Section 5 divides the labour
deliberately: the digest is a summary of the exchange, not of the artifact, and no
runner has the exchange. Rejected as inventing content.

**Write `secrets_hash` over an empty filter.** It would make the field present and every
gate green, and it would assert that a digest passed through a filter that does not
exist. Rejected as the worst available option, and named here because it is the one a
schema invites.

**Default `model` and `tool` to something plausible when the block is absent.**
`claude-code` is the only tool this repository has used. A35's whole argument is that a
plausible value in a field nobody produced is worse than an absent one, so this is
rejected by the row being amended.

<!-- xeno:section:impact -->
## Impact

`internal/model`: the `agent` block on `Project`.

`internal/runner`: `agent()`, a digest writer, `Finish` taking the summary, `SectionSet`
writing the two fields. The frontmatter builder is reused rather than copied, which is
what keeps one field order.

`cmd/xeno`: `--summary` on `phase finish`, with the stdin form `section set` already
has, and the usage line.

`ASSUMPTIONS.md`: A35 amended to three fields.

Tests: the digest's content and field order, the absent `secrets_hash`, the unchanged
behaviour without a summary, a phase green with a written digest, the two fields in
`output.md`, and a project file with no agent block.

What a reader gains: `phase finish --summary` produces a digest that passes the gates,
so the frontmatter of a digest stops being something a person types. What they still
cannot get: a filtered digest, a `tool_version`, or a `rules_hash`.

What this costs: `Finish` writes a file before it judges, so a failure between the write
and the verdict leaves a digest with no verdict. That state is already reachable and
already harmless, since the next finish rewrites the digest and judges again.

Sixty existing intents are untouched, and `gate verify` is the check on that.
