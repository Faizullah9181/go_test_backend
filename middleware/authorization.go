package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/golang-jwt/jwt"
)

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")

	// Normally Authorization HTTP header.
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}

	return ""

}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(os.Getenv("JWT_SECRET")), nil
}

func JWTProtected() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token, err := verifyToken(c)
		if err != nil {
			return jwtError(c, err)
		}
		if !token.Valid {
			return jwtError(c, jwt.ErrSignatureInvalid)
		}
		// Set the user token in the context
		c.Locals("user", token)
		return c.Next()
	}
}

func jwtError(c *fiber.Ctx, err error) error {
	return c.Status(401).JSON(fiber.Map{
		"error": "Unauthorized",
	})
}
