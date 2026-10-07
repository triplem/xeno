# xeno

Requirements, design, code and evidence are produced separately, at different times,
partly by people and partly by models. The only thing holding them together is their
binding to an intent. Xeno makes that binding provable: it runs one change through six
phases, writes an artifact per phase into the repository the change lives in, and judges
each one with deterministic gates that make no network call.

What it produces is evidence. Serving as an evidence source for ISO/IEC 42001
certification is the declared next milestone; Xeno supplies the trail and does not run a
management system. It targets projects with their own infrastructure, mostly brownfield,
and the artifacts live in the project's own repository rather than in a service.

The name is a recording technique: two tracks that were never played together, and never
at the same tempo, laid over one another afterwards. What comes out sounds coherent even
though it never happened at once, and the coherence is in the assignment rather than in
the recording. The process definition's Appendix C says it properly.

## Reading further

**[docs/](docs/) is the index.** Two of the documents there are normative, and
[the process definition](docs/process-definition.md) is the one to read first: it fixes
which phases exist, which artifacts they produce, which gates apply and how learning
happens. Where it and the code disagree, it wins.

## Running it

    go build ./cmd/xeno
    go test ./...

Setting a repository up, then a first intent, in the order the commands are run:

    xeno init           --vendor --project OWNER/REPO
    xeno intent start   --for ISSUE
    xeno phase start    --intent KEY --phase 00
    xeno phase finish   --intent KEY --phase 00

Between the last two, the phase's sections are written with `xeno section set` and the
gates run on `phase finish`. Every command that changes something ends by naming the
next step of the working sequence, so the sequence does not have to be memorised;
`--no-next` leaves that out, and the commands a pipeline or a git hook runs never print
it. Every command takes `--root DIR`, which defaults to the working directory.

`xeno --help` lists them all, and [docs/commands.md](docs/commands.md) is the same text
on the web.

Xeno collects no telemetry. Nothing it writes leaves the repository it writes in, and
there is no endpoint for it to leave towards.

## This repository runs on it

Xeno governs its own construction, so `.xeno/` here is the longest worked example there
is: every change since M0 runs six phases and writes real artifacts, and the whole trail
is readable with `xeno intent status --all`. Three things about it read oddly to
somebody expecting a released tool's output — what the version fields say, why G-Supply
reports `not-implemented` throughout, and why the earliest intents stop after their
intake — and
[docs/the-trail-in-this-repository.md](docs/the-trail-in-this-repository.md) says why,
with the coverage against the plan's acceptance criteria.

The assumptions and decisions taken while building are in
[docs/assumptions.md](docs/assumptions.md), one row each, and the open ones want
confirming. Its last two sections are what is built and what is not.
