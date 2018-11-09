// Manage project analyses.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type ProjectAnalysesService struct {
	client *Client
}

// [TODO] cannot generate resp struct of <api/project_analyses>,you should manual edit the file sonar/project_analyses_service.go,esspecial method response
type ProjectAnalyses struct{}

type ProjectAnalysesCreateEventOption struct {
	Analysis string `url:"analysis,omitempty"` // Description:"Analysis key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Category string `url:"category,omitempty"` // Description:"Category",ExampleValue:""
	Name     string `url:"name,omitempty"`     // Description:"Name",ExampleValue:"5.6"
}

// Create_event Create a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be created.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
func (s *ProjectAnalysesService) CreateEvent(opt *ProjectAnalysesCreateEventOption) (resp *ProjectAnalyses, err error) {
	err := s.ValidateCreateEventOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_analyses/create_event", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectAnalysesDeleteOption struct {
	Analysis string `url:"analysis,omitempty"` // Description:"Analysis key",ExampleValue:"AU-TpxcA-iU5OvuD2FL1"
}

// Delete Delete a project analysis.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the project of the specified analysis</li></ul>
func (s *ProjectAnalysesService) Delete(opt *ProjectAnalysesDeleteOption) (resp *string, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_analyses/delete", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectAnalysesDeleteEventOption struct {
	Event string `url:"event,omitempty"` // Description:"Event key",ExampleValue:"AU-TpxcA-iU5OvuD2FLz"
}

// Delete_event Delete a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be deleted.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
func (s *ProjectAnalysesService) DeleteEvent(opt *ProjectAnalysesDeleteEventOption) (resp *string, err error) {
	err := s.ValidateDeleteEventOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_analyses/delete_event", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectAnalysesSearchOption struct {
	Category string `url:"category,omitempty"` // Description:"Event category. Filter analyses that have at least one event of the category specified.",ExampleValue:"OTHER"
	From     string `url:"from,omitempty"`     // Description:"Filter analyses created after the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided",ExampleValue:"2013-05-01"
	P        string `url:"p,omitempty"`        // Description:"1-based page number",ExampleValue:"42"
	Project  string `url:"project,omitempty"`  // Description:"Project key",ExampleValue:"my_project"
	Ps       string `url:"ps,omitempty"`       // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	To       string `url:"to,omitempty"`       // Description:"Filter analyses created before the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
}

// Search Search a project analyses and attached events.<br>Requires the following permission: 'Browse' on the specified project
func (s *ProjectAnalysesService) Search(opt *ProjectAnalysesSearchOption) (resp *ProjectAnalyses, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_analyses/search", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectAnalysesUpdateEventOption struct {
	Event string `url:"event,omitempty"` // Description:"Event key",ExampleValue:"AU-TpxcA-iU5OvuD2FL5"
	Name  string `url:"name,omitempty"`  // Description:"New name",ExampleValue:"5.6"
}

// Update_event Update a project analysis event.<br>Only events of category 'VERSION' and 'OTHER' can be updated.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
func (s *ProjectAnalysesService) UpdateEvent(opt *ProjectAnalysesUpdateEventOption) (resp *ProjectAnalyses, err error) {
	err := s.ValidateUpdateEventOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_analyses/update_event", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
