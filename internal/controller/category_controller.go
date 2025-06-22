package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/usecase"
)

type CategoryController struct {
	usecase usecase.CategoryUsecase
}

func NewCategoryController(router fiber.Router, usecase usecase.CategoryUsecase) {
	ctrl := &CategoryController{usecase}
	router.Post("/", ctrl.Create)
	router.Put("/:id", ctrl.Update)
	router.Delete("/:id", ctrl.Delete)
	router.Get("/", ctrl.GetAll)
}

func (ctrl *CategoryController) Create(c *fiber.Ctx) error {
	var input domain.Category
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	if input.Name == "" {
		return c.Status(400).JSON(fiber.Map{"error": "name is required"})
	}

	if err := ctrl.usecase.Create(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(input)
}

func (ctrl *CategoryController) GetAll(c *fiber.Ctx) error {
	categories, err := ctrl.usecase.GetAll()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(categories)
}

func (ctrl *CategoryController) Update(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid category ID"})
	}

	var input domain.Category
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	input.ID = uint(id)
	if err := ctrl.usecase.Update(&input); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "category updated"})
}

func (ctrl *CategoryController) Delete(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid category ID"})
	}

	if err := ctrl.usecase.Delete(uint(id)); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "category deleted"})
}
