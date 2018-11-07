// Manage quality gates, including conditions and project association.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type QualitygatesService struct {
	client *Client
}

// [TODO] you should call the <Copy> func manually and complete the fields of this struct
type QualitygatesCopyResp struct{}
type QualitygatesCopyOption struct {
	Id           string `url:"id,omitempty"`           // Description:"The ID of the source quality gate",ExampleValue:"1"
	Name         string `url:"name,omitempty"`         // Description:"The name of the quality gate to create",ExampleValue:"My Quality Gate"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Copy Copy a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Copy(opt *QualitygatesCopyOption) (Resp *QualitygatesCopyResp, err error) {
	err := s.ValidateCopyOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/copy", Opt)
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
type QualitygatesCreateResp struct{}
type QualitygatesCreateOption struct {
	Name         string `url:"name,omitempty"`         // Description:"The name of the quality gate to create",ExampleValue:"My Quality Gate"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Create Create a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Create(opt *QualitygatesCreateOption) (Resp *QualitygatesCreateResp, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/create", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <CreateCondition> func manually and complete the fields of this struct
type QualitygatesCreateConditionResp struct{}
type QualitygatesCreateConditionOption struct {
	Error        string `url:"error,omitempty"`        // Description:"Condition error threshold",ExampleValue:"10"
	GateId       string `url:"gateId,omitempty"`       // Description:"ID of the quality gate",ExampleValue:"1"
	Metric       string `url:"metric,omitempty"`       // Description:"Condition metric",ExampleValue:"blocker_violations"
	Op           string `url:"op,omitempty"`           // Description:"Condition operator:<br/><ul><li>EQ = equals</li><li>NE = is not</li><li>LT = is lower than</li><li>GT = is greater than</li></ui>",ExampleValue:"EQ"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
	Period       string `url:"period,omitempty"`       // Description:"Condition period. If not set, the absolute value is considered.",ExampleValue:""
	Warning      string `url:"warning,omitempty"`      // Description:"Condition warning threshold",ExampleValue:"5"
}

// Create_condition Add a new condition to a quality gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Create_condition(opt *QualitygatesCreateConditionOption) (Resp *QualitygatesCreateConditionResp, err error) {
	err := s.ValidateCreate_conditionOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/create_condition", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <DeleteCondition> func manually and complete the fields of this struct
type QualitygatesDeleteConditionResp struct{}
type QualitygatesDeleteConditionOption struct {
	Id           string `url:"id,omitempty"`           // Description:"Condition ID",ExampleValue:"2"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Delete_condition Delete a condition from a quality gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Delete_condition(opt *QualitygatesDeleteConditionOption) (Resp *QualitygatesDeleteConditionResp, err error) {
	err := s.ValidateDelete_conditionOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/delete_condition", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Deselect> func manually and complete the fields of this struct
type QualitygatesDeselectResp struct{}
type QualitygatesDeselectOption struct {
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
	ProjectId    string `url:"projectId,omitempty"`    // Description:"Project id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ProjectKey   string `url:"projectKey,omitempty"`   // Description:"Project key",ExampleValue:"my_project"
}

// Deselect Remove the association of a project from a quality gate.<br>Requires one of the following permissions:<ul><li>'Administer Quality Gates'</li><li>'Administer' rights on the project</li></ul>
func (s *QualitygatesService) Deselect(opt *QualitygatesDeselectOption) (Resp *QualitygatesDeselectResp, err error) {
	err := s.ValidateDeselectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/deselect", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Destroy> func manually and complete the fields of this struct
type QualitygatesDestroyResp struct{}
type QualitygatesDestroyOption struct {
	Id           string `url:"id,omitempty"`           // Description:"ID of the quality gate to delete",ExampleValue:"1"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Destroy Delete a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Destroy(opt *QualitygatesDestroyOption) (Resp *QualitygatesDestroyResp, err error) {
	err := s.ValidateDestroyOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/destroy", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <GetByProject> func manually and complete the fields of this struct
type QualitygatesGetByProjectResp struct{}
type QualitygatesGetByProjectOption struct {
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
	Project      string `url:"project,omitempty"`      // Description:"Project key",ExampleValue:"my_project"
}

// Get_by_project Get the quality gate of a project.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
func (s *QualitygatesService) Get_by_project(opt *QualitygatesGetByProjectOption) (Resp *QualitygatesGetByProjectResp, err error) {
	err := s.ValidateGet_by_projectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualitygates/get_by_project", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <List> func manually and complete the fields of this struct
type QualitygatesListResp struct{}
type QualitygatesListOption struct {
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// List Get a list of quality gates
func (s *QualitygatesService) List(opt *QualitygatesListOption) (Resp *QualitygatesListResp, err error) {
	err := s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualitygates/list", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <ProjectStatus> func manually and complete the fields of this struct
type QualitygatesProjectStatusResp struct{}
type QualitygatesProjectStatusOption struct {
	AnalysisId string `url:"analysisId,omitempty"` // Description:"Analysis id",ExampleValue:"AU-TpxcA-iU5OvuD2FL1"
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Project_status Get the quality gate status of a project or a Compute Engine task.<br />Either 'analysisId', 'projectId' or 'projectKey' must be provided<br />The different statuses returned are: OK, WARN, ERROR, NONE. The NONE status is returned when there is no quality gate associated with the analysis.<br />Returns an HTTP code 404 if the analysis associated with the task is not found or does not exist.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
func (s *QualitygatesService) Project_status(opt *QualitygatesProjectStatusOption) (Resp *QualitygatesProjectStatusResp, err error) {
	err := s.ValidateProject_statusOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualitygates/project_status", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Rename> func manually and complete the fields of this struct
type QualitygatesRenameResp struct{}
type QualitygatesRenameOption struct {
	Id           string `url:"id,omitempty"`           // Description:"ID of the quality gate to rename",ExampleValue:"1"
	Name         string `url:"name,omitempty"`         // Description:"New name of the quality gate",ExampleValue:"My Quality Gate"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Rename Rename a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Rename(opt *QualitygatesRenameOption) (Resp *QualitygatesRenameResp, err error) {
	err := s.ValidateRenameOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/rename", Opt)
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
type QualitygatesSearchResp struct{}
type QualitygatesSearchOption struct {
	GateId       string `url:"gateId,omitempty"`       // Description:"Quality Gate ID",ExampleValue:"1"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
	Page         string `url:"page,omitempty"`         // Description:"Page number",ExampleValue:"2"
	PageSize     string `url:"pageSize,omitempty"`     // Description:"Page size",ExampleValue:"10"
	Query        string `url:"query,omitempty"`        // Description:"To search for projects containing this string. If this parameter is set, "selected" is set to "all".",ExampleValue:"abc"
	Selected     string `url:"selected,omitempty"`     // Description:"Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).",ExampleValue:""
}

// Search Search for projects associated (or not) to a quality gate.<br/>Only authorized projects for current user will be returned.
func (s *QualitygatesService) Search(opt *QualitygatesSearchOption) (Resp *QualitygatesSearchResp, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualitygates/search", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Select> func manually and complete the fields of this struct
type QualitygatesSelectResp struct{}
type QualitygatesSelectOption struct {
	GateId       string `url:"gateId,omitempty"`       // Description:"Quality gate id",ExampleValue:"1"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
	ProjectKey   string `url:"projectKey,omitempty"`   // Description:"Project key",ExampleValue:"my_project"
}

// Select Associate a project to a quality gate.<br>The 'projectId' or 'projectKey' must be provided.<br>Project id as a numeric value is deprecated since 6.1. Please use the id similar to 'AU-TpxcA-iU5OvuD2FLz'.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Select(opt *QualitygatesSelectOption) (Resp *QualitygatesSelectResp, err error) {
	err := s.ValidateSelectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/select", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <SetAsDefault> func manually and complete the fields of this struct
type QualitygatesSetAsDefaultResp struct{}
type QualitygatesSetAsDefaultOption struct {
	Id           string `url:"id,omitempty"`           // Description:"ID of the quality gate to set as default",ExampleValue:"1"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Set_as_default Set a quality gate as the default quality gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Set_as_default(opt *QualitygatesSetAsDefaultOption) (Resp *QualitygatesSetAsDefaultResp, err error) {
	err := s.ValidateSet_as_defaultOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/set_as_default", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Show> func manually and complete the fields of this struct
type QualitygatesShowResp struct{}
type QualitygatesShowOption struct {
	Id           string `url:"id,omitempty"`           // Description:"ID of the quality gate. Either id or name must be set",ExampleValue:"1"
	Name         string `url:"name,omitempty"`         // Description:"Name of the quality gate. Either id or name must be set",ExampleValue:"My Quality Gate"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Show Display the details of a quality gate
func (s *QualitygatesService) Show(opt *QualitygatesShowOption) (Resp *QualitygatesShowResp, err error) {
	err := s.ValidateShowOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualitygates/show", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <UnsetDefault> func manually and complete the fields of this struct
type QualitygatesUnsetDefaultResp struct{}

// Unset_default This webservice is no more available : a default quality gate is mandatory.
func (s *QualitygatesService) Unset_default() (Resp *QualitygatesUnsetDefaultResp, err error) {
	req, err := s.client.NewRequest("POST", "qualitygates/unset_default", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <UpdateCondition> func manually and complete the fields of this struct
type QualitygatesUpdateConditionResp struct{}
type QualitygatesUpdateConditionOption struct {
	Error        string `url:"error,omitempty"`        // Description:"Condition error threshold",ExampleValue:"10"
	Id           string `url:"id,omitempty"`           // Description:"Condition ID",ExampleValue:"10"
	Metric       string `url:"metric,omitempty"`       // Description:"Condition metric",ExampleValue:"blocker_violations"
	Op           string `url:"op,omitempty"`           // Description:"Condition operator:<br/><ul><li>EQ = equals</li><li>NE = is not</li><li>LT = is lower than</li><li>GT = is greater than</li></ui>",ExampleValue:"EQ"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
	Period       string `url:"period,omitempty"`       // Description:"Condition period. If not set, the absolute value is considered.",ExampleValue:""
	Warning      string `url:"warning,omitempty"`      // Description:"Condition warning threshold",ExampleValue:"5"
}

// Update_condition Update a condition attached to a quality gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Update_condition(opt *QualitygatesUpdateConditionOption) (Resp *QualitygatesUpdateConditionResp, err error) {
	err := s.ValidateUpdate_conditionOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/update_condition", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
