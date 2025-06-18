package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/usecase"
)

type UserController struct {
	usecase usecase.UserUsecase
}

func NewUserController(router fiber.Router, u usecase.UserUsecase) {
	ctrl := &UserController{u}
	router.Get("/me", ctrl.GetProfile)
	router.Get("/", ctrl.UpdateProfile)
}

func (ctrl *UserController) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	user, err := ctrl.usecase.GetProfile(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "user not found"})
	}
	user.Password = "" // hide password
	return c.JSON(user)
}

func (ctrl *UserController) UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var input domain.User
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	input.ID = userID
	if err := ctrl.usecase.UpdateProfile(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "profile updated"})
}
