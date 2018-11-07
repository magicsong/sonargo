// Manage permission templates, and the granting and revoking of permissions at the global and project levels.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type PermissionsService struct {
	client *Client
}

// [TODO] you should call the <AddGroup> func manually and complete the fields of this struct
type PermissionsAddGroupResp struct{}
type PermissionsAddGroupOption struct {
	GroupId    string `url:"groupId,omitempty"`    // Description:"Group id",ExampleValue:"42"
	GroupName  string `url:"groupName,omitempty"`  // Description:"Group name or 'anyone' (case insensitive)",ExampleValue:"sonar-administrators"
	Permission string `url:"permission,omitempty"` // Description:"Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Add_group Add permission to a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group name or group id must be provided. <br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *PermissionsService) Add_group(opt *PermissionsAddGroupOption) (Resp *PermissionsAddGroupResp, err error) {
	err := s.ValidateAdd_groupOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/add_group", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <AddGroupToTemplate> func manually and complete the fields of this struct
type PermissionsAddGroupToTemplateResp struct{}
type PermissionsAddGroupToTemplateOption struct {
	GroupId      string `url:"groupId,omitempty"`      // Description:"Group id",ExampleValue:"42"
	GroupName    string `url:"groupName,omitempty"`    // Description:"Group name or 'anyone' (case insensitive)",ExampleValue:"sonar-administrators"
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Add_group_to_template Add a group to a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) Add_group_to_template(opt *PermissionsAddGroupToTemplateOption) (Resp *PermissionsAddGroupToTemplateResp, err error) {
	err := s.ValidateAdd_group_to_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/add_group_to_template", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <AddProjectCreatorToTemplate> func manually and complete the fields of this struct
type PermissionsAddProjectCreatorToTemplateResp struct{}
type PermissionsAddProjectCreatorToTemplateOption struct {
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Add_project_creator_to_template Add a project creator to a permission template.<br>Requires the following permission: 'Administer System'.
func (s *PermissionsService) Add_project_creator_to_template(opt *PermissionsAddProjectCreatorToTemplateOption) (Resp *PermissionsAddProjectCreatorToTemplateResp, err error) {
	err := s.ValidateAdd_project_creator_to_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/add_project_creator_to_template", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <AddUser> func manually and complete the fields of this struct
type PermissionsAddUserResp struct{}
type PermissionsAddUserOption struct {
	Login      string `url:"login,omitempty"`      // Description:"User login",ExampleValue:"g.hopper"
	Permission string `url:"permission,omitempty"` // Description:"Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Add_user Add permission to a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *PermissionsService) Add_user(opt *PermissionsAddUserOption) (Resp *PermissionsAddUserResp, err error) {
	err := s.ValidateAdd_userOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/add_user", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <AddUserToTemplate> func manually and complete the fields of this struct
type PermissionsAddUserToTemplateResp struct{}
type PermissionsAddUserToTemplateOption struct {
	Login        string `url:"login,omitempty"`        // Description:"User login",ExampleValue:"g.hopper"
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Add_user_to_template Add a user to a permission template.<br /> Requires the following permission: 'Administer System'.
func (s *PermissionsService) Add_user_to_template(opt *PermissionsAddUserToTemplateOption) (Resp *PermissionsAddUserToTemplateResp, err error) {
	err := s.ValidateAdd_user_to_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/add_user_to_template", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <ApplyTemplate> func manually and complete the fields of this struct
type PermissionsApplyTemplateResp struct{}
type PermissionsApplyTemplateOption struct {
	ProjectId    string `url:"projectId,omitempty"`    // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey   string `url:"projectKey,omitempty"`   // Description:"Project key",ExampleValue:"my_project"
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Apply_template Apply a permission template to one project.<br>The project id or project key must be provided.<br>The template id or name must be provided.<br>Requires the following permission: 'Administer System'.
func (s *PermissionsService) Apply_template(opt *PermissionsApplyTemplateOption) (Resp *PermissionsApplyTemplateResp, err error) {
	err := s.ValidateApply_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/apply_template", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <BulkApplyTemplate> func manually and complete the fields of this struct
type PermissionsBulkApplyTemplateResp struct{}
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
func (s *PermissionsService) Bulk_apply_template(opt *PermissionsBulkApplyTemplateOption) (Resp *PermissionsBulkApplyTemplateResp, err error) {
	err := s.ValidateBulk_apply_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/bulk_apply_template", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <CreateTemplate> func manually and complete the fields of this struct
type PermissionsCreateTemplateResp struct{}
type PermissionsCreateTemplateOption struct {
	Description       string `url:"description,omitempty"`       // Description:"Description",ExampleValue:"Permissions for all projects related to the financial service"
	Name              string `url:"name,omitempty"`              // Description:"Name",ExampleValue:"Financial Service Permissions"
	ProjectKeyPattern string `url:"projectKeyPattern,omitempty"` // Description:"Project key pattern. Must be a valid Java regular expression",ExampleValue:".*\.finance\..*"
}

// Create_template Create a permission template.<br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) Create_template(opt *PermissionsCreateTemplateOption) (Resp *PermissionsCreateTemplateResp, err error) {
	err := s.ValidateCreate_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/create_template", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <DeleteTemplate> func manually and complete the fields of this struct
type PermissionsDeleteTemplateResp struct{}
type PermissionsDeleteTemplateOption struct {
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Delete_template Delete a permission template.<br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) Delete_template(opt *PermissionsDeleteTemplateOption) (Resp *PermissionsDeleteTemplateResp, err error) {
	err := s.ValidateDelete_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/delete_template", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <RemoveGroup> func manually and complete the fields of this struct
type PermissionsRemoveGroupResp struct{}
type PermissionsRemoveGroupOption struct {
	GroupId    string `url:"groupId,omitempty"`    // Description:"Group id",ExampleValue:"42"
	GroupName  string `url:"groupName,omitempty"`  // Description:"Group name or 'anyone' (case insensitive)",ExampleValue:"sonar-administrators"
	Permission string `url:"permission,omitempty"` // Description:"Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Remove_group Remove a permission from a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group id or group name must be provided, not both.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *PermissionsService) Remove_group(opt *PermissionsRemoveGroupOption) (Resp *PermissionsRemoveGroupResp, err error) {
	err := s.ValidateRemove_groupOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/remove_group", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <RemoveGroupFromTemplate> func manually and complete the fields of this struct
type PermissionsRemoveGroupFromTemplateResp struct{}
type PermissionsRemoveGroupFromTemplateOption struct {
	GroupId      string `url:"groupId,omitempty"`      // Description:"Group id",ExampleValue:"42"
	GroupName    string `url:"groupName,omitempty"`    // Description:"Group name or 'anyone' (case insensitive)",ExampleValue:"sonar-administrators"
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Remove_group_from_template Remove a group from a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) Remove_group_from_template(opt *PermissionsRemoveGroupFromTemplateOption) (Resp *PermissionsRemoveGroupFromTemplateResp, err error) {
	err := s.ValidateRemove_group_from_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/remove_group_from_template", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <RemoveProjectCreatorFromTemplate> func manually and complete the fields of this struct
type PermissionsRemoveProjectCreatorFromTemplateResp struct{}
type PermissionsRemoveProjectCreatorFromTemplateOption struct {
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Remove_project_creator_from_template Remove a project creator from a permission template.<br>Requires the following permission: 'Administer System'.
func (s *PermissionsService) Remove_project_creator_from_template(opt *PermissionsRemoveProjectCreatorFromTemplateOption) (Resp *PermissionsRemoveProjectCreatorFromTemplateResp, err error) {
	err := s.ValidateRemove_project_creator_from_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/remove_project_creator_from_template", Opt)
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
type PermissionsRemoveUserResp struct{}
type PermissionsRemoveUserOption struct {
	Login      string `url:"login,omitempty"`      // Description:"User login",ExampleValue:"g.hopper"
	Permission string `url:"permission,omitempty"` // Description:"Permission<ul><li>Possible values for global permissions: admin, profileadmin, gateadmin, scan, provisioning</li><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Remove_user Remove permission from a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *PermissionsService) Remove_user(opt *PermissionsRemoveUserOption) (Resp *PermissionsRemoveUserResp, err error) {
	err := s.ValidateRemove_userOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/remove_user", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <RemoveUserFromTemplate> func manually and complete the fields of this struct
type PermissionsRemoveUserFromTemplateResp struct{}
type PermissionsRemoveUserFromTemplateOption struct {
	Login        string `url:"login,omitempty"`        // Description:"User login",ExampleValue:"g.hopper"
	Permission   string `url:"permission,omitempty"`   // Description:"Permission<ul><li>Possible values for project permissions admin, codeviewer, issueadmin, scan, user</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Remove_user_from_template Remove a user from a permission template.<br /> Requires the following permission: 'Administer System'.
func (s *PermissionsService) Remove_user_from_template(opt *PermissionsRemoveUserFromTemplateOption) (Resp *PermissionsRemoveUserFromTemplateResp, err error) {
	err := s.ValidateRemove_user_from_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/remove_user_from_template", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <SearchGlobalPermissions> func manually and complete the fields of this struct
type PermissionsSearchGlobalPermissionsResp struct{}

// Search_global_permissions List global permissions. <br />Requires the following permission: 'Administer System'
func (s *PermissionsService) Search_global_permissions() (Resp *PermissionsSearchGlobalPermissionsResp, err error) {
	req, err := s.client.NewRequest("GET", "permissions/search_global_permissions", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <SearchProjectPermissions> func manually and complete the fields of this struct
type PermissionsSearchProjectPermissionsResp struct{}
type PermissionsSearchProjectPermissionsOption struct {
	P          string `url:"p,omitempty"`          // Description:"1-based page number",ExampleValue:"42"
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"ce4c03d6-430f-40a9-b777-ad877c00aa4d"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Ps         string `url:"ps,omitempty"`         // Description:"Page size. Must be greater than 0.",ExampleValue:"20"
	Q          string `url:"q,omitempty"`          // Description:"Limit search to: <ul><li>project names that contain the supplied string</li><li>project keys that are exactly the same as the supplied string</li></ul>",ExampleValue:"apac"
	Qualifier  string `url:"qualifier,omitempty"`  // Description:"Project qualifier. Filter the results with the specified qualifier. Possible values are:<ul><li>TRK - Projects</li></ul>",ExampleValue:""
}

// Search_project_permissions List project permissions. A project can be a technical project, a view or a developer.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *PermissionsService) Search_project_permissions(opt *PermissionsSearchProjectPermissionsOption) (Resp *PermissionsSearchProjectPermissionsResp, err error) {
	err := s.ValidateSearch_project_permissionsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "permissions/search_project_permissions", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <SearchTemplates> func manually and complete the fields of this struct
type PermissionsSearchTemplatesResp struct{}
type PermissionsSearchTemplatesOption struct {
	Q string `url:"q,omitempty"` // Description:"Limit search to permission template names that contain the supplied string.",ExampleValue:"defau"
}

// Search_templates List permission templates.<br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) Search_templates(opt *PermissionsSearchTemplatesOption) (Resp *PermissionsSearchTemplatesResp, err error) {
	err := s.ValidateSearch_templatesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "permissions/search_templates", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <SetDefaultTemplate> func manually and complete the fields of this struct
type PermissionsSetDefaultTemplateResp struct{}
type PermissionsSetDefaultTemplateOption struct {
	Qualifier    string `url:"qualifier,omitempty"`    // Description:"Project qualifier. Filter the results with the specified qualifier. Possible values are:<ul><li>TRK - Projects</li></ul>",ExampleValue:""
	TemplateId   string `url:"templateId,omitempty"`   // Description:"Template id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	TemplateName string `url:"templateName,omitempty"` // Description:"Template name",ExampleValue:"Default Permission Template for Projects"
}

// Set_default_template Set a permission template as default.<br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) Set_default_template(opt *PermissionsSetDefaultTemplateOption) (Resp *PermissionsSetDefaultTemplateResp, err error) {
	err := s.ValidateSet_default_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/set_default_template", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <UpdateTemplate> func manually and complete the fields of this struct
type PermissionsUpdateTemplateResp struct{}
type PermissionsUpdateTemplateOption struct {
	Description       string `url:"description,omitempty"`       // Description:"Description",ExampleValue:"Permissions for all projects related to the financial service"
	Id                string `url:"id,omitempty"`                // Description:"Id",ExampleValue:"af8cb8cc-1e78-4c4e-8c00-ee8e814009a5"
	Name              string `url:"name,omitempty"`              // Description:"Name",ExampleValue:"Financial Service Permissions"
	ProjectKeyPattern string `url:"projectKeyPattern,omitempty"` // Description:"Project key pattern. Must be a valid Java regular expression",ExampleValue:".*\.finance\..*"
}

// Update_template Update a permission template.<br />Requires the following permission: 'Administer System'.
func (s *PermissionsService) Update_template(opt *PermissionsUpdateTemplateOption) (Resp *PermissionsUpdateTemplateResp, err error) {
	err := s.ValidateUpdate_templateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "permissions/update_template", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
