package repositories_test

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/ecommerce/src/repositories"
	"backend/apps/ecommerce/src/structs"
	"backend/libs/configuration"
	"backend/libs/databases"
	testutil "backend/libs/testing"
)

const (
	CREATE   = "Create"
	UPDATE   = "Update"
	DELETE   = "Delete"
	READ_ONE = "ReadOne"
	READ_ALL = "ReadAll"
)

func TestProductRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Repository > Product")
}

var _ = Describe("ProductRepository", func() {
	var productRepository repositories.Repository
	var db databases.Database

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "ecommerce.sql")
		productRepository = repositories.NewProductRepository(db)
	})

	Describe(READ_ALL, func() {
		It("Should returns all items", func() {
			search := structs.Search{
				Pagination: structs.Pagination{
					PerPage: 10,
					Page:    1,
				},
			}
			result, err := productRepository.ReadAll(search)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Pagination.Total).To(Equal(int32(2)))

			products := result.Items.([]structs.Product)
			Expect(len(products)).To(Equal(2))

			product := products[1]
			Expect(product.ID).To(Equal("a0d354bc-b8c3-4038-a18d-12307e80759b"))
			Expect(product.Stock).To(Equal(int32(1)))
			Expect(product.ProductType).To(Equal("assinatura"))

			history, err := product.GetHistory()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(history).To(HaveLen(1))

			paymentsTypes, err := product.GetPaymentsTypes()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(paymentsTypes).To(HaveLen(1))
			Expect(paymentsTypes[0].ID).To(Equal("creditcard"))
			Expect(paymentsTypes[0].Name).To(Equal("Cartão de Crédito"))
			Expect(paymentsTypes[0].Installments).To(Equal(int32(12)))
			Expect(paymentsTypes[0].Price).To(Equal(int32(39999)))
		})

		It("Should returns all items skipping pagination when PerPage is zero", func() {
			search := structs.Search{
				Pagination: structs.Pagination{
					Page: 100,
				},
			}
			result, err := productRepository.ReadAll(search)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Pagination.Total).To(Equal(int32(2)))

			products := result.Items.([]structs.Product)
			Expect(len(products)).To(Equal(2))
		})

		It("Should returns paginated items", func() {
			search := structs.Search{
				Pagination: structs.Pagination{
					PerPage: 1,
					Page:    1,
				},
			}
			result, err := productRepository.ReadAll(search)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.Pagination.Total).To(Equal(int32(2)))

			products := result.Items.([]structs.Product)
			Expect(len(products)).To(Equal(1))
		})
	})

	Describe(READ_ONE, func() {
		It("Should returns the item by ID", func() {
			product := structs.Product{
				ID: "a0d354bc-b8c3-4038-a18d-12307e80759b",
			}

			result, err := productRepository.ReadOne(product)
			Expect(err).ShouldNot(HaveOccurred())
			item := result.(structs.Product)
			Expect(item.ID).To(Equal(product.ID))
			Expect(item.Name).To(Equal("Elearning Audio"))
			Expect(item.Description).To(Equal(""))
			Expect(item.Stock).To(Equal(int32(1)))

			history, err := item.GetHistory()
			Expect(err).ShouldNot(HaveOccurred())
			Expect(history[0].Name).To(Equal("TRT audios"))
			Expect(history[0].Stock).To(Equal(int32(5)))
			Expect(history[0].IsPublished).Should(BeFalse())
		})

		It("Should returns error when ID not found", func() {
			product := structs.Product{
				ID: "2b9f9a8a-fa5d-4627-a7fa-cbd8a50e3ec6",
			}
			result, err := productRepository.ReadOne(product)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeNil())
		})

	})

	Describe(CREATE, func() {
		It("Should create a valid product and returns itself", func() {
			creditcard := structs.PaymentType{
				ID:           "creditcard",
				Name:         "Cartão de Crédito",
				Installments: 12,
				Price:        12399,
			}
			paymentsTypes := []structs.PaymentType{creditcard}
			types, _ := json.Marshal(paymentsTypes)
			str := string(types)
			items := `[123, 456]`

			product := structs.Product{
				Name:          "Curso preparatório para TRT",
				Description:   "Melhor curso do mercado com 100% de aproveitamento",
				Stock:         2,
				ProductType:   "pacote",
				PaymentsTypes: &str,
				Items:         &items,
			}

			result, err := productRepository.Create(product)
			Expect(err).ShouldNot(HaveOccurred())
			saved := result.(structs.Product)
			Expect(saved.ID).To(HaveLen(36))
		})

		It("Should fail when name is empty", func() {
			product := structs.Product{}

			result, err := productRepository.Create(product)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeNil())
		})

		It("Should fail when name length is less than 3 chars", func() {
			product := structs.Product{Name: "xx"}

			result, err := productRepository.Create(product)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeNil())
		})
	})

	Describe(UPDATE, func() {
		It("Should record history and returns itself", func() {
			product := structs.Product{
				ID:          "575e7695-241d-4b96-8061-9f8a3771f3cc",
				Name:        "Pacotão TRT",
				Description: "O melhor curso do mercado",
				Stock:       int32(5),
			}
			result, err := productRepository.Update(product)
			Expect(err).ShouldNot(HaveOccurred())

			saved := result.(structs.Product)
			Expect(saved.ID).To(Equal(product.ID))
		})
	})
})
