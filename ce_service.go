// Get information on Compute Engine tasks.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type CeService struct {
	client *Client
}

// [TODO] you should call the <Activity> func manually and complete the fields of this struct
type CeActivityResp struct{}
type CeActivityOption struct {
	ComponentId    string `url:"componentId,omitempty"`    // Description:"Id of the component (project) to filter on",ExampleValue:"AU-TpxcA-iU5OvuD2FL0"
	MaxExecutedAt  string `url:"maxExecutedAt,omitempty"`  // Description:"Maximum date of end of task processing (inclusive)",ExampleValue:"2017-10-19T13:00:00+0200"
	MinSubmittedAt string `url:"minSubmittedAt,omitempty"` // Description:"Minimum date of task submission (inclusive)",ExampleValue:"2017-10-19T13:00:00+0200"
	OnlyCurrents   string `url:"onlyCurrents,omitempty"`   // Description:"Filter on the last tasks (only the most recent finished task by project)",ExampleValue:""
	Ps             string `url:"ps,omitempty"`             // Description:"Page size. Must be greater than 0 and less or equal than 1000",ExampleValue:"20"
	Q              string `url:"q,omitempty"`              // Description:"Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li><li>task ids that are exactly the same as the supplied string</li></ul>Must not be set together with componentId",ExampleValue:"Apache"
	Status         string `url:"status,omitempty"`         // Description:"Comma separated list of task statuses",ExampleValue:"IN_PROGRESS,SUCCESS"
	Type           string `url:"type,omitempty"`           // Description:"Task type",ExampleValue:"REPORT"
}

// Activity Search for tasks.<br> Requires the system administration permission, or project administration permission if componentId is set.
func (s *CeService) Activity(opt *CeActivityOption) (Resp *CeActivityResp, err error) {
	err := s.ValidateActivityOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "ce/activity", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Component> func manually and complete the fields of this struct
type CeComponentResp struct{}
type CeComponentOption struct {
	Component   string `url:"component,omitempty"`   // Description:"",ExampleValue:"my_project"
	ComponentId string `url:"componentId,omitempty"` // Description:"",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
}

// Component Get the pending tasks, in-progress tasks and the last executed task of a given component (usually a project).<br>Requires the following permission: 'Browse' on the specified component.<br>Either 'componentId' or 'component' must be provided.
func (s *CeService) Component(opt *CeComponentOption) (Resp *CeComponentResp, err error) {
	err := s.ValidateComponentOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "ce/component", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Task> func manually and complete the fields of this struct
type CeTaskResp struct{}
type CeTaskOption struct {
	AdditionalFields string `url:"additionalFields,omitempty"` // Description:"Comma-separated list of the optional fields to be returned in response.",ExampleValue:""
	Id               string `url:"id,omitempty"`               // Description:"Id of task",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
}

// Task Give Compute Engine task details such as type, status, duration and associated component.<br />Requires 'Administer System' or 'Execute Analysis' permission.<br/>Since 6.1, field "logs" is deprecated and its value is always false.
func (s *CeService) Task(opt *CeTaskOption) (Resp *CeTaskResp, err error) {
	err := s.ValidateTaskOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "ce/task", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
