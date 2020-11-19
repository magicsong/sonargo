// Manage project existence.
package sonargo

import "net/http"

type AlmSettingsService struct {
	client *Client
}

type AlmSettingsSetGitlabBindingOption struct {
	AlmSetting string `json:"almSetting,omitempty"` // Description:"GitLab ALM setting key"
	Project    string `json:"project,omitempty"`    // Description: "Project key"
	Repository string `json:"repository,omitempty"` // Description: "GitLab project ID"
}

// Bind a GitLab instance to a project.<br />If the project was already bound to a previous Gitlab ALM instance, the binding will be updated to the new one.Requires the 'Administer' permission on the project
func (s *AlmSettingsService) SetGitlabBinding(opt *AlmSettingsSetGitlabBindingOption) (resp *http.Response, err error) {
	err = s.ValidateSetGitlabBindingOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "alm_settings/set_gitlab_binding", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return resp, err
	}
	return
}
