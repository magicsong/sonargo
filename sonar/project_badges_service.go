// Generate badges based on quality gates or measures
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type ProjectBadgesService struct {
	client *Client
}

type ProjectBadgesMeasureOption struct {
	Branch      string `url:"branch,omitempty"`      // Description:"Branch key",ExampleValue:"feature/my_branch"
	Metric      string `url:"metric,omitempty"`      // Description:"Metric key",ExampleValue:""
	Project     string `url:"project,omitempty"`     // Description:"Project key",ExampleValue:"my_project"
	PullRequest string `url:"pullRequest,omitempty"` // Description:"Pull request id",ExampleValue:"5461"
}

// Measure Generate badge for project's measure as an SVG.<br/>Requires 'Browse' permission on the specified project.
func (s *ProjectBadgesService) Measure(opt *ProjectBadgesMeasureOption) (resp *string, err error) {
	err := s.ValidateMeasureOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_badges/measure", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectBadgesQualityGateOption struct {
	Branch      string `url:"branch,omitempty"`      // Description:"Branch key",ExampleValue:"feature/my_branch"
	Project     string `url:"project,omitempty"`     // Description:"Project key",ExampleValue:"my_project"
	PullRequest string `url:"pullRequest,omitempty"` // Description:"Pull request id",ExampleValue:"5461"
}

// Quality_gate Generate badge for project's quality gate as an SVG.<br/>Requires 'Browse' permission on the specified project.
func (s *ProjectBadgesService) QualityGate(opt *ProjectBadgesQualityGateOption) (resp *string, err error) {
	err := s.ValidateQualityGateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_badges/quality_gate", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
