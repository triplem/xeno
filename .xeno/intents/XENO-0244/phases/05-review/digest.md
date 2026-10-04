---
intent: github.com/triplem/xeno#221
phase: 05-review
created: "2026-10-04T10:59:08Z"
schema_version: "1.0"
runner_version: dev+24becc3.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a7b87bcb8345b652e04680abc3344d171e4255dc9b61ee9724f6d13e4e2cc338
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The review answers the three shipped rules, two met and one not applicable, and checks
the one claim the whole intent rests on: that no `artifacts_hash` covers
`evidence/attached.yaml`. It was read out of `DirHash` in the intake rather than out of
the prose, and then measured — `gate verify` at exit 0 over 342 verdicts before the
change, after the fifteen corrections and after the gate, with the middle reading the
one that proves it, since a wrong reading would have shown fifteen divergences there and
nowhere else. After the four attachments the trail stands at 344 verdicts, still exit 0.
The release notes say what changed for a reader who did not follow the work: an
attachment's result is judged by G-Evidence on the file that carries it, an out-of-set
value is declined at the attach with its own remedy in its own sentence, the fifteen are
corrected in the one file no hash covers, and the order of those two steps is the change
rather than a detail of it. The migration note is the rename of `Result.Unbindable` to
`Declined`, internal to the module, with the behaviour beside it named: the findings it
makes possible are new for a tree that has none, because the records that would have
produced them were corrected in the same commit. Seven residual risks are recorded. The
first is D-1's named exclusion and is filed: an attachment with no result at all on a
kind that needs one is still unjudged. The second is that G-Test, the reader that would
have caught the fifteen, is still absent. The third is that the correction sets a
precedent for a hand edit whose only safeguard is the figure above. The last is the one
worth re-reading later: two gates now compare against the same closed set in two places,
deliberately, and if a third reader appears the argument for keeping them separate
should be read again rather than repeated. This phase's four attachments arrived through
the new check, all four reading `pass`, which is the accepted path proved against real
manifests rather than against a fixture. Read in this phase: the four rule files, the
attached reports and `attached.yaml` as they arrived, P4's verdict before and after, and
the diff of the fifteen once more.
