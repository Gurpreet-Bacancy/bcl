package middleware

import (
	"crypto/rsa"
	"fmt"

	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

// check/parse token and check messagepack header
func CustomKeyFunc(priateKey *rsa.PrivateKey) jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		// Always check the signing method
		if t.Method.Alg() != jwtware.RS256 {
			return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
		}

		// TODO checking request body type here

		signingKey := priateKey.Public()

		return signingKey, nil
	}
}
