package structs_test

import (
	"net/url"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/src/structs"
)

func TestSearch(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Client > Elearning > Albums")
}

var _ = Describe("Search", func() {

	It("Should set order and sortby asc", func() {
		v := url.Values{}
		v.Add("sort", "+created_at")
		search := structs.NewSearch(v)
		Expect(search.Pagination.Order).To(Equal("created_at"))
		Expect(search.Pagination.SortBy).To(Equal("asc"))
	})

	It("Should set order and sortby desc", func() {
		v := url.Values{}
		v.Add("sort", "-created_at")
		search := structs.NewSearch(v)
		Expect(search.Pagination.Order).To(Equal("created_at"))
		Expect(search.Pagination.SortBy).To(Equal("desc"))
	})

	It("Should set raw search value when it contains ()", func() {
		param := "(id[gt]:'1')"
		v := url.Values{}
		v.Add("search", param)
		search := structs.NewSearch(v)
		Expect(search.Raw).To(Equal(param))
	})

	It("Should dont set raw search value when it does not contains ()", func() {
		param := "id[gt]:'1'"
		v := url.Values{}
		v.Add("search", param)
		search := structs.NewSearch(v)
		Expect(search.Raw).To(BeEmpty())
	})
})
