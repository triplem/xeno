---
intent: github.com/triplem/xeno#226
phase: 02-design
created: "2026-10-09T12:58:16Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 802a4f07e4a3168105ffad82743cae22a21cb1d987c9a57094587756b2c0b11a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: design@1.0.0
strings_hash: f42fab544ec4c03fdba606e3326a55706a60d63378e1bb7f57a1a92e5c93bfb3
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Design

<!-- xeno:section:decisions -->
## Decisions

**One configuration file at the root, in Zensical's own TOML.** `zensical.toml` beside
`renovate.json` and `.golangci.yml`, which is where this repository keeps a tool's
policy; `docs_dir` is `docs`, `site_dir` is `.xeno/local/site`, so the build lands where
every other pipeline product lands and git never sees it. The navigation is three groups,
the overview, the four documents, and the pages about working with the tool, in the
order a reader meets them; the plan's "the documents part" is the second group and the
rest is already there to be listed. The configuration is the one the comparison branch
carried, kept in this intent so that the deleted branch's evidence survives in the
decision it led to.

**Two palettes and a toggle, by media query first.** `default` and `slate`, Material's
light and dark, each `primary = "black"` so the site is the documents and not a colour,
each with a toggle icon and name, and the first palette following the reader's system
preference, which is the shape Zensical's own template ships. That is the maintainer's
"light and a dark mode" with the choice left to the reader, which is what the plan
paragraph now says.

**Strict on every pull request, deploy only on the push to `main`.** Two jobs in one
workflow. `docs` checks out, pins Python by major through `actions/setup-python`,
installs `zensical==0.0.69`, builds with `--strict` and uploads the site as a Pages
artifact; it runs on every pull request and on the push and holds `contents: read`.
`deploy` needs `docs`, runs only on the push to `main`, holds `pages: write` and
`id-token: write`, and deploys the artifact. The split is the one `xeno.yml` makes
between `verify` and `report` and for the same reason: the job that runs on a fork's
pull request must not be the job that can publish. A broken link in a pull request is
a red `docs` check, which is what makes WP16's "consistency is mechanical" true for
cross references from the first day.

**`configure-pages` with `enablement: true`, and the setting in the script too.** The
action can create the Pages site on the first run if the token may; where it may not,
the deploy fails loudly on the first push and the script has the same setting as a
reported-and-set block, `GET /pages` then `POST /pages` with `build_type: workflow`,
in the shape the private vulnerability reporting block has. Both roads lead to the
same state and the script is the record.

**Python is pinned by major, as Go and Node are.** `python-version: '3.14'`, the
interpreter the comparison ran under, resolved by `setup-python` to the newest patch,
which is D-5's reasoning for the Go directive: a patch reaches the pipeline without
anybody editing a file. The test gains `pythonVersion` and a `Python toolchain` entry in
`boundByHand`, because `3.14` has no shape `versionPin` reads.

**Zensical is pinned exactly and fetched from PyPI.** `pip install zensical==0.0.69`,
a row with its fetch named; `pip` resolves its fourteen dependencies unpinned, which the
row's paragraph says rather than hides, since a lockfile for a documentation job is the
cost the issue warned about for Astro and the decision not to carry one is a decision.

**The three links go to the host.** `docs/README.md` points at the three root files on
`github.com/triplem/xeno/blob/main/`, which is where a reader of the site finds them and
where a reader of the repository finds them too. Not a copy of the three into `docs/`:
`CLAUDE.md` is sent with every session and belongs where it is, and a copy is the second
friendlier version WP16 refuses.

**The job id is `docs`, and it becomes required.** Added to `scripts/github-settings.sh`'s
contexts and left to the maintainer to apply, as #341's were; until then the check runs
and reports and does not block, which is the honest intermediate state.

**The register.** The not-built paragraph loses the documentation site and says what
stands in its place; A106 records the three implementation decisions the plan left open:
strict on pull requests, deploy on the push alone, Python by major.

<!-- xeno:section:alternatives -->
## Alternatives

**A `docs/` subdirectory configuration, `docs/zensical.toml`.** Keeps the site's file
with the site's content. Declined: `renovate.json`, `.golangci.yml` and `.gosec.json`
are at the root and a reader looks there for a tool's policy; and Zensical's template
puts it at the root.

**Deploy from the pull request too, to a preview.** Pages has one site per repository and
no previews without a second deployment target; and a job that can publish must not run
on a fork's pull request. Declined; the strict build is the pull request's check.

**A lockfile for the Python environment.** `pip freeze` into a requirements file, so the
fourteen packages are pinned. Declined for now: the issue named a lockfile as the cost
that counted against Astro, Zensical's own pin is exact, and the row says the dependencies
resolve unpinned; the day a transitive package breaks the build is the day a lockfile
earns its place, and the build is strict enough to say so.

**Copying the three root files into `docs/`.** Makes the links internal. Declined, for
the reason in the decisions: a second copy of `CONTRIBUTING.md` is the drift WP16 exists
to prevent.

**Enabling Pages from this session.** Declined before trying: the host's branch
protection change was refused by the permission classifier yesterday and this is the
same class of action; the script carries it and `configure-pages` may do it on the
first run.

**Pinning Python to a patch.** `3.14.7`, the exact interpreter. Declined for D-5's
reason, applied to Go and Node already.

**A theme override or a logo.** Declined; nothing in WP16 asks for one, and a site that
is the documents needs no mark.

<!-- xeno:section:impact -->
## Impact

**For a pull request.** One more check, `docs`, red on a broken link or anchor anywhere
in the eleven documents, which nothing checked before. Required once the maintainer
applies the script; reporting from the first run.

**For the default branch.** Every merge publishes the site; the first deploy either
creates the Pages site or fails loudly until the maintainer runs the script's block.

**For the tree.** `zensical.toml`, `.github/workflows/docs.yml`, three links in
`docs/README.md`, six rows and a paragraph in `docs/supply-chain.md`, a regex and a
map entry in `supply_chain_test.go`, a block and a context in the settings script, a
line in `CONTRIBUTING.md`, a paragraph and a row in the register. No Go package moves.

**For the pin table.** Four action shas, a toolchain and a tool arrive, the largest
addition since the release pipeline; renovate's `github-actions` manager sees the shas,
and nothing proposes the next Zensical, which is a regex manager for another day.

**For the documents.** Nothing. They are published as they are, and the first reader of
the site meets the four documents and the pages about working with the tool.

**For WP16.** The documents part stands up, and every later part has a build to land in
and a strict check to fail; the ids for the seven documents without frontmatter are the
next step and a specification commit for two of them.
