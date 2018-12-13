// Copyright 2017 HootSuite Media Inc.
//
// Licensed under the Apache License, Version 2.0 (the License);
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an AS IS BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Modified hereafter by contributors to runatlantis/atlantis.

package vcs

import (
	"fmt"
	"net"
	"net/url"
	"strings"

	"github.com/hashicorp/go-version"
	"github.com/pkg/errors"
	"github.com/runatlantis/atlantis/server/logging"

	"github.com/lkysow/go-gitlab"
	"github.com/runatlantis/atlantis/server/events/models"
)

type GitlabClient struct {
	Client *gitlab.Client
	// Version is set to the server version.
	Version *version.Version
}

// commonMarkSupported is a version constraint that is true when this version of
// GitLab supports CommonMark, a markdown specification.
// See https://about.gitlab.com/2018/07/22/gitlab-11-1-released/
var commonMarkSupported = MustConstraint(">=11.1")

// gitlabClientUnderTest is true if we're running under go test.
var gitlabClientUnderTest = false

// NewGitlabClient returns a valid GitLab client.
func NewGitlabClient(hostname string, token string, logger *logging.SimpleLogger) (*GitlabClient, error) {
	client := &GitlabClient{
		Client: gitlab.NewClient(nil, token),
	}

	// If not using gitlab.com we need to set the URL to the API.
	if hostname != "gitlab.com" {
		// We assume the url will be over HTTPS if the user doesn't specify a scheme.
		absoluteURL := hostname
		if !strings.HasPrefix(hostname, "http://") && !strings.HasPrefix(hostname, "https://") {
			absoluteURL = "https://" + absoluteURL
		}

		url, err := url.Parse(absoluteURL)
		if err != nil {
			return nil, errors.Wrapf(err, "parsing URL %q", absoluteURL)
		}

		// Warn if this hostname isn't resolvable. The GitLab client
		// doesn't give good error messages in this case.
		ips, err := net.LookupIP(url.Hostname())
		if err != nil {
			logger.Warn("unable to resolve %q: %s", url.Hostname(), err)
		} else if len(ips) == 0 {
			logger.Warn("found no IPs while resolving %q", url.Hostname())
		}

		// Now we're ready to construct the client.
		absoluteURL = strings.TrimSuffix(absoluteURL, "/")
		apiURL := fmt.Sprintf("%s/api/v4/", absoluteURL)
		if err := client.Client.SetBaseURL(apiURL); err != nil {
			return nil, errors.Wrapf(err, "setting GitLab API URL: %s", apiURL)
		}
	}

	// Determine which version of GitLab is running.
	if !gitlabClientUnderTest {
		var err error
		client.Version, err = client.GetVersion()
		if err != nil {
			return nil, err
		}
		logger.Info("determined GitLab is running version %s", client.Version.String())
	}

	return client, nil
}

// GetModifiedFiles returns the names of files that were modified in the merge request.
// The names include the path to the file from the repo root, ex. parent/child/file.txt.
func (g *GitlabClient) GetModifiedFiles(repo models.Repo, pull models.PullRequest) ([]string, error) {
	const maxPerPage = 100
	var files []string
	nextPage := 1
	// Constructing the api url by hand so we can do pagination.
	apiURL := fmt.Sprintf("projects/%s/merge_requests/%d/changes", url.QueryEscape(repo.FullName), pull.Num)
	for {
		opts := gitlab.ListOptions{
			Page:    nextPage,
			PerPage: maxPerPage,
		}
		req, err := g.Client.NewRequest("GET", apiURL, opts, nil)
		if err != nil {
			return nil, err
		}
		mr := new(gitlab.MergeRequest)
		resp, err := g.Client.Do(req, mr)
		if err != nil {
			return nil, err
		}

		for _, f := range mr.Changes {
			files = append(files, f.NewPath)
		}
		if resp.NextPage == 0 {
			break
		}
		nextPage = resp.NextPage
	}

	return files, nil
}

// CreateComment creates a comment on the merge request.
func (g *GitlabClient) CreateComment(repo models.Repo, pullNum int, comment string) error {
	_, _, err := g.Client.Notes.CreateMergeRequestNote(repo.FullName, pullNum, &gitlab.CreateMergeRequestNoteOptions{Body: gitlab.String(comment)})
	return err
}

// PullIsApproved returns true if the merge request was approved.
func (g *GitlabClient) PullIsApproved(repo models.Repo, pull models.PullRequest) (bool, error) {
	approvals, _, err := g.Client.MergeRequests.GetMergeRequestApprovals(repo.FullName, pull.Num)
	if err != nil {
		return false, err
	}
	if approvals.ApprovalsLeft > 0 {
		return false, nil
	}
	return true, nil
}

// UpdateStatus updates the build status of a commit.
func (g *GitlabClient) UpdateStatus(repo models.Repo, pull models.PullRequest, state models.CommitStatus, description string) error {
	const statusContext = "Atlantis"

	gitlabState := gitlab.Failed
	switch state {
	case models.PendingCommitStatus:
		gitlabState = gitlab.Pending
	case models.FailedCommitStatus:
		gitlabState = gitlab.Failed
	case models.SuccessCommitStatus:
		gitlabState = gitlab.Success
	}
	_, _, err := g.Client.Commits.SetCommitStatus(repo.FullName, pull.HeadCommit, &gitlab.SetCommitStatusOptions{
		State:       gitlabState,
		Context:     gitlab.String(statusContext),
		Description: gitlab.String(description),
	})
	return err
}

func (g *GitlabClient) GetMergeRequest(repoFullName string, pullNum int) (*gitlab.MergeRequest, error) {
	mr, _, err := g.Client.MergeRequests.GetMergeRequest(repoFullName, pullNum)
	return mr, err
}

// GetVersion returns the version of the Gitlab server this client is using.
func (g *GitlabClient) GetVersion() (*version.Version, error) {
	req, err := g.Client.NewRequest("GET", "/version", nil, nil)
	if err != nil {
		return nil, err
	}
	versionResp := new(gitlab.Version)
	_, err = g.Client.Do(req, versionResp)
	if err != nil {
		return nil, err
	}
	// We need to strip any "-ee" or similar from the resulting version because go-version
	// uses that in its constraints and it breaks the comparison we're trying
	// to do for Common Mark.
	split := strings.Split(versionResp.Version, "-")
	parsedVersion, err := version.NewVersion(split[0])
	if err != nil {
		return nil, errors.Wrapf(err, "parsing response to /version: %q", versionResp.Version)
	}
	return parsedVersion, nil
}

// SupportsCommonMark returns true if the version of Gitlab this client is
// using supports the CommonMark markdown format.
func (g *GitlabClient) SupportsCommonMark() bool {
	// This function is called even if we didn't construct a gitlab client
	// so we need to handle that case.
	if g == nil {
		return false
	}

	return commonMarkSupported.Check(g.Version)
}

// MustConstraint returns a constraint. It panics on error.
func MustConstraint(constraint string) version.Constraints {
	c, err := version.NewConstraint(constraint)
	if err != nil {
		panic(err)
	}
	return c
}
