---
intent: github.com/triplem/xeno#121
phase: 00-intake
created: "2026-09-28T19:39:51Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c2ba6b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 10e63414aa28bd9b452cf5c56bd72c426d48ea82a21b73f480af9f292288e42a
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

Nothing has been released since v0.18.0, tagged on 2026-09-26. The release workflow has
failed on every merge to `main` since #86 landed that afternoon, and each run computes
v0.19.0 and discards it.

`@semantic-release/git` commits `CHANGELOG.md` back to `main` with a direct push, and
the protection refuses it:

    remote: error: GH006: Protected branch update failed for refs/heads/main.
    remote: - Required status check "verify" is expected.
    ! [remote rejected] HEAD -> main (protected branch hook declined)

The push can never carry that check, and `release.yml` states the reason one line above
the plugin that depends on it: a push made with `GITHUB_TOKEN` triggers no workflow. So
`verify` is never reported for the changelog commit, and a required check that is never
reported is never satisfied. With `enforce_admins: true` there is nothing left to step
over.

Nothing is half published. The failure falls in the `prepare` step, before `publish`, so
a failed run creates no tag, no release and no assets. One run in the window reported
success and published nothing either, for an unrelated reason the log names: the local
branch was behind the remote one, which is two merges racing.

Two deliberate decisions collide. A23 assumed the loop was closed and its own status
column names the condition that broke it, a protected `main` refusing the commit unless
the token is on the bypass list. #86 then removed the bypass list on purpose, because a
required check an administrator can step over is a report rather than a barrier. Neither
side is misconfigured.

<!-- xeno:section:scope -->
## Scope

`@semantic-release/git` and `@semantic-release/changelog` leave `.releaserc.json`.
Nothing pushes to `main` and the protection stays as #86 set it.

`CHANGELOG.md` is removed rather than left frozen at 0.18.0. Its own header says it is
derived from the commit history, so a file stopped four releases short reads as a
repository that stopped releasing.

`SUPPLY-CHAIN.md` and `audit.yml` stop pinning two plugins that are no longer installed,
and `release.yml` loses the `extra_plugins` block that brought them.

The comments in `release.yml` that describe writing the changelog back, and the property
that a `GITHUB_TOKEN` push triggers nothing, are corrected. That property is now the
reason the arrangement could not work rather than the reason it was safe, and a comment
still claiming the second would mislead precisely the reader trying to understand the
first.

A23 is superseded, not amended. Its subject was that the changelog is written back; that
is no longer what happens.

Not the protection, and not a credential. The alternative was a bypass actor or a token
that can push past a required check, which reopens what #86 closed. Refused in #121 and
not revisited here.

Not the plan. `docs/implementation-plan.md` was changed first, in #122, by a person,
which is what let this intent exist at all.

<!-- xeno:section:context-rationale -->
## Why this context

**The specification went first, and that is why this is a configuration change.** The
plan promised a `CHANGELOG.md` derived from the commit history. While that sentence
stood, dropping the file would have been the code disagreeing with the specification,
which the first standing rule settles in the specification's favour. #122 changed the
sentence in its own commit and its own pull request, because squash is the only merge
method here and a bundled change would have collapsed the ordering the rule is about.

**The file goes rather than freezes, because a stale generated file lies.**
`CHANGELOG.md` opens by saying it is derived from the commit history and that editing it
by hand changes the rendering and not what it renders. Left in the tree it would say
that about a rendering four releases out of date, and the next reader would take it for
the current state or, worse, correct it by hand. The content is not lost: it is in the
release notes of every release and in the history of the file itself.

**Two plugins, not one.** `@semantic-release/git` is what pushes, and
`@semantic-release/changelog` is what writes the file it pushes. Removing only the first
would leave a plugin generating a file nothing commits, which is a step that produces
nothing and a dependency that has to be audited anyway.

**The comments are part of the change and not tidying.** `release.yml` explains that a
`GITHUB_TOKEN` push triggers no workflow, offered as the reason the changelog commit
starts no loop. The same property is why the commit could never satisfy a required
check. A reader arriving at that comment after this change would find the correct fact
supporting the wrong conclusion, which is worse than no comment, and the same argument
the misplaced doc comment of
#111 made about godoc.

**A23 is superseded rather than amended.** The rows amended in this repository were
sentences that were right and unfinished. A23 says the changelog is written back to
`main`, which stops being true, so the row records what replaced it instead of growing a
qualification that contradicts its own subject.

**What stays broken on purpose.** A release is still cut from a tree whose protection
nothing can bypass, so any future step wanting to write to `main` from a pipeline meets
the same wall. That is the arrangement #86 chose and this change accepts rather than
works around.
