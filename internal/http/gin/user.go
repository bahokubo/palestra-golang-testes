package gin

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"user-crud/internal/http/presenter"
	"user-crud/user"

	"github.com/gin-gonic/gin"
)

func createUsers(s user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("[Handler] Starting create users")
		var users []*user.User

		if err := c.BindJSON(&users); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("error handling with body: %v", err))
			return
		}

		for i, u := range users {
			if strings.ToUpper(u.Type) != user.ADMIN && strings.ToUpper(u.Type) != user.DBA {
				c.String(http.StatusBadRequest, fmt.Sprintf("this user type is not valid: %v", i))
				return
			}
		}

		users, err := s.Create(users)
		if err != nil && len(users) == 0 {
			log.Println(fmt.Sprintf("[Handler] createUsers error couldn't create users: %v", err))
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

		var u []*user.User

		if err := c.BindJSON(&u); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("error handling with body: %v", err))
			return
		}

		users, err := s.List()
		if err != nil && len(users) == 0 {
			log.Println(fmt.Sprintf("[Handler] ListUsers error couldn't list users: %v", err))
			c.String(http.StatusInternalServerError, fmt.Sprintf("couldn't create users: %v", err))
			return
		}

		createUserPresenter := presenter.CreateUserResponse{}

		resp := *createUserPresenter.Parse(users, err)

		c.JSON(http.StatusOK, resp)
	}
}

func updateUsers(s user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {

		var u *user.User

		log.Println("[Handler] updateUsers initialize")

		if err := c.BindJSON(&u); err != nil {
			fmt.Sprintf("[Handler] UpdateUser handling with body error: %v", err)
			c.String(http.StatusBadRequest, fmt.Sprintf("error handling with body: %v", err))
			return
		}

		updatedUser, err := s.Update(u)

		if err != nil {
			fmt.Sprintf("[Handler] UpdateUser error: %v", err)
			c.String(http.StatusNotFound, "User not found")
			return
		}

		if err != nil {
			fmt.Sprintf("[Handler] UpdateUser error: %v", err)
			c.String(http.StatusInternalServerError, fmt.Sprintf("couldn't update user: %v", err))
			return
		}

		log.Println("[Handler] UpdateUser succeeded")

		c.JSON(http.StatusOK, updatedUser)
	}
}

func deleteUser(s user.UseCase) gin.HandlerFunc {
	return func(c *gin.Context) {
		uId := c.Param("id")

		fmt.Sprintf("[Handler] deleteUser for product id %s", uId)

		err := s.Delete(uId)

		if err != nil && err.Error() == string(user.UserNotFound) {
			fmt.Sprintf("[Handler] deleteUser error: %v", err)
			c.String(http.StatusNotFound, string(user.UserNotFound))
			return
		}

		if err != nil {
			fmt.Sprintf("[Handler] deleteUser error: %v", err)
			c.String(http.StatusInternalServerError, fmt.Sprintf("couldn't delete user: %v", err))
			return
		}

		fmt.Sprintf("[Handler] deleteUser succeeded for product id %s", uId)

		c.String(http.StatusOK, "User deleted")
	}
}

func MakeUserHandler(r *gin.RouterGroup, s user.UseCase) {
	r.Handle(http.MethodPost, "users", createUsers(s))
	r.Handle(http.MethodGet, "users", ListUsers(s))
	r.Handle(http.MethodPut, "users", updateUsers(s))
	r.Handle(http.MethodDelete, "user/:id", deleteUser(s))
}
