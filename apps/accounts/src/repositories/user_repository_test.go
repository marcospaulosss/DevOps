package repositories_test

import (
	"testing"

	"backend/libs/configuration"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/accounts/src/repositories"
	"backend/apps/accounts/src/structs"
	"backend/libs/databases"
	testutil "backend/libs/testing"
)

func TestUserRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	if !testing.Short() {
		RunSpecs(t, "Repository > User")
	}
}

var _ = Describe("UserRepository", func() {
	var db databases.Database
	var userRepository repositories.Repository

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "accounts.sql")
		userRepository = repositories.NewUserRepository(db)
	})

	Describe("#Create", func() {
		It("Should return an user when created successfully", func() {
			user := structs.User{
				Email: "leonardmccoy@unitedfederationofplanets.com",
				Phone: "5511900000000",
				Name:  "Leonard Mccoy",
			}

			savedInterface, err := userRepository.Create(user)
			saved := savedInterface.(*structs.User)

			Expect(saved.ID).NotTo(BeEmpty())
			Expect(err).ToNot(HaveOccurred())

		})
		It("Should return error when user with same email already exists", func() {
			user := structs.User{
				Email: "usersfirst@estrategiaconcursos.com.br",
				Phone: "5511900000000",
				Name:  "Leonard Mccoy",
			}

			saved, err := userRepository.Create(user)

			Expect(saved).To(BeNil())
			Expect(err).To(HaveOccurred())
		})
		It("Should return error when user with same phone already exists", func() {
			user := structs.User{
				Email: "leonardmccoy@unitedfederationofplanets.com",
				Phone: "5511999999999",
				Name:  "Leonard Mccoy",
			}

			saved, err := userRepository.Create(user)

			Expect(saved).To(BeNil())
			Expect(err).To(HaveOccurred())
		})
	})
})
