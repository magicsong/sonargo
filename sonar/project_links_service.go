// Manage projects links.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type ProjectLinksService struct {
	client *Client
}

type ProjectLinks struct {
	Link struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"link,omitempty"`
	Links []struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
		Type string `json:"type,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"links,omitempty"`
}

type ProjectLinksCreateOption struct {
	Name       string `url:"name,omitempty"`       // Description:"Link name",ExampleValue:"Custom"
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Url        string `url:"url,omitempty"`        // Description:"Link url",ExampleValue:"http://example.com"
}

// Create Create a new project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *ProjectLinksService) Create(opt *ProjectLinksCreateOption) (resp *ProjectLinks, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_links/create", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectLinksDeleteOption struct {
	Id string `url:"id,omitempty"` // Description:"Link id",ExampleValue:"17"
}

// Delete Delete existing project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *ProjectLinksService) Delete(opt *ProjectLinksDeleteOption) (resp *string, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_links/delete", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type ProjectLinksSearchOption struct {
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project Id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project Key",ExampleValue:"my_project"
}

// Search List links of a project.<br>The 'projectId' or 'projectKey' must be provided.<br>Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
func (s *ProjectLinksService) Search(opt *ProjectLinksSearchOption) (resp *ProjectLinks, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_links/search", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
