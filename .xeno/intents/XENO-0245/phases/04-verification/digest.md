---
intent: github.com/triplem/xeno#217
phase: 04-verification
created: "2026-10-04T21:28:02Z"
schema_version: "1.0"
runner_version: dev+fbf8a72.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 4f20aae59686cb170806a7342c8ee3859732261e74c5b10053cdea647cc6840b
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Verification maps every acceptance criterion of P1 to what asserts it, and the figures
are taken rather than claimed. The suite is at exit 0 over eighteen packages with two
carrying no tests, gofmt and go vet are silent, the build is clean, and gate verify is
at exit 0 over 349 verdicts, against 346 before the change — the three added being this
intent's own P0, P1 and P2. No verdict in the hundred historical intents moved and no
divergence was reported at any point, which is the figure the whole of D-2 rests on: a
requirement in a command rather than in a gate reaches no sealed phase. The rename is
complete, by grep: no ContextProfile, model.Profile or context-profile outside vendor
and outside the trail, and no "context profile" in either normative document. The
trail's own sealed artifacts keep the old name where they described the state before
fbf8a72, which is correct, because they are records of what was true when they were
written. Two figures are dogfooding results rather than test results. P1 to P4 each
resolve 43 files and 904,078 bytes against the declared budget of 50 and 1,000,000, the
first non-empty information base in the repository, and P1 carries no approval and no
override, which is the criterion P1 set to check that the fix addressed what actually
happened. And phase start for this phase printed the files P3 changed, which is
ChangedSince producing output for the first time since it was built. Evidence declared:
a test-report of format go-test-json, 2,369 lines with no fail action in them, produced
by go test -json ./..., and a build-log. Six gaps are recorded rather than left. The
staleness half has no local coverage by design, because both limits need a commit range
and section 12 forbids inferring one, so only CI exercises it and nobody running gate
run locally will see it fire. Its silence cannot be reported, which is #235, and this
intent now has three findings wanting that same mechanism. xeno scope set has not been
run against a real intent, deliberately, because the fresh created in its header would
restart P3's cascade; the only scope in existence was written by hand and the next
intent is the measure of whether the writer is usable. The budget has never been
exceeded, so the one reader whose behaviour is disputed between the code and section 5
is still the one that has produced nothing. G-Supply, G-Secret and G-Test remain
not-implemented, as for every intent in the trail, so the green above is read for what
it is. And CLAUSE-READERS.md is now stale in exactly the way it predicted of itself,
since this change renamed symbols it names; re-running that audit is its own act and not
this one.
