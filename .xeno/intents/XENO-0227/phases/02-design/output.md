---
intent: github.com/triplem/xeno#179
phase: 02-design
created: "2026-10-03T12:15:18Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+4d472b6.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: b1e3edaa171f608ccf4ff3637c40cc4a6486be32deed051931e2e8996a135d5a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**The command is `xeno intent start --for ISSUE [--intent KEY]`.** `--for` is the issue, and it is
the only required argument. `--intent` is optional, which is the one departure from the signature
the issue proposed, and the reason is in the acceptance criteria: the key is the next number of the
sequence the intents directory holds, so requiring it would leave one of the two guessed inputs
guessed.

**`needsKey` is false for it, alone among the commands that name an intent.** Everywhere else the
key says which intent is meant and there is nothing to derive it from; here the intent does not
exist yet. `intent status` and `gate verify` already take `--intent` optionally, so the dispatch
field is there and the command is an entry in the table rather than a branch in it.

**The qualified id is built in `internal/model`.** Section 3 is intent identity and the package
holds the layout and the formats of the process definition. It is `Tracker.Qualified(issue)`, a
method on the configuration block rather than a function over three strings, so the error can name
the field that is missing.

**One rule for `--for`, not two.** The argument is split at the last `#`; what is left of it is the
project where it is there and `tracker.project` where it is not, and the project is prefixed with
the host where its first segment carries no dot. So `179`, `triplem/xeno#179` and
`github.com/triplem/xeno#179` all arrive at the same id, and a repository with no tracker block has
a way through. Two rules — a number or the whole thing — would have made the middle unsayable.

**The host is `tracker.base_url`'s host with a leading `api.` removed (A78).** It is the only field
in the block that names a machine. The prefix is the one place an API address and an id host differ:
GitHub serves github.com's API at api.github.com, while GitHub Enterprise, gitlab.com and every
self-managed deployment of either serve it under the host itself.

**The key is read off the directory (A79).** The prefix the intent directories share and the number
after the highest, padded as that highest key is. Where there is no key to continue from, or where
the directory holds two prefixes, it refuses and asks for `--intent`.

**What is checked about the key is that it holds no space.** Section 3 says the key need not be a
number, so its shape is not this tool's business; a space is, because YAML reads " #" as the start
of a comment and `qualified()` exists because an id cut off by one is read back shorter without
anybody noticing.

**`Tracker` becomes a named type on `model.Project`,** and `EnforcementCheck`'s anonymous struct
reads that type instead of declaring the three fields again. Two declarations of one block is how a
field comes to be spelled differently in each.

**The refusal on an existing directory is checked after the id and the key are resolved,** so a
command that was going to fail on its argument says so about the argument.

<!-- xeno:section:alternatives -->
## Alternatives

**`--intent` required, as the issue wrote it.** Refused for the reason the issue itself gives about
the file's fields: the key is derivable. This repository's convention is "the next number of a
sequence of its own, padded to four digits", and that sequence is in `.xeno/intents/`. Requiring
the key would have left the command taking two inputs where one of them is already written down
seventy-nine times.

**A `tracker.host` or `tracker.key_prefix` field in `project.yaml`.** Refused by the second
standing rule: everything the artifacts carry is enumerated, and an addition is a specification
change first. Both values exist in the repository already — one in `base_url`, one in the directory
names — so a field for either would be a second place the answer lives, and the two would
eventually disagree.

**Mapping `tracker.adapter` to a host, `github` to `github.com`.** Refused because it is a table
that is right until the first self-managed deployment, which is the case `base_url` exists for.
`BranchRulesFor` already refuses to default an address for exactly this reason, in those words: a
tool that fills in an address where the configuration is silent is a tool with one host.

**Measuring the key from the count of intent directories.** Refused for the reason
`nextAssumptionID` gives about the register: counting entries hands a removed record's id to the
next one. The highest number, not the number of names.

**Padding to four digits as a constant.** Refused as this repository leaking into the runner. The
padding comes from the key the number was taken from, so a project padding to three keeps three and
this one keeps four, which is D-7's decision read rather than reimplemented.

**Writing `assumptions.yaml` alongside.** Refused: section 4 lists it, but nothing refuses for its
absence and not every intent here has one. An empty register written to satisfy a list is an
artifact with no reader.

**Asking the host whether the issue exists.** Refused. It would put a network call in the one path
that has to work offline, and it is the open half of WP12.

<!-- xeno:section:impact -->
## Impact

**`internal/model/identity.go`, new.** `Tracker` with the three fields of Appendix A's block,
`Qualified`, the host derivation, `NextKey` and `splitKey`. Pure functions over strings and a
configuration block, which is what lets the table tests assert the derivations without a
repository.

**`internal/model/model.go`.** `Tracker` is added to `Project`, where the two readers of the block
now find one shape.

**`internal/runner/runner.go`.** `IntentStart` beside `IntentClose`, and `nextKey`, which reads the
directory and hands the names to `model.NextKey`. A `projectConfig` constant for the path, which
the four readers of it in this file spelled out twice each.

**`internal/runner/enforcement.go`.** Its anonymous struct's tracker block becomes
`model.Tracker`.

**`internal/runner/next.go`.** The suggestion for an intent that does not exist names the command.
The command is in the sentence and not in `Command`, because the issue is the one thing nothing
there knows and a command offered with a placeholder in it would be typed as it stands.

**`cmd/xeno/main.go`.** The usage line, the `--for` flag, the table entry with `needsKey` false,
and `cmdIntentStart`, which prints the key it chose — the key is now an output and not only an
input.

**No gate changes and no rule changes.** `rules_hash` and the gate list are untouched, so no verdict
behind this intent is recomputed differently. G-Complete reads `intent.yaml` as before, because the
file has the same shape whoever wrote it.

**`README.md` and `ASSUMPTIONS.md`.** The command in the surface list, a paragraph on what it
derives, the WP7 row, and A78 and A79 for the two derivations.

**What is not touched.** The seventy-nine existing intents, section 5, Appendix A, and
`model.Intent`'s field order.
