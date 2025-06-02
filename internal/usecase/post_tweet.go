package usecase

import "uala-microblog-challenge/internal/domain"

type PostTweetUsecase struct {
	tweetRepo TweetRepository
	userRepo  UserRepository
}

func NewPostTweetUsecase(tr TweetRepository, ur UserRepository) *PostTweetUsecase {
	return &PostTweetUsecase{tweetRepo: tr, userRepo: ur}
}

func (uc *PostTweetUsecase) PostTweet(userID, content string) error {
	user, err := uc.userRepo.GetByID(userID) // Validamos que el usuario exista
	if err != nil || user == nil {
		return err
	}

	tweet, err := domain.NewTweet(userID, content)
	if err != nil {
		return err
	}

	return uc.tweetRepo.Save(tweet)
}
