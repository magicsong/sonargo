// Manage SonarSource commercial editions.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type EditionsService struct {
	client *Client
}

type Editions struct {
	CurrentEditionKey  string `json:"currentEditionKey,omitempty"`
	InstallationStatus string `json:"installationStatus,omitempty"`
	NextEditionKey     string `json:"nextEditionKey,omitempty"`
	PreviewStatus      string `json:"previewStatus,omitempty"`
}

type EditionsApplyLicenseOption struct {
	License string `url:"license,omitempty"` // Description:"the license",ExampleValue:""
}

// Apply_license Apply changes to SonarQube to match the specified license. Clear error message of previous automatic install of an edition, if there is any. Require 'Administer System' permission.
func (s *EditionsService) ApplyLicense(opt *EditionsApplyLicenseOption) (resp *Editions, err error) {
	err := s.ValidateApplyLicenseOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "editions/apply_license", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Clear_error_message Clear error message of last install of an edition (if any). Require 'Administer System' permission.
func (s *EditionsService) ClearErrorMessage() (resp *string, err error) {
	req, err := s.client.NewRequest("POST", "editions/clear_error_message", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type EditionsPreviewOption struct {
	License string `url:"license,omitempty"` // Description:"the license",ExampleValue:""
}

// Preview Preview the changes to SonarQube to match the specified license. Requires 'Administer System' permission.
func (s *EditionsService) Preview(opt *EditionsPreviewOption) (resp *Editions, err error) {
	err := s.ValidatePreviewOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "editions/preview", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Status Provide status of SonarSource commercial edition of the current SonarQube. Requires 'Administer System' permission.
func (s *EditionsService) Status() (resp *Editions, err error) {
	req, err := s.client.NewRequest("GET", "editions/status", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Uninstall Uninstall the currently installed edition. Requires 'Administer System' permission.
func (s *EditionsService) Uninstall() (resp *string, err error) {
	req, err := s.client.NewRequest("POST", "editions/uninstall", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
