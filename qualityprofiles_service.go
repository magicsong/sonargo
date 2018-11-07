// Manage quality profiles.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type QualityprofilesService struct {
	client *Client
}

// [TODO] you should call the <ActivateRule> func manually and complete the fields of this struct
type QualityprofilesActivateRuleResp struct{}
type QualityprofilesActivateRuleOption struct {
	Key      string `url:"key,omitempty"`      // Description:"Quality Profile key. Can be obtained through <code>api/qualityprofiles/search</code>",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Params   string `url:"params,omitempty"`   // Description:"Parameters as semi-colon list of <code>key=value</code>. Ignored if parameter reset is true.",ExampleValue:"params=key1=v1;key2=v2"
	Reset    string `url:"reset,omitempty"`    // Description:"Reset severity and parameters of activated rule. Set the values defined on parent profile or from rule default values.",ExampleValue:""
	Rule     string `url:"rule,omitempty"`     // Description:"Rule key",ExampleValue:"squid:AvoidCycles"
	Severity string `url:"severity,omitempty"` // Description:"Severity. Ignored if parameter reset is true.",ExampleValue:""
}

// Activate_rule Activate a rule on a Quality Profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) Activate_rule(opt *QualityprofilesActivateRuleOption) (Resp *QualityprofilesActivateRuleResp, err error) {
	err := s.ValidateActivate_ruleOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/activate_rule", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <ActivateRules> func manually and complete the fields of this struct
type QualityprofilesActivateRulesResp struct{}
type QualityprofilesActivateRulesOption struct {
	Activation       string `url:"activation,omitempty"`        // Description:"Filter rules that are activated or deactivated on the selected Quality profile. Ignored if the parameter 'qprofile' is not set.",ExampleValue:""
	ActiveSeverities string `url:"active_severities,omitempty"` // Description:"Comma-separated list of activation severities, i.e the severity of rules in Quality profiles.",ExampleValue:"CRITICAL,BLOCKER"
	Asc              string `url:"asc,omitempty"`               // Description:"Ascending sort",ExampleValue:""
	AvailableSince   string `url:"available_since,omitempty"`   // Description:"Filters rules added since date. Format is yyyy-MM-dd",ExampleValue:"2014-06-22"
	Inheritance      string `url:"inheritance,omitempty"`       // Description:"Comma-separated list of values of inheritance for a rule within a quality profile. Used only if the parameter 'activation' is set.",ExampleValue:"INHERITED,OVERRIDES"
	IsTemplate       string `url:"is_template,omitempty"`       // Description:"Filter template rules",ExampleValue:""
	Languages        string `url:"languages,omitempty"`         // Description:"Comma-separated list of languages",ExampleValue:"java,js"
	Q                string `url:"q,omitempty"`                 // Description:"UTF-8 search query",ExampleValue:"xpath"
	Qprofile         string `url:"qprofile,omitempty"`          // Description:"Quality profile key to filter on. Used only if the parameter 'activation' is set.",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Repositories     string `url:"repositories,omitempty"`      // Description:"Comma-separated list of repositories",ExampleValue:"checkstyle,findbugs"
	RuleKey          string `url:"rule_key,omitempty"`          // Description:"Key of rule to search for",ExampleValue:"squid:S001"
	S                string `url:"s,omitempty"`                 // Description:"Sort field",ExampleValue:"name"
	Severities       string `url:"severities,omitempty"`        // Description:"Comma-separated list of default severities. Not the same than severity of rules in Quality profiles.",ExampleValue:"CRITICAL,BLOCKER"
	Statuses         string `url:"statuses,omitempty"`          // Description:"Comma-separated list of status codes",ExampleValue:"READY"
	Tags             string `url:"tags,omitempty"`              // Description:"Comma-separated list of tags. Returned rules match any of the tags (OR operator)",ExampleValue:"security,java8"
	TargetKey        string `url:"targetKey,omitempty"`         // Description:"Quality Profile key on which the rule activation is done. To retrieve a quality profile key please see <code>api/qualityprofiles/search</code>",ExampleValue:"AU-TpxcA-iU5OvuD2FL0"
	TargetSeverity   string `url:"targetSeverity,omitempty"`    // Description:"Severity to set on the activated rules",ExampleValue:""
	TemplateKey      string `url:"template_key,omitempty"`      // Description:"Key of the template rule to filter on. Used to search for the custom rules based on this template.",ExampleValue:"java:S001"
	Types            string `url:"types,omitempty"`             // Description:"Comma-separated list of types. Returned rules match any of the tags (OR operator)",ExampleValue:"BUG"
}

// Activate_rules Bulk-activate rules on one quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) Activate_rules(opt *QualityprofilesActivateRulesOption) (Resp *QualityprofilesActivateRulesResp, err error) {
	err := s.ValidateActivate_rulesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/activate_rules", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <AddProject> func manually and complete the fields of this struct
type QualityprofilesAddProjectResp struct{}
type QualityprofilesAddProjectOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	Project        string `url:"project,omitempty"`        // Description:"Project key",ExampleValue:"my_project"
	ProjectUuid    string `url:"projectUuid,omitempty"`    // Description:"Project ID. Either this parameter or 'project' must be set.",ExampleValue:"AU-TpxcA-iU5OvuD2FL5"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Add_project Associate a project with a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li>  <li>Administer right on the specified project</li></ul>
func (s *QualityprofilesService) Add_project(opt *QualityprofilesAddProjectOption) (Resp *QualityprofilesAddProjectResp, err error) {
	err := s.ValidateAdd_projectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/add_project", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Backup> func manually and complete the fields of this struct
type QualityprofilesBackupResp struct{}
type QualityprofilesBackupOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Backup Backup a quality profile in XML form. The exported profile can be restored through api/qualityprofiles/restore.
func (s *QualityprofilesService) Backup(opt *QualityprofilesBackupOption) (Resp *QualityprofilesBackupResp, err error) {
	err := s.ValidateBackupOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/backup", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <ChangeParent> func manually and complete the fields of this struct
type QualityprofilesChangeParentResp struct{}
type QualityprofilesChangeParentOption struct {
	Key                  string `url:"key,omitempty"`                  // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language             string `url:"language,omitempty"`             // Description:"Quality profile language",ExampleValue:""
	ParentKey            string `url:"parentKey,omitempty"`            // Description:"New parent profile key.<br> If no profile is provided, the inheritance link with current parent profile (if any) is broken, which deactivates all rules which come from the parent and are not overridden.",ExampleValue:"AU-TpxcA-iU5OvuD2FLz"
	ParentQualityProfile string `url:"parentQualityProfile,omitempty"` // Description:"Quality profile name. If this parameter is set, 'parentKey' must not be set and 'language' must be set to disambiguate.",ExampleValue:"Sonar way"
	QualityProfile       string `url:"qualityProfile,omitempty"`       // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Change_parent Change a quality profile's parent.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) Change_parent(opt *QualityprofilesChangeParentOption) (Resp *QualityprofilesChangeParentResp, err error) {
	err := s.ValidateChange_parentOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/change_parent", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Changelog> func manually and complete the fields of this struct
type QualityprofilesChangelogResp struct{}
type QualityprofilesChangelogOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	P              string `url:"p,omitempty"`              // Description:"1-based page number",ExampleValue:"42"
	Ps             string `url:"ps,omitempty"`             // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
	Since          string `url:"since,omitempty"`          // Description:"Start date for the changelog. <br>Either a date (server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	To             string `url:"to,omitempty"`             // Description:"End date for the changelog. <br>Either a date (server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
}

// Changelog Get the history of changes on a quality profile: rule activation/deactivation, change in parameters/severity. Events are ordered by date in descending order (most recent first).
func (s *QualityprofilesService) Changelog(opt *QualityprofilesChangelogOption) (Resp *QualityprofilesChangelogResp, err error) {
	err := s.ValidateChangelogOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/changelog", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Copy> func manually and complete the fields of this struct
type QualityprofilesCopyResp struct{}
type QualityprofilesCopyOption struct {
	FromKey string `url:"fromKey,omitempty"` // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ToName  string `url:"toName,omitempty"`  // Description:"Name for the new quality profile.",ExampleValue:"My Sonar way"
}

// Copy Copy a quality profile.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) Copy(opt *QualityprofilesCopyOption) (Resp *QualityprofilesCopyResp, err error) {
	err := s.ValidateCopyOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/copy", Opt)
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
type QualityprofilesCreateResp struct{}
type QualityprofilesCreateOption struct {
	BackupSonarlintVsCsFake string `url:"backup_sonarlint-vs-cs-fake,omitempty"` // Description:"A configuration file for Technical importer for the MSBuild SonarQube Scanner.",ExampleValue:""
	Language                string `url:"language,omitempty"`                    // Description:"Quality profile language",ExampleValue:"js"
	Name                    string `url:"name,omitempty"`                        // Description:"Quality profile name",ExampleValue:"My Sonar way"
}

// Create Create a quality profile.<br>Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) Create(opt *QualityprofilesCreateOption) (Resp *QualityprofilesCreateResp, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/create", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <DeactivateRule> func manually and complete the fields of this struct
type QualityprofilesDeactivateRuleResp struct{}
type QualityprofilesDeactivateRuleOption struct {
	Key  string `url:"key,omitempty"`  // Description:"Quality Profile key. Can be obtained through <code>api/qualityprofiles/search</code>",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Rule string `url:"rule,omitempty"` // Description:"Rule key",ExampleValue:"squid:AvoidCycles"
}

// Deactivate_rule Deactivate a rule on a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) Deactivate_rule(opt *QualityprofilesDeactivateRuleOption) (Resp *QualityprofilesDeactivateRuleResp, err error) {
	err := s.ValidateDeactivate_ruleOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/deactivate_rule", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <DeactivateRules> func manually and complete the fields of this struct
type QualityprofilesDeactivateRulesResp struct{}
type QualityprofilesDeactivateRulesOption struct {
	Activation       string `url:"activation,omitempty"`        // Description:"Filter rules that are activated or deactivated on the selected Quality profile. Ignored if the parameter 'qprofile' is not set.",ExampleValue:""
	ActiveSeverities string `url:"active_severities,omitempty"` // Description:"Comma-separated list of activation severities, i.e the severity of rules in Quality profiles.",ExampleValue:"CRITICAL,BLOCKER"
	Asc              string `url:"asc,omitempty"`               // Description:"Ascending sort",ExampleValue:""
	AvailableSince   string `url:"available_since,omitempty"`   // Description:"Filters rules added since date. Format is yyyy-MM-dd",ExampleValue:"2014-06-22"
	Inheritance      string `url:"inheritance,omitempty"`       // Description:"Comma-separated list of values of inheritance for a rule within a quality profile. Used only if the parameter 'activation' is set.",ExampleValue:"INHERITED,OVERRIDES"
	IsTemplate       string `url:"is_template,omitempty"`       // Description:"Filter template rules",ExampleValue:""
	Languages        string `url:"languages,omitempty"`         // Description:"Comma-separated list of languages",ExampleValue:"java,js"
	Q                string `url:"q,omitempty"`                 // Description:"UTF-8 search query",ExampleValue:"xpath"
	Qprofile         string `url:"qprofile,omitempty"`          // Description:"Quality profile key to filter on. Used only if the parameter 'activation' is set.",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Repositories     string `url:"repositories,omitempty"`      // Description:"Comma-separated list of repositories",ExampleValue:"checkstyle,findbugs"
	RuleKey          string `url:"rule_key,omitempty"`          // Description:"Key of rule to search for",ExampleValue:"squid:S001"
	S                string `url:"s,omitempty"`                 // Description:"Sort field",ExampleValue:"name"
	Severities       string `url:"severities,omitempty"`        // Description:"Comma-separated list of default severities. Not the same than severity of rules in Quality profiles.",ExampleValue:"CRITICAL,BLOCKER"
	Statuses         string `url:"statuses,omitempty"`          // Description:"Comma-separated list of status codes",ExampleValue:"READY"
	Tags             string `url:"tags,omitempty"`              // Description:"Comma-separated list of tags. Returned rules match any of the tags (OR operator)",ExampleValue:"security,java8"
	TargetKey        string `url:"targetKey,omitempty"`         // Description:"Quality Profile key on which the rule deactivation is done. To retrieve a profile key please see <code>api/qualityprofiles/search</code>",ExampleValue:"AU-TpxcA-iU5OvuD2FL1"
	TemplateKey      string `url:"template_key,omitempty"`      // Description:"Key of the template rule to filter on. Used to search for the custom rules based on this template.",ExampleValue:"java:S001"
	Types            string `url:"types,omitempty"`             // Description:"Comma-separated list of types. Returned rules match any of the tags (OR operator)",ExampleValue:"BUG"
}

// Deactivate_rules Bulk deactivate rules on Quality profiles.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) Deactivate_rules(opt *QualityprofilesDeactivateRulesOption) (Resp *QualityprofilesDeactivateRulesResp, err error) {
	err := s.ValidateDeactivate_rulesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/deactivate_rules", Opt)
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
type QualityprofilesDeleteResp struct{}
type QualityprofilesDeleteOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Delete Delete a quality profile and all its descendants. The default quality profile cannot be deleted.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) Delete(opt *QualityprofilesDeleteOption) (Resp *QualityprofilesDeleteResp, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/delete", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Export> func manually and complete the fields of this struct
type QualityprofilesExportResp struct{}
type QualityprofilesExportOption struct {
	ExporterKey    string `url:"exporterKey,omitempty"`    // Description:"Output format. If left empty, the same format as api/qualityprofiles/backup is used. Possible values are described by api/qualityprofiles/exporters.",ExampleValue:""
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:"py"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name to export. If left empty, the default profile for the language is exported.",ExampleValue:"My Sonar way"
}

// Export Export a quality profile.
func (s *QualityprofilesService) Export(opt *QualityprofilesExportOption) (Resp *QualityprofilesExportResp, err error) {
	err := s.ValidateExportOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/export", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Exporters> func manually and complete the fields of this struct
type QualityprofilesExportersResp struct{}

// Exporters Lists available profile export formats.
func (s *QualityprofilesService) Exporters() (Resp *QualityprofilesExportersResp, err error) {
	req, err := s.client.NewRequest("GET", "qualityprofiles/exporters", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Importers> func manually and complete the fields of this struct
type QualityprofilesImportersResp struct{}

// Importers List supported importers.
func (s *QualityprofilesService) Importers() (Resp *QualityprofilesImportersResp, err error) {
	req, err := s.client.NewRequest("GET", "qualityprofiles/importers", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Inheritance> func manually and complete the fields of this struct
type QualityprofilesInheritanceResp struct{}
type QualityprofilesInheritanceOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Inheritance Show a quality profile's ancestors and children.
func (s *QualityprofilesService) Inheritance(opt *QualityprofilesInheritanceOption) (Resp *QualityprofilesInheritanceResp, err error) {
	err := s.ValidateInheritanceOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/inheritance", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Projects> func manually and complete the fields of this struct
type QualityprofilesProjectsResp struct{}
type QualityprofilesProjectsOption struct {
	Key      string `url:"key,omitempty"`      // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	P        string `url:"p,omitempty"`        // Description:"1-based page number",ExampleValue:"42"
	Ps       string `url:"ps,omitempty"`       // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q        string `url:"q,omitempty"`        // Description:"Limit search to projects that contain the supplied string.",ExampleValue:"sonar"
	Selected string `url:"selected,omitempty"` // Description:"Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).",ExampleValue:""
}

// Projects List projects with their association status regarding a quality profile
func (s *QualityprofilesService) Projects(opt *QualityprofilesProjectsOption) (Resp *QualityprofilesProjectsResp, err error) {
	err := s.ValidateProjectsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/projects", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <RemoveProject> func manually and complete the fields of this struct
type QualityprofilesRemoveProjectResp struct{}
type QualityprofilesRemoveProjectOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	Project        string `url:"project,omitempty"`        // Description:"Project key",ExampleValue:"my_project"
	ProjectUuid    string `url:"projectUuid,omitempty"`    // Description:"Project ID. Either this parameter, or 'project' must be set.",ExampleValue:"AU-TpxcB-iU5OvuD2FL6"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Remove_project Remove a project's association with a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li>  <li>Administer right on the specified project</li></ul>
func (s *QualityprofilesService) Remove_project(opt *QualityprofilesRemoveProjectOption) (Resp *QualityprofilesRemoveProjectResp, err error) {
	err := s.ValidateRemove_projectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/remove_project", Opt)
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
type QualityprofilesRenameResp struct{}
type QualityprofilesRenameOption struct {
	Key  string `url:"key,omitempty"`  // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Name string `url:"name,omitempty"` // Description:"New quality profile name",ExampleValue:"My Sonar way"
}

// Rename Rename a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) Rename(opt *QualityprofilesRenameOption) (Resp *QualityprofilesRenameResp, err error) {
	err := s.ValidateRenameOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/rename", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Restore> func manually and complete the fields of this struct
type QualityprofilesRestoreResp struct{}
type QualityprofilesRestoreOption struct {
	Backup string `url:"backup,omitempty"` // Description:"A profile backup file in XML format, as generated by api/qualityprofiles/backup or the former api/profiles/backup.",ExampleValue:""
}

// Restore Restore a quality profile using an XML file. The restored profile name is taken from the backup file, so if a profile with the same name and language already exists, it will be overwritten.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) Restore(opt *QualityprofilesRestoreOption) (Resp *QualityprofilesRestoreResp, err error) {
	err := s.ValidateRestoreOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/restore", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <RestoreBuiltIn> func manually and complete the fields of this struct
type QualityprofilesRestoreBuiltInResp struct{}

// Restore_built_in This web service has no effect since 6.4. It's no more possible to restore built-in quality profiles because they are automatically updated and read only. Returns HTTP code 410.
func (s *QualityprofilesService) Restore_built_in() (Resp *QualityprofilesRestoreBuiltInResp, err error) {
	req, err := s.client.NewRequest("POST", "qualityprofiles/restore_built_in", Opt)
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
type QualityprofilesSearchResp struct{}
type QualityprofilesSearchOption struct {
	Defaults       string `url:"defaults,omitempty"`       // Description:"If set to true, return only the quality profiles marked as default for each language",ExampleValue:""
	Language       string `url:"language,omitempty"`       // Description:"Language key. If provided, only profiles for the given language are returned.",ExampleValue:""
	Project        string `url:"project,omitempty"`        // Description:"Project key",ExampleValue:"my_project"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"SonarQube Way"
}

// Search Search quality profiles
func (s *QualityprofilesService) Search(opt *QualityprofilesSearchOption) (Resp *QualityprofilesSearchResp, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/search", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <SetDefault> func manually and complete the fields of this struct
type QualityprofilesSetDefaultResp struct{}
type QualityprofilesSetDefaultOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Set_default Select the default profile for a given language.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) Set_default(opt *QualityprofilesSetDefaultOption) (Resp *QualityprofilesSetDefaultResp, err error) {
	err := s.ValidateSet_defaultOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/set_default", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
