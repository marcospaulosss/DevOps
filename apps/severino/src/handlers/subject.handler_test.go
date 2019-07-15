package handlers_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"backend/apps/severino/mocks"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/src/structs"
	"backend/apps/severino/testutil"
	"backend/libs/errors"
)

func TestSubjectHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers > Elearning > Subjects")
}

var _ = Describe("SubjectHandler", func() {
	var url string
	var fake structs.Subject
	var err error
	var stub structs.Subject

	BeforeEach(func() {
		err = errors.NewError("Generic error")
		stub = structs.Subject{ID: 1}

		fake = structs.Subject{
			Title: "História",
		}
	})

	Context("Create", func() {

		BeforeEach(func() {
			url = "/api/subjects"
		})

		It("Should returns ID when item is created", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithSubjectServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusCreated))
			saved := testutil.DecodeSubject(body.Data)
			Expect(saved.ID).To(Equal(stub.ID))
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			service := mock.ServiceContainer.Subject.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "Create", fake)
		})

		It("Should returns error when title length < 1", func() {
			fake.Title = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithSubjectServiceMocked("Create", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when service fails", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithSubjectServiceMocked("Create", nil, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})

	Context("Update", func() {
		BeforeEach(func() {
			fake.ID = uint64(1)
			url = "/api/subjects/1"
		})

		It("Should returns ID when item is updated", func() {
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			mock := testutil.WithSubjectServiceMocked("Update", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			saved := testutil.DecodeSubject(body.Data)
			Expect(saved.ID).To(Equal(stub.ID))
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			service := mock.ServiceContainer.Subject.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "Update", fake)
		})

		It("Should returns error when title length < 1", func() {
			fake.Title = ""
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithSubjectServiceMocked("Update", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when id < 1", func() {
			url = "/api/subjects/0"
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithSubjectServiceMocked("Update", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when id is string", func() {
			url = "/api/subjects/a"
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithSubjectServiceMocked("Update", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when service fails", func() {
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithSubjectServiceMocked("Update", nil, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})

	Context("ReadOne", func() {
		BeforeEach(func() {
			fake.ID = uint64(1)
			fake.Title = ""
			url = "/api/subjects/1"
			stub = structs.Subject{
				ID:        1,
				Title:     "Português",
				CreatedAt: time.Now().String(),
			}
		})

		It("Should returns subject when item is found", func() {
			req := testutil.DoGET(url)
			mock := testutil.WithSubjectServiceMocked("ReadOne", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			saved := testutil.DecodeSubject(body.Data)
			Expect(saved.ID).To(Equal(stub.ID))
			Expect(saved.Title).To(Equal(stub.Title))
			Expect(saved.CreatedAt).ShouldNot(BeZero())
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			service := mock.ServiceContainer.Subject.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "ReadOne", fake)
		})

		It("Should returns error when id is string", func() {
			url = "/api/subjects/a"
			req := testutil.DoGET(url)
			resp, body := testutil.WithSubjectServiceMocked("ReadOne", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when service fails", func() {
			req := testutil.DoGET(url)
			resp, body := testutil.WithSubjectServiceMocked("ReadOne", nil, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})

	Context("ReadAll", func() {
		BeforeEach(func() {
			url = "/api/subjects"
		})

		item1 := structs.Subject{
			ID:    1,
			Title: "Português",
		}
		item2 := structs.Subject{
			ID:    2,
			Title: "Inglês",
		}
		items := []structs.Subject{item1, item2}
		stub := structs.Result{
			Items: items,
		}

		It("Should returns items using default search params", func() {
			req := testutil.DoGET(url)
			mock := testutil.WithSubjectServiceMocked("ReadAll", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())

			data := testutil.DecodeSubjects(body.Data)
			Expect(data).To(HaveLen(len(items)))

			defaultPagination := structs.Pagination{
				Page:    0,
				PerPage: 0,
				Order:   "created_at",
				SortBy:  "desc",
			}
			search := structs.Search{Pagination: defaultPagination}
			service := mock.ServiceContainer.Subject.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "ReadAll", search)
		})

		It("Should returns pagination data in response's meta", func() {
			req := testutil.DoGET(url)
			pagination := structs.Pagination{
				Page:    10,
				PerPage: 100,
				Order:   "created_at",
				SortBy:  "asc",
			}
			stub.Pagination = pagination
			resp, body := testutil.WithSubjectServiceMocked("ReadAll", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())

			data := testutil.DecodeSubjects(body.Data)
			Expect(data).To(HaveLen(len(items)))

			meta := testutil.DecodePagination(body.Meta)
			Expect(meta).Should(BeEquivalentTo(pagination))
		})

		It("Should returns error and status code according of gprc code", func() {
			req := testutil.DoGET(url)
			err := errors.NewGrpcError(errors.Internal, "Something was wrong")
			resp, body := testutil.WithSubjectServiceMocked("ReadAll", structs.Result{}, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})

	Context("Delete", func() {
		It("Should returns no content when deleted successfully", func() {
			url := fmt.Sprintf("%s/%d", url, stub.ID)
			req := testutil.DoDELETE(url)
			mock := testutil.WithSubjectServiceMocked("Delete", nil, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			Expect(body.Err).Should(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			service := mock.ServiceContainer.Subject.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "Delete", stub)
		})

		It("Should returns error when id < 1", func() {
			url = "/api/subjects/0"
			req := testutil.DoDELETE(url)
			resp, body := testutil.WithSubjectServiceMocked("Delete", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when id is string", func() {
			url = "/api/subjects/a"
			req := testutil.DoDELETE(url)
			resp, body := testutil.WithSubjectServiceMocked("Delete", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when service fails", func() {
			url = "/api/subjects/1"
			req := testutil.DoDELETE(url)
			resp, body := testutil.WithSubjectServiceMocked("Delete", nil, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})
})
