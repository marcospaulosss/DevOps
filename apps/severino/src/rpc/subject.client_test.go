package rpc_test

import (
	"context"
	"testing"
	"time"

	"backend/apps/severino/src/structs"
	"backend/libs/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/mocks"
	"backend/apps/severino/src/adapters"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/testutil"
	pb "backend/proto"
)

func TestSubjectClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Client > Elearning > Subject")
}

var _ = Describe("SubjectClient", func() {
	var ctx = context.Background()
	fake := structs.Subject{Title: "subject 1"}
	err := errors.NewGrpcError(errors.Internal, "Failed")

	var resp pb.SubjectResponse

	BeforeEach(func() {
		resp.Subject = &pb.Subject{Id: 1}
	})

	Context("Create", func() {
		It("Should returns created item", func() {
			service := testutil.MockSubjectServiceClient("Create", ctx, &resp, nil)
			client := rpc.NewSubjectClient(service, ctx)
			result, err := client.Create(fake)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.(structs.Subject).ID).To(Equal(resp.GetSubject().GetId()))

			req := &pb.SubjectRequest{Subject: adapters.ToProtoSubject(fake)}
			mock := service.(*mocks.SubjectServiceClient)
			mock.AssertCalled(GinkgoT(), "Create", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockSubjectServiceClient("Create", ctx, nil, err)
			client := rpc.NewSubjectClient(service, ctx)
			result, err := client.Create(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("Update", func() {
		It("Should returns updated item", func() {
			service := testutil.MockSubjectServiceClient("Update", ctx, &resp, nil)
			client := rpc.NewSubjectClient(service, ctx)
			result, err := client.Update(fake)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.(structs.Subject).ID).To(Equal(resp.GetSubject().GetId()))

			req := &pb.SubjectRequest{Subject: adapters.ToProtoSubject(fake)}
			mock := service.(*mocks.SubjectServiceClient)
			mock.AssertCalled(GinkgoT(), "Update", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockSubjectServiceClient("Update", ctx, nil, err)
			client := rpc.NewSubjectClient(service, ctx)
			result, err := client.Update(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("ReadOne", func() {
		BeforeEach(func() {
			resp.Subject = &pb.Subject{
				Id:        1,
				Title:     "Teste Unitário",
				CreatedAt: time.Now().String(),
			}
		})

		It("Should returns found item", func() {
			service := testutil.MockSubjectServiceClient("ReadOne", ctx, &resp, nil)
			client := rpc.NewSubjectClient(service, ctx)
			result, err := client.ReadOne(fake)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result.(structs.Subject).ID).To(Equal(resp.GetSubject().GetId()))
			Expect(result.(structs.Subject).Title).To(Equal(resp.GetSubject().GetTitle()))
			Expect(result.(structs.Subject).CreatedAt).To(Equal(resp.GetSubject().GetCreatedAt()))

			req := &pb.SubjectRequest{Subject: adapters.ToProtoSubject(fake)}
			mock := service.(*mocks.SubjectServiceClient)
			mock.AssertCalled(GinkgoT(), "ReadOne", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockSubjectServiceClient("ReadOne", ctx, nil, err)
			client := rpc.NewSubjectClient(service, ctx)
			result, err := client.ReadOne(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

	Context("ReadAll", func() {
		It("Should returns an item list", func() {
			subjects := []*pb.Subject{
				{
					Title: "Português",
				},
			}

			resp := &pb.SubjectsResponse{Subjects: subjects}
			service := testutil.MockSubjectServiceClientSearch("ReadAll", ctx, resp, nil)
			client := rpc.NewSubjectClient(service, ctx)

			search := structs.Search{
				Pagination: structs.Pagination{
					PerPage: 9,
					Page:    1,
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
						Page:    1,
						Order:   "id",
						SortBy:  "asc",
					},
					Raw: "",
				},
			}
			mock := service.(*mocks.SubjectServiceClient)
			mock.AssertCalled(GinkgoT(), "ReadAll", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockSubjectServiceClientSearch("ReadAll", ctx, nil, err)
			client := rpc.NewSubjectClient(service, ctx)
			result, err := client.ReadAll(structs.Search{})
			Expect(err).ShouldNot(BeZero())
			Expect(result).Should(BeZero())
		})
	})

	Context("Delete", func() {
		It("Should returns deleted item", func() {
			resp := &pb.SubjectResponse{
				Subject: &pb.Subject{Id: 1},
			}
			service := testutil.MockSubjectServiceClient("Delete", ctx, resp, nil)
			client := rpc.NewSubjectClient(service, ctx)
			fake.ID = 1
			result, err := client.Delete(fake)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeNil())

			req := &pb.SubjectRequest{Subject: adapters.ToProtoSubject(fake)}
			mock := service.(*mocks.SubjectServiceClient)
			mock.AssertCalled(GinkgoT(), "Delete", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockSubjectServiceClient("Delete", ctx, nil, err)
			client := rpc.NewSubjectClient(service, ctx)
			result, err := client.Delete(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})
})
