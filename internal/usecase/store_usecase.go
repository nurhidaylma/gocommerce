package usecase

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/dto"
	"github.com/nurhidaylma/gocommerce/internal/repository"
)

type StoreUsecase interface {
	GetByUserID(userID uint) (*domain.Store, error)
	Update(store *dto.UpdateStoreRequest) error
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

func (s *storeUsecase) Update(request *dto.UpdateStoreRequest) error {
	store, err := s.repo.GetByUserID(request.UserID)
	if err != nil {
		return err
	}

	if request.Name != "" {
		store.Name = request.Name
	}
	if request.Logo != "" {
		store.Logo = request.Logo
	}
	return s.repo.Update(store)
}
