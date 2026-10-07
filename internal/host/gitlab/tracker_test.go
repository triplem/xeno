// SPDX-License-Identifier: Apache-2.0

package gitlab

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// seen is what a fake transport recorded about the one request it was given. The path
// matters here more than on the other host: this one takes a URL encoded namespace path
// where the other takes owner and repository as path segments, and a slash left unescaped
// reaches a different endpoint rather than failing.
type seen struct {
	method, path, token string
	body                []byte
}

func answering(t *testing.T, status int, body any) (*Adapter, *seen) {
	t.Helper()
	got := &seen{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.method, got.path = r.Method, r.URL.EscapedPath()
		got.token = r.Header.Get("PRIVATE-TOKEN")
		got.body, _ = io.ReadAll(r.Body)
		w.WriteHeader(status)
		if body != nil {
			_ = json.NewEncoder(w).Encode(body)
		}
	}))
	t.Cleanup(srv.Close)
	return New(srv.Client(), srv.URL), got
}

// description is this host's word for what the other calls the body, which is the kind of
// difference that keeps the field names inside each adapter.
func TestAnIssueIsReadFromThisHostsOwnFieldNames(t *testing.T) {
	a, got := answering(t, http.StatusOK, map[string]string{
		"title":       "The tracker half of WP12 is unbuilt",
		"description": "Found by an audit.",
	})
	issue, err := a.Issue("group/proj", "288", "tok")
	if err != nil {
		t.Fatal(err)
	}
	if issue.Title == "" || issue.Body != "Found by an audit." {
		t.Fatalf("read %+v", issue)
	}
	if got.path != "/projects/group%2Fproj/issues/288" {
		t.Errorf("asked %s, which is not the namespace path encoded", got.path)
	}
	if got.token != "tok" {
		t.Errorf("the token did not reach the request: %q", got.token)
	}
}

func TestAForbiddenOrAbsentIssueIsAnAnswerAndNotAnError(t *testing.T) {
	for _, status := range []int{http.StatusForbidden, http.StatusNotFound} {
		a, _ := answering(t, status, nil)
		issue, err := a.Issue("group/proj", "288", "tok")
		if err != nil {
			t.Fatalf("%d was treated as an error: %v", status, err)
		}
		if issue.Found() || issue.Reason == "" {
			t.Fatalf("%d read as %+v", status, issue)
		}
	}
}

// A note, at the issue's notes endpoint, with the body in this host's own field.
func TestANoteIsOneCallCarryingTheBody(t *testing.T) {
	a, got := answering(t, http.StatusCreated, map[string]any{"id": 7})
	w, err := a.Comment("group/proj", "288", "tok", "## Xeno: 05-review is green")
	if err != nil {
		t.Fatal(err)
	}
	if w.Reason != "" {
		t.Fatalf("a created note reported a reason: %q", w.Reason)
	}
	// This host's answer carries no web address, so the field stays empty rather than
	// holding something composed here that nobody could follow.
	if w.URL != "" {
		t.Errorf("an address appeared where this host sends none: %q", w.URL)
	}
	if got.method != http.MethodPost || got.path != "/projects/group%2Fproj/issues/288/notes" {
		t.Errorf("wrote with %s %s", got.method, got.path)
	}
	var sent map[string]string
	if err := json.Unmarshal(got.body, &sent); err != nil {
		t.Fatal(err)
	}
	if sent["body"] != "## Xeno: 05-review is green" {
		t.Errorf("the host was sent %q", sent["body"])
	}
}

func TestARefusedNoteIsAnAnswerAndARejectedTokenIsNot(t *testing.T) {
	for _, status := range []int{http.StatusForbidden, http.StatusNotFound} {
		a, _ := answering(t, status, nil)
		w, err := a.Comment("group/proj", "288", "tok", "body")
		if err != nil {
			t.Fatalf("%d was treated as an error: %v", status, err)
		}
		if w.Reason == "" {
			t.Fatalf("%d says nothing about why nothing was written", status)
		}
	}
	a, _ := answering(t, http.StatusUnauthorized, nil)
	if _, err := a.Comment("group/proj", "288", "tok", "body"); err == nil {
		t.Error("a rejected token read as a note that was written")
	}
	if _, err := a.Issue("group/proj", "288", "tok"); err == nil {
		t.Error("a rejected token read as an answer about the issue")
	}
}
