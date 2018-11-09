// Manage pull request (only available when the Branch plugin is installed)
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type ProjectPullRequestsService struct {
	client *Client
}

type ProjectPullRequests struct {
	PullRequests []struct {
		AnalysisDate string `json:"analysisDate,omitempty"`
		Base         string `json:"base,omitempty"`
		Branch       string `json:"branch,omitempty"`
		Key          string `json:"key,omitempty"`
		Status       struct {
			Bugs              int64  `json:"bugs,omitempty"`
			CodeSmells        int64  `json:"codeSmells,omitempty"`
			QualityGateStatus string `json:"qualityGateStatus,omitempty"`
			Vulnerabilities   int64  `json:"vulnerabilities,omitempty"`
		} `json:"status,omitempty"`
		Title string `json:"title,omitempty"`
		URL   string `json:"url,omitempty"`
	} `json:"pullRequests,omitempty"`
}

type ProjectPullRequestsDeleteOption struct {
	Project     string `url:"project,omitempty"`     // Description:"Project key",ExampleValue:"my_project"
	PullRequest string `url:"pullRequest,omitempty"` // Description:"Pull request id",ExampleValue:"1543"
}

// Delete Delete a pull request.<br/>Requires 'Administer' rights on the specified project.
func (s *ProjectPullRequestsService) Delete(opt *ProjectPullRequestsDeleteOption) (resp *string, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_pull_requests/delete", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectPullRequestsListOption struct {
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// List List the pull requests of a project.<br/>One of the following permissions is required: <ul><li>'Browse' rights on the specified project</li><li>'Execute Analysis' rights on the specified project</li></ul>
func (s *ProjectPullRequestsService) List(opt *ProjectPullRequestsListOption) (resp *ProjectPullRequests, err error) {
	err := s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_pull_requests/list", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
