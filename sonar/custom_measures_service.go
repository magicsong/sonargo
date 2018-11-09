// Manage custom measures for a project. See also api/metrics.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type CustomMeasuresService struct {
	client *Client
}

type CustomMeasures struct {
	CustomMeasures []struct {
		Description string `json:"description,omitempty"`
		Metric      struct {
			Domain string `json:"domain,omitempty"`
			Key    string `json:"key,omitempty"`
			Name   string `json:"name,omitempty"`
			Type   string `json:"type,omitempty"`
		} `json:"metric,omitempty"`
		Pending    bool   `json:"pending,omitempty"`
		ProjectID  string `json:"projectId,omitempty"`
		ProjectKey string `json:"projectKey,omitempty"`
		User       struct {
			Active bool   `json:"active,omitempty"`
			Email  string `json:"email,omitempty"`
			Login  string `json:"login,omitempty"`
			Name   string `json:"name,omitempty"`
		} `json:"user,omitempty"`
		Value string `json:"value,omitempty"`
	} `json:"customMeasures,omitempty"`
	P     int64 `json:"p,omitempty"`
	Ps    int64 `json:"ps,omitempty"`
	Total int64 `json:"total,omitempty"`
}

type CustomMeasuresCreateOption struct {
	Description string `url:"description,omitempty"` // Description:"Description",ExampleValue:"Team size growing."
	MetricId    string `url:"metricId,omitempty"`    // Description:"Metric id",ExampleValue:"16"
	MetricKey   string `url:"metricKey,omitempty"`   // Description:"Metric key",ExampleValue:"ncloc"
	ProjectId   string `url:"projectId,omitempty"`   // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey  string `url:"projectKey,omitempty"`  // Description:"Project key",ExampleValue:"my_project"
	Value       string `url:"value,omitempty"`       // Description:"Measure value. Value type depends on metric type:<ul><li>INT - type: integer</li><li>FLOAT - type: double</li><li>PERCENT - type: double</li><li>BOOL - the possible values are true or false</li><li>STRING - type: string</li><li>MILLISEC - type: integer</li><li>DATA - type: string</li><li>LEVEL - the possible values are OK, WARN, ERROR</li><li>DISTRIB - type: string</li><li>RATING - type: double</li><li>WORK_DUR - long representing the number of minutes</li></ul>",ExampleValue:"47"
}

// Create Create a custom measure.<br /> The project id or the project key must be provided (only project and module custom measures can be created). The metric id or the metric key must be provided.<br/>Requires 'Administer' permission on the project.
func (s *CustomMeasuresService) Create(opt *CustomMeasuresCreateOption) (resp *string, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "custom_measures/create", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type CustomMeasuresDeleteOption struct {
	Id string `url:"id,omitempty"` // Description:"Id",ExampleValue:"24"
}

// Delete Delete a custom measure.<br /> Requires 'Administer System' permission or 'Administer' permission on the project.
func (s *CustomMeasuresService) Delete(opt *CustomMeasuresDeleteOption) (resp *string, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "custom_measures/delete", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type CustomMeasuresSearchOption struct {
	F          string `url:"f,omitempty"`          // Description:"Comma-separated list of the fields to be returned in response. All the fields are returned by default.",ExampleValue:""
	P          string `url:"p,omitempty"`          // Description:"1-based page number",ExampleValue:"42"
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Ps         string `url:"ps,omitempty"`         // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
}

// Search List custom measures. The project id or project key must be provided.<br />Requires 'Administer System' permission or 'Administer' permission on the project.
func (s *CustomMeasuresService) Search(opt *CustomMeasuresSearchOption) (resp *CustomMeasures, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "custom_measures/search", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type CustomMeasuresUpdateOption struct {
	Description string `url:"description,omitempty"` // Description:"",ExampleValue:"Team size growing."
	Id          string `url:"id,omitempty"`          // Description:"id",ExampleValue:"42"
	Value       string `url:"value,omitempty"`       // Description:"Measure value. Value type depends on metric type:<ul><li>INT - type: integer</li><li>FLOAT - type: double</li><li>PERCENT - type: double</li><li>BOOL - the possible values are true or false</li><li>STRING - type: string</li><li>MILLISEC - type: integer</li><li>DATA - type: string</li><li>LEVEL - the possible values are OK, WARN, ERROR</li><li>DISTRIB - type: string</li><li>RATING - type: double</li><li>WORK_DUR - long representing the number of minutes</li></ul>",ExampleValue:"true"
}

// Update Update a custom measure. Value and/or description must be provided<br />Requires 'Administer System' permission or 'Administer' permission on the project.
func (s *CustomMeasuresService) Update(opt *CustomMeasuresUpdateOption) (resp *string, err error) {
	err := s.ValidateUpdateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "custom_measures/update", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
