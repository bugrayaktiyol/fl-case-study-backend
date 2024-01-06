package users

import "github.com/gofiber/fiber/v2"

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
	users, err := h.service.ListUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// GetUserHandler retrieves a specific user
func (h *Handler) GetUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.GetUser(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

// CreateUserHandler creates a new user
func (h *Handler) CreateUserHandler(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	createdUser, err := h.service.CreateUser(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(createdUser)
}

// UpdateUserHandler updates a specific user
func (h *Handler) UpdateUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedUser User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	_, err := h.service.UpdateUser(id, &updatedUser)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(updatedUser)
}

// DeleteUserHandler deletes a specific user
func (h *Handler) DeleteUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
