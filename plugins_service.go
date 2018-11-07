// Manage the plugins on the server, including installing, uninstalling, and upgrading.
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type PluginsService struct {
	client *Client
}

// [TODO] you should call the <Available> func manually and complete the fields of this struct
type PluginsAvailableResp struct{}

// Available Get the list of all the plugins available for installation on the SonarQube instance, sorted by plugin name.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.<br/>Update status values are: <ul><li>COMPATIBLE: plugin is compatible with current SonarQube instance.</li><li>INCOMPATIBLE: plugin is not compatible with current SonarQube instance.</li><li>REQUIRES_SYSTEM_UPGRADE: plugin requires SonarQube to be upgraded before being installed.</li><li>DEPS_REQUIRE_SYSTEM_UPGRADE: at least one plugin on which the plugin is dependent requires SonarQube to be upgraded.</li></ul>Require 'Administer System' permission.
func (s *PluginsService) Available() (Resp *PluginsAvailableResp, err error) {
	req, err := s.client.NewRequest("GET", "plugins/available", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <CancelAll> func manually and complete the fields of this struct
type PluginsCancelAllResp struct{}

// Cancel_all Cancels any operation pending on any plugin (install, update or uninstall)<br/>Requires user to be authenticated with Administer System permissions
func (s *PluginsService) Cancel_all() (Resp *PluginsCancelAllResp, err error) {
	req, err := s.client.NewRequest("POST", "plugins/cancel_all", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Install> func manually and complete the fields of this struct
type PluginsInstallResp struct{}
type PluginsInstallOption struct {
	Key string `url:"key,omitempty"` // Description:"The key identifying the plugin to install",ExampleValue:""
}

// Install Installs the latest version of a plugin specified by its key.<br/>Plugin information is retrieved from Update Center.<br/>Requires user to be authenticated with Administer System permissions
func (s *PluginsService) Install(opt *PluginsInstallOption) (Resp *PluginsInstallResp, err error) {
	err := s.ValidateInstallOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "plugins/install", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Installed> func manually and complete the fields of this struct
type PluginsInstalledResp struct{}
type PluginsInstalledOption struct {
	F string `url:"f,omitempty"` // Description:"Comma-separated list of the additional fields to be returned in response. No additional field is returned by default. Possible values are:<ul><li>category - category as defined in the Update Center. A connection to the Update Center is needed</li></lu>",ExampleValue:""
}

// Installed Get the list of all the plugins installed on the SonarQube instance, sorted by plugin name.
func (s *PluginsService) Installed(opt *PluginsInstalledOption) (Resp *PluginsInstalledResp, err error) {
	err := s.ValidateInstalledOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "plugins/installed", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Pending> func manually and complete the fields of this struct
type PluginsPendingResp struct{}

// Pending Get the list of plugins which will either be installed or removed at the next startup of the SonarQube instance, sorted by plugin name.<br/>Require 'Administer System' permission.
func (s *PluginsService) Pending() (Resp *PluginsPendingResp, err error) {
	req, err := s.client.NewRequest("GET", "plugins/pending", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Uninstall> func manually and complete the fields of this struct
type PluginsUninstallResp struct{}
type PluginsUninstallOption struct {
	Key string `url:"key,omitempty"` // Description:"The key identifying the plugin to uninstall",ExampleValue:""
}

// Uninstall Uninstalls the plugin specified by its key.<br/>Requires user to be authenticated with Administer System permissions.
func (s *PluginsService) Uninstall(opt *PluginsUninstallOption) (Resp *PluginsUninstallResp, err error) {
	err := s.ValidateUninstallOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "plugins/uninstall", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Update> func manually and complete the fields of this struct
type PluginsUpdateResp struct{}
type PluginsUpdateOption struct {
	Key string `url:"key,omitempty"` // Description:"The key identifying the plugin to update",ExampleValue:""
}

// Update Updates a plugin specified by its key to the latest version compatible with the SonarQube instance.<br/>Plugin information is retrieved from Update Center.<br/>Requires user to be authenticated with Administer System permissions
func (s *PluginsService) Update(opt *PluginsUpdateOption) (Resp *PluginsUpdateResp, err error) {
	err := s.ValidateUpdateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "plugins/update", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Updates> func manually and complete the fields of this struct
type PluginsUpdatesResp struct{}

// Updates Lists plugins installed on the SonarQube instance for which at least one newer version is available, sorted by plugin name.<br/>Each newer version is listed, ordered from the oldest to the newest, with its own update/compatibility status.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.<br/>Update status values are: [COMPATIBLE, INCOMPATIBLE, REQUIRES_UPGRADE, DEPS_REQUIRE_UPGRADE].<br/>Require 'Administer System' permission.
func (s *PluginsService) Updates() (Resp *PluginsUpdatesResp, err error) {
	req, err := s.client.NewRequest("GET", "plugins/updates", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
