package mongo_test

import (
	"context"
	"testing"
	"user-crud/internal/mongo"
	"user-crud/user"
	userRepository "user-crud/user/mongo"

	"github.com/stretchr/testify/assert"
)

// TestUserStorage should test all functions related to the user repository.
// The tests are made by connecting to a testcontainer.
// Should setup and cleanup test data.
func TestUserStorage(t *testing.T) {

	// DB setup
	ctx := context.Background()
	var mongoTestContainer *mongo.TestContainer
	var err error

	mongoTestContainer, err = mongo.StartMongoContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer mongoTestContainer.Terminate(ctx)

	mongoURI := mongoTestContainer.URI

	dbConn, _ := mongo.Open(mongoURI)
	db := dbConn.Database("users")
	assert.Nil(t, err)

	// Repos
	userRepository := userRepository.NewUserStorage(db, ctx)

	// Repo funcs
	t.Run("create user", func(t *testing.T) {
		seed := &user.User{
			Name:     "Teste",
			Username: "teste",
			Password: "1234",
			Type:     "DBA",
			Email:    "teste@teste.com",
		}

		returnedUsers, _ := userRepository.Create([]*user.User{seed})

		expectedCreatedUser := &user.User{
			ID:       returnedUsers[0].ID,
			Name:     "Teste",
			Username: "teste",
			Password: "1234",
			Type:     "DBA",
			Email:    "teste@teste.com",
		}

		assert.NotEmpty(t, returnedUsers)
		assert.Equal(t, expectedCreatedUser, returnedUsers[0])
		assert.Nil(t, err)

		cleanup(userRepository)
	})

}

func seedDB(userRepository *userRepository.UserStorage) user.User {
	seed := &user.User{
		Name:     "Teste",
		Username: "teste",
		Password: "1234",
		Type:     "DBA",
		Email:    "teste@teste.com",
	}

	returnedUsers, _ := userRepository.Create([]*user.User{seed})

	return *returnedUsers[0]
}

func cleanup(userRepository *userRepository.UserStorage) {
	users, _ := userRepository.List()

	for _, u := range users {
		userRepository.Delete(u.ID)
	}
}
