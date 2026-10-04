// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/triplem/xeno/internal/gates"
	"github.com/triplem/xeno/internal/hashing"
	"github.com/triplem/xeno/internal/model"
)

// The declaration of section 4, which the section calls the act that binds and the only
// one Xeno knows. Running a test is not done here and judging is a gate's; what this file
// does is write down that a result exists, with its kind, its location and its hash.
//
// It is the third writer of a frontmatter block and borrows exchange.go's two helpers for
// it: one field is amended and the body is left byte for byte, because re-rendering
// through SectionSet would make a declaration depend on a template resolution that has
// nothing to do with it. Nothing here reads gate.yaml either, for the reason that file
// records: a write after a verdict changes the artifact and the next phase finish says so.
//
// What is new here is the hash. The item is bound by its sha256 and by nothing else, so on
// the form where the bytes are present the command computes it and no flag supplies it.
// That is the property #208 asks for in the words of the clause it cites — a command that
// changes a hashed file recomputes what it invalidates — and until now SectionSet did it
// for context_hash and nothing did it anywhere else.

// DeclareEvidence writes one item into the evidence block of a phase's output.md.
//
// The three forms are section 4's three states and differ only in which fields are filled.
// An item that exists arrives as file, is copied into the phase's evidence/ and declares
// the path and the hash of what was copied. An item in an artifact store arrives as uri
// with the hash the pipeline published. An item a pipeline has yet to produce arrives as
// neither and declares its kind and its job, which is the pair evidence.Attach looks it up
// by later.
//
// The copy and the declaration are one act. Two commands, or a flag naming a path somebody
// had already copied, would make G-Schema's finding about an undeclared file in evidence/
// reachable by using the tool correctly.
func (r *Runner) DeclareEvidence(key, phase, file string, e model.EvidenceItem) (*model.EvidenceItem, error) {
	if err := declarable(file, e); err != nil {
		return nil, err
	}
	front, o, body, err := r.artifact(key, phase)
	if err != nil {
		return nil, err
	}
	for _, d := range o.Evidence {
		if d.Kind == e.Kind && d.Job == e.Job {
			return nil, refuse("%s already declares %s: an item is identified by its kind and "+
				"its job, which is the pair an attachment is bound to, so one of two "+
				"declarations of it would stay pending with nothing saying why", phase, pair(e))
		}
	}
	// Read and hash before anything is written, so that a refusal below leaves no copy in
	// evidence/ and no half declared item.
	var content []byte
	if file != "" {
		content, err = os.ReadFile(file)
		if err != nil {
			return nil, refuse("the report named by --file cannot be read: %v", err)
		}
		e.Path = "evidence/" + filepath.Base(file)
		e.SHA256 = hashing.Hex(hashing.Normalise(content))
	}
	if err := shapeRefusal(gates.EvidenceShape(
		model.PhaseDir(key, phase)+"/output.md", model.Output{Evidence: []model.EvidenceItem{e}})); err != nil {
		return nil, err
	}
	if file != "" {
		if err := r.copyEvidence(key, phase, e, content); err != nil {
			return nil, err
		}
	}
	if err := r.amendFront(key, phase, front, "evidence", append(o.Evidence, e), body); err != nil {
		return nil, err
	}
	return &e, nil
}

// declarable judges the flags against each other, which is the part no gate can judge: the
// gate reads a declaration and these are the ways of arriving at one that cannot be
// written. Everything about the item itself is left to EvidenceShape.
func declarable(file string, e model.EvidenceItem) error {
	switch {
	case e.Kind == "":
		return refuse("--kind is required, from section 4's set: %s",
			strings.Join(model.EvidenceKinds, ", "))
	case file != "" && e.URI != "":
		return refuse("--file and --uri are the two places an item can live, and section 4 " +
			"decides between them by size: a small textual result is copied into evidence/ " +
			"and a large or binary one is referenced where it lies. Declare one of them")
	case file != "" && e.SHA256 != "":
		return refuse("--sha256 is not given with --file: the bytes are here, so the hash is " +
			"computed from them. A hash supplied beside the file can only agree, which is " +
			"noise, or disagree, which is a declaration saying what somebody chose rather " +
			"than what the file is")
	case e.URI != "" && e.SHA256 == "":
		return refuse("--uri needs --sha256: a uri is bound by its hash alone, because " +
			"resolving one needs a network the gate path never has")
	case e.SHA256 != "" && e.URI == "":
		return refuse("--sha256 names a hash and nothing it belongs to: give it with --uri, " +
			"or declare --file, which computes it")
	}
	if file != "" || e.URI != "" {
		return nil
	}
	switch {
	case e.Job == "":
		return refuse("--job is required on an item a pipeline has yet to produce: it names " +
			"the job expected to produce it, and the kind and the job are what the " +
			"attachment is bound to when it arrives")
	case e.Result != "":
		return refuse("a pending item carries no --result: the run that would report one has " +
			"not happened, and the value arrives in evidence/attached.yaml as the job's own " +
			"verdict")
	}
	return nil
}

// copyEvidence puts the report beside the artifact that declares it.
//
// It refuses to write over a file whose content differs. A file in evidence/ is bound by a
// hash in some declaration, and nothing here can tell whether that declaration is this
// item's or another phase's reading of the same name, so the overwrite would break a
// verdict this command was not asked about. Identical content is not a conflict: the
// declaration being written binds exactly the bytes that are already there.
func (r *Runner) copyEvidence(key, phase string, e model.EvidenceItem, content []byte) error {
	rel := model.PhaseDir(key, phase) + "/" + e.Path
	if got, err := hashing.FileHash(r.abs(rel)); err == nil && got != e.SHA256 {
		return refuse("%s already holds other content, and a file there is bound by a hash in "+
			"some declaration: give the report a name of its own", rel)
	}
	if err := os.MkdirAll(filepath.Dir(r.abs(rel)), 0o755); err != nil {
		return err
	}
	return os.WriteFile(r.abs(rel), content, 0o644)
}

// pair is how an item is named in a sentence: the two fields that identify it.
func pair(e model.EvidenceItem) string {
	if e.Job == "" {
		return e.Kind
	}
	return e.Kind + "/" + e.Job
}
