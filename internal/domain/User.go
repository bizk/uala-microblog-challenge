package domain

import "errors"

type User struct {
	ID        string
	Following map[string]bool // Usamos el bool por que go no te permite usar un set y el map es rapido para buscar
}

func NewUser(id string) *User {
	return &User{
		ID:        id,
		Following: make(map[string]bool),
	}
}

func (u *User) Follow(userID string) error {
	if userID == "" {
		return errors.New("userID cannot be empty")
	}
	if userID == u.ID {
		return errors.New("cannot follow yourself")
	}
	u.Following[userID] = true
	return nil
}

func (u *User) GetFollowing() []string {
	ids := make([]string, 0, len(u.Following))
	for id := range u.Following {
		ids = append(ids, id)
	}
	return ids
}
