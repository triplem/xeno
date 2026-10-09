---
intent: github.com/triplem/xeno#226
phase: 04-verification
created: "2026-10-09T13:05:25Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a43435785b3bc3ba1758608e95609e0671ce7c9d3c17e76412cad1aae1edc09e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| criterion | how it was checked | result |
|---|---|---|
| 1. the plan names the generator before the code | `git log` on the branch: `5056d9b docs(plan): WP16 builds the site with Zensical, not MkDocs` is the first commit after `main`; the paragraph read back as a paragraph, revision 10 | met |
| 2. the site builds strictly on the branch | `zensical build --clean --strict` with 0.0.69: "No issues found", 0.8 s, exit 0; `git status --short` shows nothing under `.xeno/local/site` | met |
| 3. a broken anchor fails the build | scratch copy of `docs/` and `zensical.toml` with `[here](process-definition.md#no-such-anchor)` appended to `commands.md`: "anchor does not exist" at `commands.md:98:49`, "Aborted because --strict flag is set", exit 1 | met |
| 4. light and dark, with a toggle | `tomllib` reads three palette entries; the built `index.html` carries `data-md-color-scheme="default"` and `="slate"` and four toggle option elements | met, with the count P3's deviation explains |
| 5. the documents are published unrewritten | `git diff main -- docs/`: `implementation-plan.md` the paragraph and revision, `README.md` the three links, `assumptions.md` the paragraph and A106; the other three normative documents untouched | met |
| 6. the three links resolve | the strict build reports nothing about `README.md`; the three hrefs are `github.com/triplem/xeno/blob/main/…` | met |
| 7. the job has the scanners' pins and the Pages shape | `docs.yml` read: two jobs, the five `uses:` lines by sha with tag, `zensical==0.0.69`, `--strict`, `deploy` with `if: github.event_name == 'push'`, `needs: docs`, `pages: write`, `id-token: write`; `docs` under the workflow's `contents: read` | met |
| 8. the pin table is held both ways | `go test ./internal/model/` green with the six rows and `Python toolchain` bound by hand; both directions pass | met |
| 9. the host settings are in the script | `sh -n` clean; the `pages` block and `docs` in `contexts`; neither applied from this session | met in the tree; the host half is the maintainer's, as the pull request says |
| 10. the register says what moved | `grep` finds no "documentation site" in the not-built paragraph; A106 present | met |
| 11. the suite and the gates stay green | `go test ./...` exit 0, `gofmt -l` empty, `go vet` clean, `xeno gate verify` 558 verdicts | met |
| 12. a reader is told | `CONTRIBUTING.md` names `docs` and what it is red on | met |

<!-- xeno:section:results -->
## Results

Twelve criteria, twelve met, one with a count that differs from the design and is
explained there.

    zensical build --clean --strict     No issues found, 0.8 s, exit 0
    broken anchor on a scratch copy     anchor does not exist, commands.md:98:49, exit 1
    go test ./...                       ok, exit 0
    gofmt -l . | grep -v vendor/        nothing
    go vet ./...                        clean
    ./xeno gate verify                  verified 558 verdicts
    sh -n scripts/github-settings.sh    clean
    docs.yml                            parses, jobs docs and deploy
    zensical.toml                       parses, three palette entries

**The strict build is the check, and it was exercised both ways.** On the branch it
passes, which it did not on `main`'s `docs/README.md` until the three links moved; on a
scratch copy with one broken anchor it aborts and names the file, the line and the
column. That is criterion 3 and the property WP16 chose the generator for.

**Light and dark are in the output, not only in the configuration.** Both scheme
attributes and the toggle markup are in the built index; what the toggle looks like
is a thing to open the site for, which `zensical serve` does locally and the first
deploy does for everybody.

**Timings, for #117.** `gate run` was not timed on this intent; the figure from
yesterday's scratch run, 0.118 s, stands. `gate verify` over 558 verdicts: 36 s, the
same per-verdict cost as yesterday's 552.

<!-- xeno:section:gaps -->
## Gaps

**The workflow has not run on the runner.** Every command in `docs` ran locally with the
same versions; `setup-python` resolving `3.14`, `pip` on the runner, the artifact upload
and the deploy are the parts a local run does not touch. The pull request's own `docs`
check is the first run of the first job; the first push to `main` is the first run of
`deploy`, and the Pages site does not exist yet, so that run either creates it through
`enablement: true` or fails loudly until the maintainer runs the script's block.

**`docs` is not required until the maintainer applies the script.** It reports on this
pull request and blocks nothing; a broken link merged before the script runs would
publish a 404. The pull request names the step.

**The unpinned dependencies are a claim about today.** Fourteen packages resolved on
2026-10-09 built the site; the row says they are unpinned and the build is strict, and
nothing here proves the next resolution builds the same site.

**Mermaid and snippets are configured and never exercised.** No document carries either,
so the two extensions are a declaration until the first diagram.

**The seven documents without frontmatter have no id**, which WP16 asks for; the site
builds without them and the non-goals say why they follow rather than lead.

**The comparison's evidence is gone.** The three branches were deleted on instruction;
the comment on #226 and this intent's `zensical.toml` are what remains, which the intake's
learning records.
