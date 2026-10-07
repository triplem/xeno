// SPDX-License-Identifier: Apache-2.0

package host

import (
	"strings"
	"testing"

	"github.com/triplem/xeno/internal/model"
)

// #104's first criterion: project.yaml selects the host and no default host name appears in
// code. The absence of a default is the property, so it is asserted rather than left to be
// noticed if somebody puts one back.
func TestThereIsNoDefaultHost(t *testing.T) {
	if _, err := BranchRulesFor("", "https://example.invalid", nil); err == nil {
		t.Fatal("an empty adapter name selected a host instead of refusing")
	} else if !strings.Contains(err.Error(), "tracker.adapter") {
		t.Errorf("the refusal does not name the field: %v", err)
	}
}

// An address is a setting and not a constant, because a self managed deployment has no
// canonical one. A missing one is a refusal rather than a guess at the hosted service.
func TestTheAddressIsRequiredRatherThanAssumed(t *testing.T) {
	if _, err := BranchRulesFor("github", "", nil); err == nil {
		t.Fatal("an empty base URL was filled in instead of refused")
	} else if !strings.Contains(err.Error(), "tracker.base_url") {
		t.Errorf("the refusal does not name the field: %v", err)
	}
}

func TestAnUnknownAdapterSaysWhatThereIs(t *testing.T) {
	_, err := BranchRulesFor("bitbucket", "https://example.invalid", nil)
	if err == nil {
		t.Fatal("an unknown adapter was accepted")
	}
	if !strings.Contains(err.Error(), "github") {
		t.Errorf("the error does not say what there is: %v", err)
	}
}

func TestTheConfiguredAdapterIsSelected(t *testing.T) {
	r, err := BranchRulesFor("github", "https://example.invalid", nil)
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatal("no adapter was returned")
	}
}

// Every adapter answers the whole of section 12's contract and not only the branch rules
// port. The claim this package makes is that a second host is an entry in the table and a
// package beside github; a table whose entries implemented different halves of the contract
// would make that claim true of nothing in particular.
func TestEveryAdapterAnswersTheWholeContract(t *testing.T) {
	for _, name := range Adapters() {
		h, err := For(model.Tracker{Adapter: name, BaseURL: "https://example.invalid"}, nil)
		if err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		// Most of the assertion is the compiler's: Host carries all three ports, so an
		// adapter missing one is not in the table at all. What is left for a test is that
		// each name yields a value and that the three ports are reachable through it.
		var _ BranchRules = h
		var _ Issues = h
		var _ Comments = h
	}
}

// Credentials is the fourth operation. The token comes from the environment and never from
// project.yaml, so what the file holds is the name of a variable and what this reads is
// what that variable says.
func TestTheTokenComesFromTheVariableTheFileNames(t *testing.T) {
	t.Setenv("XENO_TRACKER_TOKEN", "from-the-environment")
	got, err := Credentials(model.Auth{Scheme: "token", SecretEnv: "XENO_TRACKER_TOKEN"})
	if err != nil {
		t.Fatal(err)
	}
	if got != "from-the-environment" {
		t.Errorf("read %q", got)
	}
}

// An empty variable is a credential nobody renewed, which the plan's credential table
// separates from a tracker nobody configured. The message is what distinguishes them, so
// the name of the variable has to be in it.
func TestAnEmptyVariableSaysWhichOneItIs(t *testing.T) {
	t.Setenv("XENO_TRACKER_TOKEN", "")
	_, err := Credentials(model.Auth{Scheme: "token", SecretEnv: "XENO_TRACKER_TOKEN"})
	if err == nil {
		t.Fatal("an empty variable passed for a token")
	}
	if !strings.Contains(err.Error(), "XENO_TRACKER_TOKEN") {
		t.Errorf("the error does not name the variable: %v", err)
	}
}

func TestAnAbsentSecretEnvNamesTheField(t *testing.T) {
	_, err := Credentials(model.Auth{Scheme: "token"})
	if err == nil {
		t.Fatal("a block naming no variable yielded a token")
	}
	if !strings.Contains(err.Error(), "tracker.auth.secret_env") {
		t.Errorf("the error does not name the field: %v", err)
	}
}

// token is the only scheme either adapter implements, so another one is refused here rather
// than sent as a bearer and failed at the host, where the reason would be somebody else's.
func TestAnUnimplementedSchemeIsRefusedHereRatherThanAtTheHost(t *testing.T) {
	t.Setenv("XENO_TRACKER_TOKEN", "t")
	_, err := Credentials(model.Auth{Scheme: "oauth", SecretEnv: "XENO_TRACKER_TOKEN"})
	if err == nil {
		t.Fatal("a scheme nothing implements was accepted")
	}
}

// Adapters is what the errors above list, and it is sorted so that the message does not
// depend on map iteration order.
func TestAdaptersIsSortedAndNotEmpty(t *testing.T) {
	got := Adapters()
	if len(got) == 0 {
		t.Fatal("no adapters at all")
	}
	for i := 1; i < len(got); i++ {
		if got[i-1] > got[i] {
			t.Fatalf("Adapters is not sorted: %v", got)
		}
	}
}
