package usecase

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/repository"
)

type AddressUsecase interface {
	Create(address *domain.Address) error
	GetByUser(userID uint) ([]domain.Address, error)
	Update(address *domain.Address, userID uint) error
	Delete(id, userID uint) error
}

type addressUsecase struct {
	repo repository.AddressRepository
}

func NewAddressUsecase(r repository.AddressRepository) AddressUsecase {
	return &addressUsecase{r}
}

func (u *addressUsecase) Create(a *domain.Address) error {
	return u.repo.Create(a)
}

func (u *addressUsecase) GetByUser(userID uint) ([]domain.Address, error) {
	return u.repo.GetByUser(userID)
}

func (u *addressUsecase) Update(a *domain.Address, userID uint) error {
	existing, err := u.repo.GetByID(a.ID, userID)
	if err != nil || existing.UserID != userID {
		return err
	}

	existing.UserID = userID
	if a.City != "" {
		existing.City = a.City
	}
	if a.District != "" {
		existing.District = a.District
	}
	if a.PostalCode != "" {
		existing.PostalCode = a.PostalCode
	}
	if a.Province != "" {
		existing.Province = a.Province
	}
	if a.Detail != "" {
		existing.Detail = a.Detail
	}

	return u.repo.Update(existing)
}

func (u *addressUsecase) Delete(id, userID uint) error {
	return u.repo.Delete(id, userID)
}
