---
intent: github.com/triplem/xeno#111
phase: 00-intake
created: "2026-09-28T17:31:43Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0e77cc2.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 503330549dcdf03e764b0b0ea05f55fd793ce1d78319c88c3924896d5ee1e27c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

Two readings of the same afternoon, neither of them a wrong verdict.

**A doc comment sits on the wrong function.** In `cmd/xeno/main.go` the paragraph
beginning "printInit says what was done, what was left alone, and what a person still
has to do" stands after `printInit`, immediately before the comment that belongs to
`printEnforcement`, with no blank line between the two. Godoc concatenates them, so the
sentence carrying the reason — a first contact that leaves the project believing the
gate is binding when it is not is worse than no first contact — is documentation of the
enforcement report, where it is simply false. The comment says something true about a
function nobody reading it is looking at.

**`undeclaredEvidence` keys its set on a basename.** In `internal/gates/gates.go` the
line `known[filepath.Base(e.Path)] = e.Path != ""` carries two defects, both harmless
today.

A declaration with no path writes `known["."] = false`, because `filepath.Base("")` is
`"."`. No directory entry is ever named `.`, so nothing follows from it; the map carries
an entry that means nothing, and the reader has to work out that it is inert.

The key is a basename and the declaration carries a path relative to the phase
directory, prefixed `evidence/`. The comparison is against `os.ReadDir` of `evidence/`,
which returns one level and is flat today, so the two agree by accident of the data
rather than by construction. They stop agreeing as soon as an item is declared under a
subdirectory: two items with the same basename in different directories stand in for
each other, and one declared deeper marks a same named file at the top level as
declared. Section 4 forbids neither nesting.

Both are found by reading and neither changes a verdict on any artifact this repository
carries. What they change is what the next reader believes, which is the thing a record
of this kind is for.

<!-- xeno:section:scope -->
## Scope

The comment moves above `printInit`, with the blank line that separates it from
`printEnforcement`'s own. Nothing else about either function changes.

`undeclaredEvidence` keys its set on the path relative to `evidence/` and skips a
declaration with no path rather than recording one. For the key to mean anything against
the directory, the directory is walked rather than read one level deep, so a file
declared under a subdirectory is compared against the path that names it.

Tests for both halves of the second fix: a declaration under a subdirectory that matches
only the file it names, and two files of the same basename in different directories that
do not stand in for each other.

Not the shape of a declaration. Whether a `path` may point outside `evidence/` at all is
section 4's question, and this change treats one that does as declaring nothing in
`evidence/`, which is what it already did.

Not a new finding for a nested directory. A directory whose files are all declared is
not a finding today and is not one after this.

No documentation change. Section 4 permits the nesting already; the code did not.

<!-- xeno:section:context-rationale -->
## Why this context

**Two unrelated defects in one intent, because one reading found them.** They share no
file and no gate. What they share is the reason they were invisible: both are correct
against the data that exists and wrong against the data the specification allows.
Splitting them into two intents would double the record of a single reading and hide
that they were found the same way, which is the thing worth remembering.

**The comment is the more serious of the two.** Nothing executes differently either way,
and that is exactly why it lasts: a comment attached to the wrong function is not caught
by a test, a gate or a reviewer reading the function it was written for. Godoc will keep
publishing the sentence about first contact as a claim about the enforcement report
until somebody reads the two paragraphs together and notices the missing blank line.

**The basename key is a latent defect, and it is named as one.** It cannot produce a
wrong finding on any artifact in this repository, because nothing has ever put a
subdirectory under `evidence/`. `Attach` writes `evidence/<name>` at one level, so the
only way to nest one today is by hand. That is why the issue is `wp1` rather than a
verdict fix, and why the tests matter more than the change: they are what makes the
defect impossible to reintroduce once the data that exposes it finally exists.

**Walking the directory is part of the fix, not an extension of it.** Keying on the
relative path while reading one level deep would trade one disagreement for another: the
key would name `logs/a.txt` and the entry would be `logs`, so a declared file under a
subdirectory would leave its directory reported as undeclared. The two halves have to
change together or the check says something new that is also wrong.

**What this intent does not touch.** No gate, no field, no dependency, nothing under
`docs/`. Section 4 already permits what the code failed to survive, which makes this the
code catching up rather than a decision about the process.
