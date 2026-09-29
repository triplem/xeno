---
intent: github.com/triplem/xeno#110
phase: 02-design
created: "2026-09-29T19:03:27Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+454cfbf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d529c3ec4cadebcbd40fe40fcd78c7053f3f20c31b54f54c3739ccee64a20108
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Five alternatives and the first is the one that had to be argued rather than dismissed: reassigning os.Stdout
in the tests needs no production change at all and buys a package that stays hard to test, which is the whole
reason this one is at zero coverage. The decision that will matter later is extracting the dispatch cases
from the usage string rather than listing them again, so a command added to one and not the other fails.
