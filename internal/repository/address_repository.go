package repository

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"gorm.io/gorm"
)

type AddressRepository interface {
	Create(address *domain.Address) error
	GetByUser(userID uint) ([]domain.Address, error)
	GetByID(id, userID uint) (*domain.Address, error)
	Update(address *domain.Address) error
	Delete(id, userID uint) error
}

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepo{db}
}

func (r *addressRepo) Create(a *domain.Address) error {
	return r.db.Create(a).Error
}

func (r *addressRepo) GetByUser(userID uint) ([]domain.Address, error) {
	var addresses []domain.Address
	err := r.db.Where("user_id = ?", userID).Find(&addresses).Error
	return addresses, err
}

func (r *addressRepo) GetByID(id, userID uint) (*domain.Address, error) {
	var addr domain.Address
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&addr).Error
	return &addr, err
}

func (r *addressRepo) Update(a *domain.Address) error {
	return r.db.Save(a).Error
}

func (r *addressRepo) Delete(id, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&domain.Address{}).Error
}
