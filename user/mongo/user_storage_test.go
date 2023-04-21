package mongo_test

import (
	"context"
	"testing"
	"user-crud/internal/mongo"
	"user-crud/user"
	userRepository "user-crud/user/mongo"

	"github.com/stretchr/testify/assert"
)

type userTest struct {
	userRepository *userRepository.UserStorage
}

func newUserTest(userRepository *userRepository.UserStorage) *userTest {
	return &userTest{
		userRepository: userRepository,
	}
}

// TestUserStorage should test all functions related to the user repository.
// The tests are made by connecting to a testcontainer.
// Should setup and cleanup test data.
func TestUserStorage(t *testing.T) {

	// DB setup
	ctx := context.Background()

	mongoTestContainer, err := mongo.StartMongoContainer(ctx)
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

	test := newUserTest(userRepository)

	// Repo funcs
	t.Run("create users", func(t *testing.T) {
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

		test.cleanup()
	})

	t.Run("update user", func(t *testing.T) {
		createdUser := test.setup()

		createdUser.Email = "updated.email@email.com"

		returnedUser, err := userRepository.Update(&createdUser)

		assert.NotEmpty(t, returnedUser)
		assert.Equal(t, "updated.email@email.com", returnedUser.Email)
		assert.Nil(t, err)

		test.cleanup()
	})

	t.Run("list users", func(t *testing.T) {
		createdUser := test.setup()

		returnedUsers, err := userRepository.List()

		assert.NotEmpty(t, returnedUsers)
		assert.Equal(t, createdUser, *returnedUsers[0])
		assert.Nil(t, err)

		test.cleanup()
	})

	t.Run("delete users", func(t *testing.T) {
		createdUser := test.setup()

		_, err := userRepository.Delete(createdUser.ID)
		assert.Nil(t, err)

		returnedUsers, err := userRepository.List()

		assert.Empty(t, returnedUsers)
		assert.Nil(t, err)
	})

}

func (ut userTest) setup() user.User {
	seed := &user.User{
		Name:     "Teste",
		Username: "teste",
		Password: "1234",
		Type:     "DBA",
		Email:    "teste@teste.com",
	}

	returnedUsers, _ := ut.userRepository.Create([]*user.User{seed})

	return *returnedUsers[0]
}

func (ut userTest) cleanup() {
	users, _ := ut.userRepository.List()

	for _, u := range users {
		ut.userRepository.Delete(u.ID)
	}
}
