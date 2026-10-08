---
intent: github.com/triplem/xeno#324
phase: 00-intake
created: "2026-10-08T13:45:47Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9d6187edd7282aa7adc606604cf1b79fec6d80f34582be4d046df02538eb7b91
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

> **github.com/triplem/xeno#324** — The generated GitHub wrapper cannot name a self-hosted runner's uid
>
> The image #282 publishes runs as uid 1001, which is the uid a hosted GitHub runner's work
> directory belongs to, so `actions/checkout` can write it from inside the container. A
> self-hosted runner, or an Actions Runner Controller deployment, may own that directory as
> another uid, and then nothing the image does can help.

The issue as `xeno phase start` read it, quoted rather than summarised.

What the tree says. `internal/scaffold/files/ci-github.yml` renders a job that names the
image through `container: image:` and says nothing about the user it runs as. The image's
own `safe.directory` exemption answers the other half of the uid problem — git refusing to
read a foreign-owned repository's config — and cannot answer this one, because a write is a
filesystem permission and not a git setting.

So an adopter on a self-hosted runner meets a permission failure from `actions/checkout`,
before any gate exists, with nothing in the generated file to connect it to the container
user. That is the whole of the defect: not a missing capability, a missing sentence in the
one file they will be looking at.

<!-- xeno:section:scope -->
## Scope

This intent adds a commented `options: --user` line, with its reason, to the GitHub wrapper
`xeno init` generates, and a test asserting the rendered wrapper carries it.

What it does not do.

**It does not default the uid.** 1001 is right for a hosted runner and wrong for anything
else, and this repository can neither read an adopter's runner nor guess it. A line that
guessed would break the common case to help the uncommon one.

**It does not become a `project.yaml` field.** Appendix A enumerates that file, so a
`runner_uid` would be an invented field. The uid belongs to the adopter's runner rather than
to their project.

**It leaves GitLab alone.** Its docker executor leaves the build directory world writable
under `umask 0000`, so writes succeed as any uid there and the image's own exemption is what
makes that host work. A commented line in that wrapper would describe a problem that host
does not have.
