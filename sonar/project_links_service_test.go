package sonargo_test

import (
	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func createLink() string {
	opt := &ProjectLinksCreateOption{
		Name:       "testlink",
		ProjectId:  "",
		ProjectKey: AnalyzedProject,
		Url:        "https://www.google.com",
	}
	v, resp, err := client.ProjectLinks.Create(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Link.URL).To(Equal("https://www.google.com"))
	return v.Link.ID
}

func deleteLink(id string) {
	opt := &ProjectLinksDeleteOption{Id: id}
	assertNoBody(client.ProjectLinks.Delete(opt))
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Create in api/project_links", func() {
		var id string
		It("Should be ok to create link", func() {
			id = createLink()
		})
		AfterEach(func() {
			if id != "" {
				deleteLink(id)
			}
		})

	})
	Describe("Test Delete in api/project_links", func() {
		It("Should be ok", func() {
			deleteLink(createLink())
		})
	})
	Describe("Test Search in api/project_links", func() {
		It("Should be ok", func() {
			opt := &ProjectLinksSearchOption{
				ProjectId:  "",
				ProjectKey: testProject,
			}
			v, resp, err := client.ProjectLinks.Search(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeZero())
		})
	})
})
