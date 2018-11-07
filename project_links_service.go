// Manage projects links.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type Project_linksService struct {
	client *Client
}

// [TODO] you should call the <Create> func manually and complete the fields of this struct
type ProjectLinksCreateResp struct{}
type ProjectLinksCreateOption struct {
	Name       string `url:"name,omitempty"`       // Description:"Link name",ExampleValue:"Custom"
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Url        string `url:"url,omitempty"`        // Description:"Link url",ExampleValue:"http://example.com"
}

// Create Create a new project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *Project_linksService) Create(opt *ProjectLinksCreateOption) (Resp *ProjectLinksCreateResp, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_links/create", Opt)
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
type ProjectLinksDeleteResp struct{}
type ProjectLinksDeleteOption struct {
	Id string `url:"id,omitempty"` // Description:"Link id",ExampleValue:"17"
}

// Delete Delete existing project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *Project_linksService) Delete(opt *ProjectLinksDeleteOption) (Resp *ProjectLinksDeleteResp, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_links/delete", Opt)
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
type ProjectLinksSearchResp struct{}
type ProjectLinksSearchOption struct {
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project Id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project Key",ExampleValue:"my_project"
}

// Search List links of a project.<br>The 'projectId' or 'projectKey' must be provided.<br>Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
func (s *Project_linksService) Search(opt *ProjectLinksSearchOption) (Resp *ProjectLinksSearchResp, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_links/search", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
