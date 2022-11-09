package middleware

import (
	"crypto/rsa"
	"fmt"

	"github.com/Gurpreet-Bacancy/bcl/helper"
	"github.com/gofiber/fiber/v2"
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

		signingKey := priateKey.Public()

		return signingKey, nil
	}
}

func ValidateMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		msgpckHeader := c.Get("content-type")
		if msgpckHeader != "application/octet-stream" {
			return helper.HandleError(c, 400, nil, "Invalid messagepack request, Please provide messagepack request")
		}

		return c.Next()
	}
}
