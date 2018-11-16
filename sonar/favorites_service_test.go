package sonargo_test

import (
	"fmt"
	"strings"

	. "github.com/magicsong/sonargo/sonar"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SonarCLI integration test", func() {
	addTestFavorite := func() {
		opt := &FavoritesAddOption{Component: testProject}
		resp, err := client.Favorites.Add(opt)
		if err != nil && strings.Contains(err.Error(), "already a favorite") {
			fmt.Println("Already Favorites")
			return
		}
		assertNoBody(resp, err)
	}
	BeforeEach(func() {

	})
	AfterEach(func() {
		removeOpt := &FavoritesRemoveOption{Component: testProject}
		client.Favorites.Remove(removeOpt)
	})
	Describe("Test Add-Remove in api/favorites", func() {
		It("Should be ok", func() {
			addTestFavorite()
			removeOpt := &FavoritesRemoveOption{Component: testProject}
			assertNoBody(client.Favorites.Remove(removeOpt))
		})
		It("Should get error", func() {
			opt := &FavoritesAddOption{Component: "sonargo-not-exist"}
			resp, err := client.Favorites.Add(opt)
			Expect(strings.Contains(err.Error(), "'sonargo-not-exist' not found")).To(BeTrue())
			Expect(resp.StatusCode).To(Equal(404))
		})
	})
	Describe("Test Search in api/favorites", func() {
		It("Should get the favorites", func() {
			addTestFavorite()
			opt := &FavoritesSearchOption{
				P:  0,
				Ps: 0,
			}
			v, resp, err := client.Favorites.Search(opt)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.StatusCode).To(Equal(200))
			Expect(v.Favorites[0].Key).To(Equal(testProject))
		})
	})
})
