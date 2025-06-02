package usecase

import (
	"errors"
	"testing"
	"uala-microblog-challenge/internal/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// --- Mocks ---
// No soy muy fan de usar mocks honestamente: Suelen quedar deprecados facilmente y consumir mucho tiempo para mantenerlos vs lo que aportan (IMO)
// Arme unos pocos tests para ser demostrativo.

type mockUser struct {
	mock.Mock
	id string
}

func (m *mockUser) Follow(followedID string) error {
	args := m.Called(followedID)
	return args.Error(0)
}

type mockUserRepo struct {
	mock.Mock
}

func (m *mockUserRepo) GetByID(id string) (*domain.User, error) {
	args := m.Called(id)
	user, _ := args.Get(0).(*domain.User)
	return user, args.Error(1)
}

func (m *mockUserRepo) Save(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

// --- Tests ---

func TestFollowUserUsecase_Follow_Success(t *testing.T) {
	repo := new(mockUserRepo)
	follower := new(mockUser)
	followed := new(mockUser)

	repo.On("GetByID", "followerID").Return(follower, nil)
	repo.On("GetByID", "followedID").Return(followed, nil)
	follower.On("Follow", "followedID").Return(nil)
	repo.On("Save", follower).Return(nil)

	uc := NewFollowUserUsecase(repo)
	err := uc.Follow("followerID", "followedID")
	assert.NoError(t, err)
}

func TestFollowUserUsecase_Follow_FollowerNotFound(t *testing.T) {
	repo := new(mockUserRepo)
	repo.On("GetByID", "followerID").Return(nil, errors.New("not found"))

	uc := NewFollowUserUsecase(repo)
	err := uc.Follow("followerID", "followedID")
	assert.Error(t, err)
}
