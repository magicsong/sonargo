// Manage branch (only available when the Branch plugin is installed)
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type Project_branchesService struct {
	client *Client
}

// [TODO] you should call the <Delete> func manually and complete the fields of this struct
type ProjectBranchesDeleteResp struct{}
type ProjectBranchesDeleteOption struct {
	Branch  string `url:"branch,omitempty"`  // Description:"Name of the branch",ExampleValue:"branch1"
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Delete Delete a non-main branch of a project.<br/>Requires 'Administer' rights on the specified project.
func (s *Project_branchesService) Delete(opt *ProjectBranchesDeleteOption) (Resp *ProjectBranchesDeleteResp, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_branches/delete", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <List> func manually and complete the fields of this struct
type ProjectBranchesListResp struct{}
type ProjectBranchesListOption struct {
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// List List the branches of a project.<br/>Requires 'Browse' or 'Execute analysis' rights on the specified project.
func (s *Project_branchesService) List(opt *ProjectBranchesListOption) (Resp *ProjectBranchesListResp, err error) {
	err := s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_branches/list", Opt)
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
type ProjectBranchesRenameResp struct{}
type ProjectBranchesRenameOption struct {
	Name    string `url:"name,omitempty"`    // Description:"New name of the main branch",ExampleValue:"branch1"
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Rename Rename the main branch of a project.<br/>Requires 'Administer' permission on the specified project.
func (s *Project_branchesService) Rename(opt *ProjectBranchesRenameOption) (Resp *ProjectBranchesRenameResp, err error) {
	err := s.ValidateRenameOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_branches/rename", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
