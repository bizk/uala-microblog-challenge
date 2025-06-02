package usecase

import (
	"sort"
)

type GetTimelineUsecase struct {
	tweetRepo TweetRepository
	userRepo  UserRepository
}

func NewGetTimelineUsecase(tr TweetRepository, ur UserRepository) *GetTimelineUsecase {
	return &GetTimelineUsecase{tweetRepo: tr, userRepo: ur}
}

func (uc *GetTimelineUsecase) GetTimeline(userID string) ([]TweetDTO, error) {
	user, err := uc.userRepo.GetByID(userID)
	if err != nil || user == nil {
		return nil, err
	}

	userIDs := user.GetFollowing()
	userIDs = append(userIDs, userID)

	tweets, err := uc.tweetRepo.GetByUserIDs(userIDs)
	if err != nil {
		return nil, err
	}

	// Ordenar por fecha descendente
	// TODO potencial optimizacion aca
	sort.Slice(tweets, func(i, j int) bool {
		return tweets[i].CreatedAt.After(tweets[j].CreatedAt)
	})

	// Separamos la logica de la base de datos de la logica de la api
	result := make([]TweetDTO, 0, len(tweets))
	for _, t := range tweets {
		result = append(result, TweetDTO{
			UserID:  t.UserID,
			Content: t.Content,
			Time:    t.CreatedAt.Format("2025-06-02 15:04:05"),
		})
	}

	return result, nil
}

type TweetDTO struct {
	UserID  string
	Content string
	Time    string
}
