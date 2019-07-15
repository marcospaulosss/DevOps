package rpc_test

import (
	"context"
	"testing"

	"backend/apps/elearning/src/structs"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/elearning/src/repositories"
	"backend/apps/elearning/src/rpc"
	"backend/libs/configuration"
	"backend/libs/databases"
	testutil "backend/libs/testing"
	pb "backend/proto"
)

func TestPreferenceServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Server > Preference")
}

var _ = Describe("PreferenceServer", func() {
	ctx := context.Background()
	var db databases.Database
	var server *rpc.PreferenceServer

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "elearning.sql")
		repos := repositories.Container{
			PreferenceRepository: repositories.NewPreferenceRepository(db),
		}
		server = rpc.NewPreferenceServer(nil, repos)
	})

	Describe("Update", func() {
		var item *pb.Preference

		BeforeEach(func() {
			item = &pb.Preference{
				Type:    "home",
				Content: `{"shelves": [2, 1, 3]}`,
			}
		})

		It("Should update home configuration", func() {
			req := &pb.PreferenceRequest{Preference: item}
			res, err := server.Update(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetPreference().GetType()).To(Equal(item.Type))
			Expect(res.GetPreference().GetContent()).To(Equal(item.Content))

			var preference structs.Preference
			testutil.QueryRow("select type, content from preferences where type = 'home' limit 1").StructScan(&preference)
			Expect(preference.Type).To(Equal(item.Type))
			Expect(preference.Content).To(Equal(item.Content))
		})
	})

})
