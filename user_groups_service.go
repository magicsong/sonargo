// Manage user groups.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type User_groupsService struct {
	client *Client
}

// [TODO] you should call the <AddUser> func manually and complete the fields of this struct
type UserGroupsAddUserResp struct{}
type UserGroupsAddUserOption struct {
	Id    string `url:"id,omitempty"`    // Description:"Group id",ExampleValue:"42"
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"g.hopper"
	Name  string `url:"name,omitempty"`  // Description:"Group name",ExampleValue:"sonar-administrators"
}

// Add_user Add a user to a group.<br />'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'.
func (s *User_groupsService) Add_user(opt *UserGroupsAddUserOption) (Resp *UserGroupsAddUserResp, err error) {
	err := s.ValidateAdd_userOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_groups/add_user", Opt)
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
type UserGroupsCreateResp struct{}
type UserGroupsCreateOption struct {
	Description string `url:"description,omitempty"` // Description:"Description for the new group. A group description cannot be larger than 200 characters.",ExampleValue:"Default group for new users"
	Name        string `url:"name,omitempty"`        // Description:"Name for the new group. A group name cannot be larger than 255 characters and must be unique. The value 'anyone' (whatever the case) is reserved and cannot be used.",ExampleValue:"sonar-users"
}

// Create Create a group.<br>Requires the following permission: 'Administer System'.
func (s *User_groupsService) Create(opt *UserGroupsCreateOption) (Resp *UserGroupsCreateResp, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_groups/create", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Delete> func manually and complete the fields of this struct
type UserGroupsDeleteResp struct{}
type UserGroupsDeleteOption struct {
	Id   string `url:"id,omitempty"`   // Description:"Group id",ExampleValue:"42"
	Name string `url:"name,omitempty"` // Description:"Group name",ExampleValue:"sonar-administrators"
}

// Delete Delete a group. The default groups cannot be deleted.<br/>'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'.
func (s *User_groupsService) Delete(opt *UserGroupsDeleteOption) (Resp *UserGroupsDeleteResp, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_groups/delete", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <RemoveUser> func manually and complete the fields of this struct
type UserGroupsRemoveUserResp struct{}
type UserGroupsRemoveUserOption struct {
	Id    string `url:"id,omitempty"`    // Description:"Group id",ExampleValue:"42"
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"g.hopper"
	Name  string `url:"name,omitempty"`  // Description:"Group name",ExampleValue:"sonar-administrators"
}

// Remove_user Remove a user from a group.<br />'id' or 'name' must be provided.<br>Requires the following permission: 'Administer System'.
func (s *User_groupsService) Remove_user(opt *UserGroupsRemoveUserOption) (Resp *UserGroupsRemoveUserResp, err error) {
	err := s.ValidateRemove_userOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_groups/remove_user", Opt)
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
type UserGroupsSearchResp struct{}
type UserGroupsSearchOption struct {
	F  string `url:"f,omitempty"`  // Description:"Comma-separated list of the fields to be returned in response. All the fields are returned by default.",ExampleValue:""
	P  string `url:"p,omitempty"`  // Description:"1-based page number",ExampleValue:"42"
	Ps string `url:"ps,omitempty"` // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q  string `url:"q,omitempty"`  // Description:"Limit search to names that contain the supplied string.",ExampleValue:"sonar-users"
}

// Search Search for user groups.<br>Requires the following permission: 'Administer System'.
func (s *User_groupsService) Search(opt *UserGroupsSearchOption) (Resp *UserGroupsSearchResp, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "user_groups/search", Opt)
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
type UserGroupsUpdateResp struct{}
type UserGroupsUpdateOption struct {
	Description string `url:"description,omitempty"` // Description:"New optional description for the group. A group description cannot be larger than 200 characters. If value is not defined, then description is not changed.",ExampleValue:"Default group for new users"
	Id          string `url:"id,omitempty"`          // Description:"Identifier of the group.",ExampleValue:"42"
	Name        string `url:"name,omitempty"`        // Description:"New optional name for the group. A group name cannot be larger than 255 characters and must be unique. Value 'anyone' (whatever the case) is reserved and cannot be used. If value is empty or not defined, then name is not changed.",ExampleValue:"my-group"
}

// Update Update a group.<br>Requires the following permission: 'Administer System'.
func (s *User_groupsService) Update(opt *UserGroupsUpdateOption) (Resp *UserGroupsUpdateResp, err error) {
	err := s.ValidateUpdateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_groups/update", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Users> func manually and complete the fields of this struct
type UserGroupsUsersResp struct{}
type UserGroupsUsersOption struct {
	Id       string `url:"id,omitempty"`       // Description:"Group id",ExampleValue:"42"
	Name     string `url:"name,omitempty"`     // Description:"Group name",ExampleValue:"sonar-administrators"
	P        string `url:"p,omitempty"`        // Description:"1-based page number",ExampleValue:"42"
	Ps       string `url:"ps,omitempty"`       // Description:"Page size. Must be greater than 0.",ExampleValue:"20"
	Q        string `url:"q,omitempty"`        // Description:"Limit search to names or logins that contain the supplied string.",ExampleValue:"freddy"
	Selected string `url:"selected,omitempty"` // Description:"Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).",ExampleValue:""
}

// Users Search for users with membership information with respect to a group.<br>Requires the following permission: 'Administer System'.
func (s *User_groupsService) Users(opt *UserGroupsUsersOption) (Resp *UserGroupsUsersResp, err error) {
	err := s.ValidateUsersOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "user_groups/users", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
