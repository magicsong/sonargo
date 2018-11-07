// List, create, and delete a user's access tokens.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type User_tokensService struct {
	client *Client
}

// [TODO] you should call the <Generate> func manually and complete the fields of this struct
type UserTokensGenerateResp struct{}
type UserTokensGenerateOption struct {
	Login string `url:"login,omitempty"` // Description:"User login. If not set, the token is generated for the authenticated user.",ExampleValue:"g.hopper"
	Name  string `url:"name,omitempty"`  // Description:"Token name",ExampleValue:"Project scan on Travis"
}

// Generate Generate a user access token. <br />Please keep your tokens secret. They enable to authenticate and analyze projects.<br />If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func (s *User_tokensService) Generate(opt *UserTokensGenerateOption) (Resp *UserTokensGenerateResp, err error) {
	err := s.ValidateGenerateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_tokens/generate", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Revoke> func manually and complete the fields of this struct
type UserTokensRevokeResp struct{}
type UserTokensRevokeOption struct {
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"g.hopper"
	Name  string `url:"name,omitempty"`  // Description:"Token name",ExampleValue:"Project scan on Travis"
}

// Revoke Revoke a user access token. <br/>If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func (s *User_tokensService) Revoke(opt *UserTokensRevokeOption) (Resp *UserTokensRevokeResp, err error) {
	err := s.ValidateRevokeOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_tokens/revoke", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Search> func manually and complete the fields of this struct
type UserTokensSearchResp struct{}
type UserTokensSearchOption struct {
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"g.hopper"
}

// Search List the access tokens of a user.<br>The login must exist and active.<br>If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func (s *User_tokensService) Search(opt *UserTokensSearchOption) (Resp *UserTokensSearchResp, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "user_tokens/search", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
