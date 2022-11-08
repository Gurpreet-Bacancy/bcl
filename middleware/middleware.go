package middleware

import (
	"crypto/rsa"
	"fmt"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

// check/parse token
func CustomKeyFunc(priateKey *rsa.PrivateKey) jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		// Always check the signing method
		if t.Method.Alg() != jwtware.RS256 {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}

		// msgpckHeader := c.Get("content-type")
		// if msgpckHeader != "application/octet-stream" {
		// 	return nil, helper.HandleError(c, 400, nil, "Invalid messagepack request, Please provide messagepack request")
		// }

		// TODO checking request body type here

		signingKey := priateKey.Public()

		return signingKey, nil
	}
}
