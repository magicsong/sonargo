package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var testAnalysisKey = ""

func searchAnalysis() string {
	opt := &ProjectAnalysesSearchOption{
		Category: "",
		From:     "",
		P:        0,
		Project:  AnalyzedProject,
		Ps:       0,
		To:       "",
	}
	v, resp, err := client.ProjectAnalyses.Search(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Analyses).NotTo(BeNil())
	return v.Analyses[0].Key
}

func createEvent(name string) string {
	opt := &ProjectAnalysesCreateEventOption{
		Analysis: testAnalysisKey,
		Category: OtherEventCatagory,
		Name:     name,
	}
	v, resp, err := client.ProjectAnalyses.CreateEvent(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Event.Name).To(Equal(name))
	return v.Event.Key
}

func deleteEvent(key string) {
	opt := &ProjectAnalysesDeleteEventOption{Event: key}
	resp, err := client.ProjectAnalyses.DeleteEvent(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "not found")).To(BeTrue())
		return
	}
	assertNoBody(resp, err)
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {
		testAnalysisKey = searchAnalysis()
	})
	JustBeforeEach(func() {})
	Describe("Test CreateEvent in api/project_analyses", func() {
		testEventName := "test-event"
		var testEventKey string
		It("Should be ok to create event", func() {
			testEventKey = createEvent(testEventName)
		})
		AfterEach(func() {
			if testEventKey != "" {
				deleteEvent(testEventKey)
			}
		})
	})
	Describe("Test Delete in api/project_analyses", func() {
		It("Should be ok to delete analysis", func() {
			// right now we cannot delete an exsit analysis because creating an analysis is too long
			opt := &ProjectAnalysesDeleteOption{Analysis: "whatever"}
			resp, err := client.ProjectAnalyses.Delete(opt)
			if err != nil {
				Expect(strings.Contains(err.Error(), "not found")).To(BeTrue())
				return
			}
			assertNoBody(resp, err)
		})
	})
	Describe("Test DeleteEvent in api/project_analyses", func() {
		It("Should be ok", func() {
			key := createEvent("TestEvent")
			deleteEvent(key)
		})
	})
	Describe("Test Search in api/project_analyses", func() {
		It("Should be ok", func() {
			searchAnalysis()
		})
	})
	Describe("Test UpdateEvent in api/project_analyses", func() {
		var key string
		It("Should be ok", func() {
			key = createEvent("updateEventTest")
			opt := &ProjectAnalysesUpdateEventOption{
				Event: key,
				Name:  "updateEventTest2",
			}
			v, resp, err := client.ProjectAnalyses.UpdateEvent(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Event.Name).To(Equal("updateEventTest2"))
		})
		AfterEach(func() {
			if key != "" {
				deleteEvent(key)
			}
		})
	})
})
