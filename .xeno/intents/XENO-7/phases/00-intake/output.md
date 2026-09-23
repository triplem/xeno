---
intent: github.com/triplem/xeno#7
phase: 00-intake
created: 2026-09-23T19:30:53Z
schema_version: "1.0"
runner_version: 0.1.0-dev
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: 6f7e2c0f93b0bc3f452aebae2c7ee8b05cd2b1a5a3c1edbfe8127c87a8ced844
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@0.1.0
strings_hash: by-hand
rules_hash: by-hand
open_questions:
  - key: Q-1
    text: What does a decision command do when the verdict it writes into is stale?
    options:
      - text: Refuse, and say that the phase has to be finished again first
        consequence: the against of every decision is then a hash that matched when it
          was given, which is what makes a decision checkable at all
        recommended: true
      - text: Write the decision against the current directory hash, recomputing silently
        consequence: convenient, and it releases a finding somebody may not have looked
          at in the state it is now in
      - text: Write it and record both hashes, the one judged and the one in force
        consequence: honest but a schema change, and the gate result has no field for it
      - text: Something else
        free: true
---

# Intake

WP7, the runner and its entry point. Most of the package was built by hand before M0;
what is left is the part that writes a decision, and it is the part this whole process
exists to record.

## Scope

`xeno gate approve <finding-id> --reason`, `xeno gate override <finding-id> --reason`
and `xeno obligation close <finding-id>`. The process definition names these three as
the only writers of a decision block and none of them exists, so a red finding can today
be fixed and nothing else. The governance function is unexercisable, which is a strange
state for the tool that exists to record it.

Then `xeno intent close --reason` for the abandoned case, with G-Complete in its second
mode and the intent level `gate.yaml` it writes.

And `xeno check commit-message`, which the hook and the gate both call so that one
pattern decides in both places.

## Non goals

`xeno init`, which is WP9. The CI wrapper and `xeno enforcement check`, which are WP10.
G-Complete's first mode, which runs as part of P5 and belongs with the gate list rather
than with the entry point.

## Open questions

Q-1 above. A decision carries `against`, the `artifacts_hash` it was made on, and that
is what makes it verifiable rather than a claim. What a command should do when the
directory has moved since the verdict is not settled anywhere in the process definition,
and every reading of it is defensible.

## Note

This intent is the first worked on a branch and merged through a pull request. The
convention has asked for it since the reference moved into the footer; until now the
work went straight to the default branch.
