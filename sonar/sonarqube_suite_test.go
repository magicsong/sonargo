package sonargo_test

import (
	"net/http"
	"os"
	"testing"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var client *Client
var (
	testUser1  string
	testUser2  string
	testUser3  string
	testGroup1 int
	testGroup2 int

	testProject string

	testQualityGate    int
	testQualityProfile string
	testTemplate       string

	testComponentId string
	testTaskId      string

	testIssueKey string
)

const AnalyzedProject = "sonargo"

func assertNoBody(resp *http.Response, err error) {
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(204))
	Expect(resp.ContentLength).To(BeZero())
}

var _ = BeforeSuite(func() {
	sonarURL := os.Getenv("SONAR_URL")
	if sonarURL == "" {
		Fail("SONAR_URL should not be empty")
	}
	c, err := NewClient(sonarURL+"/api", "admin", "admin")
	Expect(err).ShouldNot(HaveOccurred())
	client = c
	//create some users
	testUser1 = createTestUser("magicsong", "magicsong@yunify.com", "passw0rd")
	testUser2 = createTestUser("testone", "testone@test.com", "testone")
	testUser3 = createTestUser("testtwo", "testtwo@test.com", "testtwo")

	//create some groups
	testGroup1 = createGroup("test1group")
	testGroup2 = createGroup("test2group")
	//create test project
	testProject = createProject("test-project1")
	//create test quality gate
	testQualityGate = createQualityGate("testQualityGate")
	//create permission template
	testTemplate = createTemplate("testTemplate")

	v, _, err := client.Ce.Activity(new(CeActivityOption))
	Expect(err).ShouldNot(HaveOccurred())
	testTaskId = v.Tasks[0].ID
	testComponentId = v.Tasks[0].ComponentID
	testQualityProfile = createQualityProfile("test sonar way", "java")
	//get test issue
	testIssueKey = getIssue()
})

func TestSonar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sonar Integration test suite")
}
