// fiber.go
package main

import (
	"github.com/bugrayaktiyol/fl-case-study-backend/api/routes"
	"github.com/bugrayaktiyol/fl-case-study-backend/api/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupFiberApp(userRepository *users.UserRepository) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())

	routes.SetupRoutes(app, userRepository)

	return app
}
