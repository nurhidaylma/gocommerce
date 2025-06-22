package repository

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *domain.Category) error
	Update(category *domain.Category) error
	Delete(id uint) error
	GetAll() ([]domain.Category, error)
}

type categoryRepo struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepo{db}
}

func (r *categoryRepo) Create(c *domain.Category) error {
	return r.db.Create(c).Error
}

func (r *categoryRepo) Update(c *domain.Category) error {
	return r.db.Save(c).Error
}

func (r *categoryRepo) Delete(id uint) error {
	return r.db.Delete(&domain.Category{}, id).Error
}

func (r *categoryRepo) GetAll() ([]domain.Category, error) {
	var cats []domain.Category
	return cats, r.db.Find(&cats).Error
}
