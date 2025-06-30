package controller

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/usecase"
	"github.com/nurhidaylma/gocommerce/middleware"
)

type ProductController struct {
	usecase usecase.ProductUsecase
}

func NewProductController(router fiber.Router, u usecase.ProductUsecase) {
	h := &ProductController{u}
	router.Post("/", h.Create)
	router.Get("/", h.GetAll)
	router.Get("/:id", h.GetByID)
	router.Put("/:id", h.Update)
	router.Delete("/:id", h.Delete)
}

func (h *ProductController) Create(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "image required"})
	}

	// pastikan upload folder ada
	if _, err := os.Stat("./uploads"); os.IsNotExist(err) {
		err = os.Mkdir("./uploads", os.ModePerm)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "failed to create upload dir"})
		}
	}

	safeFileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	path := "./uploads/" + safeFileName

	if err := c.SaveFile(file, path); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "upload failed"})
	}

	price, _ := strconv.ParseFloat(c.FormValue("price"), 64)
	stock, _ := strconv.Atoi(c.FormValue("stock"))
	categoryID, _ := strconv.Atoi(c.FormValue("category_id"))

	product := domain.Product{
		Name:        c.FormValue("name"),
		CategoryID:  uint(categoryID),
		Price:       price,
		Stock:       stock,
		ImageURL:    path,
		Description: c.FormValue("description"),
		UserID:      userID,
	}

	if err := h.usecase.Create(&product); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(product)
}

func (h *ProductController) GetAll(c *fiber.Ctx) error {
	filter := c.Query("search")
	categoryID, _ := strconv.Atoi(c.Query("category_id"))
	page, _ := strconv.Atoi(c.Query("page"))
	limit := c.QueryInt("limit", 10)
	offset := (page - 1) * limit

	products, err := h.usecase.GetAll(filter, uint(categoryID), limit, offset)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(products)
}

func (h *ProductController) GetByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	product, err := h.usecase.GetByID(uint(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}
	return c.JSON(product)
}

func (h *ProductController) Update(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))

	var input domain.Product
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

func (h *ProductController) Delete(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint)
	id, _ := strconv.Atoi(c.Params("id"))
	if err := h.usecase.Delete(uint(id), userID); err != nil {
		return c.Status(403).JSON(fiber.Map{"error": "unauthorized or not found"})
	}
	return c.JSON(fiber.Map{"message": "deleted"})
}
