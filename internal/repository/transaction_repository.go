package repository

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(tx *domain.Transaction) error
	GetByUser(userID uint) ([]domain.Transaction, error)
	GetByID(id, userID uint) (*domain.Transaction, error)
	UpdateStatus(id uint, status string) error
}

type transactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepo{db}
}

func (r *transactionRepo) Create(tx *domain.Transaction) error {
	return r.db.Create(tx).Error
}

func (r *transactionRepo) GetByUser(userID uint) ([]domain.Transaction, error) {
	var txs []domain.Transaction
	err := r.db.Preload("Items.LogProduct").Where("user_id = ?", userID).Find(&txs).Error
	return txs, err
}

func (r *transactionRepo) GetByID(id, userID uint) (*domain.Transaction, error) {
	var tx domain.Transaction
	err := r.db.Preload("Items.LogProduct").Where("id = ? AND user_id = ?", id, userID).First(&tx).Error
	return &tx, err
}

func (r *transactionRepo) UpdateStatus(id uint, status string) error {
	return r.db.Model(&domain.Transaction{}).Where("id = ?", id).Update("status", status).Error
}
