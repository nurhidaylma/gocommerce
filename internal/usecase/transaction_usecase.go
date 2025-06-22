package usecase

import (
	"errors"

	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/dto"
	"github.com/nurhidaylma/gocommerce/internal/repository"
)

type TransactionUsecase interface {
	Create(userID uint, input dto.CreateTransactionInput) error
	GetByUser(userID uint) ([]domain.Transaction, error)
	GetByID(id, userID uint) (*domain.Transaction, error)
	CancelTransaction(id, userID uint) error
}

type transactionUsecase struct {
	transactionRepo repository.TransactionRepository
	productRepo     repository.ProductRepository
}

func NewTransactionUsecase(transactionRepo repository.TransactionRepository, productRepo repository.ProductRepository) TransactionUsecase {
	return &transactionUsecase{transactionRepo, productRepo}
}

func (u *transactionUsecase) Create(userID uint, input dto.CreateTransactionInput) error {
	tx := domain.Transaction{
		UserID:    userID,
		AddressID: input.AddressID,
		Status:    domain.Pending,
	}

	var total int
	for _, i := range input.Items {
		product, err := u.productRepo.GetByID(i.ProductID)
		if err != nil || product.Stock < i.Quantity {
			return errors.New("product not found or insufficient stock")
		}

		// Reduce stock
		product.Stock -= i.Quantity
		u.productRepo.Update(product)

		// Add to transaction items
		item := domain.TransactionItem{
			ProductID: i.ProductID,
			Quantity:  i.Quantity,
			Price:     int(product.Price),
			LogProduct: domain.LogProduct{
				Name:        product.Name,
				Description: product.Description,
				Price:       int(product.Price),
				ImageURL:    product.ImageURL,
			},
		}

		tx.Items = append(tx.Items, item)
		total += i.Quantity * int(product.Price)
	}

	tx.Total = total
	return u.transactionRepo.Create(&tx)
}

func (u *transactionUsecase) GetByUser(userID uint) ([]domain.Transaction, error) {
	return u.transactionRepo.GetByUser(userID)
}

func (u *transactionUsecase) GetByID(id, userID uint) (*domain.Transaction, error) {
	return u.transactionRepo.GetByID(id, userID)
}

func (u *transactionUsecase) CancelTransaction(id, userID uint) error {
	tx, err := u.transactionRepo.GetByID(id, userID)
	if err != nil {
		return err
	}

	if tx.Status != domain.Pending {
		return errors.New("only pending transactions can be cancelled")
	}

	// Kembalikan stok produk
	for _, item := range tx.Items {
		product, err := u.productRepo.GetByID(item.ProductID)
		if err != nil {
			return err
		}
		product.Stock += item.Quantity
		if err := u.productRepo.Update(product); err != nil {
			return err
		}
	}

	// Ubah status transaksi
	return u.transactionRepo.UpdateStatus(id, "cancelled")
}
