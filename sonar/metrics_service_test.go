package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func createMetric(key string) string {
	optCreate := &MetricsCreateOption{
		Description: "Test custom metric",
		Domain:      "Tests",
		Key:         key,
		Name:        key,
		Type:        MetricTypeINT,
	}
	v, resp, err := client.Metrics.Create(optCreate)
	if err != nil {
		Expect(strings.Contains(err.Error(), "already exist")).To(BeTrue())
		opt := &MetricsSearchOption{
			F:        "",
			IsCustom: "true",
			P:        0,
			Ps:       0,
		}
		v, resp, err := client.Metrics.Search(opt)
		Expect(err).ShouldNot(HaveOccurred())
		Expect(resp.StatusCode).To(Equal(200))
		for _, item := range v.Metrics {
			if item.Key == key {
				return item.ID
			}
		}
		return v.Metrics[0].ID
	}
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Key).To(Equal(key))
	return v.ID
}

func deleteMetric(key string) {
	opt := &MetricsDeleteOption{
		Ids:  "",
		Keys: key,
	}
	assertNoBody(client.Metrics.Delete(opt))
}

var _ = Describe("SonarCLI integration test", func() {
	testKey := "team_size1"
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Create/Delete in api/metrics", func() {
		It("Should be ok", func() {
			createMetric(testKey)
			deleteMetric(testKey)
		})
	})
	Describe("Test Domains in api/metrics", func() {
		It("Should be ok", func() {
			//create before test
			createMetric(testKey)
			v, resp, err := client.Metrics.Domains()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(len(v.Domains)).NotTo(BeZero())
			deleteMetric(testKey)
		})
	})
	Describe("Test Search in api/metrics", func() {
		It("Should be ok", func() {
			opt := &MetricsSearchOption{
				F:        "",
				IsCustom: "",
				P:        0,
				Ps:       0,
			}
			v, resp, err := client.Metrics.Search(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(len(v.Metrics)).NotTo(BeZero())
		})
	})
	Describe("Test Types in api/metrics", func() {
		It("Should be ok", func() {
			v, resp, err := client.Metrics.Types()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(len(v.Types)).NotTo(BeZero())
		})
	})
	Describe("Test Update in api/metrics", func() {
		updateKey := testKey + "0"
		It("Should be ok", func() {
			//create
			id := createMetric(updateKey)
			testUpdate := "Hello World"
			opt := &MetricsUpdateOption{
				Description: testUpdate,
				Domain:      "",
				ID:          id,
				Key:         "",
				Name:        "",
				Type:        "",
			}
			v, resp, err := client.Metrics.Update(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Description).To(Equal(testUpdate))
			deleteMetric(updateKey)
		})
	})
})
