// Manage user groups.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type UserGroupsService struct {
	client *Client
}

type UserGroups struct {
	Group struct {
		Default      bool   `json:"default,omitempty"`
		Description  string `json:"description,omitempty"`
		ID           string `json:"id,omitempty"`
		MembersCount int64  `json:"membersCount,omitempty"`
		Name         string `json:"name,omitempty"`
		Organization string `json:"organization,omitempty"`
	} `json:"group,omitempty"`
	Groups []struct {
		Default      bool   `json:"default,omitempty"`
		Description  string `json:"description,omitempty"`
		ID           int64  `json:"id,omitempty"`
		MembersCount int64  `json:"membersCount,omitempty"`
		Name         string `json:"name,omitempty"`
	} `json:"groups,omitempty"`
	Paging struct {
		PageIndex int64 `json:"pageIndex,omitempty"`
		PageSize  int64 `json:"pageSize,omitempty"`
		Total     int64 `json:"total,omitempty"`
	} `json:"paging,omitempty"`
	Users []struct {
		Email    string `json:"email,omitempty"`
		Login    string `json:"login,omitempty"`
		Name     string `json:"name,omitempty"`
		Selected bool   `json:"selected,omitempty"`
	} `json:"users,omitempty"`
}

type UserGroupsAddUserOption struct {
	Id    string `url:"id,omitempty"`    // Description:"Group id",ExampleValue:"42"
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"g.hopper"
	Name  string `url:"name,omitempty"`  // Description:"Group name",ExampleValue:"sonar-administrators"
}

// Add_user Add a user to a group.<br />'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'.
func (s *UserGroupsService) AddUser(opt *UserGroupsAddUserOption) (resp *string, err error) {
	err := s.ValidateAddUserOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_groups/add_user", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type UserGroupsCreateOption struct {
	Description string `url:"description,omitempty"` // Description:"Description for the new group. A group description cannot be larger than 200 characters.",ExampleValue:"Default group for new users"
	Name        string `url:"name,omitempty"`        // Description:"Name for the new group. A group name cannot be larger than 255 characters and must be unique. The value 'anyone' (whatever the case) is reserved and cannot be used.",ExampleValue:"sonar-users"
}

// Create Create a group.<br>Requires the following permission: 'Administer System'.
func (s *UserGroupsService) Create(opt *UserGroupsCreateOption) (resp *UserGroups, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_groups/create", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type UserGroupsDeleteOption struct {
	Id   string `url:"id,omitempty"`   // Description:"Group id",ExampleValue:"42"
	Name string `url:"name,omitempty"` // Description:"Group name",ExampleValue:"sonar-administrators"
}

// Delete Delete a group. The default groups cannot be deleted.<br/>'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'.
func (s *UserGroupsService) Delete(opt *UserGroupsDeleteOption) (resp *string, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_groups/delete", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type UserGroupsRemoveUserOption struct {
	Id    string `url:"id,omitempty"`    // Description:"Group id",ExampleValue:"42"
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"g.hopper"
	Name  string `url:"name,omitempty"`  // Description:"Group name",ExampleValue:"sonar-administrators"
}

// Remove_user Remove a user from a group.<br />'id' or 'name' must be provided.<br>Requires the following permission: 'Administer System'.
func (s *UserGroupsService) RemoveUser(opt *UserGroupsRemoveUserOption) (resp *string, err error) {
	err := s.ValidateRemoveUserOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_groups/remove_user", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type UserGroupsSearchOption struct {
	F  string `url:"f,omitempty"`  // Description:"Comma-separated list of the fields to be returned in response. All the fields are returned by default.",ExampleValue:""
	P  string `url:"p,omitempty"`  // Description:"1-based page number",ExampleValue:"42"
	Ps string `url:"ps,omitempty"` // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q  string `url:"q,omitempty"`  // Description:"Limit search to names that contain the supplied string.",ExampleValue:"sonar-users"
}

// Search Search for user groups.<br>Requires the following permission: 'Administer System'.
func (s *UserGroupsService) Search(opt *UserGroupsSearchOption) (resp *UserGroups, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "user_groups/search", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type UserGroupsUpdateOption struct {
	Description string `url:"description,omitempty"` // Description:"New optional description for the group. A group description cannot be larger than 200 characters. If value is not defined, then description is not changed.",ExampleValue:"Default group for new users"
	Id          string `url:"id,omitempty"`          // Description:"Identifier of the group.",ExampleValue:"42"
	Name        string `url:"name,omitempty"`        // Description:"New optional name for the group. A group name cannot be larger than 255 characters and must be unique. Value 'anyone' (whatever the case) is reserved and cannot be used. If value is empty or not defined, then name is not changed.",ExampleValue:"my-group"
}

// Update Update a group.<br>Requires the following permission: 'Administer System'.
func (s *UserGroupsService) Update(opt *UserGroupsUpdateOption) (resp *string, err error) {
	err := s.ValidateUpdateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "user_groups/update", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type UserGroupsUsersOption struct {
	Id       string `url:"id,omitempty"`       // Description:"Group id",ExampleValue:"42"
	Name     string `url:"name,omitempty"`     // Description:"Group name",ExampleValue:"sonar-administrators"
	P        string `url:"p,omitempty"`        // Description:"1-based page number",ExampleValue:"42"
	Ps       string `url:"ps,omitempty"`       // Description:"Page size. Must be greater than 0.",ExampleValue:"20"
	Q        string `url:"q,omitempty"`        // Description:"Limit search to names or logins that contain the supplied string.",ExampleValue:"freddy"
	Selected string `url:"selected,omitempty"` // Description:"Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).",ExampleValue:""
}

// Users Search for users with membership information with respect to a group.<br>Requires the following permission: 'Administer System'.
func (s *UserGroupsService) Users(opt *UserGroupsUsersOption) (resp *UserGroups, err error) {
	err := s.ValidateUsersOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "user_groups/users", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
