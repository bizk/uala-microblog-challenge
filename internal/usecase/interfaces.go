package usecase

import "uala-microblog-challenge/internal/domain"

// Estos son los repositorios que usaremos para interactuar con la base de datos
// Usamos las interfaces para aislar la capa de dominio de la infraestructura

type TweetRepository interface {
	Save(tweet *domain.Tweet) error
	GetByUserIDs(userIDs []string) ([]*domain.Tweet, error)
}

type UserRepository interface {
	GetByID(userID string) (*domain.User, error)
	Save(user *domain.User) error
}
