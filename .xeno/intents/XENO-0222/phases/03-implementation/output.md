---
intent: github.com/triplem/xeno#169
phase: 03-implementation
created: "2026-10-03T09:08:50Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+93322c5.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 64bc330e66fca0dbd671b9b4869976c8e323ab98caa3572b1b1a3289a6f24883
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

Two manifests, seven skills, one hook file, one vendoring change, one register row, and a test
package that checks prose against the tree it describes.

**`.claude-plugin/marketplace.json`**, in the repository root, naming one plugin whose source is
`./.xeno/plugin`. It carries a description because the client warns without one.

**`.xeno/plugin/.claude-plugin/plugin.json`**: name, description, version `0.1.0`, author,
homepage, licence and four keywords. The version is the shared one section 13 asks for.

**`.xeno/plugin/hooks/hooks.json`**: the Stop event wired to `xeno cost turn`, which is how a
project that installs the plugin records a turn's cost without writing a settings file. It says
`xeno` rather than `./xeno`, because a project receives the runner as a platform binary; this
repository keeps its own `.claude/settings.json` entry, which runs the binary it just built.

**Seven skills under `.xeno/plugin/skills/`**, named as section 13 names them, each the same five
parts: what the phase is for in two sentences, the sections it owes with what each one is, the
commands in the order they are run, what its gate refuses gate by gate, and what it hands to the
next phase.

`xeno-intake` has the three sections and the five gates that apply from P0. `xeno-requirements`
adds the assumption commands and says that an open assumption blocks the phase, which is the
point. `xeno-design` says to cite a precedent in the tree rather than argue from first principles
for something already settled, and to write out the reading the specification supports and the
phase did not take. `xeno-implementation` says a deviation names what it departed from and that a
deviation discovered while implementing is the most valuable thing the phase produces.
`xeno-verification` says to run the thing against the real repository on a throwaway copy where
that is possible, because fixtures carry the current shape of everything and the tree carries
every earlier shape — and it names `G-Test` with the fact that it reports `not-implemented` in
this release, so an agent is not told to expect a verdict it will never see.
`xeno-review` has the checklist in the frontmatter as well as the section, the three results, the
note rule, the lens entry that answers no rule, and the two commands a second person uses.
`xeno-learning` is the record's shape, the four categories, `no_finding: true` as the honest empty
case, what makes an entry worth writing, and that a learning goes nowhere automatically because
section 10 routes it through a merge request.

**`internal/runner/init.go`**: `vendorTree(res, "skills")` beside the templates, and `vendorFile`
for `.claude-plugin/plugin.json` and `hooks/hooks.json`. An absent file is a plugin released
before that file existed, which is every release before this one, so it is not an error. No
`mcp.json`, because a declaration of a server that does not exist is a client failing at startup.

**`internal/plugin`**, a package with no code and ten tests, because three things in a skill can
be wrong mechanically. The seven names are the seven of section 13, no more and no fewer. Every
skill names itself in its frontmatter and carries a description of at least twelve words that says
when to use it. Every command a skill names exists in the dispatch table, read out of `main.go`.
Every section a phase skill asks for is a section of that phase's template, read through
`template.Load`. Every gate a skill names is in the table, and a skill naming one of the three
that report `not-implemented` has to say so. No skill names this project's own world, from a list
of seven words. No skill mentions a server that does not exist. The hook wiring is the Stop event
and the command on the path. And both manifests are read rather than described, because the client
loads them and a field a test invented would pass while the client refused the plugin.

**`ASSUMPTIONS.md`, A77**: where the manifests sit, what the client requires, what section 13
draws instead, and that the correction to the document is a person's commit.

**The skills are wrapped to 88 columns** like every other piece of prose here, with the
frontmatter description left on one line because the format requires it.

<!-- xeno:section:deviations -->
## Deviations from the design

**The first version of the verification skill promised a verdict nobody will see.** It named
`G-Test` as what the gate refuses, and `G-Test` reports `not-implemented`. The test I had written
for exactly this case exempted the verification skill by name, which is a test shaped around the
defect rather than against it. Both are corrected: the skill says the gate is unimplemented in
this release and why the mapping is owed anyway, and the check now allows a skill to name an
unimplemented gate only if it says so in the same file. Recorded because the exemption was mine
and it would have read as deliberate.

**The skills were written at the wrong width and rewrapped by a script.** I wrote them by eye at
about ninety-five characters, and the convention is eighty-eight. Seventy-six prose lines were
over. A rewrapping script in the scratchpad fixed them, preserving the frontmatter, the code
blocks and the list indents. Worth recording because the first count said eighty-four lines and
the real number was seventy-six: `awk length` counts bytes and an em dash is three of them, so the
measurement needed correcting before the prose did.

**`claude plugin marketplace add .` is refused and `./` is accepted.** The client wants
`owner/repo`, a URL, or a path beginning with `./`. A bare dot is "Invalid marketplace source
format". That is a detail of the client rather than of this piece, and it belongs in whatever
WP16 writes for an adopter, because it is the first command they will type.

**A stale marketplace outlived its removal.** After the scratch experiment I removed the test
marketplace and the listing showed four, as it had before. A later `plugin install xeno@xeno`
nevertheless reported "Plugin xeno not found in marketplace xeno", which means something still
knew the name. The install against the real marketplace then worked, and the final state is
clean — four marketplaces, no xeno plugin — but the intermediate state was not what the listing
said it was. Recorded rather than explained, because I did not establish where it was cached.

**The vendoring list is section 13's and one entry of it is still missing.** `--vendor` now copies
the manifest, the skills and the hook wiring; `mcp.json` is in the document's list and not in the
tree. The function says so in a comment rather than copying nothing silently, which is the
difference between a list that is incomplete and a list that is wrong.

**`internal/plugin` holds no code, which is an unusual shape for this tree.** A package that
exists for its tests is a thing a reviewer will ask about. The alternative was putting the checks
in `internal/runner`, where they would have been tests about prose in a package about commands.
The package comment says why it is there.
