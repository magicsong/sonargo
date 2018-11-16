package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func createUserToken(name string) string {
	opt := &UserTokensGenerateOption{
		Login: testUser1,
		Name:  name,
	}
	v, resp, err := client.UserTokens.Generate(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "already exists")).To(BeTrue())
		return ""
	}
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Name).To(Equal(name))
	return v.Token
}

func revokeToken(name string) {
	opt := &UserTokensRevokeOption{
		Login: "",
		Name:  name,
	}
	assertNoBody(client.UserTokens.Revoke(opt))
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test Generate in api/user_tokens", func() {
		It("Should be ok to add/remove token", func() {
			createUserToken("testToken")
			revokeToken("testToken")
		})
	})

	Describe("Test Search in api/user_tokens", func() {
		It("Should be ok", func() {
			opt := &UserTokensSearchOption{Login: testUser1}
			v, resp, err := client.UserTokens.Search(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.UserTokens).NotTo(BeNil())
		})
	})
})
