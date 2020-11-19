// Manage project existence.
package sonargo

import "net/http"

type AlmIntegrationsService struct {
	client *Client
}

type AlmIntegrationsProjectObject struct {
	Key        string `json:"key,omitempty"`
	Name       string `json:"name,omitempty"`
	Qualifier  string `json:"qualifier,omitempty"`
	Visibility string `json:"visibility,omitempty"`
}

type AlmIntegrationsImportGitlabProjectObject struct {
	Project *Project `json:"project,omitempty"`
}

type AlmIntegrationsImportGitlabProjectOption struct {
	AlmSetting      string `url:"almSetting,omitempty"`      // Description:"ALM setting key"
	GitlabProjectID string `url:"gitlabProjectId,omitempty"` // Description: "GitLab project ID"
}

// Import a GitLab project to SonarQube, creating a new project and configuring MR decoration<br />Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) ImportGitlabProject(opt *AlmIntegrationsImportGitlabProjectOption) (v *AlmIntegrationsImportGitlabProjectObject, resp *http.Response, err error) {
	err = s.ValidateImportGitlabProjectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "alm_integrations/import_gitlab_project", opt)
	if err != nil {
		return
	}
	v = new(AlmIntegrationsImportGitlabProjectObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return v, resp, err
	}
	return
}
