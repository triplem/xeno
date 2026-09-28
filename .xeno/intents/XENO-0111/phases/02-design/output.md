---
intent: github.com/triplem/xeno#111
phase: 02-design
created: "2026-09-28T17:33:06Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0e77cc2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: e02a19406aa6874c48bd01e7c970e421ee70e3219f29cad3beba057342256bf9
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

**The comment moves, with a blank line after it.** Above `printInit`, separated from
`printEnforcement`'s own comment by the blank line whose absence caused the defect. The
blank line is the fix; moving the paragraph without it would put both paragraphs on
`printInit` instead.

**The set is keyed on the path relative to `evidence/`.** A declaration carries a path
relative to the phase directory, prefixed `evidence/`, which is what the G-Evidence
closure resolves against. Stripping that prefix yields a key in the same coordinates as
a walk of `evidence/`, with no lookup and no dependence on the tree being judged.

**A path that is not under `evidence/` contributes no key.** That covers the empty path,
which is the defect that wrote `known["."]`, and it covers a path pointing elsewhere in
the phase directory. Both mean the same thing to this check: nothing in `evidence/` is
declared by them. One predicate rather than a special case for the empty string.

**The directory is walked, and only files are compared.** `filepath.Walk` from
`evidence/`, skipping the root and every directory, keyed by the path relative to the
root with forward slashes. A directory is not an entry that can be declared, so AC5
holds by construction rather than by a rule that excuses directories after the fact.

**`filepath.Walk` rather than `filepath.WalkDir`.** This package names its findings
variable `fs` in nearly every function, and `WalkDir` takes an `fs.DirEntry`. Importing
`io/fs` would shadow the name here and read as a typo everywhere else in the file. The
older signature costs a `Stat` per entry on a directory that holds a handful of files.

**`attached.yaml` stays keyed by its own name.** It is the one file in `evidence/` that
belongs there and that no declaration points at. Its key is `attached.yaml`, which is
already relative to `evidence/`, so it needs nothing said about it.

<!-- xeno:section:alternatives -->
## Alternatives

**Key on the relative path and keep reading one level.** The smallest diff, and it fails
AC5. The key would name `logs/a.txt` while the entry reads `logs`, so a declared file
under a subdirectory would leave its directory reported as undeclared. It trades a check
that is wrong about which file is declared for one that is wrong about whether a
directory may exist, which is not progress.

**Keep the basename key and add a finding for any subdirectory under `evidence/`.** It
makes the flat assumption true by enforcing it, and it is a new rule. Section 4 permits
the nesting, so forbidding it is a specification change first, by the second standing
rule. Rejected on authority, not on merit.

**Compare phase-relative paths on both sides.** Prefix each walked path with `evidence/`
rather than stripping the prefix from each declaration. Equivalent, and it puts the
constructed string on the side with more entries while leaving the declaration's own
coordinates alone. Chosen against because the check's subject is `evidence/`, so its
keys read better relative to it, and because `attached.yaml` would then need the prefix
too.

**Fix only the `known["."]` entry and leave the basename key.** What the issue could
have been read as asking. It removes the inert entry and leaves the defect that can
produce a wrong finding, which is the wrong half to do alone.

**Split the intent in two.** One for the comment, one for the check. Defensible, and it
loses the one thing the pair records: that both were found in a single reading, and that
both were correct against today's data.

<!-- xeno:section:impact -->
## Impact

`cmd/xeno/main.go`. One paragraph moves up eighteen lines and gains a blank line after
it. No code changes.

`internal/gates/gates.go`. `undeclaredEvidence` gains a predicate that turns a declared
path into a key relative to `evidence/`, and its directory read becomes a walk. The
comment on the function gains the reason for both, since a walk where a `ReadDir` would
do looks like more work than the check needs until the reader is told which case it is
for.

`internal/gates/*_test.go`. Tests for AC2 to AC5. The nested cases need a phase fixture
that writes a subdirectory under `evidence/`, which no existing fixture does.

Behaviour on everything that exists: identical. `Attach` writes one flat level, so the
walk returns what the read returned and the stripped keys equal the basenames. `gate
verify` is the check on that claim, over every verdict the repository carries.

Performance: one `Stat` per file in one directory per gate run, on directories that hold
a handful of files. Not measured, and named because a walk in a gate is worth a
sentence.

What a later reader gains: the check now says what it means. Keyed on the path a
declaration carries, against the paths the directory holds, with nothing in the middle
that happens to agree.
