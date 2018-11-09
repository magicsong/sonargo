// Manage project existence.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type ProjectsService struct {
	client *Client
}

type Projects interface{}

type ProjectsBulkDeleteOption struct {
	AnalyzedBefore    string `url:"analyzedBefore,omitempty"`    // Description:"Filter the projects for which last analysis is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	OnProvisionedOnly string `url:"onProvisionedOnly,omitempty"` // Description:"Filter the projects that are provisioned",ExampleValue:""
	ProjectIds        string `url:"projectIds,omitempty"`        // Description:"Comma-separated list of project ids. Only the 1'000 first ids are used. Others are silently ignored.",ExampleValue:"AU-Tpxb--iU5OvuD2FLy,AU-TpxcA-iU5OvuD2FLz"
	Projects          string `url:"projects,omitempty"`          // Description:"Comma-separated list of project keys",ExampleValue:"my_project,another_project"
	Q                 string `url:"q,omitempty"`                 // Description:"Limit to: <ul><li>component names that contain the supplied string</li><li>component keys that contain the supplied string</li></ul>",ExampleValue:"sonar"
	Qualifiers        string `url:"qualifiers,omitempty"`        // Description:"Comma-separated list of component qualifiers. Filter the results with the specified qualifiers",ExampleValue:""
}

// Bulk_delete Delete one or several projects.<br />Requires 'Administer System' permission.
func (s *ProjectsService) BulkDelete(opt *ProjectsBulkDeleteOption) (resp *string, err error) {
	err := s.ValidateBulkDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/bulk_delete", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectsBulkUpdateKeyOption struct {
	DryRun    string `url:"dryRun,omitempty"`    // Description:"Simulate bulk update. No component key is updated.",ExampleValue:""
	From      string `url:"from,omitempty"`      // Description:"String to match in components keys",ExampleValue:"_old"
	Project   string `url:"project,omitempty"`   // Description:"Project or module key",ExampleValue:"my_old_project"
	ProjectId string `url:"projectId,omitempty"` // Description:"Project or module ID",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	To        string `url:"to,omitempty"`        // Description:"String replacement in components keys",ExampleValue:"_new"
}

// Bulk_update_key Bulk update a project or module key and all its sub-components keys. The bulk update allows to replace a part of the current key by another string on the current project and all its sub-modules.<br>It's possible to simulate the bulk update by setting the parameter 'dryRun' at true. No key is updated with a dry run.<br>Ex: to rename a project with key 'my_project' to 'my_new_project' and all its sub-components keys, call the WS with parameters:<ul>  <li>project: my_project</li>  <li>from: my_</li>  <li>to: my_new_</li></ul>Either 'projectId' or 'project' must be provided.<br> Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *ProjectsService) BulkUpdateKey(opt *ProjectsBulkUpdateKeyOption) (resp *Projects, err error) {
	err := s.ValidateBulkUpdateKeyOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/bulk_update_key", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectsCreateOption struct {
	Branch     string `url:"branch,omitempty"`     // Description:"SCM Branch of the project. The key of the project will become key:branch, for instance 'SonarQube:branch-5.0'",ExampleValue:"branch-5.0"
	Name       string `url:"name,omitempty"`       // Description:"Name of the project",ExampleValue:"SonarQube"
	Project    string `url:"project,omitempty"`    // Description:"Key of the project",ExampleValue:"my_project"
	Visibility string `url:"visibility,omitempty"` // Description:"Whether the created project should be visible to everyone, or only specific user/groups.<br/>If no visibility is specified, the default project visibility of the organization will be used.",ExampleValue:""
}

// Create Create a project.<br/>Requires 'Create Projects' permission
func (s *ProjectsService) Create(opt *ProjectsCreateOption) (resp *Projects, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/create", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectsDeleteOption struct {
	Project   string `url:"project,omitempty"`   // Description:"Project key",ExampleValue:"my_project"
	ProjectId string `url:"projectId,omitempty"` // Description:"Project ID",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
}

// Delete Delete a project.<br> Requires 'Administer System' permission or 'Administer' permission on the project.
func (s *ProjectsService) Delete(opt *ProjectsDeleteOption) (resp *string, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/delete", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectsGhostsOption struct {
	F  string `url:"f,omitempty"`  // Description:"Comma-separated list of the fields to be returned in response. All the fields are returned by default.",ExampleValue:""
	P  string `url:"p,omitempty"`  // Description:"1-based page number",ExampleValue:"42"
	Ps string `url:"ps,omitempty"` // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q  string `url:"q,omitempty"`  // Description:"Limit search to names or keys that contain the supplied string.",ExampleValue:"sonar"
}

// Ghosts List ghost projects.<br> With the current architecture, it's no more possible to have invisible ghost projects. Therefore, the web service is deprecated.<br> Requires 'Administer System' permission.
func (s *ProjectsService) Ghosts(opt *ProjectsGhostsOption) (resp *Projects, err error) {
	err := s.ValidateGhostsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "projects/ghosts", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectsIndexOption struct {
	Desc        string `url:"desc,omitempty"`        // Description:"Since 6.3, this parameter has no effect",ExampleValue:""
	Format      string `url:"format,omitempty"`      // Description:"Only json response format is available",ExampleValue:""
	Libs        string `url:"libs,omitempty"`        // Description:"Since 6.3, this parameter has no effect",ExampleValue:""
	Project     string `url:"project,omitempty"`     // Description:"key or ID of the project",ExampleValue:"my_project"
	Search      string `url:"search,omitempty"`      // Description:"Substring of project name, case insensitive. Ignored if the parameter key is set",ExampleValue:"Sonar"
	Subprojects string `url:"subprojects,omitempty"` // Description:"Load sub-projects. Ignored if the parameter key is set",ExampleValue:""
	Versions    string `url:"versions,omitempty"`    // Description:"Since 6.3, this parameter has no effect",ExampleValue:""
	Views       string `url:"views,omitempty"`       // Description:"Since 6.3, this parameter has no effect",ExampleValue:""
}

// Index This web service is deprecated, please use api/components/search instead
func (s *ProjectsService) Index(opt *ProjectsIndexOption) (resp *Projects, err error) {
	err := s.ValidateIndexOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "projects/index", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectsProvisionedOption struct {
	F  string `url:"f,omitempty"`  // Description:"Comma-separated list of the fields to be returned in response. All the fields are returned by default.",ExampleValue:""
	P  string `url:"p,omitempty"`  // Description:"1-based page number",ExampleValue:"42"
	Ps string `url:"ps,omitempty"` // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q  string `url:"q,omitempty"`  // Description:"Limit search to names or keys that contain the supplied string.",ExampleValue:"sonar"
}

// Provisioned Get the list of provisioned projects.<br> Web service is deprecated. Use api/projects/search instead, with onProvisionedOnly=true.<br> Require 'Create Projects' permission.
func (s *ProjectsService) Provisioned(opt *ProjectsProvisionedOption) (resp *Projects, err error) {
	err := s.ValidateProvisionedOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "projects/provisioned", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectsSearchOption struct {
	AnalyzedBefore    string `url:"analyzedBefore,omitempty"`    // Description:"Filter the projects for which last analysis is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	OnProvisionedOnly string `url:"onProvisionedOnly,omitempty"` // Description:"Filter the projects that are provisioned",ExampleValue:""
	P                 string `url:"p,omitempty"`                 // Description:"1-based page number",ExampleValue:"42"
	ProjectIds        string `url:"projectIds,omitempty"`        // Description:"Comma-separated list of project ids",ExampleValue:"AU-Tpxb--iU5OvuD2FLy,AU-TpxcA-iU5OvuD2FLz"
	Projects          string `url:"projects,omitempty"`          // Description:"Comma-separated list of project keys",ExampleValue:"my_project,another_project"
	Ps                string `url:"ps,omitempty"`                // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q                 string `url:"q,omitempty"`                 // Description:"Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that contain the supplied string</li></ul>",ExampleValue:"sonar"
	Qualifiers        string `url:"qualifiers,omitempty"`        // Description:"Comma-separated list of component qualifiers. Filter the results with the specified qualifiers",ExampleValue:""
}

// Search Search for projects or views to administrate them.<br>Requires 'System Administrator' permission
func (s *ProjectsService) Search(opt *ProjectsSearchOption) (resp *Projects, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "projects/search", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectsUpdateKeyOption struct {
	From      string `url:"from,omitempty"`      // Description:"Project or module key",ExampleValue:"my_old_project"
	ProjectId string `url:"projectId,omitempty"` // Description:"Project or module id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	To        string `url:"to,omitempty"`        // Description:"New component key",ExampleValue:"my_new_project"
}

// Update_key Update a project or module key and all its sub-components keys.<br>Either 'from' or 'projectId' must be provided.<br> Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *ProjectsService) UpdateKey(opt *ProjectsUpdateKeyOption) (resp *string, err error) {
	err := s.ValidateUpdateKeyOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/update_key", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectsUpdateVisibilityOption struct {
	Project    string `url:"project,omitempty"`    // Description:"Project key",ExampleValue:"my_project"
	Visibility string `url:"visibility,omitempty"` // Description:"New visibility",ExampleValue:""
}

// Update_visibility Updates visibility of a project.<br>Requires 'Project administer' permission on the specified project
func (s *ProjectsService) UpdateVisibility(opt *ProjectsUpdateVisibilityOption) (resp *string, err error) {
	err := s.ValidateUpdateVisibilityOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/update_visibility", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
