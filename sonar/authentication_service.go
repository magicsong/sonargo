// Handle authentication.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type AuthenticationService struct {
	client *Client
}

type Authentication struct {
	Valid bool `json:"valid,omitempty"`
}

type AuthenticationLoginOption struct {
	Login    string `url:"login,omitempty"`    // Description:"Login of the user",ExampleValue:""
	Password string `url:"password,omitempty"` // Description:"Password of the user",ExampleValue:""
}

// Login Authenticate a user.
func (s *AuthenticationService) Login(opt *AuthenticationLoginOption) (resp *string, err error) {
	err := s.ValidateLoginOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "authentication/login", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Logout Logout a user.
func (s *AuthenticationService) Logout() (resp *string, err error) {
	req, err := s.client.NewRequest("POST", "authentication/logout", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Validate Check credentials.
func (s *AuthenticationService) Validate() (resp *Authentication, err error) {
	req, err := s.client.NewRequest("GET", "authentication/validate", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
