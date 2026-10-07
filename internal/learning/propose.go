// SPDX-License-Identifier: Apache-2.0

// Package learning carries the proposal of a learning record to the rule set, which is the
// route section 10 describes and the one clause of WP14 that had never been exercised: a
// learning takes effect after review and not on being noticed.
//
// It generates a merge request and does not open one. WP12 fixes the tracker adapter at four
// operations, and creating a merge request is not one of them, so what this produces is what a
// person pushes and opens: the rule files at the paths they take in the tree, a patch that
// creates them, and a description saying which records were carried and which were not. The
// bundle goes to a directory the caller names and never into the repository it read, because
// an edit made where a learning was noticed is the one thing section 10 forbids.
//
// Every record read is listed with what became of it. That is the whole of #289's complaint:
// a record says what was proposed and nothing said whether anybody had read it, so a route
// that carried the easy ones and dropped the rest silently would reproduce the defect while
// looking like a fix.
//
// Nothing here is an artifact. Section 5 enumerates what the trail carries and a merge request
// is not in it, so the bundle invents no field: it is a patch, a description and the rule files
// section 9 already defines.
package learning

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"go.yaml.in/yaml/v3"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/rules"
)

// ProjectRules is where section 10 says a rule the engine applies is named: the project level
// of the learned tree. It is composed from the rules package's own constants rather than
// written out, so that a tree moved there is moved here.
const ProjectRules = rules.ConfigDir + "/" + rules.Learned + "/project"

// Disposition is what the route did with one entry. Three values and no fourth: either the
// target resolves to a rule file and the bundle holds it, or section 10's one exception
// applies and a person words the line, or it resolves to neither and is listed with the
// reason.
type Disposition string

const (
	Carried    Disposition = "carried"
	ForWording Disposition = "for-wording"
	Unresolved Disposition = "unresolved"
)

// Item is one learning entry with where it was found and what became of it. The provenance is
// three fields rather than one string because the description names the record by its path and
// its position in it, and a reader looking for the entry has to be able to open the file.
type Item struct {
	Intent string
	Phase  string // empty for the record an intent writes when it closes
	Path   string // relative to the repository root
	Index  int    // the entry's position in its record, counted from one; zero where the record itself is the finding
	Entry  model.LearningEntry

	Disposition Disposition
	// Reason is what the ledger prints beside the record: the resolved file for a proposal
	// that was carried or is for a person to word, and the cause for one that was not. One
	// field for all three because they answer the one question a reader asks of a line in the
	// ledger, which is where this proposal went.
	Reason string
	Rule   string // the id of the rule it was carried into, where it was carried
}

// File is one generated file, at the path it takes in the repository rather than in the output
// directory, because that is the path the patch has to carry and the path a reviewer reads.
type File struct {
	Path    string
	Content string
}

// Bundle is the merge request before it is a merge request. Records and NoFinding count the
// population the figures are taken from: a count of the trail names what it counted, or the
// reader cannot tell a route that read twelve records from one that read five hundred.
type Bundle struct {
	Root      string
	Records   int
	NoFinding int
	Items     []Item
	Files     []File

	Description string
	Patch       string
}

// Count is how many items took one disposition. The three counts and the entry total are the
// same number read two ways, which a test asserts: an item that fell out of all three is an
// item the route dropped.
func (b *Bundle) Count(d Disposition) int {
	n := 0
	for _, it := range b.Items {
		if it.Disposition == d {
			n++
		}
	}
	return n
}

// Plan reads the records of the given intents, in the order the keys are given, and builds the
// bundle without writing anything. Reading and writing are separate so that a caller can show
// what a route would do before it does it, and so that the writing can refuse a destination
// without the reading having to be undone.
//
// The keys are passed in rather than read here, because the order they are listed in is the
// order the intents were created, which is a question about intent.yaml and belongs to the
// caller that already answers it.
func Plan(root string, keys []string) *Bundle {
	b := &Bundle{Root: root}
	for _, key := range keys {
		for _, rel := range recordPaths(key) {
			b.read(root, key, rel)
		}
	}
	for i := range b.Items {
		classify(root, &b.Items[i])
	}
	b.generate(root)
	b.Description = b.describe()
	b.Patch = b.patch()
	return b
}

// recordPaths is every place section 10 puts a record for one intent, in the order a reader
// walks them: the phases in their sequence, and then the record an abandoned intent writes
// when it closes, which comes after the phases it closed on.
func recordPaths(key string) []string {
	var out []string
	for _, p := range model.Phases {
		out = append(out, model.PhaseDir(key, p)+"/learning.yaml")
	}
	return append(out, model.IntentDir(key)+"/learning.yaml")
}

// read turns one record into items. A record that is not there is the ordinary case for a
// phase that has not run and for every merged intent's closing record, so it is not reported;
// a record that is there and cannot be read is reported as an item of its own, because the
// alternative is a file full of proposals disappearing from the ledger with the error.
func (b *Bundle) read(root, key, rel string) {
	var rec model.Learning
	err := fm.ReadYAML(filepath.Join(root, rel), &rec)
	if os.IsNotExist(err) {
		return
	}
	phase := phaseOf(rel)
	b.Records++
	if err != nil {
		b.Items = append(b.Items, Item{Intent: key, Phase: phase, Path: rel,
			Disposition: Unresolved, Reason: "the record cannot be read: " + err.Error()})
		return
	}
	if len(rec.Learnings) == 0 {
		b.NoFinding++
		return
	}
	for i, e := range rec.Learnings {
		b.Items = append(b.Items, Item{Intent: key, Phase: phase, Path: rel, Index: i + 1, Entry: e})
	}
}

func phaseOf(rel string) string {
	for _, p := range model.Phases {
		if strings.Contains(rel, "/phases/"+p+"/") {
			return p
		}
	}
	return ""
}

// classify decides the route of one entry against section 10, which fixes two shapes of
// target and leaves nothing for a third.
//
// A rule the engine applies names a path under the project level of the learned tree. Until
// that engine exists, a learning whose category is project-convention may instead name the
// file a project already reads before it acts, and that is the one exception: the same latitude
// for the other three categories would make target mean "where this belongs" again, which is
// what the 51 prose targets of #289 already were.
//
// A target it cannot place is listed rather than refused. Refusing would stop the route on the
// first record of a trail that has hundreds, and a trail nobody can run the route over is the
// state this started in; so the bundle is generated from what resolves and the rest is named,
// with the reason, in the description a person reads before pushing anything.
func classify(root string, it *Item) {
	if it.Disposition != "" {
		return
	}
	t := strings.TrimSpace(it.Entry.Target)
	unresolved := func(reason string) {
		it.Disposition, it.Reason = Unresolved, reason
	}
	switch {
	case t == "":
		unresolved("the entry names no target")
		return
	case strings.ContainsAny(t, " ,"):
		unresolved("the target reads as prose rather than as a path, so no file follows from it")
		return
	}
	clean := path.Clean(filepath.ToSlash(t))
	if dir, file := path.Split(clean); path.Clean(dir) == ProjectRules {
		switch {
		case !strings.HasSuffix(file, ".yaml") && !strings.HasSuffix(file, ".yml"):
			unresolved("the target lies in the rule tree but is not a .yaml file, which is the only thing read there")
		case strings.TrimSuffix(strings.TrimSuffix(file, ".yaml"), ".yml") == "":
			unresolved("the target names no file, only the directory rules live in")
		default:
			it.Disposition = Carried
		}
		return
	}
	if strings.HasPrefix(clean, rules.ConfigDir+"/") || strings.HasPrefix(clean, rules.PluginDir+"/") {
		unresolved("the target lies in the rule tree outside " + ProjectRules +
			", which is the one level section 10 says a rule the engine applies is named at")
		return
	}
	info, err := os.Stat(filepath.Join(root, filepath.FromSlash(clean)))
	switch {
	case err != nil:
		unresolved("the target is no path in the tree and names no rule under " + ProjectRules)
	case info.IsDir():
		unresolved("the target names a directory, and a merge request changes a file")
	case it.Entry.Category != "project-convention":
		unresolved("a " + it.Entry.Category + " learning names a file outside the rule tree; " +
			"section 10 allows that for project-convention only")
	default:
		it.Disposition, it.Reason = ForWording, clean
	}
}

// generate writes one rule file per distinct target, in the order the targets were first
// named. Several entries naming one file is the ordinary case for a convention noticed twice,
// and their proposals are carried into one statement rather than one file overwriting the
// other, which is how a route loses a record while reporting it as carried.
func (b *Bundle) generate(root string) {
	var order []string
	byTarget := map[string][]int{}
	for i, it := range b.Items {
		if it.Disposition != Carried {
			continue
		}
		t := path.Clean(filepath.ToSlash(strings.TrimSpace(it.Entry.Target)))
		if _, seen := byTarget[t]; !seen {
			order = append(order, t)
		}
		byTarget[t] = append(byTarget[t], i)
	}
	for _, t := range order {
		idx := byTarget[t]
		id := strings.TrimSuffix(strings.TrimSuffix(path.Base(t), ".yaml"), ".yml")
		for _, i := range idx {
			b.Items[i].Rule = id
			b.Items[i].Reason = t
		}
		b.Files = append(b.Files, File{Path: t, Content: b.rule(root, t, id, idx)})
	}
}

// rule is one generated rule file, as section 9's format and as text rather than as marshalled
// YAML: the comment saying which records it came from is the thing a reviewer reads first, and
// a marshaller drops comments.
//
// kind is review throughout. A proposal is a sentence somebody wrote about their own phase, and
// turning one into a predicate would mean choosing a predicate type nobody proposed; section 9
// says a review rule produces an item in the P5 checklist that a person answers, which is what
// a sentence can be held against. A reviewer who sees a check in it writes the check.
func (b *Bundle) rule(root, target, id string, idx []int) string {
	var s strings.Builder
	s.WriteString("# Proposed by the learning route, from the records named below. Section 10: a learning\n")
	s.WriteString("# takes effect after review and not on being noticed, so this file is reviewed as a merge\n")
	s.WriteString("# and was never written where it was noticed.\n#\n")
	for _, i := range idx {
		it := b.Items[i]
		s.WriteString(fmt.Sprintf("# %s entry %d, category %s\n", it.Path, it.Index, it.Entry.Category))
	}
	fmt.Fprintf(&s, "id: %s\n", id)
	fmt.Fprintf(&s, "version: %d\n", version(root, target))
	s.WriteString("scope: project\n")
	s.WriteString("kind: review\n")
	fmt.Fprintf(&s, "applies_to: [%s]\n", strings.Join(phases(b.Items, idx), ", "))
	s.WriteString("statement: >\n")
	var proposals []string
	for _, i := range idx {
		proposals = append(proposals, strings.Join(strings.Fields(b.Items[i].Entry.Proposal), " "))
	}
	for _, line := range wrap(strings.Join(proposals, " "), 94) {
		s.WriteString("  " + line + "\n")
	}
	return s.String()
}

// version counts changes to the file and claims nothing about compatibility, so a file the
// route has not seen before starts at 1 and one already in the tree is proposed at one more
// than it carries. The tree is read and not written: the route's own destination is the output
// directory, and this is the one place it looks at the real path at all.
func version(root, target string) int {
	var on struct {
		Version int `yaml:"version"`
	}
	b, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(target)))
	if err != nil || yaml.Unmarshal(b, &on) != nil || on.Version < 1 {
		return 1
	}
	return on.Version + 1
}

// phases is applies_to, which section 9 requires and will not take empty. The phases are the
// ones the contributing records were written in, in process order.
//
// An entry from the record an intent writes when it closes belongs to no phase, so it proposes
// the rule for the whole sequence and the reviewer narrows it. The alternative would be to drop
// such an entry or to guess a phase for it, and a rule applying too widely is visible in the
// checklist, where somebody answers it and says so.
func phases(items []Item, idx []int) []string {
	have := map[string]bool{}
	for _, i := range idx {
		if items[i].Phase == "" {
			for _, p := range model.Phases {
				have[p] = true
			}
			continue
		}
		have[items[i].Phase] = true
	}
	var out []string
	for _, p := range model.Phases {
		if have[p] {
			out = append(out, p)
		}
	}
	return out
}

// wrap breaks a statement into lines of at most n characters, on word boundaries. A folded
// scalar reads as one paragraph whatever the breaks are, and the shipped rules are wrapped, so
// a generated file that was not would be the one file in the tree nobody can read in a diff.
func wrap(s string, n int) []string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return []string{""}
	}
	var out []string
	line := words[0]
	for _, w := range words[1:] {
		if len(line)+1+len(w) > n {
			out = append(out, line)
			line = w
			continue
		}
		line += " " + w
	}
	return append(out, line)
}

// Write puts the bundle in a directory of its own and returns what the rules package makes of
// the files it generated.
//
// The generated tree is loaded back as a rule tree, because a route that produced a file
// G-Rules rejects would have carried a record to a rule set that will not have it, and the
// place to find that out is here rather than in the reviewer's first gate run.
//
// It refuses a destination inside the trail it read. Section 10's route never edits where the
// learning was noticed, and an output directory pointed at .xeno would be exactly that edit
// with a flag in front of it.
func (b *Bundle) Write(out string) ([]rules.Problem, error) {
	if strings.TrimSpace(out) == "" {
		return nil, fmt.Errorf("a destination directory is required: the bundle is what somebody pushes, " +
			"and it is not written into the tree it was read from")
	}
	absOut, err := filepath.Abs(out)
	if err != nil {
		return nil, err
	}
	absXeno, err := filepath.Abs(filepath.Join(b.Root, ".xeno"))
	if err != nil {
		return nil, err
	}
	if absOut == absXeno || strings.HasPrefix(absOut, absXeno+string(filepath.Separator)) {
		return nil, fmt.Errorf("%s lies inside %s: a learning takes effect after review, so the route "+
			"writes a merge request somewhere else and never into the trail it read", out, absXeno)
	}
	for _, f := range b.Files {
		p := filepath.Join(absOut, filepath.FromSlash(f.Path))
		if err := os.MkdirAll(filepath.Dir(p), 0o755); err != nil {
			return nil, err
		}
		if err := os.WriteFile(p, []byte(f.Content), 0o644); err != nil {
			return nil, err
		}
	}
	if err := os.MkdirAll(absOut, 0o755); err != nil {
		return nil, err
	}
	for name, content := range map[string]string{
		DescriptionFile: b.Description,
		PatchFile:       b.Patch,
	} {
		if err := os.WriteFile(filepath.Join(absOut, name), []byte(content), 0o644); err != nil {
			return nil, err
		}
	}
	loaded, problems := rules.Load(absOut)
	_, more := rules.Effective(loaded)
	return append(problems, more...), nil
}

// What a reader of the output directory finds beside the rule files. The description is the
// merge request's text and the patch is the change; both are what a person pushes, which is
// why neither is written into the repository.
const (
	DescriptionFile = "merge-request.md"
	PatchFile       = "learned-rules.patch"
)

// patch is a unified diff creating the generated files, in the form git apply reads. A patch
// rather than a branch: making a branch means writing to the repository the route is forbidden
// to write to, and a patch is reviewable as text before anything has happened at all.
func (b *Bundle) patch() string {
	var s strings.Builder
	for _, f := range b.Files {
		lines := strings.Split(strings.TrimSuffix(f.Content, "\n"), "\n")
		fmt.Fprintf(&s, "diff --git a/%s b/%s\n", f.Path, f.Path)
		s.WriteString("new file mode 100644\n")
		fmt.Fprintf(&s, "--- /dev/null\n+++ b/%s\n", f.Path)
		fmt.Fprintf(&s, "@@ -0,0 +1,%d @@\n", len(lines))
		for _, l := range lines {
			s.WriteString("+" + l + "\n")
		}
	}
	return s.String()
}

// describe is the merge request's description: what it carries, what it does not, and in both
// cases which record. The unresolved list is the longer half and is deliberately not an
// appendix, because the figure that matters about this trail is how many of its proposals name
// something a tool can act on.
func (b *Bundle) describe() string {
	var s strings.Builder
	s.WriteString("# Learned rules proposed from the trail\n\n")
	for _, p := range []string{
		fmt.Sprintf("Generated by `xeno learning propose` from %d learning records, of which %d state "+
			"no finding and %d carry %d proposals between them. Section 10 routes a learning through a "+
			"merge request against the rule set, so that it takes effect after review and not on being "+
			"noticed; this is that merge request before anybody has pushed it.",
			b.Records, b.NoFinding, b.Records-b.NoFinding, len(b.Items)),
		fmt.Sprintf("Of the %d proposals, %d name a rule file and are carried below, %d are conventions "+
			"for a person to word into a file the project already reads, and %d name a target this route "+
			"cannot resolve to a file. Every one of the %d is listed. A record this route dropped in "+
			"silence would leave nobody able to say which proposals were read, which is the defect the "+
			"route exists to end.",
			len(b.Items), b.Count(Carried), b.Count(ForWording), b.Count(Unresolved), len(b.Items)),
		"Merging this adopts the rules below and nothing else. The proposals under the two lists " +
			"after them are not adopted by it: a convention is adopted by the wording a person writes " +
			"into the file it names, and an unresolved target is adopted by nothing until the record " +
			"names a file.",
	} {
		s.WriteString(para(p))
	}

	s.WriteString("## The rules this carries\n\n")
	if len(b.Files) == 0 {
		s.WriteString("None: no proposal in this trail names a file under `" + ProjectRules + "`.\n\n")
	} else {
		s.WriteString("| rule | phases | from |\n|---|---|---|\n")
		for _, f := range b.Files {
			var from []string
			id := ""
			for _, it := range b.Items {
				if it.Disposition == Carried && it.Reason == f.Path {
					id = it.Rule
					from = append(from, fmt.Sprintf("`%s` entry %d", it.Path, it.Index))
				}
			}
			fmt.Fprintf(&s, "| `%s` | %s | %s |\n", id, appliesOf(f.Content), strings.Join(from, ", "))
		}
		s.WriteString("\n")
	}

	s.WriteString("## Conventions for a person to word\n\n")
	s.WriteString(para("Section 10 lets a project-convention name the file a project already " +
		"reads before it acts, and says what decides whether it earns a line there: that file is sent with " +
		"every request of every session, so a line in it costs more than a rule the engine loads when it " +
		"applies. That is a judgement about a sentence and about a budget, so the route carries the " +
		"proposal here and writes no line."))
	s.WriteString(wordings(b))

	s.WriteString("## Proposals this route could not carry\n\n")
	s.WriteString(table(b, Unresolved, "why"))

	s.WriteString("### The same list by target\n\n")
	s.WriteString(para("How many proposals named each unresolved target. It is the figure a reader of "+
		"the list above wants next, because one prose target named forty times is one decision about one "+
		"habit and not forty separate ones.") + "```\n")
	for _, l := range b.Targets(Unresolved) {
		s.WriteString(l + "\n")
	}
	s.WriteString("```\n\n")

	s.WriteString("## Applying it\n\n```\ngit checkout -b learned-rules\ngit apply " + PatchFile +
		"\ngit add " + ProjectRules + "\n```\n\n")
	s.WriteString(para("Then commit, push and open the merge request with this file as its " +
		"description. The route does not open one: the tracker adapter has four operations and this is not " +
		"among them, and a route that opened its own merge request would be deciding the thing the review " +
		"is for."))
	return s.String()
}

// para is one paragraph of the description, wrapped as this project wraps prose and followed
// by the blank line that separates it from the next.
func para(s string) string { return strings.Join(wrap(s, 88), "\n") + "\n\n" }

// appliesOf reads applies_to back out of a generated file, so that the table and the file
// cannot disagree about which phases a rule is proposed for.
func appliesOf(content string) string {
	for _, l := range strings.Split(content, "\n") {
		if rest, ok := strings.CutPrefix(l, "applies_to: ["); ok {
			return strings.TrimSuffix(rest, "]")
		}
	}
	return ""
}

// table is one disposition's records, with the entry's own target and the route's reason beside
// it. The target is printed as the record wrote it, because a reader comparing the two columns
// is the person who has to decide what the record should have said.
func table(b *Bundle, d Disposition, last string) string {
	var s strings.Builder
	fmt.Fprintf(&s, "| record | category | target | %s |\n|---|---|---|---|\n", last)
	n := 0
	for _, it := range b.Items {
		if it.Disposition != d {
			continue
		}
		n++
		fmt.Fprintf(&s, "| `%s` entry %d | %s | `%s` | %s |\n",
			it.Path, it.Index, cell(it.Entry.Category), cell(it.Entry.Target), cell(it.Reason))
	}
	if n == 0 {
		s.Reset()
		s.WriteString("None.\n")
	}
	s.WriteString("\n")
	return s.String()
}

// cell escapes what a table cell cannot hold. A target is text somebody typed, and one
// containing a pipe would end the column it is in and shift every cell after it, which is a
// record misreported rather than a record formatted badly.
func cell(s string) string {
	return strings.ReplaceAll(strings.Join(strings.Fields(s), " "), "|", `\|`)
}

// wordings is the project-convention list, grouped by the file each one names and carrying the
// proposal in full. A table would not do here: the person reading this list has to word a
// sentence from the proposal, so the proposal is the column that matters and it is a paragraph,
// which is the one thing a table cell cannot hold.
func wordings(b *Bundle) string {
	var order []string
	byFile := map[string][]Item{}
	for _, it := range b.Items {
		if it.Disposition != ForWording {
			continue
		}
		if _, seen := byFile[it.Reason]; !seen {
			order = append(order, it.Reason)
		}
		byFile[it.Reason] = append(byFile[it.Reason], it)
	}
	if len(order) == 0 {
		return "None.\n\n"
	}
	var s strings.Builder
	for _, f := range order {
		fmt.Fprintf(&s, "### %s\n\n", f)
		for _, it := range byFile[f] {
			fmt.Fprintf(&s, "`%s` entry %d:\n\n", it.Path, it.Index)
			for _, l := range wrap(it.Entry.Proposal, 86) {
				s.WriteString("> " + l + "\n")
			}
			s.WriteString("\n")
		}
	}
	return s.String()
}

// Targets is how often each distinct target was named, under each disposition, which is the
// figure #289 asked for and the figure a reader of a long unresolved list wants next: the same
// prose target named forty times is one decision, not forty.
func (b *Bundle) Targets(d Disposition) []string {
	n := map[string]int{}
	for _, it := range b.Items {
		if it.Disposition == d {
			n[it.Entry.Target]++
		}
	}
	var targets []string
	for t := range n {
		targets = append(targets, t)
	}
	// Most named first, and alphabetically where two were named as often. The order is the
	// one a reader deciding what to fix reads in, and ties break on the target so that two
	// runs over one trail print the same list.
	sort.Slice(targets, func(i, j int) bool {
		if n[targets[i]] != n[targets[j]] {
			return n[targets[i]] > n[targets[j]]
		}
		return targets[i] < targets[j]
	})
	var out []string
	for _, t := range targets {
		out = append(out, fmt.Sprintf("%4d  %s", n[t], t))
	}
	return out
}
