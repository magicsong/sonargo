// Get system details, and perform some management actions, such as restarting, and initiating a database migration (as part of a system upgrade).
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type SystemService struct {
	client *Client
}

type System struct {
	Causes []struct {
		Message string `json:"message,omitempty"`
	} `json:"causes,omitempty"`
	Health  string `json:"health,omitempty"`
	ID      string `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
	Nodes   []struct {
		Causes []struct {
			Message string `json:"message,omitempty"`
		} `json:"causes,omitempty"`
		Health    string `json:"health,omitempty"`
		Host      string `json:"host,omitempty"`
		Name      string `json:"name,omitempty"`
		Port      int64  `json:"port,omitempty"`
		StartedAt string `json:"startedAt,omitempty"`
		Type      string `json:"type,omitempty"`
	} `json:"nodes,omitempty"`
	StartedAt           string `json:"startedAt,omitempty"`
	State               string `json:"state,omitempty"`
	Status              string `json:"status,omitempty"`
	UpdateCenterRefresh string `json:"updateCenterRefresh,omitempty"`
	Upgrades            []struct {
		ChangeLogURL string `json:"changeLogUrl,omitempty"`
		Description  string `json:"description,omitempty"`
		DownloadURL  string `json:"downloadUrl,omitempty"`
		Plugins      struct {
			Incompatible []struct {
				Category         string `json:"category,omitempty"`
				Description      string `json:"description,omitempty"`
				EditionBundled   bool   `json:"editionBundled,omitempty"`
				Key              string `json:"key,omitempty"`
				License          string `json:"license,omitempty"`
				Name             string `json:"name,omitempty"`
				OrganizationName string `json:"organizationName,omitempty"`
				OrganizationURL  string `json:"organizationUrl,omitempty"`
			} `json:"incompatible,omitempty"`
			RequireUpdate []struct {
				Category              string `json:"category,omitempty"`
				Description           string `json:"description,omitempty"`
				EditionBundled        bool   `json:"editionBundled,omitempty"`
				Key                   string `json:"key,omitempty"`
				License               string `json:"license,omitempty"`
				Name                  string `json:"name,omitempty"`
				OrganizationName      string `json:"organizationName,omitempty"`
				OrganizationURL       string `json:"organizationUrl,omitempty"`
				TermsAndConditionsURL string `json:"termsAndConditionsUrl,omitempty"`
				Version               string `json:"version,omitempty"`
			} `json:"requireUpdate,omitempty"`
		} `json:"plugins,omitempty"`
		ReleaseDate string `json:"releaseDate,omitempty"`
		Version     string `json:"version,omitempty"`
	} `json:"upgrades,omitempty"`
	Version string `json:"version,omitempty"`
}

type SystemChangeLogLevelOption struct {
	Level string `url:"level,omitempty"` // Description:"The new level. Be cautious: DEBUG, and even more TRACE, may have performance impacts.",ExampleValue:""
}

// Change_log_level Temporarily changes level of logs. New level is not persistent and is lost when restarting server. Requires system administration permission.
func (s *SystemService) ChangeLogLevel(opt *SystemChangeLogLevelOption) (resp *string, err error) {
	err := s.ValidateChangeLogLevelOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "system/change_log_level", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Db_migration_status Display the database migration status of SonarQube.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>
func (s *SystemService) DbMigrationStatus() (resp *System, err error) {
	req, err := s.client.NewRequest("GET", "system/db_migration_status", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Health Provide health status of SonarQube.<p>Require 'Administer System' permission or authentication with passcode</p><p>  <ul> <li>GREEN: SonarQube is fully operational</li> <li>YELLOW: SonarQube is usable, but it needs attention in order to be fully operational</li> <li>RED: SonarQube is not operational</li> </ul></p>
func (s *SystemService) Health() (resp *System, err error) {
	req, err := s.client.NewRequest("GET", "system/health", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type SystemLogsOption struct {
	Process string `url:"process,omitempty"` // Description:"Process to get logs from",ExampleValue:""
}

// Logs Get system logs in plain-text format. Requires system administration permission.
func (s *SystemService) Logs(opt *SystemLogsOption) (resp *string, err error) {
	err := s.ValidateLogsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "system/logs", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Migrate_db Migrate the database to match the current version of SonarQube.<br/>Sending a POST request to this URL starts the DB migration. It is strongly advised to <strong>make a database backup</strong> before invoking this WS.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>
func (s *SystemService) MigrateDb() (resp *System, err error) {
	req, err := s.client.NewRequest("POST", "system/migrate_db", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Ping Answers "pong" as plain-text
func (s *SystemService) Ping() (resp *string, err error) {
	req, err := s.client.NewRequest("GET", "system/ping", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Restart Restart server. Require 'Administer System' permission. Perform a full restart of the Web, Search and Compute Engine Servers processes.
func (s *SystemService) Restart() (resp *string, err error) {
	req, err := s.client.NewRequest("POST", "system/restart", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Status Get state information about SonarQube.<p>status: the running status <ul> <li>STARTING: SonarQube Web Server is up and serving some Web Services (eg. api/system/status) but initialization is still ongoing</li> <li>UP: SonarQube instance is up and running</li> <li>DOWN: SonarQube instance is up but not running because migration has failed (refer to WS /api/system/migrate_db for details) or some other reason (check logs).</li> <li>RESTARTING: SonarQube instance is still up but a restart has been requested (refer to WS /api/system/restart for details).</li> <li>DB_MIGRATION_NEEDED: database migration is required. DB migration can be started using WS /api/system/migrate_db.</li> <li>DB_MIGRATION_RUNNING: DB migration is running (refer to WS /api/system/migrate_db for details)</li> </ul></p>
func (s *SystemService) Status() (resp *System, err error) {
	req, err := s.client.NewRequest("GET", "system/status", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

// Upgrades Lists available upgrades for the SonarQube instance (if any) and for each one, lists incompatible plugins and plugins requiring upgrade.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.
func (s *SystemService) Upgrades() (resp *System, err error) {
	req, err := s.client.NewRequest("GET", "system/upgrades", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
