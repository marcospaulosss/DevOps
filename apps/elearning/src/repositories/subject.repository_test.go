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

func TestSubjectRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	if !testing.Short() {
		RunSpecs(t, "Repository > Subject")
	}
}

var _ = Describe("SubjectRepository", func() {
	var repository repositories.SubjectRepository
	var db databases.Database
	var config = configuration.Get()
	var databaseURL = config.GetEnvConfString("database_url")
	var subject structs.Subject

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "elearning.sql")
		repository = repositories.NewSubjectRepository(db).(repositories.SubjectRepository)

		subject = structs.Subject{
			Title: "TesteUnitario",
		}
	})

	Describe("Create", func() {
		It("Should return its own ID after it was created", func() {
			result, err := repository.Create(subject)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.(structs.Subject).ID).To(Equal(uint64(4)))
		})

		It("Should return error in statement", func() {
			db.GetConnection().Exec("DROP TABLE subjects_tracks")
			db.GetConnection().Exec("DROP TABLE subjects")
			result, err := repository.Create(subject)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})

		It("Should returns error when title already exists ignoring case", func() {
			subject.Title = "Português"
			result, err := repository.Create(subject)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Unknown))
			Expect(result).Should(BeZero())
		})
	})

	Describe("Update", func() {
		BeforeEach(func() {
			subject.ID = 1
		})

		It("Should return its own ID after it was created", func() {
			result, err := repository.Update(subject)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.(structs.Subject).ID).To(Equal(subject.ID))
		})

		It("Should returns error when id not exist", func() {
			subject.ID = uint64(999999)
			result, err := repository.Update(subject)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.NotFound))
			Expect(result).Should(BeZero())
		})

		It("Should returns error when title already exists ignoring case", func() {
			subject.Title = "Matématica"
			result, err := repository.Update(subject)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Unknown))
			Expect(result).Should(BeZero())
		})
	})

	Describe("ReadOne", func() {
		BeforeEach(func() {
			subject.ID = 1
			subject.Title = "Português"
		})

		It("Should return subject found to id ", func() {
			result, err := repository.ReadOne(subject)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.(structs.Subject).ID).To(Equal(subject.ID))
			Expect(result.(structs.Subject).Title).To(Equal(subject.Title))
			Expect(result.(structs.Subject).CreatedAt).ShouldNot(BeZero())
		})

		It("Should returns error when id not exist", func() {
			subject.ID = uint64(999999)
			result, err := repository.ReadOne(subject)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.NotFound))
			Expect(result).Should(BeZero())
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
			Expect(result.Pagination.PerPage).To(Equal(search.Pagination.PerPage))
			Expect(result.Pagination.Page).To(Equal(search.Pagination.Page))
			Expect(result.Pagination.Total).To(Equal(int32(3)))
			items := result.Items.([]structs.Subject)
			Expect(items).Should(HaveLen(3))
			Expect(items[0].ID).To(Equal(uint64(1)))
			Expect(items[1].ID).To(Equal(uint64(2)))
		})

		It("Should returns a subject using filter", func() {
			search := structs.Search{
				Pagination: structs.Pagination{
					PerPage: 10,
					Page:    1,
				},
				Raw: "(title[eq]:'Matématica')",
			}
			result, err := repository.ReadAll(search)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Pagination.PerPage).To(Equal(search.Pagination.PerPage))
			Expect(result.Pagination.Page).To(Equal(search.Pagination.Page))
			Expect(result.Pagination.Total).To(Equal(int32(1)))
			items := result.Items.([]structs.Subject)
			Expect(items).Should(HaveLen(1))
			Expect(items[0].ID).To(Equal(uint64(2)))
		})

		It("Should returns error when have problem in database", func() {
			db.GetConnection().Exec("DROP TABLE subjects_tracks")
			db.GetConnection().Exec("DROP TABLE subjects")
			search := structs.Search{}
			result, err := repository.ReadAll(search)
			Expect(err).Should(HaveOccurred())
			Expect(result.Items).Should(BeZero())
		})

		It("Should returns empty items when not found subjects in database", func() {
			db.GetConnection().Exec("DELETE FROM subjects")
			search := structs.Search{}
			result, err := repository.ReadAll(search)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Pagination.Total).To(Equal(int32(0)))
			items := result.Items.([]structs.Subject)
			Expect(items).Should(HaveLen(0))
		})
	})

	Describe("Delete", func() {
		It("Should returns error when ID not found", func() {
			item := structs.Subject{ID: 1234567}
			_, err := repository.Delete(item)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.NotFound))
		})

		It("Should not returns error after deleted", func() {
			item := structs.Subject{ID: 2}
			id, err := repository.Delete(item)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(id).To(Equal(item.ID))
		})
	})
})
