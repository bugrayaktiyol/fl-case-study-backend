// api/routes.go
package routes

import (
	"github.com/bugrayaktiyol/fl-case-study-backend/api/users"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, userRepository *users.UserRepository) {
	userHandler := users.NewHandler(userRepository)

	app.Get("/users", userHandler.ListUsersHandler)
	app.Get("/users/:id", userHandler.GetUserHandler)
	app.Post("/users", userHandler.CreateUserHandler)
	app.Put("/users/:id", userHandler.UpdateUserHandler)
	app.Delete("/users/:id", userHandler.DeleteUserHandler)
}
