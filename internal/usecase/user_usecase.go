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
	repo      repository.UserRepository
	storeRepo repository.StoreRepository
}

func NewUserUsecase(r repository.UserRepository, s repository.StoreRepository) UserUsecase {
	return &userUsecase{
		repo:      r,
		storeRepo: s,
	}
}

func (u *userUsecase) GetProfile(id uint) (*domain.User, error) {
	user, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	store, err := u.storeRepo.GetByUserID(user.ID)
	if err != nil {
		return nil, err
	}

	user.Store = *store

	return user, nil
}

func (u *userUsecase) UpdateProfile(user *domain.User) error {
	updates := map[string]interface{}{}

	if user.Name != "" {
		updates["name"] = user.Name
	}
	if user.Email != "" {
		updates["email"] = user.Email
	}
	if user.Phone != "" {
		updates["phone"] = user.Phone
	}
	if user.Password != "" {
		updates["password"] = user.Password
	}
	if user.Role != "" {
		updates["role"] = user.Role
	}

	if len(updates) == 0 {
		return nil // nothing to update
	}

	return u.repo.Update(user.ID, updates)
}
