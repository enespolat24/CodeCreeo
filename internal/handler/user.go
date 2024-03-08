package handler

import (
	"codecreeo/internal/model"
	"codecreeo/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo}
}

func (uh *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	if err := uh.userRepo.Create(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "User creation failed",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}
