---
intent: github.com/triplem/xeno#162
phase: 01-requirements
created: "2026-10-01T15:51:18Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+75f3667.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 967407d2d2819f32b34f62cf1c160a4da60ed85d244074e215531025070afa9d
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Each criterion is a rule file, a repository state and a verdict; the four commit types get two each,
because having a range and not having one are unrelated failures. `section-implies-section` holds the
implication and nothing else: an empty antecedent is green whatever the consequent says, since a gate
that read it otherwise would require every section of every template, and a section the template does
not have is a configuration error in the rule. `commit-message` judges every subject against a named
pattern with merge commits exempt where the rule says so and an unknown pattern name red with the list
of what is shipped. `commit-trailer` requires its key on every commit, compared case insensitively as
git reads trailers. `commit-signature` requires a valid signature, where valid is one recorded choice
because git answers in five letters. `approver-not-author` compares every decision's `by` against the
authors of the range, and no decisions is green, because a rule about approvals is not a rule requiring
one. A rule needing a range it did not get is red for all four, as its own criterion. A git failure is
a finding, not a panic. Non-goals: no rule ships, no sixth type, no expression in a rule file, no
inferred range, no trust chain, no identity resolution.
