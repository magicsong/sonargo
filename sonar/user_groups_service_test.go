package sonargo_test

import (
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func createGroup(name string) int {
	opt := &UserGroupsCreateOption{
		Description: "description",
		Name:        name,
	}
	v, resp, err := client.UserGroups.Create(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "already exists")).To(BeTrue())
		return searchGroup(name)
	}
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(v.Group.Name).To(Equal(name))
	return v.Group.ID
}

func deleteGroup(name string) {
	//delete
	opt := &UserGroupsDeleteOption{
		Name: name,
	}
	resp, err := client.UserGroups.Delete(opt)
	if err != nil {
		Expect(strings.Contains(err.Error(), "404")).To(BeTrue())
		return
	}
	assertNoBody(resp, err)
}

func searchGroup(name string) int {
	opt := &UserGroupsSearchOption{
		F: "",
		Q: name,
	}
	v, resp, err := client.UserGroups.Search(opt)
	Expect(err).ShouldNot(HaveOccurred())
	Expect(resp.StatusCode).To(Equal(200))
	Expect(len(v.Groups)).NotTo(BeZero())
	return v.Groups[0].ID
}

var _ = Describe("SonarCLI integration test", func() {
	BeforeEach(func() {})
	JustBeforeEach(func() {})
	Describe("Test AddUser in api/user_groups", func() {
		It("Should be ok", func() {
			opt := &UserGroupsAddUserOption{
				Id:    0,
				Login: "magicsong",
				Name:  "test1group",
			}
			assertNoBody(client.UserGroups.AddUser(opt))

		})
		AfterEach(func() {
			opt := &UserGroupsRemoveUserOption{
				Id:    0,
				Login: "magicsong",
				Name:  "test1group",
			}
			assertNoBody(client.UserGroups.RemoveUser(opt))
		})

	})
	Describe("Test Create/Delete in api/user_groups", func() {
		testGroupName := "testcreategroup"
		It("Should be ok to create/delete a group", func() {
			createGroup(testGroupName)
		})
		AfterEach(func() {
			deleteGroup(testGroupName)
		})
	})
	Describe("Test Search in api/user_groups", func() {
		It("Should be ok", func() {
			searchGroup("test1group")
		})
	})
	Describe("Test Update in api/user_groups", func() {
		It("Should be ok", func() {
			old := "oldDesciption"
			new := "newDesciption"
			opt := &UserGroupsUpdateOption{
				Description: new,
				Id:          testGroup1,
				Name:        "",
			}
			v, resp, err := client.UserGroups.Update(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Group.Description).To(Equal(new))

			//change back

			opt.Description = old
			v, _, err = client.UserGroups.Update(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(v.Group.Description).To(Equal(old))

		})
	})
	Describe("Test Users in api/user_groups", func() {
		It("Should be ok", func() {
			opt := &UserGroupsUsersOption{
				Id:       testGroup1,
				Name:     "",
				Q:        "",
				Selected: "",
			}
			v, resp, err := client.UserGroups.Users(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v).NotTo(BeNil())
		})
	})
})
