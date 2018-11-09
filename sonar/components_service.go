// Get information about a component (file, directory, project, ...) and its ancestors or descendants. Update a project or module key.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type ComponentsService struct {
	client *Client
}

type Components struct {
	Ancestors []struct {
		AnalysisDate string   `json:"analysisDate,omitempty"`
		Description  string   `json:"description,omitempty"`
		ID           string   `json:"id,omitempty"`
		Key          string   `json:"key,omitempty"`
		Name         string   `json:"name,omitempty"`
		Organization string   `json:"organization,omitempty"`
		Path         string   `json:"path,omitempty"`
		Qualifier    string   `json:"qualifier,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		Version      string   `json:"version,omitempty"`
		Visibility   string   `json:"visibility,omitempty"`
	} `json:"ancestors,omitempty"`
	BaseComponent struct {
		ID           string `json:"id,omitempty"`
		Key          string `json:"key,omitempty"`
		Name         string `json:"name,omitempty"`
		Organization string `json:"organization,omitempty"`
		Qualifier    string `json:"qualifier,omitempty"`
	} `json:"baseComponent,omitempty"`
	Component struct {
		AnalysisDate   string `json:"analysisDate,omitempty"`
		ID             string `json:"id,omitempty"`
		Key            string `json:"key,omitempty"`
		Language       string `json:"language,omitempty"`
		LeakPeriodDate string `json:"leakPeriodDate,omitempty"`
		Name           string `json:"name,omitempty"`
		Organization   string `json:"organization,omitempty"`
		Path           string `json:"path,omitempty"`
		Qualifier      string `json:"qualifier,omitempty"`
		Version        string `json:"version,omitempty"`
	} `json:"component,omitempty"`
	Components []struct {
		ID           string `json:"id,omitempty"`
		Key          string `json:"key,omitempty"`
		Language     string `json:"language,omitempty"`
		Name         string `json:"name,omitempty"`
		Organization string `json:"organization,omitempty"`
		Path         string `json:"path,omitempty"`
		Project      string `json:"project,omitempty"`
		Qualifier    string `json:"qualifier,omitempty"`
	} `json:"components,omitempty"`
	Paging struct {
		PageIndex int64 `json:"pageIndex,omitempty"`
		PageSize  int64 `json:"pageSize,omitempty"`
		Total     int64 `json:"total,omitempty"`
	} `json:"paging,omitempty"`
}

type ComponentsSearchOption struct {
	Language   string `url:"language,omitempty"`   // Description:"Language key. If provided, only components for the given language are returned.",ExampleValue:"py"
	P          string `url:"p,omitempty"`          // Description:"1-based page number",ExampleValue:"42"
	Ps         string `url:"ps,omitempty"`         // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q          string `url:"q,omitempty"`          // Description:"Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li></ul>",ExampleValue:"sonar"
	Qualifiers string `url:"qualifiers,omitempty"` // Description:"Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>BRC - Sub-projects</li><li>DIR - Directories</li><li>FIL - Files</li><li>TRK - Projects</li><li>UTS - Test Files</li></ul>",ExampleValue:""
}

// Search Search for components
func (s *ComponentsService) Search(opt *ComponentsSearchOption) (resp *Components, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "components/search", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ComponentsShowOption struct {
	Component   string `url:"component,omitempty"`   // Description:"Component key",ExampleValue:"my_project"
	ComponentId string `url:"componentId,omitempty"` // Description:"Component id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
}

// Show Returns a component (file, directory, project, view…) and its ancestors. The ancestors are ordered from the parent to the root project. The 'componentId' or 'component' parameter must be provided.<br>Requires the following permission: 'Browse' on the project of the specified component.
func (s *ComponentsService) Show(opt *ComponentsShowOption) (resp *Components, err error) {
	err := s.ValidateShowOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "components/show", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ComponentsTreeOption struct {
	Asc         string `url:"asc,omitempty"`         // Description:"Ascending sort",ExampleValue:""
	Component   string `url:"component,omitempty"`   // Description:"Base component key. The search is based on this component.",ExampleValue:"my_project"
	ComponentId string `url:"componentId,omitempty"` // Description:"Base component id. The search is based on this component.",ExampleValue:"AU-TpxcA-iU5OvuD2FLz"
	P           string `url:"p,omitempty"`           // Description:"1-based page number",ExampleValue:"42"
	Ps          string `url:"ps,omitempty"`          // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q           string `url:"q,omitempty"`           // Description:"Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li></ul>",ExampleValue:"FILE_NAM"
	Qualifiers  string `url:"qualifiers,omitempty"`  // Description:"Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>BRC - Sub-projects</li><li>DIR - Directories</li><li>FIL - Files</li><li>TRK - Projects</li><li>UTS - Test Files</li></ul>",ExampleValue:""
	S           string `url:"s,omitempty"`           // Description:"Comma-separated list of sort fields",ExampleValue:"name, path"
	Strategy    string `url:"strategy,omitempty"`    // Description:"Strategy to search for base component descendants:<ul><li>children: return the children components of the base component. Grandchildren components are not returned</li><li>all: return all the descendants components of the base component. Grandchildren are returned.</li><li>leaves: return all the descendant components (files, in general) which don't have other children. They are the leaves of the component tree.</li></ul>",ExampleValue:""
}

// Tree Navigate through components based on the chosen strategy. The componentId or the component parameter must be provided.<br>Requires the following permission: 'Browse' on the specified project.<br>When limiting search with the q parameter, directories are not returned.
func (s *ComponentsService) Tree(opt *ComponentsTreeOption) (resp *Components, err error) {
	err := s.ValidateTreeOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "components/tree", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
