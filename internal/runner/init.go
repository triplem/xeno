// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/plugin"
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
	// PluginFrom is where the vendored plugin came from, printed because the answer decides
	// whether G-Supply can pass: a release carries the bytes its digest was taken over, and
	// a directory is whatever happens to be in it.
	PluginFrom string
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
		"create the label " + model.ApprovedLabel + ", with .xeno/plugin/bin/xeno-labels.sh or by hand; " +
			"an issue becomes an intent only with it and a comment beginning " + model.ApprovedWord,
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
	return r.createMode(res, rel, content, 0o644)
}

// createMode is create with the file mode named, for the one thing init writes that has to be
// executable: the plugin's entry point, whose mode cannot travel through an embed.FS.
func (r *Runner) createMode(res *InitResult, rel, content string, mode fs.FileMode) error {
	if fm.Exists(r.abs(rel)) {
		res.Kept = append(res.Kept, rel)
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(r.abs(rel)), 0o755); err != nil {
		return err
	}
	if err := os.WriteFile(r.abs(rel), []byte(content), mode); err != nil {
		return err
	}
	res.Created = append(res.Created, rel)
	return nil
}

// appendGitignore is the one place init writes into a file it did not create, and it
// only appends lines that are not there. A .gitignore belongs to the project.
func (r *Runner) appendGitignore(res *InitResult) error {
	const rel = ".gitignore"
	// The default location and not model.LocalDir's answer, which is the one site #205 left
	// unresolved on purpose: a .gitignore describes this repository, not where a person
	// redirected their own state with XENO_PLUGIN_DATA. A directory outside the tree needs no
	// entry, and one inside it is covered by this.
	want := []string{model.LocalDefault + "/"}
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
// One walk over whatever the plugin is, rather than a list of the parts of it. It was three
// functions with a named tree each, a list of single files and a depth assumption apiece, and
// it had been wrong twice over: `secrets.yaml` was in the comment naming section 13's list and
// in none of the calls, and `bin/`, added with the entry point, was in neither. An adopter
// therefore received a tree whose digest could not be the one a released runner expects, so
// G-Supply would have failed for every adopter on every phase — which is the gate this walk
// had to exist for before that gate could ship (#199).
//
// A walk cannot go stale the next time the plugin gains a directory, which a list does, twice.
func (r *Runner) vendorPlugin(res *InitResult) error {
	src, from, err := r.pluginSource()
	if err != nil {
		return err
	}
	res.PluginFrom = from
	return fs.WalkDir(src, ".", func(rel string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		b, rerr := fs.ReadFile(src, rel)
		if rerr != nil {
			return rerr
		}
		return r.createMode(res, plugin.Dir+"/"+rel, string(b), vendoredMode(rel))
	})
}

// pluginSource is the plugin this binary will vendor, and where it came from.
//
// The embedded copy wins where there is one. A release carries the bytes its digest was taken
// over, and vendoring from anywhere else would produce a tree G-Supply fails — so a released
// runner must not be talked into copying a directory that happens to be lying about.
//
// A build that carries none falls back to the directory, which is what `--plugin-from` names
// and what this repository develops against. Where that is not a plugin either, the refusal
// says which build this is rather than which file was missing.
func (r *Runner) pluginSource() (fs.FS, string, error) {
	if shipped, ok := plugin.Shipped(); ok {
		return shipped, "this release", nil
	}
	src := os.DirFS(r.PluginSource)
	if _, err := fs.Stat(src, manifestRel); err != nil {
		return nil, "", refuse("this build carries no plugin, and %s is not one either: %v.\n"+
			"A release carries the plugin it was built with; a development build needs "+
			"--plugin-from pointing at one.", r.PluginSource, err)
	}
	return src, r.PluginSource, nil
}

// manifestRel is the file that makes a directory a plugin rather than a directory. A client
// loads a plugin by reading it, which is what A77 established against the installed client.
const manifestRel = ".claude-plugin/plugin.json"

// vendoredMode is 0755 under bin/ and 0644 elsewhere. The executable bit cannot travel: an
// embed.FS reports every file as read-only regardless of what was committed, so a script the
// plugin's hook invokes has to be made executable on the way out rather than copied as found.
func vendoredMode(rel string) fs.FileMode {
	if strings.HasPrefix(rel, "bin/") {
		return 0o755
	}
	return 0o644
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
