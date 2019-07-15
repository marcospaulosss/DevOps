package handlers_test

import (
	"net/http"
	"testing"

	"backend/apps/severino/testutil"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/src/structs"
)

func TestPreferenceHandler(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handlers > Elearning > Preference")
}

var _ = Describe("PreferenceHandler", func() {
	var url string
	var fake structs.Preference

	BeforeEach(func() {
		url = "/api/preferences/home"
		fake = structs.Preference{
			Shelves: []uint64{1, 2, 3, 4},
		}
	})

	Context("Update", func() {
		It("Should returns error when content doesnt have shelves IDs", func() {
			fake.Shelves = []uint64{}
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithPreferenceServiceMocked("Update", fake, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusBadRequest))
			Expect(body.Err).ShouldNot(BeZero())
			Expect(body.Data).Should(BeZero())
			Expect(body.Meta).Should(BeZero())
		})

		It("Should update home configuration and return shelves ids", func() {
			req := testutil.DoPUT(url, testutil.ToJSON(fake))
			resp, body := testutil.WithPreferenceServiceMocked("Update", fake, nil).ServeHTTP(req)
			Expect(resp.StatusCode).To(Equal(http.StatusOK))
			Expect(body.Err).Should(BeZero())
			Expect(body.Data).ShouldNot(BeZero())
			Expect(body.Meta).Should(BeZero())
		})
	})
})
