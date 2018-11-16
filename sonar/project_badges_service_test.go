package sonargo_test

import (
	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Measure in api/project_badges", func() {
		It("Should be ok", func() {
			opt := &ProjectBadgesMeasureOption{
				Branch:      "",
				Metric:      "bugs",
				Project:     AnalyzedProject,
				PullRequest: "",
			}
			v, resp, err := client.ProjectBadges.Measure(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*v).NotTo(BeZero())
		})
	})
	Describe("Test QualityGate in api/project_badges", func() {
		It("Should be ok", func() {
			opt := &ProjectBadgesQualityGateOption{
				Branch:      "",
				Project:     AnalyzedProject,
				PullRequest: "",
			}
			v, resp, err := client.ProjectBadges.QualityGate(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*v).NotTo(BeZero())
		})
	})
})
