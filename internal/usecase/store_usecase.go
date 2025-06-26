package usecase

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/repository"
)

type StoreUsecase interface {
	GetByUserID(userID uint) (*domain.Store, error)
	Update(userID uint, name, logo string) error
}

type storeUsecase struct {
	repo repository.StoreRepository
}

func NewStoreUsecase(r repository.StoreRepository) StoreUsecase {
	return &storeUsecase{r}
}

func (s *storeUsecase) GetByUserID(userID uint) (*domain.Store, error) {
	return s.repo.GetByUserID(userID)
}

func (s *storeUsecase) Update(userID uint, name, logo string) error {
	store, err := s.repo.GetByUserID(userID)
	if err != nil {
		return err
	}
	store.Name = name
	if logo != "" {
		store.Logo = logo
	}
	return s.repo.Update(store)
}
