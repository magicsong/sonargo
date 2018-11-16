package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SonarCLI integration test", func() {
	Describe("Test Search in api/components", func() {
		It("Should be ok", func() {
			opt := &ComponentsSearchOption{
				Language:   "",
				P:          "",
				Ps:         "",
				Q:          "",
				Qualifiers: QualifierProject,
			}
			v, resp, err := client.Components.Search(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Components[0].Qualifier).To(Equal("TRK"))
		})
		It("Should get nothing", func() {
			opt := &ComponentsSearchOption{
				Language:   "",
				P:          "",
				Ps:         "",
				Q:          "whatever-not-found",
				Qualifiers: QualifierSubProject,
			}
			v, resp, err := client.Components.Search(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(len(v.Components)).To(Equal(0))
		})
	})
	Describe("Test Show in api/components", func() {
		It("Should be ok", func() {
			opt := &ComponentsShowOption{
				Component:   "",
				ComponentId: testComponentId,
			}
			v, resp, err := client.Components.Show(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Component.ID).To(Equal(testComponentId))
		})

		It("Should get error", func() {
			opt := &ComponentsShowOption{
				Component:   "",
				ComponentId: "whatever",
			}
			v, resp, err := client.Components.Show(opt)
			Expect(strings.Contains(err.Error(), "'whatever' not found")).To(BeTrue())
			Expect(resp.StatusCode).To(Equal(404))
			Expect(v).To(BeNil())
		})
	})
	Describe("Test Tree in api/components", func() {
		It("Should be ok", func() {
			opt := &ComponentsTreeOption{
				Asc:         "",
				Component:   "",
				ComponentId: testComponentId,
				P:           "",
				Ps:          "",
				Q:           "",
				Qualifiers:  "",
				S:           "",
				Strategy:    "",
			}
			v, resp, err := client.Components.Tree(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.BaseComponent.ID).To(Equal(testComponentId))
		})

		It("Should get error", func() {
			opt := &ComponentsTreeOption{
				Asc:         "",
				Component:   "",
				ComponentId: "whatever",
				P:           "",
				Ps:          "",
				Q:           "",
				Qualifiers:  "",
				S:           "",
				Strategy:    "",
			}
			v, resp, err := client.Components.Tree(opt)
			Expect(strings.Contains(err.Error(), "'whatever' not found")).To(BeTrue())
			Expect(resp.StatusCode).To(Equal(404))
			Expect(v).To(BeNil())
		})
	})
})
