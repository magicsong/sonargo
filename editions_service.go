// Manage SonarSource commercial editions.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type EditionsService struct {
	client *Client
}

// [TODO] you should call the <ApplyLicense> func manually and complete the fields of this struct
type EditionsApplyLicenseResp struct{}
type EditionsApplyLicenseOption struct {
	License string `url:"license,omitempty"` // Description:"the license",ExampleValue:""
}

// Apply_license Apply changes to SonarQube to match the specified license. Clear error message of previous automatic install of an edition, if there is any. Require 'Administer System' permission.
func (s *EditionsService) Apply_license(opt *EditionsApplyLicenseOption) (Resp *EditionsApplyLicenseResp, err error) {
	err := s.ValidateApply_licenseOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "editions/apply_license", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <ClearErrorMessage> func manually and complete the fields of this struct
type EditionsClearErrorMessageResp struct{}

// Clear_error_message Clear error message of last install of an edition (if any). Require 'Administer System' permission.
func (s *EditionsService) Clear_error_message() (Resp *EditionsClearErrorMessageResp, err error) {
	req, err := s.client.NewRequest("POST", "editions/clear_error_message", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Preview> func manually and complete the fields of this struct
type EditionsPreviewResp struct{}
type EditionsPreviewOption struct {
	License string `url:"license,omitempty"` // Description:"the license",ExampleValue:""
}

// Preview Preview the changes to SonarQube to match the specified license. Requires 'Administer System' permission.
func (s *EditionsService) Preview(opt *EditionsPreviewOption) (Resp *EditionsPreviewResp, err error) {
	err := s.ValidatePreviewOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "editions/preview", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Status> func manually and complete the fields of this struct
type EditionsStatusResp struct{}

// Status Provide status of SonarSource commercial edition of the current SonarQube. Requires 'Administer System' permission.
func (s *EditionsService) Status() (Resp *EditionsStatusResp, err error) {
	req, err := s.client.NewRequest("GET", "editions/status", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Uninstall> func manually and complete the fields of this struct
type EditionsUninstallResp struct{}

// Uninstall Uninstall the currently installed edition. Requires 'Administer System' permission.
func (s *EditionsService) Uninstall() (Resp *EditionsUninstallResp, err error) {
	req, err := s.client.NewRequest("POST", "editions/uninstall", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
