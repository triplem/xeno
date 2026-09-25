---
intent: github.com/triplem/xeno#45
phase: 00-intake
created: "2026-09-25T19:40:30Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+73849dd.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 344a481655da2b2b5b4cece43943c7b9b0e88c797b936ac6956dcfc2bc604ba3
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

"The instance" was the self managed GitLab server. It stopped being the target when
GitHub became one, and the word is still in about twenty places, used as if it named
something. So are Gitaly, the Community Edition, and a paragraph describing a module
path move that was declared void by the change that left it standing.

<!-- xeno:section:scope -->
## Scope

Every use of the word that means a host this project does not have, and the GitLab
specific terms around it. Each site judged rather than substituted: some become a self
hosted runner, some the host, and some are deleted because they described a plan that
no longer exists.

GitLab in 1.1 becomes the Enterprise variant, which changes the reasoning attached to
editions: Enterprise has the approval rules the Community Edition lacked.

<!-- xeno:section:context-rationale -->
## Why this context

This is the same failure as the passages repaired in #42, one level up. There, words
were left standing around a replaced sentence; here, terms are left standing around a
replaced concept. Both times the change searched for what it could name and reported
the work finished after fixing what it found.

The word is the one that carries the assumption without naming the host, which is why
searching for the host missed it.
