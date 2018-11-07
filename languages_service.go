// Get the list of programming languages supported in this instance.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type LanguagesService struct {
	client *Client
}

// [TODO] you should call the <List> func manually and complete the fields of this struct
type LanguagesListResp struct{}
type LanguagesListOption struct {
	Ps string `url:"ps,omitempty"` // Description:"The size of the list to return, 0 for all languages",ExampleValue:"25"
	Q  string `url:"q,omitempty"`  // Description:"A pattern to match language keys/names against",ExampleValue:"java"
}

// List List supported programming languages
func (s *LanguagesService) List(opt *LanguagesListOption) (Resp *LanguagesListResp, err error) {
	err := s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "languages/list", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
