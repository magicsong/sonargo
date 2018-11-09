package sonargo

import (
	"net/http"
	"net/url"
)

type Client struct {
	baseURL                   *url.URL
	username, password, token string
	authType                  authType
	httpClient                *http.Client
	Authentication            *AuthenticationService
	Ce                        *CeService
	Components                *ComponentsService
	CustomMeasures            *CustomMeasuresService
	Duplications              *DuplicationsService
	Editions                  *EditionsService
	Favorites                 *FavoritesService
	Issues                    *IssuesService
	Languages                 *LanguagesService
	Measures                  *MeasuresService
	Metrics                   *MetricsService
	Notifications             *NotificationsService
	Permissions               *PermissionsService
	Plugins                   *PluginsService
	ProjectAnalyses           *ProjectAnalysesService
	ProjectBadges             *ProjectBadgesService
	ProjectBranches           *ProjectBranchesService
	ProjectLinks              *ProjectLinksService
	ProjectPullRequests       *ProjectPullRequestsService
	ProjectTags               *ProjectTagsService
	Projects                  *ProjectsService
	Qualitygates              *QualitygatesService
	Qualityprofiles           *QualityprofilesService
	Rules                     *RulesService
	Server                    *ServerService
	Settings                  *SettingsService
	Sources                   *SourcesService
	System                    *SystemService
	UserGroups                *UserGroupsService
	UserTokens                *UserTokensService
	Users                     *UsersService
	Webhooks                  *WebhooksService
	Webservices               *WebservicesService
}

func NewClient(endpoint, username, password string) (*Client, error) {
	c := &Client{username: username, password: password, authType: basicAuth, httpClient: http.DefaultClient}
	if endpoint == "" {
		c.SetBaseURL(defaultBaseURL)
	} else {
		if err := c.SetBaseURL(endpoint); err != nil {
			return nil, err
		}
	}
	c.Authentication = &AuthenticationService{client: c}
	c.Ce = &CeService{client: c}
	c.Components = &ComponentsService{client: c}
	c.CustomMeasures = &CustomMeasuresService{client: c}
	c.Duplications = &DuplicationsService{client: c}
	c.Editions = &EditionsService{client: c}
	c.Favorites = &FavoritesService{client: c}
	c.Issues = &IssuesService{client: c}
	c.Languages = &LanguagesService{client: c}
	c.Measures = &MeasuresService{client: c}
	c.Metrics = &MetricsService{client: c}
	c.Notifications = &NotificationsService{client: c}
	c.Permissions = &PermissionsService{client: c}
	c.Plugins = &PluginsService{client: c}
	c.ProjectAnalyses = &ProjectAnalysesService{client: c}
	c.ProjectBadges = &ProjectBadgesService{client: c}
	c.ProjectBranches = &ProjectBranchesService{client: c}
	c.ProjectLinks = &ProjectLinksService{client: c}
	c.ProjectPullRequests = &ProjectPullRequestsService{client: c}
	c.ProjectTags = &ProjectTagsService{client: c}
	c.Projects = &ProjectsService{client: c}
	c.Qualitygates = &QualitygatesService{client: c}
	c.Qualityprofiles = &QualityprofilesService{client: c}
	c.Rules = &RulesService{client: c}
	c.Server = &ServerService{client: c}
	c.Settings = &SettingsService{client: c}
	c.Sources = &SourcesService{client: c}
	c.System = &SystemService{client: c}
	c.UserGroups = &UserGroupsService{client: c}
	c.UserTokens = &UserTokensService{client: c}
	c.Users = &UsersService{client: c}
	c.Webhooks = &WebhooksService{client: c}
	c.Webservices = &WebservicesService{client: c}
	return c, nil
}
