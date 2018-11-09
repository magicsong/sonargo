// Get the list of programming languages supported in this instance.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type LanguagesService struct {
	client *Client
}

type Languages struct {
	Languages []struct {
		Key  string `json:"key,omitempty"`
		Name string `json:"name,omitempty"`
	} `json:"languages,omitempty"`
}

type LanguagesListOption struct {
	Ps string `url:"ps,omitempty"` // Description:"The size of the list to return, 0 for all languages",ExampleValue:"25"
	Q  string `url:"q,omitempty"`  // Description:"A pattern to match language keys/names against",ExampleValue:"java"
}

// List List supported programming languages
func (s *LanguagesService) List(opt *LanguagesListOption) (resp *Languages, err error) {
	err := s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "languages/list", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
