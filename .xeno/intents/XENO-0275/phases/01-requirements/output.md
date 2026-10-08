---
intent: github.com/triplem/xeno#324
phase: 01-requirements
created: "2026-10-08T13:46:33Z"
schema_version: "1.0"
runner_version: dev+042c9bc.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2bd295150addaebea823988662af0f50e79b924c03454152a21d4cd4c23adfed
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. **The rendered GitHub wrapper carries the line, commented.** `xeno init --host github`
   produces a workflow whose `container:` block holds a commented `options: --user <uid>`
   beside the image, so an adopter reading the file meets it without being told.

2. **The reason is beside it, not elsewhere.** The comment says what the uid is for — the
   workspace the host mounts is owned by the runner's own account, and checkout runs inside
   the container — and that a hosted runner needs nothing because the image already matches
   it at 1001.

3. **The rendered file still parses as YAML with the line commented.** A comment that broke
   the document would be worse than the gap it fills.

4. **A test asserts it, not a reader.** The wrapper tests already assert that every rendered
   wrapper names the runner image; the same test file asserts this line for the GitHub
   wrapper, so a template edit that dropped it fails rather than passing quietly.

5. **GitLab's wrapper is unchanged**, and the test says so rather than leaving it to
   coincidence: a line about a uid would describe a problem that host does not have.
