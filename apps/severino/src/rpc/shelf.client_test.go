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

func TestShelfClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Client > Elearning > Shelf")
}

var _ = Describe("ShelfClient", func() {
	var ctx = context.Background()
	fake := structs.Shelf{Title: "shelf 1"}
	err := errors.NewGrpcError(errors.Internal, "Failed")

	Context("Create", func() {
		It("Should returns created item", func() {
			resp := &pb.ShelfResponse{
				Shelf: &pb.Shelf{Id: 1},
			}
			service := testutil.MockShelfServiceClient("Create", ctx, resp, nil)
			client := rpc.NewShelfClient(service, ctx)
			result, err := client.Create(fake)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeZero())

			req := &pb.ShelfRequest{Shelf: adapters.ToProtoShelf(fake)}
			mock := service.(*mocks.ShelfServiceClient)
			mock.AssertCalled(GinkgoT(), "Create", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockShelfServiceClient("Create", ctx, nil, err)
			client := rpc.NewShelfClient(service, ctx)
			result, err := client.Create(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("Update", func() {
		It("Should returns updated item", func() {
			resp := &pb.ShelfResponse{
				Shelf: &pb.Shelf{Id: 1},
			}
			service := testutil.MockShelfServiceClient("Update", ctx, resp, nil)
			client := rpc.NewShelfClient(service, ctx)
			fake.ID = 1
			result, err := client.Update(fake)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeZero())

			req := &pb.ShelfRequest{Shelf: adapters.ToProtoShelf(fake)}
			mock := service.(*mocks.ShelfServiceClient)
			mock.AssertCalled(GinkgoT(), "Update", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockShelfServiceClient("Update", ctx, nil, err)
			client := rpc.NewShelfClient(service, ctx)
			result, err := client.Update(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("Delete", func() {
		It("Should returns deleted item", func() {
			resp := &pb.ShelfResponse{
				Shelf: &pb.Shelf{Id: 1},
			}
			service := testutil.MockShelfServiceClient("Delete", ctx, resp, nil)
			client := rpc.NewShelfClient(service, ctx)
			fake.ID = 1
			result, err := client.Delete(fake)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeNil())

			req := &pb.ShelfRequest{Shelf: adapters.ToProtoShelf(fake)}
			mock := service.(*mocks.ShelfServiceClient)
			mock.AssertCalled(GinkgoT(), "Delete", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockShelfServiceClient("Delete", ctx, nil, err)
			client := rpc.NewShelfClient(service, ctx)
			result, err := client.Delete(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("ReadOne", func() {
		It("Should returns the found item", func() {
			resp := &pb.ShelfResponse{
				Shelf: &pb.Shelf{
					Id:    1,
					Title: "shelf Test",
					Albums: []*pb.Album{
						{Id: 1, Title: "album 1"},
						{Id: 2, Title: "album 2"},
					},
				},
			}

			service := testutil.MockShelfServiceClient("ReadOne", ctx, resp, nil)
			client := rpc.NewShelfClient(service, ctx)
			result, err := client.ReadOne(structs.Shelf{ID: 1})
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeNil())

			req := &pb.ShelfRequest{Shelf: adapters.ToProtoShelf(structs.Shelf{ID: 1})}
			mock := service.(*mocks.ShelfServiceClient)
			mock.AssertCalled(GinkgoT(), "ReadOne", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockShelfServiceClient("ReadOne", ctx, nil, err)
			client := rpc.NewShelfClient(service, ctx)
			result, err := client.ReadOne(structs.Shelf{ID: 1})
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("ReadAll", func() {
		It("Should returns an item list", func() {
			shelves := []*pb.Shelf{
				{
					Title: "shelf Test",
					Albums: []*pb.Album{
						{Id: 1, Title: "album 1"},
					},
				},
			}

			resp := &pb.ShelvesResponse{Shelves: shelves}
			service := testutil.MockShelfServiceClientSearch("ReadAll", ctx, resp, nil)
			client := rpc.NewShelfClient(service, ctx)

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
			mock := service.(*mocks.ShelfServiceClient)
			mock.AssertCalled(GinkgoT(), "ReadAll", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockShelfServiceClientSearch("ReadAll", ctx, nil, err)
			client := rpc.NewShelfClient(service, ctx)
			result, err := client.ReadAll(structs.Search{})
			Expect(err).ShouldNot(BeZero())
			Expect(result).Should(BeZero())
		})
	})

})
