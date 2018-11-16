package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SonarCLI integration test", func() {
	Describe("Test Activity in api/ce", func() {
		It("Should be ok", func() {
			opt := &CeActivityOption{
				ComponentId:    testComponentId,
				MaxExecutedAt:  "",
				MinSubmittedAt: "",
				OnlyCurrents:   "",
				Ps:             "",
				Q:              "",
				Status:         "",
				Type:           "",
			}
			v, resp, err := client.Ce.Activity(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Tasks[0].ComponentID).To(Equal(testComponentId))
		})
		It("Should get the error", func() {
			opt := &CeActivityOption{
				ComponentId:    "whatever",
				MaxExecutedAt:  "",
				MinSubmittedAt: "",
				OnlyCurrents:   "",
				Ps:             "",
				Q:              "",
				Status:         "",
				Type:           "",
			}
			v, resp, err := client.Ce.Activity(opt)
			Expect(strings.Contains(err.Error(), "Component 'whatever' does not exist")).Should(BeTrue())
			Expect(resp.StatusCode).To(Equal(404))
			Expect(v).To(BeNil())
		})
	})
	Describe("Test Component in api/ce", func() {
		It("Should be ok", func() {
			opt := &CeComponentOption{
				Component:   AnalyzedProject,
				ComponentId: "",
			}
			v, resp, err := client.Ce.Component(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Current.ComponentID).To(Equal(testComponentId))
		})
		It("Should get Error", func() {
			opt := &CeComponentOption{
				Component:   "",
				ComponentId: "whatever",
			}
			v, resp, err := client.Ce.Component(opt)
			Expect(strings.Contains(err.Error(), "'whatever' not found")).Should(BeTrue())
			Expect(resp.StatusCode).To(Equal(404))
			Expect(v).To(BeNil())
		})
	})
	Describe("Test Task in api/ce", func() {
		It("Should be ok", func() {
			opt := &CeTaskOption{
				AdditionalFields: "",
				Id:               testTaskId,
			}
			v, resp, err := client.Ce.Task(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Task.ComponentID).To(Equal(testComponentId))
		})
	})
})
