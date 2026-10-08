// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"bytes"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"go.yaml.in/yaml/v3"

	"github.com/triplem/xeno/internal/fm"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
)

// The context scope section 5 defines, and the command that writes it. P0 produces it as an
// artifact of its own and a phase reads what it names, so the scope is the only input to the
// lock's files list and through it to the budget, the link check, the staleness half of
// G-Freshness and Changed.
//
// It was specified and had no writer for a hundred intents (#217), which is the family #188,
// #195, #208 and #220 closed: an artifact nobody can produce with a command does not get
// produced. Unlike those four this one failed silently, because the readers treat an absent
// scope as a project that declared nothing rather than as a phase that owes one.

// ScopeResolved is what a set of patterns selects now, which the writer reports so that a
// budget is a counted figure. It is not recorded anywhere: the lock records the resolution a
// phase was given, and a figure from the moment the scope was written would describe neither.
type ScopeResolved struct {
	Files int
	Bytes int64
}

// ScopeSet writes the intent's context scope and reports what its patterns resolve to.
//
// The entry arrives as the YAML section 5 defines, for RecordQuestion's reason: include
// patterns, exclude patterns, a link per component and two budget numbers are a nested
// structure however they are spelled, and the shape read is then the shape written. Unknown
// fields are refused rather than dropped, which is the standing rule about invented fields.
//
// There is no phase argument. The scope is P0's artifact and is read from P0 by every phase
// of the intent, so a flag would offer a choice the format does not have.
func (r *Runner) ScopeSet(key string, entry []byte) (*ScopeResolved, error) {
	dec := yaml.NewDecoder(bytes.NewReader(entry))
	dec.KnownFields(true)
	var s model.Scope
	if err := dec.Decode(&s); err != nil {
		return nil, refuse("the scope could not be read as a %s: %v", model.ContextScope, err)
	}
	// The header is provenance and the runner's to write, as it is in every other artifact.
	// A35's rule is why it is refused rather than overwritten: a plausible value in a field
	// nobody produced is worse than an absent one, and silently replacing one a person typed
	// would leave them believing they had set it.
	if s.Common != (model.Common{}) {
		return nil, refuse("a scope's header is written by the runner and not given: " +
			"leave intent, phase, created, schema_version, runner_version and plugin_version out")
	}
	if len(s.Include) == 0 {
		return nil, refuse("a scope needs include: a scope that resolves to no files reaches " +
			"every reader as the absent one it is meant to replace")
	}
	dir := model.PhaseDir(key, model.Phases[0])
	if !fm.Exists(r.abs(dir)) {
		return nil, refuse("%s has not started; run xeno phase start --intent %s --phase 00 first",
			model.Phases[0], key)
	}
	common, err := r.common(key, model.Phases[0])
	if err != nil {
		return nil, err
	}
	s.Common = common
	if err := fm.WriteYAML(r.contentPath(key, model.Phases[0], model.ContextScope), s); err != nil {
		return nil, err
	}
	files, err := r.resolveScope(&s)
	if err != nil {
		return nil, err
	}
	out := &ScopeResolved{Files: len(files)}
	for _, f := range files {
		out.Bytes += f.Bytes
	}
	return out, nil
}

// readScope reads the intent's scope, or nil where none is written.
//
// An absent scope is not an error here. It is a phase that owes one, which phase finish
// refuses at P0, and the three gate readers each say what its absence means for them.
func (r *Runner) readScope(key string) (*model.Scope, error) {
	var s model.Scope
	path := r.contentSource(key, model.Phases[0], model.ContextScope)
	if err := fm.ReadYAML(path, &s); err != nil {
		return nil, nil
	}
	return &s, nil
}

// resolveScope is the information base a scope names: one walk, a path, a hash and a size per
// file. Shared by phase start, which records it in the lock, and by ScopeSet, which reports
// the figure a budget is then set from. Two resolutions would be two answers to one question,
// and the budget is set from the first and judged against the second.
func (r *Runner) resolveScope(s *model.Scope) ([]model.ContextFile, error) {
	// One walk, and the files are bucketed by the include pattern that claimed them first. The
	// order the buckets are emitted in is the scope's own, because section 5 asks for an order
	// of volatility and the project is what knows which of its directories is stable. A file two
	// patterns match takes the position of the first: a stable prefix is decided by the first
	// thing that claims it.
	buckets := make([][]model.ContextFile, len(s.Include))
	seen := map[string]bool{}
	err := filepath.WalkDir(r.Root, func(abs string, d os.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		rel, rerr := filepath.Rel(r.Root, abs)
		if rerr != nil {
			return rerr
		}
		rel = filepath.ToSlash(rel)
		if strings.HasPrefix(rel, ".git/") || seen[rel] {
			return nil
		}
		i := firstMatch(s.Include, rel)
		if i < 0 || matchesAny(s.Exclude, rel) {
			return nil
		}
		h, herr := hashing.FileHash(abs)
		if herr != nil {
			return herr
		}
		// The size comes from the entry the walk already has, rather than from a second stat:
		// two reads of one file can see two versions of it, and the recorded size exists to
		// describe the same read as the hash.
		info, ierr := d.Info()
		if ierr != nil {
			return ierr
		}
		seen[rel] = true
		buckets[i] = append(buckets[i], model.ContextFile{Path: rel, SHA256: h, Bytes: info.Size()})
		return nil
	})
	if err != nil {
		return nil, err
	}
	var files []model.ContextFile
	for _, b := range buckets {
		sort.Slice(b, func(i, j int) bool { return b[i].Path < b[j].Path })
		files = append(files, b...)
	}
	// A declared link's document is part of what the phase was given, whether or not include
	// matches it. Section 5 has links declared and never inferred, and a declaration that put
	// nothing in the base would be ornamental. It goes last: a link is the most specific thing
	// in a scope and therefore the most likely to move.
	for _, l := range s.Links {
		if l.Docs == "" || seen[l.Docs] {
			continue
		}
		h, herr := hashing.FileHash(r.abs(l.Docs))
		if herr != nil {
			continue // a link naming a document that is not there is G-Schema's finding
		}
		info, ierr := os.Stat(r.abs(l.Docs))
		if ierr != nil {
			continue // gone between the hash and the stat, which the same finding covers
		}
		seen[l.Docs] = true
		files = append(files, model.ContextFile{Path: l.Docs, SHA256: h, Bytes: info.Size()})
	}
	return files, nil
}
