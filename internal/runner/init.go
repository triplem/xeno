// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/scaffold"
)

// InitOptions are the three things xeno init asks and the flags around them.
//
// Three, and the rest defaulted. Appendix A is long enough that a questionnaire loses
// the first interested party around the fourth item, and adoption is decided here more
// than anywhere else in the plan.
type InitOptions struct {
	TrackerKey string // the project or repository the intents belong to
	Model      string // the default model a phase uses
	Language   string // the language artifacts are written in
	Vendor     bool   // copy the plugin in and pin it
	Host       string // which wrapper to generate; required, since a default is an assumed host
}

// InitResult is what init did and what it could not do, so that the caller prints both
// rather than deciding which matters.
type InitResult struct {
	Created  []string // paths written
	Kept     []string // paths left alone because they already existed
	Manual   []string // settings a person has to make on the host
	Outstand []string // what could not be determined
}

// Init prepares a repository for Xeno. It is every user's first contact with the tool.
//
// It creates what is missing and touches nothing else. A second run therefore changes
// nothing, which is not a convenience: init is the command somebody runs again after
// reading the output of the first one, and a version of it that overwrote would eat
// whatever they changed in between.
func (r *Runner) Init(o InitOptions) (*InitResult, error) {
	res := &InitResult{}
	if o.Language == "" {
		o.Language = "en"
	}

	// The version check comes first. Proceeding with whatever is on the machine and
	// then reporting a mismatch would have written files from the wrong runner by the
	// time anybody read it.
	cfgPath := ".xeno/config/project.yaml"
	if fm.Exists(r.abs(cfgPath)) {
		var p model.Project
		if err := fm.ReadYAML(r.abs(cfgPath), &p); err != nil {
			return nil, refuse("%s cannot be read: %v", cfgPath, err)
		}
		if p.RunnerVersion != "" && p.RunnerVersion != model.RunnerVersion {
			return nil, refuse("this repository is pinned to runner %s and this is %s.\n"+
				"Install the pinned version; a different one produces a different hash or a different verdict than CI.",
				p.RunnerVersion, model.RunnerVersion)
		}
	}

	// The host is resolved once, before either generated file is written: the wrapper and
	// the tracker block are the same choice, and resolving it twice is how they came to be
	// able to disagree.
	h, wrapper, err := Wrapper(r.Root, o.Host, model.RunnerVersion)
	if err != nil {
		return nil, refuse("%v", err)
	}

	cfg, err := r.projectYAML(o, h)
	if err != nil {
		return nil, err
	}
	if err := r.create(res, cfgPath, cfg); err != nil {
		return nil, err
	}
	if err := r.appendGitignore(res); err != nil {
		return nil, err
	}
	if o.Vendor {
		if err := r.vendorPlugin(res); err != nil {
			return nil, err
		}
	}
	if err := r.create(res, h.Path, wrapper); err != nil {
		return nil, err
	}

	// Named and not created, both of them deliberately. They are somebody's act on the
	// host and a command that claimed to have made them would be lying.
	res.Manual = []string{
		"make the pipeline required for a merge on the protected default branch",
		"require a review from somebody other than the author",
		"issue a token that can read the protected branch settings, for xeno enforcement check",
		"schedule xeno enforcement check daily, so that a setting changed back is noticed",
		h.Squash,
	}
	res.Outstand = []string{
		"what the host actually offers: the enforcement block is written from defaults, " +
			"and xeno enforcement check compares it once there is a token",
	}
	// A development build is a poor thing to pin to. The stamp names a commit and
	// whether that tree was clean, so no other machine can install it and this one
	// stops matching after the next commit. Said here rather than refused: refusing
	// would make the tool unusable to the people building it.
	if strings.Contains(model.RunnerVersion, "-dev") {
		res.Outstand = append(res.Outstand,
			"runner_version is pinned to a development build, "+model.RunnerVersion+
				"; a project pins a release, since nobody else can install this one")
	}
	return res, nil
}

// create writes a file that is not there and records either way. It never edits one it
// did not create, which is the property that makes a second run safe.
func (r *Runner) create(res *InitResult, rel, content string) error {
	if fm.Exists(r.abs(rel)) {
		res.Kept = append(res.Kept, rel)
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(r.abs(rel)), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(r.abs(rel), []byte(content), 0o644); err != nil {
		return err
	}
	res.Created = append(res.Created, rel)
	return nil
}

// appendGitignore is the one place init writes into a file it did not create, and it
// only appends lines that are not there. A .gitignore belongs to the project.
func (r *Runner) appendGitignore(res *InitResult) error {
	const rel = ".gitignore"
	want := []string{".xeno/local/"}
	existing := ""
	if b, err := os.ReadFile(r.abs(rel)); err == nil {
		existing = string(b)
	}
	var add []string
	for _, w := range want {
		found := false
		for _, line := range strings.Split(existing, "\n") {
			if strings.TrimSpace(line) == w {
				found = true
				break
			}
		}
		if !found {
			add = append(add, w)
		}
	}
	if len(add) == 0 {
		res.Kept = append(res.Kept, rel)
		return nil
	}
	out := existing
	if out != "" && !strings.HasSuffix(out, "\n") {
		out += "\n"
	}
	out += strings.Join(add, "\n") + "\n"
	if err := os.WriteFile(r.abs(rel), []byte(out), 0o644); err != nil {
		return err
	}
	res.Created = append(res.Created, rel+" ("+strings.Join(add, ", ")+")")
	return nil
}

// vendorPlugin copies the shipped set in so that the repository carries what it renders
// from and what it is judged by. Pinned means it is in the repository and moves only when
// somebody commits it.
//
// Two trees, not one. The templates were vendored from the start; the rules were not, because
// until WP4 there were none, and a shipped set that does not arrive is not shipped. They are
// copied the same way and through the same create, so a second run still changes nothing.
func (r *Runner) vendorPlugin(res *InitResult) error {
	if err := r.vendorTree(res, "templates"); err != nil {
		return err
	}
	// The rule tree is one level deeper and absent in a plugin that carries no rules, which is
	// why it is not an error here: section 9's layout is given/builtin, so the walk is over
	// whatever directories the plugin has under rules.
	return r.vendorRules(res)
}

// vendorTree copies one directory of the plugin, one level of subdirectories deep.
func (r *Runner) vendorTree(res *InitResult, tree string) error {
	src := filepath.Join(r.PluginSource, tree)
	entries, err := os.ReadDir(src)
	if err != nil {
		return refuse("no plugin to vendor at %s: %v", src, err)
	}
	for _, id := range entries {
		files, err := os.ReadDir(filepath.Join(src, id.Name()))
		if err != nil {
			return err
		}
		for _, f := range files {
			b, err := os.ReadFile(filepath.Join(src, id.Name(), f.Name()))
			if err != nil {
				return err
			}
			rel := filepath.ToSlash(filepath.Join(".xeno/plugin", tree, id.Name(), f.Name()))
			if err := r.create(res, rel, string(b)); err != nil {
				return err
			}
		}
	}
	return nil
}

// vendorRules copies the shipped rule set, which lives at rules/given/builtin so that it is
// covered by the plugin hash and is not a level a project maintains.
func (r *Runner) vendorRules(res *InitResult) error {
	const rel = "rules/given/builtin"
	src := filepath.Join(r.PluginSource, filepath.FromSlash(rel))
	files, err := os.ReadDir(src)
	if err != nil {
		if os.IsNotExist(err) {
			return nil // a plugin with no shipped rules, which is every release before WP4
		}
		return err
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		b, err := os.ReadFile(filepath.Join(src, f.Name()))
		if err != nil {
			return err
		}
		if err := r.create(res, ".xeno/plugin/"+rel+"/"+f.Name(), string(b)); err != nil {
			return err
		}
	}
	return nil
}

// projectYAML renders the initial configuration from the scaffold, so that a repository
// or an organisation can replace it without forking the runner. The three answers xeno
// init asks for are named fields rather than positional arguments: a fourth in the wrong
// order used to produce a file that looked right.
//
// The host comes in as the row the wrapper was generated from, so the tracker block names
// the host whose pipeline the repository got.
func (r *Runner) projectYAML(o InitOptions, h WrapperHost) (string, error) {
	key := o.TrackerKey
	if key == "" {
		key = "<owner/repository>"
	}
	mdl := o.Model
	if mdl == "" {
		mdl = "<model identifier>"
	}
	body, _, err := scaffold.RenderProject(r.Root, scaffold.Project{
		RunnerVersion: model.RunnerVersion, Language: o.Language, Model: mdl, TrackerProject: key,
		Adapter: h.Adapter, APIBase: h.APIBase,
	})
	return body, err
}
