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

func TestSubjectServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Server > Subject")
}

var _ = Describe("SubjectServer", func() {
	ctx := context.Background()
	var db databases.Database
	var server *rpc.SubjectServer

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	var item *pb.Subject

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "elearning.sql")
		repos := repositories.Container{
			SubjectRepository: repositories.NewSubjectRepository(db),
		}
		server = rpc.NewSubjectServer(nil, repos)

		item = &pb.Subject{
			Id:    uint64(1),
			Title: "TRT",
		}
	})

	Describe("Create", func() {
		It("Should create subject", func() {
			req := &pb.SubjectRequest{Subject: item}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetSubject().GetId()).To(Equal(uint64(4)))
		})

		It("Should error when title subject is exist in database", func() {
			item.Title = "Português"
			req := &pb.SubjectRequest{Subject: item}
			res, err := server.Create(ctx, req)
			Expect(res.GetSubject()).Should(BeZero())
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("Update", func() {
		It("Should update subject", func() {
			req := &pb.SubjectRequest{Subject: item}
			res, err := server.Update(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetSubject().GetId()).To(Equal(item.Id))
		})

		It("Should error when title already exists in database", func() {
			item.Title = "Matématica"
			req := &pb.SubjectRequest{Subject: item}
			res, err := server.Update(ctx, req)
			Expect(res.GetSubject()).Should(BeZero())
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("ReadOne", func() {
		It("Should find subject to id", func() {
			req := &pb.SubjectRequest{Subject: item}
			res, err := server.ReadOne(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetSubject().GetId()).To(Equal(item.Id))
			Expect(res.GetSubject().GetTitle()).To(Equal("Português"))
			Expect(res.GetSubject().GetCreatedAt()).ShouldNot(BeZero())
		})

		It("Should error when id not found", func() {
			item.Id = 999999
			req := &pb.SubjectRequest{Subject: item}
			res, err := server.ReadOne(ctx, req)
			Expect(res.GetSubject()).Should(BeZero())
			Expect(err).Should(HaveOccurred())
		})
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

		It("Should returns all subjects ordered by album position ASC", func() {
			req := &pb.SearchRequest{Search: search}
			res, err := server.ReadAll(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())

			subjects := res.GetSubjects()
			Expect(subjects).Should(HaveLen(3))

			Expect(subjects[1].GetId()).To(Equal(uint64(2)))
		})

		It("Should returns error when have occurred problem in database", func() {
			db.GetConnection().Exec("DROP TABLE subjects_tracks")
			db.GetConnection().Exec("DROP TABLE subjects")
			req := &pb.SearchRequest{Search: search}
			res, err := server.ReadAll(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(res.GetSubjects()).Should(BeZero())
			Expect(res.GetPagination()).Should(BeZero())
		})

		It("Should returns all subjects where title equal English", func() {
			search.Raw = "(title[eq]:'Matématica')"
			req := &pb.SearchRequest{Search: search}
			res, err := server.ReadAll(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())

			subjects := res.GetSubjects()
			Expect(subjects).Should(HaveLen(1))

			Expect(subjects[0].GetId()).To(Equal(uint64(2)))
		})
	})

	Describe("Delete", func() {
		var item *pb.Subject

		BeforeEach(func() {
			item = &pb.Subject{
				Id: 1,
			}
		})

		It("Should delete subject", func() {
			req := &pb.SubjectRequest{Subject: item}
			res, err := server.Delete(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetSubject().GetId()).To(Equal(item.Id))

			var total int
			testutil.QueryRow("select count(id) from subjects where id=1 limit 1").Scan(&total)
			Expect(total).To(Equal(0))
		})

		It("Should error when subject not exist", func() {
			item.Id = 0
			req := &pb.SubjectRequest{Subject: item}
			res, err := server.Delete(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(res.GetSubject()).Should(BeZero())
		})
	})
})
