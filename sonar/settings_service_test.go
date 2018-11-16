package sonargo_test

import (
	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func resetProjectSetting(key string) {
	opt := &SettingsResetOption{
		Component: AnalyzedProject,
		Keys:      key,
	}
	assertNoBody(client.Settings.Reset(opt))
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test ListDefinitions in api/settings", func() {
		It("Should be ok", func() {
			opt := &SettingsListDefinitionsOption{Component: testProject}
			v, resp, err := client.Settings.ListDefinitions(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Definitions).NotTo(BeNil())
		})
	})
	Describe("Test Reset in api/settings", func() {
		It("Should be ok", func() {
			resetProjectSetting("sonar.css.file.suffixes")
		})
	})
	Describe("Test Set in api/settings", func() {
		It("Should be ok", func() {
			opt := &SettingsSetOption{
				Component:   AnalyzedProject,
				FieldValues: "",
				Key:         "sonar.go.file.suffixes",
				Value:       "",
				Values:      []string{".go", ".go1"},
			}
			assertNoBody(client.Settings.Set(opt))
			resetProjectSetting("sonar.go.file.suffixes")
		})
	})
	Describe("Test Values in api/settings", func() {
		It("Should be ok", func() {
			opt := &SettingsValuesOption{
				Component: AnalyzedProject,
				Keys:      "sonar.go.file.suffixes",
			}
			v, resp, err := client.Settings.Values(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Settings[0].Values[0]).To(Equal(".go"))
		})
	})
})
