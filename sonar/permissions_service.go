// Manage permission templates, and the granting and revoking of permissions at the global and project levels.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type PermissionsService struct {
	client *Client
}

type Permissions struct {
	DefaultTemplates []struct {
		Qualifier  string `json:"qualifier,omitempty"`
		TemplateID string `json:"templateId,omitempty"`
	} `json:"defaultTemplates,omitempty"`
	Paging struct {
		PageIndex int64 `json:"pageIndex,omitempty"`
		PageSize  int64 `json:"pageSize,omitempty"`
		Total     int64 `json:"total,omitempty"`
	} `json:"paging,omitempty"`
	PermissionTemplate struct {
		CreatedAt         string `json:"createdAt,omitempty"`
		Description       string `json:"description,omitempty"`
		ID                string `json:"id,omitempty"`
		Name              string `json:"name,omitempty"`
		ProjectKeyPattern string `json:"projectKeyPattern,omitempty"`
		UpdatedAt         string `json:"updatedAt,omitempty"`
	} `json:"permissionTemplate,omitempty"`
	PermissionTemplates []struct {
		CreatedAt   string `json:"createdAt,omitempty"`
		Description string `json:"description,omitempty"`
		ID          string `json:"id,omitempty"`
		Name        string `json:"name,omitempty"`
		Permissions []struct {
			GroupsCount        int64  `json:"groupsCount,omitempty"`
			Key                string `json:"key,omitempty"`
			UsersCount         int64  `json:"usersCount,omitempty"`
			WithProjectCreator bool   `json:"withProjectCreator,omitempty"`
		} `json:"permissions,omitempty"`
		ProjectKeyPattern string `json:"projectKeyPattern,omitempty"`
		UpdatedAt         string `json:"updatedAt,omitempty"`
	} `json:"permissionTemplates,omitempty"`
	Permissions []struct {
		Description string `json:"description,omitempty"`
		GroupsCount int64  `json:"groupsCount,omitempty"`
		Key         string `json:"key,omitempty"`
		Name        string `json:"name,omitempty"`
		UsersCount  int64  `json:"usersCount,omitempty"`
	} `json:"permissions,omitempty"`
	Projects []struct {
		ID          string `json:"id,omitempty"`
		Key         string `json:"key,omitempty"`
		Name        string `json:"name,omitempty"`
		Permissions []struct {
			GroupsCount int64  `json:"groupsCount,omitempty"`
			Key         string `json:"key,omitempty"`
			UsersCount  int64  `json:"usersCount,omitempty"`
		} `json:"permissions,omitempty"`
		Qualifier string `json:"qualifier,omitempty"`
	} `json:"projects,omitempty"`
}

type PermissionsAddGroupOption struct {
	GroupId    string `url:"groupId,omitempty"`    // Description:"Group id",ExampleValue:"42"
	GroupName  string `url:"groupName,omitempty"`  // Description:"Group name or 'anyone' (case insensitive)",ExampleValue:"sonar-administrators"
	Permission string `url:"permission,omitempty"` // Description:"Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Add_group Add permission to a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group name or group id must be provided. <br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *PermissionsService) AddGroup(opt *PermissionsAddGroupOption) (resp *string, err error) {
	err := s.ValidateAddGroupOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/add_group", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsAddGroupToTemplateOption struct {
	GroupId      string `url:"groupId,omitempty"`      // Description:"Group id",ExampleValue:"42"
	GroupName    string `url:"groupName,omitempty"`    // Description:"Group name or 'anyone' (case insensitive)",ExampleValue:"sonar-administrators"
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Add_group_to_template Add a group to a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) AddGroupToTemplate(opt *PermissionsAddGroupToTemplateOption) (resp *string, err error) {
	err := s.ValidateAddGroupToTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/add_group_to_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsAddProjectCreatorToTemplateOption struct {
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Add_project_creator_to_template Add a project creator to a permission template.<br>Requires the following permission: 'Administer System'.
func (s *PermissionsService) AddProjectCreatorToTemplate(opt *PermissionsAddProjectCreatorToTemplateOption) (resp *string, err error) {
	err := s.ValidateAddProjectCreatorToTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/add_project_creator_to_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsAddUserOption struct {
	Login      string `url:"login,omitempty"`      // Description:"User login",ExampleValue:"g.hopper"
	Permission string `url:"permission,omitempty"` // Description:"Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Add_user Add permission to a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *PermissionsService) AddUser(opt *PermissionsAddUserOption) (resp *string, err error) {
	err := s.ValidateAddUserOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/add_user", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsAddUserToTemplateOption struct {
	Login        string `url:"login,omitempty"`        // Description:"User login",ExampleValue:"g.hopper"
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Add_user_to_template Add a user to a permission template.<br /> Requires the following permission: 'Administer System'.
func (s *PermissionsService) AddUserToTemplate(opt *PermissionsAddUserToTemplateOption) (resp *string, err error) {
	err := s.ValidateAddUserToTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/add_user_to_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsApplyTemplateOption struct {
	ProjectId    string `url:"projectId,omitempty"`    // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey   string `url:"projectKey,omitempty"`   // Description:"Project key",ExampleValue:"my_project"
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Apply_template Apply a permission template to one project.<br>The project id or project key must be provided.<br>The template id or name must be provided.<br>Requires the following permission: 'Administer System'.
func (s *PermissionsService) ApplyTemplate(opt *PermissionsApplyTemplateOption) (resp *string, err error) {
	err := s.ValidateApplyTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/apply_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsBulkApplyTemplateOption struct {
	AnalyzedBefore    string `url:"analyzedBefore,omitempty"`    // Description:"Filter the projects for which last analysis is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	OnProvisionedOnly string `url:"onProvisionedOnly,omitempty"` // Description:"Filter the projects that are provisioned",ExampleValue:""
	Projects          string `url:"projects,omitempty"`          // Description:"Comma-separated list of project keys",ExampleValue:"my_project,another_project"
	Q                 string `url:"q,omitempty"`                 // Description:"Limit search to: <ul><li>project names that contain the supplied string</li><li>project keys that are exactly the same as the supplied string</li></ul>",ExampleValue:"apac"
	Qualifiers        string `url:"qualifiers,omitempty"`        // Description:"Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>TRK - Projects</li></ul>",ExampleValue:""
	TemplateId        string `url:"templateId,omitempty"`        // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName      string `url:"templateName,omitempty"`      // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Bulk_apply_template Apply a permission template to several projects.<br />The template id or name must be provided.<br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) BulkApplyTemplate(opt *PermissionsBulkApplyTemplateOption) (resp *string, err error) {
	err := s.ValidateBulkApplyTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/bulk_apply_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsCreateTemplateOption struct {
	Description       string `url:"description,omitempty"`       // Description:"Description",ExampleValue:"Permissions for all projects related to the financial service"
	Name              string `url:"name,omitempty"`              // Description:"Name",ExampleValue:"Financial Service Permissions"
	ProjectKeyPattern string `url:"projectKeyPattern,omitempty"` // Description:"Project key pattern. Must be a valid Java regular expression",ExampleValue:".*\.finance\..*"
}

// Create_template Create a permission template.<br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) CreateTemplate(opt *PermissionsCreateTemplateOption) (resp *Permissions, err error) {
	err := s.ValidateCreateTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/create_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsDeleteTemplateOption struct {
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Delete_template Delete a permission template.<br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) DeleteTemplate(opt *PermissionsDeleteTemplateOption) (resp *string, err error) {
	err := s.ValidateDeleteTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/delete_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsRemoveGroupOption struct {
	GroupId    string `url:"groupId,omitempty"`    // Description:"Group id",ExampleValue:"42"
	GroupName  string `url:"groupName,omitempty"`  // Description:"Group name or 'anyone' (case insensitive)",ExampleValue:"sonar-administrators"
	Permission string `url:"permission,omitempty"` // Description:"Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Remove_group Remove a permission from a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group id or group name must be provided, not both.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *PermissionsService) RemoveGroup(opt *PermissionsRemoveGroupOption) (resp *string, err error) {
	err := s.ValidateRemoveGroupOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/remove_group", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsRemoveGroupFromTemplateOption struct {
	GroupId      string `url:"groupId,omitempty"`      // Description:"Group id",ExampleValue:"42"
	GroupName    string `url:"groupName,omitempty"`    // Description:"Group name or 'anyone' (case insensitive)",ExampleValue:"sonar-administrators"
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Remove_group_from_template Remove a group from a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) RemoveGroupFromTemplate(opt *PermissionsRemoveGroupFromTemplateOption) (resp *string, err error) {
	err := s.ValidateRemoveGroupFromTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/remove_group_from_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsRemoveProjectCreatorFromTemplateOption struct {
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Remove_project_creator_from_template Remove a project creator from a permission template.<br>Requires the following permission: 'Administer System'.
func (s *PermissionsService) RemoveProjectCreatorFromTemplate(opt *PermissionsRemoveProjectCreatorFromTemplateOption) (resp *string, err error) {
	err := s.ValidateRemoveProjectCreatorFromTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/remove_project_creator_from_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsRemoveUserOption struct {
	Login      string `url:"login,omitempty"`      // Description:"User login",ExampleValue:"g.hopper"
	Permission string `url:"permission,omitempty"` // Description:"Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Remove_user Remove permission from a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *PermissionsService) RemoveUser(opt *PermissionsRemoveUserOption) (resp *string, err error) {
	err := s.ValidateRemoveUserOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/remove_user", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsRemoveUserFromTemplateOption struct {
	Login        string `url:"login,omitempty"`        // Description:"User login",ExampleValue:"g.hopper"
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Remove_user_from_template Remove a user from a permission template.<br /> Requires the following permission: 'Administer System'.
func (s *PermissionsService) RemoveUserFromTemplate(opt *PermissionsRemoveUserFromTemplateOption) (resp *string, err error) {
	err := s.ValidateRemoveUserFromTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/remove_user_from_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Search_global_permissions List global permissions. <br />Requires the following permission: 'Administer System'
func (s *PermissionsService) SearchGlobalPermissions() (resp *Permissions, err error) {
	req, err := s.client.NewRequest("GET", "permissions/search_global_permissions", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsSearchProjectPermissionsOption struct {
	P          string `url:"p,omitempty"`          // Description:"1-based page number",ExampleValue:"42"
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Ps         string `url:"ps,omitempty"`         // Description:"Page size. Must be greater than 0.",ExampleValue:"20"
	Q          string `url:"q,omitempty"`          // Description:"Limit search to: <ul><li>project names that contain the supplied string</li><li>project keys that are exactly the same as the supplied string</li></ul>",ExampleValue:"apac"
	Qualifier  string `url:"qualifier,omitempty"`  // Description:"Project qualifier. Filter the results with the specified qualifier. Possible values are:<ul><li>TRK - Projects</li></ul>",ExampleValue:""
}

// Search_project_permissions List project permissions. A project can be a technical project, a view or a developer.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *PermissionsService) SearchProjectPermissions(opt *PermissionsSearchProjectPermissionsOption) (resp *Permissions, err error) {
	err := s.ValidateSearchProjectPermissionsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "permissions/search_project_permissions", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsSearchTemplatesOption struct {
	Q string `url:"q,omitempty"` // Description:"Limit search to permission template names that contain the supplied string.",ExampleValue:"defau"
}

// Search_templates List permission templates.<br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) SearchTemplates(opt *PermissionsSearchTemplatesOption) (resp *Permissions, err error) {
	err := s.ValidateSearchTemplatesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "permissions/search_templates", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsSetDefaultTemplateOption struct {
	Qualifier    string `url:"qualifier,omitempty"`    // Description:"Project qualifier. Filter the results with the specified qualifier. Possible values are:<ul><li>TRK - Projects</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Set_default_template Set a permission template as default.<br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) SetDefaultTemplate(opt *PermissionsSetDefaultTemplateOption) (resp *string, err error) {
	err := s.ValidateSetDefaultTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/set_default_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type PermissionsUpdateTemplateOption struct {
	Description       string `url:"description,omitempty"`       // Description:"Description",ExampleValue:"Permissions for all projects related to the financial service"
	Id                string `url:"id,omitempty"`                // Description:"Id",ExampleValue:"af8cb8cc-1e78-4c4e-8c00-ee8e814009a5"
	Name              string `url:"name,omitempty"`              // Description:"Name",ExampleValue:"Financial Service Permissions"
	ProjectKeyPattern string `url:"projectKeyPattern,omitempty"` // Description:"Project key pattern. Must be a valid Java regular expression",ExampleValue:".*\.finance\..*"
}

// Update_template Update a permission template.<br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) UpdateTemplate(opt *PermissionsUpdateTemplateOption) (resp *Permissions, err error) {
	err := s.ValidateUpdateTemplateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/update_template", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
