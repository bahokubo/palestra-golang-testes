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
			fmt.Sprintf("[Handler] CreateProducts error couldn't create products: %v", err)
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
}
