package domain

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/google/uuid"
)

const MaxTweetLength = 280

type Tweet struct {
	ID        string
	UserID    string
	Content   string
	CreatedAt time.Time
}

func NewTweet(userID string, content string) (*Tweet, error) {
	if utf8.RuneCountInString(content) > MaxTweetLength { // Valida que el tweet no exceda el largo máximo
		return nil, errors.New("tweet content exceeds max length")
	}

	if userID == "" { // Valida que el usuario no sea vacío
		return nil, errors.New("user ID is required")
	}

	// TODO Podriamos validar que el userId exista en efecto en la base de datos

	return &Tweet{
		ID:        uuid.New().String(),
		UserID:    userID,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
