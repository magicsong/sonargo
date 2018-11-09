// Get information on automatic metrics, and manage custom metrics. See also api/custom_measures.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type MetricsService struct {
	client *Client
}

type Metrics struct {
	Domains []string `json:"domains,omitempty"`
	Metrics []struct {
		Custom      bool   `json:"custom,omitempty"`
		Description string `json:"description,omitempty"`
		Direction   int64  `json:"direction,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Hidden      bool   `json:"hidden,omitempty"`
		ID          string `json:"id,omitempty"`
		Key         string `json:"key,omitempty"`
		Name        string `json:"name,omitempty"`
		Qualitative bool   `json:"qualitative,omitempty"`
		Type        string `json:"type,omitempty"`
	} `json:"metrics,omitempty"`
	P     int64    `json:"p,omitempty"`
	Ps    int64    `json:"ps,omitempty"`
	Total int64    `json:"total,omitempty"`
	Types []string `json:"types,omitempty"`
}

type MetricsCreateOption struct {
	Description string `url:"description,omitempty"` // Description:"Description",ExampleValue:"Size of the team"
	Domain      string `url:"domain,omitempty"`      // Description:"Domain",ExampleValue:"Tests"
	Key         string `url:"key,omitempty"`         // Description:"Key",ExampleValue:"team_size"
	Name        string `url:"name,omitempty"`        // Description:"Name",ExampleValue:"Team Size"
	Type        string `url:"type,omitempty"`        // Description:"Metric type key",ExampleValue:"INT"
}

// Create Create custom metric.<br /> Requires 'Administer System' permission.
func (s *MetricsService) Create(opt *MetricsCreateOption) (resp *string, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "metrics/create", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type MetricsDeleteOption struct {
	Ids  string `url:"ids,omitempty"`  // Description:"Metrics ids to delete.",ExampleValue:"5, 23, 42"
	Keys string `url:"keys,omitempty"` // Description:"Metrics keys to delete",ExampleValue:"team_size, business_value"
}

// Delete Delete metrics and associated measures. Delete only custom metrics.<br />Ids or keys must be provided. <br />Requires 'Administer System' permission.
func (s *MetricsService) Delete(opt *MetricsDeleteOption) (resp *string, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "metrics/delete", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Domains List all custom metric domains.
func (s *MetricsService) Domains() (resp *Metrics, err error) {
	req, err := s.client.NewRequest("GET", "metrics/domains", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type MetricsSearchOption struct {
	F        string `url:"f,omitempty"`        // Description:"Comma-separated list of the fields to be returned in response. All the fields are returned by default.",ExampleValue:""
	IsCustom string `url:"isCustom,omitempty"` // Description:"Choose custom metrics following 3 cases:<ul><li>true: only custom metrics are returned</li><li>false: only non custom metrics are returned</li><li>not specified: all metrics are returned</li></ul>",ExampleValue:"true"
	P        string `url:"p,omitempty"`        // Description:"1-based page number",ExampleValue:"42"
	Ps       string `url:"ps,omitempty"`       // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
}

// Search Search for metrics
func (s *MetricsService) Search(opt *MetricsSearchOption) (resp *Metrics, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "metrics/search", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Types List all available metric types.
func (s *MetricsService) Types() (resp *Metrics, err error) {
	req, err := s.client.NewRequest("GET", "metrics/types", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type MetricsUpdateOption struct {
	Description string `url:"description,omitempty"` // Description:"Description",ExampleValue:"Size of the team"
	Domain      string `url:"domain,omitempty"`      // Description:"Domain",ExampleValue:"Tests"
	Id          string `url:"id,omitempty"`          // Description:"Id of the custom metric to update",ExampleValue:"42"
	Key         string `url:"key,omitempty"`         // Description:"Key",ExampleValue:"team_size"
	Name        string `url:"name,omitempty"`        // Description:"Name",ExampleValue:""
	Type        string `url:"type,omitempty"`        // Description:"Metric type key",ExampleValue:"INT"
}

// Update Update a custom metric.<br /> Requires 'Administer System' permission.
func (s *MetricsService) Update(opt *MetricsUpdateOption) (resp *string, err error) {
	err := s.ValidateUpdateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "metrics/update", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
