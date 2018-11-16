package sonargo_test

import (
	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//Only Available in Commercial Edition, so skip
var _ = XDescribe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Delete in api/project_branches", func() {
		It("Should be ok", func() {
			opt := &ProjectBranchesDeleteOption{
				Branch:  "branch1",
				Project: AnalyzedProject,
			}
			assertNoBody(client.ProjectBranches.Delete(opt))
		})
	})
	Describe("Test List in api/project_branches", func() {
		It("Should be ok", func() {
			opt := &ProjectBranchesListOption{Project: AnalyzedProject}
			v, resp, err := client.ProjectBranches.List(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Branches).NotTo(BeZero())
		})
	})
	Describe("Test Rename in api/project_branches", func() {
		It("Should be ok", func() {
			opt := &ProjectBranchesRenameOption{
				Name:    "test-branche",
				Project: AnalyzedProject,
			}
			assertNoBody(client.ProjectBranches.Rename(opt))
		})
	})
})
