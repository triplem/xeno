---
intent: github.com/triplem/xeno#176
phase: 00-intake
created: "2026-10-03T11:34:08Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+f001058.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: f7ba15bbe7766e6bd043ff29ce40f2058496fd833e90095ac8641f10af85fa0e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The budget is the one number in the trail that can change after a phase is sealed: G-Schema takes the
recorded context's size from the tree when the check runs, so a phase inside its budget when it ran is
over it a week later if a file grew, and `gate verify` says so about a phase nobody touched. It is an
exception by omission — A66 and A74 both rest on the opposite rule, that a phase is judged against what
its own lock recorded, and A74 exists because four rule files turned twenty-four sealed verdicts red.
The specification now has the field: the commit before this one adds `bytes` to the lock's `files` and
the sentence that says why, which is the order the first standing rule requires. What the code must be
careful about is the trail behind it: every lock in seventy-nine intents records no sizes, and a check
that summed absent sizes would read each one as a zero-byte context and therefore permanently inside
its budget — quietly, which is worse than the finding it replaced. A lock that records nothing has to
produce no byte finding, which is the distinction A74 drew for rules and the register's closure drew
for decisions: absent is not empty.
