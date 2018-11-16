package sonargo_test

import (
	"fmt"
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func getGateIDByName(name string) int {
	opt := &QualitygatesShowOption{
		Name:         name,
		Organization: "",
	}
	v, resp, err := client.Qualitygates.Show(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Name).To(Equal(name))
	return v.ID
}

func createQualityGate(name string) int {
	opt := &QualitygatesCreateOption{
		Name:         name,
		Organization: "",
	}
	v, resp, err := client.Qualitygates.Create(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "already been taken")).To(BeTrue())
		return getGateIDByName(name)
	}
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Name).To(Equal(name))
	return v.ID
}

func deleteQualityGate(id int) {
	opt := &QualitygatesDestroyOption{
		Id:           id,
		Organization: "",
	}
	resp, err := client.Qualitygates.Destroy(opt)
	if err != nil {
		fmt.Println(err.Error())
		Expect(strings.Contains(err.Error(), "not found")).To(BeTrue())
		return
	}
	assertNoBody(resp, err)
}

func createCondition() int {
	opt := &QualitygatesCreateConditionOption{
		Error:        "10",
		GateId:       testQualityGate,
		Metric:       "blocker_violations",
		Op:           "GT",
		Organization: "",
		Period:       "",
		Warning:      "",
	}
	v, resp, err := client.Qualitygates.CreateCondition(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Op).To(Equal("GT"))
	return v.ID
}

func deleteCondition(id int) {
	opt := &QualitygatesDeleteConditionOption{
		ConditionID:  id,
		Organization: "",
	}
	assertNoBody(client.Qualitygates.DeleteCondition(opt))

}
func setDefaultGate(id int) {
	opt := &QualitygatesSetAsDefaultOption{
		Id:           id,
		Organization: "",
	}
	assertNoBody(client.Qualitygates.SetAsDefault(opt))
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Copy in api/qualitygates", func() {
		It("Should be ok to copy a quality gate", func() {
			opt := &QualitygatesCopyOption{
				Id:           testQualityGate,
				Name:         "testQualityGate-copy2",
				Organization: "",
			}
			v, resp, err := client.Qualitygates.Copy(opt)
			if err != nil {
				Expect(strings.Contains(err.Error(), "already been taken")).To(BeTrue())
				deleteQualityGate(getGateIDByName("testQualityGate-copy2"))
				return
			}
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Name).To(Equal("testQualityGate-copy2"))
			//delete
			deleteQualityGate(v.ID)
		})
	})
	Describe("Test Create in api/qualitygates", func() {
		It("Should be ok to create a gate", func() {
			id := createQualityGate("testCreateQualityGate")
			deleteQualityGate(id)
		})
	})
	Describe("Test CreateCondition in api/qualitygates", func() {
		It("Should be ok to create/delete a condition", func() {
			id := createCondition()
			deleteCondition(id)
		})
	})
	Describe("Test GetByProject in api/qualitygates", func() {
		It("Should be ok", func() {
			opt := &QualitygatesGetByProjectOption{
				Organization: "",
				Project:      testProject,
			}
			v, resp, err := client.Qualitygates.GetByProject(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test List in api/qualitygates", func() {
		It("Should be ok to list qualitygates", func() {
			opt := &QualitygatesListOption{Organization: ""}
			v, resp, err := client.Qualitygates.List(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test ProjectStatus in api/qualitygates", func() {
		It("Should be ok", func() {
			opt := &QualitygatesProjectStatusOption{
				AnalysisId: searchAnalysis(),
				ProjectId:  "whateveris",
				ProjectKey: AnalyzedProject,
			}
			v, resp, err := client.Qualitygates.ProjectStatus(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.ProjectStatus).NotTo(BeNil())
		})
	})
	Describe("Test Rename in api/qualitygates", func() {
		It("Should be ok", func() {
			id := createQualityGate("testRename")
			opt := &QualitygatesRenameOption{
				Id:           id,
				Name:         "newName",
				Organization: "",
			}
			v, resp, err := client.Qualitygates.Rename(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Name).To(Equal("newName"))
			deleteQualityGate(id)
		})
	})
	Describe("Test Search in api/qualitygates", func() {
		It("Should be ok to use search", func() {
			opt := &QualitygatesSearchOption{
				GateId:       testQualityGate,
				Organization: "",
				Page:         "",
				PageSize:     "",
				Query:        "",
				Selected:     "",
			}
			v, resp, err := client.Qualitygates.Search(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
	Describe("Test Select/Deselect in api/qualitygates", func() {
		It("Should be ok", func() {
			opt := &QualitygatesSelectOption{
				GateId:       testQualityGate,
				Organization: "",
				ProjectKey:   testProject,
				ProjectID:    "",
			}
			assertNoBody(client.Qualitygates.Select(opt))
			//deselect
			optDeselect := &QualitygatesDeselectOption{
				Organization: "",
				ProjectId:    "",
				ProjectKey:   testProject,
			}
			assertNoBody(client.Qualitygates.Deselect(optDeselect))
		})
	})
	Describe("Test SetAsDefault in api/qualitygates", func() {
		It("Should be ok", func() {
			setDefaultGate(testQualityGate)
			//change back
			setDefaultGate(1)
		})
	})
	Describe("Test Show in api/qualitygates", func() {
		It("Should be ok to show a qualitygate", func() {
			opt := &QualitygatesShowOption{
				Id:           testQualityGate,
				Name:         "",
				Organization: "",
			}
			v, resp, err := client.Qualitygates.Show(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.ID).To(Equal(testQualityGate))
		})
	})
	Describe("Test UpdateCondition in api/qualitygates", func() {
		It("Should be ok to update condition", func() {
			id := createCondition()
			opt := &QualitygatesUpdateConditionOption{
				Error:        "25",
				Id:           id,
				Metric:       "blocker_violations",
				Op:           "GT",
				Organization: "",
				Period:       "",
				Warning:      "",
			}
			v, resp, err := client.Qualitygates.UpdateCondition(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Error).To(Equal("25"))
			deleteCondition(id)
		})
	})
})
