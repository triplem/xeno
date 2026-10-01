// SPDX-License-Identifier: Apache-2.0

// Package rules reads the rule tree of section 9: the rule file, the two axes it is laid out
// on, the configuration errors that make a file unusable as a rule, the precedence that
// reduces the tree to one rule per id, and the hash over that effective set.
//
// It evaluates no predicate. A rule's check is read as a named type with its parameters and
// handed on, because whether a predicate holds is a question about an artifact or a commit
// range, and whether a rule set resolves is a question about the tree. The second has to be
// answered first: a set that does not resolve has no business running anything.
//
// Nothing here words a verdict. Load and Effective return what they could not use, each with
// a path, a cause and a next step, and the caller decides what that is worth. Three callers
// are coming — G-Rules, G-Policy and the writer of rules_hash — and a package that printed or
// refused would decide for all three. A61 took the same shape for evidence attachment.
package rules

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"go.yaml.in/yaml/v3"

	"github.com/triplem/xeno/internal/hashing"
)

// Where section 9 puts the two trees, relative to a repository root. They are constants and
// not configuration: a path a project could move would make scope against the path
// uncheckable, which is the one thing the gate is for.
const (
	PluginDir = ".xeno/plugin/rules"
	ConfigDir = ".xeno/config/rules"
)

// The two axes. The path determines both, and the frontmatter repeats the level as scope so
// that a moved file cannot silently change its reach.
const (
	Given   = "given"
	Learned = "learned"
)

// Levels from least to most specific, which is also the precedence order of section 9.
var Levels = []string{"builtin", "provider", "org", "project"}

func rank(level string) int {
	for i, l := range Levels {
		if l == level {
			return i
		}
	}
	return -1
}

// Check is a rule's predicate: a named type and the parameters it is given. The parameters are
// kept as read rather than typed, because the types are implemented by the piece that
// evaluates them and a shape fixed here would be fixed before anything used it.
type Check struct {
	Type   string
	Params map[string]any
}

// UnmarshalYAML splits the type from the rest. Section 9 writes a check as a mapping whose
// type names the predicate and whose remaining keys parameterise it, so the two cannot be
// separate fields in the file.
func (c *Check) UnmarshalYAML(n *yaml.Node) error {
	var raw map[string]any
	if err := n.Decode(&raw); err != nil {
		return err
	}
	c.Params = map[string]any{}
	for k, v := range raw {
		if k == "type" {
			c.Type, _ = v.(string)
			continue
		}
		c.Params[k] = v
	}
	return nil
}

// Rule is one rule file, plus where it was found. Origin, Level and Path are not in the file:
// the path carries them, and Scope is the file's own claim about the level, which is what
// gets compared.
type Rule struct {
	ID        string   `yaml:"id"`
	Version   int      `yaml:"version"`
	Scope     string   `yaml:"scope"`
	Binding   bool     `yaml:"binding"`
	Kind      string   `yaml:"kind"`
	AppliesTo []string `yaml:"applies_to"`
	Statement string   `yaml:"statement"`
	Abstract  string   `yaml:"abstract"`
	Check     *Check   `yaml:"check"`

	Origin string `yaml:"-"`
	Level  string `yaml:"-"`
	Path   string `yaml:"-"` // relative to the repository root, as a finding has to name it
}

// Problem is a file, or a pair of files, that cannot be used as section 9 requires. It carries
// what a red verdict has to name, so that the gate words it once and invents nothing.
type Problem struct {
	Path  string
	Cause string
	Next  string
}

// Kinds, and what each one does with a check.
const (
	Checked = "checked"
	Review  = "review"
)

// Load walks both trees and returns every rule it could read and every file it could not.
//
// A repository with no rule tree yields no rules and no problems. That is the state of every
// repository before a plugin is vendored and of every project that maintains no rules of its
// own, and it is not a broken rule set.
func Load(root string) ([]Rule, []Problem) {
	var rules []Rule
	var problems []Problem
	for _, origin := range []string{Given, Learned} {
		for _, level := range Levels {
			dir := treeDir(origin, level)
			if dir == "" {
				continue
			}
			abs := filepath.Join(root, dir)
			entries, err := os.ReadDir(abs)
			if err != nil {
				continue // an absent directory is a level nobody maintains
			}
			// learned/builtin does not exist: section 9 says nothing learns into a pinned,
			// hashed package, and the way into the shipped set is a pull request. The
			// finding is on the directory, because its presence is the error.
			if origin == Learned && level == "builtin" {
				problems = append(problems, Problem{dir,
					"learned/builtin/ does not exist; nothing learns into the shipped set",
					"move these rules to learned/org/ or propose them for the shipped set, as section 15 describes"})
				continue
			}
			for _, e := range entries {
				if e.IsDir() || !isRuleFile(e.Name()) {
					continue
				}
				rel := filepath.Join(dir, e.Name())
				r, probs := read(root, rel, origin, level)
				problems = append(problems, probs...)
				if r != nil {
					rules = append(rules, *r)
				}
			}
		}
	}
	problems = append(problems, duplicates(rules)...)
	return rules, problems
}

// treeDir is where one axis combination lives. The shipped set sits with the plugin, so it is
// covered by the plugin hash and cannot be edited in place in a project; everything else is
// the project's own.
func treeDir(origin, level string) string {
	if level == "builtin" {
		if origin == Given {
			return PluginDir + "/given/builtin"
		}
		return ConfigDir + "/learned/builtin" // only to be reported, never read
	}
	return ConfigDir + "/" + origin + "/" + level
}

func isRuleFile(name string) bool {
	return strings.HasSuffix(name, ".yaml") || strings.HasSuffix(name, ".yml")
}

// read parses one file and checks it against what sections 7 and 9 require of it. A file that
// fails any of these is not returned as a rule: the verdict is red either way, and resolving a
// rule whose scope contradicts its path would mean choosing which of the two to believe.
func read(root, rel, origin, level string) (*Rule, []Problem) {
	b, err := os.ReadFile(filepath.Join(root, rel))
	if err != nil {
		return nil, []Problem{{rel, "cannot be read: " + err.Error(), "make the file readable, or take it out of the tree"}}
	}
	var r Rule
	if err := yaml.Unmarshal(b, &r); err != nil {
		return nil, []Problem{{rel, "is not a readable rule file: " + err.Error(),
			"write the fields section 9 lists, as YAML"}}
	}
	r.Origin, r.Level, r.Path = origin, level, rel

	var ps []Problem
	add := func(cause, next string) { ps = append(ps, Problem{rel, cause, next}) }

	// The format's own completeness. A file missing one of these cannot be read as a rule at
	// all, which is why it is reported here rather than treated as a rule with a gap.
	if r.ID == "" {
		add("has no id", "give the rule an id; it is what precedence resolves and what the checklist anchors on")
	}
	if r.Version < 1 {
		add("has no version", "write version as a counter starting at 1; it claims nothing about compatibility")
	}
	if r.Statement == "" {
		add("has no statement", "write what the rule requires, in one sentence a person can answer against")
	}
	if len(r.AppliesTo) == 0 {
		add("names no phase in applies_to", "list the phases the rule applies to")
	}

	switch r.Scope {
	case "":
		add("has no scope", "write scope as the level its path gives it: "+level)
	case level:
	default:
		add(fmt.Sprintf("claims scope %s where its path gives it %s", r.Scope, level),
			"correct the scope, or move the file to the level it claims")
	}

	switch r.Kind {
	case Checked:
		if r.Check == nil {
			add("is kind checked and carries no check",
				"give it a check, or write it as kind review for a person to answer")
		}
	case Review:
		if r.Check != nil {
			add("is kind review and carries a check",
				"take the check out, or write it as kind checked")
		}
	case "":
		add("has no kind", "write kind as checked for a predicate or review for a person to answer")
	default:
		add("has kind "+r.Kind, "write kind as checked or review; section 9 fixes the two")
	}

	if r.Binding && origin != Given {
		add("is binding under learned/",
			"binding is only allowed under given/; take it out or move the rule to given/")
	}
	if r.Binding && origin == Given && level == "project" {
		add("is binding under given/project/",
			"binding has nothing left to bind at the project level; take it out")
	}
	if origin == Learned && (level == "provider" || level == "org") && r.Abstract == "" {
		add("is a learned rule above the project level and carries no abstract",
			"write an abstract: a transferable summary with no project or code specifics")
	}

	if len(ps) > 0 {
		return nil, ps
	}
	return &r, nil
}

// duplicates reports one id claimed twice on one level by one origin. The same id on two
// levels is the ordinary case precedence exists for; twice in one directory is a collision
// nothing can resolve.
func duplicates(rules []Rule) []Problem {
	seen := map[string]Rule{}
	var ps []Problem
	for _, r := range rules {
		k := r.Origin + "/" + r.Level + "/" + r.ID
		if first, ok := seen[k]; ok {
			ps = append(ps, Problem{r.Path,
				fmt.Sprintf("claims id %s, which %s already claims at the same level", r.ID, first.Path),
				"give one of them a different id, or take one out"})
			continue
		}
		seen[k] = r
	}
	return ps
}

// Effective reduces the tree to one rule per id, in the precedence order of section 9: more
// specific beats less specific, given beats the matching learned at the same level, and a
// binding rule under given/builtin/, given/provider/ or given/org/ cannot be overridden at
// all.
//
// Two binding rules colliding on different levels leave the id out of the set. Section 9 says
// the gate is red and stays red, and that this is a matter between two organisations and not
// something a gate may resolve, so the resolver picks neither rather than the more specific
// one.
func Effective(rules []Rule) ([]Rule, []Problem) {
	byID := map[string][]Rule{}
	var ids []string
	for _, r := range rules {
		if _, seen := byID[r.ID]; !seen {
			ids = append(ids, r.ID)
		}
		byID[r.ID] = append(byID[r.ID], r)
	}
	sort.Strings(ids)

	var out []Rule
	var ps []Problem
	for _, id := range ids {
		cands := byID[id]
		var binding []Rule
		for _, r := range cands {
			if r.Binding {
				binding = append(binding, r)
			}
		}
		switch {
		case len(binding) > 1:
			sort.Slice(binding, func(i, j int) bool { return rank(binding[i].Level) < rank(binding[j].Level) })
			var where []string
			for _, r := range binding {
				where = append(where, r.Path)
			}
			ps = append(ps, Problem{binding[0].Path,
				fmt.Sprintf("binding rule %s collides with a binding rule on another level: %s",
					id, strings.Join(where, ", ")),
				"settle it between the levels that declared it; a gate may not resolve a binding collision"})
		case len(binding) == 1:
			out = append(out, binding[0])
		default:
			sort.SliceStable(cands, func(i, j int) bool { return more(cands[i], cands[j]) })
			out = append(out, cands[0])
		}
	}
	return out, ps
}

// more reports whether a beats b: the more specific level first, and given over learned where
// the level is the same.
func more(a, b Rule) bool {
	if ra, rb := rank(a.Level), rank(b.Level); ra != rb {
		return ra > rb
	}
	return a.Origin == Given && b.Origin == Learned
}

// Render is the canonical rendering rules_hash is taken over: one line per effective rule,
// sorted by id, tab separated, newline terminated.
//
// Appendix B fixes every other hash to the byte and delegates this one to the package that
// writes it, as it did for secrets_hash in A62. The fields are the rule's content and not the
// bytes of the file it came from, so a comment, a file name and the order the tree was walked
// in leave the value alone, while an edited statement changes it. What the field has to prove
// is which rule set was in force, and the set in force is the effective one: a rule that lost
// its precedence contest never applied and is not in here.
func Render(effective []Rule) string {
	rs := make([]Rule, len(effective))
	copy(rs, effective)
	sort.Slice(rs, func(i, j int) bool { return rs[i].ID < rs[j].ID })
	var b strings.Builder
	for _, r := range rs {
		phases := make([]string, len(r.AppliesTo))
		copy(phases, r.AppliesTo)
		sort.Strings(phases)
		fmt.Fprintf(&b, "%s\t%d\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			r.ID, r.Version, r.Origin, r.Level, r.Kind, strconv.FormatBool(r.Binding),
			strings.Join(phases, ","), oneLine(r.Statement), renderCheck(r.Check))
	}
	return b.String()
}

// Hash is rules_hash over the rendering above. An empty effective set hashes the empty
// rendering: the field is required by section 5, so leaving it out is not available the way it
// is for secrets_hash in a digest, and what the value then says is that a tree was resolved
// and the set came out empty.
func Hash(effective []Rule) string {
	return hashing.Hex([]byte(Render(effective)))
}

// oneLine normalises a statement's whitespace. Section 9 writes statements as folded YAML
// scalars, so the same sentence can arrive with different line breaks from two editors, and a
// line break is not a change to what a rule requires.
func oneLine(s string) string { return strings.Join(strings.Fields(s), " ") }

// renderCheck flattens a check to its type and its parameters as sorted dotted keys. A nested
// map is flattened rather than marshalled, so that the ordering is this function's and not a
// YAML library's.
func renderCheck(c *Check) string {
	if c == nil {
		return ""
	}
	parts := []string{"type=" + c.Type}
	parts = append(parts, flatten("", c.Params)...)
	sort.Strings(parts[1:])
	return strings.Join(parts, " ")
}

func flatten(prefix string, v any) []string {
	switch t := v.(type) {
	case map[string]any:
		var out []string
		for k, val := range t {
			out = append(out, flatten(join(prefix, k), val)...)
		}
		return out
	case []any:
		var out []string
		for i, val := range t {
			out = append(out, flatten(join(prefix, strconv.Itoa(i)), val)...)
		}
		return out
	default:
		return []string{prefix + "=" + fmt.Sprint(v)}
	}
}

func join(prefix, k string) string {
	if prefix == "" {
		return k
	}
	return prefix + "." + k
}
