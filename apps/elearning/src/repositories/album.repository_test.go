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

func TestAlbumRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	if !testing.Short() {
		RunSpecs(t, "Repository > Album")
	}
}

var _ = Describe("AlbumRepository", func() {
	var repository repositories.AlbumRepository
	var db databases.Database

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "elearning.sql")
		repository = repositories.NewAlbumRepository(db).(repositories.AlbumRepository)
	})

	Describe("ReadAll", func() {
		It("Should returns all items", func() {
			search := structs.Search{
				Pagination: structs.Pagination{
					PerPage: 10,
					Page:    1,
				},
			}
			result, err := repository.ReadAll(search)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Pagination.Total).To(Equal(int32(4)))

			items := result.Items.([]structs.Album)
			Expect(items).To(HaveLen(4))

			item := items[3]
			Expect(item.ID).To(Equal(uint64(7)))
			Expect(item.Title).To(Equal("Regência verbal e nominal"))
		})

		It("Should returns albums filtering by ID", func() {
			search := structs.Search{
				Pagination: structs.Pagination{
					PerPage: 10,
					Page:    1,
				},
				Raw: "(id[eq]:'4')",
			}
			result, err := repository.ReadAll(search)
			Expect(err).ShouldNot(HaveOccurred())

			items := result.Items.([]structs.Album)
			Expect(items[0].ID).To(Equal(uint64(4)))
		})

		It("Should returns albums filtering by title", func() {
			search := structs.Search{
				Pagination: structs.Pagination{
					PerPage: 10,
					Page:    1,
				},
				Raw: "(title[icontains]:'%forças%')",
			}
			result, err := repository.ReadAll(search)
			Expect(err).ShouldNot(HaveOccurred())

			items := result.Items.([]structs.Album)
			Expect(items).Should(HaveLen(1))
			Expect(items[0].Title).To(Equal("Forças armadas"))
		})
	})

	Describe("ReadOne", func() {
		id := uint64(7)
		It("Should returns item by ID", func() {
			result, err := repository.ReadOne(structs.Album{ID: id})
			Expect(err).ShouldNot(HaveOccurred())

			album := result.(structs.Album)
			Expect(album.ID).To(Equal(id))
			Expect(album.Title).To(Equal("Regência verbal e nominal"))
		})

		It("Should returns error when ID was not found", func() {
			id := uint64(1234556)
			_, err := repository.ReadOne(structs.Album{ID: id})
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.NotFound))
		})

		It("Should returns album with empty sections", func() {
			testutil.Exec("delete FROM sections")
			result, err := repository.ReadOne(structs.Album{ID: id})
			Expect(err).ShouldNot(HaveOccurred())
			album := result.(structs.Album)
			Expect(album.Sections).Should(BeZero())
			Expect(album.Teachers).Should(BeZero())
		})

		It("Should returns album with empty teachers", func() {
			testutil.Exec("Update tracks SET teachers = ''")
			result, err := repository.ReadOne(structs.Album{ID: id})
			Expect(err).ShouldNot(HaveOccurred())
			album := result.(structs.Album)
			Expect(album.Teachers).Should(BeZero())
		})
	})

	Describe("Create", func() {
		var item structs.Album

		BeforeEach(func() {
			trackA1 := structs.Track{ID: uint64(10), Title: "Infrações disciplinares e Órgãos da OAB"}
			trackA2 := structs.Track{ID: uint64(11), Title: "Princípios; Inscrição na OAB; Direitos do Advogado"}
			sectionA1 := structs.Section{Title: "Captitulo 1", Tracks: []structs.Track{trackA1, trackA2}}
			item = structs.Album{
				Title:       "Album 1",
				Description: "Primeiro album",
				Image:       "a7b3c1f2-9526-4cd6-8953-49f3f1b22469.jpg",
				Sections:    []structs.Section{sectionA1},
			}
		})

		It("Should returns error when title length < 1", func() {
			item.Title = ""
			_, err := repository.Create(item)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
		})

		It("Should not returns error when it does not have a shelf", func() {
			item.Shelves = []structs.Shelf{}
			_, err := repository.Create(item)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Should not returns error when shelf ID is zero or empty", func() {
			item.Shelves = []structs.Shelf{structs.Shelf{Title: "Shelf 1"}}
			_, err := repository.Create(item)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Should returns error when shelf ID does not exists", func() {
			item.Shelves = []structs.Shelf{structs.Shelf{ID: uint64(12345)}}
			_, err := repository.Create(item)
			Expect(err).ShouldNot(HaveOccurred())
		})

		It("Should returns its own ID after it was created and associated with shelves and sections", func() {
			result, err := repository.Create(item)
			Expect(err).ShouldNot(HaveOccurred())
			saved := result.(structs.Album)
			Expect(saved.ID).To(Equal(uint64(1000)))
		})
	})

	Describe("Update", func() {
		var item structs.Album

		BeforeEach(func() {
			trackA1 := structs.Track{ID: uint64(10), Title: "Infrações disciplinares e Órgãos da OAB"}
			trackA2 := structs.Track{ID: uint64(11), Title: "Princípios; Inscrição na OAB; Direitos do Advogado"}
			sectionA1 := structs.Section{Title: "Captitulo 1", Tracks: []structs.Track{trackA1, trackA2}}
			item = structs.Album{
				ID:          uint64(7),
				Title:       "Album 1",
				Description: "Primeiro album",
				Image:       "a7b3c1f2-9526-4cd6-8953-49f3f1b22469.jpg",
				Sections:    []structs.Section{sectionA1},
			}
		})

		It("Should returns error when title length < 1", func() {
			item.Title = ""
			_, err := repository.Update(item)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
		})

		It("Should returns its own ID after updated", func() {
			result, err := repository.Update(item)
			Expect(err).ShouldNot(HaveOccurred())
			saved := result.(structs.Album)
			Expect(saved.ID).To(Equal(item.ID))
		})
	})

	Describe("Delete", func() {
		It("Should returns error when ID not found", func() {
			item := structs.Album{ID: 1234567}
			_, err := repository.Delete(item)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.NotFound))
		})

		It("Should not returns error after deleted", func() {
			item := structs.Album{ID: 7}
			id, err := repository.Delete(item)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(id).To(Equal(item.ID))
		})
	})

	Describe("FindAlbumsByTrackID", func() {
		It("Should returns albums according of track ID", func() {
			albums, err := repository.FindAlbumsByTrackID(uint64(15))
			Expect(err).ShouldNot(HaveOccurred())
			Expect(albums).Should(HaveLen(2))
		})
	})

	Context("Publication", func() {
		var item structs.Album

		BeforeEach(func() {
			item = structs.Album{ID: uint64(6)}
		})

		Describe("Publish", func() {
			It("Should returns album after it is published", func() {
				album, err := repository.Publish(item)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(album.ID).To(Equal(item.ID))

				var isPublished bool
				testutil.QueryRow("select is_published from albums where id=6").Scan(&isPublished)
				Expect(isPublished).To(BeTrue())
			})
		})

		Describe("Unpublish", func() {
			It("Should returns album after it is unpublished", func() {
				item.ID = uint64(4)
				album, err := repository.Unpublish(item)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(album.ID).To(Equal(item.ID))

				var isPublished bool
				testutil.QueryRow("select is_published from albums where id=6").Scan(&isPublished)
				Expect(isPublished).To(BeFalse())
			})
		})
	})
})
