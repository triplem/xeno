---
intent: github.com/triplem/xeno#355
phase: 02-design
created: "2026-10-10T12:18:00Z"
schema_version: "1.0"
runner_version: dev+90c7227.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: a3044769b7b891b8cba111c1407e2b1d851e0e4546870058ac818bf45c649d2a
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276 (Claude Code)
---
The design is four edits to `release.yml` and one row plus two sentences on
`docs/supply-chain.md`. The `Dockerfile` is not edited, which is the point of it.

Nine alternatives are written down and the two that nearly got built are worth naming.
`docker pull`, `docker tag`, `docker push` is the three-command version every runner can
run and it loses the whole property: a pull resolves the index to one platform, so the
mirror would hold a different digest and the pin would become two pins kept equal by hand.
And widening all five publishing conditions is the simpler-looking edit; gating the
`release` step on `github.event_name == 'push'` instead leaves the four binary-channel
conditions untouched and false by construction, because a step that did not run has no
outputs. One line rather than four, and the four that matter are the ones not edited.

The mirror's hit check and the proof that the copy preserved the digest are one command,
run twice. `imagetools inspect "$IMAGE/base@$digest"` decides whether a copy is needed and,
after a copy, fails with the reference it could not resolve if the index was re-serialised.
So criterion 2 is a check the pipeline makes on every miss rather than a premise: the
property belongs to buildx, the local docker has no buildx plugin and no daemon, and a
design resting on somebody else's behaviour should fail on the day it stops holding.

One defect was found by running the thing rather than reading it. The dispatch's download
was simulated against v0.60.2, and the asset arrives as mode 0644 where `go build` produces
0755. `COPY` preserves the mode, so a dispatch that only downloaded would have published an
image whose `/usr/local/bin/xeno` cannot be executed, and every gate job in it would fail
on a permission rather than on anything a reader would trace back to a release workflow.
The dispatch path chmods, and the reason is in the design rather than in a comment.

Two things are taken rather than measured and both are written in the impact section, since
this intent has no assumption register: that `imagetools create` copies an index through
unmodified, which the inspect above turns into a check; and that a container package
`GITHUB_TOKEN` creates is readable by the same token on a later run, whose failure mode is
a copy on every release rather than a broken one.

Two sections, no open question, no decision.
