package user_test

import (
	"testing"
	"user-crud/user"
	"user-crud/user/mock"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
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
		createdUsers, err := s.Create([]*user.User{})
		assert.Nil(t, createdUsers)
		assert.ErrorIs(t, expectedError, err)
	})
}

func TestList(t *testing.T) {
	t.Run("it should return all users and list be called for all of them", func(t *testing.T) {
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
			Return(expectedUsers, nil)

		s := user.NewService(repo)

		users, err := s.List()

		assert.Nil(t, err)
		assert.EqualValues(t, expectedUsers, users)
	})

	t.Run("it should return the expected repository error", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepository(ctrl)
		expectedError := errors.New("error when list in repo")

		repo.EXPECT().
			List().
			Return(nil, expectedError).
			Times(1)

		s := user.NewService(repo)
		listedUsers, err := s.List()
		assert.Nil(t, listedUsers)
		assert.ErrorIs(t, expectedError, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("it should return updated user", func(t *testing.T) {
		expectedUser := &user.User{
			Name:     "Teste",
			Username: "teste",
			Password: "1234",
			Type:     "DBA",
			Email:    "teste@teste.com",
		}

		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepository(ctrl)

		repo.EXPECT().
			Update(expectedUser).
			Return(expectedUser, nil)

		s := user.NewService(repo)

		user, err := s.Update(expectedUser)

		assert.Nil(t, err)
		assert.EqualValues(t, expectedUser, user)
	})

	t.Run("it should return the expected repository error for update", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepository(ctrl)
		expectedError := errors.New("error when update in repo")

		userToUpdate := &user.User{
			Name:     "Teste",
			Username: "teste",
			Password: "1234",
			Type:     "",
			Email:    "teste@teste.com",
		}

		repo.EXPECT().
			Update(userToUpdate).
			Return(nil, expectedError).
			Times(1)

		s := user.NewService(repo)
		userUpdated, err := s.Update(userToUpdate)
		assert.Nil(t, userUpdated)
		assert.ErrorIs(t, expectedError, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("it should return a message because user was deleted", func(t *testing.T) {
		userId := "12345"

		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepository(ctrl)
		repo.EXPECT().
			Delete(userId).
			Return(1, nil)

		s := user.NewService(repo)
		err := s.Delete(userId)

		assert.Nil(t, err)
	})

	t.Run("it should return the expected repository error for delete", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepository(ctrl)
		expectedError := errors.New("error when delete in repo")
		userId := "wrong string"
		repo.EXPECT().
			Delete(userId).
			Return(1, expectedError).
			Times(1)

		s := user.NewService(repo)
		err := s.Delete(userId)
		assert.ErrorIs(t, expectedError, err)
	})

	t.Run("it should return a user not found error when try to delete", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		repo := mock.NewMockRepository(ctrl)
		userId := "wrong string"
		expectedError := errors.New("error when delete in repo")

		repo.EXPECT().
			Delete(userId).
			Return(0, expectedError).
			Times(1)

		s := user.NewService(repo)
		err := s.Delete(userId)
		assert.Errorf(t, err, user.UserNotFound)
	})
}
