---
intent: github.com/triplem/xeno#167
phase: 00-intake
created: "2026-10-01T16:46:28Z"
schema_version: "1.0"
runner_version: 0.1.0-dev+15693cf.dirty
plugin_version: 0.1.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: c78b7950bdda2a62b83b35f36dab47ec87f6fd6989f1fcefb08c74b19a54c755
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
Thirteen gates read this repository and one statement cannot be made at all: the one a project's own
tool makes. Section 14 fixes the mechanism and nothing in the tree reads the declaration, runs the
command or produces such a check. The invariants for it exist and have never been met by anything:
`ExternalProvenance`, `carryForward` declining to attach a decision to an external finding, and
`Invariants` refusing one that carries it anyway all came from #66, and the comment at the verdict says
they were put there so that "a second path into it, an external gate above all, meets the same rule as
the first". There has never been a second path. Three things the documents require that the tree cannot
express: the declaration, which `model.Project` does not carry; the refusal, since "it runs only when
the hash matches" is a sentence about what must not happen and WP4 makes it a criterion; and the
contract, which is as far as "receiving and returning JSON" goes and is the load-bearing part, because
it is the only piece a project's tool has to agree with. And one thing neither document bounds: how long
a foreign command may take. The marking is the point and the reason this is awkward — there is no
sandbox, so what makes foreign code in the gate path acceptable is that the trail says which statement
was not Xeno's.
