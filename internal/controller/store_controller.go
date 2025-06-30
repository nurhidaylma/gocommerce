package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nurhidaylma/gocommerce/internal/dto"
	"github.com/nurhidaylma/gocommerce/internal/usecase"
)

type StoreController struct {
	usecase usecase.StoreUsecase
}

func NewStoreController(r fiber.Router, s usecase.StoreUsecase) {
	h := &StoreController{s}
	r.Get("/", h.Get)
	r.Put("/", h.Update)
}

func (h *StoreController) Get(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	store, err := h.usecase.GetByUserID(userID)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "store not found"})
	}
	return c.JSON(store)
}

func (h *StoreController) Update(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	name := c.FormValue("name")

	var logoPath string
	file, err := c.FormFile("logo")
	if file == nil {
		// logo optional, nothing to do
	} else if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	} else {
		logoPath = fmt.Sprintf("uploads/%d_%s", userID, file.Filename)
		if err := c.SaveFile(file, logoPath); err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to save file"})
		}
	}

	request := dto.UpdateStoreRequest{
		UserID: userID,
		Name:   name,
		Logo:   logoPath,
	}
	if err := h.usecase.Update(&request); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"message": "store updated"})
}
