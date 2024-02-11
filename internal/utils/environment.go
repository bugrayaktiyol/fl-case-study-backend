package utils

import (
	_ "github.com/bugrayaktiyol/fl-case-study-backend/docs"
	"github.com/bugrayaktiyol/fl-case-study-backend/internal/users"
	"github.com/bugrayaktiyol/fl-case-study-backend/internal/users/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	swagger "github.com/gofiber/swagger"
)

func SetupFiberApp(userHandler *users.Handler) *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(logger.New())

	routes.SetupRoutes(app, userHandler)

	return app
}
