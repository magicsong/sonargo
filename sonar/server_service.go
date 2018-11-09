//
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type ServerService struct {
	client *Client
}

// Version Version of SonarQube in plain text
func (s *ServerService) Version() (resp *string, err error) {
	req, err := s.client.NewRequest("GET", "server/version", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
