// Manage project tags
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type Project_tagsService struct {
	client *Client
}

// [TODO] you should call the <Search> func manually and complete the fields of this struct
type ProjectTagsSearchResp struct{}
type ProjectTagsSearchOption struct {
	Ps string `url:"ps,omitempty"` // Description:"Page size. Must be greater than 0 and less or equal than 100",ExampleValue:"20"
	Q  string `url:"q,omitempty"`  // Description:"Limit search to tags that contain the supplied string.",ExampleValue:"off"
}

// Search Search tags
func (s *Project_tagsService) Search(opt *ProjectTagsSearchOption) (Resp *ProjectTagsSearchResp, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_tags/search", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Set> func manually and complete the fields of this struct
type ProjectTagsSetResp struct{}
type ProjectTagsSetOption struct {
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Tags    string `url:"tags,omitempty"`    // Description:"Comma-separated list of tags",ExampleValue:"finance, offshore"
}

// Set Set tags on a project.<br>Requires the following permission: 'Administer' rights on the specified project
func (s *Project_tagsService) Set(opt *ProjectTagsSetOption) (Resp *ProjectTagsSetResp, err error) {
	err := s.ValidateSetOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_tags/set", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
