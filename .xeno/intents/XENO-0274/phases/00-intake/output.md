---
intent: github.com/triplem/xeno#95
phase: 00-intake
created: "2026-10-08T13:35:25Z"
schema_version: "1.0"
runner_version: dev+042c9bc
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d4322d4f833c38bd066c03cfe35aee5892929c227585f62a269b33a10e188322
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

> **github.com/triplem/xeno#95** — Replace dependabot with renovatebot
>
> Renovatebot does offer a dependency dashboard. This could be used instead of prs.

The issue as it stood at 2026-10-08T13:35:25Z, read by `xeno phase start` and quoted
rather than summarised.

What the tree says. There is no `dependabot.yml` under `.github/`, so nothing in this
repository configures either bot: whatever runs today runs from a host setting nobody
here can read, which is the same class of dependency as the protected branch settings
that `scripts/github-settings.sh` exists to write down. So the change is not a swap of
one file for another. It is the first in-repository statement of how dependencies are
kept current.

What makes this project unusual for such a bot. Every action is pinned to a commit sha
and every tool to an exact version, and since #317 two tests hold those pins against the
rows of `docs/supply-chain.md` in both directions. A bot that bumps a sha and leaves the
page alone therefore produces a pull request that fails `verify` — not on the bump, on
the page. That is the fact this intent has to design around, and it is why the dashboard
the issue asks for is the interesting half: a dashboard proposes, a person moves the row.

<!-- xeno:section:scope -->
## Scope

This intent writes `renovate.json`: what renovate watches in this repository, how it
groups what it finds, and that its default is the dependency dashboard rather than a pull
request per update.

What it does not do, and why.

**It does not switch dependabot off.** Dependabot is enabled, if it is enabled, in the
host's settings and not in a file here. Turning it off is an administrator's act, in the
same class as the protected branch and the squash setting, and belongs in
`scripts/github-settings.sh` or in the sentence that says what an administrator arranges
— not in a config file that cannot reach it.

**It does not teach the bot to move a supply-chain row.** A pull request that bumps a
pinned sha fails the pin tests until the page moves with it, and renovate has no reader
for a Markdown table. This intent records the consequence and leaves the remedy to a
person: either the dashboard is used as a list a person works from, which is what the
issue proposes, or a later intent makes the page derivable from the tree so there is one
copy to bump.

**It does not change a pin.** Nothing is bumped here. The config describes what would be
proposed; proposing it is the bot's run, not this commit.
