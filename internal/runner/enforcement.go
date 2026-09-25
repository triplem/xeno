// SPDX-License-Identifier: Apache-2.0

package runner

import (
	"net/http"
	"time"

	"github.com/triplem/xeno/internal/enforcement"
	"github.com/triplem/xeno/internal/fm"
)

// ReportPath is where the enforcement report is written. Under .xeno/local/, which is
// gitignored, because the report is a pipeline artifact and a job status and never a
// commit: CI verifies artifacts that exist and produces none.
const ReportPath = ".xeno/local/enforcement.yaml"

// EnforcementCheck compares what the project declares it requires of its host against
// what the host is configured to do.
//
// It is a command of its own and never part of xeno gate. It needs the network, and the
// gate path's freedom from it is what makes a verdict reproducible; a gate that called
// out would depend on whether a service answered.
//
// It cannot make itself binding. Where the pipeline is not required for a merge, this
// check failing stops nothing either, which is a circularity with no way out from
// inside the repository. What is left is visibility: the job status, the report, and a
// scheduled run whose failure shows in the pipeline overview.
func (r *Runner) EnforcementCheck(branch, token string) (*enforcement.Report, error) {
	var p struct {
		Tracker struct {
			Project string `yaml:"project"`
			BaseURL string `yaml:"base_url"`
		} `yaml:"tracker"`
		Enforcement enforcement.Declared `yaml:"enforcement"`
	}
	const cfg = ".xeno/config/project.yaml"
	if err := fm.ReadYAML(r.abs(cfg), &p); err != nil {
		return nil, refuse("%s cannot be read; run xeno init first: %v", cfg, err)
	}
	if p.Tracker.Project == "" {
		return nil, refuse("%s names no tracker project, so there is nothing to ask the host about", cfg)
	}
	if token == "" {
		return nil, refuse("no enforcement token; set XENO_ENFORCEMENT_TOKEN.\n" +
			"It is a token that can read the protected branch settings, and it is never in the repository.")
	}
	base := p.Tracker.BaseURL
	if base == "" {
		base = "https://api.github.com"
	}
	if branch == "" {
		branch = "main"
	}

	prot, err := enforcement.Fetch(&http.Client{Timeout: 20 * time.Second},
		base, p.Tracker.Project, branch, token)
	if err != nil {
		return nil, err
	}
	rep := enforcement.Compare(p.Tracker.Project, branch, p.Enforcement, prot, r.Now())
	return &rep, fm.WriteYAML(r.abs(ReportPath), rep)
}
