// SPDX-License-Identifier: Apache-2.0

package host

import (
	"strings"
	"testing"
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
