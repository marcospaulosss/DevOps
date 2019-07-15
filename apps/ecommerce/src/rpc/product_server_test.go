package rpc_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"backend/apps/ecommerce/src/repositories"
	"backend/apps/ecommerce/src/rpc"
	"backend/libs/configuration"
	testutil "backend/libs/testing"
	pb "backend/proto"
)

const (
	CREATE   = "Create"
	READ_ONE = "ReadOne"
	READ_ALL = "ReadAll"
)

func TestProductHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Server > Product")
}

var _ = Describe("ProductServer", func() {
	var server *rpc.ProductServer
	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db := testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "ecommerce.sql")
		repos := repositories.Container{
			ProductRepository: repositories.NewProductRepository(db),
		}
		server = rpc.NewProductServer(nil, repos)
	})

	Context(CREATE, func() {
		It("Should returns a product with only ID when created", func() {
			creditcard := &pb.PaymentType{
				Id:           "creditcard",
				Name:         "Cartão de Crédito",
				Installments: int32(12),
				Price:        int32(12399),
			}
			req := &pb.ProductRequest{Product: &pb.Product{
				Name:          "Pacote com 5 aulas",
				Description:   "Imperdível",
				Stock:         2,
				ProductType:   "pacote",
				PaymentsTypes: []*pb.PaymentType{creditcard},
				Items:         `[123456]`,
			}}

			resp, err := server.Create(context.Background(), req)
			Expect(err).ShouldNot(HaveOccurred())

			saved := resp.GetProduct()
			Expect(saved.Id).To(HaveLen(36))
		})

		It("Should returns error when name is empty", func() {
			req := &pb.ProductRequest{Product: &pb.Product{
				Description: "Imperdível",
				Stock:       2,
			}}

			resp, err := server.Create(context.Background(), req)
			Expect(err).Should(HaveOccurred())
			Expect(resp).Should(BeNil())
		})

		It("Should returns error when field items is empty", func() {
			req := &pb.ProductRequest{Product: &pb.Product{
				Description: "Imperdível",
				Stock:       2,
				Items:       "",
			}}

			resp, err := server.Create(context.Background(), req)
			Expect(err).Should(HaveOccurred())

			st, _ := status.FromError(err)
			Expect(st.Code()).To(Equal(codes.Internal))
			Expect(resp).Should(BeNil())
		})
	})

	Context(READ_ONE, func() {
		It("Should returns a product by its ID", func() {
			req := &pb.ProductRequest{Product: &pb.Product{
				Id: "a0d354bc-b8c3-4038-a18d-12307e80759b",
			}}

			resp, err := server.ReadOne(context.Background(), req)
			Expect(err).ShouldNot(HaveOccurred())

			item := resp.GetProduct()
			Expect(item.GetId()).To(HaveLen(36))
			Expect(item.GetName()).To(Equal("Elearning Audio"))
			Expect(item.GetDescription()).To(Equal(""))
			Expect(item.GetStock()).To(Equal(int32(1)))
			Expect(item.GetHistory()).ShouldNot(BeEmpty())
		})
	})

	Context(READ_ALL, func() {
		It("Should returns all items when pagination was not defined", func() {
			req := &pb.SearchRequest{
				Search: &pb.Search{
					Pagination: &pb.Pagination{
						Order:  "id",
						SortBy: "asc",
					},
				},
			}
			resp, err := server.ReadAll(context.Background(), req)
			Expect(err).ShouldNot(HaveOccurred())

			// pagination := resp.GetPagination()
			// Expect(pagination.Total).To(Equal(int32(2)))

			products := resp.GetProducts()
			Expect(products).To(HaveLen(2))

			item := products[1]
			Expect(item.GetId()).To(HaveLen(36))
			Expect(item.GetName()).To(Equal("Elearning Audio"))
			Expect(item.GetStock()).To(Equal(int32(1)))
			Expect(item.GetPaymentsTypes()).ShouldNot(BeZero())

			Expect(item.GetItems()).ShouldNot(BeZero())
		})
	})
})
