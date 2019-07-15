package repositories_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/elearning/src/repositories"
	"backend/apps/elearning/src/structs"
	"backend/libs/configuration"
	"backend/libs/databases"
	"backend/libs/errors"
	testutil "backend/libs/testing"
)

func TestShelfRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	if !testing.Short() {
		RunSpecs(t, "Repository > Shelf")
	}
}

var _ = Describe("ShelfRepository", func() {
	var repository repositories.Repository
	var db databases.Database

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "elearning.sql")
		repository = repositories.NewShelfRepository(db)
	})

	Describe("ReadOne", func() {
		id := uint64(2)
		It("Should returns item by ID", func() {
			result, err := repository.ReadOne(structs.Shelf{ID: id})
			Expect(err).ShouldNot(HaveOccurred())

			shelf := result.(structs.Shelf)
			Expect(shelf.ID).To(Equal(id))
			Expect(shelf.Title).To(Equal("Estatuto e Ética da OAB"))
		})
	})

	Describe("ReadAll", func() {
		It("Should returns paginated items", func() {
			search := structs.Search{
				Pagination: structs.Pagination{
					PerPage: 10,
					Page:    1,
				},
			}
			result, err := repository.ReadAll(search)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Pagination.Total).To(Equal(int32(3)))
			items := result.Items.([]structs.Shelf)
			Expect(items).To(HaveLen(3))
			Expect(items[0].ID).To(Equal(uint64(1)))
			Expect(items[1].ID).To(Equal(uint64(2)))
		})
	})

	Describe("Create", func() {
		var item structs.Shelf

		BeforeEach(func() {
			item = structs.Shelf{
				Title: "TRT",
			}
		})

		It("Should returns error when title length < 1", func() {
			item.Title = ""
			_, err := repository.Create(item)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
		})

		It("Should returns error when title already exists ignoring case", func() {
			item.Title = "estatuto e ética da oab"
			_, err := repository.Create(item)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
		})

		It("Should returns its own ID after it was created", func() {
			result, err := repository.Create(item)
			Expect(err).ShouldNot(HaveOccurred())
			saved := result.(structs.Shelf)
			Expect(saved.ID).To(Equal(uint64(1000)))
		})
	})

	Describe("Update", func() {
		var item structs.Shelf

		BeforeEach(func() {
			item = structs.Shelf{
				ID:    uint64(2),
				Title: "Teste",
			}
		})

		It("Should returns error when title length < 1", func() {
			item.Title = ""
			_, err := repository.Update(item)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
		})

		It("Should returns error when ID not found", func() {
			item.ID = uint64(12345)
			_, err := repository.Update(item)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
		})

		It("Should returns its own ID after it was updated", func() {
			result, err := repository.Update(item)
			Expect(err).ShouldNot(HaveOccurred())
			saved := result.(structs.Shelf)
			Expect(saved.ID).To(Equal(item.ID))
		})
	})

	Describe("Delete", func() {
		It("Should returns error when ID not found", func() {
			item := structs.Shelf{ID: 1234567}
			_, err := repository.Delete(item)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.NotFound))
		})

		It("Should not returns error after deleted", func() {
			item := structs.Shelf{ID: 2}
			id, err := repository.Delete(item)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(id).To(Equal(item.ID))
		})
	})

})
