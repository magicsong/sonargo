// Get duplication information for a project.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type DuplicationsService struct {
	client *Client
}

// [TODO] you should call the <Show> func manually and complete the fields of this struct
type DuplicationsShowResp struct{}
type DuplicationsShowOption struct {
	Key  string `url:"key,omitempty"`  // Description:"File key",ExampleValue:"my_project:/src/foo/Bar.php"
	Uuid string `url:"uuid,omitempty"` // Description:"File ID. If provided, 'key' must not be provided.",ExampleValue:"584a89f2-8037-4f7b-b82c-8b45d2d63fb2"
}

// Show Get duplications. Require Browse permission on file's project
func (s *DuplicationsService) Show(opt *DuplicationsShowOption) (Resp *DuplicationsShowResp, err error) {
	err := s.ValidateShowOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "duplications/show", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
