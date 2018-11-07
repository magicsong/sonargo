// Handle authentication.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type AuthenticationService struct {
	client *Client
}

// [TODO] you should call the <Login> func manually and complete the fields of this struct
type AuthenticationLoginResp struct{}
type AuthenticationLoginOption struct {
	Login    string `url:"login,omitempty"`    // Description:"Login of the user",ExampleValue:""
	Password string `url:"password,omitempty"` // Description:"Password of the user",ExampleValue:""
}

// Login Authenticate a user.
func (s *AuthenticationService) Login(opt *AuthenticationLoginOption) (Resp *AuthenticationLoginResp, err error) {
	err := s.ValidateLoginOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "authentication/login", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Logout> func manually and complete the fields of this struct
type AuthenticationLogoutResp struct{}

// Logout Logout a user.
func (s *AuthenticationService) Logout() (Resp *AuthenticationLogoutResp, err error) {
	req, err := s.client.NewRequest("POST", "authentication/logout", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Validate> func manually and complete the fields of this struct
type AuthenticationValidateResp struct{}

// Validate Check credentials.
func (s *AuthenticationService) Validate() (Resp *AuthenticationValidateResp, err error) {
	req, err := s.client.NewRequest("GET", "authentication/validate", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
