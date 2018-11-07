// Generate badges based on quality gates or measures
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type Project_badgesService struct {
	client *Client
}

// [TODO] you should call the <Measure> func manually and complete the fields of this struct
type ProjectBadgesMeasureResp struct{}
type ProjectBadgesMeasureOption struct {
	Branch      string `url:"branch,omitempty"`      // Description:"Branch key",ExampleValue:"feature/my_branch"
	Metric      string `url:"metric,omitempty"`      // Description:"Metric key",ExampleValue:""
	Project     string `url:"project,omitempty"`     // Description:"Project key",ExampleValue:"my_project"
	PullRequest string `url:"pullRequest,omitempty"` // Description:"Pull request id",ExampleValue:"5461"
}

// Measure Generate badge for project's measure as an SVG.<br/>Requires 'Browse' permission on the specified project.
func (s *Project_badgesService) Measure(opt *ProjectBadgesMeasureOption) (Resp *ProjectBadgesMeasureResp, err error) {
	err := s.ValidateMeasureOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_badges/measure", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <QualityGate> func manually and complete the fields of this struct
type ProjectBadgesQualityGateResp struct{}
type ProjectBadgesQualityGateOption struct {
	Branch      string `url:"branch,omitempty"`      // Description:"Branch key",ExampleValue:"feature/my_branch"
	Project     string `url:"project,omitempty"`     // Description:"Project key",ExampleValue:"my_project"
	PullRequest string `url:"pullRequest,omitempty"` // Description:"Pull request id",ExampleValue:"5461"
}

// Quality_gate Generate badge for project's quality gate as an SVG.<br/>Requires 'Browse' permission on the specified project.
func (s *Project_badgesService) Quality_gate(opt *ProjectBadgesQualityGateOption) (Resp *ProjectBadgesQualityGateResp, err error) {
	err := s.ValidateQuality_gateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_badges/quality_gate", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
