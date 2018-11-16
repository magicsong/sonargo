package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	spew "github.com/davecgh/go-spew/spew"
	. "github.com/magicsong/sonargo/sonar"
)

var client *Client

func init() {
	sonarURL := os.Getenv("SONAR_URL")
	if sonarURL == "" {
		fmt.Println("Sonar URL has not been set")
		os.Exit(1)
	}
	c, err := NewClient(sonarURL+"/api", "admin", "admin")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	client = c

}

// You should MANUALLY add the test func in here
func main() {
	// RulesCreateFunc()
	// PermissionsSearchTemplatesFunc()
	//RulesDeleteFunc()
	// RulesUpdateFunc()
	QualityProfilesBackupFunc()
	return
}

// CeActivityFunc testing Search for tasks.<br> Requires the system administration permission, or project administration permission if componentId is set.
func CeActivityFunc() {
	opt := &CeActivityOption{
		ComponentId:    "",
		MaxExecutedAt:  "",
		MinSubmittedAt: "",
		OnlyCurrents:   "",
		Ps:             "",
		Q:              "",
		Status:         "",
		Type:           "",
	}
	v, resp, err := client.Ce.Activity(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// CeComponentFunc testing Get the pending tasks, in-progress tasks and the last executed task of a given component (usually a project).<br>Requires the following permission: 'Browse' on the specified component.<br>Either 'componentId' or 'component' must be provided.
func CeComponentFunc() {
	opt := &CeComponentOption{
		Component:   "",
		ComponentId: "",
	}
	v, resp, err := client.Ce.Component(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// CeTaskFunc testing Give Compute Engine task details such as type, status, duration and associated component.<br />Requires 'Administer System' or 'Execute Analysis' permission.<br/>Since 6.1, field "logs" is deprecated and its value is always false.
func CeTaskFunc() {
	opt := &CeTaskOption{
		AdditionalFields: "",
		Id:               "MUST_EDIT_IT",
	}
	v, resp, err := client.Ce.Task(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ComponentsSearchFunc testing Search for components
func ComponentsSearchFunc() {
	opt := &ComponentsSearchOption{
		Language:   "",
		P:          "",
		Ps:         "",
		Q:          "",
		Qualifiers: "MUST_EDIT_IT",
	}
	v, resp, err := client.Components.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ComponentsShowFunc testing Returns a component (file, directory, project, view…) and its ancestors. The ancestors are ordered from the parent to the root project. The 'componentId' or 'component' parameter must be provided.<br>Requires the following permission: 'Browse' on the project of the specified component.
func ComponentsShowFunc() {
	opt := &ComponentsShowOption{
		Component:   "",
		ComponentId: "",
	}
	v, resp, err := client.Components.Show(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ComponentsTreeFunc testing Navigate through components based on the chosen strategy. The componentId or the component parameter must be provided.<br>Requires the following permission: 'Browse' on the specified project.<br>When limiting search with the q parameter, directories are not returned.
func ComponentsTreeFunc() {
	opt := &ComponentsTreeOption{
		Asc:         "",
		Component:   "",
		ComponentId: "",
		P:           "",
		Ps:          "",
		Q:           "",
		Qualifiers:  "",
		S:           "",
		Strategy:    "",
	}
	v, resp, err := client.Components.Tree(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// DuplicationsShowFunc testing Get duplications. Require Browse permission on file's project
func DuplicationsShowFunc() {
	opt := &DuplicationsShowOption{
		Key:  "",
		Uuid: "",
	}
	v, resp, err := client.Duplications.Show(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// FavoritesAddFunc testing Add a component (project, directory, file etc.) as favorite for the authenticated user.<br>Requires authentication and the following permission: 'Browse' on the project of the specified component.
func FavoritesAddFunc() {
	opt := &FavoritesAddOption{Component: "MUST_EDIT_IT"}
	resp, err := client.Favorites.Add(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// FavoritesRemoveFunc testing Remove a component (project, directory, file etc.) as favorite for the authenticated user.<br>Requires authentication.
func FavoritesRemoveFunc() {
	opt := &FavoritesRemoveOption{Component: "MUST_EDIT_IT"}
	resp, err := client.Favorites.Remove(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// FavoritesSearchFunc testing Search for the authenticated user favorites.<br>Requires authentication.
func FavoritesSearchFunc() {
	opt := &FavoritesSearchOption{
		P:  0,
		Ps: 0,
	}
	v, resp, err := client.Favorites.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesAddCommentFunc testing Add a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
func IssuesAddCommentFunc() {
	opt := &IssuesAddCommentOption{
		Issue: "AWcai7lPiJioIHvADp15",
		Text:  "Write a comment",
	}
	v, resp, err := client.Issues.AddComment(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesAssignFunc testing Assign/Unassign an issue. Requires authentication and Browse permission on project
func IssuesAssignFunc() {
	opt := &IssuesAssignOption{
		Assignee: "",
		Issue:    "MUST_EDIT_IT",
	}
	v, resp, err := client.Issues.Assign(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesAuthorsFunc testing Search SCM accounts which match a given query
func IssuesAuthorsFunc() {
	opt := &IssuesAuthorsOption{
		Ps: 0,
		Q:  "",
	}
	v, resp, err := client.Issues.Authors(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesBulkChangeFunc testing Bulk change on issues.<br/>Requires authentication.
func IssuesBulkChangeFunc() {
	opt := &IssuesBulkChangeOption{
		AddTags:           "",
		Assign:            "",
		Comment:           "",
		DoTransition:      "",
		Issues:            "MUST_EDIT_IT",
		Plan:              "",
		RemoveTags:        "",
		SendNotifications: "",
		SetSeverity:       "",
		SetType:           "",
	}
	v, resp, err := client.Issues.BulkChange(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesChangelogFunc testing Display changelog of an issue.<br/>Requires the 'Browse' permission on the project of the specified issue.
func IssuesChangelogFunc() {
	opt := &IssuesChangelogOption{Issue: "MUST_EDIT_IT"}
	v, resp, err := client.Issues.Changelog(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesDeleteCommentFunc testing Delete a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
func IssuesDeleteCommentFunc() {
	opt := &IssuesDeleteCommentOption{Comment: "MUST_EDIT_IT"}
	v, resp, err := client.Issues.DeleteComment(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesDoTransitionFunc testing Do workflow transition on an issue. Requires authentication and Browse permission on project.<br/>The transitions 'wontfix' and 'falsepositive' require the permission 'Administer Issues'.
func IssuesDoTransitionFunc() {
	opt := &IssuesDoTransitionOption{
		Issue:      "AWcai7gLiJioIHvADp0q",
		Transition: "close",
	}
	v, resp, err := client.Issues.DoTransition(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesEditCommentFunc testing Edit a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
func IssuesEditCommentFunc() {
	opt := &IssuesEditCommentOption{
		Comment: "MUST_EDIT_IT",
		Text:    "MUST_EDIT_IT",
	}
	v, resp, err := client.Issues.EditComment(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesSearchFunc testing Search for issues.<br>At most one of the following parameters can be provided at the same time: componentKeys, componentUuids, components, componentRootUuids, componentRoots.<br>Requires the 'Browse' permission on the specified project(s).
func IssuesSearchFunc() {
	opt := &IssuesSearchOption{
		AdditionalFields:   "",
		Asc:                "",
		Assigned:           "",
		Assignees:          "",
		Authors:            "",
		ComponentKeys:      "",
		ComponentRootUuids: "",
		ComponentRoots:     "",
		ComponentUuids:     "",
		Components:         "",
		CreatedAfter:       "",
		CreatedAt:          "",
		CreatedBefore:      "",
		CreatedInLast:      "",
		Issues:             "",
		Languages:          "",
		P:                  "",
		Ps:                 "",
		Resolutions:        "",
		Resolved:           "",
		Rules:              "",
		S:                  "",
		Severities:         "MINOR",
		SinceLeakPeriod:    "",
		Statuses:           "",
		Tags:               "",
		Types:              "VULNERABILITY",
	}
	v, resp, err := client.Issues.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v.Issues[0])
}

// IssuesSetSeverityFunc testing Change severity.<br/>Requires the following permissions:<ul>  <li>'Authentication'</li>  <li>'Browse' rights on project of the specified issue</li>  <li>'Administer Issues' rights on project of the specified issue</li></ul>
func IssuesSetSeverityFunc() {
	opt := &IssuesSetSeverityOption{
		Issue:    "MUST_EDIT_IT",
		Severity: "MUST_EDIT_IT",
	}
	v, resp, err := client.Issues.SetSeverity(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesSetTagsFunc testing Set tags on an issue. <br/>Requires authentication and Browse permission on project
func IssuesSetTagsFunc() {
	opt := &IssuesSetTagsOption{
		Issue: "MUST_EDIT_IT",
		Tags:  "",
	}
	v, resp, err := client.Issues.SetTags(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesSetTypeFunc testing Change type of issue, for instance from 'code smell' to 'bug'.<br/>Requires the following permissions:<ul>  <li>'Authentication'</li>  <li>'Browse' rights on project of the specified issue</li>  <li>'Administer Issues' rights on project of the specified issue</li></ul>
func IssuesSetTypeFunc() {
	opt := &IssuesSetTypeOption{
		Issue: "MUST_EDIT_IT",
		Type:  "MUST_EDIT_IT",
	}
	v, resp, err := client.Issues.SetType(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// IssuesTagsFunc testing List tags matching a given query
func IssuesTagsFunc() {
	opt := &IssuesTagsOption{
		Ps: 0,
		Q:  "",
	}
	v, resp, err := client.Issues.Tags(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// LanguagesListFunc testing List supported programming languages
func LanguagesListFunc() {
	opt := &LanguagesListOption{
		Ps: 0,
		Q:  "",
	}
	v, resp, err := client.Languages.List(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// MeasuresComponentFunc testing Return component with specified measures. The componentId or the component parameter must be provided.<br>Requires the following permission: 'Browse' on the project of specified component.
func MeasuresComponentFunc() {
	opt := &MeasuresComponentOption{
		AdditionalFields: "",
		Component:        "",
		ComponentId:      "AWcai51MgFEHqqIBJRwT",
		MetricKeys:       "ncloc,complexity,violations",
	}
	v, resp, err := client.Measures.Component(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// MeasuresComponentTreeFunc testing Navigate through components based on the chosen strategy with specified measures. The baseComponentId or the component parameter must be provided.<br>Requires the following permission: 'Browse' on the specified project.<br>When limiting search with the q parameter, directories are not returned.
func MeasuresComponentTreeFunc() {
	opt := &MeasuresComponentTreeOption{
		AdditionalFields: "",
		Asc:              "",
		BaseComponentId:  "",
		Component:        "",
		MetricKeys:       "MUST_EDIT_IT",
		MetricPeriodSort: "",
		MetricSort:       "",
		MetricSortFilter: "",
		P:                "",
		Ps:               "",
		Q:                "",
		Qualifiers:       "",
		S:                "",
		Strategy:         "",
	}
	v, resp, err := client.Measures.ComponentTree(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// MeasuresSearchHistoryFunc testing Search measures history of a component.<br>Measures are ordered chronologically.<br>Pagination applies to the number of measures for each metric.<br>Requires the following permission: 'Browse' on the specified component
func MeasuresSearchHistoryFunc() {
	opt := &MeasuresSearchHistoryOption{
		Component: "MUST_EDIT_IT",
		From:      "",
		Metrics:   "MUST_EDIT_IT",
		P:         "",
		Ps:        "",
		To:        "",
	}
	v, resp, err := client.Measures.SearchHistory(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// MetricsCreateFunc testing Create custom metric.<br /> Requires 'Administer System' permission.
func MetricsCreateFunc() {
	opt := &MetricsCreateOption{
		Description: "Test custom metric",
		Domain:      "Tests",
		Key:         "team_size5",
		Name:        "Team Size5",
		Type:        MetricTypeINT,
	}
	_, resp, err := client.Metrics.Create(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	str, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Print(string(str))
}

// MetricsDeleteFunc testing Delete metrics and associated measures. Delete only custom metrics.<br />Ids or keys must be provided. <br />Requires 'Administer System' permission.
func MetricsDeleteFunc() {
	opt := &MetricsDeleteOption{
		Ids:  "",
		Keys: "",
	}
	resp, err := client.Metrics.Delete(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// MetricsDomainsFunc testing List all custom metric domains.
func MetricsDomainsFunc() {
	v, resp, err := client.Metrics.Domains()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// MetricsSearchFunc testing Search for metrics
func MetricsSearchFunc() {
	opt := &MetricsSearchOption{
		F:        "",
		IsCustom: "",
		P:        0,
		Ps:       0,
	}
	v, resp, err := client.Metrics.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// MetricsTypesFunc testing List all available metric types.
func MetricsTypesFunc() {
	v, resp, err := client.Metrics.Types()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// MetricsUpdateFunc testing Update a custom metric.<br /> Requires 'Administer System' permission.
func MetricsUpdateFunc() {
	testKey := "hello"
	defer func() {
		optDelete := &MetricsDeleteOption{
			Ids:  "",
			Keys: testKey + "0",
		}
		client.Metrics.Delete(optDelete)
	}()
	optCreate := &MetricsCreateOption{
		Description: "Test custom metric for update",
		Domain:      "Tests",
		Key:         testKey + "0",
		Name:        "Team Size1",
		Type:        MetricTypeINT,
	}
	v, _, err := client.Metrics.Create(optCreate)
	opt := &MetricsUpdateOption{
		Description: "Hello World",
		Domain:      "",
		ID:          v.ID,
		Key:         "",
		Name:        "",
		Type:        "",
	}
	_, resp, err := client.Metrics.Update(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	str, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(str))
}

// NotificationsAddFunc testing Add a notification for the authenticated user.<br>Requires one of the following permissions:<ul> <li>Authentication if no login is provided. If a project is provided, requires the 'Browse' permission on the specified project.</li> <li>System administration if a login is provided. If a project is provided, requires the 'Browse' permission on the specified project.</li></ul>
func NotificationsAddFunc() {
	opt := &NotificationsAddOption{
		Channel: "",
		Login:   "",
		Project: "",
		Type:    "MUST_EDIT_IT",
	}
	resp, err := client.Notifications.Add(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// NotificationsListFunc testing List notifications of the authenticated user.<br>Requires one of the following permissions:<ul>  <li>Authentication if no login is provided</li>  <li>System administration if a login is provided</li></ul>
func NotificationsListFunc() {
	opt := &NotificationsListOption{Login: ""}
	v, resp, err := client.Notifications.List(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// NotificationsRemoveFunc testing Remove a notification for the authenticated user.<br>Requires one of the following permissions:<ul>  <li>Authentication if no login is provided</li>  <li>System administration if a login is provided</li></ul>
func NotificationsRemoveFunc() {
	opt := &NotificationsRemoveOption{
		Channel: "",
		Login:   "",
		Project: "",
		Type:    "MUST_EDIT_IT",
	}
	resp, err := client.Notifications.Remove(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsAddGroupFunc testing Add permission to a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group name or group id must be provided. <br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func PermissionsAddGroupFunc() {
	opt := &PermissionsAddGroupOption{
		GroupId:    0,
		GroupName:  "",
		Permission: "MUST_EDIT_IT",
		ProjectId:  "",
		ProjectKey: "",
	}
	resp, err := client.Permissions.AddGroup(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsAddGroupToTemplateFunc testing Add a group to a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'.
func PermissionsAddGroupToTemplateFunc() {
	opt := &PermissionsAddGroupToTemplateOption{
		GroupId:      0,
		GroupName:    "",
		Permission:   "MUST_EDIT_IT",
		TemplateId:   "",
		TemplateName: "",
	}
	resp, err := client.Permissions.AddGroupToTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsAddProjectCreatorToTemplateFunc testing Add a project creator to a permission template.<br>Requires the following permission: 'Administer System'.
func PermissionsAddProjectCreatorToTemplateFunc() {
	opt := &PermissionsAddProjectCreatorToTemplateOption{
		Permission:   "MUST_EDIT_IT",
		TemplateId:   "",
		TemplateName: "",
	}
	resp, err := client.Permissions.AddProjectCreatorToTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsAddUserFunc testing Add permission to a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func PermissionsAddUserFunc() {
	opt := &PermissionsAddUserOption{
		Login:      "MUST_EDIT_IT",
		Permission: "MUST_EDIT_IT",
		ProjectId:  "",
		ProjectKey: "",
	}
	resp, err := client.Permissions.AddUser(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsAddUserToTemplateFunc testing Add a user to a permission template.<br /> Requires the following permission: 'Administer System'.
func PermissionsAddUserToTemplateFunc() {
	opt := &PermissionsAddUserToTemplateOption{
		Login:        "MUST_EDIT_IT",
		Permission:   "MUST_EDIT_IT",
		TemplateId:   "",
		TemplateName: "",
	}
	resp, err := client.Permissions.AddUserToTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsApplyTemplateFunc testing Apply a permission template to one project.<br>The project id or project key must be provided.<br>The template id or name must be provided.<br>Requires the following permission: 'Administer System'.
func PermissionsApplyTemplateFunc() {
	opt := &PermissionsApplyTemplateOption{
		ProjectId:    "",
		ProjectKey:   "",
		TemplateId:   "",
		TemplateName: "",
	}
	resp, err := client.Permissions.ApplyTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsBulkApplyTemplateFunc testing Apply a permission template to several projects.<br />The template id or name must be provided.<br />Requires the following permission: 'Administer System'.
func PermissionsBulkApplyTemplateFunc() {
	opt := &PermissionsBulkApplyTemplateOption{
		AnalyzedBefore:    "",
		OnProvisionedOnly: "",
		Projects:          "",
		Q:                 "",
		Qualifiers:        "",
		TemplateId:        "",
		TemplateName:      "",
	}
	resp, err := client.Permissions.BulkApplyTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsCreateTemplateFunc testing Create a permission template.<br />Requires the following permission: 'Administer System'.
func PermissionsCreateTemplateFunc() {
	opt := &PermissionsCreateTemplateOption{
		Description:       "",
		Name:              "MUST_EDIT_IT",
		ProjectKeyPattern: "",
	}
	v, resp, err := client.Permissions.CreateTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// PermissionsDeleteTemplateFunc testing Delete a permission template.<br />Requires the following permission: 'Administer System'.
func PermissionsDeleteTemplateFunc() {
	opt := &PermissionsDeleteTemplateOption{
		TemplateId:   "",
		TemplateName: "",
	}
	resp, err := client.Permissions.DeleteTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsRemoveGroupFunc testing Remove a permission from a group.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> The group id or group name must be provided, not both.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func PermissionsRemoveGroupFunc() {
	opt := &PermissionsRemoveGroupOption{
		GroupId:    0,
		GroupName:  "",
		Permission: "MUST_EDIT_IT",
		ProjectId:  "",
		ProjectKey: "",
	}
	resp, err := client.Permissions.RemoveGroup(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsRemoveGroupFromTemplateFunc testing Remove a group from a permission template.<br /> The group id or group name must be provided. <br />Requires the following permission: 'Administer System'.
func PermissionsRemoveGroupFromTemplateFunc() {
	opt := &PermissionsRemoveGroupFromTemplateOption{
		GroupId:      0,
		GroupName:    "",
		Permission:   "MUST_EDIT_IT",
		TemplateId:   "",
		TemplateName: "",
	}
	resp, err := client.Permissions.RemoveGroupFromTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsRemoveProjectCreatorFromTemplateFunc testing Remove a project creator from a permission template.<br>Requires the following permission: 'Administer System'.
func PermissionsRemoveProjectCreatorFromTemplateFunc() {
	opt := &PermissionsRemoveProjectCreatorFromTemplateOption{
		Permission:   "MUST_EDIT_IT",
		TemplateId:   "",
		TemplateName: "",
	}
	resp, err := client.Permissions.RemoveProjectCreatorFromTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsRemoveUserFunc testing Remove permission from a user.<br /> This service defaults to global permissions, but can be limited to project permissions by providing project id or project key.<br /> Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func PermissionsRemoveUserFunc() {
	opt := &PermissionsRemoveUserOption{
		Login:      "MUST_EDIT_IT",
		Permission: "MUST_EDIT_IT",
		ProjectId:  "",
		ProjectKey: "",
	}
	resp, err := client.Permissions.RemoveUser(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsRemoveUserFromTemplateFunc testing Remove a user from a permission template.<br /> Requires the following permission: 'Administer System'.
func PermissionsRemoveUserFromTemplateFunc() {
	opt := &PermissionsRemoveUserFromTemplateOption{
		Login:        "MUST_EDIT_IT",
		Permission:   "MUST_EDIT_IT",
		TemplateId:   "",
		TemplateName: "",
	}
	resp, err := client.Permissions.RemoveUserFromTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsSearchTemplatesFunc testing List permission templates.<br />Requires the following permission: 'Administer System'.
func PermissionsSearchTemplatesFunc() {
	opt := &PermissionsSearchTemplatesOption{Q: ""}
	v, resp, err := client.Permissions.SearchTemplates(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// PermissionsSetDefaultTemplateFunc testing Set a permission template as default.<br />Requires the following permission: 'Administer System'.
func PermissionsSetDefaultTemplateFunc() {
	opt := &PermissionsSetDefaultTemplateOption{
		Qualifier:    "",
		TemplateId:   "",
		TemplateName: "",
	}
	resp, err := client.Permissions.SetDefaultTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PermissionsUpdateTemplateFunc testing Update a permission template.<br />Requires the following permission: 'Administer System'.
func PermissionsUpdateTemplateFunc() {
	opt := &PermissionsUpdateTemplateOption{
		Description:       "",
		Id:                "MUST_EDIT_IT",
		Name:              "",
		ProjectKeyPattern: "",
	}
	v, resp, err := client.Permissions.UpdateTemplate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// PluginsAvailableFunc testing Get the list of all the plugins available for installation on the SonarQube instance, sorted by plugin name.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.<br/>Update status values are: <ul><li>COMPATIBLE: plugin is compatible with current SonarQube instance.</li><li>INCOMPATIBLE: plugin is not compatible with current SonarQube instance.</li><li>REQUIRES_SYSTEM_UPGRADE: plugin requires SonarQube to be upgraded before being installed.</li><li>DEPS_REQUIRE_SYSTEM_UPGRADE: at least one plugin on which the plugin is dependent requires SonarQube to be upgraded.</li></ul>Require 'Administer System' permission.
func PluginsAvailableFunc() {
	v, resp, err := client.Plugins.Available()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// PluginsCancelAllFunc testing Cancels any operation pending on any plugin (install, update or uninstall)<br/>Requires user to be authenticated with Administer System permissions
func PluginsCancelAllFunc() {
	resp, err := client.Plugins.CancelAll()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PluginsInstallFunc testing Installs the latest version of a plugin specified by its key.<br/>Plugin information is retrieved from Update Center.<br/>Requires user to be authenticated with Administer System permissions
func PluginsInstallFunc() {
	opt := &PluginsInstallOption{Key: "MUST_EDIT_IT"}
	resp, err := client.Plugins.Install(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PluginsInstalledFunc testing Get the list of all the plugins installed on the SonarQube instance, sorted by plugin name.
func PluginsInstalledFunc() {
	opt := &PluginsInstalledOption{F: ""}
	v, resp, err := client.Plugins.Installed(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// PluginsPendingFunc testing Get the list of plugins which will either be installed or removed at the next startup of the SonarQube instance, sorted by plugin name.<br/>Require 'Administer System' permission.
func PluginsPendingFunc() {
	v, resp, err := client.Plugins.Pending()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// PluginsUninstallFunc testing Uninstalls the plugin specified by its key.<br/>Requires user to be authenticated with Administer System permissions.
func PluginsUninstallFunc() {
	opt := &PluginsUninstallOption{Key: "MUST_EDIT_IT"}
	resp, err := client.Plugins.Uninstall(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PluginsUpdateFunc testing Updates a plugin specified by its key to the latest version compatible with the SonarQube instance.<br/>Plugin information is retrieved from Update Center.<br/>Requires user to be authenticated with Administer System permissions
func PluginsUpdateFunc() {
	opt := &PluginsUpdateOption{Key: "MUST_EDIT_IT"}
	resp, err := client.Plugins.Update(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// PluginsUpdatesFunc testing Lists plugins installed on the SonarQube instance for which at least one newer version is available, sorted by plugin name.<br/>Each newer version is listed, ordered from the oldest to the newest, with its own update/compatibility status.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.<br/>Update status values are: [COMPATIBLE, INCOMPATIBLE, REQUIRES_UPGRADE, DEPS_REQUIRE_UPGRADE].<br/>Require 'Administer System' permission.
func PluginsUpdatesFunc() {
	v, resp, err := client.Plugins.Updates()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectAnalysesCreateEventFunc testing Create a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be created.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
func ProjectAnalysesCreateEventFunc() {
	opt := &ProjectAnalysesCreateEventOption{
		Analysis: "MUST_EDIT_IT",
		Category: "",
		Name:     "MUST_EDIT_IT",
	}
	v, resp, err := client.ProjectAnalyses.CreateEvent(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectAnalysesDeleteFunc testing Delete a project analysis.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the project of the specified analysis</li></ul>
func ProjectAnalysesDeleteFunc() {
	opt := &ProjectAnalysesDeleteOption{Analysis: "MUST_EDIT_IT"}
	resp, err := client.ProjectAnalyses.Delete(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// ProjectAnalysesDeleteEventFunc testing Delete a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be deleted.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
func ProjectAnalysesDeleteEventFunc() {
	opt := &ProjectAnalysesDeleteEventOption{Event: "MUST_EDIT_IT"}
	resp, err := client.ProjectAnalyses.DeleteEvent(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// ProjectAnalysesSearchFunc testing Search a project analyses and attached events.<br>Requires the following permission: 'Browse' on the specified project
func ProjectAnalysesSearchFunc() {
	opt := &ProjectAnalysesSearchOption{
		Category: "",
		From:     "",
		P:        0,
		Project:  "MUST_EDIT_IT",
		Ps:       0,
		To:       "",
	}
	v, resp, err := client.ProjectAnalyses.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectAnalysesUpdateEventFunc testing Update a project analysis event.<br>Only events of category 'VERSION' and 'OTHER' can be updated.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
func ProjectAnalysesUpdateEventFunc() {
	opt := &ProjectAnalysesUpdateEventOption{
		Event: "MUST_EDIT_IT",
		Name:  "",
	}
	v, resp, err := client.ProjectAnalyses.UpdateEvent(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectBadgesMeasureFunc testing Generate badge for project's measure as an SVG.<br/>Requires 'Browse' permission on the specified project.
func ProjectBadgesMeasureFunc() {
	opt := &ProjectBadgesMeasureOption{
		Branch:      "",
		Metric:      "MUST_EDIT_IT",
		Project:     "MUST_EDIT_IT",
		PullRequest: "",
	}
	v, resp, err := client.ProjectBadges.Measure(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectBadgesQualityGateFunc testing Generate badge for project's quality gate as an SVG.<br/>Requires 'Browse' permission on the specified project.
func ProjectBadgesQualityGateFunc() {
	opt := &ProjectBadgesQualityGateOption{
		Branch:      "",
		Project:     "MUST_EDIT_IT",
		PullRequest: "",
	}
	v, resp, err := client.ProjectBadges.QualityGate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectBranchesDeleteFunc testing Delete a non-main branch of a project.<br/>Requires 'Administer' rights on the specified project.
func ProjectBranchesDeleteFunc() {
	opt := &ProjectBranchesDeleteOption{
		Branch:  "MUST_EDIT_IT",
		Project: "MUST_EDIT_IT",
	}
	resp, err := client.ProjectBranches.Delete(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// ProjectBranchesListFunc testing List the branches of a project.<br/>Requires 'Browse' or 'Execute analysis' rights on the specified project.
func ProjectBranchesListFunc() {
	opt := &ProjectBranchesListOption{Project: "MUST_EDIT_IT"}
	v, resp, err := client.ProjectBranches.List(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectBranchesRenameFunc testing Rename the main branch of a project.<br/>Requires 'Administer' permission on the specified project.
func ProjectBranchesRenameFunc() {
	opt := &ProjectBranchesRenameOption{
		Name:    "MUST_EDIT_IT",
		Project: "MUST_EDIT_IT",
	}
	resp, err := client.ProjectBranches.Rename(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// ProjectLinksCreateFunc testing Create a new project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func ProjectLinksCreateFunc() {
	opt := &ProjectLinksCreateOption{
		Name:       "MUST_EDIT_IT",
		ProjectId:  "",
		ProjectKey: "",
		Url:        "MUST_EDIT_IT",
	}
	v, resp, err := client.ProjectLinks.Create(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectLinksDeleteFunc testing Delete existing project link.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func ProjectLinksDeleteFunc() {
	opt := &ProjectLinksDeleteOption{Id: "MUST_EDIT_IT"}
	resp, err := client.ProjectLinks.Delete(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// ProjectLinksSearchFunc testing List links of a project.<br>The 'projectId' or 'projectKey' must be provided.<br>Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
func ProjectLinksSearchFunc() {
	opt := &ProjectLinksSearchOption{
		ProjectId:  "",
		ProjectKey: "",
	}
	v, resp, err := client.ProjectLinks.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectPullRequestsDeleteFunc testing Delete a pull request.<br/>Requires 'Administer' rights on the specified project.
func ProjectPullRequestsDeleteFunc() {
	opt := &ProjectPullRequestsDeleteOption{
		Project:     "MUST_EDIT_IT",
		PullRequest: 0,
	}
	resp, err := client.ProjectPullRequests.Delete(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// ProjectPullRequestsListFunc testing List the pull requests of a project.<br/>One of the following permissions is required: <ul><li>'Browse' rights on the specified project</li><li>'Execute Analysis' rights on the specified project</li></ul>
func ProjectPullRequestsListFunc() {
	opt := &ProjectPullRequestsListOption{Project: "MUST_EDIT_IT"}
	v, resp, err := client.ProjectPullRequests.List(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectTagsSearchFunc testing Search tags
func ProjectTagsSearchFunc() {
	opt := &ProjectTagsSearchOption{
		Ps: 0,
		Q:  "",
	}
	v, resp, err := client.ProjectTags.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectTagsSetFunc testing Set tags on a project.<br>Requires the following permission: 'Administer' rights on the specified project
func ProjectTagsSetFunc() {
	opt := &ProjectTagsSetOption{
		Project: "MUST_EDIT_IT",
		Tags:    nil,
	}
	resp, err := client.ProjectTags.Set(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// ProjectsBulkDeleteFunc testing Delete one or several projects.<br />Requires 'Administer System' permission.
func ProjectsBulkDeleteFunc() {
	opt := &ProjectsBulkDeleteOption{
		AnalyzedBefore:    "",
		OnProvisionedOnly: "",
		ProjectIds:        "",
		Projects:          "",
		Q:                 "",
		Qualifiers:        "",
	}
	resp, err := client.Projects.BulkDelete(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// ProjectsBulkUpdateKeyFunc testing Bulk update a project or module key and all its sub-components keys. The bulk update allows to replace a part of the current key by another string on the current project and all its sub-modules.<br>It's possible to simulate the bulk update by setting the parameter 'dryRun' at true. No key is updated with a dry run.<br>Ex: to rename a project with key 'my_project' to 'my_new_project' and all its sub-components keys, call the WS with parameters:<ul>  <li>project: my_project</li>  <li>from: my_</li>  <li>to: my_new_</li></ul>Either 'projectId' or 'project' must be provided.<br> Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func ProjectsBulkUpdateKeyFunc() {
	opt := &ProjectsBulkUpdateKeyOption{
		DryRun:    "",
		From:      "MUST_EDIT_IT",
		Project:   "",
		ProjectId: "",
		To:        "MUST_EDIT_IT",
	}
	v, resp, err := client.Projects.BulkUpdateKey(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectsCreateFunc testing Create a project.<br/>Requires 'Create Projects' permission
func ProjectsCreateFunc() {
	opt := &ProjectsCreateOption{
		Branch:     "",
		Name:       "MUST_EDIT_IT",
		Project:    "MUST_EDIT_IT",
		Visibility: "",
	}
	v, resp, err := client.Projects.Create(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectsDeleteFunc testing Delete a project.<br> Requires 'Administer System' permission or 'Administer' permission on the project.
func ProjectsDeleteFunc() {
	opt := &ProjectsDeleteOption{
		Project:   "",
		ProjectId: "",
	}
	resp, err := client.Projects.Delete(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// ProjectsSearchFunc testing Search for projects or views to administrate them.<br>Requires 'System Administrator' permission
func ProjectsSearchFunc() {
	opt := &ProjectsSearchOption{
		AnalyzedBefore:    "",
		OnProvisionedOnly: "",
		P:                 "",
		ProjectIds:        "",
		Projects:          "",
		Ps:                "",
		Q:                 "",
		Qualifiers:        "",
	}
	v, resp, err := client.Projects.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// ProjectsUpdateKeyFunc testing Update a project or module key and all its sub-components keys.<br>Either 'from' or 'projectId' must be provided.<br> Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func ProjectsUpdateKeyFunc() {
	opt := &ProjectsUpdateKeyOption{
		From:      "",
		ProjectId: "",
		To:        "MUST_EDIT_IT",
	}
	resp, err := client.Projects.UpdateKey(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// ProjectsUpdateVisibilityFunc testing Updates visibility of a project.<br>Requires 'Project administer' permission on the specified project
func ProjectsUpdateVisibilityFunc() {
	opt := &ProjectsUpdateVisibilityOption{
		Project:    "MUST_EDIT_IT",
		Visibility: "MUST_EDIT_IT",
	}
	resp, err := client.Projects.UpdateVisibility(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualitygatesCopyFunc testing Copy a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func QualitygatesCopyFunc() {
	opt := &QualitygatesCopyOption{
		Id:           1,
		Name:         "test-gate-copy",
		Organization: "",
	}
	_, resp, err := client.Qualitygates.Copy(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	printBody(resp.Body)
}

// QualitygatesCreateFunc testing Create a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func QualitygatesCreateFunc() {
	opt := &QualitygatesCreateOption{
		Name:         "MUST_EDIT_IT",
		Organization: "",
	}
	v, resp, err := client.Qualitygates.Create(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualitygatesCreateConditionFunc testing Add a new condition to a quality gate.<br>Requires the 'Administer Quality Gates' permission.
func QualitygatesCreateConditionFunc() {
	opt := &QualitygatesCreateConditionOption{
		Error:        "",
		GateId:       0,
		Metric:       "MUST_EDIT_IT",
		Op:           "",
		Organization: "",
		Period:       "",
		Warning:      "",
	}
	v, resp, err := client.Qualitygates.CreateCondition(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualitygatesDeleteConditionFunc testing Delete a condition from a quality gate.<br>Requires the 'Administer Quality Gates' permission.
func QualitygatesDeleteConditionFunc() {
	opt := &QualitygatesDeleteConditionOption{
		ConditionID:  0,
		Organization: "",
	}
	resp, err := client.Qualitygates.DeleteCondition(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualitygatesDeselectFunc testing Remove the association of a project from a quality gate.<br>Requires one of the following permissions:<ul><li>'Administer Quality Gates'</li><li>'Administer' rights on the project</li></ul>
func QualitygatesDeselectFunc() {
	opt := &QualitygatesDeselectOption{
		Organization: "",
		ProjectId:    "",
		ProjectKey:   "",
	}
	resp, err := client.Qualitygates.Deselect(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualitygatesDestroyFunc testing Delete a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func QualitygatesDestroyFunc() {
	opt := &QualitygatesDestroyOption{
		Id:           0,
		Organization: "",
	}
	resp, err := client.Qualitygates.Destroy(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualitygatesGetByProjectFunc testing Get the quality gate of a project.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
func QualitygatesGetByProjectFunc() {
	opt := &QualitygatesGetByProjectOption{
		Organization: "",
		Project:      "MUST_EDIT_IT",
	}
	v, resp, err := client.Qualitygates.GetByProject(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualitygatesListFunc testing Get a list of quality gates
func QualitygatesListFunc() {
	opt := &QualitygatesListOption{Organization: ""}
	v, resp, err := client.Qualitygates.List(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualitygatesProjectStatusFunc testing Get the quality gate status of a project or a Compute Engine task.<br />Either 'analysisId', 'projectId' or 'projectKey' must be provided<br />The different statuses returned are: OK, WARN, ERROR, NONE. The NONE status is returned when there is no quality gate associated with the analysis.<br />Returns an HTTP code 404 if the analysis associated with the task is not found or does not exist.<br />Requires one of the following permissions:<ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li><li>'Browse' on the specified project</li></ul>
func QualitygatesProjectStatusFunc() {
	opt := &QualitygatesProjectStatusOption{
		AnalysisId: "",
		ProjectId:  "",
		ProjectKey: "",
	}
	v, resp, err := client.Qualitygates.ProjectStatus(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualitygatesRenameFunc testing Rename a Quality Gate.<br>Requires the 'Administer Quality Gates' permission.
func QualitygatesRenameFunc() {
	opt := &QualitygatesRenameOption{
		Id:           0,
		Name:         "MUST_EDIT_IT",
		Organization: "",
	}
	_, resp, err := client.Qualitygates.Rename(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualitygatesSearchFunc testing Search for projects associated (or not) to a quality gate.<br/>Only authorized projects for current user will be returned.
func QualitygatesSearchFunc() {
	opt := &QualitygatesSearchOption{
		GateId:       0,
		Organization: "",
		Page:         "",
		PageSize:     "",
		Query:        "",
		Selected:     "",
	}
	v, resp, err := client.Qualitygates.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualitygatesSelectFunc testing Associate a project to a quality gate.<br>The 'projectId' or 'projectKey' must be provided.<br>Project id as a numeric value is deprecated since 6.1. Please use the id similar to 'AU-TpxcA-iU5OvuD2FLz'.<br>Requires the 'Administer Quality Gates' permission.
func QualitygatesSelectFunc() {
	opt := &QualitygatesSelectOption{
		GateId:       0,
		Organization: "",
		ProjectKey:   "",
	}
	resp, err := client.Qualitygates.Select(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualitygatesSetAsDefaultFunc testing Set a quality gate as the default quality gate.<br>Requires the 'Administer Quality Gates' permission.
func QualitygatesSetAsDefaultFunc() {
	opt := &QualitygatesSetAsDefaultOption{
		Id:           0,
		Organization: "",
	}
	resp, err := client.Qualitygates.SetAsDefault(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualitygatesShowFunc testing Display the details of a quality gate
func QualitygatesShowFunc() {
	opt := &QualitygatesShowOption{
		Id:           0,
		Name:         "",
		Organization: "",
	}
	v, resp, err := client.Qualitygates.Show(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualitygatesUpdateConditionFunc testing Update a condition attached to a quality gate.<br>Requires the 'Administer Quality Gates' permission.
func QualitygatesUpdateConditionFunc() {
	opt := &QualitygatesUpdateConditionOption{
		Error:        "",
		Id:           0,
		Metric:       "MUST_EDIT_IT",
		Op:           "",
		Organization: "",
		Period:       "",
		Warning:      "",
	}
	_, resp, err := client.Qualitygates.UpdateCondition(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesActivateRuleFunc testing Activate a rule on a Quality Profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func QualityProfilesActivateRuleFunc() {
	opt := &QualityProfilesActivateRuleOption{
		Key:      "MUST_EDIT_IT",
		Params:   "",
		Reset:    "",
		Rule:     "MUST_EDIT_IT",
		Severity: "",
	}
	resp, err := client.QualityProfiles.ActivateRule(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesActivateRulesFunc testing Bulk-activate rules on one quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func QualityProfilesActivateRulesFunc() {
	opt := &QualityProfilesActivateRulesOption{
		Activation:       "",
		ActiveSeverities: "",
		Asc:              "",
		AvailableSince:   "",
		Inheritance:      "",
		IsTemplate:       "",
		Languages:        "",
		Q:                "",
		Qprofile:         "",
		Repositories:     "",
		RuleKey:          "",
		S:                "",
		Severities:       "",
		Statuses:         "",
		Tags:             "",
		TargetKey:        "MUST_EDIT_IT",
		TargetSeverity:   "",
		TemplateKey:      "",
		Types:            "",
	}
	_, resp, err := client.QualityProfiles.ActivateRules(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesAddProjectFunc testing Associate a project with a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li>  <li>Administer right on the specified project</li></ul>
func QualityProfilesAddProjectFunc() {
	opt := &QualityProfilesAddProjectOption{
		Key:            "",
		Language:       "",
		Project:        "",
		ProjectUuid:    "",
		QualityProfile: "",
	}
	resp, err := client.QualityProfiles.AddProject(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesBackupFunc testing Backup a quality profile in XML form. The exported profile can be restored through api/qualityprofiles/restore.
func QualityProfilesBackupFunc() {
	//http://192.168.98.8:9000/api/qualityprofiles/backup?profileKey=AWdZhiracZBxYk_CF_We
	opt := &QualityProfilesBackupOption{
		ProfileKey: "AWdZhiracZBxYk_CF_We",
	}
	v, resp, err := client.QualityProfiles.Backup(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualityProfilesChangeParentFunc testing Change a quality profile's parent.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func QualityProfilesChangeParentFunc() {
	opt := &QualityProfilesChangeParentOption{
		Key:                  "",
		Language:             "",
		ParentKey:            "",
		ParentQualityProfile: "",
		QualityProfile:       "",
	}
	resp, err := client.QualityProfiles.ChangeParent(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesChangelogFunc testing Get the history of changes on a quality profile: rule activation/deactivation, change in parameters/severity. Events are ordered by date in descending order (most recent first).
func QualityProfilesChangelogFunc() {
	opt := &QualityProfilesChangelogOption{
		Key:            "",
		Language:       "",
		P:              "",
		Ps:             "",
		QualityProfile: "",
		Since:          "",
		To:             "",
	}
	v, resp, err := client.QualityProfiles.Changelog(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualityProfilesCopyFunc testing Copy a quality profile.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func QualityProfilesCopyFunc() {
	opt := &QualityProfilesCopyOption{
		FromKey: "MUST_EDIT_IT",
		ToName:  "MUST_EDIT_IT",
	}
	v, resp, err := client.QualityProfiles.Copy(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualityProfilesCreateFunc testing Create a quality profile.<br>Requires to be logged in and the 'Administer Quality Profiles' permission.
func QualityProfilesCreateFunc() {
	opt := &QualityProfilesCreateOption{
		BackupSonarlintVsCsFake: "",
		Language:                "MUST_EDIT_IT",
		Name:                    "MUST_EDIT_IT",
	}
	v, resp, err := client.QualityProfiles.Create(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualityProfilesDeactivateRuleFunc testing Deactivate a rule on a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func QualityProfilesDeactivateRuleFunc() {
	opt := &QualityProfilesDeactivateRuleOption{
		Key:  "MUST_EDIT_IT",
		Rule: "MUST_EDIT_IT",
	}
	resp, err := client.QualityProfiles.DeactivateRule(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesDeactivateRulesFunc testing Bulk deactivate rules on Quality profiles.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func QualityProfilesDeactivateRulesFunc() {
	opt := &QualityProfilesDeactivateRulesOption{
		Activation:       "",
		ActiveSeverities: "",
		Asc:              "",
		AvailableSince:   "",
		Inheritance:      "",
		IsTemplate:       "",
		Languages:        "",
		Q:                "",
		Qprofile:         "",
		Repositories:     "",
		RuleKey:          "",
		S:                "",
		Severities:       "",
		Statuses:         "",
		Tags:             "",
		TargetKey:        "MUST_EDIT_IT",
		TemplateKey:      "",
		Types:            "",
	}
	_, resp, err := client.QualityProfiles.DeactivateRules(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesDeleteFunc testing Delete a quality profile and all its descendants. The default quality profile cannot be deleted.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func QualityProfilesDeleteFunc() {
	opt := &QualityProfilesDeleteOption{
		Key:            "",
		Language:       "",
		QualityProfile: "",
	}
	resp, err := client.QualityProfiles.Delete(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesExportFunc testing Export a quality profile.
func QualityProfilesExportFunc() {
	opt := &QualityProfilesExportOption{
		ExporterKey:    "",
		Key:            "",
		Language:       "",
		QualityProfile: "",
	}
	v, resp, err := client.QualityProfiles.Export(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualityProfilesExportersFunc testing Lists available profile export formats.
func QualityProfilesExportersFunc() {
	v, resp, err := client.QualityProfiles.Exporters()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualityProfilesImportersFunc testing List supported importers.
func QualityProfilesImportersFunc() {
	v, resp, err := client.QualityProfiles.Importers()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualityProfilesInheritanceFunc testing Show a quality profile's ancestors and children.
func QualityProfilesInheritanceFunc() {
	opt := &QualityProfilesInheritanceOption{
		Key:            "",
		Language:       "",
		QualityProfile: "",
	}
	v, resp, err := client.QualityProfiles.Inheritance(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualityProfilesProjectsFunc testing List projects with their association status regarding a quality profile
func QualityProfilesProjectsFunc() {
	opt := &QualityProfilesProjectsOption{
		Key:      "MUST_EDIT_IT",
		P:        0,
		Ps:       0,
		Q:        "",
		Selected: "",
	}
	v, resp, err := client.QualityProfiles.Projects(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualityProfilesRemoveProjectFunc testing Remove a project's association with a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li>  <li>Administer right on the specified project</li></ul>
func QualityProfilesRemoveProjectFunc() {
	opt := &QualityProfilesRemoveProjectOption{
		Key:            "",
		Language:       "",
		Project:        "",
		ProjectUuid:    "",
		QualityProfile: "",
	}
	resp, err := client.QualityProfiles.RemoveProject(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesRenameFunc testing Rename a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func QualityProfilesRenameFunc() {
	opt := &QualityProfilesRenameOption{
		Key:  "MUST_EDIT_IT",
		Name: "MUST_EDIT_IT",
	}
	resp, err := client.QualityProfiles.Rename(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesRestoreFunc testing Restore a quality profile using an XML file. The restored profile name is taken from the backup file, so if a profile with the same name and language already exists, it will be overwritten.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func QualityProfilesRestoreFunc() {
	opt := &QualityProfilesRestoreOption{Backup: "MUST_EDIT_IT"}
	resp, err := client.QualityProfiles.Restore(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesRestoreBuiltInFunc testing This web service has no effect since 6.4. It's no more possible to restore built-in quality profiles because they are automatically updated and read only. Returns HTTP code 410.
func QualityProfilesRestoreBuiltInFunc() {
	resp, err := client.QualityProfiles.RestoreBuiltIn()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// QualityProfilesSearchFunc testing Search quality profiles
func QualityProfilesSearchFunc() {
	opt := &QualityProfilesSearchOption{
		Defaults:       "",
		Language:       "",
		Project:        "",
		QualityProfile: "",
	}
	v, resp, err := client.QualityProfiles.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// QualityProfilesSetDefaultFunc testing Select the default profile for a given language.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func QualityProfilesSetDefaultFunc() {
	opt := &QualityProfilesSetDefaultOption{
		Key:            "",
		Language:       "",
		QualityProfile: "",
	}
	resp, err := client.QualityProfiles.SetDefault(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// RulesCreateFunc testing Create a custom rule.<br>Requires the 'Administer Quality Profiles' permission
func RulesCreateFunc() {
	name := "magic_test_rule1"
	opt := &RulesCreateOption{
		CustomKey:           name,
		ManualKey:           "",
		MarkdownDescription: "This is the description of `rule`",
		Name:                name,
		Params:              "",
		PreventReactivation: "",
		Severity:            "INFO",
		Status:              "BETA",
		TemplateKey:         "squid:S3546",
		Type:                IssueTypeCodeSmell,
	}
	_, resp, err := client.Rules.Create(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	//printBody(resp.Body)
}

// RulesDeleteFunc testing Delete custom rule.<br/>Requires the 'Administer Quality Profiles' permission
func RulesDeleteFunc() {
	opt := &RulesDeleteOption{Key: "squid:magic_test_rule1"}
	resp, err := client.Rules.Delete(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	printBody(resp.Body)
}

// RulesRepositoriesFunc testing List available rule repositories
func RulesRepositoriesFunc() {
	opt := &RulesRepositoriesOption{
		Language: "",
		Q:        "",
	}
	v, resp, err := client.Rules.Repositories(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// RulesSearchFunc testing Search for a collection of relevant rules matching a specified query.<br/>Since 5.5, following fields in the response have been deprecated :<ul><li>"effortToFixDescription" becomes "gapDescription"</li><li>"debtRemFnCoeff" becomes "remFnGapMultiplier"</li><li>"defaultDebtRemFnCoeff" becomes "defaultRemFnGapMultiplier"</li><li>"debtRemFnOffset" becomes "remFnBaseEffort"</li><li>"defaultDebtRemFnOffset" becomes "defaultRemFnBaseEffort"</li><li>"debtOverloaded" becomes "remFnOverloaded"</li></ul>
func RulesSearchFunc() {
	opt := &RulesSearchOption{
		Activation:       "",
		ActiveSeverities: "",
		Asc:              "",
		AvailableSince:   "",
		Facets:           "",
		Inheritance:      "",
		IsTemplate:       "",
		Languages:        "",
		P:                "",
		Ps:               "",
		Q:                "",
		Qprofile:         "",
		Repositories:     "",
		RuleKey:          "",
		S:                "",
		Severities:       "",
		Statuses:         "",
		Tags:             "",
		TemplateKey:      "",
		Types:            "",
	}
	v, resp, err := client.Rules.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// RulesShowFunc testing Get detailed information about a rule<br>Since 5.5, following fields in the response have been deprecated :<ul><li>"effortToFixDescription" becomes "gapDescription"</li><li>"debtRemFnCoeff" becomes "remFnGapMultiplier"</li><li>"defaultDebtRemFnCoeff" becomes "defaultRemFnGapMultiplier"</li><li>"debtRemFnOffset" becomes "remFnBaseEffort"</li><li>"defaultDebtRemFnOffset" becomes "defaultRemFnBaseEffort"</li><li>"debtOverloaded" becomes "remFnOverloaded"</li></ul>In 7.1, the field 'scope' has been added.
func RulesShowFunc() {
	opt := &RulesShowOption{
		Actives: "",
		Key:     "MUST_EDIT_IT",
	}
	v, resp, err := client.Rules.Show(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// RulesTagsFunc testing List rule tags
func RulesTagsFunc() {
	opt := &RulesTagsOption{
		Ps: 0,
		Q:  "",
	}
	v, resp, err := client.Rules.Tags(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// RulesUpdateFunc testing Update an existing rule.<br>Requires the 'Administer Quality Profiles' permission
func RulesUpdateFunc() {
	opt := &RulesUpdateOption{
		DebtRemediationFnOffset: "",
		DebtRemediationFnType:   "",
		DebtRemediationFyCoeff:  "",
		DebtSubCharacteristic:   "",
		Key:                        "squid:magic_test_rule1",
		MarkdownDescription:        "",
		MarkdownNote:               "",
		Name:                       "",
		Params:                     "",
		RemediationFnBaseEffort:    "",
		RemediationFnType:          "",
		RemediationFyGapMultiplier: "",
		Severity:                   SeverityCRITICAL,
		Status:                     "",
		Tags:                       "",
	}
	_, resp, err := client.Rules.Update(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	printBody(resp.Body)
}

// ServerVersionFunc testing Version of SonarQube in plain text
func ServerVersionFunc() {
	v, resp, err := client.Server.Version()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SettingsListDefinitionsFunc testing List settings definitions.<br>Requires 'Browse' permission when a component is specified<br/>
func SettingsListDefinitionsFunc() {
	opt := &SettingsListDefinitionsOption{Component: ""}
	v, resp, err := client.Settings.ListDefinitions(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SettingsResetFunc testing Remove a setting value.<br>The settings defined in config/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>
func SettingsResetFunc() {
	opt := &SettingsResetOption{
		Component: "",
		Keys:      "MUST_EDIT_IT",
	}
	resp, err := client.Settings.Reset(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// SettingsSetFunc testing Update a setting value.<br>Either 'value' or 'values' must be provided.<br> The settings defined in config/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>
func SettingsSetFunc() {
	opt := &SettingsSetOption{
		Component:   "",
		FieldValues: "",
		Key:         "MUST_EDIT_IT",
		Value:       "",
	}
	resp, err := client.Settings.Set(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// SettingsValuesFunc testing List settings values.<br>If no value has been set for a setting, then the default value is returned.<br>The settings from conf/sonar.properties are excluded from results.<br>Requires 'Browse' or 'Execute Analysis' permission when a component is specified<br/>
func SettingsValuesFunc() {
	opt := &SettingsValuesOption{
		Component: "",
		Keys:      "",
	}
	v, resp, err := client.Settings.Values(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SourcesRawFunc testing Get source code as raw text. Require 'See Source Code' permission on file
func SourcesRawFunc() {
	opt := &SourcesRawOption{Key: "MUST_EDIT_IT"}
	v, resp, err := client.Sources.Raw(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SourcesScmFunc testing Get SCM information of source files. Require See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Author of the commit</li><li>Datetime of the commit (before 5.2 it was only the Date)</li><li>Revision of the commit (added in 5.2)</li></ol>
func SourcesScmFunc() {
	opt := &SourcesSCMOption{
		CommitsByLine: "",
		From:          "",
		Key:           "MUST_EDIT_IT",
		To:            "",
	}
	v, resp, err := client.Sources.SCM(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SourcesShowFunc testing Get source code. Require See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Content of the line</li></ol>
func SourcesShowFunc() {
	opt := &SourcesShowOption{
		From: "",
		Key:  "MUST_EDIT_IT",
		To:   "",
	}
	v, resp, err := client.Sources.Show(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SystemChangeLogLevelFunc testing Temporarily changes level of logs. New level is not persistent and is lost when restarting server. Requires system administration permission.
func SystemChangeLogLevelFunc() {
	opt := &SystemChangeLogLevelOption{Level: "MUST_EDIT_IT"}
	resp, err := client.System.ChangeLogLevel(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// SystemDbMigrationStatusFunc testing Display the database migration status of SonarQube.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>
func SystemDbMigrationStatusFunc() {
	v, resp, err := client.System.DbMigrationStatus()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SystemHealthFunc testing Provide health status of SonarQube.<p>Require 'Administer System' permission or authentication with passcode</p><p>  <ul> <li>GREEN: SonarQube is fully operational</li> <li>YELLOW: SonarQube is usable, but it needs attention in order to be fully operational</li> <li>RED: SonarQube is not operational</li> </ul></p>
func SystemHealthFunc() {
	v, resp, err := client.System.Health()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SystemLogsFunc testing Get system logs in plain-text format. Requires system administration permission.
func SystemLogsFunc() {
	opt := &SystemLogsOption{Process: ""}
	v, resp, err := client.System.Logs(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SystemMigrateDbFunc testing Migrate the database to match the current version of SonarQube.<br/>Sending a POST request to this URL starts the DB migration. It is strongly advised to <strong>make a database backup</strong> before invoking this WS.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>
func SystemMigrateDbFunc() {
	v, resp, err := client.System.MigrateDb()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SystemPingFunc testing Answers "pong" as plain-text
func SystemPingFunc() {
	v, resp, err := client.System.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SystemRestartFunc testing Restart server. Require 'Administer System' permission. Perform a full restart of the Web, Search and Compute Engine Servers processes.
func SystemRestartFunc() {
	resp, err := client.System.Restart()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// SystemStatusFunc testing Get state information about SonarQube.<p>status: the running status <ul> <li>STARTING: SonarQube Web Server is up and serving some Web Services (eg. api/system/status) but initialization is still ongoing</li> <li>UP: SonarQube instance is up and running</li> <li>DOWN: SonarQube instance is up but not running because migration has failed (refer to WS /api/system/migrate_db for details) or some other reason (check logs).</li> <li>RESTARTING: SonarQube instance is still up but a restart has been requested (refer to WS /api/system/restart for details).</li> <li>DB_MIGRATION_NEEDED: database migration is required. DB migration can be started using WS /api/system/migrate_db.</li> <li>DB_MIGRATION_RUNNING: DB migration is running (refer to WS /api/system/migrate_db for details)</li> </ul></p>
func SystemStatusFunc() {
	v, resp, err := client.System.Status()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// SystemUpgradesFunc testing Lists available upgrades for the SonarQube instance (if any) and for each one, lists incompatible plugins and plugins requiring upgrade.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.
func SystemUpgradesFunc() {
	v, resp, err := client.System.Upgrades()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// UserGroupsAddUserFunc testing Add a user to a group.<br />'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'.
func UserGroupsAddUserFunc() {
	opt := &UserGroupsAddUserOption{
		Id:    0,
		Login: "",
		Name:  "",
	}
	resp, err := client.UserGroups.AddUser(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// UserGroupsCreateFunc testing Create a group.<br>Requires the following permission: 'Administer System'.
func UserGroupsCreateFunc() {
	opt := &UserGroupsCreateOption{
		Description: "",
		Name:        "test1",
	}
	v, resp, err := client.UserGroups.Create(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// UserGroupsDeleteFunc testing Delete a group. The default groups cannot be deleted.<br/>'id' or 'name' must be provided.<br />Requires the following permission: 'Administer System'.
func UserGroupsDeleteFunc() {
	opt := &UserGroupsDeleteOption{
		Id:   0,
		Name: "test1",
	}
	resp, err := client.UserGroups.Delete(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// UserGroupsRemoveUserFunc testing Remove a user from a group.<br />'id' or 'name' must be provided.<br>Requires the following permission: 'Administer System'.
func UserGroupsRemoveUserFunc() {
	opt := &UserGroupsRemoveUserOption{
		Id:    0,
		Login: "",
		Name:  "",
	}
	resp, err := client.UserGroups.RemoveUser(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// UserGroupsSearchFunc testing Search for user groups.<br>Requires the following permission: 'Administer System'.
func UserGroupsSearchFunc() {
	opt := &UserGroupsSearchOption{
		F:  "",
		P:  0,
		Ps: 0,
		Q:  "",
	}
	v, resp, err := client.UserGroups.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

func printBody(body io.Reader) {
	str, _ := ioutil.ReadAll(body)
	fmt.Print(string(str))
}

// UserGroupsUpdateFunc testing Update a group.<br>Requires the following permission: 'Administer System'.
func UserGroupsUpdateFunc() {
	opt := &UserGroupsUpdateOption{
		Description: "hdadasd",
		Id:          4,
		Name:        "",
	}
	_, resp, err := client.UserGroups.Update(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	printBody(resp.Body)
}

// UserGroupsUsersFunc testing Search for users with membership information with respect to a group.<br>Requires the following permission: 'Administer System'.
func UserGroupsUsersFunc() {
	opt := &UserGroupsUsersOption{
		Id:       0,
		Name:     "",
		P:        0,
		Ps:       0,
		Q:        "",
		Selected: "",
	}
	v, resp, err := client.UserGroups.Users(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// UserTokensGenerateFunc testing Generate a user access token. <br />Please keep your tokens secret. They enable to authenticate and analyze projects.<br />If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func UserTokensGenerateFunc() {
	opt := &UserTokensGenerateOption{
		Login: "",
		Name:  "MUST_EDIT_IT",
	}
	v, resp, err := client.UserTokens.Generate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// UserTokensRevokeFunc testing Revoke a user access token. <br/>If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func UserTokensRevokeFunc() {
	opt := &UserTokensRevokeOption{
		Login: "",
		Name:  "MUST_EDIT_IT",
	}
	resp, err := client.UserTokens.Revoke(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// UserTokensSearchFunc testing List the access tokens of a user.<br>The login must exist and active.<br>If the login is set, it requires administration permissions. Otherwise, a token is generated for the authenticated user.
func UserTokensSearchFunc() {
	opt := &UserTokensSearchOption{Login: ""}
	v, resp, err := client.UserTokens.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// UsersChangePasswordFunc testing Update a user's password. Authenticated users can change their own password, provided that the account is not linked to an external authentication system. Administer System permission is required to change another user's password.
func UsersChangePasswordFunc() {
	opt := &UsersChangePasswordOption{
		Login:            "MUST_EDIT_IT",
		Password:         "MUST_EDIT_IT",
		PreviousPassword: "",
	}
	resp, err := client.Users.ChangePassword(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// UsersCreateFunc testing Create a user.<br/>If a deactivated user account exists with the given login, it will be reactivated.<br/>Requires Administer System permission
func UsersCreateFunc() {
	opt := &UsersCreateOption{
		Email:      "",
		Local:      "",
		Login:      "MUST_EDIT_IT",
		Name:       "MUST_EDIT_IT",
		Password:   "",
		ScmAccount: "",
	}
	v, resp, err := client.Users.Create(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// UsersDeactivateFunc testing Deactivate a user. Requires Administer System permission
func UsersDeactivateFunc() {
	opt := &UsersDeactivateOption{Login: "testone"}
	v, resp, err := client.Users.Deactivate(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// UsersGroupsFunc testing Lists the groups a user belongs to. <br/>Requires Administer System permission.
func UsersGroupsFunc() {
	opt := &UsersGroupsOption{
		Login:    "MUST_EDIT_IT",
		P:        0,
		Ps:       0,
		Q:        "",
		Selected: "",
	}
	v, resp, err := client.Users.Groups(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// UsersSearchFunc testing Get a list of active users. <br/>Administer System permission is required to show the 'groups' field.<br/>When accessed anonymously, only logins and names are returned.
func UsersSearchFunc() {
	opt := &UsersSearchOption{
		F:  "",
		P:  0,
		Ps: 0,
		Q:  "",
	}
	v, resp, err := client.Users.Search(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// UsersUpdateFunc testing Update a user. If a deactivated user account exists with the given login, it will be reactivated. Requires Administer System permission
func UsersUpdateFunc() {
	opt := &UsersUpdateOption{
		Email:      "",
		Login:      "testone",
		Name:       "testone",
		ScmAccount: "",
	}
	v, resp, err := client.Users.Update(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// WebhooksCreateFunc testing Create a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func WebhooksCreateFunc() {
	opt := &WebhooksCreateOption{
		Name:    "MUST_EDIT_IT",
		Project: "",
		Url:     "MUST_EDIT_IT",
	}
	v, resp, err := client.Webhooks.Create(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// WebhooksDeleteFunc testing Delete a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func WebhooksDeleteFunc() {
	opt := &WebhooksDeleteOption{Webhook: "MUST_EDIT_IT"}
	resp, err := client.Webhooks.Delete(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}

// WebhooksDeliveriesFunc testing Get the recent deliveries for a specified project or Compute Engine task.<br/>Require 'Administer' permission on the related project.<br/>Note that additional information are returned by api/webhooks/delivery.
func WebhooksDeliveriesFunc() {
	opt := &WebhooksDeliveriesOption{
		CeTaskId:     "",
		ComponentKey: "",
		P:            "",
		Ps:           "",
		Webhook:      "",
	}
	v, resp, err := client.Webhooks.Deliveries(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// WebhooksDeliveryFunc testing Get a webhook delivery by its id.<br/>Require 'Administer System' permission.<br/>Note that additional information are returned by api/webhooks/delivery.
func WebhooksDeliveryFunc() {
	opt := &WebhooksDeliveryOption{DeliveryId: "MUST_EDIT_IT"}
	v, resp, err := client.Webhooks.Delivery(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// WebhooksListFunc testing Search for global webhooks or project webhooks. Webhooks are ordered by name.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func WebhooksListFunc() {
	opt := &WebhooksListOption{Project: ""}
	v, resp, err := client.Webhooks.List(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
	spew.Dump(v)
}

// WebhooksUpdateFunc testing Update a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func WebhooksUpdateFunc() {
	opt := &WebhooksUpdateOption{
		Name:    "MUST_EDIT_IT",
		Url:     "MUST_EDIT_IT",
		Webhook: "MUST_EDIT_IT",
	}
	resp, err := client.Webhooks.Update(opt)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(resp.StatusCode)
}
