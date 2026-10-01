---
intent: github.com/triplem/xeno#167
phase: 02-design
created: "2026-10-01T16:48:23Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15693cf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 29387d7f9ab77f267b52216c43fa30c81256709e45ac9ac2916ae37ce72fa43b
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

**`internal/external` runs the commands and `internal/gates` keeps the rules.** One function,
`Run(root, declared, phase, now)`, returning a `model.Check` per declared gate that applies.
`internal/gates` already holds everything an external check must satisfy and gains nothing here.
The import runs from the runner to both, so the package that reads cannot start a process even
by accident.

**The runner injects the producer rather than the gates package calling it.** `gates.Ctx` gains
an `External func(phase string) []model.Check`, which `compute` supplies and which `Run` calls
after the table. So the external checks pass through `carryForward` with every other check, which
is what gives them their ids and drops their decisions, and the verdict has one assembly point
rather than two.

**The hash is read and compared before anything is started, every time.** `sha256` over the file
as it is on disk, against the declared value, with no cache and no memory of a previous run.
Appendix A says "checked before every run" and that is a frequency rather than an optimisation
hint.

**A mismatch is a check that failed, not a run that was skipped.** `fail`, one finding, naming the
declared hash and the one found, and the next step saying to re-declare the hash if the change was
intended. The alternative — reporting `not-implemented`, or leaving the check out — would let a
modified tool turn a red verdict green by being modified, which is the opposite of what the clause
is for.

**The contract: an object in, an object out.** On stdin, `{"intent", "phase", "phase_dir",
"artifacts_hash", "root"}` — what the command needs to find the artifacts and name them back.
On stdout, `{"findings": [{"file", "cause", "next"}]}`, and nothing else is read. `cause` is
required of a finding, `file` and `next` are optional, and anything else in the object is
ignored so that a tool may answer a superset. Written in the register because neither document
specifies it and a project's tool has to agree with it.

**The exit status decides and the findings do not.** Zero is a pass, anything else is a fail. A
command that exits zero with findings produces a pass carrying them, which is how a tool reports
something it does not consider fatal; a command that exits non-zero with none gets a synthesised
finding naming the status, because `Status` refuses a fail with no finding and a foreign tool must
not be able to make a verdict unrepresentable.

**Sixty seconds, stated twice.** The limit is in the register and in the configuration comment,
the process is killed when it passes, and the finding names the limit. Nothing in the documents
bounds a foreign command; the first one somebody writes will have a bug in it, and an unbounded
gate is not deterministic in the only sense this project can claim.

**A foreign answer is text and nothing more.** The cause and the next step are trimmed to a
length a verdict can carry, the file is recorded as the command wrote it and not resolved against
the repository, and nothing read from the answer reaches a path, a command or a template.

**A declaration is validated before it is used.** An empty id, an empty path, an empty hash, a
hash that is not sixty-four hex characters, or a phase that is not one of the six: each is a
finding against the declaration rather than an attempt to run it. A project's configuration
mistake should read as a configuration mistake.

<!-- xeno:section:alternatives -->
## Alternatives

**Let `internal/gates` run the commands.** One package for everything about gates, no injected
function, no new import edge. Rejected because that package's claim is that it reads: it was
narrowed once in #162 to admit one `git log`, and admitting arbitrary foreign code would leave the
comment saying nothing. The injected function also has a second benefit the tests use — an external
producer can be supplied directly, so the verdict path is testable without a subprocess.

**Report a hash mismatch as `not-implemented`.** It reads as honest — the gate did not run — and it
matches the vocabulary four other rows already use. Rejected because `not-implemented` is a
statement about this runner and a mismatch is a statement about a project's file, and because a
modified tool would then make a red verdict quieter rather than louder.

**Cache the hash for the duration of a run.** Two phases in one `gate verify` would hash the same
file twice. Rejected on the sentence: "checked before every run". The cost is a sha256 over a
binary per phase, which is nothing beside starting the process.

**Take the verdict from the findings rather than the exit status.** Findings present means fail,
absent means pass, which needs no synthesised finding and no convention about zero. Rejected
because section 14 says the exit status decides, and because it would make a tool that reports
warnings unable to pass.

**No timeout.** Nothing in the documents asks for one, and inventing a limit is inventing a rule.
Rejected: a limit that is written down is a decision, and an unbounded subprocess in the gate path
is a property nobody chose. The register says what it is and the configuration comment repeats it,
so a project that needs longer knows what it is arguing with — and that argument is a cheap change
rather than a discovery.

**Pass the artifacts to the command on stdin rather than a path.** It would let a gate run where
the repository is not on disk, and it would bound what the command sees. Rejected as a pretence:
the command runs in the repository with the pipeline's own access, so handing it the text as well
as the path would suggest a boundary that is not there.

**Let a project name a shell line rather than a path.** More convenient and it is what most tools
want. Rejected because a shell line cannot be hashed in any meaningful sense, and the hash is the
whole mechanism: a path to a file is the only thing whose content can be pinned.

<!-- xeno:section:impact -->
## Impact

**The chain of trust opens, by declaration only.** A project that declares nothing keeps the closed
chain: vendored plugin, hash, signed release, nothing executed. A project that declares one has
foreign code with repository access inside its gate path, and what the trail gains is the mark on
every statement that code made.

**`gate verify` runs foreign code too.** It recomputes verdicts, so the commands run there as well,
which means a project's CI verification step executes its own tool once per phase. That follows
from an external gate being a gate rather than from anything decided here, and it is the sentence a
project should read before declaring one.

**The invariants get their second path, which is what they were written for.** #66 added the rule
that an external finding keeps no decision and the comment saying it was checked at the verdict so
that an external gate would meet it. From here that is exercised rather than asserted.

**A foreign tool's wording becomes part of the finding ids.** An external finding's id is derived
from its cause, so a project that rewords its tool's messages invalidates every decision taken on
them. Section 4 says exactly this and calls the honest alternative a rule of the project's own.
It is now observable.

**`model.Project` gains a block that makes the runner execute something.** Every other field in
that file changes what the runner reads. This one changes what it runs, which is worth the comment
it gets in the scaffold rather than the schema line it would otherwise be.

**WP4 is complete after this.** The rule file format, the four levels, the two kinds, the
predicates, the two gates, the shipped set with its examples and hooks, and external gates. What
the plan called the substance of M1 is then built, and what remains of M1 is WP5 and WP8, both of
which already exist in part.
