---
intent: github.com/triplem/xeno#226
phase: 01-requirements
created: "2026-10-09T12:58:16Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a70db35a7e681b2cacb18b2caa41b2e234a266c1f8fe3ff63695d96fa354dada
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

Each criterion is a state of the tree and a verdict over it, read on the branch.

1. **The plan names the generator before the code does.** `docs/implementation-plan.md`'s
   WP16 tooling paragraph names Zensical, the reason, the two alternatives measured and
   the light and dark mode, at revision 10, in a commit that precedes every other commit
   of this branch.

2. **The site builds strictly on the branch.** `zensical.toml` at the root; `zensical
   build --strict` with Zensical 0.0.69 exits 0 and writes the eleven documents under
   `.xeno/local/site`, which `git status` does not show.

3. **A broken anchor fails the build.** A deliberately broken anchor added on a scratch
   copy makes `zensical build --strict` exit 1 and name the file, line and anchor.

4. **Light and dark, with a toggle.** `zensical.toml` carries two palettes, `default` and
   `slate`, each with a toggle; the built `index.html` contains both scheme attributes
   and the toggle markup.

5. **The documents are published unrewritten.** The diff of `docs/` on the branch is the
   plan paragraph, the three links in `docs/README.md` and `docs/assumptions.md`; the
   four normative documents are byte-identical to `main` apart from the plan paragraph
   and its revision.

6. **The three links resolve.** `docs/README.md` links to `CONTRIBUTING.md`, `SECURITY.md`
   and `CLAUDE.md` on the host, and the strict build reports nothing about them.

7. **The job has the scanners' pins and the Pages shape.** `.github/workflows/docs.yml`,
   job id `docs`, runs on every pull request and on a push to `main`, pins
   `actions/checkout`, `actions/setup-python`, `actions/configure-pages`,
   `actions/upload-pages-artifact` and `actions/deploy-pages` by sha with the tag beside,
   installs `zensical==0.0.69`, builds with `--strict`; a `deploy` job runs only on the
   push to `main`, needs `docs`, and holds `pages: write` and `id-token: write` where
   `docs` holds `contents: read`.

8. **The pin table is held to the tree in both directions.** `go test ./internal/model/`
   is green with the four new action rows, the Zensical row and a Python toolchain row
   bound by hand from `python-version` in `docs.yml`, as the Go and Node rows are.

9. **The host settings are in the script.** `scripts/github-settings.sh` reports the Pages
   site and creates it with `build_type: workflow` where it is absent, and lists `docs`
   among the required checks; the pull request names both as the maintainer's to apply.

10. **The register says what moved.** `docs/assumptions.md` no longer lists the site as
    not built and carries A106 for strict on pull requests, deploy on the push, and the
    Python toolchain pinned by major.

11. **The suite and the gates stay green.** `go test ./...`, `gofmt -l`, `go vet ./...`
    and `./xeno gate verify` exit 0 on the branch.

12. **A reader is told.** `CONTRIBUTING.md` names `docs` among the checks that gate a
    merge.

<!-- xeno:section:non-goals -->
## Non goals

- **The rest of WP16.** Reference pages, gate pages, the limitations page, the worked
  intent, onboarding, adaptation and removal are each their own intent; the plan says
  the documents part goes up first and this is it.
- **Ids for the seven documents without frontmatter.** Two of the eleven are normative
  and a frontmatter block on them is a specification commit; Zensical builds without
  them, so the ids follow the site.
- **Diagrams and snippets.** Configured, unused: no document carries either yet.
- **The source hash binding and translation.** WP16 puts both with the first translated
  page.
- **Enabling Pages or changing the required checks from this session.** Both are host
  settings the maintainer applies; the script carries them and the pull request says so.
- **A custom domain, analytics, or a theme override.** Nothing WP16 asks for.

<!-- xeno:section:constraints -->
## Constraints

- The first standing rule: the plan paragraph is the specification and precedes the code;
  nothing here reads it more loosely.
- The four documents are published as they are; no line of them changes for the site.
- Every action pinned by sha with the tag beside it, every tool to an exact version, every
  pin a row of `docs/supply-chain.md`, held by the test in both directions.
- The job id is the check name; a required context that never reports blocks every merge,
  so `docs` is added to the script and applied by the maintainer, as #341 was.
- The built site lands under `.xeno/local/`, which is gitignored, for the reason A39
  gives; nothing generated is committed.
- One vendored Go dependency; this adds a Python tool to the pipeline and nothing to the
  binary.
- The deploy job alone holds the Pages permissions, and runs only on a push to `main`:
  a pull request from a fork must never be able to publish.
- `CLAUDE.md`: prose at 88, a paragraph changed twice is replaced, a comment says what
  the construction is and why.
