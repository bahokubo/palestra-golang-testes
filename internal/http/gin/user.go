package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-crud/internal/http/presenter"
	"user-crud/user"
)

func createUsers(s user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		var users []*user.User

		if err := c.BindJSON(&users); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("error handling with body: %v", err))
			return
		}

		users, err := s.Create(users)
		if err != nil && len(users) == 0 {
			fmt.Sprintf("[Handler] createUsers error couldn't create users: %v", err)
			c.String(http.StatusInternalServerError, fmt.Sprintf("couldn't create users: %v", err))
			return
		}

		createUserPresenter := presenter.CreateUserResponse{}

		resp := *createUserPresenter.Parse(users, err)

		c.JSON(http.StatusOK, resp)
	}
}

func ListUsers(s user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		var u []user.User

		if err := c.BindJSON(&u); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("error handling with body: %v", err))
			return
		}

		users, err := s.List()
		if err != nil && len(users) == 0 {
			fmt.Sprintf("[Handler] ListUsers error couldn't list users: %v", err)
			c.String(http.StatusInternalServerError, fmt.Sprintf("couldn't create users: %v", err))
			return
		}

		createUserPresenter := presenter.CreateUserResponse{}

		resp := *createUserPresenter.Parse(users, err)

		c.JSON(http.StatusOK, resp)
	}
}

func MakeProductHandler(r *gin.RouterGroup, s user.UseCase) {
	r.Handle(http.MethodPost, "users", createUsers(s))
	r.Handle(http.MethodGet, "users", ListUsers(s))
}
