package controllers

import (
	"go_test_backend/utils"

	"github.com/gofiber/fiber/v2"

	"fmt"

	"go_test_backend/models"

	"strconv"

	db "go_test_backend/config"
)

func CreateTitan(c *fiber.Ctx) error {
	// Extract the user from the token
	user, err := utils.UserExtractor(c)
	if err != nil {
		return err
	}

	fmt.Println("userID:", user.UserID)
	fmt.Println("userEmail:", user.Email)

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid post request",
		})
	}

	//get titan name age and power from the request
	name := data["name"]
	age := data["age"]
	power := data["power"]

	//check validity of the request
	if name == "" || age == "" || power == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Name, Age and Power are required",
		})
	}

	//create a new titan
	ageInt, _ := strconv.Atoi(age)
	userIDInt, err := strconv.Atoi(user.UserID)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid age value",
		})
	}

	titan := models.Titan{
		Name:        name,
		Age:         ageInt,
		Power:       power,
		CreatedByID: uint(userIDInt),
	}

	//save the titan
	result := db.DB.Create(&titan)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to create titan",
		})
	}

	// Preload the associated user
	db.DB.Preload("User").Find(&titan)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Titan created successfully",
		"data":    titan,
	})

}

func GetTitans(c *fiber.Ctx) error {
	//get user from the token
	_, err := utils.UserExtractor(c)

	if err != nil {
		return err
	}

	var titans []models.Titan
	db.DB.Preload("User").Find(&titans)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Titans retrieved successfully",
		"data":    titans,
	})
}
