package filtering_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/libs/filtering"
)

func TestSearch(t *testing.T) {
	RegisterFailHandler(Fail)
	if !testing.Short() {
		RunSpecs(t, "Libs > Search")
	}
}

var _ = Describe("Search.Where()", func() {
	It("When there is blank", func() {
		s := filtering.Search{Raw: ""}
		sql := s.Where("")
		Expect(sql).To(Equal(""))
	})

	It("When there is only one field", func() {
		s := filtering.Search{Raw: "(id[eq]:'1')"}
		sql := s.Where("")
		Expect(sql).To(Equal("id = '1'"))
	})

	It("When there are multi fields separated by comma", func() {
		s := filtering.Search{Raw: "(id[gt]:'1',title[ne]:'hello')"}
		sql := s.Where("a.")
		Expect(sql).To(Equal("a.id > '1' OR a.title <> 'hello'"))
	})

	It("When value contains %", func() {
		s := filtering.Search{Raw: "(id[gt]:'1',title[contains]:'%hello%')"}
		sql := s.Where("a.")
		Expect(sql).To(Equal("a.id > '1' OR a.title LIKE '%hello%'"))
	})

	It("When uses AND & OR", func() {
		s := filtering.Search{Raw: "(id[gt]:'1'+created_at[gt]:'2019-01-01'+created_at[lt]:'2019-10-10',title[contains]:'%hello%')"}
		sql := s.Where("a.")
		Expect(sql).To(Equal("a.id > '1' OR a.title LIKE '%hello%' AND a.created_at > '2019-01-01' AND a.created_at < '2019-10-10'"))
	})

	Describe("Prevent Where Injection", func() {
		It("Returns empty when contains =", func() {
			s := filtering.Search{Raw: `(id[gt]:'' OR 1=1)`}
			sql := s.Where("a.")
			Expect(sql).To(Equal(""))
		})
	})
})
