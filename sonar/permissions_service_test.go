package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func getTemplateIdByName(name string) string {
	opt := &PermissionsSearchTemplatesOption{Q: name}
	v, resp, err := client.Permissions.SearchTemplates(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v).NotTo(BeZero())
	return v.PermissionTemplates[0].ID
}
func createTemplate(name string) string {
	opt := &PermissionsCreateTemplateOption{
		Description:       "Test Template",
		Name:              name,
		ProjectKeyPattern: ".*sonar.*",
	}
	v, resp, err := client.Permissions.CreateTemplate(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "already exists")).To(BeTrue())
	} else {
		Expect(err).ShouldNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(200))
		Expect(v.PermissionTemplate.Name).To(Equal(name))
	}
	return name
}
func deleteTemplate(name string) {
	opt := &PermissionsDeleteTemplateOption{
		TemplateId:   "",
		TemplateName: name,
	}
	assertNoBody(client.Permissions.DeleteTemplate(opt))

}

func setDefaultTemplate(name string) {
	opt := &PermissionsSetDefaultTemplateOption{
		Qualifier:    "",
		TemplateId:   "",
		TemplateName: name,
	}
	assertNoBody(client.Permissions.SetDefaultTemplate(opt))
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test AddGroup in api/permissions", func() {
		It("Should be ok set global permission", func() {
			opt := &PermissionsAddGroupOption{
				GroupId:    testGroup1,
				GroupName:  "",
				Permission: "scan",
				ProjectId:  "",
				ProjectKey: "",
			}
			assertNoBody(client.Permissions.AddGroup(opt))
		})

		It("Should be ok set/unset project permission", func() {
			opt := &PermissionsAddGroupOption{
				GroupId:    testGroup1,
				GroupName:  "",
				Permission: "issueadmin",
				ProjectId:  "",
				ProjectKey: testProject,
			}
			assertNoBody(client.Permissions.AddGroup(opt))
			By("Change Back")
			optRemove := PermissionsRemoveGroupOption(*opt)
			assertNoBody(client.Permissions.RemoveGroup(&optRemove))
		})
	})
	Describe("Test AddGroupToTemplate in api/permissions", func() {
		It("Should be ok to add/remove group to template", func() {
			opt := &PermissionsAddGroupToTemplateOption{
				GroupId:      testGroup2,
				GroupName:    "",
				Permission:   "issueadmin",
				TemplateId:   "",
				TemplateName: testTemplate,
			}
			assertNoBody(client.Permissions.AddGroupToTemplate(opt))

			By("Change Back")
			optRemove := PermissionsRemoveGroupFromTemplateOption(*opt)
			assertNoBody(client.Permissions.RemoveGroupFromTemplate(&optRemove))

		})
	})
	Describe("Test AddProjectCreatorToTemplate in api/permissions", func() {
		It("Should be ok to add project creator to a template", func() {
			opt := &PermissionsAddProjectCreatorToTemplateOption{
				Permission:   "issueadmin",
				TemplateId:   "",
				TemplateName: testTemplate,
			}
			assertNoBody(client.Permissions.AddProjectCreatorToTemplate(opt))
			By("Change Back")
			optRemove := PermissionsRemoveProjectCreatorFromTemplateOption(*opt)
			assertNoBody(client.Permissions.RemoveProjectCreatorFromTemplate(&optRemove))
		})
	})
	Describe("Test AddUser in api/permissions", func() {
		It("Should be ok add/remove user in permissions", func() {
			opt := &PermissionsAddUserOption{
				Login:      testUser1,
				Permission: "issueadmin",
				ProjectId:  "",
				ProjectKey: testProject,
			}
			assertNoBody(client.Permissions.AddUser(opt))
			By("Change Back")
			optRemove := PermissionsRemoveUserOption(*opt)
			assertNoBody(client.Permissions.RemoveUser(&optRemove))
		})
	})
	Describe("Test AddUserToTemplate in api/permissions", func() {
		It("Should be ok", func() {
			opt := &PermissionsAddUserToTemplateOption{
				Login:        testUser1,
				Permission:   "issueadmin",
				TemplateId:   "",
				TemplateName: testTemplate,
			}
			assertNoBody(client.Permissions.AddUserToTemplate(opt))
			optRemove := PermissionsRemoveUserFromTemplateOption(*opt)
			assertNoBody(client.Permissions.RemoveUserFromTemplate(&optRemove))
		})
	})
	Describe("Test ApplyTemplate in api/permissions", func() {
		It("Should be ok", func() {
			opt := &PermissionsApplyTemplateOption{
				ProjectId:    "",
				ProjectKey:   testProject,
				TemplateId:   "",
				TemplateName: testTemplate,
			}
			assertNoBody(client.Permissions.ApplyTemplate(opt))
		})
	})
	Describe("Test BulkApplyTemplate in api/permissions", func() {
		It("Should be ok", func() {
			opt := &PermissionsBulkApplyTemplateOption{
				AnalyzedBefore:    "",
				OnProvisionedOnly: "",
				Projects:          testProject,
				Q:                 "",
				Qualifiers:        "",
				TemplateId:        "",
				TemplateName:      testTemplate,
			}
			assertNoBody(client.Permissions.BulkApplyTemplate(opt))
		})
	})
	Describe("Test CreateTemplate in api/permissions", func() {
		It("Should be ok to create/delete template", func() {
			createTemplate("testcreatedelete")
			deleteTemplate("testcreatedelete")
		})
	})

	Describe("Test SearchTemplates in api/permissions", func() {
		It("Should be ok", func() {
			getTemplateIdByName(testTemplate)
		})
	})
	Describe("Test SetDefaultTemplate in api/permissions", func() {
		It("Should be ok", func() {
			setDefaultTemplate(testTemplate)
			setDefaultTemplate("Default template")
		})
	})
	Describe("Test UpdateTemplate in api/permissions", func() {
		It("Should be ok", func() {
			opt := &PermissionsUpdateTemplateOption{
				Description:       "",
				Id:                getTemplateIdByName(testTemplate),
				Name:              testTemplate,
				ProjectKeyPattern: "whatever",
			}
			v, resp, err := client.Permissions.UpdateTemplate(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.PermissionTemplate.ProjectKeyPattern).To(Equal("whatever"))
			By("Change back")

			opt.ProjectKeyPattern = ".*sonar.*"
			_, resp, err = client.Permissions.UpdateTemplate(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
		})
	})
})
