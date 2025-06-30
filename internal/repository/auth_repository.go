package repository

import (
	"errors"

	"github.com/nurhidaylma/gocommerce/internal/domain"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(user *domain.User) error
	FindByEmail(email string) (*domain.User, error)
	FindByEmailOrPhone(email, phone string) (*domain.User, error)
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

func (r *authRepo) FindByEmailOrPhone(email, phone string) (*domain.User, error) {
	var user domain.User

	if email == "" || phone == "" {
		return nil, nil
	}

	err := r.db.Where("email = ? OR phone = ?", email, phone).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // no user found, not an error
		} else {
			return nil, err // actual DB error
		}
	}

	return &user, nil // user found
}
