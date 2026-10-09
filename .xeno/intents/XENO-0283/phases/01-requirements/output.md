---
intent: github.com/triplem/xeno#346
phase: 01-requirements
created: "2026-10-09T15:11:28Z"
schema_version: "1.0"
runner_version: dev+9590797
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: be2c748b2140cb5cf14b7a6bfa8b5f25455f9df0b4160e6b0b042c2ce2a314e8
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

1. **The GitHub read asks by name.** The line before `gh label create` reads `gh label
   list --search "$label"` and keeps `grep -qx "$label"` on the names it returns, so a
   label whose name merely contains the string, or whose description does, is not
   taken for the label.
2. **The GitLab read names its page.** The line before `glab label create` carries
   `--per-page 100`, with the bound and its reason on the line.
3. **The second run exits 0.** `sh .xeno/plugin/bin/xeno-labels.sh github triplem/xeno`,
   with the label present and the repository past thirty labels, prints *exists on
   triplem/xeno; nothing changed* and exits 0; recorded in P4 from a real run.
4. **The test holds the read's shape.** `TestTheLabelScriptCreatesTheLabelTheClauseNames`
   fails on the script as it stands on `main` and passes on the fixed one, by requiring
   `--search` and `--per-page` among the strings the script carries.
5. **Nothing else moves.** The constants, the description, the colour, the usage line,
   the init line and section 12 are as they were; `git diff --stat` shows the script
   and the test, and the trail.
6. **The suite and the gates stay green.** `go test ./...`, `gofmt -l`, `go vet ./...`,
   `sh -n` on the script and `./xeno gate verify` exit 0.

<!-- xeno:section:non-goals -->
## Non goals

- **Paging the GitLab list.** A parser in POSIX shell for a case past a hundred labels, which no tracker using labels for work packages has; the bound is said on the line instead.
- **Running a host in the test.** Section 7 keeps the network and the credential out of the gate path, and a test that reached a host would put both in the suite.
- **A read-back in `xeno init` or `enforcement check`.** A107 says why the label is never read back by the runner; this intent is about the script's own read.
- **Correcting XENO-0282's sealed gap paragraph.** It is sealed and was right as far as it went; this intake says what it missed.

<!-- xeno:section:constraints -->
## Constraints

- Section 12 names the script and says it creates the label; it says nothing about how the script reads, so no clause moves and the first standing rule is not touched.
- The script stays POSIX shell with `set -eu` and no dependency beyond `gh` or `glab`, which is the shape the plugin's `bin/` has and `G-Supply` will compare on a release.
- The test stays a shape test over the script's text, beside the entry point's, so that the suite proves the same thing offline as it did before.
- Prose at 88, commit subjects as Conventional Commits, `Refs #346` in the footer and `Closes #346` on the pull request.
