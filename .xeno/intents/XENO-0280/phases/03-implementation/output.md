---
intent: github.com/triplem/xeno#226
phase: 03-implementation
created: "2026-10-09T13:05:25Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 199a3d8f522fb8e8ab61a951875f21bcb3b5a67872490fc66979528d3c3b8d62
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: implementation@1.0.0
strings_hash: 2a598b5919612a4456684ef6d3aac06c0c8c5be86bb4732c3738eb810e8949fa
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Implementation

<!-- xeno:section:changes -->
## Changes

Nine files, by concern, plus the plan paragraph that preceded them in `5056d9b`.

**The configuration: `zensical.toml`.** The comparison branch's file, with a header that
says what it is and why it is strict, `docs_dir` `docs`, `site_dir` `.xeno/local/site`,
the eleven documents in three navigation groups, permalinks, the Mermaid superfence and
snippets with path checking, and the Material theme with three palette entries: one
following the system preference with a toggle, then `default` and `slate` with primary
black and a toggle each. The built `index.html` carries both scheme attributes.

**The workflow: `.github/workflows/docs.yml`.** `docs` on every pull request and the push
to `main`: checkout and `setup-python` pinned by sha, `python-version: '3.14'`, `pip
install zensical==0.0.69`, `zensical build --clean --strict`, and the site uploaded as
the Pages artifact. `deploy` only on the push, needing `docs`, with `pages: write` and
`id-token: write` in the `github-pages` environment: `configure-pages` with `enablement:
true`, then `deploy-pages`. Three Pages actions and `setup-python`, each by sha with the
tag beside it, read off the tags on 2026-10-09.

**The three links: `docs/README.md`.** The paragraph naming the three root files now
points at them on the host and says why, in a paragraph replaced rather than edited.

**The pin table: `docs/supply-chain.md`.** Six rows after the Node row: `setup-python`, the
Python toolchain by major, `zensical` with its fetch and the sentence about its
dependencies, and the three Pages actions; and a paragraph before the `docker` one on
why the one unpinned install is tolerable here and would not be on the gate path.

**The test: `internal/model/supply_chain_test.go`.** A `pythonVersion` shape beside
`nodeVersion`, a `python` field on `treePins`, the same one-row check the Node major has,
and `Python toolchain` in `boundByHand`. `go test ./internal/model/` green with the six
rows.

**The host: `scripts/github-settings.sh`.** A `pages` block before the branch protection,
reporting the site where it is there and creating it with `build_type: workflow` where
it is not, in the shape the vulnerability-reporting block has; and `docs` as the eighth
required context. `sh -n` clean. Not run from this session, for the reason the design
gives.

**The reader: `CONTRIBUTING.md`.** The sentence listing the checks names `docs` and what
it is red on.

**The register: `docs/assumptions.md`.** The not-built paragraph loses the documentation
site, and A106 records strict on pull requests, deploy on the push alone, Python by
major and the unpinned dependencies.

**Measured on the branch.** `zensical build --clean --strict` in 0.8 s, no issues, the
site under `.xeno/local/site` and `git status` silent about it; a deliberately broken
anchor on a scratch copy aborted the strict build naming `commands.md:98:49`; `go test
./...` green, `gofmt` and `go vet` clean, `xeno gate verify` 558 verdicts; the workflow
parses with the two jobs and the TOML with three palette entries.

<!-- xeno:section:deviations -->
## Deviations from the design

One, against the design's second decision, and it is a count.

**Three palette entries, not two.** The design said two palettes, `default` and `slate`,
"the first following the reader's system preference". In Zensical's shape, as in
Material's, following the system preference is a third entry with no scheme of its own
and a toggle that moves to the next, so the file carries three entries for two schemes.
What the reader gets is what the design meant: light and dark, the system's choice
first, and a toggle through all three states. The light and dark attributes are both in
the built page.

Nothing else departs. The strict job, the deploy split, the enablement in the action and
in the script, Python by major and bound by hand, Zensical pinned with its dependencies
named as unpinned, the three links to the host, the required context left to the
maintainer and the register rows are as decided.
