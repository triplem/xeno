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
		// From #330 a read of the issue is followed by a read of what was said under it,
		// which is a list; a fake that serves one body to every call answers that read
		// with an empty one, and records nothing about it, so that the tests about the issue
		// stay about the issue.
		if r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/comments") {
			_ = json.NewEncoder(w).Encode([]any{})
			return
		}
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

// routing answers each path from a map, and records every path asked in order. The one
// read of section 12 is up to three calls from #330, so a transport that serves one body
// to every request cannot assert it.
func routing(t *testing.T, bodies map[string]any) (*Adapter, *[]string) {
	t.Helper()
	asked := &[]string{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*asked = append(*asked, r.URL.Path+"?"+r.URL.RawQuery)
		body, ok := bodies[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		_ = json.NewEncoder(w).Encode(body)
	}))
	t.Cleanup(srv.Close)
	return New(srv.Client(), srv.URL), asked
}

// Section 12's "Starting an intent" reads five things through the one operation: the
// labels, the milestone and the comments of the issue, and the repository's open
// milestones where the issue carries one, in this host's own field names.
func TestAnIssueIsReadWithWhatTheHostSaysAroundIt(t *testing.T) {
	a, asked := routing(t, map[string]any{
		"/repos/triplem/xeno/issues/330": map[string]any{
			"title": "t", "body": "b",
			"labels":    []map[string]string{{"name": "wp12"}, {"name": "xeno-approved"}},
			"milestone": map[string]any{"title": "1.1", "due_on": "2026-12-31T00:00:00Z", "number": 2, "state": "open"},
		},
		"/repos/triplem/xeno/issues/330/comments": []map[string]any{
			{"user": map[string]string{"login": "someone"}, "created_at": "2026-10-08T10:00:00Z", "body": "a question"},
			{"user": map[string]string{"login": "triplem"}, "created_at": "2026-10-08T15:10:28Z", "body": "/xeno approved\n\nbecause"},
		},
		"/repos/triplem/xeno/milestones": []map[string]any{
			{"title": "1.0", "due_on": "2026-11-30T00:00:00Z", "number": 1, "state": "open"},
			{"title": "1.1", "due_on": "2026-12-31T00:00:00Z", "number": 2, "state": "open"},
			{"title": "later", "due_on": nil, "number": 3, "state": "open"},
		},
	})
	issue, err := a.Issue("triplem/xeno", "330", "tok")
	if err != nil {
		t.Fatal(err)
	}
	if strings.Join(issue.Labels, ",") != "wp12,xeno-approved" {
		t.Errorf("labels %v", issue.Labels)
	}
	if issue.Milestone == nil || issue.Milestone.Title != "1.1" || issue.Milestone.Due != "2026-12-31T00:00:00Z" ||
		issue.Milestone.Number != 2 || issue.Milestone.Closed {
		t.Errorf("milestone %+v", issue.Milestone)
	}
	if len(issue.Comments) != 2 || issue.Comments[1].Author != "triplem" ||
		issue.Comments[1].At != "2026-10-08T15:10:28Z" || issue.Comments[1].Body != "/xeno approved\n\nbecause" {
		t.Errorf("comments %+v", issue.Comments)
	}
	if len(issue.OpenMilestones) != 3 || issue.OpenMilestones[2].Due != "" || issue.OpenMilestones[0].Number != 1 {
		t.Errorf("open milestones %+v", issue.OpenMilestones)
	}
	want := []string{
		"/repos/triplem/xeno/issues/330?",
		"/repos/triplem/xeno/issues/330/comments?per_page=100&page=1",
		"/repos/triplem/xeno/milestones?state=open&per_page=100",
	}
	if strings.Join(*asked, " ") != strings.Join(want, " ") {
		t.Errorf("asked %v\nwant  %v", *asked, want)
	}
}

// An issue with no milestone costs two calls and not three, and reads as carrying none.
func TestAnIssueWithoutAMilestoneListsNone(t *testing.T) {
	a, asked := routing(t, map[string]any{
		"/repos/triplem/xeno/issues/330":          map[string]any{"title": "t", "milestone": nil},
		"/repos/triplem/xeno/issues/330/comments": []map[string]any{},
	})
	issue, err := a.Issue("triplem/xeno", "330", "tok")
	if err != nil {
		t.Fatal(err)
	}
	if issue.Milestone != nil || issue.OpenMilestones != nil || len(*asked) != 2 {
		t.Errorf("read %+v after %v", issue, *asked)
	}
}

// Comments are paged at a hundred and the approval is usually the last one, so the read
// follows pages until a short one.
func TestCommentsBeyondTheFirstPageAreFollowed(t *testing.T) {
	full := make([]map[string]any, pageSize)
	for i := range full {
		full[i] = map[string]any{"user": map[string]string{"login": "x"}, "body": "noise"}
	}
	calls := 0
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/comments") {
			_ = json.NewEncoder(w).Encode(map[string]any{"title": "t"})
			return
		}
		calls++
		if r.URL.Query().Get("page") == "1" {
			_ = json.NewEncoder(w).Encode(full)
			return
		}
		_ = json.NewEncoder(w).Encode([]map[string]any{
			{"user": map[string]string{"login": "triplem"}, "body": "/xeno approved"}})
	}))
	t.Cleanup(srv.Close)
	issue, err := New(srv.Client(), srv.URL).Issue("triplem/xeno", "330", "tok")
	if err != nil {
		t.Fatal(err)
	}
	if calls != 2 || len(issue.Comments) != pageSize+1 || issue.Comments[pageSize].Body != "/xeno approved" {
		t.Errorf("%d calls read %d comments", calls, len(issue.Comments))
	}
}
