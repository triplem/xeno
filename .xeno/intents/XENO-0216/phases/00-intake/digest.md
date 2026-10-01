---
intent: github.com/triplem/xeno#156
phase: 00-intake
created: "2026-10-01T14:36:08Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+2dd4dc9.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 8f7531e869730170e6dba53236684807fa7b080d5dd92eef13871018546816a2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
`ASSUMPTIONS.md` and `README.md` describe a runner that has moved on, and they are the two files read
first by anyone orienting. The worst of it is the README's trail section, which says the intents stop
after P0 because G-Freshness is half implemented: A6 closed with #63 and twenty intents from
`XENO-0107` to `XENO-0215` run all six phases and finish green, so the file asserts that a tool whose
claim is a judged trail has never judged a phase past intake. The register's "Not built" paragraph
lists the secret filter and the digest writer, both built and wired into `phase finish`, and calls
WP12 the GitHub adapter where the work is GitLab. Both files' lists are additive and neither was added
to. Behind all of it is a gap in the plan: WP16 covers the four documents under `docs/` and the site,
nobody owns these two, so they are maintained by whoever notices. That is written down as a finding
rather than fixed here, because naming an owner is a plan change and the first standing rule makes it
a person's commit.
