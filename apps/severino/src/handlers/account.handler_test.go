package handlers_test

import (
	"net/http"
	"testing"
	"time"

	"backend/libs/errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/src/structs"
	"backend/apps/severino/testutil"
)

func TestAccountHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers > Accounts > account")
}

var _ = Describe("AccountHandler", func() {
	var url string
	var fake structs.Account
	var err error

	BeforeEach(func() {
		err = errors.NewError("Generic error")
	})

	Context("Create", func() {
		stub := structs.Account{ID: "691ee763-ea36-4b15-bef4-76c03bb838c4"}

		BeforeEach(func() {
			url = "/auth/code/generate"
			fake = structs.Account{
				Email:     "testeUnitario@estrategiaconcursos.com.br",
				Phone:     "11987654321",
				Type:      "email",
				Exists:    false,
				CreatedAt: time.Now().String(),
			}
		})

		It("Should returns ID when item is created", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			saved := testutil.DecodeAccount(body.Data)
			Expect(saved.ID).To(Equal(stub.ID))
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			//service := mock.ServiceContainer.account.(*mocks.Service)
			//service.AssertCalled(GinkgoT(), "Create", fake)
		})

		It("Should returns error when json is incorrect", func() {
			req := testutil.DoPOST(url, `{"teste":"teste",}`)
			mock := testutil.WithAccountServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when type is different email or phone", func() {
			fake.Type = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when email is incorrect", func() {
			fake.Email = "teste.com.br"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone length < 11", func() {
			fake.Phone = "1198574673"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone length > 11", func() {
			fake.Phone = "1198574673122"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone is not numeric", func() {
			fake.Phone = "119888d8888"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone is not numeric", func() {
			fake.Email = ""
			fake.Phone = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when service fails", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("Create", nil, err)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})

	Context("Validate", func() {
		stub := structs.Account{ID: "691ee763-ea36-4b15-bef4-76c03bb838c4"}

		BeforeEach(func() {
			url = "/auth/code/validate"
			fake = structs.Account{
				Email:     "testeUnitario@estrategiaconcursos.com.br",
				Code:      "123456",
				Phone:     "11987654321",
				Type:      "email",
				Exists:    false,
				CreatedAt: time.Now().String(),
			}
		})

		It("Should returns ID when item is validate", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("ReadOne", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			saved := testutil.DecodeAccount(body.Data)
			Expect(saved.ID).To(Equal(stub.ID))
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when json is incorrect", func() {
			req := testutil.DoPOST(url, `{"teste":"teste",}`)
			mock := testutil.WithAccountServiceMocked("ReadOne", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when type is different email or phone", func() {
			fake.Type = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("ReadOne", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when email is incorrect", func() {
			fake.Email = "teste.com.br"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("ReadOne", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone length < 11", func() {
			fake.Phone = "1198574673"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("ReadOne", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone length > 11", func() {
			fake.Phone = "1198574673122"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("ReadOne", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone is not numeric", func() {
			fake.Phone = "119888d8888"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("ReadOne", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone is not numeric", func() {
			fake.Email = ""
			fake.Phone = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("ReadOne", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when service fails", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAccountServiceMocked("ReadOne", nil, err)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})

	Context("LoginApp", func() {
		stub := structs.Account{UserID: "691ee763-ea36-4b15-bef4-76c03bb838c4"}

		BeforeEach(func() {
			url = "/auth/login/app/validate"
			fake = structs.Account{
				Email:     "userssix@estrategiaconcursos.com.br",
				Code:      "123456",
				Phone:     "11987604121",
				Type:      "email",
				CreatedAt: time.Now().String(),
			}
		})

		It("Should returns ID when item is validate", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())
			Expect(body.Data).ShouldNot(BeZero())
			Expect(body.Data.(map[string]interface{})["id"]).ShouldNot(BeZero())
			Expect(body.Data.(map[string]interface{})["token"]).ShouldNot(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when json is incorrect", func() {
			req := testutil.DoPOST(url, `{"teste":"teste",}`)
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when email is incorrect", func() {
			fake.Email = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when type is different email or phone", func() {
			fake.Type = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when email is incorrect", func() {
			fake.Email = "teste.com.br"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone length < 11", func() {
			fake.Phone = "1198574673"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone length > 11", func() {
			fake.Phone = "1198574673122"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone is not numeric", func() {
			fake.Phone = "119888d8888"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone is not numeric", func() {
			fake.Email = ""
			fake.Phone = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when service fails", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", nil, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})

	Context("LoginWeb", func() {
		stub := structs.Account{UserID: "691ee763-ea36-4b15-bef4-76c03bb838c4"}

		BeforeEach(func() {
			url = "/auth/login/web/validate"
			fake = structs.Account{
				Email:     "userssix@estrategiaconcursos.com.br",
				Code:      "123456",
				Phone:     "11987604121",
				Type:      "email",
				CreatedAt: time.Now().String(),
			}
		})

		It("Should returns ID when item is validate", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())
			Expect(body.Data).ShouldNot(BeZero())
			Expect(body.Data.(map[string]interface{})["id"]).ShouldNot(BeZero())
			Expect(resp.Header.Get("Set-Cookie")).ShouldNot(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when json is incorrect", func() {
			req := testutil.DoPOST(url, `{"teste":"teste",}`)
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when email is incorrect", func() {
			fake.Email = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when type is different email or phone", func() {
			fake.Type = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when email is incorrect", func() {
			fake.Email = "teste.com.br"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone length < 11", func() {
			fake.Phone = "1198574673"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone length > 11", func() {
			fake.Phone = "1198574673122"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone is not numeric", func() {
			fake.Phone = "119888d8888"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when phone is not numeric", func() {
			fake.Email = ""
			fake.Phone = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when service fails", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAccountServiceMocked("ReadOne", nil, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})
})
