package util_test

import (
	"errors"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/src/util"
)

func TestUtil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Util")
}

var _ = Describe("Util", func() {
	Describe("ToJsonString", func() {
		type SomeStruct struct {
			ID          int    `json:"id"`
			Description string `json:"desc"`
		}

		Context("When receives a struct", func() {
			It("Should returns a json string", func() {
				instance := SomeStruct{
					ID:          1,
					Description: "Hello",
				}
				expected := "{\"id\":1,\"desc\":\"Hello\"}"
				result := util.ToJsonString(instance)
				Expect(result).To(Equal(expected))
			})
		})
		Context("When receives an error", func() {
			It("Should return an empty json string", func() {
				err := errors.New("Something is wrong")
				result := util.ToJsonString(err)
				Expect(result).To(Equal("{}"))
			})
		})
		Context("When receives a string", func() {
			It("Should return the string itself", func() {
				result := util.ToJsonString("hi!")
				Expect(result).To(Equal("\"hi!\""))
			})
		})
	})

	Describe("Response", func() {
		Context("When receives an error", func() {
			It("Should put it in the error key", func() {
				err := errors.New("Something is wrong")
				result := util.Response(err)
				Expect(result["err"]).To(Equal(err))
			})
		})
	})
})
