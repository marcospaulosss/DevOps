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

func TestTrackHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers > Elearning > Tracks")
}

var _ = Describe("TrackHandler", func() {
	var url string
	var fake structs.Track

	BeforeEach(func() {
		url = "/api/tracks"
	})

	Context("Create", func() {
		err := errors.NewError("Generic error")
		stub := structs.Track{ID: 1}

		BeforeEach(func() {
			fake = structs.Track{
				Title:       "Curso de português para o TRT",
				Description: "Curso com carga horária de 10h",
				Teachers:    "Maria",
				Duration:    300,
				Media:       "43ED6E3A-A8E2-05EE-C2B9-B86FBC133ABB.jpeg",
				Subject:     uint64(1),
			}
		})

		It("Should returns error when title length <= 5", func() {
			fake.Title = "xxz"
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithTrackServiceMocked("Create", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when track is empty", func() {
			fake.Media = ""
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithTrackServiceMocked("Create", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when service fails", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			resp, body := testutil.WithTrackServiceMocked("Create", nil, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusConflict))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns ID when item is created", func() {
			req := testutil.DoPOST(url, testutil.ToJSON(fake))
			mock := testutil.WithTrackServiceMocked("Create", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusCreated))
			saved := testutil.DecodeTrack(body.Data)
			Expect(saved.ID).To(Equal(stub.ID))
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			service := mock.ServiceContainer.Track.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "Create", fake)
		})
	})

	Context("Update", func() {
		stub := structs.Album{ID: 1}

		BeforeEach(func() {
			fake = structs.Track{
				ID:          1,
				Title:       "Curso de português para o TRT",
				Description: "Curso com carga horária de 10h",
				Teachers:    "Maria",
				Duration:    300,
				Media:       "43ED6E3A-A8E2-05EE-C2B9-B86FBC133ABB.jpeg",
				Subject:     uint64(1),
			}
		})

		It("Should returns error when ID is non numeric", func() {
			url := fmt.Sprintf("%s/xxxxx", url)
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithTrackServiceMocked("Update", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns error when ID < 1", func() {
			url := fmt.Sprintf("%s/0", url)
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithTrackServiceMocked("Update", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should returns ID when update successfully", func() {
			url := fmt.Sprintf("%s/%d", url, fake.ID)
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			mock := testutil.WithTrackServiceMocked("Update", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			data := testutil.DecodeTrack(body.Data)
			Expect(data.ID).To(Equal(fake.ID))

			service := mock.ServiceContainer.Track.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "Update", fake)
		})
	})

	Context("Delete", func() {
		stub := structs.Track{ID: 1}

		It("Should returns no content when deleted successfully", func() {
			url := fmt.Sprintf("%s/%d", url, stub.ID)
			req := testutil.DoDELETE(url)
			mock := testutil.WithTrackServiceMocked("Delete", nil, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusNoContent))
			Expect(body.Err).Should(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())

			service := mock.ServiceContainer.Track.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "Delete", stub)
		})
	})

	Context("ReadAll", func() {
		filename := "43ED6E3A-A8E2-05EE-C2B9-B86FBC133ABB"
		item1 := structs.Track{
			ID:          1,
			Title:       "Track 1",
			Description: "Curso com carga horária de 10h",
			Teachers:    "Maria",
			Duration:    300,
			Media:       filename + ".mp4",
		}
		item2 := structs.Track{
			ID:          2,
			Title:       "Track 2",
			Description: "Treinamen com carga horária de 20h",
			Teachers:    "Mariana",
			Duration:    500,
			Media:       "89EA6E9A-A8E2-05EE-C2C9-C86FCC199ACC.mp4",
		}
		items := []structs.Track{item1, item2}
		stub := structs.Result{
			Items: items,
		}

		It("Should returns items using default search params", func() {
			req := testutil.DoGET(url)
			mock := testutil.WithTrackServiceMocked("ReadAll", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())

			data := testutil.DecodeTracks(body.Data)
			Expect(data).To(HaveLen(len(items)))

			defaultPagination := structs.Pagination{
				Page:    0,
				PerPage: 0,
				Order:   "created_at",
				SortBy:  "desc",
			}
			search := structs.Search{Pagination: defaultPagination}
			service := mock.ServiceContainer.Track.(*mocks.Service)
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
			resp, body := testutil.WithTrackServiceMocked("ReadAll", stub, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())

			data := testutil.DecodeTracks(body.Data)
			Expect(data).To(HaveLen(len(items)))

			first := data[0]
			Expect(first.Title).To(Equal(item1.Title))
			Expect(first.Description).To(Equal(item1.Description))
			Expect(first.Teachers).To(Equal(item1.Teachers))
			Expect(first.Duration).To(Equal(item1.Duration))
			Expect(first.Media).To(Equal(filename + ".mp4"))

			meta := testutil.DecodePagination(body.Meta)
			Expect(meta).Should(BeEquivalentTo(pagination))
		})

		It("Should returns error and status code according of gprc code", func() {
			req := testutil.DoGET(url)
			err := errors.NewGrpcError(errors.Internal, "Something was wrong")
			resp, body := testutil.WithTrackServiceMocked("ReadAll", structs.Result{}, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusInternalServerError))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})

	Context("ReadOne", func() {
		filename := "43ED6E3A-A8E2-05EE-C2B9-B86FBC133ABB"
		stub := structs.Track{
			ID:          1,
			Title:       "Track 1",
			Description: "Curso com carga horária de 10h",
			Teachers:    "Maria",
			Duration:    300,
			Media:       filename + ".mp4",
		}

		It("Should returns item when it was found", func() {
			url := fmt.Sprintf("%s/%d", url, stub.ID)
			req := testutil.DoGET(url)
			mock := testutil.WithTrackServiceMocked("ReadOne", stub, nil)
			resp, body := mock.ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())
			data := testutil.DecodeTrack(body.Data)
			Expect(data.Media).To(Equal(filename + ".mp4"))

			item := structs.Track{ID: 1}
			service := mock.ServiceContainer.Track.(*mocks.Service)
			service.AssertCalled(GinkgoT(), "ReadOne", item)
		})

		It("Should returns error when service fails", func() {
			err := errors.NewGrpcError(errors.NotFound, "Item not found")
			url := fmt.Sprintf("%s/%d", url, stub.ID)
			req := testutil.DoGET(url)
			resp, body := testutil.WithTrackServiceMocked("ReadOne", structs.Track{}, err).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusNotFound))
			Expect(body.Data).Should(BeZero())
			Expect(body.Err).ShouldNot(BeZero())
		})
	})
})
