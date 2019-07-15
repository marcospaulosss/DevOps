package repositories_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/accounts/src/repositories"
	"backend/apps/accounts/src/structs"
	"backend/libs/configuration"
	"backend/libs/databases"
	testutil "backend/libs/testing"
)

func TestAccountRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	if !testing.Short() {
		RunSpecs(t, "Repository > Account")
	}
}

var _ = Describe("AccountRepository", func() {
	var db databases.Database
	var repository repositories.AccountRepository

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "accounts.sql")
		repository = repositories.NewAccountRepository(db)
	})

	Describe("Create", func() {
		var item structs.Account

		BeforeEach(func() {
			item = structs.Account{
				Type:      "email",
				Email:     "testeUnitario@estrategiaconcursos.com.br",
				EmailCode: "020783",
				Phone:     "5511775539485",
				PhoneCode: "04072017",
			}
		})

		It("Should returns its own ID when params is valid", func() {
			result, err := repository.Create(&item)
			Expect(err).ShouldNot(HaveOccurred())
			saved := result.(*structs.Account)
			Expect(saved.ID).ShouldNot(BeZero())
			Expect(saved.ID).ShouldNot(BeNil())
			Expect(saved.ID).ShouldNot(BeEmpty())
		})

		It("Should returns its own ID when items is valid and user login", func() {
			item.Exists = true
			result, err := repository.Create(&item)
			Expect(err).ShouldNot(HaveOccurred())
			saved := result.(*structs.Account)
			Expect(saved.ID).ShouldNot(BeZero())
			Expect(saved.ID).ShouldNot(BeNil())
			Expect(saved.ID).ShouldNot(BeEmpty())
		})

		It("Should returns error when params account is invalid", func() {
			itemEmpty := structs.Account{}
			result, err := repository.Create(&itemEmpty)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Describe("ReadOne", func() {
		var item structs.Account

		BeforeEach(func() {
			item = structs.Account{
				Type:      "email",
				Email:     "code@estrategiaconcursos.com.br",
				EmailCode: "987654",
				Phone:     "5511999999999",
				PhoneCode: "123456",
			}
		})

		It("Should returns account when email and code is found", func() {
			result, err := repository.ReadOne(&item)
			Expect(err).ShouldNot(HaveOccurred())
			saved := result.(*structs.Account)
			Expect(saved.ID).ShouldNot(BeZero())
			Expect(saved.Email).To(Equal(item.Email))
			Expect(saved.EmailCode).To(Equal(item.EmailCode))
			Expect(saved.Phone).To(Equal(item.Phone))
			Expect(saved.PhoneCode).To(Equal(item.PhoneCode))
			Expect(saved.CreatedAt).ShouldNot(BeZero())
		})

		It("Should returns error when params account is invalid", func() {
			itemEmpty := structs.Account{}
			result, err := repository.ReadOne(&itemEmpty)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})

		It("Should returns error when account not found", func() {
			item.Email = "teste"
			result, err := repository.ReadOne(&item)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Describe("Delete", func() {
		var item structs.Account

		BeforeEach(func() {
			item = structs.Account{
				Type:      "email",
				Email:     "code@estrategiaconcursos.com.br",
				EmailCode: "987654",
				Phone:     "5511999999999",
				PhoneCode: "123456",
			}
		})

		It("Should returns id account when account is deleted", func() {
			result, err := repository.Delete(&item)
			Expect(err).ShouldNot(HaveOccurred())
			saved := result.(*structs.Account)
			Expect(saved.ID).ShouldNot(BeZero())
		})

		It("Should returns error when params account is invalid", func() {
			itemEmpty := structs.Account{}
			result, err := repository.Delete(&itemEmpty)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Describe("FetchCodeExisting", func() {
		var item structs.Account

		BeforeEach(func() {
			item = structs.Account{
				Type:      "email",
				Email:     "code@estrategiaconcursos.com.br",
				EmailCode: "987654",
				Phone:     "5511999999999",
				PhoneCode: "123456",
			}
		})

		It("Should returns account when is found", func() {
			result, err := repository.FetchCodeExisting(&item)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.ID).ShouldNot(BeZero())
			Expect(result.Email).To(Equal(item.Email))
			Expect(result.EmailCode).To(Equal(item.EmailCode))
			Expect(result.Phone).To(Equal(item.Phone))
			Expect(result.PhoneCode).To(Equal(item.PhoneCode))
			Expect(result.CreatedAt).ShouldNot(BeZero())
		})

		It("Should returns error when params account is invalid", func() {
			item.Email = "teste"
			result, err := repository.FetchCodeExisting(&item)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})

		It("Should returns error when params account is invalid", func() {
			var itemEmpty structs.Account
			result, err := repository.FetchCodeExisting(&itemEmpty)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})
})
