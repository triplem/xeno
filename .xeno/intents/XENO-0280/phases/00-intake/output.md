---
intent: github.com/triplem/xeno#226
phase: 00-intake
created: "2026-10-09T12:56:02Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 028c682e00440bc21cbd179f1ff5d1b7be3cddebd9b9add5ff74637bd2a32de1
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Intake

<!-- xeno:section:problem -->
## Problem

Approved by @triplem on 2026-10-09T07:02:00Z: Use zensical, if possible with a light and a dark mode

> **github.com/triplem/xeno#226** — Is MkDocs still the right generator against Astro and Zensical?
>
> The question, as asked: would Astro or Zensical be a better fit than MkDocs for the
> documentation, since they look better?
>
> ## What the plan committed to, and why
>
> WP16 does not leave the generator open. It names **MkDocs with the Material theme**, puts
> the choice in that work package deliberately rather than in a later decision, and gives
> three reasons: per page frontmatter, anchor and link checking, and Mermaid are all there
> without assembling them; the source hash binding and the extraction of code examples both
> depend on the generator; and this is "the one place where a Python dependency in a
> documentation job costs nothing that matters". The alternative it names is Hugo — a single
> static binary, closer to the rest of this project — rejected because its checking would
> have to be put together by hand.
>
> So the first standing rule applies: this is a change to a normative document, and it is the
> maintainer's commit, made before any code follows from it.
>
> ## What is built so far
>
> Nothing. There is no `mkdocs.yml` in the repository and no documentation job. WP16 is
> unstarted, which is the cheapest possible moment to revisit the choice — the decision has
> cost one paragraph so far and no pipeline.
>
> ## What the comparison has to answer
>
> Appearance is a real criterion and it is not the one WP16 chose on, so a replacement has to
> be measured against what it chose on:
>
> | what WP16 depends on | the question for a candidate |
> |---|---|
> | anchor and link checking out of the box | does it fail the build on a broken cross reference, or does a plugin have to be found? |
> | per page frontmatter with a stable id and revision | can a cross reference point at an id rather than a path? WP16 says renaming a file has already broken links once |
> | Mermaid, from sources in the repository | rendered without a browser in the pipeline? the plan forbids committed images |
> | generated pages | the gate pages, the reference and the limitations page are generated from declarations; how does the generator take generated Markdown? |
> | extraction of code examples | the worked intent is the WP17 golden fixture, shown rather than retyped |
> | published from CI to Pages on every merge | one job, no runtime in the pipeline beyond the build |
> | the four documents published as they are | 2 284 and 2 009 line Markdown files with tables and deep cross references, rendered unrewritten |
>
> Two notes on the candidates before anybody measures them. Astro is a JavaScript toolchain,
> so adopting it trades WP16's Python dependency for a node dependency and a lockfile in a
> repository whose entire supply chain story is one vendored Go dependency; `SUPPLY-CHAIN.md`
> is where that cost lands, not in the documentation job. Zensical is the newer project from
> the Material for MkDocs authors, which makes it the most interesting candidate and also the
> one whose maturity has to be established rather than assumed — the state of its link
> checking, its frontmatter handling and its Mermaid support on the day somebody looks is the
> whole of the answer, and this issue should not pretend to know it.
>
> ## Done when
>
> The table above is filled in for MkDocs, Astro and Zensical by somebody who has built a
> page in each, and `docs/implementation-plan.md`'s WP16 paragraph either stays as it is with
> the comparison recorded beside it, or is changed in a commit of its own. A generator chosen
> for how it looks and discovered later to have no link checking would cost WP16 its
> "consistency is mechanical where it can be" section, which is the part of that package that
> does the work.
>

The issue as it stood at 2026-10-09T12:56:02Z, read by `xeno phase start` and quoted rather than summarised. What this phase concludes about it belongs below.

The issue as `xeno phase start` read it, with the approval sentence above it: the label
and the comment were both there, read from the host. The decision the comment records
was reached on the issue itself, where three branches had built a page with each
generator on 2026-10-08 and the table the issue asked for was filled in from what was
built: MkDocs Material and Zensical both check links and anchors out of the box and
rendered the eleven documents unrewritten; Astro Starlight needed three things assembled
and a link-rewriting plugin before the documents would publish; and Material for MkDocs
prints on every build that MkDocs 2.0 removes the plugin system it depends on, that its
authors expect no further 1.x updates, and that they build Zensical as the replacement.
The maintainer chose Zensical and asked for a light and a dark mode. The three
comparison branches and their worktrees were deleted on the same instruction, and the
comment on the issue is their record.

**The specification moved first.** WP16's tooling paragraph in
`docs/implementation-plan.md` named MkDocs with the Material theme and now names
Zensical, with the reason, the two alternatives measured and the light and dark mode,
in commit `5056d9b` on this branch, revision 9 to 10, written by the agent on the
maintainer's instruction and recorded here so that the exception to the first standing
rule stays one.

**What is in the tree today.** Nothing: no site configuration, no documentation job, and
`docs/assumptions.md` lists the documentation site as not built. WP16 says the
documents part can be built before anything else in the package and should be, because
the four documents already cross reference each other by section and by anchor and
nothing checks those references. The strict build of the comparison found exactly three
broken links across the eleven documents, all of them `docs/README.md` pointing at
`CONTRIBUTING.md`, `SECURITY.md` and `CLAUDE.md` in the repository root, which the site
has no page for.

**What the measurement fixed.** Zensical 0.0.69 builds the eleven documents in 0.8 s
with fourteen packages behind it, reports a missing page or anchor with line and column,
and `--strict` aborts on either; its configuration is a TOML of the keys `mkdocs.yml`
has; a Pages deployment needs three actions of GitHub's own and a Python toolchain, each
a row of the pin table; and the repository has no Pages site yet, which is a host
setting.

<!-- xeno:section:scope -->
## Scope

This intent builds the documents part of WP16 with Zensical, as the plan now says, and
publishes it from CI to Pages on every merge to the default branch.

- **The configuration.** `zensical.toml` at the root: the eleven documents under `docs/`
  in a navigation of three groups, the Material theme with a light and a dark palette
  and a toggle, permalinks on headings, the Mermaid superfence and the snippets extension
  with path checking, and the site built into `.xeno/local/site`, which is gitignored.
- **The job.** `.github/workflows/docs.yml`: a `docs` job that installs Zensical at an
  exact version and builds with `--strict` on every pull request and on a push to `main`,
  so a broken link or anchor fails a pull request; and a `deploy` job that runs only on
  the push to `main`, uploads the built site and deploys it to Pages through GitHub's own
  three actions, each pinned by sha.
- **The three links.** `docs/README.md`'s links to `CONTRIBUTING.md`, `SECURITY.md` and
  `CLAUDE.md` point at the files on the host, since the site has no page for them and a
  strict build refuses a link to a page that does not exist.
- **The pin table and its test.** Rows for the three Pages actions, `actions/setup-python`,
  the Python toolchain by major as Go and Node are, and Zensical's version;
  `supply_chain_test.go` learns to bind the Python row by hand as it binds the other two
  toolchains.
- **The host.** `scripts/github-settings.sh` gains the Pages site, built from a workflow,
  reported and set where it is absent, and `docs` joins the required checks, so that a
  broken cross reference stops a merge. Both are the maintainer's to apply; the pull
  request names them.
- **The register.** `docs/assumptions.md` stops listing the site as not built and gains a
  row for what the plan left to the implementation: strict on pull requests, deploy on
  the push, the Python toolchain pinned by major.

What it does not do.

**It does not build the rest of WP16.** The generated reference, the gate pages, the
limitations page, the worked intent and the hand-written parts are each their own work;
this is the part the plan says to stand up first.

**It does not give the seven documents without frontmatter an id.** WP16 asks for a
stable id and a revision per document; four have them and seven do not, and adding them
to the two normative documents is a specification commit. Zensical does not need them to
build, so the site goes up first and the ids follow.

**It does not add Mermaid diagrams or snippets.** Both are configured and neither is
used, because no document carries a diagram or an extracted example yet.

**It does not check the links in the four documents by id.** They point at files and
anchors, which the strict build checks; pointing them at ids is the hash binding WP16
puts with the first translated page.

<!-- xeno:section:context-rationale -->
## Why this context

Read for this phase, and why.

- Issue #226 with its comments, which hold the question, the filled-in table, the
  recommendation and the maintainer's answer; and `docs/implementation-plan.md` WP16,
  the paragraph that changed and the paragraphs around it on frontmatter, publishing from
  CI and mechanical consistency, which bound what the site has to do.
- `docs/README.md`, the one document whose links the strict build refused.
- `docs/supply-chain.md` and `internal/model/supply_chain_test.go`, because four actions,
  a toolchain and a tool arrive and each is a row the test holds to the tree in both
  directions, and the toolchain row needs the by-hand binding the Go and Node rows have.
- `scripts/github-settings.sh`, for the shape of a host setting reported and set, and the
  required check list; `CONTRIBUTING.md`, for the list of checks a reader is told about.
- `.github/workflows/trivy.yml` and `lint.yml`, for the pinning and the job id
  conventions a new workflow follows; `.gitignore`, for where a built site may land.
- `docs/assumptions.md`, for the not-built paragraph and the shape of a row.
- Zensical 0.0.69's own `zensical new` template, its `--help` and the comparison branch's
  configuration, read on the day rather than remembered.
