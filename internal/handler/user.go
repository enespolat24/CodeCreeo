package handler

import (
	"codecreeo/internal/model"
	"codecreeo/internal/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userRepo *repository.UserRepository
}

func NewUserHandler(userRepo *repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo}
}

func (uh *UserHandler) GetUserByID(c *fiber.Ctx) error {
	userID := c.Params("userID")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user, err := uh.userRepo.GetById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
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

func (uh *UserHandler) UpdateUser(c *fiber.Ctx) error {
	userID := c.Params("userID")

	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}
	user.ID = uint(id)

	if err := uh.userRepo.Update(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "User update failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (uh *UserHandler) DeleteUser(c *fiber.Ctx) error {
	userID := c.Params("userID")
	id, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid user ID",
		})
	}

	user := model.User{ID: uint(id)}
	if err := uh.userRepo.Delete(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "User deletion failed",
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{})
}
