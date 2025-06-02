package domain

import "testing"

func TestUser_Follow(t *testing.T) {
	user := NewUser("123")
	err := user.Follow("456")
	if err != nil {
		t.Errorf("Error al seguir al usuario: %v", err)
	}

	if len(user.Following) != 1 {
		t.Errorf("El usuario no sigue al usuario 456")
	}
}

func TestUser_GetFollowing(t *testing.T) {
	user := NewUser("123")
	user.Follow("456")
	user.Follow("789")

	following := user.GetFollowing()
	if len(following) != 2 {
		t.Errorf("El usuario no sigue a 2 usuarios")
	}
}
