// Get details on source files. See also api/tests.
package sonargo // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type SourcesService struct {
	client *Client
}

type Sources struct {
	Scm     [][]string      `json:"scm,omitempty"`
	Sources [][]interface{} `json:"sources,omitempty"`
}

type SourcesRawOption struct {
	Key string `url:"key,omitempty"` // Description:"File key",ExampleValue:"my_project:src/foo/Bar.php"
}

// Raw Get source code as raw text. Require 'See Source Code' permission on file
func (s *SourcesService) Raw(opt *SourcesRawOption) (resp *string, err error) {
	err := s.ValidateRawOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "sources/raw", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type SourcesScmOption struct {
	CommitsByLine string `url:"commits_by_line,omitempty"` // Description:"Group lines by SCM commit if value is false, else display commits for each line, even if two consecutive lines relate to the same commit.",ExampleValue:""
	From          string `url:"from,omitempty"`            // Description:"First line to return. Starts at 1",ExampleValue:"10"
	Key           string `url:"key,omitempty"`             // Description:"File key",ExampleValue:"my_project:/src/foo/Bar.php"
	To            string `url:"to,omitempty"`              // Description:"Last line to return (inclusive)",ExampleValue:"20"
}

// Scm Get SCM information of source files. Require See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Author of the commit</li><li>Datetime of the commit (before 5.2 it was only the Date)</li><li>Revision of the commit (added in 5.2)</li></ol>
func (s *SourcesService) Scm(opt *SourcesScmOption) (resp *Sources, err error) {
	err := s.ValidateScmOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "sources/scm", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}

type SourcesShowOption struct {
	From string `url:"from,omitempty"` // Description:"First line to return. Starts at 1",ExampleValue:"10"
	Key  string `url:"key,omitempty"`  // Description:"File key",ExampleValue:"my_project:/src/foo/Bar.php"
	To   string `url:"to,omitempty"`   // Description:"Last line to return (inclusive)",ExampleValue:"20"
}

// Show Get source code. Require See Source Code permission on file's project<br/>Each element of the result array is composed of:<ol><li>Line number</li><li>Content of the line</li></ol>
func (s *SourcesService) Show(opt *SourcesShowOption) (resp *Sources, err error) {
	err := s.ValidateShowOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "sources/show", opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, resp)
	if err != nil {
		return
	}
	return
}
