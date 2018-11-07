// Get information on the web api supported on this instance.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type WebservicesService struct {
	client *Client
}

// [TODO] you should call the <List> func manually and complete the fields of this struct
type WebservicesListResp struct{}
type WebservicesListOption struct {
	IncludeInternals string `url:"include_internals,omitempty"` // Description:"Include web services that are implemented for internal use only. Their forward-compatibility is not assured",ExampleValue:""
}

// List List web services
func (s *WebservicesService) List(opt *WebservicesListOption) (Resp *WebservicesListResp, err error) {
	err := s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "webservices/list", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <ResponseExample> func manually and complete the fields of this struct
type WebservicesResponseExampleResp struct{}
type WebservicesResponseExampleOption struct {
	Action     string `url:"action,omitempty"`     // Description:"Action of the web service",ExampleValue:"search"
	Controller string `url:"controller,omitempty"` // Description:"Controller of the web service",ExampleValue:"api/issues"
}

// Response_example Display web service response example
func (s *WebservicesService) Response_example(opt *WebservicesResponseExampleOption) (Resp *WebservicesResponseExampleResp, err error) {
	err := s.ValidateResponse_exampleOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "webservices/response_example", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
