package repository

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
}

type authRepo struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepo{db}
}

func (r *authRepo) Register(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *authRepo) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Preload("Store").Where("email = ?", email).First(&user).Error
	return &user, err
}
