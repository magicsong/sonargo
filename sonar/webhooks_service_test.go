package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func createWebhook(name string) string {
	opt := &WebhooksCreateOption{
		Name:    name,
		Project: AnalyzedProject,
		Url:     "http://localhost:1080",
	}
	v, resp, err := client.Webhooks.Create(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "already been taken")).To(BeTrue())
		return ""
	}
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Webhook.Name).To(Equal(name))
	return v.Webhook.Key
}

func deleteWebhook(key string) {
	opt := &WebhooksDeleteOption{Webhook: key}
	assertNoBody(client.Webhooks.Delete(opt))
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Create/delete in api/webhooks", func() {
		It("Should be ok to create/delete webhook", func() {
			key := createWebhook("testWebhook")
			deleteWebhook(key)
		})
	})
	Describe("Test Deliveries in api/webhooks", func() {
		It("Should be ok", func() {
			opt := &WebhooksDeliveriesOption{
				CeTaskId:     "",
				ComponentKey: AnalyzedProject,
				P:            "",
				Ps:           "",
				Webhook:      "",
			}
			v, resp, err := client.Webhooks.Deliveries(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Deliveries).NotTo(BeNil())
		})
	})
	Describe("Test Delivery in api/webhooks", func() {
		It("Should be ok to get a Delivery", func() {
			opt := &WebhooksDeliveryOption{DeliveryId: "AWdYMuOPd-5YlXz1HENB"}
			v, resp, err := client.Webhooks.Delivery(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Delivery.ID).To(Equal("AWdYMuOPd-5YlXz1HENB"))
		})
	})
	Describe("Test List in api/webhooks", func() {
		It("Should be ok", func() {
			opt := &WebhooksListOption{Project: AnalyzedProject}
			v, resp, err := client.Webhooks.List(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Webhooks).NotTo(BeNil())
		})
	})
	Describe("Test Update in api/webhooks", func() {
		It("Should be ok", func() {
			name := "testwebhook"
			key := createWebhook(name)
			opt := &WebhooksUpdateOption{
				Name:    name,
				Url:     "http://localhost:1010",
				Webhook: key,
			}
			resp, err := client.Webhooks.Update(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(204))
			Expect(resp.ContentLength).To(BeZero())
			deleteWebhook(key)
		})
	})
})
