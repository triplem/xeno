---
intent: github.com/triplem/xeno#346
phase: 00-intake
created: "2026-10-09T15:11:09Z"
schema_version: "1.0"
runner_version: dev+9590797
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 87865ec7e33a1a31be6160e7b377f4d1859407074bdbcb17ae6af3dff0b16932
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake for #346: the label script reads the first page of labels and no more, so on a repository past thirty its second run tries to create the label again and exits 1, reproduced on this one. The fix is a read by name on GitHub and a per-page bound on GitLab, with the plugin test holding both flags; nothing in the clause moves.
