// Get information on the web api supported on this instance.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type WebservicesService struct {
	client *Client
}

type Webservices struct {
	Example     string `json:"example,omitempty"`
	Format      string `json:"format,omitempty"`
	WebServices []struct {
		Actions []struct {
			Changelog []struct {
				Description string `json:"description,omitempty"`
				Version     string `json:"version,omitempty"`
			} `json:"changelog,omitempty"`
			DeprecatedSince    string `json:"deprecatedSince,omitempty"`
			Description        string `json:"description,omitempty"`
			HasResponseExample bool   `json:"hasResponseExample,omitempty"`
			Internal           bool   `json:"internal,omitempty"`
			Key                string `json:"key,omitempty"`
			Params             []struct {
				DefaultValue       string   `json:"defaultValue,omitempty"`
				DeprecatedKey      string   `json:"deprecatedKey,omitempty"`
				DeprecatedKeySince string   `json:"deprecatedKeySince,omitempty"`
				DeprecatedSince    string   `json:"deprecatedSince,omitempty"`
				Description        string   `json:"description,omitempty"`
				ExampleValue       string   `json:"exampleValue,omitempty"`
				Internal           bool     `json:"internal,omitempty"`
				Key                string   `json:"key,omitempty"`
				MaximumLength      int64    `json:"maximumLength,omitempty"`
				MaximumValue       int64    `json:"maximumValue,omitempty"`
				MinimumLength      int64    `json:"minimumLength,omitempty"`
				PossibleValues     []string `json:"possibleValues,omitempty"`
				Required           bool     `json:"required,omitempty"`
				Since              string   `json:"since,omitempty"`
			} `json:"params,omitempty"`
			Post  bool   `json:"post,omitempty"`
			Since string `json:"since,omitempty"`
		} `json:"actions,omitempty"`
		Description string `json:"description,omitempty"`
		Path        string `json:"path,omitempty"`
		Since       string `json:"since,omitempty"`
	} `json:"webServices,omitempty"`
}

type WebservicesListOption struct {
	IncludeInternals string `url:"include_internals,omitempty"` // Description:"Include web services that are implemented for internal use only. Their forward-compatibility is not assured",ExampleValue:""
}

// List List web services
func (s *WebservicesService) List(opt *WebservicesListOption) (resp *Webservices, err error) {
	err := s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "webservices/list", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type WebservicesResponseExampleOption struct {
	Action     string `url:"action,omitempty"`     // Description:"Action of the web service",ExampleValue:"search"
	Controller string `url:"controller,omitempty"` // Description:"Controller of the web service",ExampleValue:"api/issues"
}

// Response_example Display web service response example
func (s *WebservicesService) ResponseExample(opt *WebservicesResponseExampleOption) (resp *Webservices, err error) {
	err := s.ValidateResponseExampleOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "webservices/response_example", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
