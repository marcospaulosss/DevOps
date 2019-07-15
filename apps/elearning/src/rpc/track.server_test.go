package rpc_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/elearning/src/repositories"
	"backend/apps/elearning/src/rpc"
	"backend/apps/elearning/src/structs"
	"backend/libs/configuration"
	"backend/libs/databases"
	"backend/libs/errors"
	testutil "backend/libs/testing"
	pb "backend/proto"
)

func TestTrackServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Server > Track")
}

var _ = Describe("TrackServer", func() {
	ctx := context.Background()
	var db databases.Database
	var server *rpc.TrackServer

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "elearning.sql")
		repos := repositories.Container{
			TrackRepository: repositories.NewTrackRepository(db),
			AlbumRepository: repositories.NewAlbumRepository(db),
		}
		server = rpc.NewTrackServer(nil, repos)
	})

	Describe("Create", func() {
		var item *pb.Track

		BeforeEach(func() {
			item = &pb.Track{
				Title:       "Faixa 1",
				Description: "My first track",
				Teachers:    "Maria,Leila",
				Duration:    120,
				Media:       "3f6ce0c7-993d-495a-aa8c-c9abad82e90e.mp3",
				Subject:     &pb.Subject{Id: uint64(1)},
			}
		})

		It("Should create a track and returns its ID", func() {
			req := &pb.TrackRequest{Track: item}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetTrack().GetId()).To(Equal(uint64(1000)))

			var track structs.Track
			testutil.QueryRow("select * from tracks where id=1000").StructScan(&track)
			Expect(track.Title).To(Equal(item.Title))
			Expect(track.Description).To(Equal(item.Description))
			Expect(track.Teachers).To(Equal(item.Teachers))
			Expect(track.Duration).To(Equal(item.Duration))
			Expect(track.Media).To(Equal(item.Media))
		})

		It("Should returns error when title is empty", func() {
			item.Title = ""
			req := &pb.TrackRequest{Track: item}
			_, err := server.Create(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
		})

		It("Should returns error when title is empty", func() {
			item.Subject.Id = 0
			req := &pb.TrackRequest{Track: item}
			_, err := server.Create(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
		})
	})

	Describe("Update", func() {
		var item *pb.Track

		BeforeEach(func() {
			item = &pb.Track{
				Id:          uint64(8),
				Title:       "Faixa 1",
				Description: "My first track",
				Teachers:    "Maria,Leila",
				Duration:    120,
				Media:       "3f6ce0c7-993d-495a-aa8c-c9abad82e90e.mp3",
				Subject:     &pb.Subject{},
			}
		})

		It("Should update the track and returns its ID", func() {
			req := &pb.TrackRequest{Track: item}
			res, err := server.Update(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetTrack().GetId()).To(Equal(item.Id))

			var track structs.Track
			testutil.QueryRow("select * from tracks where id=8").StructScan(&track)
			Expect(track.Title).To(Equal(item.Title))
			Expect(track.Description).To(Equal(item.Description))
			Expect(track.Teachers).To(Equal(item.Teachers))
			Expect(track.Duration).To(Equal(item.Duration))
			Expect(track.Media).To(Equal(item.Media))
		})

		It("Should returns error when title is empty", func() {
			item.Title = ""
			req := &pb.TrackRequest{Track: item}
			_, err := server.Create(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.Invalid))
		})
	})

	Describe("Delete", func() {
		var item *pb.Track

		BeforeEach(func() {
			item = &pb.Track{
				Id:      uint64(8),
				Subject: &pb.Subject{},
			}
		})

		It("Should returns error when ID was not found", func() {
			item.Id = uint64(1234566)
			req := &pb.TrackRequest{Track: item}
			_, err := server.Delete(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.NotFound))
		})

		It("Should returns the ID when item was deleted", func() {
			req := &pb.TrackRequest{Track: item}
			result, err := server.Delete(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.GetTrack().GetId()).To(Equal(item.Id))

			var total int
			testutil.QueryRow("select count(track_id) from sections_tracks where track_id=8").Scan(&total)
			Expect(total).To(Equal(0))
		})
	})

	Describe("ReadOne", func() {
		var item *pb.Track

		BeforeEach(func() {
			item = &pb.Track{
				Id:      uint64(8),
				Subject: &pb.Subject{},
			}
		})

		It("Should returns error when ID was not found", func() {
			item.Id = uint64(1234566)
			req := &pb.TrackRequest{Track: item}
			_, err := server.ReadOne(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(errors.GetCodeFrom(err)).To(Equal(errors.NotFound))
		})

		It("Should returns the item found by ID", func() {
			item.Id = uint64(15)
			req := &pb.TrackRequest{Track: item}
			result, err := server.ReadOne(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			track := result.GetTrack()
			Expect(track.GetId()).To(Equal(item.Id))
			Expect(track.Albums).Should(HaveLen(2))
		})
	})

	Describe("ReadAll", func() {
		var search *pb.Search

		BeforeEach(func() {
			pagination := &pb.Pagination{
				PerPage: 10,
				Page:    1,
			}
			search = &pb.Search{
				Pagination: pagination,
			}
		})

		It("Should returns the paginated items", func() {
			req := &pb.SearchRequest{Search: search}
			result, err := server.ReadAll(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Tracks).Should(HaveLen(int(search.Pagination.PerPage)))
			Expect(result.Tracks[0].Albums).Should(HaveLen(1))
			Expect(result.Pagination.Total).To(Equal(int32(12)))
		})

		It("Should returns filtered items", func() {
			search.Raw = `(teachers[icontains]:'%daniela%')`
			req := &pb.SearchRequest{Search: search}
			result, err := server.ReadAll(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Tracks).Should(HaveLen(3))
			Expect(result.Tracks[0].Albums).Should(HaveLen(1))
			Expect(result.Pagination.Total).To(Equal(int32(3)))
		})
	})
})
