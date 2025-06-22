package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nurhidaylma/gocommerce/config"
	"github.com/nurhidaylma/gocommerce/internal/controller"
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/repository"
	"github.com/nurhidaylma/gocommerce/internal/usecase"
	"github.com/nurhidaylma/gocommerce/middleware"
)

func main() {
	db := config.InitDB()
	db.AutoMigrate(&domain.User{}, &domain.Store{}, &domain.Product{}, &domain.Address{}, &domain.Category{})

	app := fiber.New()

	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	controller.NewAuthController(app, authUsecase)

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	controller.NewUserController(app.Group("/api/v1/user", middleware.JWTProtected()), userUsecase)

	productRepo := repository.NewProductRepository(db)
	productUsecase := usecase.NewProductUsecase(productRepo)
	controller.NewProductController(app.Group("/api/v1/products", middleware.JWTProtected()), productUsecase)

	addressRepo := repository.NewAddressRepository(db)
	addressUC := usecase.NewAddressUsecase(addressRepo)
	controller.NewAddressController(app.Group("/api/v1/address", middleware.JWTProtected()), addressUC)

	categoryRepo := repository.NewCategoryRepository(db)
	categoryUsecase := usecase.NewCategoryUsecase(categoryRepo)
	controller.NewCategoryController(app.Group("/api/v1/category", middleware.JWTProtected(), middleware.AdminOnly()), categoryUsecase)

	app.Use(middleware.JWTProtected())

	log.Fatal(app.Listen(":8080"))
}
