package rpc_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/elearning/src/repositories"
	"backend/apps/elearning/src/rpc"
	"backend/libs/configuration"
	"backend/libs/databases"
	testutil "backend/libs/testing"
	pb "backend/proto"
)

func TestShelfServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Server > Shelf")
}

var _ = Describe("ShelfServer", func() {
	ctx := context.Background()
	var db databases.Database
	var server *rpc.ShelfServer

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "elearning.sql")
		repos := repositories.Container{
			ShelfRepository: repositories.NewShelfRepository(db),
			AlbumRepository: repositories.NewAlbumRepository(db),
		}
		server = rpc.NewShelfServer(nil, repos)
	})

	Describe("Create", func() {
		var item *pb.Shelf

		BeforeEach(func() {
			album := &pb.Album{Id: 5}
			item = &pb.Shelf{
				Title:  "TRT",
				Albums: []*pb.Album{album},
			}
		})

		It("Should create shelf and does not associate it with albums", func() {
			item.Albums = nil
			req := &pb.ShelfRequest{Shelf: item}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetShelf().GetId()).To(Equal(uint64(1000)))

			var total int
			testutil.QueryRow("select count(album_id) from albums_shelves where shelf_id=1000 limit 1").Scan(&total)
			Expect(total).To(Equal(0))
		})

		It("Should create shelf and associate it with albums", func() {
			req := &pb.ShelfRequest{Shelf: item}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetShelf().GetId()).To(Equal(uint64(1000)))

			var id uint64
			testutil.QueryRow("select album_id from albums_shelves where shelf_id=1000 limit 1").Scan(&id)
			Expect(id).To(Equal(item.Albums[0].Id))
		})

		It("Should create shelf and does not associate it when album does not exists", func() {
			album := &pb.Album{Id: 12345678}
			item.Albums = []*pb.Album{album}
			req := &pb.ShelfRequest{Shelf: item}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetShelf().GetId()).To(Equal(uint64(1000)))

			var total int
			testutil.QueryRow("select count(album_id) from albums_shelves where shelf_id=1000 limit 1").Scan(&total)
			Expect(total).To(Equal(0))
		})
	})

	Describe("Update", func() {
		var item *pb.Shelf

		BeforeEach(func() {
			album := &pb.Album{Id: 5}
			item = &pb.Shelf{
				Id:     1,
				Title:  "Shelf 1",
				Albums: []*pb.Album{album},
			}
		})

		It("Should update shelf and reassociate it with albums", func() {
			req := &pb.ShelfRequest{Shelf: item}
			res, err := server.Update(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetShelf().GetId()).To(Equal(item.Id))

			var total int
			testutil.QueryRow("select count(album_id) from albums_shelves where shelf_id=1 limit 1").Scan(&total)
			Expect(total).To(Equal(1))
		})
	})

	Describe("Delete", func() {
		var item *pb.Shelf

		BeforeEach(func() {
			item = &pb.Shelf{
				Id: 1,
			}
		})

		It("Should delete shelf and its associations with albums", func() {
			req := &pb.ShelfRequest{Shelf: item}
			res, err := server.Delete(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetShelf().GetId()).To(Equal(item.Id))

			var total int
			testutil.QueryRow("select count(album_id) from albums_shelves where shelf_id=1 limit 1").Scan(&total)
			Expect(total).To(Equal(0))
		})
	})

	Describe("ReadOne", func() {
		item := &pb.Shelf{
			Id: 2,
		}

		It("Should returns a shelf with its albums ordered by album position ASC", func() {
			req := &pb.ShelfRequest{Shelf: item}
			res, err := server.ReadOne(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())

			found := res.GetShelf()
			Expect(found.GetId()).To(Equal(item.Id))

			albums := found.GetAlbums()
			Expect(albums).Should(HaveLen(2))
			Expect(albums[0].GetId()).To(Equal(uint64(6)))
			Expect(albums[1].GetId()).To(Equal(uint64(5)))
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
					SortBy:  "asc",
				},
			}
		})

		It("Should returns all shelves with its albums ordered by album position ASC", func() {
			req := &pb.SearchRequest{Search: search}
			res, err := server.ReadAll(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())

			shelves := res.GetShelves()
			Expect(shelves).Should(HaveLen(3))

			item := shelves[1]
			Expect(item.Id).To(Equal(uint64(2)))
			albums := item.GetAlbums()
			Expect(albums[0].Id).To(Equal(uint64(6)))
			Expect(albums[1].Id).To(Equal(uint64(5)))
		})
	})
})
