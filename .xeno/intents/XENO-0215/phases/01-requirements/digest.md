---
intent: github.com/triplem/xeno#151
phase: 01-requirements
created: "2026-10-01T13:53:05Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+c54db93.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: dde464e9b230bdeee958110c72948291dcf333bfca30d997d35c8104f497ed04
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Eight criteria. The fourth is the load bearing one: absent, unreadable, malformed and stale all return no index and
no error, each with a reason, because a phase must run without one and absence is the state every repository is in.
That turned out to constrain the reader's signature rather than its behaviour — an error return invites a caller to
treat an absent index as a problem. The seventh holds A42's property against the first package that could break it
by being useful, since an index is a tempting thing for a gate to consult and exactly what must not happen. The
sixth has this repository produce an index for its own Go source from the standard library, which exercises the
format against a tool's output rather than a fixture the author wrote, and the non-goals say plainly that this is
the project being a project and not Xeno shipping an indexer.
