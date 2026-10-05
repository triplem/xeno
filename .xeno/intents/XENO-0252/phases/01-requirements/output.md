---
intent: github.com/triplem/xeno#205
phase: 01-requirements
created: "2026-10-05T16:27:02Z"
schema_version: "1.0"
runner_version: dev+97d07da.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4634f3e2c35ad99626d574711f32f5fd6d16b26d7bd07dc8891864bb9d34789f
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. `model.LocalDir(root)` returns the local data location: the value of `XENO_PLUGIN_DATA`
   where set and non-empty, `.xeno/local` under `root` where not. An absolute value is returned
   as it stands; a relative one is joined to `root`.

2. `model.LocalPath(root, parts...)` joins inside that directory, and the name of the variable
   and the default are exported constants so that nothing writes either string twice.

3. All six literals are gone. The run marker, `phase.env` in both places that build it, the
   enforcement report and the cost ledger resolve through the helper; `grep` for the literal
   finds only the `.gitignore` entry and comments.

4. `cmd/xeno/main.go` no longer repeats the `phase.env` path. One expression builds it.

5. The `.gitignore` entry `xeno init` writes stays `.xeno/local/`, with a comment saying why it
   is the one site that does not resolve.

6. Setting `XENO_PLUGIN_DATA` moves the run marker, `phase.env`, the enforcement report and the
   ledger, which is the behaviour the export promises and is asserted rather than described.

7. Unsetting it, or setting it empty, leaves every path exactly where it is today. This is the
   criterion that protects every existing repository and it is asserted per path.

8. `./xeno gate verify` exits 0 with the 387 verdicts that exist now intact. Nothing under this
   directory is hashed, so the expectation is that the change is invisible to it, and the command
   is what confirms the expectation rather than the reasoning.

9. `go build`, `go test ./...`, `go vet ./...` pass and `gofmt -l` outside `vendor/` prints
   nothing.

10. `docs/assumptions.md` gains a row: the variable is read, the directory holds nothing hashed,
    and that is why A89's objection does not reach it. A reader meeting the difference between
    this and `XENO_PLUGIN_ROOT` should find it written down.

11. One commit, `Closes #205`, and the issue carries `wp7`.

<!-- xeno:section:non-goals -->
## Non goals

Not `XENO_PLUGIN_ROOT`. A89 removed it from section 7 and gave the measurement for why; this
intent does not reopen it and the register row says what makes the two different.

Not a flag. Section 7 normalises through the environment and the entry point exports this one
already; `--local-dir` would be a second way to say the same thing and an addition to the command
surface.

Not a migration. A person who sets the variable gets an empty directory; nothing moves existing
state or looks for it in the old place. `retention` is `local_days: 30`, so what is under there
is transient by design.

Not a change to the entry point. `xeno-env.sh` already sets the variable to an absolute path
built from the git root, which is exactly what the resolver has to accept, so the script is read
and left alone.

Not a change to section 7, or to what the variable means. "Local data location" is what it was
and what it now is; the only thing that changes is that something reads it.

Not validation of the value. A path that cannot be created fails where it is used, as it does
today for `.xeno/local`, and a resolver that checked would be checking at a moment when nothing
is being written.

Not the `.gitignore` entry. It describes this repository rather than a person's redirection, and
criterion 5 keeps it fixed deliberately.

<!-- xeno:section:constraints -->
## Constraints

The exported value is absolute. `xeno-env.sh` sets `${XENO_PLUGIN_DATA:=$root/.xeno/local}` from
the git root, so the value the runner meets in practice is an absolute path, and a resolver that
joined it to the root would produce a path inside a path. That is the detail the issue does not
contain and section 7 does not either; it comes from reading the script.

Nothing under the directory is hashed, and that is the licence for the whole change. The run
marker, `phase.env`, the ledger and the enforcement report are all outside `artifacts_hash`, so
no verdict can be made to depend on the environment. A89's measurement is the counter-example to
stay away from: an override of the plugin root leaves `gate verify` at exit 0 while G-Policy
silently stops judging, because `rules_hash` resolves from that tree.

Backwards compatibility is absolute here. Every existing repository has state under
`.xeno/local/` and no variable set, so the unset path has to resolve to exactly what the literals
produced. Criterion 7 asserts it per path rather than once.

`internal/cost` must not gain an internal dependency beyond `internal/model`. It imports only
the standard library and `yaml` today; `internal/model` imports only the standard library, so the
resolver can live there and nowhere that would make `cost` depend on the runner.

A relocated directory may not exist yet. The writers create what they need today —
`phase start` makes the marker's parent — so the resolver returns a path and creates nothing,
which keeps it a pure function and keeps the error at the write.

One intent, one branch, `Closes #205`, and the issue carries `wp7`.
