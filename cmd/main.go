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
	db.AutoMigrate(&domain.User{}, &domain.Store{})

	app := fiber.New()

	authRepo := repository.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	controller.NewAuthHandler(app, authUsecase)

	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	controller.NewUserController(app.Group("/api/v1/user", middleware.JWTProtected()), userUsecase)

	app.Use(middleware.JWTProtected())

	log.Fatal(app.Listen(":8080"))
}
