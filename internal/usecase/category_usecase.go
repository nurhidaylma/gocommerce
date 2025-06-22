package usecase

import (
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/repository"
)

type CategoryUsecase interface {
	Create(c *domain.Category) error
	Update(c *domain.Category) error
	Delete(id uint) error
	GetAll() ([]domain.Category, error)
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repo}
}

func (u *categoryUsecase) Create(c *domain.Category) error {
	return u.repo.Create(c)
}

func (u *categoryUsecase) Update(c *domain.Category) error {
	return u.repo.Update(c)
}

func (u *categoryUsecase) Delete(id uint) error {
	return u.repo.Delete(id)
}

func (u *categoryUsecase) GetAll() ([]domain.Category, error) {
	return u.repo.GetAll()
}
