package main

import (
	"github.com/bugrayaktiyol/fl-case-study-backend/api/routes"
	"github.com/bugrayaktiyol/fl-case-study-backend/api/users"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupFiberApp(userRepository *users.UserRepository) *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))
	app.Use(logger.New())

	routes.SetupRoutes(app, userRepository)

	return app
}
