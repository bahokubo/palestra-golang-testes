package user_test

import (
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"user-crud/user"
	"user-crud/user/mock"
)

func TestCreate(t *testing.T) {
	t.Run("it should return all users and create be called for all of them", func(t *testing.T) {
		expectedUsers := []*user.User{
			{
				Name:     "Nichene",
				Email:    "ni@gmail.com",
				Type:     "ADMIN",
				Username: "ni",
			},
			{
				Name:     "Barbara",
				Email:    "ba@gmail.com",
				Username: "bcasac",
				Type:     "USER",
			},
		}

		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepository(ctrl)

		repo.EXPECT().
			Create(expectedUsers).
			Return(expectedUsers, nil)

		s := user.NewService(repo)

		users, err := s.Create(expectedUsers)

		assert.Nil(t, err)
		assert.EqualValues(t, expectedUsers, users)
	})

	t.Run("it should return the expected repository error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepository(ctrl)
		expectedError := errors.New("error when creating in repo")

		repo.EXPECT().
			Create(gomock.Any()).
			Return(nil, expectedError).
			Times(1)

		s := user.NewService(repo)
		validProducts, err := s.Create([]*user.User{})
		assert.Nil(t, validProducts)
		assert.ErrorIs(t, expectedError, err)
	})
}

func TestList(t *testing.T) {
	t.Run("it should return the expected slice of Users", func(t *testing.T) {
		expectedUsers := []*user.User{
			{
				Name:     "Nichene",
				Email:    "ni@gmail.com",
				Type:     "ADMIN",
				Username: "ni",
			},
			{
				Name:     "Barbara",
				Email:    "ba@gmail.com",
				Username: "bcasac",
				Type:     "USER",
			},
		}

		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepository(ctrl)

		repo.EXPECT().
			List().
			Return(expectedUsers, nil).
			Times(1)

		s := user.NewService(repo)
		users, err := s.List()

		assert.Nil(t, err)
		assert.EqualValues(t, expectedUsers[0].Email, users[0].Email)
	})

	t.Run("it should return expected error with nil slices", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepository(ctrl)
		expectedError := errors.New("repo error")

		repo.EXPECT().
			List().
			Return(nil, expectedError).
			Times(1)

		s := user.NewService(repo)
		users, err := s.List()

		assert.Nil(t, users)
		assert.ErrorIs(t, expectedError, err)
	})
}
