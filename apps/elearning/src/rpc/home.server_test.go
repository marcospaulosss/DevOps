package rpc_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/elearning/src/repositories"
	"backend/apps/elearning/src/rpc"
	"backend/libs/configuration"
	"backend/libs/databases"
	testutil "backend/libs/testing"
	pb "backend/proto"
)

func TestHomeServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Server > Home")
}

var _ = Describe("HomeServer", func() {
	ctx := context.Background()
	var db databases.Database
	var server *rpc.HomeServer

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "elearning.sql")
		repos := repositories.Container{
			PreferenceRepository: repositories.NewPreferenceRepository(db),
			ShelfRepository:      repositories.NewShelfRepository(db),
			AlbumRepository:      repositories.NewAlbumRepository(db),
		}
		server = rpc.NewHomeServer(nil, repos)
	})

	Describe("ReadAll", func() {
		var search *pb.Search

		BeforeEach(func() {
			search = &pb.Search{
				Pagination: &pb.Pagination{
					Page:    1,
					PerPage: 10,
					Order:   "id",
					SortBy:  "asc",
				},
			}
		})

		It("Should return shelves with at least 1 related album", func() {
			req := &pb.SearchRequest{Search: search}
			resp, err := server.ReadAll(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(resp.Shelves).ShouldNot(BeEmpty())
		})
	})
})
