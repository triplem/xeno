---
intent: github.com/triplem/xeno#199
phase: 01-requirements
created: "2026-10-03T18:22:38Z"
schema_version: "1.0"
runner_version: dev+a897f2a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 924f00a8283ae31a812cf1e4b876b4348b6339523a0bd9b1440ddd996d41cf32
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
A released binary in an empty repository vendors the plugin and says where it came from; that tree's
digest is the one the binary carries, so G-Supply passes against what the release just installed —
the property the whole change exists for and the one that was impossible before it. The adopter's
`plugin_version` is the release's number, stamped into the shipped copy, so section 13's one shared
number holds for anybody who receives a release. The digest is taken after the stamp, which is the
one thing here that fails silently if it is wrong: taken before, no adopter's tree could match, and
the release that made the mistake would pass while every project installing it failed. Every file is
vendored at any depth with nothing named, so `secrets.yaml` and `bin/xeno-env.sh` arrive and so does
whatever the plugin gains next; the entry point is executable, because an `embed.FS` reports every
file read-only. A development build carries no plugin and says so, naming `--plugin-from` rather than
the file it could not find, which is the answer `ExpectedDigest` already gives in the same situation.
The embedded copy wins over the directory, because a release must not be talked into vendoring a tree
whose digest it does not carry. The repository's manifest says `0.0.0-dev` and is never bumped again,
and `claude plugin validate` accepts it. The test compares digests rather than file lists, since a
list passes while a byte differs. Out of scope: the manifest check, offered and declined; the
resolution order; a signature; confirming a published binary; and `mcp.json`, which arrives the day
it exists because the walk copies what is there.
