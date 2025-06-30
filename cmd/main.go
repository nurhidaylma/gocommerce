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
	db.AutoMigrate(&domain.User{}, &domain.Store{}, &domain.Product{}, &domain.Address{}, &domain.Category{},
		&domain.Transaction{}, &domain.TransactionItem{}, &domain.LogProduct{}, &domain.Store{},
	)

	app := fiber.New()

	storeRepo := repository.NewStoreRepository(db)
	storeUsecase := usecase.NewStoreUsecase(storeRepo)
	controller.NewStoreController(app.Group("/api/v1/store", middleware.JWTProtected()), storeUsecase)

	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo, storeRepo)
	controller.NewAuthController(app, authUsecase)

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo, storeRepo)
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

	transactionRepo := repository.NewTransactionRepository(db)
	transactionUsecase := usecase.NewTransactionUsecase(transactionRepo, productRepo)
	controller.NewTransactionController(app.Group("/api/v1/transactions", middleware.JWTProtected()), transactionUsecase)

	app.Use(middleware.JWTProtected())

	log.Fatal(app.Listen(":8080"))
}
