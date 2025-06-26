package repository

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"gorm.io/gorm"
)

type StoreRepository interface {
	Create(store *domain.Store) error
	GetByUserID(userID uint) (*domain.Store, error)
	Update(store *domain.Store) error
}

type storeRepo struct {
	db *gorm.DB
}

func NewStoreRepository(db *gorm.DB) StoreRepository {
	return &storeRepo{db}
}

func (r *storeRepo) Create(store *domain.Store) error {
	return r.db.Create(store).Error
}

func (r *storeRepo) GetByUserID(userID uint) (*domain.Store, error) {
	var store domain.Store
	err := r.db.Where("user_id = ?", userID).First(&store).Error
	return &store, err
}

func (r *storeRepo) Update(store *domain.Store) error {
	return r.db.Save(store).Error
}
