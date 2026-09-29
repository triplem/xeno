# Hooks

What a harness has to be told so that a phase can record what it cost.

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

For Claude Code, in `.claude/settings.json` of the repository:

```json
{
  "hooks": {
    "Stop": [
      { "hooks": [ { "type": "command", "command": "./xeno cost turn" } ] }
    ]
  }
}
```

`Stop` fires once when a turn ends, which is the event that corresponds to work having
been done. A tool-level event would fire hundreds of times a turn and see the same
numbers.

## What it does not do

It attributes a turn to the phase that was open, and nothing else. A turn spent with no
phase open is recorded as `none` rather than being assigned to one, which is deliberate:
measured against this repository, attributing by each phase's own window captures about
seven per cent of what a session spends, because the work is done before `phase start`
is called. A63 records it.

So the ledger measures adherence to section 6's working sequence as much as cost. A
phase whose work happened inside it has a figure; one composed first and written
afterwards does not.

## Codex

Not written. The reader here is Claude Code's transcript, chosen by the `agent.tool` a
project records in `project.yaml`, and a second reader is additive.
