package memory

import (
	"sync"
	"uala-microblog-challenge/internal/domain"
)

type InMemoryTweetRepository struct {
	mu     sync.RWMutex
	tweets []*domain.Tweet
}

func NewInMemoryTweetRepository() *InMemoryTweetRepository {
	return &InMemoryTweetRepository{
		tweets: make([]*domain.Tweet, 0),
	}
}

func (r *InMemoryTweetRepository) Save(tweet *domain.Tweet) error {
	r.mu.Lock()         // Basicamente lo que estamos haciendo es que cuando se llama a Save, se bloquea el acceso a la memoria para que no se pueda acceder a la misma desde otro lado
	defer r.mu.Unlock() // Defer se ejecuta cuando la funcion termina

	r.tweets = append(r.tweets, tweet)
	return nil
}

func (r *InMemoryTweetRepository) GetByUserIDs(userIDs []string) ([]*domain.Tweet, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Index rápido para búsqueda
	index := make(map[string]bool)
	for _, id := range userIDs {
		index[id] = true
	}

	var result []*domain.Tweet
	for _, t := range r.tweets {
		if index[t.UserID] {
			result = append(result, t)
		}
	}
	return result, nil
}
