---
intent: github.com/triplem/xeno#176
phase: 04-verification
created: "2026-10-03T11:41:10Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+bdf4e26.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0ed0b5201dbe6c9ed073e266532c53111a40d637ca55246c7cb26180e9d69f03
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
437 cases pass, `gofmt` and `go vet` clean, `gate verify` at exit 0 over 228 verdicts — a trail in
which every lock records no sizes, which is the criterion that protects it. A lock written now carries
all three fields: this intent's 03-implementation records `repo_commit`, four `rules_applied`, and a
base of nothing, because this repository writes no profile. The demonstration, on a copy with a profile
over `docs/**` and section 5's example budget: four files, 320,480 bytes recorded with each size beside
its hash, inside the budget, phase green. Then a document grew by 300,000 bytes — what a month of
writing does here — and the same phase was judged again with no budget finding, where before this
intent the same growth would have reported a sealed phase over budget by a quarter. The inverse holds:
a lost file keeps its recorded size, so a deletion cannot quietly bring a phase inside its budget
either. Gaps: a link's document is hashed and sized by two reads; nothing checks a recorded size
against its file, which is the same trust the process places in a hand-written hash; the budget bounds
the declaration and not the reading; and seventy-nine intents can never be judged against a byte
budget, because their sizes were never recorded.
