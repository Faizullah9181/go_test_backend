package utils

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type User struct {
	UserID string
	Email  string
}

func UserExtractor(c *fiber.Ctx) (*User, error) {
	claims, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return nil, c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "User token not found in context",
		})
	}

	if claims == nil || !claims.Valid {
		return nil, c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
		})
	}

	claimsMap, ok := claims.Claims.(jwt.MapClaims)
	if !ok {
		return nil, c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to parse token claims",
		})
	}

	userID, ok := claimsMap["user_id"].(string)
	if !ok {
		return nil, c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to extract userID from token claims",
		})
	}

	userEmail, ok := claimsMap["email"].(string)
	if !ok {
		return nil, c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to extract userEmail from token claims",
		})
	}

	return &User{
		UserID: userID,
		Email:  userEmail,
	}, nil
}
