---
intent: github.com/triplem/xeno#95
phase: 01-requirements
created: "2026-10-09T16:47:20Z"
schema_version: "1.0"
runner_version: dev+d2bc411.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 96b20eddde2429059cf00105ca3f91d1228819255f06e8371bbf7b4986d6a730
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
template: requirements@1.1.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. **Every pin `docs/supply-chain.md` records as fetched and movable is reached by a
   manager.** The pipeline table holds thirty rows and the module table one. They account
   for as follows, and the accounting is the criterion rather than a count:

   | | rows | reached by |
   |---|---|---|
   | actions pinned by sha | 13 | `github-actions` |
   | the one module, vendored | 1 | `gomod` |
   | the `go` directive | 1 | `gomod`, minor by D-5 |
   | tool versions inside workflow text | 6 | a `customManager` each, today |
   | **the image's base, `gitleaks`, `zensical`** | **3** | **a `customManager` each, after this change** |
   | Node and Python, a chosen major resolved by `setup-*` | 2 | nothing, deliberately |
   | gitleaks' rules, translated into `secrets.yaml` and not fetched | 1 | nothing, deliberately |
   | not pinnable at all | 4 | nothing, and the page says why |

   The four that cannot be pinned are trivy's and govulncheck's vulnerability databases,
   the packages installed into the image, and the docker that builds it. After this change
   no row is unreached except the seven the page itself explains.

2. **Every `customManager` resolves to the version the tree actually carries.** Not "the
   regex is plausible" — the match is run against the files in scope and each entry's
   resolved value is recorded beside the value in the tree. Nine entries, nine values, and
   the six that exist today are included rather than assumed, because an entry that matches
   nothing is indistinguishable from an entry that was never run.

3. **The base image reference names the tag the page names.** `docs/supply-chain.md` calls
   that pin `debian` 13-slim and the `Dockerfile` carries a bare digest, so the page names a
   stream the tree does not. After this change the reference carries both, and a reader of
   either finds the same thing.

4. **`internal/model/supply_chain_test.go` still holds the page against the tree in both
   directions.** The `Dockerfile` edit inserts a tag before `@sha256:`, which the digest
   pattern either tolerates or does not. The test passes, or the change is wrong.

5. **No commit footer points at a closed issue.** `commitBody` carries no `Refs #95` once
   this intent closes #95.

6. **Nothing about the decided shape moves.** `automerge` stays false everywhere, the
   dashboard stays on beside pull requests, `postUpgradeTasks` still runs `go mod vendor`,
   and no proposal gains a `Signed-off-by`. Those are the answers given on #95 and Q-1 of
   XENO-0277, and this intent is not entitled to revisit them.

7. **The whole gate suite passes on the change.** `go build`, `go test ./...`, `gofmt -l`
   empty, `go vet`, and `xeno gate verify` at exit 0, run on the tree as it stands with the
   change applied rather than on a tree it was reverted out of.

<!-- xeno:section:non-goals -->
## Non goals

**Switching renovate on.** `RENOVATE_TOKEN` is a secret only the maintainer can add, and
#350 carries it with what to check on the first run. D-1 is that #95 does not wait for it.

**Enabling the `dockerfile` and `pip` managers.** Both would be the obvious way to reach two
of the three pins and neither fits: the base image is the value of `ARG BASE` rather than of
`FROM`, and `zensical` is an argument to `pip install` in a workflow rather than a line of a
requirements file. A manager that resolves nothing is worse than none, because its presence
reads as coverage. `enabledManagers` stays as it is and the three pins are reached by regex,
which is what the rest of this configuration already does.

**Pinning what the page says cannot be pinned.** Two vulnerability databases, the packages
apt installs into the image, and the runner's docker. Each has a row saying so and a reason.

**Moving any pin.** This intent makes pins visible to a bot. It bumps nothing: every version
in the tree after it is the version that was there before, which is also what makes the
verification readable — nine managers resolving to nine unchanged values.

**Proposing the Node or Python major.** `node-version: '24'` and `python-version: '3.14'`
are a choice this project made, resolved to the newest of the line by the setup actions. A
manager offering 25 or 3.15 would be proposing a decision dressed as an update.

**Answering whether the gosec pseudo-version should go back to a release.** #343 pinned a
commit because release 2.29.0 cannot read Go 1.27.2's export data, and the existing manager
will propose 2.30.0 when it exists. That proposal is the mechanism working, and whether to
take it belongs to whoever reads it.

<!-- xeno:section:constraints -->
## Constraints

**`docs/supply-chain.md` is held against the tree in both directions.** #317 added the
second direction and `internal/model/supply_chain_test.go` runs both, so a pin that moves
without its row, or a row without its pin, is a red `verify` rather than a stale page. The
`Dockerfile` edit is inside that check.

**A gitleaks proposal cannot be complete, and that is accepted rather than solved.**
Renovate can move `GITLEAKS_VERSION` and cannot recompute `GITLEAKS_SHA256` beside it, so
`sha256sum -c` in the install step fails until a person recomputes the checksum. This is
the bargain the configuration already struck for `github-actions`, where a bumped sha needs
a Markdown row moved by hand and the existing `packageRules` entry says so in as many
words. Nothing automerges, so the failure is loud and waits for somebody; the alternative
is a pin that goes stale with nothing to notice, which is what this intent exists to fix.

**The decided shape is not this intent's to revisit.** Self-hosted, no automerge, dashboard
plus pull requests, and a person re-signs each proposal. Three of those are answers in #95's
comments and the fourth is Q-1 of XENO-0277. A change here that moved any of them would be
XENO-0274's failure again, from the other direction.

**`commitBody` is not where a sign-off goes.** Dropping it must not read as reopening the
DCO question. The key carried `Refs #95` and never a `Signed-off-by`; Q-1's answer is why
there is no sign-off, and the footer is a separate thing that happens to live in the same
key.

**The regex dialect is renovate's, not Go's.** `managerFilePatterns` are `/…/`-delimited
and `matchStrings` use JavaScript named groups, so a regex is checked by running it the way
renovate runs it rather than by compiling it in a Go test. There is no renovate binary in
this repository and adding one would be a second dependency, which this project treats as a
decision; so the check is a script run in the verification phase and recorded as evidence,
not a test in the suite.
