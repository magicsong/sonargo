package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func createProject(name string) string {
	opt := &ProjectsCreateOption{
		Branch:     "",
		Name:       name,
		Project:    name,
		Visibility: "",
	}
	v, resp, err := client.Projects.Create(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "already exists")).To(BeTrue())
		return name
	}
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Project.Name).To(Equal(name))
	return name
}

func deleteProject(name string) {
	opt := &ProjectsDeleteOption{
		Project:   name,
		ProjectId: "",
	}
	resp, err := client.Projects.Delete(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "not found")).To(BeTrue())
		return
	}
	assertNoBody(resp, err)
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test BulkDelete in api/projects", func() {
		BeforeEach(func() {
			createProject("test1")
			createProject("test2")
			createProject("test3")
		})
		It("Should be ok to delete projects in bulk mode", func() {
			opt := &ProjectsBulkDeleteOption{
				AnalyzedBefore:    "",
				OnProvisionedOnly: "",
				ProjectIds:        "",
				Projects:          "test1,test2,test3",
				Q:                 "",
				Qualifiers:        "",
			}
			assertNoBody(client.Projects.BulkDelete(opt))
		})
	})
	Describe("Test BulkUpdateKey in api/projects", func() {
		It("Should be ok to bulkupdate a project", func() {
			opt := &ProjectsBulkUpdateKeyOption{
				DryRun:    "true",
				From:      "sonargo",
				Project:   AnalyzedProject,
				ProjectId: "",
				To:        "pro-sonargo",
			}
			v, resp, err := client.Projects.BulkUpdateKey(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Keys[0].NewKey).To(Equal("pro-sonargo"))
		})
	})
	Describe("Test Create in api/projects", func() {
		testProjectName := "test-proj"
		It("Should be ok", func() {
			createProject(testProjectName)
		})
		AfterEach(func() {
			deleteProject(testProjectName)
		})
	})
	Describe("Test Delete in api/projects", func() {
		It("Should be ok", func() {
			deleteProject("test-proj")
		})
	})
	Describe("Test Search in api/projects", func() {
		It("Should be ok", func() {
			opt := &ProjectsSearchOption{
				AnalyzedBefore:    "",
				OnProvisionedOnly: "",
				P:                 "",
				ProjectIds:        "",
				Projects:          "",
				Ps:                "",
				Q:                 "",
				Qualifiers:        QualifierProject,
			}
			v, resp, err := client.Projects.Search(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test UpdateKey in api/projects", func() {
		It("Should be ok", func() {
			opt := &ProjectsUpdateKeyOption{
				From: AnalyzedProject,
				To:   "pro-project1",
			}
			assertNoBody(client.Projects.UpdateKey(opt))
			//change back
			opt.From = "pro-project1"
			opt.To = AnalyzedProject
			assertNoBody(client.Projects.UpdateKey(opt))
		})
	})
	Describe("Test UpdateVisibility in api/projects", func() {
		It("Should be ok", func() {
			opt := &ProjectsUpdateVisibilityOption{
				Project:    AnalyzedProject,
				Visibility: ProjectVisibilityPrivate,
			}
			assertNoBody(client.Projects.UpdateVisibility(opt))
			//change back
			opt.Visibility = ProjectVisibilityPublic
			assertNoBody(client.Projects.UpdateVisibility(opt))
		})
	})
})
