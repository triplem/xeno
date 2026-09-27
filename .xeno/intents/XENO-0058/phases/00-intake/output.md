---
intent: github.com/triplem/xeno#58
phase: 00-intake
created: "2026-09-27T08:35:30Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+081dec4
plugin_version: 0.1.0-dev
language: en
secrets_hash: by-hand
context_hash: f6264093aeb94f536c22da099d5935cbcbabbf815292c79904631b57c9ddecf3
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: intake@1.0.0
strings_hash: 5fbb37323bf455c8dbda0b543f32593eaef4e0d2fd56f6568d849f7a7c0a7fc8
rules_hash: by-hand
---

# Intake

<!-- xeno:section:problem -->
## Problem

The plan says `xeno phase start` exports the qualified intent id and the phase id into
the environment the harness reads for request headers, so that a model request made
during a phase reaches the gateway carrying both. Nothing did it: `XENO_INTENT` and
`XENO_PHASE` existed only inside the CI wrapper, where a project sets them for `gate
run`.

Without it the gateway records requests it cannot attribute. #56 established that the
attribution works, per request and per tag, and that the tags have to be sent under
LiteLLM's own header names, because a header it does not know is dropped without an
error.

<!-- xeno:section:scope -->
## Scope

The two values, written where a later process can read them and printed where a shell
can eval them. `.xeno/local/phase.env` beside the run marker, removed by `phase finish`,
and `--export` for a person.

Not the mapping into headers. Which variable a harness reads and which header a gateway
keeps are the plugin's business, WP11's, and the runner holds no agent specific logic by
that package's own condition. So this closes the runner's half of #58 and leaves the
harness half open, including the verification that a header sent this way arrives.

<!-- xeno:section:context-rationale -->
## Why this context

**Export cannot mean what it says, and the file is the honest reading.** A child process
cannot set its parent's environment. Printing assignments to eval serves a person in a
shell and gives nothing to a hook that runs later in another process, which is most of
what WP11 wires, so the file is the truth and the print is a convenience over the same
two lines. One formatter produces both, and a test asserts they are identical, because
two ways of saying the same thing is exactly where drift lives.

**The lifetime is the run marker's, for the same reason.** A file left behind after a
phase is sealed would tag the next session's requests with a phase that has ended.
Attributing nothing is a gap; attributing wrongly is a false record, and the trail's
whole claim is against the second.

**The names are the runner's own, which is a boundary and not a preference.** WP11 is
done only when a phase runs from either agent with no agent specific logic in the
runner. `ANTHROPIC_CUSTOM_HEADERS` is claude-code's, `x-litellm-tags` is LiteLLM's, and
a runner that wrote either would know a harness and a gateway. The plugin knows both
already, and the CI wrapper reads these two names today, so nothing new is invented for
this.

**The id is the qualified one and never the directory name.** `qualified` refuses to
guess, and a test asserts the key does not appear in the file: a request tagged
`XENO-0058` instead of `github.com/triplem/xeno#58` would aggregate under something no
tracker knows.