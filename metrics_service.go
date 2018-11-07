// Get information on automatic metrics, and manage custom metrics. See also api/custom_measures.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type MetricsService struct {
	client *Client
}

// [TODO] you should call the <Create> func manually and complete the fields of this struct
type MetricsCreateResp struct{}
type MetricsCreateOption struct {
	Description string `url:"description,omitempty"` // Description:"Description",ExampleValue:"Size of the team"
	Domain      string `url:"domain,omitempty"`      // Description:"Domain",ExampleValue:"Tests"
	Key         string `url:"key,omitempty"`         // Description:"Key",ExampleValue:"team_size"
	Name        string `url:"name,omitempty"`        // Description:"Name",ExampleValue:"Team Size"
	Type        string `url:"type,omitempty"`        // Description:"Metric type key",ExampleValue:"INT"
}

// Create Create custom metric.<br /> Requires 'Administer System' permission.
func (s *MetricsService) Create(opt *MetricsCreateOption) (Resp *MetricsCreateResp, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "metrics/create", Opt)
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
type MetricsDeleteResp struct{}
type MetricsDeleteOption struct {
	Ids  string `url:"ids,omitempty"`  // Description:"Metrics ids to delete.",ExampleValue:"5, 23, 42"
	Keys string `url:"keys,omitempty"` // Description:"Metrics keys to delete",ExampleValue:"team_size, business_value"
}

// Delete Delete metrics and associated measures. Delete only custom metrics.<br />Ids or keys must be provided. <br />Requires 'Administer System' permission.
func (s *MetricsService) Delete(opt *MetricsDeleteOption) (Resp *MetricsDeleteResp, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "metrics/delete", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Domains> func manually and complete the fields of this struct
type MetricsDomainsResp struct{}

// Domains List all custom metric domains.
func (s *MetricsService) Domains() (Resp *MetricsDomainsResp, err error) {
	req, err := s.client.NewRequest("GET", "metrics/domains", Opt)
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
type MetricsSearchResp struct{}
type MetricsSearchOption struct {
	F        string `url:"f,omitempty"`        // Description:"Comma-separated list of the fields to be returned in response. All the fields are returned by default.",ExampleValue:""
	IsCustom string `url:"isCustom,omitempty"` // Description:"Choose custom metrics following 3 cases:<ul><li>true: only custom metrics are returned</li><li>false: only non custom metrics are returned</li><li>not specified: all metrics are returned</li></ul>",ExampleValue:"true"
	P        string `url:"p,omitempty"`        // Description:"1-based page number",ExampleValue:"42"
	Ps       string `url:"ps,omitempty"`       // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
}

// Search Search for metrics
func (s *MetricsService) Search(opt *MetricsSearchOption) (Resp *MetricsSearchResp, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "metrics/search", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Types> func manually and complete the fields of this struct
type MetricsTypesResp struct{}

// Types List all available metric types.
func (s *MetricsService) Types() (Resp *MetricsTypesResp, err error) {
	req, err := s.client.NewRequest("GET", "metrics/types", Opt)
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
type MetricsUpdateResp struct{}
type MetricsUpdateOption struct {
	Description string `url:"description,omitempty"` // Description:"Description",ExampleValue:"Size of the team"
	Domain      string `url:"domain,omitempty"`      // Description:"Domain",ExampleValue:"Tests"
	Id          string `url:"id,omitempty"`          // Description:"Id of the custom metric to update",ExampleValue:"42"
	Key         string `url:"key,omitempty"`         // Description:"Key",ExampleValue:"team_size"
	Name        string `url:"name,omitempty"`        // Description:"Name",ExampleValue:""
	Type        string `url:"type,omitempty"`        // Description:"Metric type key",ExampleValue:"INT"
}

// Update Update a custom metric.<br /> Requires 'Administer System' permission.
func (s *MetricsService) Update(opt *MetricsUpdateOption) (Resp *MetricsUpdateResp, err error) {
	err := s.ValidateUpdateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "metrics/update", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
