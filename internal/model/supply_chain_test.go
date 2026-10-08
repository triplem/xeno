// SPDX-License-Identifier: Apache-2.0

package model

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// docs/supply-chain.md's second table is a copy of every pin in this repository, and the
// workflows point at it — "Every action is pinned to a commit sha and every tool to an exact
// version; docs/supply-chain.md says which and why" — so a reader who follows that sentence
// is told what the page says and not what the pipeline runs. Nothing held the two together,
// and the page had already drifted: `actions/setup-node` was pinned in two workflows and on
// no row of the table (#317).
//
// The two tests below read the pins out of the tree and the rows out of the page and compare
// them in both directions. The second direction is the one a reader cannot run: a row left
// behind after a tool is dropped names a pin nothing fetches any more, and it reads exactly
// like a row for a tool still in use.
//
// They live in this package because its tests are already the ones that read the repository's
// build files — TestTheReleaseStampsTheVariablesThisPackageDeclares reads release.yml and
// TestTheDockerfileCopiesTheBinaryTheReleaseBuilds reads the Dockerfile — so the tests over
// those files stay in one place rather than opening a package whose only content is this one.

// pinnedIn are the files a pin of the pipeline is written in. Versions are read from the
// workflows alone: the Dockerfile's `ARG VERSION=0.0.0-dev` is this project's own version and
// not something it fetches, and go.mod's `require` line is the first table's subject rather
// than this one's. What each file contributes is therefore named per kind below.
func pinnedIn(t *testing.T) (workflows map[string]string, dockerfile, gomod string) {
	t.Helper()
	root := filepath.Join("..", "..")
	read := func(rel string) string {
		b, err := os.ReadFile(filepath.Join(root, rel))
		if err != nil {
			t.Fatal(err)
		}
		return string(b)
	}
	entries, err := os.ReadDir(filepath.Join(root, ".github", "workflows"))
	if err != nil {
		t.Fatal(err)
	}
	workflows = map[string]string{}
	for _, e := range entries {
		if !e.IsDir() && strings.HasSuffix(e.Name(), ".yml") {
			rel := filepath.Join(".github", "workflows", e.Name())
			workflows[rel] = read(rel)
		}
	}
	// A directory that has stopped holding workflows would leave every assertion below
	// vacuously true, which is the one failure a comparison cannot report about itself.
	if len(workflows) == 0 {
		t.Fatal("no workflow was read, so there is nothing to hold the page to")
	}
	return workflows, read("Dockerfile"), read("go.mod")
}

var (
	// An action pin as a workflow writes it: the sha, and the tag that sha stood for in the
	// trailing comment. Both are compared, because the sha is what runs and the tag is what
	// the page's second column prints beside it.
	usesPin = regexp.MustCompile(`uses:\s*([A-Za-z0-9][A-Za-z0-9._/-]*)@([0-9a-f]{40})[ \t]*(?:#[ \t]*(v?[0-9][0-9A-Za-z.+-]*))?`)
	// An image pinned by digest, which is how the image's base is written. Read from the
	// Dockerfile, where it arrives as the value of ARG BASE, and from any workflow that pins
	// one the same way.
	digestPin = regexp.MustCompile(`([A-Za-z0-9][A-Za-z0-9._/-]*)@sha256:([0-9a-f]{64})`)
	// A sha256 with no reference in front of it: the gitleaks tarball's checksum, which is a
	// pin of the same kind written without the algorithm prefix.
	checksumPin = regexp.MustCompile(`(?:^|[^:0-9a-f])([0-9a-f]{64})(?:[^0-9a-f]|$)`)
	// An exact version, which is the shape every tool pin in these workflows is written in. A
	// major on its own is deliberately not matched: `node-version: '24'` and `go 1.27` cannot
	// be told from a port or a count by their shape, so the two toolchains are bound by name
	// instead, in boundByHand below.
	versionPin = regexp.MustCompile(`\bv?[0-9]+\.[0-9]+\.[0-9]+\b`)
	// What the page writes where a whole digest would be unreadable: the algorithm or nothing,
	// eight hex digits and an ellipsis. The comparison is against the page's own form rather
	// than against a form the page would have to be rewritten into, so what is checked is what
	// a reader sees; eight hex digits is thirty-two bits of the digest, which is enough to
	// bind a row to one artifact and not enough to be mistaken for prose.
	abbreviated = regexp.MustCompile(`(?:sha256:)?\b([0-9a-f]{8})\x{2026}`)
	fullSha     = regexp.MustCompile(`\b[0-9a-f]{40}\b`)
	goDirective = regexp.MustCompile(`(?m)^go ([0-9]+\.[0-9]+)`)
	nodeVersion = regexp.MustCompile(`node-version:\s*'?([0-9][0-9.]*)'?`)
)

// actionPin is one sha-pinned `uses:` together with every file that writes it, so that two
// workflows disagreeing about one action is reported as the drift it is rather than as two
// pins the page is asked to carry both of.
type actionPin struct {
	sha, tag string
	where    []string
}

// treePins is everything the tree pins, found by the shape each pin is written in rather than
// by looking up a list of tools. A list would be a third copy of the same pins and would go
// stale the way the page did; a shape finds the pin somebody adds tomorrow.
type treePins struct {
	actions   map[string]*actionPin
	digests   map[string]string // the whole digest -> the file it is written in
	checksums map[string]string // the whole sha256 -> the file it is written in
	versions  map[string]string // the version as written -> the file it is written in
	goVersion string
	node      string
	raw       string
}

func readTreePins(t *testing.T) *treePins {
	t.Helper()
	workflows, dockerfile, gomod := pinnedIn(t)
	p := &treePins{
		actions:   map[string]*actionPin{},
		digests:   map[string]string{},
		checksums: map[string]string{},
		versions:  map[string]string{},
	}
	node := map[string]string{}

	for _, rel := range sortedKeys(workflows) {
		body := workflows[rel]
		p.raw += body
		// Shas and digests are read from the raw text, because their shapes are long enough
		// that a match is the pin itself and not a paragraph about it, and because the tag
		// beside a `uses:` sha is written as a comment. Versions are read from the text with
		// the comments removed: a version in a comment is something a paragraph says, and a
		// version in a value is something a step runs.
		for _, m := range usesPin.FindAllStringSubmatch(body, -1) {
			name, sha, tag := m[1], m[2], m[3]
			if a, ok := p.actions[name]; ok {
				if a.sha != sha {
					t.Errorf("%s pins %s at %s and %s pins it at %s, so one row cannot be right about both",
						rel, name, sha, strings.Join(a.where, ", "), a.sha)
				}
				a.where = append(a.where, rel)
				continue
			}
			p.actions[name] = &actionPin{sha: sha, tag: tag, where: []string{rel}}
		}
		stripped := withoutComments(body)
		for _, m := range digestPin.FindAllStringSubmatch(stripped, -1) {
			p.digests[m[2]] = rel
		}
		for _, m := range checksumPin.FindAllStringSubmatch(stripped, -1) {
			if _, ok := p.digests[m[1]]; !ok {
				p.checksums[m[1]] = rel
			}
		}
		for _, v := range versionPin.FindAllString(stripped, -1) {
			p.versions[v] = rel
		}
		for _, m := range nodeVersion.FindAllStringSubmatch(stripped, -1) {
			node[m[1]] = rel
		}
	}

	// The Dockerfile contributes the one pin the issue names it for: the base image's digest,
	// which #305 put on the page beside the ARG and which nothing stopped the next bump moving
	// on one side only.
	p.raw += dockerfile
	for _, m := range digestPin.FindAllStringSubmatch(withoutComments(dockerfile), -1) {
		p.digests[m[2]] = "Dockerfile"
	}

	// go.mod contributes the `go` directive and nothing else. The Go toolchain's row is that
	// directive — setup-go reads it rather than taking a version of its own — so a page held
	// only to the workflows could not check that row at all.
	if m := goDirective.FindStringSubmatch(gomod); m != nil {
		p.goVersion = m[1]
	}

	// Two workflows asking for different node majors would make the single row a claim about
	// neither, so that is reported here rather than compared against the page twice.
	if len(node) > 1 {
		t.Errorf("the workflows ask for node %v and the page carries one row", sortedKeys(node))
	}
	for v := range node {
		p.node = v
	}
	return p
}

// withoutComments removes YAML and Dockerfile comments, which begin at a `#` that opens a line
// or follows whitespace. A `#` anywhere else belongs to a value — `${target#*/}` in the
// release's build loop is one — and cutting there would hide the rest of that line, which is
// how a comparison comes to pass over a pin it never read.
func withoutComments(body string) string {
	var out strings.Builder
	for _, line := range strings.Split(body, "\n") {
		for i := 0; i < len(line); i++ {
			if line[i] == '#' && (i == 0 || line[i-1] == ' ' || line[i-1] == '\t') {
				line = line[:i]
				break
			}
		}
		out.WriteString(line)
		out.WriteString("\n")
	}
	return out.String()
}

// tableRow is one line of the page's pipeline table: what is pinned, and what it is pinned to.
// The third column says where it is fetched from, which is a question about egress and not
// about drift, so nothing here reads it.
type tableRow struct {
	what, pinnedTo string
	line           int
}

// pipelineTable is the table under "In the pipeline" and not the one above it, which lists what
// ends up in the binary and is a different question. It is read as markdown, because the page is
// prose for a reader before it is data for this test.
func pipelineTable(t *testing.T) []tableRow {
	t.Helper()
	b, err := os.ReadFile(filepath.Join("..", "..", "docs", "supply-chain.md"))
	if err != nil {
		t.Fatal(err)
	}
	var rows []tableRow
	inSection := false
	for i, line := range strings.Split(string(b), "\n") {
		if strings.HasPrefix(line, "## ") {
			if len(rows) > 0 {
				break
			}
			inSection = strings.TrimSpace(line) == "## In the pipeline"
			continue
		}
		if !inSection {
			continue
		}
		if !strings.HasPrefix(line, "|") {
			// A blank line ends the table. The paragraphs below it name the same pins in
			// prose, and a paragraph is not a row.
			if len(rows) > 0 {
				break
			}
			continue
		}
		cells := strings.Split(strings.Trim(strings.TrimSpace(line), "|"), "|")
		if len(cells) < 2 {
			continue
		}
		what := strings.TrimSpace(cells[0])
		if what == "What" || strings.Trim(what, "-: ") == "" {
			continue
		}
		rows = append(rows, tableRow{what: what, pinnedTo: strings.TrimSpace(cells[1]), line: i + 1})
	}
	if len(rows) == 0 {
		t.Fatal("docs/supply-chain.md holds no pipeline table, so there is nothing to compare against")
	}
	return rows
}

// exempt names the rows that carry no pin any file in this repository can be compared against,
// each with the reason. They are held to existing and to nothing else: a tool that is dropped
// takes its row with it, which is all these rows can be asked. Naming them is the point rather
// than the loophole — the same thing the walk in cmd/xeno/sealed_test.go does for a command
// that may move a hash — and an entry nothing needs fails below, so an exemption cannot outlive
// the row it was written for.
var exempt = map[string]string{
	"trivy's vulnerability database": "not pinned, and cannot be: a scan answers what is known " +
		"today, so the database is fetched on every run and there is no version in a workflow " +
		"for the page to agree with",
	"govulncheck's vulnerability database": "not pinned, and cannot be, for the reason trivy's " +
		"is not: it is fetched on every run so that the scan answers what is known today, and " +
		"the report records the date it was built rather than a version a workflow could carry",
	"`ca-certificates` and `git`, installed into the image": "not pinned, and cannot be: apt-get " +
		"resolves both against Debian's suite on the day the release is cut, and the page's own " +
		"paragraph says why pinning them would turn each move of that suite into a failed release",
	"`docker`, which builds and pushes the image": "pinned by the hosted runner's image, which " +
		"this repository neither pins nor can read; the row exists in order to say that the pin " +
		"is somebody else's",
}

// boundByHand names the rows whose pin is a major version that a setup action resolves to a
// patch. Neither `go 1.27` nor `node-version: '24'` has a shape that can be told from a port or
// a count, so each is read out of the one key that holds it and compared here by name. These
// rows are checked rather than exempt; what they cannot do is arrive through the shapes above.
var boundByHand = map[string]func(*treePins) (string, string){
	"Go toolchain": func(p *treePins) (string, string) {
		return p.goVersion, "the `go` directive of go.mod"
	},
	"Node toolchain": func(p *treePins) (string, string) {
		return p.node, "`node-version` in the workflows"
	},
}

// TestEveryPinInTheTreeIsOnTheSupplyChainPage is the direction that was missing when
// `actions/setup-node` reached two workflows and no row. A pin the pipeline runs and the page
// does not carry makes the sentence that sends a reader to the page the one that is wrong.
func TestEveryPinInTheTreeIsOnTheSupplyChainPage(t *testing.T) {
	p := readTreePins(t)
	rows := pipelineTable(t)
	table := func() string {
		var b strings.Builder
		for _, r := range rows {
			fmt.Fprintf(&b, "  %s | %s\n", r.what, r.pinnedTo)
		}
		return b.String()
	}

	for _, name := range sortedKeys(p.actions) {
		a := p.actions[name]
		var found *tableRow
		for i := range rows {
			if strings.Contains(rows[i].what, name) {
				found = &rows[i]
				break
			}
		}
		if found == nil {
			t.Errorf("%s is pinned in %s and no row of the page names it:\n%s",
				name, strings.Join(a.where, ", "), table())
			continue
		}
		if !strings.Contains(found.pinnedTo, a.sha) {
			t.Errorf("%s runs at %s in %s and the page's row says %q",
				name, a.sha, strings.Join(a.where, ", "), found.pinnedTo)
		}
		// The tag beside the sha. It is what a reader compares when deciding whether a pin is
		// current, and it is the half that can be wrong while the sha is right.
		if a.tag != "" && versionPin.MatchString(found.pinnedTo) && !strings.Contains(found.pinnedTo, a.tag) {
			t.Errorf("%s is pinned as %s in %s and the page's row says %q",
				name, a.tag, strings.Join(a.where, ", "), found.pinnedTo)
		}
	}

	// Digests and checksums are matched by value across the whole table rather than against the
	// name beside them: a digest is unique, and a row's name is the page's and not the tree's.
	for _, kind := range []struct {
		label string
		pins  map[string]string
	}{{"digest", p.digests}, {"checksum", p.checksums}} {
		for _, sum := range sortedKeys(kind.pins) {
			if !carried(rows, sum) {
				t.Errorf("the %s %s is pinned in %s and no row of the page carries it, whole or "+
					"abbreviated:\n%s", kind.label, sum, kind.pins[sum], table())
			}
		}
	}

	for _, v := range sortedKeys(p.versions) {
		if !carried(rows, v) {
			t.Errorf("%s is the version pinned in %s and no row of the page carries it:\n%s",
				v, p.versions[v], table())
		}
	}

	for _, what := range sortedKeys(boundByHand) {
		got, source := boundByHand[what](p)
		if got == "" {
			t.Errorf("nothing in the tree answers %s, which the %q row is held to", source, what)
			continue
		}
		var found *tableRow
		for i := range rows {
			if rows[i].what == what {
				found = &rows[i]
				break
			}
		}
		if found == nil {
			t.Errorf("%s is %s and the page has no %q row:\n%s", source, got, what, table())
			continue
		}
		if !strings.Contains(found.pinnedTo, got) {
			t.Errorf("%s is %s and the page's %q row says %q", source, got, what, found.pinnedTo)
		}
	}
}

// TestEveryRowOfTheSupplyChainPageIsAPinTheTreeCarries is the direction a reader cannot run. A
// row left behind by a tool that is gone reads exactly like a row for a tool still in use, and
// the page is where somebody goes to ask what a release depends on.
func TestEveryRowOfTheSupplyChainPageIsAPinTheTreeCarries(t *testing.T) {
	p := readTreePins(t)
	rows := pipelineTable(t)
	seenExempt, seenByHand := map[string]bool{}, map[string]bool{}

	for _, r := range rows {
		if why, ok := exempt[r.what]; ok {
			seenExempt[r.what] = true
			if why == "" {
				t.Errorf("line %d: %s is exempt for no stated reason", r.line, r.what)
			}
			continue
		}
		if _, ok := boundByHand[r.what]; ok {
			seenByHand[r.what] = true
			continue
		}

		// Every identifier the row carries, each held to the tree by its own shape. The anchor
		// counted here is what catches a dropped tool: a row for something nothing fetches any
		// more carries no identifier the tree has, and a row that names none at all is either
		// written in a form this test cannot read or is not a pin.
		anchors := 0
		for _, sha := range fullSha.FindAllString(r.pinnedTo, -1) {
			anchors++
			if !pinsSha(p.actions, sha) {
				t.Errorf("line %d: %s is pinned to %s and no workflow uses that sha",
					r.line, r.what, sha)
			}
		}
		for _, m := range abbreviated.FindAllStringSubmatch(r.pinnedTo, -1) {
			anchors++
			if !hasPrefixIn(p.digests, m[1]) && !hasPrefixIn(p.checksums, m[1]) {
				t.Errorf("line %d: %s is pinned to %s… and nothing in the workflows or the "+
					"Dockerfile carries a digest beginning there", r.line, r.what, m[1])
			}
		}
		for _, v := range versionPin.FindAllString(r.pinnedTo, -1) {
			anchors++
			// Against the raw text and not the stripped text, and this direction is the weaker
			// of the two on purpose. A tag is written in the tree as the trailing comment of a
			// `uses:` line or as a `version:` input beside one, so what can be asserted is that
			// the number on the page is a number the tree says somewhere.
			// The other direction is where a version that runs is held to being printed here.
			if !strings.Contains(p.raw, strings.TrimPrefix(v, "v")) {
				t.Errorf("line %d: %s is pinned to %s and that version appears nowhere in the "+
					"workflows or the Dockerfile", r.line, r.what, v)
			}
		}
		if anchors == 0 {
			t.Errorf("line %d: the row %q names nothing that can be compared against the tree. "+
				"Either its pin is written in a form this test does not read, or it is a row left "+
				"behind by a tool that is gone; a row that genuinely cannot be compared belongs in "+
				"exempt, with the reason", r.line, r.what)
		}
	}

	// An exemption nothing needs is one the next reader will trust, so both maps are held to
	// naming a row that exists and a dropped tool takes its entry here with it.
	for _, what := range sortedKeys(exempt) {
		if !seenExempt[what] {
			t.Errorf("the page has no row %q and this test exempts one: %s", what, exempt[what])
		}
	}
	for _, what := range sortedKeys(boundByHand) {
		if !seenByHand[what] {
			t.Errorf("the page has no row %q and this test binds one by hand", what)
		}
	}
}

// carried answers whether any row names the value, whole or in the page's abbreviated form.
func carried(rows []tableRow, value string) bool {
	for _, r := range rows {
		if strings.Contains(r.pinnedTo, value) {
			return true
		}
		for _, m := range abbreviated.FindAllStringSubmatch(r.pinnedTo, -1) {
			if strings.HasPrefix(value, m[1]) {
				return true
			}
		}
	}
	return false
}

func pinsSha(actions map[string]*actionPin, sha string) bool {
	for _, a := range actions {
		if a.sha == sha {
			return true
		}
	}
	return false
}

func hasPrefixIn(pins map[string]string, prefix string) bool {
	for sum := range pins {
		if strings.HasPrefix(sum, prefix) {
			return true
		}
	}
	return false
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
