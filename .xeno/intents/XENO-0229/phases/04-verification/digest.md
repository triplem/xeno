---
intent: github.com/triplem/xeno#184
phase: 04-verification
created: "2026-10-03T12:57:50Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+e8f68b1.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 2c336b3beb1e260ab3badb389c37adcf0d4ac499fda1c94ea837f73e773250f5
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Every acceptance criterion maps to what holds it, including the two that no test can hold: the
specification commit's ordering, held by `git show --stat e8f68b1` — two documents, nineteen
insertions, no `.go` file — and "the runner reads no other variable", held by `grep -rn os.Getenv
internal cmd`, which returns the enforcement token and this one while the other three names in
section 7's list appear in no Go file at all. The flag-beats-variable test sets the variable to
`9.9.9` so a pass cannot be a coincidence of both channels carrying one string, and the
read-once criterion is held by the `reopen` helper, which a per-command read would make
unnecessary. Hermeticity is tested rather than reasoned: the suite runs with
`XENO_HARNESS_VERSION` exported in this session and the absence tests still pass. Eighteen packages
ok, `gofmt` and `go vet` silent, `gate verify` 253 at exit 0, and the session channel measured end
to end on a throwaway tree with no flag anywhere. This intent's own six phases were written with the
variable alone. Five gaps: nothing sets the variable for a session yet, so one export replaced six
flags and #183 is still what removes the typing; the runner reads one of four variables, so that
issue's substance is untouched and its title less true; the value is unverifiable, as the surface
test depends on; hermeticity rests on one line in one fixture; and the repository now holds three
provenances for one field with nothing marking which is which.
