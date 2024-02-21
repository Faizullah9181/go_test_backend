package main

import (
	"fmt"
	db "go_test_backend/config"
	"go_test_backend/routes"

	_ "go_test_backend/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title Fiber Example API
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email fiber@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3030
// @BasePath /
func main() {
	fmt.Println("app running...")
	//db connection

	app := fiber.New()
	routes.Setup(app)
	db.Connect()
	app.Use(cors.New())


	app.Listen(":3030")
}
