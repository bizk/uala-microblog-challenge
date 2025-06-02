package usecase

import (
	"uala-microblog-challenge/internal/domain"
)

type CreateUserUsecase struct {
	userRepo UserRepository
}

func NewCreateUserUsecase(userRepo UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{
		userRepo: userRepo,
	}
}

func (u *CreateUserUsecase) Create(id string) (string, error) {
	user := domain.NewUser(id)
	err := u.userRepo.Save(user)
	if err != nil {
		return "", err
	}

	return id, nil
}
