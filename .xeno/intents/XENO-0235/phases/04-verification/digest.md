---
intent: github.com/triplem/xeno#199
phase: 04-verification
created: "2026-10-03T18:24:51Z"
schema_version: "1.0"
runner_version: dev+a897f2a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 3f648b87612c8857a89f076fbdb1a4f6534428bd7fc05bb21165be93dfa04a69
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The property this change exists for has two halves and they are held differently. A test holds the
half a test can — the tree `init` writes hashes the same as the tree it copied from — and the other
half, that the binary's compiled-in value is that digest, is the release's step order, held by the
workflow and measured once by running those steps and then `phase finish` against what they
produced. With a binary built the release's way in an empty repository: `plugin from this release`,
34 files, manifest `0.33.0`, entry point `rwxr-xr-x`, `G-Supply pass`, and the first artifact
recording `plugin_version: 0.33.0`. Every one of those lines was impossible before — the first said
nothing, the second was 32, the fourth did not exist, and the fifth would have failed because the
tree was two files short of what the digest covered. `claude plugin validate` accepts `0.0.0-dev`,
which was the one thing promised before starting. Eighteen packages ok with five new cases and the
three existing vendor tests passing untouched, `gate verify` 285 at exit 0. Six gaps: the release
path is untested by a release and two of its mistakes would be quiet, with the ordering comment the
whole of the protection; nothing confirms a published binary carries the tag's digest; a partial
`cp -R` would leave the embedded copy and the hashed tree disagreeing, which three lines would catch;
a development build cannot exercise the embedded path at all; `mcp.json` is still absent and will
arrive the day it exists; and the embed directory is tracked, so a failed release can leave a
generated tree in a dirty checkout that nothing ignores.
