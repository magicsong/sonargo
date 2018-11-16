package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Available in api/plugins", func() {
		It("Should be ok to Test Available", func() {
			v, resp, err := client.Plugins.Available()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Plugins).NotTo(BeNil())
		})
	})
	Describe("Test CancelAll in api/plugins", func() {
		XIt("Should be ok but it will be skip", func() {
			assertNoBody(client.Plugins.CancelAll())
		})
	})
	Describe("Test Install in api/plugins", func() {
		It("Should be ok", func() {
			opt := &PluginsInstallOption{Key: "gitlab"}
			resp, err := client.Plugins.Install(opt)
			if err != nil {
				Expect(strings.Contains(err.Error(), "No plugin with key 'gitlab' or plugin 'gitlab' is already")).To(BeTrue())
			} else {
				assertNoBody(resp, err)
			}
		})
	})
	Describe("Test Installed in api/plugins", func() {
		It("Should be ok", func() {
			opt := &PluginsInstalledOption{F: ""}
			v, resp, err := client.Plugins.Installed(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Plugins).NotTo(BeNil())
		})
	})
	Describe("Test Pending in api/plugins", func() {
		It("Should be ok", func() {
			v, resp, err := client.Plugins.Pending()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Installing).NotTo(BeNil())
		})
	})
	Describe("Test Uninstall in api/plugins", func() {
		It("Should be ok", func() {
			opt := &PluginsUninstallOption{Key: "gitlab"}
			resp, err := client.Plugins.Uninstall(opt)
			if err != nil {
				Expect(strings.Contains(err.Error(), "installed")).To(BeTrue())
			} else {
				assertNoBody(resp, err)
			}
		})
	})
	Describe("Test Update in api/plugins", func() {
		It("Should be ok", func() {
			opt := &PluginsUpdateOption{Key: "gitlab"}
			resp, err := client.Plugins.Update(opt)
			if err != nil {
				Expect(strings.Contains(err.Error(), "No plugin with key 'gitlab' or plugin 'gitlab' is already in latest compatible version")).To(BeTrue())
			} else {
				assertNoBody(resp, err)
			}
		})
	})
	Describe("Test Updates in api/plugins", func() {
		It("Should be ok", func() {
			v, resp, err := client.Plugins.Updates()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Plugins).NotTo(BeNil())
		})
	})
})
