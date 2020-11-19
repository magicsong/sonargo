// Manage project existence.
package sonargo

import "net/http"

type AlmIntegrationsService struct {
	client *Client
}

type AlmIntegrationsImportGitlabProjectOption struct {
	AlmSetting      string `json:"almSetting,omitempty"`      // Description:"ALM setting key"
	GitlabProjectID string `json:"gitlabProjectId,omitempty"` // Description: "GitLab project ID"
}

// Import a GitLab project to SonarQube, creating a new project and configuring MR decoration<br />Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) ImportGitlabProject(opt *AlmIntegrationsImportGitlabProjectOption) (resp *http.Response, err error) {
	err = s.ValidateImportGitlabProjectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "alm_integrations/import_gitlab_project", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return nil, resp, err
	}
	return
}
