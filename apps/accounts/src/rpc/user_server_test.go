package rpc_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/accounts/src/repositories"
	"backend/apps/accounts/src/rpc"
	"backend/libs/configuration"
	"backend/libs/databases"
	testutil "backend/libs/testing"
	pb "backend/proto"
)

func TestAccountsHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC > User")
}

var _ = Describe("UserServer", func() {

	ctx := context.Background()
	var db databases.Database
	var server *rpc.UserServer

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "accounts.sql")
		repos := repositories.Container{
			AccountRepository: repositories.NewAccountRepository(db),
			UserRepository:    repositories.NewUserRepository(db),
		}
		server = rpc.NewUserServer(nil, repos)
	})

	Describe("Create", func() {
		var fake pb.User

		BeforeEach(func() {
			fake.Phone = "5511912363678"
			fake.Name = "José dos Santos"
			fake.Email = "josedossantos@estrategiaconcursos.com.br"
			fake.Active = true
		})

		Describe("Create", func() {

			It("Should rise error when phone does not exits on account", func() {
				req := &pb.UserRequest{User: &fake}
				res, err := server.Create(ctx, req)
				Expect(err).Should(HaveOccurred())
				Expect(res.GetUser().GetId()).Should(BeZero())
				Expect(res.GetUser().GetEmail()).Should(BeZero())
				Expect(res.GetUser().GetPhone()).Should(BeZero())
			})

			It("Should rise error when email does not exits on account", func() {
				req := &pb.UserRequest{User: &fake}
				res, err := server.Create(ctx, req)
				Expect(err).Should(HaveOccurred())
				Expect(res.GetUser().GetId()).Should(BeZero())
				Expect(res.GetUser().GetEmail()).Should(BeZero())
				Expect(res.GetUser().GetPhone()).Should(BeZero())
			})

			It("Should rise error when email is vinculated to registered user", func() {
				fake.Email = "userssecond@estrategiaconcursos.com.br"

				req := &pb.UserRequest{User: &fake}
				res, err := server.Create(ctx, req)
				Expect(err).Should(HaveOccurred())
				Expect(res.GetUser().GetId()).Should(BeZero())
				Expect(res.GetUser().GetEmail()).Should(BeZero())
				Expect(res.GetUser().GetPhone()).Should(BeZero())
			})

			It("Should rise error when phone is vinculated to registered user", func() {
				fake.Phone = "5511999999666"

				req := &pb.UserRequest{User: &fake}
				res, err := server.Create(ctx, req)
				Expect(err).Should(HaveOccurred())
				Expect(res.GetUser().GetId()).Should(BeZero())
				Expect(res.GetUser().GetEmail()).Should(BeZero())
				Expect(res.GetUser().GetPhone()).Should(BeZero())
			})

			It("Should created user", func() {
				fake.Email = "code3@estrategiaconcursos.com.br"
				fake.EmailCode = "987633"
				fake.Phone = "5511999999966"
				fake.PhoneCode = "123466"
				req := &pb.UserRequest{User: &fake}
				res, err := server.Create(ctx, req)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res.GetUser().GetId()).ShouldNot(BeZero())
			})
		})

		Describe("ReadAll", func() {
			var search pb.SearchRequest
			BeforeEach(func() {
				search.Search = &pb.Search{}
			})

			It("Should empty list if does not found any registers", func() {
				search.Search.Raw = "(id[eq]:'aaaaaaaa-13c8-4301-8de3-f625d622a01e')"
				res, err := server.ReadAll(ctx, &search)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res.GetUsers()).To(HaveLen(0))
			})

			It("Should return a list of users", func() {
				res, err := server.ReadAll(ctx, &search)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(len(res.GetUsers())).To(BeNumerically(">", 1))
			})

			It("Should return a list of users filtered by email", func() {
				search.Search.Raw = "(email[eq]:'usersfive@estrategiaconcursos.com.br')"
				res, err := server.ReadAll(ctx, &search)
				Expect(err).ShouldNot(HaveOccurred())
				Expect(res.GetUsers()).To(HaveLen(1))
				Expect(res.GetUsers()[0].GetEmail()).To(Equal("usersfive@estrategiaconcursos.com.br"))
			})
		})
	})
})
