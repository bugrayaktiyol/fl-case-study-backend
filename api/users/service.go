package users

import (
	"github.com/gofiber/fiber/v2"
)

type UserService struct {
	repository *UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (us *UserService) ListUsersHandler(c *fiber.Ctx) error {
	users, err := us.repository.FindAll()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (us *UserService) GetUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := us.repository.FindByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (us *UserService) CreateUserHandler(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	if err := us.repository.Create(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (us *UserService) UpdateUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var updatedUser User
	if err := c.BodyParser(&updatedUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid JSON"})
	}
	if err := us.repository.UpdateByID(id, &updatedUser); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(updatedUser)
}

func (us *UserService) DeleteUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := us.repository.DeleteByID(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
