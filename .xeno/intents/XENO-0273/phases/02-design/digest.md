---
intent: github.com/triplem/xeno#153
phase: 02-design
created: "2026-10-07T14:24:43Z"
schema_version: "1.0"
runner_version: dev+6b48c17.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: d358c4cd6b87c132b7edd0c45313b9ae720520b775490a485f36299c7febd7f4
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The page carries the file rather than a description of it, because the example's comments were
written for a reader who does not write Go and publishing them is publishing the description.
The prose above the block says only what the block cannot: that this is the format Xeno reads
and where it is fixed, that the block is the file a test reads, and where to put the result.
The test for every sentence was whether removing it would leave a reader unable to use the
block, and the field descriptions all failed it.

The prose goes above the block, as `docs/commands.md` does, because a reader who has scrolled
past a hundred lines of YAML has already decided what the page is. The test compares the page's
first fenced block against the file on disk rather than an embedded copy, so there is one
authority and one direction of dependence, and it lives in `internal/index/index_test.go`
because the thing held is that package's fixture. The entry goes under "Running it", the page
is named `symbol-index.md` and not `index.md`, and nothing in the example points back at the
page: the dependence runs one way and the test's failure message names the page.

D-1 records the choice with the three options and the reading that separated them.

The alternatives are the prose page, a bare link, generating the page — which is WP16's end
state and makes this test the stopgap — moving the example under `docs/`, putting a schema in
section 5, and doing nothing.
