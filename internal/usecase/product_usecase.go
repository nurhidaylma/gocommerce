package usecase

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/repository"
)

type ProductUsecase interface {
	Create(*domain.Product) error
	GetAll(filter string, categoryID uint, limit, offset int) ([]domain.Product, error)
	GetByID(id uint) (*domain.Product, error)
	Update(*domain.Product, uint) error
	Delete(id uint, userID uint) error
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(r repository.ProductRepository) ProductUsecase {
	return &productUsecase{r}
}

func (u *productUsecase) Create(p *domain.Product) error {
	return u.repo.Create(p)
}

func (u *productUsecase) GetAll(filter string, categoryID uint, limit, offset int) ([]domain.Product, error) {
	return u.repo.GetAll(filter, categoryID, limit, offset)
}

func (u *productUsecase) GetByID(id uint) (*domain.Product, error) {
	return u.repo.GetByID(id)
}

func (u *productUsecase) Update(p *domain.Product, userID uint) error {
	old, err := u.repo.GetByID(p.ID)
	if err != nil || old.UserID != userID {
		return err
	}
	p.UserID = userID
	return u.repo.Update(p)
}

func (u *productUsecase) Delete(id uint, userID uint) error {
	return u.repo.Delete(id, userID)
}
