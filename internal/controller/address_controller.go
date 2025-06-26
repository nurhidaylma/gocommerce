package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/usecase"
	"github.com/nurhidaylma/gocommerce/middleware"
)

type AddressController struct {
	usecase usecase.AddressUsecase
}

func NewAddressController(router fiber.Router, u usecase.AddressUsecase) {
	h := &AddressController{u}
	router.Post("/", h.Create)
	router.Get("/", h.GetAll)
	router.Put("/:id", h.Update)
	router.Delete("/:id", h.Delete)
}

func (h *AddressController) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	var addr domain.Address
	if err := c.BodyParser(&addr); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	addr.UserID = userID

	if err := h.usecase.Create(&addr); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(addr)
}

func (h *AddressController) GetAll(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	addresses, err := h.usecase.GetByUser(userID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(addresses)
}

func (h *AddressController) Update(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	var input domain.Address
	if err := c.BodyParser(&input); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid input"})
	}
	input.ID = uint(id)

	if ok := middleware.IsAuthorized(input.UserID, userID); !ok {
		return c.Status(403).JSON(fiber.Map{"error": "unauthorized"})
	}

	if err := h.usecase.Update(&input, userID); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": "unauthorized"})
	}
	return c.JSON(fiber.Map{"message": "updated"})
}

func (h *AddressController) Delete(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	if err := h.usecase.Delete(uint(id), userID); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": "unauthorized"})
	}
	return c.JSON(fiber.Map{"message": "deleted"})
}
