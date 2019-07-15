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

func TestAlbumHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers > Elearning > Album")
}

var _ = Describe("AlbumHandler", func() {
	var url string
	var fake structs.Album

	BeforeEach(func() {
		url = "/api/albums"
	})

	Context("Create", func() {
		err := errors.NewError("Generic error")
		stub := structs.Album{ID: 1}

		BeforeEach(func() {
			section := structs.Section{ID: 1}
			fake = structs.Album{
				Title:       "Curso de português para o TRT",
				Description: "Curso com carga horária de 10h",
				Image:       "43ED6E3A-A8E2-05EE-C2B9-B86FBC133ABB.jpeg",
				IsPublished: false,
				Sections:    []structs.Section{section},
			}
		})

		It("Should returns error when title length <= 1", func() {
			fake.Title = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAlbumServiceMocked("Create", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when image length <= 36", func() {
			fake.Image = "image.jpg"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAlbumServiceMocked("Create", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when service fails", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAlbumServiceMocked("Create", nil, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns ID when item is created", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithAlbumServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusCreated))
			saved := testutil.DecodeAlbum(body.Data)
			Expect(saved.ID).To(Equal(stub.ID))
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			service := mock.ServiceContainer.Album.(*mocks.AlbumService)
			service.AssertCalled(GinkgoT(), "Create", fake)
		})
	})

	Context("Update", func() {
		stub := structs.Album{ID: 1}

		BeforeEach(func() {
			section := structs.Section{ID: 1}
			fake = structs.Album{
				ID:          uint64(1),
				Title:       "Curso de português para o TRT",
				Description: "Curso com carga horária de 10h",
				Image:       "43ED6E3A-A8E2-05EE-C2B9-B86FBC133ABB.jpeg",
				IsPublished: false,
				Sections:    []structs.Section{section},
			}
		})

		It("Should returns error when ID is non numeric", func() {
			url := fmt.Sprintf("%s/xxxxx", url)
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAlbumServiceMocked("Update", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when ID < 1", func() {
			url := fmt.Sprintf("%s/0", url)
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithAlbumServiceMocked("Update", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns ID and set as IsPublished as false", func() {
			url := fmt.Sprintf("%s/%d", url, fake.ID)
			fake.IsPublished = true
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			mock := testutil.WithAlbumServiceMocked("Update", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			data := testutil.DecodeAlbum(body.Data)
			Expect(data.ID).To(Equal(fake.ID))

			service := mock.ServiceContainer.Album.(*mocks.AlbumService)
			unpublished := fake
			unpublished.IsPublished = false
			service.AssertCalled(GinkgoT(), "Update", unpublished)
		})
	})

	Context("Delete", func() {
		stub := structs.Album{ID: 1}

		It("Should returns no content when deleted successfully", func() {
			url := fmt.Sprintf("%s/%d", url, stub.ID)
			req := testutil.DoDELETE(url)
			mock := testutil.WithAlbumServiceMocked("Delete", nil, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			Expect(body.Err).Should(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			service := mock.ServiceContainer.Album.(*mocks.AlbumService)
			service.AssertCalled(GinkgoT(), "Delete", stub)
		})
	})

	Context("ReadAll", func() {
		filename := "43ED6E3A-A8E2-05EE-C2B9-B86FBC133ABB"
		section := structs.Section{ID: 1}
		item1 := structs.Album{
			ID:          1,
			Title:       "Album 1",
			Description: "Curso com carga horária de 10h",
			Image:       filename + ".jpeg",
			Sections:    []structs.Section{section},
		}
		item2 := structs.Album{
			ID:          2,
			Title:       "Album 2",
			Description: "Treinamento com carga horária de 20h",
			Image:       "89EA6E9A-A8E2-05EE-C2C9-C86FCC199ACC.jpeg",
			Sections:    []structs.Section{section},
		}
		items := []structs.Album{item1, item2}
		stub := structs.Result{
			Items: items,
		}

		It("Should returns items using default search params", func() {
			req := testutil.DoGET(url)
			mock := testutil.WithAlbumServiceMocked("ReadAll", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())

			data := testutil.DecodeAlbums(body.Data)
			Expect(data).To(HaveLen(len(items)))

			defaultPagination := structs.Pagination{
				Page:    0,
				PerPage: 0,
				Order:   "created_at",
				SortBy:  "desc",
			}
			search := structs.Search{Pagination: defaultPagination}
			service := mock.ServiceContainer.Album.(*mocks.AlbumService)
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
			resp, body := testutil.WithAlbumServiceMocked("ReadAll", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())

			data := testutil.DecodeAlbums(body.Data)
			Expect(data).To(HaveLen(len(items)))

			first := data[0]
			Expect(first.Title).To(Equal(item1.Title))
			Expect(first.Description).To(Equal(item1.Description))
			Expect(first.Image).To(Equal(item1.Image))

			meta := testutil.DecodePagination(body.Meta)
			Expect(meta).Should(BeEquivalentTo(pagination))
		})

		It("Should returns error and status code according of gprc code", func() {
			req := testutil.DoGET(url)
			err := errors.NewGrpcError(errors.Internal, "Something was wrong")
			resp, body := testutil.WithAlbumServiceMocked("ReadAll", structs.Result{}, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})

	Context("Publication", func() {
		Describe("Publish", func() {
			It("Should returns no content after publishing", func() {
				stub := structs.Album{ID: 1}
				url := fmt.Sprintf("%s/publish/%d", url, stub.ID)
				req := testutil.DoPUT(url, "")
				mock := testutil.WithAlbumServiceMocked("Publish", stub, nil)
				resp, body := mock.ServeHTTP(req)
				Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
				Expect(body.Err).Should(BeZero())
				Expect(body.Data).Should(BeZero())
				Expect(body.Meta).Should(BeZero())

				service := mock.ServiceContainer.Album.(*mocks.AlbumService)
				service.AssertCalled(GinkgoT(), "Publish", stub)
			})

			It("Should returns error when ID < 1", func() {
				stub := structs.Album{ID: 0}
				url := fmt.Sprintf("%s/publish/%d", url, stub.ID)
				req := testutil.DoPUT(url, "")
				mock := testutil.WithAlbumServiceMocked("Publish", stub, nil)
				resp, body := mock.ServeHTTP(req)
				Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(body.Err).ShouldNot(BeZero())
				Expect(body.Data).Should(BeZero())
				Expect(body.Meta).Should(BeZero())

				service := mock.ServiceContainer.Album.(*mocks.AlbumService)
				service.AssertNotCalled(GinkgoT(), "Publish")
			})

			It("Should returns error when album is not found", func() {
				stub := structs.Album{ID: 1}
				url := fmt.Sprintf("%s/publish/%d", url, stub.ID)
				req := testutil.DoPUT(url, "")
				mock := testutil.WithAlbumServiceMocked("Publish", stub, fmt.Errorf("not found"))
				resp, body := mock.ServeHTTP(req)
				Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
				Expect(body.Err).ShouldNot(BeZero())
				Expect(body.Data).Should(BeZero())
				Expect(body.Meta).Should(BeZero())

				service := mock.ServiceContainer.Album.(*mocks.AlbumService)
				service.AssertCalled(GinkgoT(), "Publish", stub)
			})
		})

		Describe("Unpublish", func() {
			It("Should returns no content after unpublishing", func() {
				stub := structs.Album{ID: 1}
				url := fmt.Sprintf("%s/unpublish/%d", url, stub.ID)
				req := testutil.DoPUT(url, "")
				mock := testutil.WithAlbumServiceMocked("Unpublish", stub, nil)
				resp, body := mock.ServeHTTP(req)
				Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
				Expect(body.Err).Should(BeZero())
				Expect(body.Data).Should(BeZero())
				Expect(body.Meta).Should(BeZero())

				service := mock.ServiceContainer.Album.(*mocks.AlbumService)
				service.AssertCalled(GinkgoT(), "Unpublish", stub)
			})

			It("Should returns error when ID < 1", func() {
				stub := structs.Album{ID: 0}
				url := fmt.Sprintf("%s/unpublish/%d", url, stub.ID)
				req := testutil.DoPUT(url, "")
				mock := testutil.WithAlbumServiceMocked("Unpublish", stub, nil)
				resp, body := mock.ServeHTTP(req)
				Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
				Expect(body.Err).ShouldNot(BeZero())
				Expect(body.Data).Should(BeZero())
				Expect(body.Meta).Should(BeZero())

				service := mock.ServiceContainer.Album.(*mocks.AlbumService)
				service.AssertNotCalled(GinkgoT(), "Unpublish")
			})

			It("Should returns error when album is not found", func() {
				stub := structs.Album{ID: 1}
				url := fmt.Sprintf("%s/unpublish/%d", url, stub.ID)
				req := testutil.DoPUT(url, "")
				mock := testutil.WithAlbumServiceMocked("Unpublish", stub, fmt.Errorf("not found"))
				resp, body := mock.ServeHTTP(req)
				Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
				Expect(body.Err).ShouldNot(BeZero())
				Expect(body.Data).Should(BeZero())
				Expect(body.Meta).Should(BeZero())

				service := mock.ServiceContainer.Album.(*mocks.AlbumService)
				service.AssertCalled(GinkgoT(), "Unpublish", stub)
			})
		})
	})

})
