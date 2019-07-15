package testutil

import (
	"github.com/dgrijalva/jwt-go"

	"backend/libs/configuration"
)

func GenerateCustomJwt(claims jwt.StandardClaims, secret string) string {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = claims

	str, err := token.SignedString([]byte(secret))
	if err != nil {
		panic("Testutil failed generating token")
	}
	return str
}

func ParseJwtToken(tokenString string) *jwt.StandardClaims {
	secret := configuration.Get().GetEnvConfString("jwt.secret")
	token, _ := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secret), nil
	})

	return token.Claims.(*jwt.StandardClaims)
}
