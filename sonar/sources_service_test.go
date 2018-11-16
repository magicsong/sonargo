package sonargo_test

import (
	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Raw in api/sources", func() {
		It("Should be ok", func() {
			opt := &SourcesRawOption{Key: AnalyzedProject + ":doc.go"}
			v, resp, err := client.Sources.Raw(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*v).NotTo(BeEmpty())
		})
	})
	Describe("Test Scm in api/sources", func() {
		It("Should be ok", func() {
			opt := &SourcesSCMOption{
				CommitsByLine: "",
				From:          "",
				Key:           AnalyzedProject + ":doc.go",
				To:            "",
			}
			v, resp, err := client.Sources.SCM(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.SCM).NotTo(BeNil())
		})
	})
	Describe("Test Show in api/sources", func() {
		It("Should be ok", func() {
			opt := &SourcesShowOption{
				From: "",
				Key:  AnalyzedProject + ":doc.go",
				To:   "",
			}
			v, resp, err := client.Sources.Show(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Sources).NotTo(BeNil())
		})
	})
})
