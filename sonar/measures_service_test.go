package sonargo_test

import (
	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Component in api/measures", func() {
		It("Should be ok", func() {
			opt := &MeasuresComponentOption{
				AdditionalFields: "",
				Component:        "",
				ComponentId:      testComponentId,
				MetricKeys:       "ncloc,complexity,violations",
			}
			v, resp, err := client.Measures.Component(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(len(v.Component.Measures)).NotTo(BeZero())
		})
	})
	Describe("Test ComponentTree in api/measures", func() {
		It("Should be ok", func() {
			opt := &MeasuresComponentTreeOption{
				AdditionalFields: "",
				Asc:              "",
				BaseComponentId:  testComponentId,
				Component:        "",
				MetricKeys:       "ncloc,complexity,violations",
				MetricPeriodSort: "",
				MetricSort:       "",
				MetricSortFilter: "",
				P:                "",
				Ps:               "",
				Q:                "",
				Qualifiers:       "",
				S:                "",
				Strategy:         "",
			}
			v, resp, err := client.Measures.ComponentTree(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.BaseComponent.ID).To(Equal(testComponentId))
		})
	})
	Describe("Test SearchHistory in api/measures", func() {
		It("Should be ok", func() {
			opt := &MeasuresSearchHistoryOption{
				Component: "sonargo",
				From:      "",
				Metrics:   "ncloc,complexity,violations",
				P:         "",
				Ps:        "",
				To:        "",
			}
			v, resp, err := client.Measures.SearchHistory(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(len(v.Measures)).NotTo(BeZero())
		})
	})
})
