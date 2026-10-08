// SPDX-License-Identifier: Apache-2.0

package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/triplem/xeno/internal/model"
)

// The tracker half of this host's adapter: reading an issue and writing a comment. It is a
// file of its own beside the branch protection reader because the two answer different
// questions with different credentials — the plan's credential table gives them a row
// each — and because what they share is an address and a header rather than any logic.

// pageSize is what this host pages lists at, and pages is how many of them one read
// follows. An approval is usually the last comment, so the first page is not enough; ten
// pages is a bound rather than a limit anybody expects to meet.
const (
	pageSize = 100
	pages    = 10
)

// Issue implements host.Issues: the issue's title and body, as the person who raised it
// wrote them, with the labels, the milestone and the comments beside them, and the
// repository's open milestones where the issue carries one. Section 12 calls that one
// operation and up to three calls, and this is where the three are made.
//
// A 404 and a 403 are answers and not errors. On this host the two are not cleanly
// separable, since a repository a token cannot see answers 404 rather than 403 on purpose,
// so the sentence for each says which of them happened without claiming to know why. A 401
// is an error, because a token the host will not accept cannot answer anything and
// reporting "no such issue" would hide a misconfiguration behind the issue.
func (a *Adapter) Issue(project, key, token string) (model.Issue, error) {
	raw, status, err := a.request(http.MethodGet,
		fmt.Sprintf("repos/%s/issues/%s", project, url.PathEscape(key)), token, nil)
	switch {
	case err != nil:
		return model.Issue{}, err
	case status == http.StatusNotFound:
		return model.Issue{Reason: fmt.Sprintf(
			"there is no issue %s in %s, or the token cannot see the repository; this host "+
				"answers alike to both", key, project)}, nil
	case status == http.StatusForbidden:
		return model.Issue{Reason: fmt.Sprintf(
			"the host refused the read of issue %s in %s", key, project)}, nil
	case status != http.StatusOK:
		return model.Issue{}, fmt.Errorf("asking for issue %s in %s answered %d",
			key, project, status)
	}
	var body struct {
		Title     string                  `json:"title"`
		Body      string                  `json:"body"`
		Labels    []struct{ Name string } `json:"labels"`
		Milestone *milestone              `json:"milestone"`
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return model.Issue{}, err
	}
	issue := model.Issue{Title: body.Title, Body: body.Body}
	for _, l := range body.Labels {
		issue.Labels = append(issue.Labels, l.Name)
	}
	if body.Milestone != nil {
		m := body.Milestone.model()
		issue.Milestone = &m
	}
	if issue.Comments, err = a.comments(project, key, token); err != nil {
		return model.Issue{}, err
	}
	// The milestones are read only where the issue carries one, because the ordering is
	// asked about in no other case and a read nobody consults is a call for nothing.
	if issue.Milestone != nil {
		if issue.OpenMilestones, err = a.milestones(project, token); err != nil {
			return model.Issue{}, err
		}
	}
	return issue, nil
}

// milestone is this host's shape of one: due_on is an RFC 3339 stamp or null, state is
// open or closed, and number is what the host orders them by.
type milestone struct {
	Title  string  `json:"title"`
	DueOn  *string `json:"due_on"`
	Number int     `json:"number"`
	State  string  `json:"state"`
}

func (m milestone) model() model.Milestone {
	out := model.Milestone{Title: m.Title, Number: m.Number, Closed: m.State == "closed"}
	if m.DueOn != nil {
		out.Due = *m.DueOn
	}
	return out
}

// comments reads the issue's comments, oldest first as the host lists them, following
// pages until a short one. Only the three fields the domain reads are kept.
func (a *Adapter) comments(project, key, token string) ([]model.Comment, error) {
	var out []model.Comment
	for page := 1; page <= pages; page++ {
		raw, status, err := a.request(http.MethodGet, fmt.Sprintf(
			"repos/%s/issues/%s/comments?per_page=%d&page=%d",
			project, url.PathEscape(key), pageSize, page), token, nil)
		if err != nil {
			return nil, err
		}
		if status != http.StatusOK {
			return nil, fmt.Errorf("asking for the comments of issue %s in %s answered %d",
				key, project, status)
		}
		var batch []struct {
			User      struct{ Login string } `json:"user"`
			CreatedAt string                 `json:"created_at"`
			Body      string                 `json:"body"`
		}
		if err := json.Unmarshal(raw, &batch); err != nil {
			return nil, err
		}
		for _, c := range batch {
			out = append(out, model.Comment{Author: c.User.Login, At: c.CreatedAt, Body: c.Body})
		}
		if len(batch) < pageSize {
			break
		}
	}
	return out, nil
}

// milestones reads the repository's open milestones. One page: a repository with more
// than a hundred open milestones is not ordering its work by them.
func (a *Adapter) milestones(project, token string) ([]model.Milestone, error) {
	raw, status, err := a.request(http.MethodGet, fmt.Sprintf(
		"repos/%s/milestones?state=open&per_page=%d", project, pageSize), token, nil)
	if err != nil {
		return nil, err
	}
	if status != http.StatusOK {
		return nil, fmt.Errorf("asking for the milestones of %s answered %d", project, status)
	}
	var batch []milestone
	if err := json.Unmarshal(raw, &batch); err != nil {
		return nil, err
	}
	out := make([]model.Milestone, 0, len(batch))
	for _, m := range batch {
		out = append(out, m.model())
	}
	return out, nil
}

// Comment implements host.Comments: one note on the issue, written and not read back.
//
// 201 is this host's answer to a write that happened. 403 is a token without the issues
// write permission, which is what a pull request from a fork gets and is the ordinary case
// rather than a fault, so it is a reason and not an error.
func (a *Adapter) Comment(project, key, token, body string) (model.Written, error) {
	payload, err := json.Marshal(map[string]string{"body": body})
	if err != nil {
		return model.Written{}, err
	}
	raw, status, err := a.request(http.MethodPost,
		fmt.Sprintf("repos/%s/issues/%s/comments", project, url.PathEscape(key)),
		token, payload)
	switch {
	case err != nil:
		return model.Written{}, err
	case status == http.StatusForbidden:
		return model.Written{Reason: fmt.Sprintf(
			"the host refused the write to issue %s in %s; that is a token without the "+
				"issues write permission, which is what a run on a fork's pull request has",
			key, project)}, nil
	case status == http.StatusNotFound || status == http.StatusGone:
		return model.Written{Reason: fmt.Sprintf(
			"issue %s in %s takes no comment: there is no such issue, it is locked, or the "+
				"token cannot see the repository", key, project)}, nil
	case status != http.StatusCreated:
		return model.Written{}, fmt.Errorf("writing to issue %s in %s answered %d",
			key, project, status)
	}
	var out struct {
		HTMLURL string `json:"html_url"`
	}
	// The address is a convenience and its absence is not a failure of the write: the
	// comment exists either way, and the status is what says so.
	_ = json.Unmarshal(raw, &out)
	return model.Written{URL: out.HTMLURL}, nil
}

// request is one call for the operations above. It is separate from transport in
// github.go because that one carries the branch protection endpoint's reading of each
// status, which is this host's pricing rather than a rule about requests; what the two
// share is the address, the Accept header and the bearer scheme.
//
// A nil body is a read. The status is returned rather than judged, because what a status
// means differs per endpoint and only the caller knows which one it asked.
func (a *Adapter) request(method, path, token string, body []byte) ([]byte, int, error) {
	u := fmt.Sprintf("%s/%s", strings.TrimRight(a.baseURL, "/"), path)
	var payload io.Reader
	if body != nil {
		payload = bytes.NewReader(body)
	}
	req, err := http.NewRequest(method, u, payload)
	if err != nil {
		return nil, 0, err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return nil, 0, fmt.Errorf("asking %s: %w", u, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusUnauthorized {
		return nil, resp.StatusCode, fmt.Errorf("the token was rejected by %s", u)
	}
	raw, err := io.ReadAll(resp.Body)
	return raw, resp.StatusCode, err
}
