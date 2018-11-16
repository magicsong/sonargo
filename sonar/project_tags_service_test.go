package sonargo_test

import (
	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Search in api/project_tags", func() {
		It("Should be ok to search tags", func() {
			opt := &ProjectTagsSearchOption{
				Ps: 0,
				Q:  "sonar",
			}
			v, resp, err := client.ProjectTags.Search(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test Set in api/project_tags", func() {
		It("Should be ok", func() {
			empty := ""
			opt := &ProjectTagsSetOption{
				Project: AnalyzedProject,
				Tags:    &empty,
			}
			assertNoBody(client.ProjectTags.Set(opt))
		})
	})
})
