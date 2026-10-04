---
intent: github.com/triplem/xeno#217
phase: 00-intake
created: "2026-10-04T20:02:38Z"
schema_version: "1.0"
runner_version: dev+fbf8a72.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 10be97648005d1f95c08630b9dc28fcc3dac60db91738a2f280ec4ccacb00841
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The intake settles a prior question before the one #217 asks, and corrects a premise of
the issue. Section 5 puts `context-scope.yaml` at P0 and calls it a budget P0 produces;
nothing in the repository has ever produced one, across 100 finished intents, 345
verdicts and 346 locks with no `files` list between them. Five readers depend on it and
all are inert: `informationBase` resolves nothing, G-Freshness's second half iterates an
empty list, `budget` has no number, `links` has no link, and `ChangedSince` has no base,
so a repeated phase is told to re-read everything by being told nothing. Three of them
return nil with a comment saying so, which is A74's distinction undrawn — an empty list
says a set was resolved and came out empty, absence says there was nothing to resolve —
and `gates.go:681` is G-Schema's only mention of the scope: it permits the file at P0
and nowhere requires it. The statement is not missing, only its machine-readable half:
every one of the 100 intakes names its reading in prose in `context-rationale`, and the
scope is per intent rather than per project, which `informationBase` is explicit about.
WP8's own blocker, "P0 cannot produce a scope until the format is fixed", is cleared —
`model.Profile` and section 5's example fix it — and the writer was never added behind
it. Q-1 put the prior question, whether the context economy is something v1 means to
have at all, against keeping it optional and cutting it to 1.1; the maintainer chose
keeping it with the scope required at P0. Q-2 then had to be put because #217's stated
cost is wrong: it claims A74 leaves the old verdicts alone, and A74 binds G-Policy alone
to the rule set a phase records, while `Verify` recomputes and compares every phase's
status — so a G-Schema check would re-judge 100 sealed P0s, and backfill cannot answer
it, because `PhaseExcluded` holds only `gate.yaml` and `cost.yaml` and the scope is
therefore inside `artifacts_hash`. That is the inverse of XENO-0244, which was orderable
only because `evidence/attached.yaml` lies in a subdirectory. The maintainer's answer,
that the scope should be required in the future, is the refusal in `phase finish`:
forward only, because commands are not re-run where verdicts are recomputed. Scope: a
writer, this intent's own scope written by hand before its P0 is judged, and the
refusal. Out of scope and filed: the 100 sealed P0s, which nothing can backfill;
validating the scope's content, which section 5 gives G-Schema and nothing does; WP8's
other half, the out-of-scope read named in the digest; a project default for a scope's
contents; WP15 and WP20. One risk is named for P2 rather than assumed: an intent whose
scope names a file it then edits has moved an input of its own earlier phases, and
G-Freshness will say so from P3 on, which no intent has ever been in a position to
discover. This intent's own scope is the first in the repository, so its P1 to P5 carry
the first non-empty `files` list and give three checks their first live input — 43 files
and 904,078 bytes, resolved at P1 and counted rather than guessed. Two things were found
by that first resolution and nothing else could have found them. A budget overrun turns
a phase red, because the check sits in G-Schema and `result` makes any finding a `fail`,
where section 5 says it is "deliberately a finding and not a red gate in the sense of
stopping work"; the specification wins, the runner has no notion of a finding that does
not fail its check, and that is a mechanism rather than a line, so it is #235 rather
than absorbed. And the budget is frozen when P0 is judged, because the scope is inside
its `artifacts_hash`: an intent that finds the number too tight at P3 cannot correct it.
This one met that at P1, rolled back an unwritten phase and corrected the number while
nothing was sealed. The scope then widened twice more, both times from what the first
resolution exposed rather than from a change of mind. The maintainer renamed the
artifact in commit fbf8a72, from the context profile to the context scope, because
profile names a set of settings one switches between and an intent has exactly one; the
code rename and this intent's own file follow, and the window for it is this intent,
since a merged scope would sit inside a sealed artifacts_hash. And that same commit
blocked the intent: it moved two documents the scope names, G-Freshness read P1's own
lock, P1 went red and A12 refused P2 on a red predecessor. Section 5 states two limits
on that check and staleReads honours neither, so #236 is in scope as well, because the
refusal this intent adds is not safe without it — a mandatory scope plus a check that
reads the gated phase's own lock makes every implementation intent report itself out of
date for doing its job. The two findings were approved by the maintainer rather than
re-run, which is what section 5 provides for a stale phase and the only route available,
since the remedy the finding prints is an act phase start refuses: #237. Read: #217,
three passages of section 5, A74 in full with A6 and A10, WP8 and WP20, `runner.go` for
`informationBase`, `phase start`, `ChangedSince`, `Verify` and four refusals, `gates.go`
for the three nil readers and line 681, `hashing.go` for `DirHash` and `PhaseExcluded`,
and `model.go` for `Profile`, `ContextLock.Files` and `MatchPath`.
