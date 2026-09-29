---
intent: github.com/triplem/xeno#141
phase: 05-review
created: "2026-09-29T19:57:59Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+6aa6f9b.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a573c831e7efda0b251b0aa27cf3b22fdccc1b84eaf74e0b325766e00f61f877
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

| item | state |
|---|---|
| The three standing rules | held. No document edited: `ASSUMPTIONS.md` is where `CLAUDE.md` sends a decision taken while building the core, not one of the normative two. No field, gate, tool or rule invented; `Requirement` and the `State` set both predate this. The change belongs to #141, labelled `wp12`, on branch `141-the-enforcement-vocabulary` |
| The row answers #104's fourth criterion exactly | yes, and no more than it. The decision is recorded, the port is not written, and the state column says which of the two it is |
| The rejected reading is recoverable from the row alone | yes. A reader with only the current tree gets both readings, the field that separates them and why the rejected one does not answer it, without opening #97 |
| The deviation is visible where a reviewer will look | yes, in `deviations` of P3 and again in the row's own account of what made it decidable. #97's schedule is not met to the letter and the row says on what ground |
| Eighty-eight columns | held in the prose of every phase. Not held in the table row, which cannot wrap; the file's existing rows run to 438 characters and P3's learning records that the exception is unstated in `CLAUDE.md` |
| The check suite | `go build`, `go vet ./...`, `go test ./...` all clean, `gofmt -l .` outside `vendor/` silent, `xeno gate verify` 145 verdicts and exit 0 |
| The diff is the diff that was promised | `ASSUMPTIONS.md | 2 ++` against `main`, outside this intent's own phase directories |
| Commit and reference convention | subject is Conventional Commits with no issue reference, `Closes #141` in the footer and in the merge request description |

**What a reviewer should push back on, if anything.** The substitution in the deviation is the
judgement call in this intent. #97 said the decision waits for GitLab to be written and it did
not wait; the argument is that #97's stated reason was wanting the second host's answer rather
than the adapter, and that the answer is served without a token. A reviewer who reads the
schedule as the thing to keep, rather than the reason behind it, should say so here, because the
alternative is real: draw the port in piece 1 against a provisional vocabulary and change
`Compare`'s tests twice, which #97 itself named as the counter-argument.

<!-- xeno:section:release-notes -->
## Release notes

No user visible change. Nothing in the binary, the artifacts or the commands is different, so
the release notes carry nothing for this intent beyond the commit subject.

For a reader of the repository rather than of the release: the open question #97 recorded is
closed. Where a host adapter arrives, it returns `[]enforcement.Requirement` and
`internal/enforcement` keeps the waive logic, the declared-or-not filter and section 13's
unchecked `merge_method` line. `Protection` and `ReviewsExpressible` are retired rather than
generalised, by the piece of #104 that writes the port.

<!-- xeno:section:residual-risk -->
## Residual risk

**The words drift, and this decision permits it.** Two adapters can phrase the same
requirement differently, where one table holds every phrasing today. The names are held
centrally and the words are not, so a report about a GitLab project and one about a GitHub
project can describe the same met requirement in language a reader would not connect. The row
records this as the cost and nothing mitigates it beyond review; if it turns out to matter,
the answer is a shared phrasing for the cases that are genuinely the same, which would be a
third reading and a new decision rather than a correction to this one.

**Two mapping rows rest on documentation.** `/approvals` answers 401 and the merge settings
are absent from an unauthenticated payload, so the spelling and shape of
`merge_requests_author_approval`, `only_allow_merge_if_pipeline_succeeds`, `merge_method` and
`squash_option` are unobserved. #104's third piece closes this against the Premium namespace.
The decision survives an error there because `allow_bypass` decides it alone, but the mapping
table would need correcting and #104 already marks those rows "to verify".

**Section 8's edition remains a claim this intent did not test.** That the Enterprise variant
has the approval rules the free edition lacks is read from the plan, and what was reachable
here was the tier that does not. An adapter written and verified only against Free would be
written against the weaker of the two, which is exactly what section 8 warns about; the
mitigation is that piece 3 is the piece that needs the fixture and it is sequenced last for
that reason.

**The vocabulary could still be wrong for a third host.** Jira is the case the plan says tests
the contract, because its keys have no dependency on a code host, and nothing here was chosen
with it in mind. Returning requirements is the more permissive of the two readings, so a third
host is likelier to fit it than the struct — but that is an argument, not evidence, and WP12's
own note accepts finding this out with a later host as long as the contract needs no redrawing.

**Not a risk.** That the row cannot be edited later. Appendix B putting the phase path inside
`artifacts_hash` is why a change of mind is a new intent, and that is the property, not a
limitation.
