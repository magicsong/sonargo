package sonargo_test

import (
	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//only available in commercial edition
var _ = XDescribe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Delete in api/project_pull_requests", func() {
		It("Should be ok", func() {
			opt := &ProjectPullRequestsDeleteOption{
				Project:     testProject,
				PullRequest: 12,
			}
			assertNoBody(client.ProjectPullRequests.Delete(opt))
		})
	})
	Describe("Test List in api/project_pull_requests", func() {
		It("Should be ok", func() {
			opt := &ProjectPullRequestsListOption{Project: testProject}
			v, resp, err := client.ProjectPullRequests.List(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.PullRequests).NotTo(BeNil())
		})
	})
})
