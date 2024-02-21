package controllers

import (
	"fmt"
	db "go_test_backend/config"
	"go_test_backend/models"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Signup(c *fiber.Ctx) error {
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "Invalid post request",
		})
	}

	//check if name is empty
	if data["name"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "Name is required",
		})
	}

	//check if email is empty
	if data["email"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "Email is required",
		})
	}

	//check if user exist
	var user models.User
	db.DB.Where("name = ?", data["name"]).First(&user)
	if user.ID != 0 {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "User already exists",
		})
	}

	//create user
	user = models.User{
		Name:  data["name"],
		Email: data["email"],
	}
	db.DB.Create(&user)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"Message": "User created successfully",
		"data":    user,
	})
}

func Login(c *fiber.Ctx) error {
	userId := c.Query("userId")
	var user models.User
	db.DB.Where("id = ?", userId).First(&user)

	// Check if user exists
	if user.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "User not found",
		})
	}

	// Create JWT claims with additional user details
	claims := jwt.MapClaims{
		"user_id":   strconv.Itoa(int(user.ID)),
		"email":     user.Email,
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(), // Expires in 1 day
		"IssuedAt":  time.Now().Unix(),
	}

	// Create JWT token with the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	fmt.Println((token))

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Failed to generate token",
		})
	}

	// Send token in response
	userData := map[string]interface{}{
		"token": tokenString,
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Token generated successfully",
		"data":    userData,
	})
}
