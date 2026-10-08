---
intent: github.com/triplem/xeno#324
phase: 00-intake
created: "2026-10-08T13:46:15Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 9d6187edd7282aa7adc606604cf1b79fec6d80f34582be4d046df02538eb7b91
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake of #324: the generated GitHub wrapper names the image and not the user it runs as.

The image runs as uid 1001 because a hosted runner work directory belongs to 1001 and
`actions/checkout` runs inside the container. A self-hosted runner may own it as another uid,
and no image-side setting can fix a write: the `safe.directory` exemption answers git reading
a foreign-owned config, which is the other half of the problem. So the adopter meets a
permission failure from checkout with nothing in the generated file to connect it to the
container user.

Scope is the commented line plus a test. Not a default uid, because 1001 is right for the
hosted case and wrong elsewhere; not a `project.yaml` field, because Appendix A enumerates
that file and the uid belongs to the runner rather than to the project; and nothing for
GitLab, whose executor leaves the build directory world writable so writes succeed as any uid.

Two sections of the five, and the key had to be given by hand because the previous intent is
still on an unmerged branch — which is this phase learning record.
