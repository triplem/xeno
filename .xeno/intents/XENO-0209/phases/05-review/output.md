---
intent: github.com/triplem/xeno#109
phase: 05-review
created: "2026-09-29T18:45:23Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+3ec2429.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 47ecde682e61606b8b1181ca72ebd07e99e3464e9fdd60219281061c12f7ad1c
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: review@1.0.0
strings_hash: 60b45763687c63ee43d3de2ee7ae0b03a7eb026db07d8e184da43d2982411738
rules_hash: by-hand
---

# Review

<!-- xeno:section:review-checklist -->
## Review checklist

**Nothing under `docs/` changed.** Section 4 enumerates the files and Appendix B claims
the property; this makes the claim true at the level that lacked it.

**Nothing was invented.** One map transcribed from section 4, one constant, one function
mirroring the one a level up, and one call. No new gate: G-Complete already has the mode
an abandoned intent meets.

**The hash is untouched, and the transcript proves it.** `bf49ef35…` on both sides of
the change, green before and red after. What changed is that something judges what the
hash covers.

**The two levels keep separate lists**, so neither check stops distinguishing what
section 4 distinguishes. A single map would have accepted `output.md` in an intent
directory.

**A directory is reported although the hash cannot descend into one**, because a
directory nobody wrote is where files appear next, and the phase level had that right
without saying why. It says why now.

**It reports rather than refuses.** Two alternatives would have made the property true
by construction and both prevent an abandonment, which section 6 deliberately writes
rather than refuses: the intent is dropped either way and the record should say what is
missing.

**#138 is fixed here and is not this intent's subject.** Two findings on a close made
`Status` refuse, so AC5 was unreachable without it. It is filed, its commit is separate,
and the deviations say why it rode along. The fix is what A25 already asked for:
findings routed through `carryForward`, and `Invariants` run on the second path into a
verdict, which is what that function exists for.

**Every verdict still matches**, 133, and the phases were written in order.

**What this does not make true is said in the gaps**: nothing here is abandoned, so the
guard has judged nothing real yet.

<!-- xeno:section:release-notes -->
## Release notes

An abandoned intent's verdict now reports what is in the intent directory and should not
be. A file that is not `intent.yaml`, `assumptions.yaml`, `learning.yaml` or `gate.yaml`
is a finding, and so is a directory other than `phases/`.

Appendix B rests the normalisation of both hashes on every file they cover being one
Xeno wrote, and calls that a checked property rather than an assumption. The check
existed for a phase and not for an intent, so the sentence was true of one of the two
hashes it covers. It is true of both now.

The hash itself is unchanged: what it covers, what it excludes and how it normalises. A
stray file still enters the hash of the run that reports it, which is section 4's
arrangement — the verdict records both.

Two findings on a close used to produce an error and no verdict at all, because `xeno
intent close` routed its findings through neither `carryForward` nor `Invariants`, so
their ids were empty and collided. Both now run on that path, as they do for a phase,
and a decision taken on an intent level finding survives a re-close.

<!-- xeno:section:residual-risk -->
## Residual risk

**It has judged nothing.** No intent in this repository is abandoned, so the check is a
guard before its data and the first abandonment is when it earns its keep. The
two-binary comparison used a scratch intent.

**Reported, not prevented**, and the transcript makes that concrete: the red verdict
carries the same hash the green one did. Whoever expects the check to keep a stray file
out of the hash will be disappointed by design, and section 6's argument about
abandonment is why.

**Two lists are transcribed by hand from a document no gate reads.** If section 4 gains
a file at either level, the check reports it as unknown until somebody updates a map.
That is the failure in the safe direction and it is still a failure, and nothing
connects the document to the code.

**Four near-identical strings across two levels.** The wordings differ only in one word,
and a change to one will not change the other. The parallel is visible to a reader
comparing them and invisible to anything else.

**#138's fix widens what `intent close` can refuse.** `Invariants` now runs there, so a
malformed finding from that path stops the close rather than being written. That is the
same trade XENO-0107 made and the same risk: the loudest response on the least examined
path, and this path has one caller and no external gate behind it yet.

**Accepted with the five named.** The state it replaces is a verdict bound to a
directory nobody had looked at, which section 7 says would otherwise be the only verdict
not bound to what it judged.
