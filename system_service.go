// Get system details, and perform some management actions, such as restarting, and initiating a database migration (as part of a system upgrade).
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type SystemService struct {
	client *Client
}

// [TODO] you should call the <ChangeLogLevel> func manually and complete the fields of this struct
type SystemChangeLogLevelResp struct{}
type SystemChangeLogLevelOption struct {
	Level string `url:"level,omitempty"` // Description:"The new level. Be cautious: DEBUG, and even more TRACE, may have performance impacts.",ExampleValue:""
}

// Change_log_level Temporarily changes level of logs. New level is not persistent and is lost when restarting server. Requires system administration permission.
func (s *SystemService) Change_log_level(opt *SystemChangeLogLevelOption) (Resp *SystemChangeLogLevelResp, err error) {
	err := s.ValidateChange_log_levelOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "system/change_log_level", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <DbMigrationStatus> func manually and complete the fields of this struct
type SystemDbMigrationStatusResp struct{}

// Db_migration_status Display the database migration status of SonarQube.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>
func (s *SystemService) Db_migration_status() (Resp *SystemDbMigrationStatusResp, err error) {
	req, err := s.client.NewRequest("GET", "system/db_migration_status", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Health> func manually and complete the fields of this struct
type SystemHealthResp struct{}

// Health Provide health status of SonarQube.<p>Require 'Administer System' permission or authentication with passcode</p><p>  <ul> <li>GREEN: SonarQube is fully operational</li> <li>YELLOW: SonarQube is usable, but it needs attention in order to be fully operational</li> <li>RED: SonarQube is not operational</li> </ul></p>
func (s *SystemService) Health() (Resp *SystemHealthResp, err error) {
	req, err := s.client.NewRequest("GET", "system/health", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Logs> func manually and complete the fields of this struct
type SystemLogsResp struct{}
type SystemLogsOption struct {
	Process string `url:"process,omitempty"` // Description:"Process to get logs from",ExampleValue:""
}

// Logs Get system logs in plain-text format. Requires system administration permission.
func (s *SystemService) Logs(opt *SystemLogsOption) (Resp *SystemLogsResp, err error) {
	err := s.ValidateLogsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "system/logs", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <MigrateDb> func manually and complete the fields of this struct
type SystemMigrateDbResp struct{}

// Migrate_db Migrate the database to match the current version of SonarQube.<br/>Sending a POST request to this URL starts the DB migration. It is strongly advised to <strong>make a database backup</strong> before invoking this WS.<br/>State values are:<ul><li>NO_MIGRATION: DB is up to date with current version of SonarQube.</li><li>NOT_SUPPORTED: Migration is not supported on embedded databases.</li><li>MIGRATION_RUNNING: DB migration is under go.</li><li>MIGRATION_SUCCEEDED: DB migration has run and has been successful.</li><li>MIGRATION_FAILED: DB migration has run and failed. SonarQube must be restarted in order to retry a DB migration (optionally after DB has been restored from backup).</li><li>MIGRATION_REQUIRED: DB migration is required.</li></ul>
func (s *SystemService) Migrate_db() (Resp *SystemMigrateDbResp, err error) {
	req, err := s.client.NewRequest("POST", "system/migrate_db", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Ping> func manually and complete the fields of this struct
type SystemPingResp struct{}

// Ping Answers "pong" as plain-text
func (s *SystemService) Ping() (Resp *SystemPingResp, err error) {
	req, err := s.client.NewRequest("GET", "system/ping", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Restart> func manually and complete the fields of this struct
type SystemRestartResp struct{}

// Restart Restart server. Require 'Administer System' permission. Perform a full restart of the Web, Search and Compute Engine Servers processes.
func (s *SystemService) Restart() (Resp *SystemRestartResp, err error) {
	req, err := s.client.NewRequest("POST", "system/restart", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Status> func manually and complete the fields of this struct
type SystemStatusResp struct{}

// Status Get state information about SonarQube.<p>status: the running status <ul> <li>STARTING: SonarQube Web Server is up and serving some Web Services (eg. api/system/status) but initialization is still ongoing</li> <li>UP: SonarQube instance is up and running</li> <li>DOWN: SonarQube instance is up but not running because migration has failed (refer to WS /api/system/migrate_db for details) or some other reason (check logs).</li> <li>RESTARTING: SonarQube instance is still up but a restart has been requested (refer to WS /api/system/restart for details).</li> <li>DB_MIGRATION_NEEDED: database migration is required. DB migration can be started using WS /api/system/migrate_db.</li> <li>DB_MIGRATION_RUNNING: DB migration is running (refer to WS /api/system/migrate_db for details)</li> </ul></p>
func (s *SystemService) Status() (Resp *SystemStatusResp, err error) {
	req, err := s.client.NewRequest("GET", "system/status", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Upgrades> func manually and complete the fields of this struct
type SystemUpgradesResp struct{}

// Upgrades Lists available upgrades for the SonarQube instance (if any) and for each one, lists incompatible plugins and plugins requiring upgrade.<br/>Plugin information is retrieved from Update Center. Date and time at which Update Center was last refreshed is provided in the response.
func (s *SystemService) Upgrades() (Resp *SystemUpgradesResp, err error) {
	req, err := s.client.NewRequest("GET", "system/upgrades", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
