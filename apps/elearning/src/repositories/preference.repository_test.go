package repositories_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/elearning/src/repositories"
	"backend/apps/elearning/src/structs"
	"backend/libs/configuration"
	"backend/libs/databases"
	testutil "backend/libs/testing"
)

func TestPreferenceRepository(t *testing.T) {
	RegisterFailHandler(Fail)
	if !testing.Short() {
		RunSpecs(t, "Repository > Preference")
	}
}

var _ = Describe("PreferenceRepository", func() {
	var repository repositories.Repository
	var db databases.Database

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "elearning.sql")
		repository = repositories.NewPreferenceRepository(db)
	})

	Describe("ReadOne", func() {
		It("Should returns preference by type", func() {
			result, err := repository.ReadOne(structs.Preference{Type: "home"})
			Expect(err).ShouldNot(HaveOccurred())

			found := result.(structs.Preference)

			Expect(found.Type).To(Equal("home"))
			Expect(found.Content).To(Equal(`{"shelves": [2, 1, 3]}`))
		})
	})

	Describe("Update", func() {

		var fake structs.Preference

		BeforeEach(func() {
			fake = structs.Preference{
				Type:    "home",
				Content: `{"shelves": [2, 1, 3]}`,
			}
		})

		It("Should returns errors when home content is not a valid json", func() {
			fake.Content = "InvalidJson"

			result, err := repository.Update(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).To(BeZero())
		})

		It("Should overrides home content", func() {
			result, err := repository.Update(fake)
			Expect(err).ShouldNot(HaveOccurred())

			preference := result.(structs.Preference)
			Expect(preference.Type).ShouldNot(BeEmpty())
			Expect(preference.Content).ShouldNot(BeEmpty())
		})
	})
})
