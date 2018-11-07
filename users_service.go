// Manage users.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type UsersService struct {
	client *Client
}

// [TODO] you should call the <ChangePassword> func manually and complete the fields of this struct
type UsersChangePasswordResp struct{}
type UsersChangePasswordOption struct {
	Login            string `url:"login,omitempty"`            // Description:"User login",ExampleValue:"myuser"
	Password         string `url:"password,omitempty"`         // Description:"New password",ExampleValue:"mypassword"
	PreviousPassword string `url:"previousPassword,omitempty"` // Description:"Previous password. Required when changing one's own password.",ExampleValue:"oldpassword"
}

// Change_password Update a user's password. Authenticated users can change their own password, provided that the account is not linked to an external authentication system. Administer System permission is required to change another user's password.
func (s *UsersService) Change_password(opt *UsersChangePasswordOption) (Resp *UsersChangePasswordResp, err error) {
	err := s.ValidateChange_passwordOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/change_password", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Create> func manually and complete the fields of this struct
type UsersCreateResp struct{}
type UsersCreateOption struct {
	Email      string `url:"email,omitempty"`      // Description:"User email",ExampleValue:"myname@email.com"
	Local      string `url:"local,omitempty"`      // Description:"Specify if the user should be authenticated from SonarQube server or from an external authentication system. Password should not be set when local is set to false.",ExampleValue:""
	Login      string `url:"login,omitempty"`      // Description:"User login",ExampleValue:"myuser"
	Name       string `url:"name,omitempty"`       // Description:"User name",ExampleValue:"My Name"
	Password   string `url:"password,omitempty"`   // Description:"User password. Only mandatory when creating local user, otherwise it should not be set",ExampleValue:"mypassword"
	ScmAccount string `url:"scmAccount,omitempty"` // Description:"List of SCM accounts. To set several values, the parameter must be called once for each value.",ExampleValue:"scmAccount=firstValue&scmAccount=secondValue&scmAccount=thirdValue"
}

// Create Create a user.<br/>If a deactivated user account exists with the given login, it will be reactivated.<br/>Requires Administer System permission
func (s *UsersService) Create(opt *UsersCreateOption) (Resp *UsersCreateResp, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/create", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Deactivate> func manually and complete the fields of this struct
type UsersDeactivateResp struct{}
type UsersDeactivateOption struct {
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"myuser"
}

// Deactivate Deactivate a user. Requires Administer System permission
func (s *UsersService) Deactivate(opt *UsersDeactivateOption) (Resp *UsersDeactivateResp, err error) {
	err := s.ValidateDeactivateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/deactivate", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Groups> func manually and complete the fields of this struct
type UsersGroupsResp struct{}
type UsersGroupsOption struct {
	Login    string `url:"login,omitempty"`    // Description:"A user login",ExampleValue:"admin"
	P        string `url:"p,omitempty"`        // Description:"1-based page number",ExampleValue:"42"
	Ps       string `url:"ps,omitempty"`       // Description:"Page size. Must be greater than 0.",ExampleValue:"20"
	Q        string `url:"q,omitempty"`        // Description:"Limit search to group names that contain the supplied string.",ExampleValue:"users"
	Selected string `url:"selected,omitempty"` // Description:"Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).",ExampleValue:""
}

// Groups Lists the groups a user belongs to. <br/>Requires Administer System permission.
func (s *UsersService) Groups(opt *UsersGroupsOption) (Resp *UsersGroupsResp, err error) {
	err := s.ValidateGroupsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "users/groups", Opt)
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
type UsersSearchResp struct{}
type UsersSearchOption struct {
	F  string `url:"f,omitempty"`  // Description:"Comma-separated list of the fields to be returned in response. All the fields are returned by default.",ExampleValue:""
	P  string `url:"p,omitempty"`  // Description:"1-based page number",ExampleValue:"42"
	Ps string `url:"ps,omitempty"` // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q  string `url:"q,omitempty"`  // Description:"Filter on login, name and email",ExampleValue:""
}

// Search Get a list of active users. <br/>Administer System permission is required to show the 'groups' field.<br/>When accessed anonymously, only logins and names are returned.
func (s *UsersService) Search(opt *UsersSearchOption) (Resp *UsersSearchResp, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "users/search", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Update> func manually and complete the fields of this struct
type UsersUpdateResp struct{}
type UsersUpdateOption struct {
	Email      string `url:"email,omitempty"`      // Description:"User email",ExampleValue:"myname@email.com"
	Login      string `url:"login,omitempty"`      // Description:"User login",ExampleValue:"myuser"
	Name       string `url:"name,omitempty"`       // Description:"User name",ExampleValue:"My Name"
	ScmAccount string `url:"scmAccount,omitempty"` // Description:"SCM accounts. To set several values, the parameter must be called once for each value.",ExampleValue:"scmAccount=firstValue&scmAccount=secondValue&scmAccount=thirdValue"
}

// Update Update a user. If a deactivated user account exists with the given login, it will be reactivated. Requires Administer System permission
func (s *UsersService) Update(opt *UsersUpdateOption) (Resp *UsersUpdateResp, err error) {
	err := s.ValidateUpdateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/update", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
