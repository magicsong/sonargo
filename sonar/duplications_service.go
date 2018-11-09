// Get duplication information for a project.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type DuplicationsService struct {
	client *Client
}

type Duplications struct {
	Duplications []struct {
		Blocks []struct {
			Ref  string `json:"_ref,omitempty"`
			From int64  `json:"from,omitempty"`
			Size int64  `json:"size,omitempty"`
		} `json:"blocks,omitempty"`
	} `json:"duplications,omitempty"`
	Files struct {
		One struct {
			Key         string `json:"key,omitempty"`
			Name        string `json:"name,omitempty"`
			ProjectName string `json:"projectName,omitempty"`
		} `json:"1,omitempty"`
		Two struct {
			Key         string `json:"key,omitempty"`
			Name        string `json:"name,omitempty"`
			ProjectName string `json:"projectName,omitempty"`
		} `json:"2,omitempty"`
		Three struct {
			Key         string `json:"key,omitempty"`
			Name        string `json:"name,omitempty"`
			ProjectName string `json:"projectName,omitempty"`
		} `json:"3,omitempty"`
	} `json:"files,omitempty"`
}

type DuplicationsShowOption struct {
	Key  string `url:"key,omitempty"`  // Description:"File key",ExampleValue:"my_project:/src/foo/Bar.php"
	Uuid string `url:"uuid,omitempty"` // Description:"File ID. If provided, 'key' must not be provided.",ExampleValue:"584a89f2-8037-4f7b-b82c-8b45d2d63fb2"
}

// Show Get duplications. Require Browse permission on file's project
func (s *DuplicationsService) Show(opt *DuplicationsShowOption) (resp *Duplications, err error) {
	err := s.ValidateShowOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "duplications/show", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
