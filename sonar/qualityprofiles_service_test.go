package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func getQualityProfileKey(name, language string) string {
	opt := &QualityProfilesSearchOption{
		Defaults:       "",
		Language:       language,
		Project:        "",
		QualityProfile: name,
	}
	v, resp, err := client.QualityProfiles.Search(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Profiles).NotTo(BeNil())
	return v.Profiles[0].Key
}
func createQualityProfile(name, language string) string {
	opt := &QualityProfilesCreateOption{
		BackupSonarlintVsCsFake: "",
		Language:                language,
		Name:                    name,
	}
	v, resp, err := client.QualityProfiles.Create(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "already exists")).To(BeTrue())
		return getQualityProfileKey(name, language)
	}
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Profile.Name).To(Equal(name))
	return v.Profile.Key
}

func deleteQualityProfile(name, language string) {
	opt := &QualityProfilesDeleteOption{
		Key:            "",
		Language:       language,
		QualityProfile: name,
	}
	resp, err := client.QualityProfiles.Delete(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(204))
	Expect(resp.ContentLength).To(BeZero())
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test ActivateRule in api/qualityprofiles", func() {
		It("Should be ok", func() {
			ruleKey := createRule("test_rule")
			opt := &QualityProfilesActivateRuleOption{
				Key:      testQualityProfile,
				Params:   "",
				Reset:    "",
				Rule:     ruleKey,
				Severity: "",
			}
			assertNoBody(client.QualityProfiles.ActivateRule(opt))
			deactiveOpt := &QualityProfilesDeactivateRuleOption{
				Key:  opt.Key,
				Rule: opt.Rule,
			}
			assertNoBody(client.QualityProfiles.DeactivateRule(deactiveOpt))
		})
	})
	Describe("Test ActivateRules in api/qualityprofiles", func() {
		It("Should be ok", func() {
			opt := &QualityProfilesActivateRulesOption{
				Activation:       "",
				ActiveSeverities: "",
				Asc:              "",
				AvailableSince:   "",
				Inheritance:      "",
				IsTemplate:       "",
				Languages:        "",
				Q:                "",
				Qprofile:         "",
				Repositories:     "",
				RuleKey:          "",
				S:                "",
				Severities:       "",
				Statuses:         "",
				Tags:             "",
				TargetKey:        testQualityProfile,
				TargetSeverity:   "",
				TemplateKey:      "",
				Types:            "",
			}
			v, resp, err := client.QualityProfiles.ActivateRules(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Succeeded).NotTo(BeNil())
		})
	})
	Describe("Test AddProject in api/qualityprofiles", func() {
		It("Should be ok", func() {
			opt := &QualityProfilesAddProjectOption{
				Key:            testQualityProfile,
				Language:       "",
				Project:        testProject,
				ProjectUuid:    "",
				QualityProfile: "",
			}
			assertNoBody(client.QualityProfiles.AddProject(opt))
			//remove
			removeOpt := QualityProfilesRemoveProjectOption(*opt)
			assertNoBody(client.QualityProfiles.RemoveProject(&removeOpt))
		})
	})
	Describe("Test Backup in api/qualityprofiles", func() {
		It("Should be ok", func() {
			opt := &QualityProfilesBackupOption{
				ProfileKey: testQualityProfile,
			}
			v, resp, err := client.QualityProfiles.Backup(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*v).NotTo(BeEmpty())
		})
	})
	Describe("Test ChangeParent in api/qualityprofiles", func() {
		It("Should be ok", func() {
			opt := &QualityProfilesChangeParentOption{
				Key:                  "",
				Language:             "java",
				ParentKey:            "",
				ParentQualityProfile: "Sonar way",
				QualityProfile:       "test sonar way",
			}
			assertNoBody(client.QualityProfiles.ChangeParent(opt))
		})
	})
	Describe("Test Changelog in api/qualityprofiles", func() {
		It("Should be ok", func() {
			opt := &QualityProfilesChangelogOption{
				Key:            testQualityProfile,
				Language:       "",
				P:              "",
				Ps:             "",
				QualityProfile: "",
				Since:          "",
				To:             "",
			}
			v, resp, err := client.QualityProfiles.Changelog(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Events).NotTo(BeNil())
		})
	})
	Describe("Test Copy in api/qualityprofiles", func() {
		It("Should be ok", func() {
			opt := &QualityProfilesCopyOption{
				FromKey: testQualityProfile,
				ToName:  "copy sonar way",
			}
			v, resp, err := client.QualityProfiles.Copy(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Name).To(Equal("copy sonar way"))
			deleteQualityProfile("copy sonar way", "java")
		})
	})
	Describe("Test Create/Delete in api/qualityprofiles", func() {
		It("Should be ok", func() {
			createQualityProfile("my sonar way", "go")
			deleteQualityProfile("my sonar way", "go")
		})
	})
	Describe("Test DeactivateRules in api/qualityprofiles", func() {
		It("Should be ok", func() {
			opt := &QualityProfilesDeactivateRulesOption{
				Activation:       "",
				ActiveSeverities: "",
				Asc:              "",
				AvailableSince:   "",
				Inheritance:      "",
				IsTemplate:       "",
				Languages:        "",
				Q:                "",
				Qprofile:         "",
				Repositories:     "",
				RuleKey:          "",
				S:                "",
				Severities:       "",
				Statuses:         "",
				Tags:             "",
				TargetKey:        testQualityProfile,
				TemplateKey:      "",
				Types:            "",
			}
			v, resp, err := client.QualityProfiles.DeactivateRules(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Succeeded).NotTo(BeNil())
		})
	})
	Describe("Test Export in api/qualityprofiles", func() {
		It("Should be ok", func() {
			createQualityProfile("test cs profile", "cs")
			opt := &QualityProfilesExportOption{
				ExporterKey:    "sonarlint-vs-cs",
				Key:            "",
				Language:       "cs",
				QualityProfile: "test cs profile",
			}
			v, resp, err := client.QualityProfiles.Export(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(*v).NotTo(BeEmpty())
			deleteQualityProfile("test cs profile", "cs")
		})
	})
	Describe("Test Exporters in api/qualityprofiles", func() {
		It("Should be ok", func() {
			v, resp, err := client.QualityProfiles.Exporters()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Exporters).NotTo(BeNil())
		})
	})
	Describe("Test Importers in api/qualityprofiles", func() {
		It("Should be ok", func() {
			v, resp, err := client.QualityProfiles.Importers()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Importers).NotTo(BeNil())
		})
	})
	Describe("Test Inheritance in api/qualityprofiles", func() {
		It("Should be ok", func() {
			opt := &QualityProfilesInheritanceOption{
				Key:            testQualityProfile,
				Language:       "",
				QualityProfile: "",
			}
			v, resp, err := client.QualityProfiles.Inheritance(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Profile).NotTo(BeNil())
		})
	})
	Describe("Test Projects in api/qualityprofiles", func() {
		It("Should be ok", func() {
			opt := &QualityProfilesProjectsOption{
				Key:      testQualityProfile,
				P:        0,
				Ps:       0,
				Q:        "",
				Selected: "",
			}
			v, resp, err := client.QualityProfiles.Projects(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Results).NotTo(BeNil())
		})
	})
	Describe("Test Rename in api/qualityprofiles", func() {
		It("Should be ok", func() {
			opt := &QualityProfilesRenameOption{
				Key:  testQualityProfile,
				Name: "new sonar way",
			}
			assertNoBody(client.QualityProfiles.Rename(opt))
			opt.Name = "test sonar way"
			assertNoBody(client.QualityProfiles.Rename(opt))
		})
	})
	Describe("Test Restore in api/qualityprofiles", func() {
		XIt("Should be ok", func() {
			opt := &QualityProfilesRestoreOption{Backup: "MUST_EDIT_IT"}
			resp, err := client.QualityProfiles.Restore(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(resp.ContentLength).To(BeZero())
		})
	})
	Describe("Test Search in api/qualityprofiles", func() {
		It("Should be ok to get all qualityprofiles", func() {
			getQualityProfileKey("test sonar way", "java")
		})
	})
	Describe("Test SetDefault in api/qualityprofiles", func() {
		It("Should be ok", func() {
			opt := &QualityProfilesSetDefaultOption{
				Key:            testQualityProfile,
				Language:       "",
				QualityProfile: "",
			}
			assertNoBody(client.QualityProfiles.SetDefault(opt))

			opt.Language = "java"
			opt.Key = ""
			opt.QualityProfile = "Sonar way"
			assertNoBody(client.QualityProfiles.SetDefault(opt))
		})
	})
})
