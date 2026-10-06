---
intent: github.com/triplem/xeno#257
phase: 01-requirements
created: "2026-10-06T07:50:46Z"
schema_version: "1.0"
runner_version: dev+8e3b29b.dirty
plugin_version: 0.0.0-dev
language: en
secrets_hash: 8002ba2da35fee01baac158b0842b29929ea7e1e5b4f8636ab6298d44cb135b9
context_hash: 15928b5c6c09dec86f414f74e44f67a832520c75d5bcb421be4229c6d6cb0433
model: claude-opus-5
tool: claude-code
tool_version: 2.1.276
template: requirements@1.0.0
strings_hash: 448045d7b8e291bd71bcb3b970d7bb362fbc36c4c7879fb7fa977a6c0adb74ff
rules_hash: ce2250ab2bb61ca8c0b7a39bab95ac1bbee7e1cbd6a4463d4b4ce535564a9721
---

# Requirements

<!-- xeno:section:acceptance-criteria -->
## Acceptance criteria

1. `examples/templates/` holds six directories, one per phase id, each with `template.yaml` and
   `strings.en.yaml`. Both files, because `Load` treats a missing bundle as an error rather than
   falling back to another language.

2. Each `template.yaml` keeps every section the shipped one defines, with the same ids, versions
   and phase, and differs only in `required` flags. Nothing is deleted, so `section set` goes on
   accepting every id and no rule can name a section the template does not know.

3. Six sections stay required: `problem`, `acceptance-criteria`, `changes`, `test-mapping`,
   `results`, `release-notes`. Eleven become `required: false`.

4. `release-notes` stays required and its file says what reads it: `release-notes-are-filled`
   through `section-non-empty`, the only section id any rule names. A candidate that dropped it
   would fail a shipped checked rule at P5.

5. The review template's file says that the `review_checklist` frontmatter is read by G-Policy
   whatever the `review-checklist` section's flag is, so making the section optional does not make
   the answers optional.

6. `acceptance-criteria` and `test-mapping` stay required, and their files say this is by decision
   rather than by a reader: nothing reads them today, and #258 settled that section 9 will make a
   criterion identifiable and a gate will read the mapping.

7. Each file says what it drops and why, in its own header, in the shape
   `examples/rules/documentation-follows-the-change.yaml` uses: what it is, why it is not enabled,
   how to adopt it.

8. A README under `examples/templates/` says how to try it, that trying it is how the question is
   answered, which figures to collect, and that the baseline to compare against is the seventeen
   recorded nine times on #117.

9. Every `template.yaml` and `strings.en.yaml` parses as the type the loader reads, and the section
   ids in each bundle's `headings` cover the ids its template defines.

10. Nothing is adopted. `.xeno/config/templates/` does not exist after this intent, no shipped
    template changes, and `gate verify` exits 0 with the verdicts that exist now intact — which is
    what proves this repository's own process is unchanged.

11. `go build`, `go test ./...`, `go vet ./...` pass and `gofmt -l` outside `vendor/` prints
    nothing.

12. One commit, `Closes #257`, and the issue carries `wp20`.

<!-- xeno:section:non-goals -->
## Non goals

Not adopting it, here or anywhere. Copying it into `.xeno/config/templates/` would change this
repository's process mid-session and make the next intent's figures incomparable with the nine
already on #117, which is the comparison the fixture exists to enable.

Not a change to any shipped template. Nothing under `.xeno/plugin/templates/` is touched, so no
adopter's process moves and the plugin digest is unaffected.

Not a recommendation that the shortcut be taken. The plan forbids settling the shape in advance of
measurement; this is the thing that makes measuring cheap, and the candidate is an argument to be
tested rather than a conclusion.

Not a second language bundle. `strings.en.yaml` only, said in the README, because a project on
another language writes its own and a fixture that shipped a German bundle it had not checked would
be worse than one that says it does not.

Not a change to what a section is. Section 9 fixes that; the candidate moves a flag the template
layer owns.

Not a deletion of any section. Every id the shipped template defines stays defined, so
`section set` keeps working for all of them and a rule naming one cannot find it missing.

Not the figures from running it. Measuring an intent against the reduced set is the next step and
needs an intent to run; this provides the set and the method, not the result.

<!-- xeno:section:constraints -->
## Constraints

A missing strings bundle is an error. `Load` says so and will not fall back to another language, so
each directory needs `strings.en.yaml` and the fixture is twelve files rather than six.

`Missing` is the only reader of the flag. It reports required sections that carry nothing, so
`required: false` is the entire lever and nothing needs deleting — which also means the flag changes
what the next-step suggestion asks for and nothing else.

One section is read by a gate and the candidate may not touch it. `release-notes` is named by
`release-notes-are-filled` through `section-non-empty`; a candidate dropping it ships a template
that fails a shipped checked rule at P5, which would make the fixture a trap.

G-Policy reads frontmatter, not a section. The `review_checklist` list is required per review rule
whatever the `review-checklist` section's flag says, so the candidate may make the prose optional
and must say that the answers are not.

A72 settles where it goes. A review rule in `given/builtin/` reaches every adopter; a template in
`.xeno/plugin/templates/` is the same reach for the same reason, so a candidate belongs under
`examples/` where it binds nobody.

The baseline has to stay comparable. Nine intents were measured against the seventeen, so this
repository must keep running them — which is why adopting the candidate here is a non-goal and why
criterion 10 asserts `.xeno/config/templates/` does not exist afterwards.

One intent, one branch, `Closes #257`, and the issue carries `wp20`.
