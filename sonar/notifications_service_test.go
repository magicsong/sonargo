package sonargo_test

import (
	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SonarCLI integration test", func() {

	JustBeforeEach(func() {})
	Describe("Test Add/Remove in api/notifications", func() {
		BeforeEach(func() {
			optRemove := &NotificationsRemoveOption{
				Channel: "",
				Login:   "",
				Project: "sonargo",
				Type:    "ChangesOnMyIssue",
			}
			client.Notifications.Remove(optRemove)
		})
		It("Should be ok", func() {
			opt := &NotificationsAddOption{
				Channel: "",
				Login:   "",
				Project: "sonargo",
				Type:    "ChangesOnMyIssue",
			}
			assertNoBody(client.Notifications.Add(opt))

			optRemove := &NotificationsRemoveOption{
				Channel: "",
				Login:   "",
				Project: "sonargo",
				Type:    "ChangesOnMyIssue",
			}
			assertNoBody(client.Notifications.Remove(optRemove))

		})
	})
	Describe("Test List in api/notifications", func() {
		It("Should be ok", func() {
			opt := &NotificationsListOption{Login: ""}
			v, resp, err := client.Notifications.List(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
})
