package handlers_test

import (
	"fmt"
	"net/http"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/mocks"
	"backend/apps/severino/src/structs"
	"backend/apps/severino/testutil"
	"backend/libs/errors"
)

func TestShelfHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers > Elearning > Shelves")
}

var _ = Describe("ShelfHandler", func() {
	var url string
	var fake structs.Shelf

	BeforeEach(func() {
		url = "/api/shelves"
	})

	Context("Create", func() {
		err := errors.NewError("Generic error")
		stub := structs.Shelf{ID: 1}

		BeforeEach(func() {
			fake = structs.Shelf{
				Title: "Receita Federal",
			}
		})

		It("Should returns error when title length < 1", func() {
			fake.Title = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithShelfServiceMocked("Create", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when title is empty", func() {
			fake.Title = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithShelfServiceMocked("Create", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when service fails", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithShelfServiceMocked("Create", nil, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns ID when item is created", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithShelfServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusCreated))
			saved := testutil.DecodeShelf(body.Data)
			Expect(saved.ID).To(Equal(stub.ID))
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			service := mock.ServiceContainer.Shelf.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "Create", fake)
		})
	})

	Context("Update", func() {
		stub := structs.Album{ID: 1}

		BeforeEach(func() {
			fake = structs.Shelf{
				ID:    1,
				Title: "Receita Federal",
			}
		})

		It("Should returns error when ID is non numeric", func() {
			url := fmt.Sprintf("%s/xxxxx", url)
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithShelfServiceMocked("Update", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when ID < 1", func() {
			url := fmt.Sprintf("%s/0", url)
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithShelfServiceMocked("Update", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns ID when update successfully", func() {
			url := fmt.Sprintf("%s/%d", url, fake.ID)
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			mock := testutil.WithShelfServiceMocked("Update", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			data := testutil.DecodeShelf(body.Data)
			Expect(data.ID).To(Equal(fake.ID))

			service := mock.ServiceContainer.Shelf.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "Update", fake)
		})
	})

	Context("Delete", func() {
		stub := structs.Shelf{ID: 1}

		It("Should returns no content when deleted successfully", func() {
			url := fmt.Sprintf("%s/%d", url, stub.ID)
			req := testutil.DoDELETE(url)
			mock := testutil.WithShelfServiceMocked("Delete", nil, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			Expect(body.Err).Should(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			service := mock.ServiceContainer.Shelf.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "Delete", stub)
		})
	})

	Context("ReadAll", func() {
		item1 := structs.Shelf{
			ID:    1,
			Title: "Receita Federal",
		}
		item2 := structs.Shelf{
			ID:    2,
			Title: "TRT",
		}
		items := []structs.Shelf{item1, item2}
		stub := structs.Result{
			Items: items,
		}

		It("Should returns items using default search params", func() {
			req := testutil.DoGET(url)
			mock := testutil.WithShelfServiceMocked("ReadAll", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())

			data := testutil.DecodeShelves(body.Data)
			Expect(data).To(HaveLen(len(items)))

			defaultPagination := structs.Pagination{
				Page:    0,
				PerPage: 0,
				Order:   "created_at",
				SortBy:  "desc",
			}
			search := structs.Search{Pagination: defaultPagination}
			service := mock.ServiceContainer.Shelf.(*mocks.Service)
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
			resp, body := testutil.WithShelfServiceMocked("ReadAll", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())

			data := testutil.DecodeShelves(body.Data)
			Expect(data).To(HaveLen(len(items)))

			meta := testutil.DecodePagination(body.Meta)
			Expect(meta).Should(BeEquivalentTo(pagination))
		})

		It("Should returns error and status code according of gprc code", func() {
			req := testutil.DoGET(url)
			err := errors.NewGrpcError(errors.Internal, "Something was wrong")
			resp, body := testutil.WithShelfServiceMocked("ReadAll", structs.Result{}, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})

	Context("ReadOne", func() {
		stub := structs.Shelf{
			ID:    1,
			Title: "Shelf 1",
		}

		It("Should returns item when it was found", func() {
			url := fmt.Sprintf("%s/%d", url, stub.ID)
			req := testutil.DoGET(url)
			mock := testutil.WithShelfServiceMocked("ReadOne", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())
			data := testutil.DecodeShelf(body.Data)
			Expect(data.Title).To(Equal(stub.Title))

			item := structs.Shelf{ID: 1}
			service := mock.ServiceContainer.Shelf.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "ReadOne", item)
		})

		It("Should returns error when service fails", func() {
			err := errors.NewGrpcError(errors.NotFound, "Item not found")
			url := fmt.Sprintf("%s/%d", url, stub.ID)
			req := testutil.DoGET(url)
			resp, body := testutil.WithShelfServiceMocked("ReadOne", structs.Shelf{}, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
			Expect(body.Data).Should(BeZero())
			Expect(body.Err).ShouldNot(BeZero())
		})
	})
})
