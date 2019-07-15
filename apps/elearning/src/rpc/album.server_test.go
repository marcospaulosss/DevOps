package rpc_test

import (
	"context"
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/elearning/src/repositories"
	"backend/apps/elearning/src/rpc"
	"backend/libs/configuration"
	"backend/libs/databases"
	"backend/libs/errors"
	testutil "backend/libs/testing"
	pb "backend/proto"
)

func TestAlbumServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Server > Album")
}

var _ = Describe("AlbumServer", func() {
	ctx := context.Background()
	var db databases.Database
	var server *rpc.AlbumServer

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "elearning.sql")
		repos := repositories.Container{
			AlbumRepository: repositories.NewAlbumRepository(db),
			ShelfRepository: repositories.NewShelfRepository(db),
		}
		server = rpc.NewAlbumServer(nil, repos)
	})

	Describe("Create", func() {
		var item *pb.Album

		BeforeEach(func() {
			trackA := &pb.Track{Id: uint64(8), Subject: &pb.Subject{}}
			section1 := &pb.Section{
				Title:       "Capitulo Final",
				Description: "Ultimo",
				Tracks:      []*pb.Track{trackA},
			}
			trackB := &pb.Track{Id: uint64(9), Subject: &pb.Subject{}}
			trackC := &pb.Track{Id: uint64(10), Subject: &pb.Subject{}}
			section2 := &pb.Section{
				Title:       "Capitulo 1",
				Description: "Primeiro",
				Tracks:      []*pb.Track{trackC, trackB},
			}
			item = &pb.Album{
				Title:       "Album 1",
				Description: "Awesome album",
				Sections:    []*pb.Section{section1, section2},
			}
		})

		It("Should create the album and its sections associated with tracks", func() {
			req := &pb.AlbumRequest{Album: item}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			id := uint64(1000)
			Expect(res.GetAlbum().GetId()).To(Equal(id))

			var totalAlbums, totalSections, totalTracks int
			query := fmt.Sprintf(`select count(a.id) as totalAlbums,
                (select count(s.id) from sections s where s.album_id = %d) as totalSections,
                (select count(st.section_id) from sections_tracks st, sections s where st.section_id = s.id and s.album_id=%d) as totalTracks
            from albums a where a.id=%d`, id, id, id)
			testutil.QueryRow(query).Scan(&totalAlbums, &totalSections, &totalTracks)
			Expect(totalAlbums).To(Equal(1))
			Expect(totalSections).To(Equal(2))
			Expect(totalTracks).To(Equal(3))
		})

		It("Should returns error when title is empty", func() {
			item.Title = ""
			req := &pb.AlbumRequest{Album: item}
			res, err := server.Create(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
			Expect(res.GetAlbum()).Should(BeNil())
		})
	})

	Describe("Update", func() {
		var item *pb.Album

		BeforeEach(func() {
			trackA := &pb.Track{Id: uint64(8), Subject: &pb.Subject{}}
			section1 := &pb.Section{
				Title:       "Capitulo Final",
				Description: "Ultimo",
				Tracks:      []*pb.Track{trackA},
			}
			trackB := &pb.Track{Id: uint64(9), Subject: &pb.Subject{}}
			trackC := &pb.Track{Id: uint64(10), Subject: &pb.Subject{}}
			section2 := &pb.Section{
				Title:       "Capitulo 1",
				Description: "Primeiro",
				Tracks:      []*pb.Track{trackC, trackB},
			}
			item = &pb.Album{
				Id:          uint64(4),
				Title:       "Album 1",
				Description: "Awesome album",
				Sections:    []*pb.Section{section1, section2},
			}
		})

		It("Should updates the album and recreates its sections", func() {
			req := &pb.AlbumRequest{Album: item}
			res, err := server.Update(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			id := uint64(4)
			Expect(res.GetAlbum().GetId()).To(Equal(id))

			var totalAlbums, totalSections, totalTracks int
			query := fmt.Sprintf(`select count(a.id) as totalAlbums,
                (select count(s.id) from sections s where s.album_id = %d) as totalSections,
                (select count(st.section_id) from sections_tracks st, sections s where st.section_id = s.id and s.album_id=%d) as totalTracks
            from albums a where a.id=%d`, id, id, id)
			testutil.QueryRow(query).Scan(&totalAlbums, &totalSections, &totalTracks)
			Expect(totalAlbums).To(Equal(1))
			Expect(totalSections).To(Equal(2))
			Expect(totalTracks).To(Equal(3))
		})

		It("Should returns error when title is empty", func() {
			item.Title = ""
			req := &pb.AlbumRequest{Album: item}
			res, err := server.Update(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
			Expect(res.GetAlbum()).Should(BeNil())
		})
	})

	Describe("Delete", func() {
		var item *pb.Album

		BeforeEach(func() {
			item = &pb.Album{
				Id: uint64(4),
			}
		})

		It("Should delete albums and its sections", func() {
			req := &pb.AlbumRequest{Album: item}
			res, err := server.Delete(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			id := uint64(4)
			Expect(res.GetAlbum().GetId()).To(Equal(id))

			var totalAlbums, totalSections, totalTracks int
			query := fmt.Sprintf(`select count(a.id) as totalAlbums,
                (select count(s.id) from sections s where s.album_id = %d) as totalSections,
                (select count(st.section_id) from sections_tracks st, sections s where st.section_id = s.id and s.album_id=%d) as totalTracks
            from albums a where a.id=%d`, id, id, id)
			testutil.QueryRow(query).Scan(&totalAlbums, &totalSections, &totalTracks)
			Expect(totalAlbums).To(Equal(0))
			Expect(totalSections).To(Equal(0))
			Expect(totalTracks).To(Equal(0))
		})

		It("Should returns error when ID was not found", func() {
			item.Id = uint64(12345)
			req := &pb.AlbumRequest{Album: item}
			res, err := server.Delete(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.NotFound))
			Expect(res.GetAlbum()).Should(BeNil())
		})
	})

	Describe("ReadOne", func() {
		var item *pb.Album

		BeforeEach(func() {
			item = &pb.Album{
				Id: uint64(7),
			}
		})

		It("Should returns album with sections and tracks", func() {
			req := &pb.AlbumRequest{Album: item}
			res, err := server.ReadOne(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())

			album := res.GetAlbum()
			Expect(album.Id).To(Equal(item.Id))
			Expect(album.Sections).Should(HaveLen(3))
			Expect(album.Sections[0].Id).To(Equal(uint64(107)))
			Expect(album.Teachers).To(Equal("Ena Loiola"))
			Expect(album.Shelves).Should(HaveLen(1))
			Expect(album.Shelves[0].Id).To(Equal(uint64(3)))
		})
	})

	Describe("ReadAll", func() {
		var search *pb.Search

		BeforeEach(func() {
			search = &pb.Search{
				Pagination: &pb.Pagination{
					Page:    1,
					PerPage: 10,
					Order:   "id",
					SortBy:  "desc",
				},
			}
		})

		It("Should returns all albums ordered by id DESC", func() {
			req := &pb.SearchRequest{Search: search}
			res, err := server.ReadAll(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())

			albums := res.GetAlbums()
			Expect(albums).Should(HaveLen(4))
			album := albums[0]
			Expect(album.Id).To(Equal(uint64(7)))
			Expect(album.Title).To(Equal("Regência verbal e nominal"))
			Expect(album.Teachers).To(Equal("Ena Loiola"))
			Expect(album.Sections).Should(HaveLen(3))
			Expect(album.Sections[0].Id).Should(Equal(uint64(107)))
			Expect(album.Sections[1].Tracks).Should(HaveLen(7))
			Expect(album.Sections[1].Tracks[0].Id).Should(Equal(uint64(14)))
			Expect(album.Shelves).Should(HaveLen(1))
			Expect(album.Shelves[0].Id).Should(Equal(uint64(3)))

		})

	})

	Describe("Publish", func() {
		var item *pb.Album

		BeforeEach(func() {
			item = &pb.Album{
				Id:          uint64(6),
				IsPublished: true,
			}
		})

		It("Should returns the same album after published", func() {
			req := &pb.AlbumRequest{Album: item}
			res, err := server.Publish(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())

			album := res.GetAlbum()
			Expect(album.Id).To(Equal(item.Id))

			var isPublished bool
			testutil.QueryRow("select is_published from albums where id=6").Scan(&isPublished)
			Expect(isPublished).To(BeTrue())
		})
	})
})
