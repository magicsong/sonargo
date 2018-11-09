// Manage user favorites
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type FavoritesService struct {
	client *Client
}

type Favorites struct {
	Favorites []struct {
		Key          string `json:"key,omitempty"`
		Name         string `json:"name,omitempty"`
		Organization string `json:"organization,omitempty"`
		Qualifier    string `json:"qualifier,omitempty"`
	} `json:"favorites,omitempty"`
	Paging struct {
		PageIndex int64 `json:"pageIndex,omitempty"`
		PageSize  int64 `json:"pageSize,omitempty"`
		Total     int64 `json:"total,omitempty"`
	} `json:"paging,omitempty"`
}

type FavoritesAddOption struct {
	Component string `url:"component,omitempty"` // Description:"Component key",ExampleValue:"my_project:/src/foo/Bar.php"
}

// Add Add a component (project, directory, file etc.) as favorite for the authenticated user.<br>Requires authentication and the following permission: 'Browse' on the project of the specified component.
func (s *FavoritesService) Add(opt *FavoritesAddOption) (resp *string, err error) {
	err := s.ValidateAddOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "favorites/add", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type FavoritesRemoveOption struct {
	Component string `url:"component,omitempty"` // Description:"Component key",ExampleValue:"my_project"
}

// Remove Remove a component (project, directory, file etc.) as favorite for the authenticated user.<br>Requires authentication.
func (s *FavoritesService) Remove(opt *FavoritesRemoveOption) (resp *string, err error) {
	err := s.ValidateRemoveOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "favorites/remove", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type FavoritesSearchOption struct {
	P  string `url:"p,omitempty"`  // Description:"1-based page number",ExampleValue:"42"
	Ps string `url:"ps,omitempty"` // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
}

// Search Search for the authenticated user favorites.<br>Requires authentication.
func (s *FavoritesService) Search(opt *FavoritesSearchOption) (resp *Favorites, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "favorites/search", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
