---
intent: github.com/triplem/xeno#205
phase: 02-design
created: "2026-10-05T16:27:41Z"
schema_version: "1.0"
runner_version: dev+97d07da.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 786279273cf1c46c69a6777c831b41a65405c774cb1e7fdccb6444211a948e4b
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

`internal/model` is the resolver's home. It is the only package all six callers already import,
and it imports nothing but the standard library, so `internal/cost` can reach it without gaining
a dependency on the runner. A new package would be a seventh import for one function.

`LocalDir(root)` and `LocalPath(root, parts...)`, with `LocalDataEnv` and `LocalDefault` as
exported constants. Two functions rather than one because every caller but `init` wants a file
inside the directory, and a caller that joined the parts itself would be writing `filepath.Join`
around a function whose whole job is to be joined to.

An absolute value is returned as it stands. The entry point exports
`${XENO_PLUGIN_DATA:=$root/.xeno/local}`, so the value in practice is absolute and joining it to
the root would nest one path inside another. `filepath.IsAbs` is the test, which makes a relative
value relative to the repository root rather than to the process's working directory — the root
is what every other path in the runner is relative to.

An empty value reads as unset. `os.Getenv` cannot distinguish them, and an exported-but-empty
variable is what a shell produces from `export XENO_PLUGIN_DATA=` or from the entry point's
default having failed; resolving that to the repository root would put the ledger and the run
marker at the top of the tree.

`Ledger` and `ReportPath` stop being path constants and become base names. They were
`.xeno/local/cost-ledger.yaml` and `.xeno/local/enforcement.yaml`, exported and referenced
outside their packages; the directory part belongs to the resolver now, so each keeps its file
name and the callers build the path. `cmd/xeno/main.go` prints the resolved path rather than the
constant, because a message naming a location the file is not in is worse than no message.

Nothing is created by the resolver. It is a pure function returning a path; the writers already
create what they need, and a resolver that made directories would do it on every read, including
the reads that only report.

The register gets a row rather than a comment, by #243's test: why `XENO_PLUGIN_DATA` is read
when `XENO_PLUGIN_ROOT` was removed is a fact every later reader of section 7 needs, and it
outlives this intent.

<!-- xeno:section:alternatives -->
## Alternatives

Removing the clause from section 7, as A89 removed the resolution order, was the maintainer's
other option and the cleanest end state: the specification would stop describing something
inert. It was not taken, and the reason is not that it is wrong but that it is the opposite
trade — the specification commit is a person's and must precede the code, so it costs a round
trip to delete a capability, where this costs one function to deliver one. It also forecloses
relocating local state, which is the only thing the variable could usefully buy.

Leaving it inert and closing #205 against A84 was the third option. Rejected because A84 records
the gap and the audit's complaint is about exactly that record: a reader cannot tell a variable
with no reader from one whose reader is a constant, and a row saying so does not change which it
is.

A shared constant without the variable was considered — one `localDir = ".xeno/local"` fixing the
duplication without reading the environment. It is half the value for most of the work, and it
would leave the export still promising something, which is the part of the defect that can
mislead a person rather than a reader.

Putting the resolver in `internal/runner` was considered and fails on `internal/cost`, which
would then depend on the runner for a path. The dependency runs the wrong way: `cost` is called
by the runner and by the hook.

Reading the variable once into the `Runner` struct was considered. It would cache the value at
construction, which is tidier for the runner and wrong for `cost`, whose functions take a root
and no runner; two sources for one path is what this intent removes.

Resolving relative to the process's working directory rather than to the repository root was
considered and rejected: every other path the runner handles is root-relative, and a value
meaning different things depending on where the command was invoked is a worse promise than no
promise.

<!-- xeno:section:impact -->
## Impact

One new function pair and six call sites. `internal/model/model.go` gains `LocalDataEnv`,
`LocalDefault`, `LocalDir` and `LocalPath`. `internal/runner/runner.go` resolves the marker and
`phase.env`; `internal/runner/enforcement.go` and `internal/cost/cost.go` keep a file name each
and lose a directory; `cmd/xeno/main.go` loses a duplicated literal and prints a resolved path.
`internal/runner/init.go` is unchanged, deliberately, with a comment.

Nothing changes for any existing repository. With the variable unset every path resolves to what
the literal produced, which criterion 7 asserts per path. The trail is untouched because nothing
under the directory is hashed, and criterion 8 is `gate verify` confirming that rather than the
argument for it.

What is gained is that the export means something. A person who sets `XENO_PLUGIN_DATA` moves
the run marker, `phase.env`, the ledger and the enforcement report, which is what section 7's
line says and what the entry point has been promising since it was written.

What is also gained, and is the smaller half, is that one path is defined in one place.
`cmd/xeno/main.go` and `internal/runner/runner.go` both wrote `.xeno/local/phase.env`; two
places agreeing by hand about a path neither owns is the kind of thing that drifts silently, and
it is already the second time this session that a duplicated definition has turned out to be the
real defect behind a stated one.

Two exported constants change meaning rather than disappearing. `cost.Ledger` and
`runner.ReportPath` were paths and become file names, so anything outside this repository reading
them as paths would break — nothing does, and the only caller outside their packages is the
command that prints one.

The honest limit is that a person can split their own state. The variable is read at each use,
so setting it between `phase start` and `phase finish` leaves the marker in one directory and the
lookup in another, and the second command reports the phase as not running. That is true of any
path made configurable, it is not detectable from inside a single command, and the register row
says so rather than leaving it to be discovered.
