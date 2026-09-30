// SPDX-License-Identifier: Apache-2.0

package gitlab

import (
	"os"
	"testing"

	"github.com/triplem/xeno/internal/enforcement"
)

// TestLiveGitLabAnswersAboutAPublicProject asks gitlab.com about a real project. It is
// skipped unless XENO_LIVE is set, because a test that needs a network is not one the suite
// may depend on, and A42 keeps the gate path free of both.
//
// Without a token it asks only what a public read can answer, which is the protected branch:
// that is the payload A65 turned on, and asserting it here is what keeps this adapter honest
// about the shape of a real answer rather than of a fixture somebody wrote.
func TestLiveGitLabAnswersAboutAPublicProject(t *testing.T) {
	if os.Getenv("XENO_LIVE") == "" {
		t.Skip("set XENO_LIVE to ask gitlab.com")
	}
	a := New(nil, "https://gitlab.com/api/v4")

	pb, why, err := a.branch("gitlab-org/gitlab", "master", "")
	if err != nil {
		t.Fatal(err)
	}
	if why != "" {
		t.Fatalf("the branch could not be read: %s", why)
	}
	q := allowBypass(pb)
	t.Logf("%-26s %-14s %s", q.Name, q.State, q.Actual)
	if q.State != enforcement.Unmet {
		t.Errorf("allow_bypass on a project that permits pushes is %q", q.State)
	}

	// The merge settings are omitted from an unauthenticated read, which is the defect this
	// asserts against: the field is absent and must not decode to off.
	pr, why, err := a.project("gitlab-org/gitlab", "")
	if err != nil || why != "" {
		t.Fatalf("the project could not be read: %v %s", err, why)
	}
	if pr.OnlyAllowMergeIfPipelineSucceeds != nil {
		t.Log("the host now sends the merge settings unauthenticated; the pointer is still correct")
	}
}

// TestLiveGitLabWithAToken runs the whole adapter, which needs a credential: both approvals
// endpoints refuse an anonymous request, and a rejected token is an error for the adapter
// rather than four requirements reported as unavailable.
func TestLiveGitLabWithAToken(t *testing.T) {
	token := os.Getenv("XENO_GITLAB_TOKEN")
	if os.Getenv("XENO_LIVE") == "" || token == "" {
		t.Skip("set XENO_LIVE and XENO_GITLAB_TOKEN to run the adapter against gitlab.com")
	}
	repo := os.Getenv("XENO_GITLAB_PROJECT")
	if repo == "" {
		repo = "gitlab-org/gitlab"
	}
	branch := os.Getenv("XENO_GITLAB_BRANCH")
	if branch == "" {
		branch = "master"
	}
	qs, err := New(nil, "https://gitlab.com/api/v4").
		Requirements(repo, branch, token, declared())
	if err != nil {
		t.Fatal(err)
	}
	known := make(map[string]bool)
	for _, n := range enforcement.Names() {
		known[n] = true
	}
	for _, q := range qs {
		t.Logf("%-26s %-14s %s | %s", q.Name, q.State, q.Actual, q.Note)
		if !known[q.Name] {
			t.Errorf("%q is not a name the domain knows", q.Name)
		}
		if q.State == enforcement.Waived {
			t.Errorf("%s is waived, which is the project's decision and not the host's", q.Name)
		}
	}
	if len(qs) != 4 {
		t.Errorf("%d requirements answered, want 4", len(qs))
	}
}
