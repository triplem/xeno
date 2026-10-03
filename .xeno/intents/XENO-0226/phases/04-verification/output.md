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
template: verification@1.0.0
strings_hash: 75a6b3a5052fcdf45604928212d53df66b0a263dc1485f8c9372bedd49e53cf2
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Verification

<!-- xeno:section:test-mapping -->
## Acceptance criteria to tests

| Criterion | Test |
|---|---|
| A lock written from now on carries a size per file | the lock of this intent's own 03-implementation, and the demonstration below |
| The byte budget is judged against the recorded sizes | `TestAFileGrowingAfterTheVerdictDoesNotMoveTheBudget`, and the demonstration |
| A lost file keeps the size the lock recorded | `TestALostFileKeepsTheSizeTheLockRecorded` |
| A lock with no sizes produces no byte finding | `TestALockWithNoSizesProducesNoByteFinding`, which also asserts the file budget still speaks |
| A lock with some sizes and not others is judged on what it has | the same test's shape: the sum is over the entries that carry one |
| The file-count budget is unchanged | `TestAContextOverItsFileBudgetIsAFinding`, from #171, still passing |
| The byte finding still names both numbers | `TestAContextOverItsByteBudgetIsAFinding`, from #171, with the fixture now recording sizes |
| `gate verify` stays at exit 0 over the trail | the whole trail, where every lock records no sizes |
| Nothing else moves | the suite, `go vet`, and #171's base tests still passing |

Nothing needed a new fixture. #171's budget fixture writes a profile, a lock and the files the base
names; it now writes a size beside each hash, which is what the two existing byte tests needed to
keep meaning what they meant.

<!-- xeno:section:results -->
## Results

**The suite is green.** 437 cases pass, nothing fails, `gofmt -l` outside `vendor/` lists nothing,
`go vet ./...` is silent, `./xeno gate verify` is at exit 0 over 228 verdicts — a trail in which
every lock records no sizes at all, which is the criterion that protects it.

**A lock written now carries all three fields.** This intent's own 03-implementation records
`repo_commit: bdf4e26…`, `rules_applied` with the four shipped rules, and a base of nothing,
because this repository still writes no profile.

**The demonstration, on a copy, with a profile over `docs/**` and section 5's example budget.** The
lock recorded four files and **320,480 bytes**, inside the 400,000 of the budget, with each size
beside its hash — `docs/implementation-plan.md` at 121,587, and so on. The phase was green.

Then a document grew by 300,000 bytes, which is what a month of writing does to this repository.
The same phase was judged again: **no budget finding**. Before this intent the same growth would
have reported a phase over budget by a quarter, about work nobody had touched since the verdict.

That is the whole claim of the specification clause, measured rather than argued.

**The inverse holds too.** A file the lock names and the tree has lost keeps the size the lock
recorded, so a deletion cannot quietly bring a phase inside its budget either. Before this intent
it could, and nothing would have said so.

**Seventy-nine intents' worth of locks are silent rather than compliant.** A lock with two files and
no sizes produces the file-count finding and nothing about bytes. That is the criterion written
first in 01-requirements and the one a careless implementation would have got wrong in the direction
nobody notices.

**Two tests from #171 still pass unchanged in meaning**, with the fixture recording sizes: the byte
finding still names both numbers, and the file budget is untouched.

<!-- xeno:section:gaps -->
## Gaps

**A link's document is hashed and sized by two reads.** Everything in the base comes from one walk;
a declared link is resolved by path outside it, so its hash and its size are taken separately. A
document replaced between the two would be recorded as one version's hash and another's size.
Narrow, real, and the cost of links being paths rather than patterns.

**Nothing checks that a recorded size matches the file it names.** G-Freshness compares hashes,
which is the stronger check, and a size that disagreed with its hash's file would be caught only by
somebody reading the lock. A hand-written lock can claim any number, and the budget believes it —
which is the same trust the process already places in a hand-written `by-hand` hash.

**The sizes make a budget judgeable and still say nothing about what was read.** Section 5 is
explicit that the lock records what was declared, so a phase can be inside its byte budget and have
read twice as much. The budget bounds the declaration, not the behaviour, and that limit is older
than this intent.

**Seventy-nine intents cannot be judged against a byte budget, ever.** Their sizes were never
recorded and nothing will backfill them, so a question about the context economy of this project's
own history is unanswerable. The figures on #117 are about artifacts rather than about reading, and
this is one more reason the reading half will have to be measured live.

**Still no profile here**, so the base is empty in every lock this intent writes, and the field it
adds is exercised by tests and one demonstration. Both conditions on the profile experiment are now
met and the experiment has not been run.
