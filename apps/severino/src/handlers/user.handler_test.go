package handlers_test

import (
	"net/http"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/mocks"
	"backend/apps/severino/src/structs"
	"backend/apps/severino/testutil"
	"backend/libs/errors"
)

func TestUserHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers > Accounts > Users")
}

var _ = Describe("UserHandler", func() {
	var url string
	var fake structs.User

	Context("Create", func() {
		err := errors.NewError("Generic error")
		stub := structs.User{
			ID: "8b306588-c12f-4ed3-9257-3d2054c3c4f0",
		}

		BeforeEach(func() {
			url = "/auth/users"
			fake = structs.User{
				Name:      "Maria das Graças Silva",
				Phone:     "11988765543",
				PhoneCode: "333444",
				Email:     "gracas@mariasss.com.br",
				EmailCode: "555666",
			}
		})

		It("Should returns error when email is not provided", func() {
			fake.Email = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithUserServiceMocked("Create", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
		It("Should returns error when phone code have len < 6", func() {
			fake.PhoneCode = "123"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithUserServiceMocked("Create", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
		It("Should returns error when service fails", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithUserServiceMocked("Create", nil, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
		It("Should returns user ID and token when it was successfully created", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithUserServiceMocked("Create", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusCreated))
			saved := testutil.DecodeUser(body.Data)
			Expect(saved.ID).To(Equal(stub.ID))
			Expect(saved.Token).ToNot(BeEmpty())
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})

	Context("ReadAll", func() {
		items := []structs.User{
			{
				ID:        "14aa53dd-7437-4716-9ab3-4478b7259552",
				Name:      "Ricardo dos Santos Silva",
				Phone:     "5511934456778",
				Email:     "ricardo.santos@gggmail.com",
				Active:    false,
				CreatedAt: time.Now().String(),
				UpdatedAt: time.Now().String(),
			},
			{
				ID:        "98f97764-2502-4cc9-a3f7-372a6cf958e9",
				Name:      "Camila das Neves Reis",
				Phone:     "5511976659008",
				Email:     "camila.reis@gggmail.com",
				Active:    true,
				CreatedAt: time.Now().String(),
				UpdatedAt: time.Now().String(),
			},
			{
				ID:        "a32242e3-8473-4bd9-a17f-c4caf67af224",
				Name:      "Astolfo de Jesus",
				Phone:     "5511990087665",
				Email:     "astolfo.jesus@gggmail.com",
				Active:    true,
				CreatedAt: time.Now().String(),
				UpdatedAt: time.Now().String(),
			},
		}
		stub := structs.Result{
			Items: items,
		}

		BeforeEach(func() {
			url = "/api/users"
		})

		It("Should return error when service fails", func() {
			req := testutil.DoGET(url)
			err := errors.NewGrpcError(errors.Internal, "Something was wrong")
			resp, body := testutil.WithUserServiceMocked("ReadAll", structs.Result{}, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
		It("Should return users when using default parameters", func() {
			req := testutil.DoGET(url)
			mock := testutil.WithUserServiceMocked("ReadAll", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())

			data := testutil.DecodeUsers(body.Data)
			Expect(data).To(HaveLen(len(items)))

			defaultPagination := structs.Pagination{
				Page:    0,
				PerPage: 0,
				Order:   "created_at",
				SortBy:  "desc",
			}
			search := structs.Search{Pagination: defaultPagination}
			service := mock.ServiceContainer.User.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "ReadAll", search)
		})
		It("Should return users when using pagination", func() {
			req := testutil.DoGET(url)
			pagination := structs.Pagination{
				Page:    10,
				PerPage: 100,
				Order:   "created_at",
				SortBy:  "asc",
			}
			stub.Pagination = pagination
			resp, body := testutil.WithUserServiceMocked("ReadAll", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())

			data := testutil.DecodeUsers(body.Data)
			Expect(data).To(HaveLen(len(items)))

			first := data[0]
			Expect(first.ID).To(Equal(items[0].ID))
			Expect(first.Name).To(Equal(items[0].Name))
			Expect(first.Phone).To(Equal(items[0].Phone))
			Expect(first.Email).To(Equal(items[0].Email))

			meta := testutil.DecodePagination(body.Meta)
			Expect(meta).Should(BeEquivalentTo(pagination))
		})
	})
})
