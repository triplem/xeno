// SPDX-License-Identifier: Apache-2.0

package gitlab

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

// The tracker half of this host's adapter. It exists for the same reason the branch rules
// half does: section 12 fixes one contract and says a second host is configuration plus an
// adapter, and a contract only half of which has a second implementation has not been
// tested against the claim. Nothing here is shared with the other host's file, and the
// differences are where the claim is worth something — this host numbers an issue by its
// per project iid rather than by a number unique to the repository, calls the body a
// description, calls a comment a note, and answers a write with a payload that carries no
// web address at all.

// Issue implements host.Issues. A 404 and a 403 are answers: there is no such issue in
// this project, or the token may not read it.
func (a *Adapter) Issue(project, key, token string) (model.Issue, error) {
	raw, status, err := a.request(http.MethodGet,
		fmt.Sprintf("projects/%s/issues/%s", ident(project), url.PathEscape(key)), token, nil)
	switch {
	case err != nil:
		return model.Issue{}, err
	case status == http.StatusNotFound:
		return model.Issue{Reason: fmt.Sprintf(
			"there is no issue %s in %s, or the token cannot see the project", key, project)}, nil
	case status == http.StatusForbidden:
		return model.Issue{Reason: fmt.Sprintf(
			"the host refused the read of issue %s in %s", key, project)}, nil
	case status != http.StatusOK:
		return model.Issue{}, fmt.Errorf("asking for issue %s in %s answered %d",
			key, project, status)
	}
	// description rather than body, which is this host's word for the same thing and the
	// reason the field names stay inside each adapter.
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return model.Issue{}, err
	}
	return model.Issue{Title: body.Title, Body: body.Description}, nil
}

// Comment implements host.Comments, as a note on the issue.
//
// The answer carries no web address for the note, so Written.URL stays empty and the
// status is the whole of what says the write happened. That is why the field is documented
// as one host's answer rather than as something the contract promises.
func (a *Adapter) Comment(project, key, token, body string) (model.Written, error) {
	payload, err := json.Marshal(map[string]string{"body": body})
	if err != nil {
		return model.Written{}, err
	}
	_, status, err := a.request(http.MethodPost,
		fmt.Sprintf("projects/%s/issues/%s/notes", ident(project), url.PathEscape(key)),
		token, payload)
	switch {
	case err != nil:
		return model.Written{}, err
	case status == http.StatusForbidden:
		return model.Written{Reason: fmt.Sprintf(
			"the host refused the write to issue %s in %s; that is a token without the api "+
				"scope, or a member without permission to comment", key, project)}, nil
	case status == http.StatusNotFound:
		return model.Written{Reason: fmt.Sprintf(
			"issue %s in %s takes no note: there is no such issue, it is locked, or the "+
				"token cannot see the project", key, project)}, nil
	case status != http.StatusCreated:
		return model.Written{}, fmt.Errorf("writing to issue %s in %s answered %d",
			key, project, status)
	}
	return model.Written{}, nil
}

// request is one call for the two operations above, separate from get in gitlab.go because
// that one reads only and treats every status but 200 as something for its caller to
// phrase; a write answers 201 and nothing here should have to remember that 201 is not an
// error for one endpoint and is for another.
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
	req.Header.Set("Accept", "application/json")
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("PRIVATE-TOKEN", token)
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
