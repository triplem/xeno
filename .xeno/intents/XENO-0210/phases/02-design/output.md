---
intent: github.com/triplem/xeno#110
phase: 02-design
created: "2026-09-29T19:03:27Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+454cfbf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d529c3ec4cadebcbd40fe40fcd78c7053f3f20c31b54f54c3739ccee64a20108
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

**`opts` gains `out` and `errw`, and `run` takes them.** `main` becomes
`os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))` and is the only place the real files
appear. `run` needs them as parameters because it prints the usage before an `opts`
exists.

**`report` and `suggest` become methods on `opts`.** Every command function already
takes `*opts`, so once it carries the writers nothing further is threaded, and the call
sites read `o.report(...)` where they read `report(...)`.

**Every `fmt.Print` becomes `fmt.Fprint` against one of the two.** Mechanical, and the
diff is then reviewable by the property that every hunk is a writer: nothing in it
changes a word, a format string or a code.

**The tests drive `run` with a `bytes.Buffer` each.** In process, one temporary
repository per case, and each assertion names the stream it read.

**A table for the staircase, a table for the dispatch.** The staircase's cases are a
refusal, an error, a red verdict, a green one and a provisional one; the dispatch's are
every name the usage string contains, which the test extracts from the string rather
than listing again, so a command added to one and not the other fails.

**The provisional case is asserted through `gate verify` on a real provisional phase.**
Not by calling `cmdGateVerify` with a fixture: the point is the exit code a wrapper
sees, so the test builds a phase that declares evidence, finishes it, and runs the
command.

**`--export` is asserted for what it does not print.** The suggestion is what it must
not carry, so the test reads standard output and requires the environment and nothing
else.

<!-- xeno:section:alternatives -->
## Alternatives

**Reassign `os.Stdout` in the tests and change no production code.** The cheapest option
by far: no refactor, no risk to fifty-five call sites, and the tests read what a command
wrote. It makes every test in the package share one mutable global, forbids `t.Parallel`
for all of them, and leaves the package's writers exactly as untestable as they are for
anybody who comes later. The issue weighed these two and chose the refactor; the saving
is real and it buys a package that stays hard to test.

**Test the built binary.** No refactor either, and it asserts the real thing including
`main`. It needs a build step inside the test, a temporary directory per case and a
subprocess per assertion, and it turns a fast package into a slow one. The issue names
it as the other option and prefers in process.

**Package level `var stdout io.Writer = os.Stdout`.** Smaller than threading, and it is
the same mutable global with a different name: two tests cannot run at once and the
writer is reachable from anywhere in the package.

**Assert coverage rather than behaviour.** A threshold in CI would stop the number
falling back. It would also be satisfied by testing whatever is longest, and the five
things the issue names are not the five longest.

**Thread an `io.Writer` through every function signature instead of putting it on
`opts`.** Explicit at every call. `opts` already reaches every command function, so this
would add a parameter to thirty functions to avoid adding a field to one.

<!-- xeno:section:impact -->
## Impact

`cmd/xeno/main.go`: two fields, one signature, two methods, and fifty-five writers. No
words, formats or codes change.

`cmd/xeno/main_test.go`: the staircase, the dispatch against the usage string, the
flags, the positional argument, and the streams.

Nothing outside `cmd/`. The runner, the gates and the model are untouched.

What a reader gains: the contract in the package comment is asserted, and the case
section 6 depends on has a test naming it.

What they do not gain: a tested `enforcement check`, which reaches a host, and a tested
`main`, which is one line and now contains the only reference to the real files.

What this costs: a diff where fifty-five hunks are the same edit, which is tedious to
review and is the one-time price of a package that was written to print rather than to
return.

The risk worth naming: the refactor's failure mode is a stream, not a code, and AC9 is
the answer to it.
