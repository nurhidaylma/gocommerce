package repository

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id uint) (*domain.User, error)
	Update(id uint, updates map[string]interface{}) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db}
}

func (r *userRepo) FindByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepo) Update(id uint, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}

	return r.db.Model(&domain.User{}).
		Where("id = ?", id).
		Updates(updates).Error
}
