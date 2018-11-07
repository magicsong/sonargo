// Manage user favorites
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type FavoritesService struct {
	client *Client
}

// [TODO] you should call the <Add> func manually and complete the fields of this struct
type FavoritesAddResp struct{}
type FavoritesAddOption struct {
	Component string `url:"component,omitempty"` // Description:"Component key",ExampleValue:"my_project:/src/foo/Bar.php"
}

// Add Add a component (project, directory, file etc.) as favorite for the authenticated user.<br>Requires authentication and the following permission: 'Browse' on the project of the specified component.
func (s *FavoritesService) Add(opt *FavoritesAddOption) (Resp *FavoritesAddResp, err error) {
	err := s.ValidateAddOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "favorites/add", Opt)
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
type FavoritesRemoveResp struct{}
type FavoritesRemoveOption struct {
	Component string `url:"component,omitempty"` // Description:"Component key",ExampleValue:"my_project"
}

// Remove Remove a component (project, directory, file etc.) as favorite for the authenticated user.<br>Requires authentication.
func (s *FavoritesService) Remove(opt *FavoritesRemoveOption) (Resp *FavoritesRemoveResp, err error) {
	err := s.ValidateRemoveOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "favorites/remove", Opt)
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
type FavoritesSearchResp struct{}
type FavoritesSearchOption struct {
	P  string `url:"p,omitempty"`  // Description:"1-based page number",ExampleValue:"42"
	Ps string `url:"ps,omitempty"` // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
}

// Search Search for the authenticated user favorites.<br>Requires authentication.
func (s *FavoritesService) Search(opt *FavoritesSearchOption) (Resp *FavoritesSearchResp, err error) {
	err := s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "favorites/search", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
