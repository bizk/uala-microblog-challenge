package usecase

type FollowUserUsecase struct {
	userRepo UserRepository
}

func NewFollowUserUsecase(ur UserRepository) *FollowUserUsecase {
	return &FollowUserUsecase{userRepo: ur}
}

func (uc *FollowUserUsecase) Follow(followerID, followedID string) error {
	follower, err := uc.userRepo.GetByID(followerID)
	if err != nil || follower == nil {
		return err
	}

	followed, err := uc.userRepo.GetByID(followedID)
	if err != nil || followed == nil {
		return err
	}

	if err := follower.Follow(followedID); err != nil {
		return err
	}

	return uc.userRepo.Save(follower)
}
