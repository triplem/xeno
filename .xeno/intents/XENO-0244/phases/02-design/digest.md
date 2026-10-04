---
intent: github.com/triplem/xeno#221
phase: 02-design
created: "2026-10-04T09:29:55Z"
schema_version: "1.0"
runner_version: dev+24becc3
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b33f0b251d52f99ef3f1a7ce158db071a54c72d467f914e1ebd75423b9e80983
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The design settles where the rule goes and how much of it there is. The check is in
G-Evidence rather than in G-Schema, because G-Schema judges the shape of one file's
frontmatter and an attachment is another file whose meaning depends on the declaration
`Collect` pairs it with; the finding is placed on `attached.yaml`, which is the sentence
that gate already carries about which file an attachment's failure belongs to; and the
wording is `EvidenceShape`'s, because one rule reported in two places has to read the
same in both. A value is judged and an absence is not, which keeps `other/trivy-db` —
the entry this repository's own vulnerability scan publishes with a file and no result —
green, and is the same asymmetry section 4 draws on the declaration side. In `Attach`
the entry is declined rather than errored, because that function is called from three
places and an error would abort a `phase start` carrying a predecessor's verdict
forward; the sentence goes into `Unbindable`, which is a field built for exactly this
and now has a third reason. The result check is lifted out of `unbindable`'s file-less
branch, since a bad result is wrong whichever way an entry is bound, and it is reported
before the binding reasons because it is the one a publisher can fix in the step that
wrote it. The correction of the fifteen is a hand edit, recorded as one: fifteen files,
one word, once, against a command that would be a permanent write path into sealed
phases for something that happens once. Eight alternatives are recorded against, the
closest being a shared function for both gates — rejected because the declaration check
carries an exemption for a pending item and the attachment check cannot, so the shared
thing is the set and it is already shared — and requiring a result on an attached
`test-report`, which is the same clause one level down and is filed rather than dropped.
The impact section names where the risk is paid: if the reading of Appendix B were
wrong, `gate verify` would report fifteen divergences in CI before the branch could
merge, which is why the figure is taken three times and the middle reading is the one
that proves it. Read in this phase: `evidence` and `Collect` in `gates.go` for where a
resolved item is read, `attach.go` around `unbindable` and its three callers, and the
trivy manifest again for the entry with no result.
