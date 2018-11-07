// Get components or children with specified measures.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type MeasuresService struct {
	client *Client
}

// [TODO] you should call the <Component> func manually and complete the fields of this struct
type MeasuresComponentResp struct{}
type MeasuresComponentOption struct {
	AdditionalFields string `url:"additionalFields,omitempty"` // Description:"Comma-separated list of additional fields that can be returned in the response.",ExampleValue:"periods,metrics"
	Component        string `url:"component,omitempty"`        // Description:"Component key",ExampleValue:"my_project"
	ComponentId      string `url:"componentId,omitempty"`      // Description:"Component id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	MetricKeys       string `url:"metricKeys,omitempty"`       // Description:"Comma-separated list of metric keys",ExampleValue:"ncloc,complexity,violations"
}

// Component Return component with specified measures. The componentId or the component parameter must be provided.<br>Requires the following permission: 'Browse' on the project of specified component.
func (s *MeasuresService) Component(opt *MeasuresComponentOption) (Resp *MeasuresComponentResp, err error) {
	err := s.ValidateComponentOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "measures/component", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <ComponentTree> func manually and complete the fields of this struct
type MeasuresComponentTreeResp struct{}
type MeasuresComponentTreeOption struct {
	AdditionalFields string `url:"additionalFields,omitempty"` // Description:"Comma-separated list of additional fields that can be returned in the response.",ExampleValue:"periods,metrics"
	Asc              string `url:"asc,omitempty"`              // Description:"Ascending sort",ExampleValue:""
	BaseComponentId  string `url:"baseComponentId,omitempty"`  // Description:"Base component id. The search is based on this component.",ExampleValue:"AU-TpxcA-iU5OvuD2FLz"
	Component        string `url:"component,omitempty"`        // Description:"Component key. The search is based on this component.",ExampleValue:"my_project"
	MetricKeys       string `url:"metricKeys,omitempty"`       // Description:"Comma-separated list of metric keys. Types DISTRIB, DATA are not allowed.",ExampleValue:"ncloc,complexity,violations"
	MetricPeriodSort string `url:"metricPeriodSort,omitempty"` // Description:"Sort measures by leak period or not ?. The 's' parameter must contain the 'metricPeriod' value.",ExampleValue:""
	MetricSort       string `url:"metricSort,omitempty"`       // Description:"Metric key to sort by. The 's' parameter must contain the 'metric' or 'metricPeriod' value. It must be part of the 'metricKeys' parameter",ExampleValue:"ncloc"
	MetricSortFilter string `url:"metricSortFilter,omitempty"` // Description:"Filter components. Sort must be on a metric. Possible values are: <ul><li>all: return all components</li><li>withMeasuresOnly: filter out components that do not have a measure on the sorted metric</li></ul>",ExampleValue:""
	P                string `url:"p,omitempty"`                // Description:"1-based page number",ExampleValue:"42"
	Ps               string `url:"ps,omitempty"`               // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q                string `url:"q,omitempty"`                // Description:"Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li></ul>",ExampleValue:"FILE_NAM"
	Qualifiers       string `url:"qualifiers,omitempty"`       // Description:"Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>BRC - Sub-projects</li><li>DIR - Directories</li><li>FIL - Files</li><li>TRK - Projects</li><li>UTS - Test Files</li></ul>",ExampleValue:""
	S                string `url:"s,omitempty"`                // Description:"Comma-separated list of sort fields",ExampleValue:"name,path"
	Strategy         string `url:"strategy,omitempty"`         // Description:"Strategy to search for base component descendants:<ul><li>children: return the children components of the base component. Grandchildren components are not returned</li><li>all: return all the descendants components of the base component. Grandchildren are returned.</li><li>leaves: return all the descendant components (files, in general) which don't have other children. They are the leaves of the component tree.</li></ul>",ExampleValue:""
}

// Component_tree Navigate through components based on the chosen strategy with specified measures. The baseComponentId or the component parameter must be provided.<br>Requires the following permission: 'Browse' on the specified project.<br>When limiting search with the q parameter, directories are not returned.
func (s *MeasuresService) Component_tree(opt *MeasuresComponentTreeOption) (Resp *MeasuresComponentTreeResp, err error) {
	err := s.ValidateComponent_treeOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "measures/component_tree", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <SearchHistory> func manually and complete the fields of this struct
type MeasuresSearchHistoryResp struct{}
type MeasuresSearchHistoryOption struct {
	Component string `url:"component,omitempty"` // Description:"Component key",ExampleValue:"my_project"
	From      string `url:"from,omitempty"`      // Description:"Filter measures created after the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	Metrics   string `url:"metrics,omitempty"`   // Description:"Comma-separated list of metric keys",ExampleValue:"ncloc,coverage,new_violations"
	P         string `url:"p,omitempty"`         // Description:"1-based page number",ExampleValue:"42"
	Ps        string `url:"ps,omitempty"`        // Description:"Page size. Must be greater than 0 and less or equal than 1000",ExampleValue:"20"
	To        string `url:"to,omitempty"`        // Description:"Filter measures created before the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
}

// Search_history Search measures history of a component.<br>Measures are ordered chronologically.<br>Pagination applies to the number of measures for each metric.<br>Requires the following permission: 'Browse' on the specified component
func (s *MeasuresService) Search_history(opt *MeasuresSearchHistoryOption) (Resp *MeasuresSearchHistoryResp, err error) {
	err := s.ValidateSearch_historyOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "measures/search_history", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
