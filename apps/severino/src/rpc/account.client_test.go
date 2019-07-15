package rpc_test

import (
	"context"
	"testing"

	"backend/libs/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/mocks"
	"backend/apps/severino/src/adapters"
	"backend/apps/severino/src/rpc"
	"backend/apps/severino/src/structs"
	"backend/apps/severino/testutil"
	pb "backend/proto"
)

var ctxSub = context.Background()

func TestAccountClient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Client > Accounts > Account")
}

var _ = Describe("AccountClient", func() {

	fake := structs.Account{
		Email: "testeUnitario@estrategiaconcursos.com.br",
		Type:  "email",
	}
	err := errors.NewGrpcError(errors.Internal, "Failed")

	Context("Create", func() {
		It("Should returns created item", func() {
			resp := &pb.AccountResponse{
				Account: &pb.Account{Id: "691ee763-ea36-4b15-bef4-76c03bb838c4"},
			}
			service := testutil.MockAccountServiceClient("Create", ctxSub, resp, nil)
			client := rpc.NewAccountClient(service, ctxSub)
			result, err := client.Create(&fake)
			Expect(err).Should(BeZero())
			Expect(result).ShouldNot(BeZero())
			Expect(result.(structs.Account).ID).To(Equal(resp.Account.Id))

			req := &pb.AccountRequest{Account: adapters.ToProtoAccount(fake)}
			mock := service.(*mocks.AccountServiceClient)
			mock.AssertCalled(GinkgoT(), "Create", ctxSub, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockAccountServiceClient("Create", ctxSub, nil, err)
			client := rpc.NewAccountClient(service, ctxSub)
			result, err := client.Create(&fake)
			Expect(err).ShouldNot(BeZero())
			Expect(result).Should(BeZero())
		})
	})

	Context("ReadOne", func() {
		It("Should returns validate item", func() {
			resp := &pb.AccountResponse{
				Account: &pb.Account{Id: "691ee763-ea36-4b15-bef4-76c03bb838c4"},
			}
			service := testutil.MockAccountServiceClient("ReadOne", ctxSub, resp, nil)
			client := rpc.NewAccountClient(service, ctxSub)
			result, err := client.ReadOne(&fake)
			Expect(err).Should(BeZero())
			Expect(result).ShouldNot(BeZero())
			Expect(result.(structs.Account).ID).To(Equal(resp.Account.Id))

			req := &pb.AccountRequest{Account: adapters.ToProtoAccount(fake)}
			mock := service.(*mocks.AccountServiceClient)
			mock.AssertCalled(GinkgoT(), "ReadOne", ctxSub, req)
		})

		It("Should returns error when the server call failed", func() {
			service := testutil.MockAccountServiceClient("ReadOne", ctxSub, nil, err)
			client := rpc.NewAccountClient(service, ctxSub)
			result, err := client.ReadOne(&fake)
			Expect(err).ShouldNot(BeZero())
			Expect(result).Should(BeZero())
		})
	})
})
