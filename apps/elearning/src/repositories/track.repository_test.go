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

func TestTrackRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	if !testing.Short() {
		RunSpecs(t, "Repository > Track")
	}
}

var _ = Describe("TrackRepository", func() {
	var repository repositories.TrackRepository
	var db databases.Database

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")
	var item structs.Track

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "elearning.sql")
		repository = repositories.NewTrackRepository(db)

		item = structs.Track{
			Title:       "Intro",
			Description: "Apenas mais uma track",
			Media:       "a7b3c1f2-9526-4cd6-8953-49f3f1b22469.mp3",
			Duration:    120,
			Teachers:    "Maria,Helena",
		}
	})

	Describe("ReadOne", func() {
		id := uint64(15)
		It("Should returns item by ID", func() {
			result, err := repository.ReadOne(structs.Track{ID: id})
			Expect(err).ShouldNot(HaveOccurred())

			track := result.(structs.Track)
			Expect(track.ID).To(Equal(id))
			Expect(track.Title).To(Equal("Conectivos"))
			Expect(track.Teachers).To(Equal("Ena Loiola"))
			Expect(track.Media).To(Equal("c5c94c85-bd07-4aa9-9c73-9dde53f42251.mp3"))
			Expect(track.Duration).To(Equal(int64(42)))
			Expect(track.Albums).Should(HaveLen(0))
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
			Expect(result.Pagination.Total).To(Equal(int32(12)))
			items := result.Items.([]structs.Track)
			Expect(items).To(HaveLen(10))
			item := items[0]
			Expect(item.ID).To(Equal(uint64(8)))
			Expect(item.Title).To(Equal("Forças Armadas (FFAA)"))
			Expect(item.Description).To(Equal("Missão Constitucional; Hierarquia e disciplina; e Comandante Supremo das Forças Armadas."))
			Expect(item.Teachers).To(Equal("Alan Hirt"))
			Expect(item.Albums).Should(HaveLen(0))
		})
	})

	Describe("Create", func() {

		It("Should returns error when title length < 1", func() {
			item.Title = ""
			_, err := repository.Create(item)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
		})

		It("Should returns its own ID after it was created", func() {
			result, err := repository.Create(item)
			Expect(err).ShouldNot(HaveOccurred())
			saved := result.(structs.Track)
			Expect(saved.ID).To(Equal(uint64(1000)))
		})
	})

	Describe("Update", func() {
		var item structs.Track

		BeforeEach(func() {
			item.ID = uint64(8)
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
			item.Title = "Intro"
			result, err := repository.Update(item)
			Expect(err).ShouldNot(HaveOccurred())
			saved := result.(structs.Track)
			Expect(saved.ID).To(Equal(item.ID))
		})
	})

	Describe("Delete", func() {
		It("Should returns error when ID not found", func() {
			item := structs.Track{ID: 1234567}
			_, err := repository.Delete(item)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.NotFound))
		})

		It("Should not returns error after deleted", func() {
			item := structs.Track{ID: 8}
			id, err := repository.Delete(item)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(id).To(Equal(item.ID))
		})
	})

	Describe("CreateSubjectsAndAssociateItWithTracks", func() {
		var item structs.Track

		BeforeEach(func() {
			item.ID = 12
			item.Subject = structs.Subject{ID: uint64(1)}
		})

		It("Should returns its own ID after it was associate", func() {
			err := repository.CreateSubjectsAndAssociateItWithTracks(item)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Should returns error when track already associated the subject", func() {
			item.ID = 8
			err := repository.CreateSubjectsAndAssociateItWithTracks(item)
			Expect(err).Should(HaveOccurred())
		})

		It("Should returns error when track already associated the subject", func() {
			item.ID = 0
			err := repository.CreateSubjectsAndAssociateItWithTracks(item)
			Expect(err).Should(HaveOccurred())
		})

		It("Should returns error when track already associated the subject", func() {
			item.Subject.ID = 0
			err := repository.CreateSubjectsAndAssociateItWithTracks(item)
			Expect(err).Should(HaveOccurred())
		})
	})
})
