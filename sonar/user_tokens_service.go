// List, create, and delete a user's access tokens.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type UserTokensService struct {
	client *Client
}

type UserTokens struct {
	CreatedAt  string `json:"createdAt,omitempty"`
	Login      string `json:"login,omitempty"`
	Name       string `json:"name,omitempty"`
	Token      string `json:"token,omitempty"`
	UserTokens []struct {
		CreatedAt string `json:"createdAt,omitempty"`
		Name      string `json:"name,omitempty"`
	} `json:"userTokens,omitempty"`
}

type UserTokensGenerateOption struct {
	Login string `url:"login,omitempty"` // Description:"User login. If not set, the token is generated for the authenticated user.",ExampleValue:"g.hopper"
	Name  string `url:"name,omitempty"`  // Description:"Token name",ExampleValue:"Project scan on Travis"
}

// Generate Generate a user access token. <br />Please keep your tokens secret. They enable to authenticate and analyze projects.<br />If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func (s *UserTokensService) Generate(opt *UserTokensGenerateOption) (resp *UserTokens, err error) {
	err := s.ValidateGenerateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_tokens/generate", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type UserTokensRevokeOption struct {
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"g.hopper"
	Name  string `url:"name,omitempty"`  // Description:"Token name",ExampleValue:"Project scan on Travis"
}

// Revoke Revoke a user access token. <br/>If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func (s *UserTokensService) Revoke(opt *UserTokensRevokeOption) (resp *string, err error) {
	err := s.ValidateRevokeOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_tokens/revoke", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type UserTokensSearchOption struct {
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"g.hopper"
}

// Search List the access tokens of a user.<br>The login must exist and active.<br>If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func (s *UserTokensService) Search(opt *UserTokensSearchOption) (resp *UserTokens, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "user_tokens/search", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
