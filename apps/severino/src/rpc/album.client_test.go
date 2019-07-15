package rpc_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/mocks"
	"backend/apps/severino/src/adapters"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/structs"
	"backend/apps/severino/testutil"
	"backend/libs/errors"
	pb "backend/proto"
)

func TestAlbumClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Client > Elearning > Albums")
}

var _ = Describe("AlbumClient", func() {
	var ctx = context.Background()
	fake := structs.Album{Title: "Album 1"}
	err := errors.NewGrpcError(errors.Internal, "Failed")

	Context("Create", func() {
		It("Should returns created item", func() {
			resp := &pb.AlbumResponse{
				Album: &pb.Album{Id: 1},
			}
			service := testutil.MockAlbumServiceClient("Create", ctx, resp, nil)
			client := rpc.NewAlbumClient(service, ctx)
			result, err := client.Create(fake)
			Expect(err).Should(BeZero())
			Expect(result).ShouldNot(BeZero())

			req := &pb.AlbumRequest{Album: adapters.ToProtoAlbum(fake)}
			mock := service.(*mocks.AlbumServiceClient)
			mock.AssertCalled(GinkgoT(), "Create", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockAlbumServiceClient("Create", ctx, nil, err)
			client := rpc.NewAlbumClient(service, ctx)
			result, err := client.Create(fake)
			Expect(err).ShouldNot(BeZero())
			Expect(result).Should(BeZero())
		})
	})

	Context("Update", func() {
		BeforeEach(func() {
			fake = structs.Album{ID: 1, Title: "Album 1"}
		})

		It("Should returns updated item", func() {
			resp := &pb.AlbumResponse{
				Album: &pb.Album{Id: 1},
			}
			service := testutil.MockAlbumServiceClient("Update", ctx, resp, nil)
			client := rpc.NewAlbumClient(service, ctx)
			result, err := client.Update(fake)
			Expect(err).Should(BeZero())
			Expect(result).ShouldNot(BeZero())

			req := &pb.AlbumRequest{Album: adapters.ToProtoAlbum(fake)}
			mock := service.(*mocks.AlbumServiceClient)
			mock.AssertCalled(GinkgoT(), "Update", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockAlbumServiceClient("Update", ctx, nil, err)
			client := rpc.NewAlbumClient(service, ctx)
			result, err := client.Update(fake)
			Expect(err).ShouldNot(BeZero())
			Expect(result).Should(BeZero())
		})
	})

	Context("Delete", func() {
		It("Should returns nil when item was deleted", func() {
			service := testutil.MockAlbumServiceClient("Delete", ctx, nil, nil)
			client := rpc.NewAlbumClient(service, ctx)
			result, err := client.Delete(fake)
			Expect(err).Should(BeZero())
			Expect(result).Should(BeZero())

			req := &pb.AlbumRequest{Album: adapters.ToProtoAlbum(fake)}
			mock := service.(*mocks.AlbumServiceClient)
			mock.AssertCalled(GinkgoT(), "Delete", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockAlbumServiceClient("Delete", ctx, nil, err)
			client := rpc.NewAlbumClient(service, ctx)
			result, err := client.Delete(fake)
			Expect(err).ShouldNot(BeZero())
			Expect(result).Should(BeZero())
		})
	})

	Context("ReadOne", func() {
		It("Should returns the found item", func() {
			resp := &pb.AlbumResponse{
				Album: &pb.Album{Id: 1},
			}
			service := testutil.MockAlbumServiceClient("ReadOne", ctx, resp, nil)
			client := rpc.NewAlbumClient(service, ctx)
			result, err := client.ReadOne(fake)
			Expect(err).Should(BeZero())
			Expect(result).ShouldNot(BeZero())

			req := &pb.AlbumRequest{Album: adapters.ToProtoAlbum(fake)}
			mock := service.(*mocks.AlbumServiceClient)
			mock.AssertCalled(GinkgoT(), "ReadOne", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockAlbumServiceClient("ReadOne", ctx, nil, err)
			client := rpc.NewAlbumClient(service, ctx)
			result, err := client.ReadOne(fake)
			Expect(err).ShouldNot(BeZero())
			Expect(result).Should(BeZero())
		})
	})

	Context("ReadAll", func() {
		It("Should returns an item list", func() {
			shelves1 := make([]*pb.Shelf, 1)
			shelves1[0] = &pb.Shelf{Id: 5}
			item1 := &pb.Album{
				Id:      1,
				Shelves: shelves1,
			}
			shelves2 := make([]*pb.Shelf, 1)
			shelves2[0] = &pb.Shelf{Id: 6}
			item2 := &pb.Album{
				Id:      2,
				Shelves: shelves2,
			}
			albums := make([]*pb.Album, 2)
			albums[0] = item1
			albums[1] = item2
			resp := &pb.AlbumsResponse{
				Albums: albums,
			}
			service := testutil.MockAlbumServiceClientSearch("ReadAll", ctx, resp, nil)
			client := rpc.NewAlbumClient(service, ctx)

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
			mock := service.(*mocks.AlbumServiceClient)
			mock.AssertCalled(GinkgoT(), "ReadAll", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockAlbumServiceClientSearch("ReadAll", ctx, nil, err)
			client := rpc.NewAlbumClient(service, ctx)
			result, err := client.ReadAll(structs.Search{})
			Expect(err).ShouldNot(BeZero())
			Expect(result).Should(BeZero())
		})
	})
})
