package routes

import (
	"go_test_backend/controllers"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/swagger"

	"go_test_backend/middleware"
)

func Setup(app *fiber.App) {
	//authentication routes
	app.Post("/signup", controllers.Signup)
	app.Post("login", controllers.Login)

	//titan routes
	app.Post("/titan",middleware.JWTProtected() ,controllers.CreateTitan)
	app.Get("/titan",middleware.JWTProtected() ,controllers.GetTitans)

	//swagger

	app.Get("/swagger/*", swagger.HandlerDefault) // default

}
