package authentication_test

import (
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"backend/apps/severino/src/libs/authentication"
	"backend/apps/severino/src/structs"
	"backend/apps/severino/testutil"
	"backend/libs/configuration"
)

func TestSearch(t *testing.T) {
	RegisterFailHandler(Fail)
	if !testing.Short() {
		RunSpecs(t, "Libs > Authentication")
	}
}

var _ = Describe("Authentication", func() {
	Describe("GenerateToken", func() {
		var testUser structs.User

		BeforeEach(func() {
			testUser = structs.User{
				ID:   "666thenumberofthebest",
				Name: "Eddie",
			}
		})

		It("Should Generate a valid Jwt token", func() {
			generatedToken, err := authentication.GenerateToken(&testUser)

			Expect(err).ShouldNot(HaveOccurred())
			expectedClaims := testutil.ParseJwtToken(generatedToken)
			Expect(expectedClaims.Subject).To(Equal(testUser.ID))
		})

		It("Should error when user has no ID", func() {
			noIdUser := structs.User{
				Name: "No Id",
			}

			generatedToken, err := authentication.GenerateToken(&noIdUser)

			Expect(err).Should(HaveOccurred())
			Expect(generatedToken).To(BeEmpty())
		})
	})

	Describe("GenerateCookieWithToken", func() {
		It("Should Create cookie when token is passed", func() {
			var token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIn0.dozjgNryP4J3jVmNHl0w5N_XgL0n3I9PlFUP0THsR8U"

			cookie := authentication.GenerateCookieWithToken(token)

			Expect(cookie.Value).To(Equal(token))
			Expect(cookie.Name).To(Equal("token"))
			Expect(cookie.Path).To(Equal("/"))
			Expect(cookie.Expires).ToNot(BeNil())
			Expect(cookie.HttpOnly).To(BeTrue())
			Expect(cookie.Secure).To(BeFalse())
		})
	})

	Describe("ParseJwtToken", func() {
		var jwtSecret = configuration.Get().GetEnvConfString("jwt.secret")
		var claims jwt.StandardClaims
		var user structs.User

		BeforeEach(func() {
			user = structs.User{
				ID:   "666thenumberofthebeast",
				Name: "Eddie",
			}

			claims = jwt.StandardClaims{
				ExpiresAt: time.Now().Add(120).Unix(),
				IssuedAt:  time.Now().Unix(),
				Subject:   user.ID,
			}
		})

		It("Should parse valid JWT token", func() {
			token := testutil.GenerateCustomJwt(claims, jwtSecret)
			parsedToken, err := authentication.ParseJwtToken(token)

			Expect(err).ShouldNot(HaveOccurred())
			Expect(parsedToken.Subject).Should(Equal(user.ID))
		})

		It("Should return error when token is expired", func() {
			tokenIssueDate := time.Date(
				2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
			claims.ExpiresAt = tokenIssueDate.Add(120).Unix()
			claims.IssuedAt = tokenIssueDate.Unix()

			token := testutil.GenerateCustomJwt(claims, jwtSecret)
			parsedToken, err := authentication.ParseJwtToken(token)

			Expect(err).Should(HaveOccurred())
			Expect(parsedToken).Should(BeNil())
		})

		It("Should should return error when token signature is invalid", func() {
			token := testutil.GenerateCustomJwt(claims, "invalid_secret")
			parsedToken, err := authentication.ParseJwtToken(token)

			Expect(err).Should(HaveOccurred())
			Expect(parsedToken).Should(BeNil())
		})

		It("Should should return error when token is invalid", func() {
			token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
			parsedToken, err := authentication.ParseJwtToken(token)

			Expect(err).Should(HaveOccurred())
			Expect(parsedToken).Should(BeNil())
		})
	})
})
