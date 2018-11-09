// Manage settings.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type SettingsService struct {
	client *Client
}

type Settings struct {
	Definitions []struct {
		Category     string `json:"category,omitempty"`
		DefaultValue string `json:"defaultValue,omitempty"`
		Description  string `json:"description,omitempty"`
		Fields       []struct {
			Description string   `json:"description,omitempty"`
			Key         string   `json:"key,omitempty"`
			Name        string   `json:"name,omitempty"`
			Options     []string `json:"options,omitempty"`
			Type        string   `json:"type,omitempty"`
		} `json:"fields,omitempty"`
		Key         string   `json:"key,omitempty"`
		MultiValues bool     `json:"multiValues,omitempty"`
		Name        string   `json:"name,omitempty"`
		Options     []string `json:"options,omitempty"`
		SubCategory string   `json:"subCategory,omitempty"`
		Type        string   `json:"type,omitempty"`
	} `json:"definitions,omitempty"`
	Settings []struct {
		FieldValues []struct {
			Boolean string `json:"boolean,omitempty"`
			Text    string `json:"text,omitempty"`
		} `json:"fieldValues,omitempty"`
		Inherited bool     `json:"inherited,omitempty"`
		Key       string   `json:"key,omitempty"`
		Value     string   `json:"value,omitempty"`
		Values    []string `json:"values,omitempty"`
	} `json:"settings,omitempty"`
}

type SettingsListDefinitionsOption struct {
	Component string `url:"component,omitempty"` // Description:"Component key",ExampleValue:"my_project"
}

// List_definitions List settings definitions.<br>Requires 'Browse' permission when a component is specified<br/>
func (s *SettingsService) ListDefinitions(opt *SettingsListDefinitionsOption) (resp *Settings, err error) {
	err := s.ValidateListDefinitionsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "settings/list_definitions", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type SettingsResetOption struct {
	Component string `url:"component,omitempty"` // Description:"Component key",ExampleValue:"my_project"
	Keys      string `url:"keys,omitempty"`      // Description:"Comma-separated list of keys",ExampleValue:"sonar.links.scm,sonar.debt.hoursInDay"
}

// Reset Remove a setting value.<br>The settings defined in config/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>
func (s *SettingsService) Reset(opt *SettingsResetOption) (resp *string, err error) {
	err := s.ValidateResetOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "settings/reset", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type SettingsSetOption struct {
	Component   string `url:"component,omitempty"`   // Description:"Component key",ExampleValue:"my_project"
	FieldValues string `url:"fieldValues,omitempty"` // Description:"Setting field values. To set several values, the parameter must be called once for each value.",ExampleValue:"fieldValues={"firstField":"first value", "secondField":"second value", "thirdField":"third value"}"
	Key         string `url:"key,omitempty"`         // Description:"Setting key",ExampleValue:"sonar.links.scm"
	Value       string `url:"value,omitempty"`       // Description:"Setting value. To reset a value, please use the reset web service.",ExampleValue:"git@github.com:SonarSource/sonarqube.git"
	Values      string `url:"values,omitempty"`      // Description:"Setting multi value. To set several values, the parameter must be called once for each value.",ExampleValue:"values=firstValue&values=secondValue&values=thirdValue"
}

// Set Update a setting value.<br>Either 'value' or 'values' must be provided.<br> The settings defined in config/sonar.properties are read-only and can't be changed.<br/>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified component</li></ul>
func (s *SettingsService) Set(opt *SettingsSetOption) (resp *string, err error) {
	err := s.ValidateSetOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "settings/set", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type SettingsValuesOption struct {
	Component string `url:"component,omitempty"` // Description:"Component key",ExampleValue:"my_project"
	Keys      string `url:"keys,omitempty"`      // Description:"List of setting keys",ExampleValue:"sonar.test.inclusions,sonar.dbcleaner.cleanDirectory"
}

// Values List settings values.<br>If no value has been set for a setting, then the default value is returned.<br>The settings from conf/sonar.properties are excluded from results.<br>Requires 'Browse' or 'Execute Analysis' permission when a component is specified<br/>
func (s *SettingsService) Values(opt *SettingsValuesOption) (resp *Settings, err error) {
	err := s.ValidateValuesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "settings/values", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
