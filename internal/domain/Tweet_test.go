package domain

import (
	"testing"
)

func TestNewTweet(t *testing.T) {
	tweet, err := NewTweet("123", "Lo viejo funciona, Juan")
	if err != nil {
		t.Errorf("Error al crear el tweet: %v", err)
	}

	if tweet.ID == "" {
		t.Errorf("El ID del tweet no puede estar vacío")
	}

	if tweet.UserID != "123" {
		t.Errorf("El ID del usuario no puede estar vacío")
	}

	if tweet.Content != "Lo viejo funciona, Juan" {
		t.Errorf("El contenido del tweet no puede estar vacío")
	}
}
