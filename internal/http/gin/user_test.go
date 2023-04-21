package gin_test

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"gotest.tools/assert"
// )

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	handler "user-crud/internal/http/gin"
	"user-crud/internal/http/presenter"
	"user-crud/internal/mongo"
	"user-crud/user"
	userRepository "user-crud/user/mongo"

	"github.com/gin-gonic/gin"
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
func TestUserE2E(t *testing.T) {
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
	assert.Nil(t, err)
	userService := user.NewService(userRepository)

	//Route configuration
	r := gin.Default()
	v1 := r.Group("/api/v1")
	handler.MakeUserHandler(v1, userService)

	t.Run("Should return 200 when users are successfully listed", func(t *testing.T) {
		createdUser := test.setup()

		res, err := http.NewRequest("GET", "/api/v1/users", strings.NewReader(``))
		assert.Nil(t, err)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, res)

		body, err := ioutil.ReadAll(w.Body)

		expectedResponse := presenter.CreateUserResponse{
			Users: []*presenter.UserPresenter{
				{
					ID:       createdUser.ID,
					Name:     createdUser.Name,
					Username: createdUser.Username,
					Password: createdUser.Password,
					Type:     createdUser.Type,
					Email:    createdUser.Email,
				},
			},
			ErrorMessage: "",
		}

		v, _ := json.Marshal(expectedResponse)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(v), string(body))
		assert.Nil(t, err)

		test.cleanup()
	})

	t.Run("Should return 200 when user is successfully created", func(t *testing.T) {
		usersToCreate := []user.User{{
			Name:     "Teste",
			Username: "teste",
			Password: "1234",
			Type:     "DBA",
			Email:    "teste@teste.com",
		}}

		usersByte, _ := json.Marshal(usersToCreate)
		req, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(usersByte))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		b, _ := io.ReadAll(w.Body)
		var createResponse *presenter.CreateUserResponse
		json.Unmarshal(b, &createResponse)

		expectedResponse := presenter.CreateUserResponse{
			Users: []*presenter.UserPresenter{
				{
					ID:       createResponse.Users[0].ID,
					Name:     usersToCreate[0].Name,
					Username: usersToCreate[0].Username,
					Password: usersToCreate[0].Password,
					Type:     usersToCreate[0].Type,
					Email:    usersToCreate[0].Email,
				},
			},
			ErrorMessage: "",
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, expectedResponse, *createResponse)
		assert.Nil(t, err)
	})

	t.Run("Should return 400 when user is not successfully created", func(t *testing.T) {
		usersToCreate := "wrong string"

		usersByte, _ := json.Marshal(usersToCreate)
		req, err := http.NewRequest("POST", "/api/v1/users", bytes.NewBuffer(usersByte))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Nil(t, err)

	})

	t.Run("Should return 200 when user is successfully updated", func(t *testing.T) {
		userToUpdate := test.setup()
		userToUpdate.Password = "a much better password because this one is long"

		usersByte, _ := json.Marshal(userToUpdate)
		req, err := http.NewRequest("PUT", "/api/v1/users", bytes.NewBuffer(usersByte))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		b, _ := io.ReadAll(w.Body)
		v, _ := json.Marshal(userToUpdate)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, string(v), string(b))
		assert.Nil(t, err)

		test.cleanup()
	})

	t.Run("Should return 400 when user is not successfully updated", func(t *testing.T) {
		usersToUpdate := "wrong string"

		usersByte, _ := json.Marshal(usersToUpdate)
		req, err := http.NewRequest("PUT", "/api/v1/users", bytes.NewBuffer(usersByte))
		assert.Nil(t, err)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Nil(t, err)

	})

	t.Run("Should return 200 when user is successfully deleted", func(t *testing.T) {
		createdUser := test.setup()
		endpoint := fmt.Sprintf("/api/v1/user/%s", createdUser.ID)

		res, err := http.NewRequest("DELETE", endpoint, strings.NewReader(``))
		assert.Nil(t, err)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, res)

		body, err := ioutil.ReadAll(w.Body)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "User deleted", string(body))
		assert.Nil(t, err)

		test.cleanup()
	})

	t.Run("Should return 404 when user is not successfully deleted", func(t *testing.T) {
		res, err := http.NewRequest("DELETE", "/api/v1/user", strings.NewReader(``))
		assert.Nil(t, err)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, res)

		body, err := ioutil.ReadAll(w.Body)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, "404 page not found", string(body))
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
