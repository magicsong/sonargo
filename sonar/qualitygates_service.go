// Manage quality gates, including conditions and project association.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type QualitygatesService struct {
	client *Client
}

type Qualitygates struct {
	Actions struct {
		AssociateProjects bool `json:"associateProjects,omitempty"`
		Copy              bool `json:"copy,omitempty"`
		Create            bool `json:"create,omitempty"`
		Delete            bool `json:"delete,omitempty"`
		ManageConditions  bool `json:"manageConditions,omitempty"`
		Rename            bool `json:"rename,omitempty"`
		SetAsDefault      bool `json:"setAsDefault,omitempty"`
	} `json:"actions,omitempty"`
	Conditions []struct {
		Error   string `json:"error,omitempty"`
		ID      int64  `json:"id,omitempty"`
		Metric  string `json:"metric,omitempty"`
		Op      string `json:"op,omitempty"`
		Period  int64  `json:"period,omitempty"`
		Warning string `json:"warning,omitempty"`
	} `json:"conditions,omitempty"`
	Default int64  `json:"default,omitempty"`
	Error   string `json:"error,omitempty"`
	Errors  []struct {
		Msg string `json:"msg,omitempty"`
	} `json:"errors,omitempty"`
	ID            int64  `json:"id,omitempty"`
	IsBuiltIn     bool   `json:"isBuiltIn,omitempty"`
	Metric        string `json:"metric,omitempty"`
	More          bool   `json:"more,omitempty"`
	Name          string `json:"name,omitempty"`
	Op            string `json:"op,omitempty"`
	ProjectStatus struct {
		Conditions []struct {
			ActualValue      string `json:"actualValue,omitempty"`
			Comparator       string `json:"comparator,omitempty"`
			ErrorThreshold   string `json:"errorThreshold,omitempty"`
			MetricKey        string `json:"metricKey,omitempty"`
			PeriodIndex      int64  `json:"periodIndex,omitempty"`
			Status           string `json:"status,omitempty"`
			WarningThreshold string `json:"warningThreshold,omitempty"`
		} `json:"conditions,omitempty"`
		IgnoredConditions bool `json:"ignoredConditions,omitempty"`
		Periods           []struct {
			Date      string `json:"date,omitempty"`
			Index     int64  `json:"index,omitempty"`
			Mode      string `json:"mode,omitempty"`
			Parameter string `json:"parameter,omitempty"`
		} `json:"periods,omitempty"`
		Status string `json:"status,omitempty"`
	} `json:"projectStatus,omitempty"`
	QualityGate struct {
		Default bool   `json:"default,omitempty"`
		ID      string `json:"id,omitempty"`
		Name    string `json:"name,omitempty"`
	} `json:"qualityGate,omitempty"`
	Qualitygates []struct {
		Actions struct {
			AssociateProjects bool `json:"associateProjects,omitempty"`
			Copy              bool `json:"copy,omitempty"`
			Delete            bool `json:"delete,omitempty"`
			ManageConditions  bool `json:"manageConditions,omitempty"`
			Rename            bool `json:"rename,omitempty"`
			SetAsDefault      bool `json:"setAsDefault,omitempty"`
		} `json:"actions,omitempty"`
		ID        int64  `json:"id,omitempty"`
		IsBuiltIn bool   `json:"isBuiltIn,omitempty"`
		IsDefault bool   `json:"isDefault,omitempty"`
		Name      string `json:"name,omitempty"`
	} `json:"qualitygates,omitempty"`
	Results []struct {
		ID       int64  `json:"id,omitempty"`
		Name     string `json:"name,omitempty"`
		Selected bool   `json:"selected,omitempty"`
	} `json:"results,omitempty"`
	Warning string `json:"warning,omitempty"`
}

type QualitygatesCopyOption struct {
	Id           string `url:"id,omitempty"`           // Description:"The ID of the source quality gate",ExampleValue:"1"
	Name         string `url:"name,omitempty"`         // Description:"The name of the quality gate to create",ExampleValue:"My Quality Gate"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Copy Copy a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Copy(opt *QualitygatesCopyOption) (resp *string, err error) {
	err := s.ValidateCopyOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/copy", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesCreateOption struct {
	Name         string `url:"name,omitempty"`         // Description:"The name of the quality gate to create",ExampleValue:"My Quality Gate"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Create Create a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Create(opt *QualitygatesCreateOption) (resp *Qualitygates, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/create", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

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
func (s *QualitygatesService) CreateCondition(opt *QualitygatesCreateConditionOption) (resp *Qualitygates, err error) {
	err := s.ValidateCreateConditionOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/create_condition", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesDeleteConditionOption struct {
	Id           string `url:"id,omitempty"`           // Description:"Condition ID",ExampleValue:"2"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Delete_condition Delete a condition from a quality gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) DeleteCondition(opt *QualitygatesDeleteConditionOption) (resp *string, err error) {
	err := s.ValidateDeleteConditionOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/delete_condition", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesDeselectOption struct {
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
	ProjectId    string `url:"projectId,omitempty"`    // Description:"Project id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ProjectKey   string `url:"projectKey,omitempty"`   // Description:"Project key",ExampleValue:"my_project"
}

// Deselect Remove the association of a project from a quality gate.<br>Requires one of the following permissions:<ul><li>'Administer Quality Gates'</li><li>'Administer' rights on the project</li></ul>
func (s *QualitygatesService) Deselect(opt *QualitygatesDeselectOption) (resp *string, err error) {
	err := s.ValidateDeselectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/deselect", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesDestroyOption struct {
	Id           string `url:"id,omitempty"`           // Description:"ID of the quality gate to delete",ExampleValue:"1"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Destroy Delete a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Destroy(opt *QualitygatesDestroyOption) (resp *string, err error) {
	err := s.ValidateDestroyOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/destroy", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesGetByProjectOption struct {
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
	Project      string `url:"project,omitempty"`      // Description:"Project key",ExampleValue:"my_project"
}

// Get_by_project Get the quality gate of a project.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
func (s *QualitygatesService) GetByProject(opt *QualitygatesGetByProjectOption) (resp *Qualitygates, err error) {
	err := s.ValidateGetByProjectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualitygates/get_by_project", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesListOption struct {
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// List Get a list of quality gates
func (s *QualitygatesService) List(opt *QualitygatesListOption) (resp *Qualitygates, err error) {
	err := s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualitygates/list", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesProjectStatusOption struct {
	AnalysisId string `url:"analysisId,omitempty"` // Description:"Analysis id",ExampleValue:"AU-TpxcA-iU5OvuD2FL1"
	ProjectId  string `url:"projectId,omitempty"`  // Description:"Project id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ProjectKey string `url:"projectKey,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Project_status Get the quality gate status of a project or a Compute Engine task.<br />Either 'analysisId', 'projectId' or 'projectKey' must be provided<br />The different statuses returned are: OK, WARN, ERROR, NONE. The NONE status is returned when there is no quality gate associated with the analysis.<br />Returns an HTTP code 404 if the analysis associated with the task is not found or does not exist.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
func (s *QualitygatesService) ProjectStatus(opt *QualitygatesProjectStatusOption) (resp *Qualitygates, err error) {
	err := s.ValidateProjectStatusOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualitygates/project_status", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesRenameOption struct {
	Id           string `url:"id,omitempty"`           // Description:"ID of the quality gate to rename",ExampleValue:"1"
	Name         string `url:"name,omitempty"`         // Description:"New name of the quality gate",ExampleValue:"My Quality Gate"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Rename Rename a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Rename(opt *QualitygatesRenameOption) (resp *string, err error) {
	err := s.ValidateRenameOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/rename", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesSearchOption struct {
	GateId       string `url:"gateId,omitempty"`       // Description:"Quality Gate ID",ExampleValue:"1"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
	Page         string `url:"page,omitempty"`         // Description:"Page number",ExampleValue:"2"
	PageSize     string `url:"pageSize,omitempty"`     // Description:"Page size",ExampleValue:"10"
	Query        string `url:"query,omitempty"`        // Description:"To search for projects containing this string. If this parameter is set, "selected" is set to "all".",ExampleValue:"abc"
	Selected     string `url:"selected,omitempty"`     // Description:"Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).",ExampleValue:""
}

// Search Search for projects associated (or not) to a quality gate.<br/>Only authorized projects for current user will be returned.
func (s *QualitygatesService) Search(opt *QualitygatesSearchOption) (resp *Qualitygates, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualitygates/search", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesSelectOption struct {
	GateId       string `url:"gateId,omitempty"`       // Description:"Quality gate id",ExampleValue:"1"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
	ProjectKey   string `url:"projectKey,omitempty"`   // Description:"Project key",ExampleValue:"my_project"
}

// Select Associate a project to a quality gate.<br>The 'projectId' or 'projectKey' must be provided.<br>Project id as a numeric value is deprecated since 6.1. Please use the id similar to 'AU-TpxcA-iU5OvuD2FLz'.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) Select(opt *QualitygatesSelectOption) (resp *string, err error) {
	err := s.ValidateSelectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/select", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesSetAsDefaultOption struct {
	Id           string `url:"id,omitempty"`           // Description:"ID of the quality gate to set as default",ExampleValue:"1"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Set_as_default Set a quality gate as the default quality gate.<br>Requires the 'Administer Quality Gates' permission.
func (s *QualitygatesService) SetAsDefault(opt *QualitygatesSetAsDefaultOption) (resp *string, err error) {
	err := s.ValidateSetAsDefaultOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/set_as_default", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualitygatesShowOption struct {
	Id           string `url:"id,omitempty"`           // Description:"ID of the quality gate. Either id or name must be set",ExampleValue:"1"
	Name         string `url:"name,omitempty"`         // Description:"Name of the quality gate. Either id or name must be set",ExampleValue:"My Quality Gate"
	Organization string `url:"organization,omitempty"` // Description:"Organization key. If no organization is provided, the default organization is used.",ExampleValue:"my-org"
}

// Show Display the details of a quality gate
func (s *QualitygatesService) Show(opt *QualitygatesShowOption) (resp *Qualitygates, err error) {
	err := s.ValidateShowOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualitygates/show", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Unset_default This webservice is no more available : a default quality gate is mandatory.
func (s *QualitygatesService) UnsetDefault() (resp *Qualitygates, err error) {
	req, err := s.client.NewRequest("POST", "qualitygates/unset_default", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

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
func (s *QualitygatesService) UpdateCondition(opt *QualitygatesUpdateConditionOption) (resp *string, err error) {
	err := s.ValidateUpdateConditionOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualitygates/update_condition", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
