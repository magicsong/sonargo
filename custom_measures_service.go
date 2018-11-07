// Manage custom measures for a project. See also api/metrics.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type Custom_measuresService struct {
	client *Client
}

// [TODO] you should call the <Create> func manually and complete the fields of this struct
type CustomMeasuresCreateResp struct{}
type CustomMeasuresCreateOption struct {
	Description string `url:"description,omitempty"` // Description:"Description",ExampleValue:"Team size growing."
	MetricId    string `url:"metricId,omitempty"`    // Description:"Metric id",ExampleValue:"16"
	MetricKey   string `url:"metricKey,omitempty"`   // Description:"Metric key",ExampleValue:"ncloc"
	ProjectId   string `url:"projectId,omitempty"`   // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey  string `url:"projectKey,omitempty"`  // Description:"Project key",ExampleValue:"my_project"
	Value       string `url:"value,omitempty"`       // Description:"Measure value. Value type depends on metric type:<ul><li>INT - type: integer</li><li>FLOAT - type: double</li><li>PERCENT - type: double</li><li>BOOL - the possible values are true or false</li><li>STRING - type: string</li><li>MILLISEC - type: integer</li><li>DATA - type: string</li><li>LEVEL - the possible values are OK, WARN, ERROR</li><li>DISTRIB - type: string</li><li>RATING - type: double</li><li>WORK_DUR - long representing the number of minutes</li></ul>",ExampleValue:"47"
}

// Create Create a custom measure.<br /> The project id or the project key must be provided (only project and module custom measures can be created). The metric id or the metric key must be provided.<br/>Requires 'Administer' permission on the project.
func (s *Custom_measuresService) Create(opt *CustomMeasuresCreateOption) (Resp *CustomMeasuresCreateResp, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "custom_measures/create", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Delete> func manually and complete the fields of this struct
type CustomMeasuresDeleteResp struct{}
type CustomMeasuresDeleteOption struct {
	Id string `url:"id,omitempty"` // Description:"Id",ExampleValue:"24"
}

// Delete Delete a custom measure.<br /> Requires 'Administer System' permission or 'Administer' permission on the project.
func (s *Custom_measuresService) Delete(opt *CustomMeasuresDeleteOption) (Resp *CustomMeasuresDeleteResp, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "custom_measures/delete", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Search> func manually and complete the fields of this struct
type CustomMeasuresSearchResp struct{}
type CustomMeasuresSearchOption struct {
	F          string `url:"f,omitempty"`          // Description:"Comma-separated list of the fields to be returned in response. All the fields are returned by default.",ExampleValue:""
	P          string `url:"p,omitempty"`          // Description:"1-based page number",ExampleValue:"42"
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Ps         string `url:"ps,omitempty"`         // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
}

// Search List custom measures. The project id or project key must be provided.<br />Requires 'Administer System' permission or 'Administer' permission on the project.
func (s *Custom_measuresService) Search(opt *CustomMeasuresSearchOption) (Resp *CustomMeasuresSearchResp, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "custom_measures/search", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Update> func manually and complete the fields of this struct
type CustomMeasuresUpdateResp struct{}
type CustomMeasuresUpdateOption struct {
	Description string `url:"description,omitempty"` // Description:"",ExampleValue:"Team size growing."
	Id          string `url:"id,omitempty"`          // Description:"id",ExampleValue:"42"
	Value       string `url:"value,omitempty"`       // Description:"Measure value. Value type depends on metric type:<ul><li>INT - type: integer</li><li>FLOAT - type: double</li><li>PERCENT - type: double</li><li>BOOL - the possible values are true or false</li><li>STRING - type: string</li><li>MILLISEC - type: integer</li><li>DATA - type: string</li><li>LEVEL - the possible values are OK, WARN, ERROR</li><li>DISTRIB - type: string</li><li>RATING - type: double</li><li>WORK_DUR - long representing the number of minutes</li></ul>",ExampleValue:"true"
}

// Update Update a custom measure. Value and/or description must be provided<br />Requires 'Administer System' permission or 'Administer' permission on the project.
func (s *Custom_measuresService) Update(opt *CustomMeasuresUpdateOption) (Resp *CustomMeasuresUpdateResp, err error) {
	err := s.ValidateUpdateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "custom_measures/update", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
