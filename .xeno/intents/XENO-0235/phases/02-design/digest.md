---
intent: github.com/triplem/xeno#199
phase: 02-design
created: "2026-10-03T18:23:23Z"
schema_version: "1.0"
runner_version: dev+a897f2a.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 0d268a4d8c7863265a55d18197a3bfff2c16ec75be64f01202ba8e091150295e
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
---
The release copies the tree into the package because a `go:embed` directive cannot reach outside its
own package directory — a constraint rather than a preference, measured in a scratch module before
the design. The target is `internal/plugin/embedded/plugin/`, beside a committed placeholder, so one
path's presence distinguishes a release from a build; the placeholder is committed because `go:embed`
fails to compile when its target is absent and a clean checkout has to build; `all:` is on the
directive because the manifest lives under a dotted directory. `Shipped()` settles the question on
the manifest rather than the directory, since `fs.Sub` succeeds on a path that is not there.
`vendorPlugin` becomes one `fs.WalkDir`, replacing three functions that each knew a tree and a depth
and two of which were silently incomplete; the source is the embedded copy where there is one,
because a release carries the bytes its digest was taken over and vendoring a directory would produce
a tree its own gate fails. `vendoredMode` is a function of the path and not of the source's mode,
which does not survive an `embed.FS` and would make a release and a development build vendor
different trees. The release stamps with inline `python3`, because the runner cannot — a command that
edits the plugin is the one thing `--vendor` exists to copy. Eight alternatives refused, two of them
not choices: the directive's reach and its behaviour on an absent target, both of which shaped the
design rather than being chosen within it.
