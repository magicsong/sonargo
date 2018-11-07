// Manage notifications of the authenticated user
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type NotificationsService struct {
	client *Client
}

// [TODO] you should call the <Add> func manually and complete the fields of this struct
type NotificationsAddResp struct{}
type NotificationsAddOption struct {
	Channel string `url:"channel,omitempty"` // Description:"Channel through which the notification is sent. For example, notifications can be sent by email.",ExampleValue:""
	Login   string `url:"login,omitempty"`   // Description:"User login",ExampleValue:""
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Type    string `url:"type,omitempty"`    // Description:"Notification type. Possible values are for:<ul>  <li>Global notifications: CeReportTaskFailure, ChangesOnMyIssue, NewAlerts, NewFalsePositiveIssue, NewIssues, SQ-MyNewIssues</li>  <li>Per project notifications: CeReportTaskFailure, ChangesOnMyIssue, NewAlerts, NewFalsePositiveIssue, NewIssues, SQ-MyNewIssues</li></ul>",ExampleValue:"SQ-MyNewIssues"
}

// Add Add a notification for the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided. If a project is provided, requires the 'Browse' permission on the specified project.</li> <li>System administration if a login is provided. If a project is provided, requires the 'Browse' permission on the specified project.</li></ul>
func (s *NotificationsService) Add(opt *NotificationsAddOption) (Resp *NotificationsAddResp, err error) {
	err := s.ValidateAddOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "notifications/add", Opt)
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
type NotificationsListResp struct{}
type NotificationsListOption struct {
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:""
}

// List List notifications of the authenticated user.<br>Requires one of the following permissions:<ul>  <li>Authentication if no login is provided</li>  <li>System administration if a login is provided</li></ul>
func (s *NotificationsService) List(opt *NotificationsListOption) (Resp *NotificationsListResp, err error) {
	err := s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "notifications/list", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Remove> func manually and complete the fields of this struct
type NotificationsRemoveResp struct{}
type NotificationsRemoveOption struct {
	Channel string `url:"channel,omitempty"` // Description:"Channel through which the notification is sent. For example, notifications can be sent by email.",ExampleValue:""
	Login   string `url:"login,omitempty"`   // Description:"User login",ExampleValue:""
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Type    string `url:"type,omitempty"`    // Description:"Notification type. Possible values are for:<ul>  <li>Global notifications: CeReportTaskFailure, ChangesOnMyIssue, NewAlerts, NewFalsePositiveIssue, NewIssues, SQ-MyNewIssues</li>  <li>Per project notifications: CeReportTaskFailure, ChangesOnMyIssue, NewAlerts, NewFalsePositiveIssue, NewIssues, SQ-MyNewIssues</li></ul>",ExampleValue:"SQ-MyNewIssues"
}

// Remove Remove a notification for the authenticated user.<br>Requires one of the following permissions:<ul>  <li>Authentication if no login is provided</li>  <li>System administration if a login is provided</li></ul>
func (s *NotificationsService) Remove(opt *NotificationsRemoveOption) (Resp *NotificationsRemoveResp, err error) {
	err := s.ValidateRemoveOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "notifications/remove", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
