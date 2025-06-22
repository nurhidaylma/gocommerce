package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nurhidaylma/gocommerce/internal/dto"
	"github.com/nurhidaylma/gocommerce/internal/usecase"
)

type TransactionController struct {
	usecase usecase.TransactionUsecase
}

func NewTransactionController(router fiber.Router, usecase usecase.TransactionUsecase) {
	ctrl := &TransactionController{usecase}
	router.Post("/", ctrl.Create)
	router.Get("/", ctrl.GetAll)
	router.Put("/:id/cancel", ctrl.Cancel)
}

func (ctrl *TransactionController) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	var input dto.CreateTransactionInput

	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}

	if err := ctrl.usecase.Create(userID, input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "transaction created"})
}

func (ctrl *TransactionController) GetAll(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	txs, err := ctrl.usecase.GetByUser(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(txs)
}

func (ctrl *TransactionController) GetByID(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
	}

	tx, err := ctrl.usecase.GetByID(uint(id), userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	return c.JSON(tx)
}

func (ctrl *TransactionController) Cancel(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := ctrl.usecase.CancelTransaction(uint(id), userID); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "transaction cancelled"})
}
