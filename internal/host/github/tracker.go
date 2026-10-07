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

// Issue implements host.Issues: the issue's title and body, as the person who raised it
// wrote them.
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
		Title string `json:"title"`
		Body  string `json:"body"`
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return model.Issue{}, err
	}
	return model.Issue{Title: body.Title, Body: body.Body}, nil
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

// request is one call for the two operations above. It is separate from transport in
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
