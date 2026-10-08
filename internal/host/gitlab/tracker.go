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
// description, calls a comment a note, lists its labels as plain strings, calls an open
// milestone active, and answers a write with a payload that carries no web address at all.

// pageSize and pages are the same bound the other adapter holds, for the same reason: an
// approval is usually the last note, and ten pages of a hundred is a bound nobody meets.
const (
	pageSize = 100
	pages    = 10
)

// Issue implements host.Issues, with the labels, the milestone and the notes beside the
// title and the description, and the project's active milestones where the issue carries
// one. A 404 and a 403 are answers: there is no such issue in this project, or the token
// may not read it.
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
		Title       string     `json:"title"`
		Description string     `json:"description"`
		Labels      []string   `json:"labels"`
		Milestone   *milestone `json:"milestone"`
	}
	if err := json.Unmarshal(raw, &body); err != nil {
		return model.Issue{}, err
	}
	issue := model.Issue{Title: body.Title, Body: body.Description, Labels: body.Labels}
	if body.Milestone != nil {
		m := body.Milestone.model()
		issue.Milestone = &m
	}
	if issue.Comments, err = a.notes(project, key, token); err != nil {
		return model.Issue{}, err
	}
	if issue.Milestone != nil {
		if issue.OpenMilestones, err = a.milestones(project, token); err != nil {
			return model.Issue{}, err
		}
	}
	return issue, nil
}

// milestone is this host's shape: due_date is a `YYYY-MM-DD` or null, state is active or
// closed, and iid is the per project number.
type milestone struct {
	Title   string  `json:"title"`
	DueDate *string `json:"due_date"`
	IID     int     `json:"iid"`
	State   string  `json:"state"`
}

func (m milestone) model() model.Milestone {
	out := model.Milestone{Title: m.Title, Number: m.IID, Closed: m.State == "closed"}
	if m.DueDate != nil {
		out.Due = *m.DueDate
	}
	return out
}

// notes reads the issue's notes oldest first, following pages until a short one. A system
// note — this host records a label change or a milestone move as a note too — is left
// out, because what the domain reads is what people wrote.
func (a *Adapter) notes(project, key, token string) ([]model.Comment, error) {
	var out []model.Comment
	for page := 1; page <= pages; page++ {
		raw, status, err := a.request(http.MethodGet, fmt.Sprintf(
			"projects/%s/issues/%s/notes?sort=asc&order_by=created_at&per_page=%d&page=%d",
			ident(project), url.PathEscape(key), pageSize, page), token, nil)
		if err != nil {
			return nil, err
		}
		if status != http.StatusOK {
			return nil, fmt.Errorf("asking for the notes of issue %s in %s answered %d",
				key, project, status)
		}
		var batch []struct {
			Author    struct{ Username string } `json:"author"`
			CreatedAt string                    `json:"created_at"`
			Body      string                    `json:"body"`
			System    bool                      `json:"system"`
		}
		if err := json.Unmarshal(raw, &batch); err != nil {
			return nil, err
		}
		for _, n := range batch {
			if n.System {
				continue
			}
			out = append(out, model.Comment{Author: n.Author.Username, At: n.CreatedAt, Body: n.Body})
		}
		if len(batch) < pageSize {
			break
		}
	}
	return out, nil
}

// milestones reads the project's active milestones, which is this host's word for open.
func (a *Adapter) milestones(project, token string) ([]model.Milestone, error) {
	raw, status, err := a.request(http.MethodGet, fmt.Sprintf(
		"projects/%s/milestones?state=active&per_page=%d", ident(project), pageSize), token, nil)
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

// request is one call for the operations above, separate from get in gitlab.go because
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
