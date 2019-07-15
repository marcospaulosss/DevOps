package rpc_test

import (
	"context"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"backend/apps/accounts/src/repositories"
	"backend/apps/accounts/src/rpc"
	"backend/libs/configuration"
	"backend/libs/databases"
	testutil "backend/libs/testing"
	pb "backend/proto"
)

func TestAccountServer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RPC Server > Account")
}

type NotificationMock struct {
	mock.Mock
}

func (_m *NotificationMock) SendEmailNotification(email string, code string) {
	_m.Called(email, code)
}
func (_m *NotificationMock) SendSmsNotification(phone string, code string) (string, error) {
	ret := _m.Called(phone, code)
	return ret.String(0), ret.Error(1)
}

var _ = Describe("AccountServer", func() {
	ctx := context.Background()
	var db databases.Database
	var server *rpc.AccountServer

	config := configuration.Get()
	databaseURL := config.GetEnvConfString("database_url")

	BeforeEach(func() {
		db = testutil.NewTestDB(databaseURL)
		testutil.Seed(db, "accounts.sql")

		notificationMock := new(NotificationMock)
		notificationMock.On("SendEmailNotification", mock.Anything, mock.Anything)
		notificationMock.On("SendSmsNotification", mock.Anything, mock.Anything).Return("1155123451234", nil)

		repos := repositories.Container{
			AccountRepository: repositories.NewAccountRepository(db),
			UserRepository:    repositories.NewUserRepository(db),
			Notifications:     notificationMock,
		}
		server = rpc.NewAccountServer(nil, repos)
	})

	Describe("Create", func() {
		var email *pb.Account
		var phone *pb.Account

		BeforeEach(func() {
			email = &pb.Account{
				Type:  "email",
				Email: "testeUnitario@estrategiaconcursos.com.br",
			}

			phone = &pb.Account{
				Type:  "phone",
				Phone: "5511989988989",
			}
		})

		It("Should create the account to email", func() {
			req := &pb.AccountRequest{Account: email}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetAccount().GetId()).ShouldNot(BeZero())
		})

		It("Should create the account to phone", func() {
			req := &pb.AccountRequest{Account: phone}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetAccount().GetId()).ShouldNot(BeZero())
		})

		It("Should return error in get the account existing", func() {
			email.Email = "code@estrategiaconcursos.com.br"
			req := &pb.AccountRequest{Account: email}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetAccount().GetId()).ShouldNot(BeZero())
		})

		It("Should return error in get the account existing and user registered", func() {
			email.Email = "usersfirst@estrategiaconcursos.com.br"
			req := &pb.AccountRequest{Account: email}
			res, err := server.Create(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(res).Should(BeNil())
		})

		It("Should return error in create the account but existing and user registered", func() {
			phone.Phone = "4511999999999"
			req := &pb.AccountRequest{Account: phone}
			res, err := server.Create(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(res).Should(BeNil())
		})

		It("Should create the account to email login", func() {
			email.Email = "usersthird@estrategiaconcursos.com.br"
			email.Exists = true
			req := &pb.AccountRequest{Account: email}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetAccount().GetId()).ShouldNot(BeZero())
		})

		It("Should create the account to email login resend", func() {
			email.Email = "usersfirst@estrategiaconcursos.com.br"
			email.Exists = true
			req := &pb.AccountRequest{Account: email}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetAccount().GetId()).ShouldNot(BeZero())
		})

		It("Should create the account to phone login", func() {
			phone.Phone = "4511999999998"
			phone.Exists = true
			req := &pb.AccountRequest{Account: phone}
			res, err := server.Create(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetAccount().GetId()).ShouldNot(BeZero())
		})

		It("Should error when user not found", func() {
			email.Email = "usersfirst@estrategiaconcursos.com.b"
			email.Exists = true
			req := &pb.AccountRequest{Account: email}
			res, err := server.Create(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(res).Should(BeNil())
		})
	})

	Describe("ReadOne", func() {
		var item *pb.Account

		BeforeEach(func() {
			item = &pb.Account{
				Type:      "email",
				Email:     "code@estrategiaconcursos.com.br",
				EmailCode: "987654",
				Phone:     "5511999999999",
				PhoneCode: "123456",
			}
		})

		It("Should validate the account to email", func() {
			req := &pb.AccountRequest{Account: item}
			res, err := server.ReadOne(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetAccount().GetId()).ShouldNot(BeZero())
			Expect(res.GetAccount().GetEmail()).To(Equal(item.Email))
			Expect(res.GetAccount().GetEmailCode()).To(Equal(item.EmailCode))
			Expect(res.GetAccount().GetCreatedAt()).ShouldNot(BeZero())
		})

		It("Should validate the account to phone", func() {
			item.Type = "phone"
			req := &pb.AccountRequest{Account: item}
			res, err := server.ReadOne(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetAccount().GetId()).ShouldNot(BeZero())
			Expect(res.GetAccount().GetPhone()).To(Equal(item.Phone))
			Expect(res.GetAccount().GetPhoneCode()).To(Equal(item.PhoneCode))
			Expect(res.GetAccount().GetCreatedAt()).ShouldNot(BeZero())
		})

		It("Should return error in validate when email not found", func() {
			item.Email = "code@estrategiaconcursos.com.b"
			req := &pb.AccountRequest{Account: item}
			res, err := server.ReadOne(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(res).Should(BeZero())
		})

		It("Should return error in get the account existing and user registered", func() {
			item.Exists = true
			req := &pb.AccountRequest{Account: item}
			res, err := server.ReadOne(ctx, req)
			Expect(err).Should(HaveOccurred())
			Expect(res).Should(BeNil())
		})

		It("Should validate the account to email and deleted account", func() {
			item.Exists = true
			item.Email = "usersfirst@estrategiaconcursos.com.br"
			item.EmailCode = "123456"
			req := &pb.AccountRequest{Account: item}
			res, err := server.ReadOne(ctx, req)
			Expect(err).ShouldNot(HaveOccurred())
			Expect(res.GetAccount().GetId()).ShouldNot(BeZero())
			Expect(res.GetAccount().GetEmail()).To(Equal(item.Email))
			Expect(res.GetAccount().GetEmailCode()).To(Equal(item.EmailCode))
			Expect(res.GetAccount().GetCreatedAt()).ShouldNot(BeZero())
		})
	})
})
