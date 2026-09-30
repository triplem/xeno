---
intent: github.com/triplem/xeno#147
phase: 00-intake
created: "2026-09-30T21:14:02Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+0a51653.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 24c37d56e0aa8ed1efbe2efbc77d3cd4cd6a1c23c0d2216b6ade5e4db525aa55
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Since #145 a repository can be initialised for GitLab and nothing answers for the adapter it names, which is
#104's remaining done-when. The intent has a second subject it does not deliver: #143 could not test whether
BranchRules is a boundary or an indirection, and this is the second adapter that decides it. What makes the
work more than transcription is allow_bypass, where GitHub has one boolean and GitLab has three lists of
grants plus allow_force_push, so reading them as one requirement is the judgement A65 put in the adapter. The
other half is the tier: both approvals endpoints answer 401 unauthenticated on Free and paid alike, so what a
real token sees there is the fact to observe, and it decides whether two of five rows report not-available with
a reason or an error that stops the command.
