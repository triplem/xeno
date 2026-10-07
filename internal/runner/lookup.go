// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"fmt"
	"time"

	"github.com/triplem/xeno/internal/index"
	"github.com/triplem/xeno/internal/model"
	"github.com/triplem/xeno/internal/template"
)

// Template resolves the template and the strings bundle a phase renders from.
//
// It exists so that fetching a template is the same act as writing a section: SectionSet
// resolves the template the same way and through the same two lines, so a caller that
// asked what the sections are cannot be told one set and have another rendered. Section 13
// lists fetch template among the six operations the MCP server exposes, and that server
// calls this rather than reading the plugin directory for itself.
func (r *Runner) Template(phase string) (*template.Resolved, error) {
	if model.PhaseIndex(phase) < 0 {
		return nil, fmt.Errorf("unknown phase %q", phase)
	}
	t, err := template.Load(r.Root, model.TemplateID(phase), r.language())
	if err != nil {
		return nil, refuse("%v", err)
	}
	return &t, nil
}

// SymbolAnswer is where a name is defined, together with what answered. The provenance
// travels with the answer rather than being asked for separately, because an index is
// allowed to be missing, stale or wrong and a location with nothing saying what produced
// it cannot be weighed against that.
type SymbolAnswer struct {
	Name  string
	Found []index.Symbol
	// Tool, ToolVersion and Age describe the index that answered, and are empty where
	// there was none.
	Tool, ToolVersion string
	Age               time.Duration
	// Why is the sentence index.Load returns when there is no usable index, empty where
	// there is one. A name the index does not hold is not this case: that is an answer
	// with no locations in it.
	Why string
}

// Symbols answers where a name is defined, reading the index through the same path a phase
// does.
//
// It cannot fail, for index.Load's reason: absent, unreadable, malformed and stale are four
// causes with one outcome, and no phase and no verdict depends on the answer. A caller gets
// an answer with no locations and the sentence saying why, which is what it would record.
func (r *Runner) Symbols(name string) *SymbolAnswer {
	now := r.Now()
	i, why := r.symbolIndex(now)
	a := &SymbolAnswer{Name: name, Why: why}
	if i == nil {
		return a
	}
	a.Tool, a.ToolVersion, a.Age = i.Tool, i.ToolVersion, i.Age(now)
	a.Found = i.Lookup(name)
	return a
}
