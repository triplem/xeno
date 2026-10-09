---
intent: github.com/triplem/xeno#226
phase: 00-intake
created: "2026-10-09T12:56:57Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 028c682e00440bc21cbd179f1ff5d1b7be3cddebd9b9add5ff74637bd2a32de1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake for #226. The decision was taken on the issue: three generators built a page
each, Zensical was recommended because MkDocs's own theme authors say to leave MkDocs
and build Zensical as its replacement, and the maintainer chose it with a light and a
dark mode. The plan's WP16 paragraph moved first, in its own commit. The scope is the
documents part of WP16: a Zensical configuration, a docs job that builds strictly on
every pull request and deploys to Pages on a push to main, the three README links the
strict build refused, the pin table rows and the Python toolchain bound by hand, the
host settings the maintainer applies, and the register. The rest of WP16, the ids for
the seven documents without them, diagrams and snippets stay out.
