// Manage pull request (only available when the Branch plugin is installed)
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type Project_pull_requestsService struct {
	client *Client
}

// [TODO] you should call the <Delete> func manually and complete the fields of this struct
type ProjectPullRequestsDeleteResp struct{}
type ProjectPullRequestsDeleteOption struct {
	Project     string `url:"project,omitempty"`     // Description:"Project key",ExampleValue:"my_project"
	PullRequest string `url:"pullRequest,omitempty"` // Description:"Pull request id",ExampleValue:"1543"
}

// Delete Delete a pull request.<br/>Requires 'Administer' rights on the specified project.
func (s *Project_pull_requestsService) Delete(opt *ProjectPullRequestsDeleteOption) (Resp *ProjectPullRequestsDeleteResp, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_pull_requests/delete", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <List> func manually and complete the fields of this struct
type ProjectPullRequestsListResp struct{}
type ProjectPullRequestsListOption struct {
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// List List the pull requests of a project.<br/>One of the following permissions is required: <ul><li>'Browse' rights on the specified project</li><li>'Execute Analysis' rights on the specified project</li></ul>
func (s *Project_pull_requestsService) List(opt *ProjectPullRequestsListOption) (Resp *ProjectPullRequestsListResp, err error) {
	err := s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_pull_requests/list", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
