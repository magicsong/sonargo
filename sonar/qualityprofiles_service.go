// Manage quality profiles.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type QualityprofilesService struct {
	client *Client
}

type Qualityprofiles struct {
	Actions struct {
		Create bool `json:"create,omitempty"`
	} `json:"actions,omitempty"`
	Ancestors []struct {
		ActiveRuleCount int64  `json:"activeRuleCount,omitempty"`
		IsBuiltIn       bool   `json:"isBuiltIn,omitempty"`
		Key             string `json:"key,omitempty"`
		Name            string `json:"name,omitempty"`
		Parent          string `json:"parent,omitempty"`
	} `json:"ancestors,omitempty"`
	Children []struct {
		ActiveRuleCount     int64  `json:"activeRuleCount,omitempty"`
		IsBuiltIn           bool   `json:"isBuiltIn,omitempty"`
		Key                 string `json:"key,omitempty"`
		Name                string `json:"name,omitempty"`
		OverridingRuleCount int64  `json:"overridingRuleCount,omitempty"`
	} `json:"children,omitempty"`
	Events []struct {
		Action      string `json:"action,omitempty"`
		AuthorLogin string `json:"authorLogin,omitempty"`
		AuthorName  string `json:"authorName,omitempty"`
		Date        string `json:"date,omitempty"`
		Params      struct {
			Format   string `json:"format,omitempty"`
			Severity string `json:"severity,omitempty"`
		} `json:"params,omitempty"`
		RuleKey  string `json:"ruleKey,omitempty"`
		RuleName string `json:"ruleName,omitempty"`
	} `json:"events,omitempty"`
	Exporters []struct {
		Key       string   `json:"key,omitempty"`
		Languages []string `json:"languages,omitempty"`
		Name      string   `json:"name,omitempty"`
	} `json:"exporters,omitempty"`
	Importers []struct {
		Key       string   `json:"key,omitempty"`
		Languages []string `json:"languages,omitempty"`
		Name      string   `json:"name,omitempty"`
	} `json:"importers,omitempty"`
	IsDefault   bool   `json:"isDefault,omitempty"`
	IsInherited bool   `json:"isInherited,omitempty"`
	Key         string `json:"key,omitempty"`
	Language    string `json:"language,omitempty"`
	More        bool   `json:"more,omitempty"`
	Name        string `json:"name,omitempty"`
	P           int64  `json:"p,omitempty"`
	ParentKey   string `json:"parentKey,omitempty"`
	Profile     struct {
		ActiveRuleCount     int64  `json:"activeRuleCount,omitempty"`
		IsBuiltIn           bool   `json:"isBuiltIn,omitempty"`
		IsDefault           bool   `json:"isDefault,omitempty"`
		IsInherited         bool   `json:"isInherited,omitempty"`
		Key                 string `json:"key,omitempty"`
		Language            string `json:"language,omitempty"`
		LanguageName        string `json:"languageName,omitempty"`
		Name                string `json:"name,omitempty"`
		OverridingRuleCount int64  `json:"overridingRuleCount,omitempty"`
		Parent              string `json:"parent,omitempty"`
	} `json:"profile,omitempty"`
	Profiles []struct {
		Actions struct {
			AssociateProjects bool `json:"associateProjects,omitempty"`
			Copy              bool `json:"copy,omitempty"`
			Delete            bool `json:"delete,omitempty"`
			Edit              bool `json:"edit,omitempty"`
			SetAsDefault      bool `json:"setAsDefault,omitempty"`
		} `json:"actions,omitempty"`
		ActiveDeprecatedRuleCount int64  `json:"activeDeprecatedRuleCount,omitempty"`
		ActiveRuleCount           int64  `json:"activeRuleCount,omitempty"`
		IsBuiltIn                 bool   `json:"isBuiltIn,omitempty"`
		IsDefault                 bool   `json:"isDefault,omitempty"`
		IsInherited               bool   `json:"isInherited,omitempty"`
		Key                       string `json:"key,omitempty"`
		Language                  string `json:"language,omitempty"`
		LanguageName              string `json:"languageName,omitempty"`
		LastUsed                  string `json:"lastUsed,omitempty"`
		Name                      string `json:"name,omitempty"`
		ParentKey                 string `json:"parentKey,omitempty"`
		ParentName                string `json:"parentName,omitempty"`
		ProjectCount              int64  `json:"projectCount,omitempty"`
		RuleUpdatedAt             string `json:"ruleUpdatedAt,omitempty"`
		UserUpdatedAt             string `json:"userUpdatedAt,omitempty"`
	} `json:"profiles,omitempty"`
	Ps      int64 `json:"ps,omitempty"`
	Results []struct {
		ID       string `json:"id,omitempty"`
		Key      string `json:"key,omitempty"`
		Name     string `json:"name,omitempty"`
		Selected bool   `json:"selected,omitempty"`
	} `json:"results,omitempty"`
	Total    int64    `json:"total,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
}

type QualityprofilesActivateRuleOption struct {
	Key      string `url:"key,omitempty"`      // Description:"Quality Profile key. Can be obtained through <code>api/qualityprofiles/search</code>",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Params   string `url:"params,omitempty"`   // Description:"Parameters as semi-colon list of <code>key=value</code>. Ignored if parameter reset is true.",ExampleValue:"params=key1=v1;key2=v2"
	Reset    string `url:"reset,omitempty"`    // Description:"Reset severity and parameters of activated rule. Set the values defined on parent profile or from rule default values.",ExampleValue:""
	Rule     string `url:"rule,omitempty"`     // Description:"Rule key",ExampleValue:"squid:AvoidCycles"
	Severity string `url:"severity,omitempty"` // Description:"Severity. Ignored if parameter reset is true.",ExampleValue:""
}

// Activate_rule Activate a rule on a Quality Profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) ActivateRule(opt *QualityprofilesActivateRuleOption) (resp *string, err error) {
	err := s.ValidateActivateRuleOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/activate_rule", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

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
func (s *QualityprofilesService) ActivateRules(opt *QualityprofilesActivateRulesOption) (resp *string, err error) {
	err := s.ValidateActivateRulesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/activate_rules", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesAddProjectOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	Project        string `url:"project,omitempty"`        // Description:"Project key",ExampleValue:"my_project"
	ProjectUuid    string `url:"projectUuid,omitempty"`    // Description:"Project ID. Either this parameter or 'project' must be set.",ExampleValue:"AU-TpxcA-iU5OvuD2FL5"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Add_project Associate a project with a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li>  <li>Administer right on the specified project</li></ul>
func (s *QualityprofilesService) AddProject(opt *QualityprofilesAddProjectOption) (resp *string, err error) {
	err := s.ValidateAddProjectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/add_project", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesBackupOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Backup Backup a quality profile in XML form. The exported profile can be restored through api/qualityprofiles/restore.
func (s *QualityprofilesService) Backup(opt *QualityprofilesBackupOption) (resp *Qualityprofiles, err error) {
	err := s.ValidateBackupOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/backup", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesChangeParentOption struct {
	Key                  string `url:"key,omitempty"`                  // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language             string `url:"language,omitempty"`             // Description:"Quality profile language",ExampleValue:""
	ParentKey            string `url:"parentKey,omitempty"`            // Description:"New parent profile key.<br> If no profile is provided, the inheritance link with current parent profile (if any) is broken, which deactivates all rules which come from the parent and are not overridden.",ExampleValue:"AU-TpxcA-iU5OvuD2FLz"
	ParentQualityProfile string `url:"parentQualityProfile,omitempty"` // Description:"Quality profile name. If this parameter is set, 'parentKey' must not be set and 'language' must be set to disambiguate.",ExampleValue:"Sonar way"
	QualityProfile       string `url:"qualityProfile,omitempty"`       // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Change_parent Change a quality profile's parent.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) ChangeParent(opt *QualityprofilesChangeParentOption) (resp *string, err error) {
	err := s.ValidateChangeParentOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/change_parent", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

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
func (s *QualityprofilesService) Changelog(opt *QualityprofilesChangelogOption) (resp *Qualityprofiles, err error) {
	err := s.ValidateChangelogOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/changelog", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesCopyOption struct {
	FromKey string `url:"fromKey,omitempty"` // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ToName  string `url:"toName,omitempty"`  // Description:"Name for the new quality profile.",ExampleValue:"My Sonar way"
}

// Copy Copy a quality profile.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) Copy(opt *QualityprofilesCopyOption) (resp *Qualityprofiles, err error) {
	err := s.ValidateCopyOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/copy", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesCreateOption struct {
	BackupSonarlintVsCsFake string `url:"backup_sonarlint-vs-cs-fake,omitempty"` // Description:"A configuration file for Technical importer for the MSBuild SonarQube Scanner.",ExampleValue:""
	Language                string `url:"language,omitempty"`                    // Description:"Quality profile language",ExampleValue:"js"
	Name                    string `url:"name,omitempty"`                        // Description:"Quality profile name",ExampleValue:"My Sonar way"
}

// Create Create a quality profile.<br>Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) Create(opt *QualityprofilesCreateOption) (resp *Qualityprofiles, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/create", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesDeactivateRuleOption struct {
	Key  string `url:"key,omitempty"`  // Description:"Quality Profile key. Can be obtained through <code>api/qualityprofiles/search</code>",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Rule string `url:"rule,omitempty"` // Description:"Rule key",ExampleValue:"squid:AvoidCycles"
}

// Deactivate_rule Deactivate a rule on a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) DeactivateRule(opt *QualityprofilesDeactivateRuleOption) (resp *string, err error) {
	err := s.ValidateDeactivateRuleOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/deactivate_rule", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

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
func (s *QualityprofilesService) DeactivateRules(opt *QualityprofilesDeactivateRulesOption) (resp *string, err error) {
	err := s.ValidateDeactivateRulesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/deactivate_rules", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesDeleteOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Delete Delete a quality profile and all its descendants. The default quality profile cannot be deleted.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) Delete(opt *QualityprofilesDeleteOption) (resp *string, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/delete", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesExportOption struct {
	ExporterKey    string `url:"exporterKey,omitempty"`    // Description:"Output format. If left empty, the same format as api/qualityprofiles/backup is used. Possible values are described by api/qualityprofiles/exporters.",ExampleValue:""
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:"py"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name to export. If left empty, the default profile for the language is exported.",ExampleValue:"My Sonar way"
}

// Export Export a quality profile.
func (s *QualityprofilesService) Export(opt *QualityprofilesExportOption) (resp *Qualityprofiles, err error) {
	err := s.ValidateExportOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/export", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Exporters Lists available profile export formats.
func (s *QualityprofilesService) Exporters() (resp *Qualityprofiles, err error) {
	req, err := s.client.NewRequest("GET", "qualityprofiles/exporters", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Importers List supported importers.
func (s *QualityprofilesService) Importers() (resp *Qualityprofiles, err error) {
	req, err := s.client.NewRequest("GET", "qualityprofiles/importers", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesInheritanceOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Inheritance Show a quality profile's ancestors and children.
func (s *QualityprofilesService) Inheritance(opt *QualityprofilesInheritanceOption) (resp *Qualityprofiles, err error) {
	err := s.ValidateInheritanceOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/inheritance", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesProjectsOption struct {
	Key      string `url:"key,omitempty"`      // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	P        string `url:"p,omitempty"`        // Description:"1-based page number",ExampleValue:"42"
	Ps       string `url:"ps,omitempty"`       // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q        string `url:"q,omitempty"`        // Description:"Limit search to projects that contain the supplied string.",ExampleValue:"sonar"
	Selected string `url:"selected,omitempty"` // Description:"Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).",ExampleValue:""
}

// Projects List projects with their association status regarding a quality profile
func (s *QualityprofilesService) Projects(opt *QualityprofilesProjectsOption) (resp *Qualityprofiles, err error) {
	err := s.ValidateProjectsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/projects", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesRemoveProjectOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	Project        string `url:"project,omitempty"`        // Description:"Project key",ExampleValue:"my_project"
	ProjectUuid    string `url:"projectUuid,omitempty"`    // Description:"Project ID. Either this parameter, or 'project' must be set.",ExampleValue:"AU-TpxcB-iU5OvuD2FL6"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Remove_project Remove a project's association with a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li>  <li>Administer right on the specified project</li></ul>
func (s *QualityprofilesService) RemoveProject(opt *QualityprofilesRemoveProjectOption) (resp *string, err error) {
	err := s.ValidateRemoveProjectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/remove_project", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesRenameOption struct {
	Key  string `url:"key,omitempty"`  // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Name string `url:"name,omitempty"` // Description:"New quality profile name",ExampleValue:"My Sonar way"
}

// Rename Rename a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) Rename(opt *QualityprofilesRenameOption) (resp *string, err error) {
	err := s.ValidateRenameOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/rename", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesRestoreOption struct {
	Backup string `url:"backup,omitempty"` // Description:"A profile backup file in XML format, as generated by api/qualityprofiles/backup or the former api/profiles/backup.",ExampleValue:""
}

// Restore Restore a quality profile using an XML file. The restored profile name is taken from the backup file, so if a profile with the same name and language already exists, it will be overwritten.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) Restore(opt *QualityprofilesRestoreOption) (resp *string, err error) {
	err := s.ValidateRestoreOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/restore", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Restore_built_in This web service has no effect since 6.4. It's no more possible to restore built-in quality profiles because they are automatically updated and read only. Returns HTTP code 410.
func (s *QualityprofilesService) RestoreBuiltIn() (resp *string, err error) {
	req, err := s.client.NewRequest("POST", "qualityprofiles/restore_built_in", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesSearchOption struct {
	Defaults       string `url:"defaults,omitempty"`       // Description:"If set to true, return only the quality profiles marked as default for each language",ExampleValue:""
	Language       string `url:"language,omitempty"`       // Description:"Language key. If provided, only profiles for the given language are returned.",ExampleValue:""
	Project        string `url:"project,omitempty"`        // Description:"Project key",ExampleValue:"my_project"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"SonarQube Way"
}

// Search Search quality profiles
func (s *QualityprofilesService) Search(opt *QualityprofilesSearchOption) (resp *Qualityprofiles, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/search", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type QualityprofilesSetDefaultOption struct {
	Key            string `url:"key,omitempty"`            // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"Sonar way"
}

// Set_default Select the default profile for a given language.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) SetDefault(opt *QualityprofilesSetDefaultOption) (resp *string, err error) {
	err := s.ValidateSetDefaultOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/set_default", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
