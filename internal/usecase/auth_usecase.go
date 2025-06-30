package usecase

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/nurhidaylma/gocommerce/internal/domain"
	"github.com/nurhidaylma/gocommerce/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(user *domain.User) error
	Login(email, password string) (string, error)
}

type authUsecase struct {
	repo      repository.AuthRepository
	storeRepo repository.StoreRepository
}

func NewAuthUsecase(r repository.AuthRepository, storeRepo repository.StoreRepository) AuthUsecase {
	return &authUsecase{
		repo:      r,
		storeRepo: storeRepo,
	}
}

func (u *authUsecase) Register(user *domain.User) error {
	existingUser, err := u.repo.FindByEmailOrPhone(user.Email, user.Phone)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashed)
	user.Role = "user"

	err = u.repo.Register(user)
	if err != nil {
		return err
	}

	// Create store automatically
	user.Store = domain.Store{
		UserID: user.ID,
		Name:   user.Name + "'s Store",
	}
	err = u.storeRepo.Create(&user.Store)
	if err != nil {
		return err
	}

	return nil
}

func (u *authUsecase) Login(email, password string) (string, error) {
	user, err := u.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString([]byte("secret"))
}
