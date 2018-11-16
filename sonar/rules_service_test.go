package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func createRule(name string) string {
	opt := &RulesCreateOption{
		CustomKey:           name,
		ManualKey:           "",
		MarkdownDescription: "This is the description of `rule`",
		Name:                name,
		Params:              "",
		PreventReactivation: "",
		Severity:            SeverityINFO,
		Status:              "BETA",
		TemplateKey:         "squid:S3546",
		Type:                IssueTypeCodeSmell,
	}
	v, resp, err := client.Rules.Create(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "already exists")).To(BeTrue())
		return getRuleKeyByName(name)
	}
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Rule.Name).To(Equal(name))
	return v.Rule.Key
}
func getRuleKeyByName(name string) string {
	opt := &RulesSearchOption{
		Activation:       "",
		ActiveSeverities: "",
		Asc:              "",
		AvailableSince:   "",
		Facets:           "",
		Inheritance:      "",
		IsTemplate:       "",
		Languages:        "",
		P:                "",
		Ps:               "",
		Q:                name,
		Qprofile:         "",
		Repositories:     "",
		RuleKey:          "",
		S:                "",
		Severities:       "",
		Statuses:         "",
		Tags:             "",
		TemplateKey:      "",
		Types:            "",
	}
	v, resp, err := client.Rules.Search(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Rules).NotTo(BeNil())
	return v.Rules[0].Key
}
func deleteRule(key string) {
	opt := &RulesDeleteOption{Key: key}
	resp, err := client.Rules.Delete(opt)
	Expect(err).ShouldNot(HaveOccurred())
	//the sonarqube api has no specifications
	Expect(resp.StatusCode).To(Equal(200))
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Create/delete in api/rules", func() {
		It("Should be ok", func() {
			key := createRule("magic_test_rule")
			deleteRule(key)
		})
	})
	Describe("Test Repositories in api/rules", func() {
		It("Should be ok to list repositories", func() {
			opt := &RulesRepositoriesOption{
				Language: "",
				Q:        "",
			}
			v, resp, err := client.Rules.Repositories(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test Search in api/rules", func() {
		It("Should be ok to search template rules", func() {

		})
	})
	Describe("Test Show in api/rules", func() {
		It("Should be ok", func() {
			opt := &RulesShowOption{
				Actives: "",
				Key:     "squid:S3546",
			}
			v, resp, err := client.Rules.Show(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Rule.Key).To(Equal("squid:S3546"))
		})
	})
	Describe("Test Tags in api/rules", func() {
		It("Should be ok", func() {
			opt := &RulesTagsOption{
				Ps: 0,
				Q:  "",
			}
			v, resp, err := client.Rules.Tags(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test Update in api/rules", func() {
		It("Should be ok", func() {
			key := createRule("magic_test_rule")
			opt := &RulesUpdateOption{
				DebtRemediationFnOffset: "",
				DebtRemediationFnType:   "",
				DebtRemediationFyCoeff:  "",
				DebtSubCharacteristic:   "",
				Key:                        key,
				MarkdownDescription:        "",
				MarkdownNote:               "",
				Name:                       "",
				Params:                     "",
				RemediationFnBaseEffort:    "",
				RemediationFnType:          "",
				RemediationFyGapMultiplier: "",
				Severity:                   SeverityBLOCKER,
				Status:                     "",
				Tags:                       "",
			}
			v, resp, err := client.Rules.Update(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Rule.Severity).To(Equal(SeverityBLOCKER))
			deleteRule(key)
		})
	})
})
