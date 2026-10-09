---
intent: github.com/triplem/xeno#226
phase: 02-design
created: "2026-10-09T12:58:16Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 802a4f07e4a3168105ffad82743cae22a21cb1d987c9a57094587756b2c0b11a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The design: one zensical.toml at the root with the comparison's configuration, two
palettes following the system preference with a toggle, a docs job building strictly on
every pull request and a deploy job alone holding the Pages permissions on the push to
main, Pages enablement through the action and in the settings script as a reported-and-
set block, Python pinned by major and bound by hand in the pin test, Zensical pinned
exactly with its dependencies resolving unpinned and the row saying so, the three README
links pointing at the host, docs added to the required checks for the maintainer to
apply, and the register saying what moved.
