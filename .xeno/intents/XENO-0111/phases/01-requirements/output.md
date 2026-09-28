---
intent: github.com/triplem/xeno#111
phase: 01-requirements
created: "2026-09-28T17:32:20Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0e77cc2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 580ecc08eed0854893b081a8ad6e3a58c8ca84b3a0cc407598e016cb7a785144
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: by-hand
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

**AC1.** The paragraph about what `init` did and what a person still has to do stands
above `printInit`, and `go doc` for `printEnforcement` shows only the paragraph about
the enforcement report.

**AC2.** A declaration with no `path` adds nothing to the set `undeclaredEvidence`
compares against. In particular no entry keyed `.`.

**AC3.** A declaration whose path is `evidence/logs/a.txt` marks `logs/a.txt` as
declared and marks nothing else. A file `a.txt` at the top of `evidence/` is reported
undeclared while only the nested one is declared.

**AC4.** Two files of the same basename in different directories under `evidence/` are
judged separately: declaring one leaves the other reported.

**AC5.** A declared file under a subdirectory leaves no finding about its directory.
Walking the tree means the directory is not itself an entry to be declared.

**AC6.** Everything already green stays green. `evidence/` as `Attach` writes it, one
flat level with `attached.yaml` and a copied result, produces the same verdicts as
before, and `gate verify` matches every verdict in the repository.

**AC7.** A declaration whose path points outside `evidence/` marks nothing inside it,
which is what it did before.

<!-- xeno:section:non-goals -->
## Non goals

Whether a declaration may nest at all. Section 4 shows `path: evidence/junit.xml` as an
example and forbids nothing further. This change makes the check survive nesting; it
does not decide that nesting is good practice, and it adds no finding against it.

A finding for an empty directory under `evidence/`. Walking means directories are not
entries, so an empty one is invisible to this check. Whether it should be is a question
for the directory check above it, not for this one.

The other users of `e.Path`. The G-Evidence verify closure resolves `dir + "/" + path`
and is correct for a nested path already. It is read during this work and left alone.

`filepath.Base` elsewhere in the tree. Only this one call sites its key on a basename;
the rest are building messages or names, and a sweep for the pattern is not this intent.

Godoc output for anything but the two functions named. The file has other comments and
they are not audited here.

<!-- xeno:section:constraints -->
## Constraints

Nothing under `docs/` changes. Section 4 already permits the nesting the code failed to
survive, and the doc comment is a comment.

One dependency, unchanged. Walking a directory is `path/filepath`, which the file
already imports.

The findings variable in this package is named `fs` throughout. That rules out importing
`io/fs` here, so the walk uses `filepath.Walk` and `os.FileInfo` rather than `WalkDir`
and `fs.DirEntry`. A shadowed package name in one function of a file that uses `fs` for
findings in twenty others is a worse trade than an older signature.

The comparison key must be derivable from what a declaration carries, with no lookup. A
declaration gives a path relative to the phase directory; the check reads `evidence/`;
the key is what remains after that prefix. Anything that needed the file system to
reconcile the two would make the check depend on the tree it is judging.

`evidence/attached.yaml` stays in the set by name, because it is the one file in the
directory that no declaration points at and that belongs there.

One intent, one work package. The issue carries `wp1`, the branch carries this intent,
the commits reference #111.
