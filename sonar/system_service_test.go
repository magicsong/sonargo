package sonargo_test

import (
	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test ChangeLogLevel in api/system", func() {
		It("Should be ok", func() {
			opt := &SystemChangeLogLevelOption{Level: LogLevelInfo}
			assertNoBody(client.System.ChangeLogLevel(opt))
		})
	})
	Describe("Test DbMigrationStatus in api/system", func() {
		It("Should be ok", func() {
			v, resp, err := client.System.DbMigrationStatus()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test Health in api/system", func() {
		It("Should be ok", func() {
			v, resp, err := client.System.Health()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test Logs in api/system", func() {
		It("Should be ok", func() {
			opt := &SystemLogsOption{Process: "app"}
			v, resp, err := client.System.Logs(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*v).NotTo(BeEmpty())
		})
	})
	Describe("Test MigrateDb in api/system", func() {
		XIt("Should be ok", func() {
			v, resp, err := client.System.MigrateDb()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test Ping in api/system", func() {
		It("Should be ok", func() {
			v, resp, err := client.System.Ping()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*v).To(Equal("pong"))
		})
	})
	Describe("Test Restart in api/system", func() {
		XIt("Should be ok", func() {
			assertNoBody(client.System.Restart())
		})
	})
	Describe("Test Status in api/system", func() {
		It("Should be ok", func() {
			v, resp, err := client.System.Status()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test Upgrades in api/system", func() {
		It("Should be ok", func() {
			v, resp, err := client.System.Upgrades()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
})
