---
intent: github.com/triplem/xeno#226
phase: 05-review
created: "2026-10-09T13:05:25Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a54bec71ce63f8383e26504597c6e1920057bcbe28663af92f35c187fe8aefad
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: met
      note: 'One deviation in P3, naming the design''s palette decision: three entries for two schemes, because the system-preference state is an entry of its own.'
    - rule: interface-change-needs-a-migration-note
      result: met
      note: The set of checks gains docs, which reports from the next push and becomes required when the maintainer applies scripts/github-settings.sh, named on the pull request; nothing an adopter uses changes.
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: No module enters go.mod; a Python tool and four actions enter the pipeline as rows of the pin table, with the tool's unpinned dependencies given a paragraph there.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**deviations-are-traceable — met.** One deviation in P3, naming the design's palette
decision and explaining the count: three entries for two schemes, because the
system-preference state is an entry of its own in the theme's shape. Nothing decided was
undone.

**interface-change-needs-a-migration-note — met.** The interface that moves is the set of
checks a pull request reports: `docs` appears, red on a broken cross reference between
the documents. A pull request open across the merge gets the check on its next push and
needs nothing else; it becomes required when the maintainer applies the settings script,
which the pull request names, and until then it reports without blocking. No command or
file an adopter uses changes.

**new-dependency-needs-a-rationale — not-applicable.** No module enters `go.mod`. A
Python tool and four actions enter the pipeline, each a row of the pin table with its
fetch named, and the tool's unpinned dependencies have a paragraph of their own there.

**What no rule asks.** Whether the decision to leave fourteen packages unpinned in a
documentation job is the right side of the line the supply chain page draws: A106 and
the page both say it is tolerable because nothing a verdict rests on comes out of that
job, and both name the day it stops being tolerable.

<!-- xeno:section:release-notes -->
## Release notes

**The documentation is a site, built with Zensical and published to Pages.** The eleven
documents under `docs/` are rendered as they are, in a light and a dark mode that follow
the reader's system preference with a toggle, from `zensical.toml` at the root. Every
pull request builds the site strictly, so a link to a page that does not exist or an
anchor that does not exist is a red `docs` check before a merge; every push to `main`
deploys what it built.

**Why Zensical.** WP16 named MkDocs with the Material theme; #226 built a page with each
of three generators and found that Material's own authors say to leave MkDocs, whose 2.0
removes the plugin system, and build Zensical as its replacement with the same
configuration keys and extensions. Astro Starlight needed three things assembled before
the documents would publish. The plan's paragraph now says so, in a commit before this
code.

**What it costs the pipeline.** A Python toolchain pinned by major, Zensical pinned to
0.0.69 with its fourteen dependencies resolving unpinned, and four of GitHub's own
actions pinned by sha, all on the pin table, which its test holds both ways.

**For the maintainer, once.** `scripts/github-settings.sh` now creates the Pages site
where it is absent and requires `docs`; run it after the merge, or let the first deploy
try to create the site itself.

**For an adopter.** Nothing. This is this repository's documentation.

<!-- xeno:section:residual-risk -->
## Residual risk

In order of weight.

**The first deploy has not happened.** The Pages site does not exist; the first push to
`main` either creates it through the action or fails until the maintainer runs the
script's block. A failed `deploy` publishes nothing and breaks nothing, and the `docs`
check on this pull request is the first time the build runs on a runner at all.

**`docs` reports and does not block until the script runs.** A broken cross reference
merged in that window publishes a 404 that the next pull request's check would have
caught. The window is as long as the maintainer leaves it.

**Fourteen packages resolve unpinned.** A documentation build that worked today may not
tomorrow, and the strict build will say so loudly rather than publish something else;
the page names the lockfile as the remedy and the day it earns its place.

**Zensical is below version one.** Its configuration is MkDocs's keys in TOML and the
move back is an afternoon, which the plan paragraph says; what it does not say is what
a breaking change in a 0.0.x release looks like on the day it arrives, which is a red
`docs` check and a version pin to move.

**The site publishes the trail's absence.** `.xeno/intents/` is in the repository and not
on the site, which #234 is about and recommends publishing the digests; the first
reader to notice will ask, and the answer is on that issue.

**Belongs to the specification, not to this code.** Ids for the seven documents without
frontmatter, which WP16 asks for and two of which are a specification commit; and
whether the pull request's uploaded artifact should be an evidence item a phase can
declare, which is section 4's shape and nobody's intent.
