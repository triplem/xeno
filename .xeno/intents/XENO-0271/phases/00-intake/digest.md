---
intent: github.com/triplem/xeno#273
phase: 00-intake
created: "2026-10-07T09:53:38Z"
schema_version: "1.0"
runner_version: dev+0768c44
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 531321fa2b5bed268d9849ad359230ee142f9f13db61771c9c7dc853ff1eefab
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
#273 asks whether the shipped template directories should carry the phase as a prefix and ends
"Does this make sense?". The answer needed measuring, because the directory name and the
template id are separate things — `Ref()` comes from `template.yaml`'s `id:` field and the
directory is only the path `Load` resolves — so a rename has two forms with different
consequences.

Renaming the directory alone keeps every check live and contradicts section 5: resolution
becomes per phase, the id decorative, and a project override moves. Renaming the `id:` with it
is what the issue asks for, and it retires the `strings_hash` recompute across all 495 artifacts
while `gate verify` goes on reporting every one verified, because `goneBundle` cannot tell a
renamed bundle from a missing one.

Both measured on throwaway copies, with the corruption that makes the difference visible. The
prefix would buy `ls` ordering and nothing else, since `template.yaml` already carries
`phase: 00-intake`.

In scope: A33's state column gains the measurement. Nothing is renamed, no normative document
is touched, and the two incidental findings — an all-digit hash scalar being skipped as a
non-string, and `goneBundle`'s blindness to a rename — are written on the issue rather than
absorbed.
