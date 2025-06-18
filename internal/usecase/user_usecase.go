package usecase

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/repository"
)

type UserUsecase interface {
	GetProfile(id uint) (*domain.User, error)
	UpdateProfile(user *domain.User) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{r}
}

func (u *userUsecase) GetProfile(id uint) (*domain.User, error) {
	return u.repo.FindByID(id)
}

func (u *userUsecase) UpdateProfile(user *domain.User) error {
	return u.repo.Update(user)
}
