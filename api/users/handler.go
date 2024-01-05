package users

import (
	"github.com/gofiber/fiber/v2"
)

// Handler struct contains functions handling HTTP requests
type Handler struct {
	service *UserService
}

// NewHandler returns a new Handler instance
func NewHandler(repository *UserRepository) *Handler {
	service := NewUserService(repository)
	return &Handler{service: service}
}

// ListUsersHandler lists all users
func (h *Handler) ListUsersHandler(c *fiber.Ctx) error {
	return h.service.ListUsersHandler(c)
}

// GetUserHandler retrieves a specific user
func (h *Handler) GetUserHandler(c *fiber.Ctx) error {
	return h.service.GetUserHandler(c)
}

// CreateUserHandler creates a new user
func (h *Handler) CreateUserHandler(c *fiber.Ctx) error {
	return h.service.CreateUserHandler(c)
}

// UpdateUserHandler updates a specific user
func (h *Handler) UpdateUserHandler(c *fiber.Ctx) error {
	return h.service.UpdateUserHandler(c)
}

// DeleteUserHandler deletes a specific user
func (h *Handler) DeleteUserHandler(c *fiber.Ctx) error {
	return h.service.DeleteUserHandler(c)
}
