---
intent: github.com/triplem/xeno#81
phase: 00-intake
created: "2026-09-26T15:03:33Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+d4f92bf
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 7d812b138befa64c0cce48c4fdf1a41ccfcfc8c1e8229772a4d71c215510ac26
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

`NOTICE` carried `Copyright 2026 conet Deutschland GmbH`, a line A17 has called a
placeholder since it was written, because the wording was for the maintainer or legal to
settle rather than for the build. The holder is javafreedom.org.

The name was in three more places, all of them normative: section 15 of the process
definition twice, once as the holder and once in the sentence about the patent grant
applying to every user, and the implementation plan under licence, copyright and cost.
Changing `NOTICE` alone would have left the tree contradicting itself.

<!-- xeno:section:scope -->
## Scope

The holder, in the two documents and in `NOTICE`. conet is replaced throughout rather
than kept as the place of origin beside a different holder, which was the maintainer's
decision between two readings the old sentences allowed.

Not the security contact. `SECURITY.md` still carries a placeholder, and it waits on
#47, because what the file should name depends on whether this repository is published.
A17 therefore closes by half.

<!-- xeno:section:context-rationale -->
## Why this context

**The documents first, in their own commit.** The standing rule fixes that order, and
here it is not a formality: `NOTICE` is what a distribution carries and section 15 is
what the process claims about itself, so a tree where the two disagree would state the
holder twice and differently.

**The old sentences said two things at once.** "Developed at and for conet, which holds
the copyright" is a statement about origin and one about ownership, and a change of
holder does not settle what happens to the first. Both readings were put to the
maintainer, who chose to replace the name throughout, so section 15 keeps its argument,
that a tool asking organisations to run it in their pipelines should say whose tool it
is, with a different answer.

**What is deliberately untouched.** The vendored package's licence and notice stay as
they are: that copyright is Canonical's, and `NOTICE` describes it rather than claiming
it. The SPDX headers in the source carry no copyright line by A16, so a change of holder
does not reach them, which is the second time that assumption has paid for itself.