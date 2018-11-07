//
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type ServerService struct {
	client *Client
}

// [TODO] you should call the <Version> func manually and complete the fields of this struct
type ServerVersionResp struct{}

// Version Version of SonarQube in plain text
func (s *ServerService) Version() (Resp *ServerVersionResp, err error) {
	req, err := s.client.NewRequest("GET", "server/version", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
