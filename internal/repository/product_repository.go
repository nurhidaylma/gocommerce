package repository

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *domain.Product) error
	GetAll(filter string, categoryID uint, limit, offset int) ([]domain.Product, error)
	GetByID(id uint) (*domain.Product, error)
	Update(product *domain.Product) error
	Delete(id uint, userID uint) error
}

type productRepo struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepo{db}
}
func (r *productRepo) Create(p *domain.Product) error {
	return r.db.Create(p).Error
}

func (r *productRepo) GetAll(filter string, categoryID uint, limit, offset int) ([]domain.Product, error) {
	var products []domain.Product
	query := r.db.Limit(limit).Offset(offset)
	if filter != "" {
		query = query.Where("name LIKE ?", "%"+filter+"%")
	}
	if categoryID != 0 {
		query = query.Where("category_id = ?", categoryID)
	}
	err := query.Find(&products).Error
	return products, err
}

func (r *productRepo) GetByID(id uint) (*domain.Product, error) {
	var product domain.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *productRepo) Update(p *domain.Product) error {
	return r.db.Save(p).Error
}

func (r *productRepo) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&domain.Product{}).Error
}
