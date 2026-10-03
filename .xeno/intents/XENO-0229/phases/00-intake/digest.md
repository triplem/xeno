---
intent: github.com/triplem/xeno#184
phase: 00-intake
created: "2026-10-03T12:54:05Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 6e42bfd05a27291c07441bb3fdc6f5b3eab080b3d1eff39ae47ff6ab432708f7
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
#181 gave `tool_version` a flag and A80 recorded it as a channel rather than the channel: the better
home is a variable in the list section 7 enumerates, and the first standing rule makes that list a
person's commit. The constraint was the rule and not the design — A80's own words are that the flag
"does not block it and would become its override" — so asking the maintainer was all that stood
between the two, and the approval has been given. What the flag leaves unsolved is remembering it: a
skill instructs and does not enforce, so a phase written without it is as red as before #181, which
is XENO-0228's first gap and its main residual risk. Six phases times one flag is six chances to
forget a value that holds for the whole session, and a variable is the channel for exactly that
kind of value. In scope: the specification change as its own commit, then the runner reading the
variable in `New`, the flag overriding it, absence unchanged, and tests that clear it rather than
inheriting what the machine exports. Out of scope: the entry point and the other three variables,
which stay with #183; any client specific variable, since normalising those is the entry point's
work by section 7's own account; removing the flag, which A80 said would become an override;
recording `XENO_HARNESS`, which is a separate argument about where `tool` comes from; and
backfilling XENO-0228, whose flag-reported values are correct. The claim is deliberately the
smaller one: this removes the remembering for anybody who exports the variable once and for nobody
who does not, because nothing yet exports it for them.
