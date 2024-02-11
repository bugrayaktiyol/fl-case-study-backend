package routes

import (
	"github.com/bugrayaktiyol/fl-case-study-backend/internal/users"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userHandler *users.Handler) {

	app.Get("/api/users", userHandler.ListUsersHandler)
	app.Get("/api/users/:id", userHandler.GetUserHandler)
	app.Post("/api/users", userHandler.CreateUserHandler)
	app.Put("/api/users/:id", userHandler.UpdateUserHandler)
	app.Delete("/api/users/:id", userHandler.DeleteUserHandler)

}
