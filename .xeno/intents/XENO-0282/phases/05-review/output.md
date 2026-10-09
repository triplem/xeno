---
intent: github.com/triplem/xeno#332
phase: 05-review
created: "2026-10-09T13:41:10Z"
schema_version: "1.0"
runner_version: dev+30b1dea.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 7155901919bb109a6007a60507a30aee638c36917bcecd33ef5a450233310fc2
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
review_checklist:
    - rule: deviations-are-traceable
      result: met
      note: P3 records none; the second pass over the fixtures is a correction inside the phase.
    - rule: interface-change-needs-a-migration-note
      result: met
      note: The names an issue must carry change; the migration is the shipped script once per tracker and a /xeno approved comment per issue, named in docs/commands.md, the init line and the refusal.
    - rule: new-dependency-needs-a-rationale
      result: not-applicable
      note: A shell script needing gh or glab; nothing enters the binary or the pipeline.
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**deviations-are-traceable — met.** P3 records none and says why the second pass over the fixtures is a correction and not a departure.

**interface-change-needs-a-migration-note — met.** The interface is what an issue has to carry: `xeno-approved` and `/xeno approved` instead of `approved` twice. The migration is one script run per tracker and one new comment per issue already approved, and the refusal names what is missing; `docs/commands.md` says so and so does the init line.

**new-dependency-needs-a-rationale — not-applicable.** A shell script that needs `gh` or `glab`, which a person who manages labels has; nothing enters the binary or the pipeline.

**What no rule asks.** Whether a rename without a transition is right for an adopter mid-flight: A107 says why two names are better than four and names the clause change a transition would be.

<!-- xeno:section:release-notes -->
## Release notes

**The approval names carry the tool's name.** An issue becomes an intent with the label `xeno-approved` and a comment whose first line is `/xeno approved`; the old `approved` pair starts nothing. The reason: a brownfield tracker may have an `approved` label of its own, a bare word is also what people write to each other, and a mention would have notified a real account.

**One script creates the label.** `.xeno/plugin/bin/xeno-labels.sh github OWNER/REPO` or `gitlab GROUP/PROJECT`, shipped with the plugin, idempotent, and named by `xeno init` among the settings a person makes on the host.

**For a project already using the old names.** Run the script, relabel the issues, and write `/xeno approved` on each one still to start; the refusal tells you which half is missing.

<!-- xeno:section:residual-risk -->
## Residual risk

**Nine open issues of this repository carry the old comment.** Relabelling at the merge moves the label; the comments stay as written, and each issue needs a new one before `intent start` accepts it. The maintainer writes those, and until then every one of them is refused.

**The script's GitLab branch is unproved.** No GitLab host is available here; the flags are `glab`'s as documented today.

**The sealed trail names the old pair** in four intakes and will do so forever; that is what it is for.

**Belongs to the specification.** Whether a transition period should ever exist, which A107 leaves to a clause change.
