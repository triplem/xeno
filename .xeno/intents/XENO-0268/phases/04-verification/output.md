---
intent: github.com/triplem/xeno#238
phase: 04-verification
created: "2026-10-07T08:51:28Z"
schema_version: "1.0"
runner_version: dev+9fd3647.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 94f1a8e73955042401ee08cb6dcd08f8e2a476de6e0836d77cc9f810b893420e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: verification@1.1.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
evidence:
    - kind: test-report
      result: pass
      produced_by: go test ./...
      sha256: 572a48ae635babe31ed639eb9936b17b6cb0d835d7f1573beb48edb377d2a260
      path: evidence/go-test.txt
      job: test
    - kind: build-log
      result: pass
      produced_by: go build, gofmt -l ., go vet ./..., xeno gate verify, git diff against main
      sha256: ce9fce2dd602524b10cddf113234ce781619d79b06adb59577a19ff8508609e3
      path: evidence/checks.txt
      job: checks
    - kind: other
      result: pass
      produced_by: the three verification answers with the pages each was read from
      sha256: e5f1669b9dc46c27b924c7cd82b892b569f8771aa5d090a365b036412ab2f0a8
      path: evidence/verification.txt
      job: verification
    - kind: other
      result: pass
      produced_by: the three passages and the register row, read in the files
      sha256: dd62abcbc3042f2a55cc545e6685829df201331d9708e3338a8229a83695c8df
      path: evidence/passages.txt
      job: passages
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

Twelve criteria, by number, all passing. Eight are read by a person, because the deliverable is
prose and a grep asserting a paragraph exists asserts nothing about whether it says anything.
Four are checks.

| # | What it asserts | What proves it |
|---|---|---|
| 1 | section 13 says the model invocation is deliberate, where the skills are enumerated | read in the file: `evidence/passages.txt` carries the paragraph and the one it follows |
| 2 | it states the skill-versus-command difference | read: the paragraph's second sentence |
| 3 | it says why a lens is the other way round | read: the third and fourth sentences |
| 4 | it names the mechanism and the client that lacks it | read: the fifth to seventh sentences, and `evidence/verification.txt` for what each rests on |
| 5 | it is worded as a fact a change deletes | read: the last sentence, and the absence of any permissive form in the paragraph |
| 6 | the budget paragraph says a prompt is not a tool for that purpose, and why | read: `evidence/passages.txt`, the second paragraph, against the budget clause printed beside it |
| 7 | it says what stays the client's choice | read: that paragraph's last sentence |
| 8 | the plan's section 9 carries all three answers in its own style | read: `evidence/passages.txt`, and the entry sits among the four that precede it in the same `-->` form |
| 9 | the six operations stay six | `evidence/passages.txt` prints the clause: "read intent, record assumption, write artifact, fetch template, run gates locally, and query the symbol index. Six" |
| 10 | no code changes and the suite passes | `evidence/checks.txt`: the diff against main is two documents and no Go file; `evidence/go-test.txt`; `gofmt`, `go vet` and `xeno gate verify` over 481 verdicts |
| 11 | the register carries one row with the dated fact | `evidence/passages.txt`: A99, five cells, state naming the decider and 2026-10-07 |
| 12 | the specification commit is its own and comes first | `git log`: `docs(spec): the phases stay skills, and a prompt is not a tool` is the only commit carrying a document change, and its message names the exception |

<!-- xeno:section:results -->
## Results

**The three answers, and what each rests on.** In `evidence/verification.txt`, each with the
page it came from rather than the conclusion alone.

Claude Code: the slash commands documentation lists, among the things a short name resolves
to, "An MCP prompt", linking the MCP page's section on prompts as commands; the name is
`/mcp__<server>__<prompt>` and the prompts are discovered from the connected server rather
than declared. The protocol's own prompts page agrees about the shape: prompts "are designed
to be **user-controlled**" and "would be triggered through user-initiated commands in the
user interface", with a slash command as its illustration.

Codex: `/mcp` inspects MCP tools and resources, `/mcp verbose` expands them, and nothing
exposes prompts. openai/codex#8342, "Expose MCP Server Prompts as Slash Commands Like Claude
Code", opened 2025-12-19, is closed as duplicate. The changelog's MCP work through 2026 is the
opt-in 2026-07-28 protocol revision in v0.147.0 with paginated tool discovery, the deprecation
and removal of `codex mcp-server`, and MCP request verification in v0.155.0. Checked
2026-10-07.

The cost: `prompts/list` returns `{ name, title?, description?, arguments? }` per prompt and
`prompts/get` returns `{ description?, messages[] }`. The messages are the prompt's text and
arrive only on invocation, so a prompt is not sent with every request the way a tool
definition is.

**The negative was checked the way the convention asks.** Codex's absence is the load bearing
answer, so the method was shown to find the thing where it exists: the same question asked of
Claude Code returns a documented mechanism and a naming pattern. Then the Codex issue's own
thread was read for its state rather than inferred from a search result, and the changelog was
read for what it does say about MCP rather than only for what it does not.

**The three passages, read in the files.** `evidence/passages.txt` prints each one whole, next
to the paragraph it follows, because the deliverable is prose and eight of the twelve criteria
are a person reading it. The budget clause is printed beside the new paragraph so that the two
can be read as one argument, and the clause enumerating the six operations is printed to show
it is the same six.

**The register row.** A99, five cells, state `approved (#238)` naming the maintainer and
2026-10-07, immediately after A97 — A98 is on the unmerged branch for #228, and P3's
deviations say why the number was left to it.

**Wrapping.** No added prose line exceeds 88 characters. Both paragraphs went in two lines
over the width on the first attempt and were replaced as whole paragraphs rather than edited
line by line, which is what `CLAUDE.md` asks for a paragraph being changed a second time, and
read back as paragraphs afterwards.

**The checks.** `go test ./...` passes across 20 packages, `internal/secrets` the slow one at
132.5 s and no failures. `go build`, `gofmt -l .` outside `vendor/` and `go vet ./...` are
clean. `xeno gate verify` recomputes and matches 481 verdicts. The diff against `main` is two
files, both documents, and no Go file — which is the fastest way to see that criterion 10
holds.

<!-- xeno:section:gaps -->
## Gaps

**The load bearing fact is a negative about somebody else's software, and it is dated.** Codex
not surfacing MCP prompts as commands is what makes #238's first branch unavailable, and it is
checked by reading that project's own issue tracker and changelog on 2026-10-07. Three things
could make it wrong: a release between the check and the merge, a feature shipped without a
changelog line, and a configuration or flag that enables it and is documented somewhere neither
read. The positive control narrows the first risk and not the other two — it shows the method
finds the feature in a client that has it, which is not the same as showing it would find a
feature documented nowhere.

**Nothing in this repository would notice if it changed.** The paragraph is worded so that a
change makes it false rather than satisfied, which puts the correction in a diff when somebody
comes to make it. It does not put it anywhere when nobody does. No check watches another
project's changelog, nothing can, and the register row says the date rather than pretending to
a watch.

**The third answer is only half measured.** The protocol fixes where a prompt's text lives:
`prompts/list` carries a name, a title, a description and the arguments, `prompts/get` carries
the messages. It does not fix what a client keeps in context from the listing, so the claim
that prompts are cheaper than tools holds for the text and is unverified for the names and
descriptions. Both paragraphs say so. If a client did place every prompt's listing in every
request, six phase prompts would be six standing costs after all, smaller than six tool
definitions and not zero.

**Eight of twelve criteria are read by a person.** The deliverable is prose, and a check that
a paragraph exists says nothing about whether it says the thing. `evidence/passages.txt` prints
each passage so that a reviewer reads what was written rather than a claim about it, which is
the most this phase can offer and is not the same as a check.

**The MCP server is not built, so nothing here is tested against the thing it decides about.**
If WP11 builds the server and the prompts question is reopened, the decision will be read
against a surface that exists; today it is read against the protocol and two clients'
documentation. That is the right way round for a decision not to build something, and it is
still a decision taken without the artifact in hand.

**A98 is missing from the register on this branch.** It is on the unmerged branch for #228, so
a reader of this branch alone finds A97 followed by A99 and no explanation. P3's deviations say
why. If #228 is abandoned the gap stays, and nothing fills it.

**Nothing verifies that `/mcp__server__prompt` is still the name.** It was read from Claude
Code's documentation, which is the right source and the same kind of dated fact as the Codex
answer. The plan's section 9 carries the name because a reader rechecking needs it; if the
pattern changes the entry is wrong in a way nobody here would catch.
