package usecase

import "uala-microblog-challenge/internal/domain"

type GetUserUsecase struct {
	userRepo UserRepository
}

func NewGetUserUsecase(userRepo UserRepository) *GetUserUsecase {
	return &GetUserUsecase{
		userRepo: userRepo,
	}
}

func (u *GetUserUsecase) Get(userID string) (*domain.User, error) {
	return u.userRepo.GetByID(userID)
}
