---
intent: github.com/triplem/xeno#355
phase: 00-intake
created: "2026-10-10T12:11:03Z"
schema_version: "1.0"
runner_version: dev+90c7227.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 57b0925079ffdfab468947b6826bfc7ee272173b5ccd7b165fa1f73eb5abdc4b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
The intake of #355, closing #356 with it, because neither half is worth releasing alone:
one is why the image step failed and the other is why the failure is permanent.

Everything to be decided was decided before this phase. Four decisions are recorded, all
the maintainer's and all taken on the two issues: the base is mirrored into `ghcr.io`
rather than pulled with a Docker Hub credential; the mirror is named by the release and
not by the `Dockerfile`, so `ARG BASE` and the regex manager from #351 stay as they are;
the release fills the mirror on a miss, which makes it a cache and retires the question of
who refreshes it; and v0.60.2 gets its image through the dispatch rather than by hand. No
open question follows from them.

The premise both issues rest on was measured here rather than carried over. `ghcr.io`
answers 200 for `0.60.0` and `0.60.1` and 404 for `0.60.2`, read in one loop with one
token and one Accept header, so the three answers differ in the tag and in nothing else.
That is what makes the 404 evidence: an unauthenticated probe, a wrong Accept header or a
mistyped repository path would have returned 404 for all three.

Two facts about the tree were read before the scope was fixed, and each moved it. The
`digestPin` regex in `internal/model/supply_chain_test.go` reads digests out of every
workflow as well as out of the `Dockerfile`, so a mirror reference carrying a digest in
`release.yml` would become a second pin the page is asked to carry — which is the reason
the release reads the digest out of the `Dockerfile` at run time. And the third column of
the page's table is read by nothing, which is what makes the row's change a documents
change with no test behind it.

What is left for later phases is one claim this phase could not check: that
`docker buildx imagetools create` copies a manifest index with its digest intact. The
whole arrangement rests on it, the local docker has no buildx plugin, and P4 is where it
is shown against a registry rather than asserted.

Six files and 72,508 bytes of context, enumerated rather than globbed. Three sections and
four decisions, of the five sections this template defines.
