// SPDX-License-Identifier: Apache-2.0

package github

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// seen is what a fake transport recorded about the one request it was given, so that a test
// can say what this host was asked as well as what it answered.
type seen struct {
	method, path, accept, authorization string
	body                                []byte
}

// answering serves one status and one body to every request, and records the request. It is
// the whole of what these tests need: the adapter makes one call per operation, and what is
// being asserted is the reading of a status and the shape of the call.
func answering(t *testing.T, status int, body any) (*Adapter, *seen) {
	t.Helper()
	got := &seen{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got.method, got.path = r.Method, r.URL.Path
		got.accept, got.authorization = r.Header.Get("Accept"), r.Header.Get("Authorization")
		got.body, _ = io.ReadAll(r.Body)
		w.WriteHeader(status)
		if body != nil {
			_ = json.NewEncoder(w).Encode(body)
		}
	}))
	t.Cleanup(srv.Close)
	return New(srv.Client(), srv.URL), got
}

func TestAnIssueIsReadAsTitleAndBodyUnchanged(t *testing.T) {
	a, got := answering(t, http.StatusOK, map[string]string{
		"title": "The tracker half of WP12 is unbuilt",
		"body":  "Found by an audit.\n\n## Done when\n\nthe issue content reaches P0.",
	})
	issue, err := a.Issue("triplem/xeno", "288", "tok")
	if err != nil {
		t.Fatal(err)
	}
	if !issue.Found() {
		t.Fatalf("an answered read reported a reason: %q", issue.Reason)
	}
	if issue.Title != "The tracker half of WP12 is unbuilt" {
		t.Errorf("title is %q", issue.Title)
	}
	// Unchanged and not reflowed or trimmed: what P0 records is the problem as stated, and a
	// body the adapter had tidied would be a paraphrase with the same provenance.
	if !strings.Contains(issue.Body, "## Done when") {
		t.Errorf("the body lost its structure: %q", issue.Body)
	}
	if got.path != "/repos/triplem/xeno/issues/288" {
		t.Errorf("asked %s", got.path)
	}
	if got.authorization != "Bearer tok" {
		t.Errorf("the token did not reach the request: %q", got.authorization)
	}
}

// Forbidden and absent are answers. The command that asked has a phase to start either
// way, and the reason is what reaches the person who has to go looking.
func TestAForbiddenOrAbsentIssueIsAnAnswerAndNotAnError(t *testing.T) {
	for _, status := range []int{http.StatusForbidden, http.StatusNotFound} {
		a, _ := answering(t, status, nil)
		issue, err := a.Issue("triplem/xeno", "288", "tok")
		if err != nil {
			t.Fatalf("%d was treated as an error: %v", status, err)
		}
		if issue.Found() {
			t.Fatalf("%d read as content", status)
		}
		if issue.Reason == "" {
			t.Fatalf("%d says nothing about why there is no content", status)
		}
	}
}

// A token the host will not accept cannot answer anything, so reporting "no such issue"
// would hide a misconfiguration behind the issue.
func TestARejectedTokenIsAnErrorOnBothOperations(t *testing.T) {
	a, _ := answering(t, http.StatusUnauthorized, nil)
	if _, err := a.Issue("triplem/xeno", "288", "tok"); err == nil {
		t.Error("a rejected token read as an answer about the issue")
	}
	if _, err := a.Comment("triplem/xeno", "288", "tok", "body"); err == nil {
		t.Error("a rejected token read as an answer about the write")
	}
}

func TestACommentIsOneCallCarryingTheBody(t *testing.T) {
	a, got := answering(t, http.StatusCreated,
		map[string]string{"html_url": "https://example.invalid/c/1"})
	w, err := a.Comment("triplem/xeno", "288", "tok", "## Xeno: 05-review is green")
	if err != nil {
		t.Fatal(err)
	}
	if w.Reason != "" {
		t.Fatalf("a created comment reported a reason: %q", w.Reason)
	}
	if w.URL != "https://example.invalid/c/1" {
		t.Errorf("the address of the comment is %q", w.URL)
	}
	if got.method != http.MethodPost || got.path != "/repos/triplem/xeno/issues/288/comments" {
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

// The ordinary case on a run from a fork's pull request: the token is read only. A job that
// failed here would report a red step over a verdict it had already carried.
func TestARefusedWriteIsAnAnswerAndNotAnError(t *testing.T) {
	for _, status := range []int{http.StatusForbidden, http.StatusNotFound, http.StatusGone} {
		a, _ := answering(t, status, nil)
		w, err := a.Comment("triplem/xeno", "288", "tok", "body")
		if err != nil {
			t.Fatalf("%d was treated as an error: %v", status, err)
		}
		if w.Reason == "" {
			t.Fatalf("%d says nothing about why nothing was written", status)
		}
	}
}

// A write that answered something nobody can read is an error, because unlike a refusal it
// says nothing about whether the comment exists.
func TestAnUnreadableAnswerToAWriteIsAnError(t *testing.T) {
	a, _ := answering(t, http.StatusInternalServerError, nil)
	if _, err := a.Comment("triplem/xeno", "288", "tok", "body"); err == nil {
		t.Fatal("a 500 read as a comment that was written")
	}
}

// The created comment exists whether or not the answer carries its address, so an answer
// without one is not a failed write.
func TestACreatedCommentWithNoAddressIsStillWritten(t *testing.T) {
	a, _ := answering(t, http.StatusCreated, map[string]string{})
	w, err := a.Comment("triplem/xeno", "288", "tok", "body")
	if err != nil || w.Reason != "" {
		t.Fatalf("w=%+v err=%v", w, err)
	}
	if w.URL != "" {
		t.Errorf("an address appeared from nowhere: %q", w.URL)
	}
}
