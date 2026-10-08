# Command reference

Every command, with its flags. This is the text `xeno --help` prints, and a test in
`cmd/xeno/main_test.go` fails if this block and the `usage` constant it comes from stop
being identical: the README used to carry its own copy of this list and had drifted by
six commands before #231, so there is one place a command name is written down and this
page is not it.

What each command is for is in the process definition's section 6, which fixes who does
what with which command and in which order. This page says only what the shapes are.

```
usage:
  xeno init           --host HOST [--vendor] [--project OWNER/REPO] [--model ID] [--language TAG]
  xeno intent start   --for ISSUE [--intent KEY]        writes intent.yaml
  xeno phase start    --intent KEY --phase NN [--evidence-from DIR] [--export]
  xeno phase finish   --intent KEY --phase NN [--summary PATH|-]   writes digest.md
  xeno gate run       --intent KEY --phase NN [--base REF --head REF] [--evidence-from DIR]
  xeno gate approve   FINDING --intent KEY --phase NN --by WHO --reason TEXT
  xeno gate override  FINDING --intent KEY --phase NN --by WHO --reason TEXT
  xeno obligation close FINDING --intent KEY --phase NN
  xeno learning record --intent KEY [--phase NN] --category C --observation T --proposal T --target P
  xeno learning record --intent KEY [--phase NN] --no-finding
  xeno learning propose --out DIR [--intent KEY]   generates the merge request against the rule set
  xeno question record --intent KEY --phase NN [--file PATH]   reads the entry on stdin
  xeno decision record --intent KEY --phase NN --chosen TEXT --reason TEXT --by WHO [--resolves KEY] [--proposed-by WHO]
  xeno decision record --intent KEY --phase NN --withdraw --resolves KEY --reason TEXT --by WHO
  xeno review answer  RULE --intent KEY --result R [--note TEXT]   answers one review rule
  xeno review lens    --intent KEY --result R --note TEXT   a lens's entry, which answers no rule
  xeno assumption record --intent KEY --phase NN --text TEXT --origin WHERE --confidence HOW [--resolves KEY]
  xeno assumption confirm ID --intent KEY --by WHO
  xeno assumption reject  ID --intent KEY --by WHO
  xeno gate verify    [--intent KEY]            recompute and compare, write nothing (CI)
  xeno intent verify  --base REF --head REF     every intent the range touches is finished or closed (CI)
  xeno enforcement check [--branch NAME]        ask the host what it enforces (needs the network)
  xeno report verdict --intent KEY [--phase NN] [--dry-run]   the verdict onto the issue (CI, needs the network)
  xeno report figures [--intent KEY] [--all]    cost, artifacts and reopens per intent, read from the trail
  xeno evidence declare --intent KEY --phase NN --kind K [--job J] [--result R]
                        [--file PATH | --uri URL --sha256 HEX] [--produced-by CMD] [--format F]
  xeno evidence attach --intent KEY --phase NN --from DIR
  xeno intent status  [--intent KEY] [--all]    without one, the last ten by creation
  xeno intent close   --intent KEY --reason TEXT
  xeno section set    SECTION --intent KEY --phase NN [--file PATH]   reads stdin without --file
  xeno scope set      --intent KEY [--file PATH]   P0's context scope, read from stdin
  xeno template show  --phase NN                the sections a phase owes, in order
  xeno symbol show    NAME                      where a name is defined, from the project's index
  xeno check commit-message [--pattern NAME] [--file PATH]   reads stdin without --file
  xeno cost turn                                reads a hook's JSON on stdin
  xeno mcp            [--root DIR]              the six process operations over stdio
  xeno version
common: --root DIR (default .), --no-next to leave out the next step
        --tool-version V on section set and phase finish, which record what wrote a phase;
        XENO_HARNESS_VERSION says the same thing for a whole session
```

## Notes on particular commands

What the shapes above do not say. These moved here from the README when #231 cut it
back, and they are about behaviour rather than about this repository's own trail.

The gate path makes no network call, which is the property the verdicts rest on.
`xeno enforcement check` is the one command that does, because asking a host what it
enforces is the one question the repository cannot answer about itself.

`xeno intent start` takes the one thing it cannot derive, which is the issue the intent
belongs to: a key, `176`, or as much of the qualified id as the configuration does not
already hold, up to the whole `github.com/triplem/xeno#176`. The key continues the
sequence the intents directory holds, and `--intent` names it instead where there is no
sequence to continue or the key is not the next one.

`xeno learning record` writes the record section 10 owes at the end of every phase, and
without `--phase` the one an intent owes when it closes. The four keys are the agent's
and the header is the runner's; `--no-finding` is the empty record said rather than left
out, and a second call appends. It was the last artifact of this process that no command
wrote (A82).

`xeno intent status` without an intent lists the last ten by creation, `--all` every one
of them. `xeno cost turn` reads a hook's JSON on stdin and attributes the turn to the
phase the run marker says is open. Every command takes `--root DIR`, which defaults to
the working directory.

Every command that changes something ends by naming the next step of the working
sequence, and so does `xeno intent status`. `--no-next` leaves it out; the commands a
pipeline or a git hook runs never print it.
