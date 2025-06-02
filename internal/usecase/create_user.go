package usecase

import (
	"uala-microblog-challenge/internal/domain"

	"github.com/google/uuid"
)

type CreateUserUsecase struct {
	userRepo UserRepository
}

func NewCreateUserUsecase(userRepo UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{
		userRepo: userRepo,
	}
}

func (u *CreateUserUsecase) Create() (string, error) {
	userID := uuid.New()
	user := domain.NewUser(userID.String())
	err := u.userRepo.Save(user)
	if err != nil {
		return "", err
	}

	return userID.String(), nil
}
