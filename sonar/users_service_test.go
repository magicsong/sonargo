package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func createTestUser(name, email, password string) string {
	opt := &UsersCreateOption{
		Email:      email,
		Local:      "true",
		Login:      name,
		Name:       name,
		Password:   password,
		ScmAccount: "",
	}
	v, resp, err := client.Users.Create(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "already exists")).To(BeTrue())
		return name
	}
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.User.Active).To(BeTrue())
	return v.User.Login
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test ChangePassword in api/users", func() {
		It("Should be ok", func() {
			old := "passw0rd"
			new := "newPassword"
			opt := &UsersChangePasswordOption{
				Login:            testUser1,
				Password:         new,
				PreviousPassword: old,
			}
			assertNoBody(client.Users.ChangePassword(opt))
			By("Change Back")
			opt.Password = old
			opt.PreviousPassword = new
			assertNoBody(client.Users.ChangePassword(opt))
		})
	})

	Describe("Test Create in api/users (This is skipped because sonarqube doesn't support delete user)", func() {
		It("Should be ok to create a user", func() {
			createTestUser("magicsong", "magicsong@yunify.com", "passw0rd")
		})
	})
	Describe("Test Deactivate in api/users", func() {
		It("Should return false when deactive a user", func() {
			opt := &UsersDeactivateOption{Login: testUser1}
			v, resp, err := client.Users.Deactivate(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.User.Active).To(BeFalse())
			//active
			testUser1 = createTestUser("magicsong", "magicsong@yunify.com", "passw0rd")
		})
	})
	Describe("Test Groups in api/users", func() {
		It("Should be ok", func() {
			opt := &UsersGroupsOption{
				Login:    testUser1,
				P:        0,
				Ps:       0,
				Q:        "",
				Selected: "",
			}
			v, resp, err := client.Users.Groups(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Groups).NotTo(BeNil())
		})
	})
	Describe("Test Search in api/users", func() {
		It("Should be ok", func() {
			opt := &UsersSearchOption{
				F:  "",
				P:  0,
				Ps: 0,
				Q:  testUser1,
			}
			v, resp, err := client.Users.Search(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(len(v.Users)).NotTo(BeZero())
		})
	})
	Describe("Test Update in api/users", func() {
		It("Should be ok", func() {
			old := "magicsong"
			new := "magicsong1"
			opt := &UsersUpdateOption{
				Email:      "",
				Login:      testUser1,
				Name:       new,
				ScmAccount: "",
			}
			v, resp, err := client.Users.Update(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.User.Name).To(Equal(new))

			opt.Name = old
			v, resp, err = client.Users.Update(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.User.Name).To(Equal(old))
		})
	})
})
