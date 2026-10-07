# Hooks

What a harness has to be told so that a phase can record what it cost, and what each
harness can be told at all.

Both of the clients section 12 supports are wired, and one file wires them. Which of
them a reader is in does not change what is here and does not have to be worked out from
the directory's contents. What it does change is in the table below, under "What each
client can observe".

## One file, both clients

Section 7 says the two clients have "comparable lifecycle events and different wiring:
JSON settings on one side, `hooks.json` or an inline `[hooks]` table in TOML on the
other", and that "a plugin can ship its own lifecycle configuration, so the hooks travel
with Xeno rather than with an installation guide". That is true of a project's own
settings. It is not true of a plugin's, and the difference is worth stating here because
it is the whole reason this directory has one file rather than two.

Read in `openai/codex` on 2026-10-07:

- `codex-rs/core-plugins/src/loader.rs` takes the hooks of a plugin from the `hooks`
  entry of its manifest and, where the manifest declares none, from
  `DEFAULT_HOOKS_CONFIG_FILE`, which is `hooks/hooks.json` under the plugin root. It
  parses that file with `serde_json` as `HooksFile`.
- `codex-rs/config/src/hook_config.rs` defines `HooksFile` as an optional `description`
  and a `hooks` table whose keys are `PreToolUse`, `PermissionRequest`, `PostToolUse`,
  `PreCompact`, `PostCompact`, `SessionStart`, `SessionEnd`, `UserPromptSubmit`,
  `SubagentStart`, `SubagentStop`, `Stop` and `Interrupt`, each a list of matcher groups
  of handlers, and a `command` handler as `{"type": "command", "command": "…"}`.
- `codex-rs/exec-server-protocol/src/protocol.rs` lists the manifest paths a plugin is
  discovered by: `.codex-plugin/plugin.json`, then `.claude-plugin/plugin.json`, then
  `.cursor-plugin/plugin.json`. The second is the one this plugin already ships, which
  A77 chose against the installed Claude Code client.
- `codex-rs/hooks/src/engine/discovery.rs` sets `PLUGIN_ROOT` and `PLUGIN_DATA` for a
  plugin's command handlers and sets `CLAUDE_PLUGIN_ROOT` and `CLAUDE_PLUGIN_DATA`
  beside them — its own comment says "for OOTB compat with existing plugins that use
  this env var" — and substitutes `${…}` of those names into the command string before
  running it.
- `codex-rs/hooks/src/events/stop.rs` names the engine it dispatches through
  `ClaudeHooksEngine`.

So a Codex plugin's lifecycle configuration is a Claude Code plugin's, deliberately, and
the file it is read from is the one Claude Code reads it from. `hooks/hooks.json` with
`${CLAUDE_PLUGIN_ROOT}` in the command is the wiring both formats allow, and shipping a
second `.codex-plugin/plugin.json` would only add a manifest that has to be kept in step
with the first.

Two things follow that are easy to undo by accident. The manifest must keep saying
nothing about hooks, because a `hooks` entry in it replaces the default path rather than
adding to it. And the command must keep naming `${CLAUDE_PLUGIN_ROOT}` rather than
`${PLUGIN_ROOT}`, because only the first is set by both clients.

## The cost hook

Section 11 puts token counts in `cost.yaml` and says the local session logs are
evaluated on phase completion. A hook is the only place those counts are visible: hook
input carries `transcript_path`, `session_id` and `cwd`, and no token counts, so
something has to read the transcript the harness already wrote.

`xeno cost turn` is that something. It reads the hook's JSON on standard input, sums the
usage of the transcript it names, and appends one line to `.xeno/local/cost-ledger.yaml`
naming the phase `.xeno/local/phase.env` says is open. It exits zero whatever happens,
because it runs on every turn and a hook that breaks a session is worse than a missing
figure.

`Stop` fires once when a turn ends, which is the event that corresponds to work having
been done, and both clients have it. A tool-level event would fire hundreds of times a
turn and see the same numbers.

A project that wants the same thing without the plugin writes it into its own settings,
which for Claude Code is `.claude/settings.json` of the repository:

```json
{
  "hooks": {
    "Stop": [
      { "hooks": [ { "type": "command", "command": "./xeno cost turn" } ] }
    ]
  }
}
```

## What each client can observe

| | Claude Code | Codex |
|---|---|---|
| where this file is read from | `hooks/hooks.json`, by default | `hooks/hooks.json`, by default, where the manifest declares no `hooks` |
| `${CLAUDE_PLUGIN_ROOT}` in a command | set | set, beside `PLUGIN_ROOT`, and substituted into the command string |
| `Stop` | fires at the end of a turn | fires at the end of a turn |
| `transcript_path` on `Stop` | the session's own transcript, JSONL, one record per message | the session's rollout file, JSONL, one `RolloutItemWire` record per line |
| what `xeno cost turn` totals from it | `message.usage`, so a line reaches the ledger | nothing it recognises, so no line reaches the ledger |
| the harness's own version | `claude --version`, read by `bin/xeno-env.sh` | nothing the entry point can ask, so `tool_version` stays absent (A35, A81) |
| before a hook fires at all | the plugin is enabled | the plugin is enabled and its hooks are trusted, which Codex asks about per hook |

The row that costs something is the transcript. Codex's `Stop` input carries
`transcript_path` under that name, required by the generated schema at
`codex-rs/hooks/schema/generated/stop.command.input.schema.json`, and it names the
session rollout, whose lines are the `type`/`payload` records of
`codex-rs/history/src/rollout_payload.rs`. Token counts are in there, in a
`token_usage_record` payload, but not where `internal/cost` looks, and a reader for the
second shape is a change to the runner rather than to this directory. So a Codex session
records no cost line today. It records no wrong one either: `xeno cost turn` writes
nothing when a transcript totals zero, which is A35's rule in the one place it would be
cheapest to break.

## What is not wired

Section 7's execution table has two hook rows. Only the first of them, `Stop` for the
cost record, is here.

The row for G-Secret on every write is not, and the reason is that it would have to
invoke a check the runner does not implement yet. Section 7's first constraint on a hook
is that it "only invokes checks the runner already implements. No exclusive logic, or
the same input would yield different results depending on which agent produced it" — so
the hook cannot carry the check itself, and until `xeno` has a command that takes one
file and answers G-Secret about it, there is nothing for a `PostToolUse` entry to call.
Both clients have `PostToolUse`, with the written file in `tool_input`, so the wiring is
a few lines the day the command exists.

The row for G-Schema on an artifact write is not here either, and it is further away: it
fires "only when a write came through the MCP operation that completes an artifact", and
there is no MCP server.

## What it does not do

It attributes a turn to the phase that was open, and nothing else. A turn spent with no
phase open is recorded as `none` rather than being assigned to one, which is deliberate:
measured against this repository, attributing by each phase's own window captures about
seven per cent of what a session spends, because the work is done before `phase start`
is called. A63 records it.

So the ledger measures adherence to section 6's working sequence as much as cost. A
phase whose work happened inside it has a figure; one composed first and written
afterwards does not.

Nothing here is binding. Section 7's second constraint says a hook's result is advisory
because hook configuration lives on the developer machine, and the binding result is the
CI run, which runs no hooks and calls the runner directly.
