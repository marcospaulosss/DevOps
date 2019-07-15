package authentication

import (
	"errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"

	"backend/apps/severino/src/structs"
	"backend/libs/configuration"
)

const CookieName string = "token"

func GenerateToken(user *structs.User) (string, error) {
	if len(user.ID) == 0 {
		return "", errors.New("User has no ID")
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set custom expiration time
	conf := configuration.Get()
	expiration := (time.Hour * 24) * 90
	if confExpiration := conf.GetEnvConfInteger("jwt.exp"); confExpiration > 0 {
		expiration = time.Second * time.Duration(confExpiration)
	}

	// Set claims
	token.Claims = jwt.StandardClaims{
		ExpiresAt: time.Now().Add(expiration).Unix(), // Token expiration date
		IssuedAt:  time.Now().Unix(),                 // Token generation date
		Subject:   user.ID,                           // User ID for whom the token was generated for (subject)
	}

	// Generate encoded token and send it as response.
	jwtSecret := conf.GetEnvConfString("jwt.secret")
	generatedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return generatedToken, nil
}

func ParseJwtToken(tokenString string) (*jwt.StandardClaims, error) {
	jwtSecret := configuration.Get().GetEnvConfString("jwt.secret")

	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(*jwt.StandardClaims), nil
}

func GenerateCookieWithToken(token string) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Path = "/"
	cookie.Name = CookieName
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.HttpOnly = true
	cookie.Secure = false
	return cookie
}
