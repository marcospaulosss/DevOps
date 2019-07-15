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

func TestProductClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Client > Elearning > Products")
}

var _ = Describe("ProductClient", func() {
	var ctx = context.Background()
	fake := structs.Product{Name: "Fake Product"}
	err := errors.NewGrpcError(errors.Internal, "Failed")

	Context("Create", func() {
		It("Should returns created item", func() {
			resp := &pb.ProductResponse{
				Product: &pb.Product{Id: "1"},
			}
			service := testutil.MockProductServiceClient("Create", ctx, resp, nil)
			client := rpc.NewProductClient(service, ctx)
			result, err := client.Create(fake)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeZero())

			req := &pb.ProductRequest{Product: adapters.ToProtoProduct(fake)}
			mock := service.(*mocks.ProductServiceClient)
			mock.AssertCalled(GinkgoT(), "Create", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockProductServiceClient("Create", ctx, nil, err)
			client := rpc.NewProductClient(service, ctx)
			result, err := client.Create(fake)

			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("Update", func() {
		It("Should returns updated item", func() {
			resp := &pb.ProductResponse{
				Product: &pb.Product{Id: "1"},
			}
			service := testutil.MockProductServiceClient("Update", ctx, resp, nil)
			client := rpc.NewProductClient(service, ctx)

			fake.Name = "New Name"
			result, err := client.Update(fake)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeZero())

			req := &pb.ProductRequest{Product: adapters.ToProtoProduct(fake)}
			mock := service.(*mocks.ProductServiceClient)
			mock.AssertCalled(GinkgoT(), "Update", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockProductServiceClient("Update", ctx, nil, err)
			client := rpc.NewProductClient(service, ctx)
			result, err := client.Update(fake)

			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("Delete", func() {
		It("Should returns nil", func() {
			service := testutil.MockProductServiceClient("Delete", ctx, nil, nil)
			client := rpc.NewProductClient(service, ctx)
			result, err := client.Delete(fake)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).Should(BeNil())
		})
	})

	Context("ReadOne", func() {
		It("Should returns the found item", func() {
			resp := &pb.ProductResponse{
				Product: &pb.Product{
					Id:   "1",
					Name: "Test Product",
					PaymentsTypes: []*pb.PaymentType{
						{Id: "1", Name: "test", Price: 666, Installments: 1},
						{Id: "2", Name: "test 2", Price: 333, Installments: 2},
					},
				},
			}

			service := testutil.MockProductServiceClient("ReadOne", ctx, resp, nil)
			client := rpc.NewProductClient(service, ctx)
			result, err := client.ReadOne(fake)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeZero())

			req := &pb.ProductRequest{Product: adapters.ToProtoProduct(fake)}
			mock := service.(*mocks.ProductServiceClient)
			mock.AssertCalled(GinkgoT(), "ReadOne", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockProductServiceClient("ReadOne", ctx, nil, err)
			client := rpc.NewProductClient(service, ctx)
			result, err := client.ReadOne(fake)

			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("ReadAll", func() {
		It("Should returns an item list", func() {
			products := []*pb.Product{
				{
					Id:   "1",
					Name: "Test Product",
					PaymentsTypes: []*pb.PaymentType{
						{Id: "1", Name: "test", Price: 666, Installments: 1},
						{Id: "2", Name: "test 2", Price: 333, Installments: 2},
					},
				},
				{
					Id:   "2",
					Name: "Test Product",
					PaymentsTypes: []*pb.PaymentType{
						{Id: "2", Name: "test 2", Price: 333, Installments: 2},
					},
				},
			}

			resp := &pb.ProductsResponse{Products: products}

			service := testutil.MockProductServiceClientSearch("ReadAll", ctx, resp, nil)
			client := rpc.NewProductClient(service, ctx)

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
			mock := service.(*mocks.ProductServiceClient)
			mock.AssertCalled(GinkgoT(), "ReadAll", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockProductServiceClientSearch("ReadAll", ctx, nil, err)
			client := rpc.NewProductClient(service, ctx)
			result, err := client.ReadAll(structs.Search{})
			Expect(err).ShouldNot(BeZero())
			Expect(result).Should(BeZero())
		})
	})
})
