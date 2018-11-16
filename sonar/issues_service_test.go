package sonargo_test

import (
	"sort"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func getIssue() string {
	opt := &IssuesSearchOption{
		AdditionalFields:   "",
		Asc:                "",
		Assigned:           "",
		Assignees:          "",
		Authors:            "",
		ComponentKeys:      AnalyzedProject,
		ComponentRootUuids: "",
		ComponentRoots:     "",
		ComponentUuids:     "",
		Components:         "",
		CreatedAfter:       "",
		CreatedAt:          "",
		CreatedBefore:      "",
		CreatedInLast:      "",
		Issues:             "",
		Languages:          "",
		P:                  "",
		Ps:                 "",
		Resolutions:        "",
		Resolved:           "",
		Rules:              "",
		S:                  "",
		Severities:         "",
		SinceLeakPeriod:    "",
		Statuses:           "OPEN",
		Tags:               "",
		Types:              "",
	}
	v, resp, err := client.Issues.Search(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Issues).NotTo(BeNil())
	return v.Issues[1].Key
}
func addComment(comment string) string {
	opt := &IssuesAddCommentOption{
		Issue: testIssueKey,
		Text:  comment,
	}
	v, resp, err := client.Issues.AddComment(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	l := len(v.Issue.Comments)
	Expect(l).ShouldNot(BeZero())
	Expect(v.Issue.Comments[l-1].Markdown).To(Equal(comment))
	return v.Issue.Comments[l-1].Key
}

func removeComment(id string) {
	removeOpt := &IssuesDeleteCommentOption{Comment: id}
	issue, resp, err := client.Issues.DeleteComment(removeOpt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(issue.Issue.Key).To(Equal(testIssueKey))
}

var _ = Describe("SonarCLI integration test", func() {
	Describe("Test Add-Remove Comment in api/issues", func() {
		It("Should be ok", func() {
			removeComment(addComment("Haha, I write a comment"))
		})
	})
	Describe("Test Assign/Remove in api/issues", func() {
		It("Should be ok use '_me'", func() {
			opt := &IssuesAssignOption{
				Assignee: "_me",
				Issue:    testIssueKey,
			}
			v, resp, err := client.Issues.Assign(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Issue.Assignee).To(Equal("admin"))

			opt.Assignee = ""
			v, resp, err = client.Issues.Assign(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Issue.Assignee).To(BeZero())
		})
	})
	Describe("Test Authors in api/issues", func() {
		It("Should be ok", func() {
			opt := &IssuesAuthorsOption{
				Ps: 0,
				Q:  "",
			}
			v, resp, err := client.Issues.Authors(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).ShouldNot(BeNil())
		})
	})
	Describe("Test BulkChange in api/issues", func() {
		It("Should get 1 change", func() {
			opt := &IssuesBulkChangeOption{
				AddTags:           "",
				Assign:            "",
				Comment:           "",
				DoTransition:      "",
				Issues:            testIssueKey,
				Plan:              "",
				RemoveTags:        "",
				SendNotifications: "",
				SetSeverity:       "CRITICAL",
				SetType:           "",
			}
			v, resp, err := client.Issues.BulkChange(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Total).To(Equal(1))
		})
	})
	Describe("Test Changelog in api/issues", func() {
		It("Should be ok", func() {
			opt := &IssuesChangelogOption{Issue: testIssueKey}
			v, resp, err := client.Issues.Changelog(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test DoTransition in api/issues", func() {
		It("Should be ok", func() {
			//try to reopen ignore the error
			opt := &IssuesDoTransitionOption{
				Issue:      testIssueKey,
				Transition: "confirm",
			}
			v, resp, err := client.Issues.DoTransition(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Issue.Status).To(Equal("CONFIRMED"))
			//Do the test
			opt.Transition = "unconfirm"
			client.Issues.DoTransition(opt)
		})
	})
	Describe("Test EditComment in api/issues", func() {
		It("Should be ok to edit comment", func() {
			commentId := addComment("this comment need to be modified")
			modifiedText := "this comment is modified"
			opt := &IssuesEditCommentOption{
				Comment: commentId,
				Text:    modifiedText,
			}
			v, resp, err := client.Issues.EditComment(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Issue.Comments[len(v.Issue.Comments)-1].Markdown).To(Equal(modifiedText))
		})
	})
	Describe("Test Search in api/issues", func() {
		It("Should be ok to use search", func() {
			getIssue()
		})
	})
	Describe("Test SetSeverity in api/issues", func() {
		It("Should be ok to set severity", func() {
			opt := &IssuesSetSeverityOption{
				Issue:    testIssueKey,
				Severity: SeverityBLOCKER,
			}
			v, resp, err := client.Issues.SetSeverity(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Issue.Key).To(Equal(testIssueKey))
		})
	})
	Describe("Test SetTags in api/issues", func() {
		It("Should be ok", func() {
			opt := &IssuesSetTagsOption{
				Issue: testIssueKey,
				Tags:  "security,cwe,misra-c",
			}
			v, resp, err := client.Issues.SetTags(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			r := []string{"security", "cwe", "misra-c"}
			sort.Strings(r)
			Expect(v.Issue.Tags).To(Equal(r))
		})
	})
	Describe("Test SetType in api/issues", func() {
		It("Should be ok", func() {
			opt := &IssuesSetTypeOption{
				Issue: testIssueKey,
				Type:  IssueTypeBug,
			}
			v, resp, err := client.Issues.SetType(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Issue.Type).To(Equal(IssueTypeBug))
		})
	})
	Describe("Test Tags in api/issues", func() {
		It("Should be ok", func() {
			opt := &IssuesTagsOption{
				Ps: 0,
				Q:  "misra-c",
			}
			v, resp, err := client.Issues.Tags(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
})
