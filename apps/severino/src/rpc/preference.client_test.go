package rpc_test

import (
	"context"
	"testing"

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

func TestPreferenceClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Client > Elearning > Preference")
}

var _ = Describe("PreferenceClient", func() {
	var ctx = context.Background()
	fake := structs.Preference{Type: "home", Content: `{"shelves": [1,2,3]}`}
	err := errors.NewGrpcError(errors.Internal, "Failed")

	Describe("Update", func() {
		It("Should returns updated item", func() {
			resp := &pb.PreferenceResponse{
				Preference: &pb.Preference{
					Type:    fake.Type,
					Content: fake.Content,
				},
			}
			service := testutil.MockPreferenceServiceClient("Update", ctx, resp, nil)
			client := rpc.NewPreferenceClient(service, ctx)
			result, err := client.Update(fake)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(result).ShouldNot(BeZero())

			req := &pb.PreferenceRequest{Preference: adapters.ToProtoPreference(fake)}
			mock := service.(*mocks.PreferenceServiceClient)
			mock.AssertCalled(GinkgoT(), "Update", ctx, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockPreferenceServiceClient("Update", ctx, nil, err)
			client := rpc.NewPreferenceClient(service, ctx)
			result, err := client.Update(fake)
			Expect(err).Should(HaveOccurred())
			Expect(result).Should(BeZero())
		})
	})

})
