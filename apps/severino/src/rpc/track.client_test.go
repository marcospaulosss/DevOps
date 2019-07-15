package rpc_test

import (
	"context"
	"testing"

	"backend/apps/severino/src/structs"
	"backend/libs/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/mocks"
	"backend/apps/severino/src/adapters"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/testutil"
	pb "backend/proto"
)

func TestTrackClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Client > Elearning > Track")
}

var _ = Describe("TrackClient", func() {
	var ctx = context.Background()
	fake := structs.Track{Title: "track 1"}
	err := errors.NewGrpcError(errors.Internal, "Failed")

	Context("Create", func() {
		It("Should returns created item", func() {
			resp := &pb.TrackResponse{
				Track: &pb.Track{Id: 1},
			}
			service := testutil.MockTrackServiceClient("Create", ctx, resp, nil)
			client := rpc.NewTrackClient(service, ctx)
			result, err := client.Create(fake)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeZero())

			req := &pb.TrackRequest{Track: adapters.ToProtoTrack(fake)}
			mock := service.(*mocks.TrackServiceClient)
			mock.AssertCalled(GinkgoT(), "Create", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockTrackServiceClient("Create", ctx, nil, err)
			client := rpc.NewTrackClient(service, ctx)
			result, err := client.Create(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("Update", func() {
		It("Should returns updated item", func() {
			resp := &pb.TrackResponse{
				Track: &pb.Track{Id: 1},
			}
			service := testutil.MockTrackServiceClient("Update", ctx, resp, nil)
			client := rpc.NewTrackClient(service, ctx)
			fake.ID = 1
			result, err := client.Update(fake)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeZero())

			req := &pb.TrackRequest{Track: adapters.ToProtoTrack(fake)}
			mock := service.(*mocks.TrackServiceClient)
			mock.AssertCalled(GinkgoT(), "Update", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockTrackServiceClient("Update", ctx, nil, err)
			client := rpc.NewTrackClient(service, ctx)
			result, err := client.Update(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("Delete", func() {
		It("Should returns deleted item", func() {
			resp := &pb.TrackResponse{
				Track: &pb.Track{Id: 1},
			}
			service := testutil.MockTrackServiceClient("Delete", ctx, resp, nil)
			client := rpc.NewTrackClient(service, ctx)
			fake.ID = 1
			result, err := client.Delete(fake)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeNil())

			req := &pb.TrackRequest{Track: adapters.ToProtoTrack(fake)}
			mock := service.(*mocks.TrackServiceClient)
			mock.AssertCalled(GinkgoT(), "Delete", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockTrackServiceClient("Delete", ctx, nil, err)
			client := rpc.NewTrackClient(service, ctx)
			result, err := client.Delete(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("ReadOne", func() {
		It("Should returns the found item", func() {
			resp := &pb.TrackResponse{
				Track: &pb.Track{
					Id:    1,
					Title: "track Test",
					Albums: []*pb.Album{
						{Id: 1, Title: "album 1"},
						{Id: 2, Title: "album 2"},
					},
				},
			}

			service := testutil.MockTrackServiceClient("ReadOne", ctx, resp, nil)
			client := rpc.NewTrackClient(service, ctx)
			result, err := client.ReadOne(structs.Track{ID: 1})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeNil())

			req := &pb.TrackRequest{Track: adapters.ToProtoTrack(structs.Track{ID: 1})}
			mock := service.(*mocks.TrackServiceClient)
			mock.AssertCalled(GinkgoT(), "ReadOne", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockTrackServiceClient("ReadOne", ctx, nil, err)
			client := rpc.NewTrackClient(service, ctx)
			result, err := client.ReadOne(structs.Track{ID: 1})
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("ReadAll", func() {
		It("Should returns an item list", func() {
			tracks := []*pb.Track{
				{
					Title: "track Test",
					Albums: []*pb.Album{
						{Id: 1, Title: "album 1"},
					},
				},
			}

			resp := &pb.TracksResponse{Tracks: tracks}
			service := testutil.MockTrackServiceClientSearch("ReadAll", ctx, resp, nil)
			client := rpc.NewTrackClient(service, ctx)

			search := structs.Search{
				Pagination: structs.Pagination{
					PerPage: 9,
					Page:    6,
					Order:   "id",
					SortBy:  "asc",
				},
				Raw: "",
			}
			result, err := client.ReadAll(search)
			Expect(err).Should(BeZero())
			Expect(result).ShouldNot(BeZero())

			req := &pb.SearchRequest{
				Search: &pb.Search{
					Pagination: &pb.Pagination{
						PerPage: 9,
						Page:    6,
						Order:   "id",
						SortBy:  "asc",
					},
					Raw: "",
				},
			}
			mock := service.(*mocks.TrackServiceClient)
			mock.AssertCalled(GinkgoT(), "ReadAll", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockTrackServiceClientSearch("ReadAll", ctx, nil, err)
			client := rpc.NewTrackClient(service, ctx)
			result, err := client.ReadAll(structs.Search{})
			Expect(err).ShouldNot(BeZero())
			Expect(result).Should(BeZero())
		})
	})

})
