package presenter

import (
	"github.com/akrennmair/slice"
	"user-crud/user"
)

type UserPresenter struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Email    string `json:"email"`
}

type CreateUserResponse struct {
	Users        []*UserPresenter `json:"user"`
	ErrorMessage string           `json:"errorMessage"`
}

func (cR *CreateUserResponse) Parse(users []*user.User, err error) *CreateUserResponse {
	var errorMessage string

	if err != nil {
		errorMessage = err.Error()
	}

	return &CreateUserResponse{
		Users: slice.Map(users, func(u *user.User) *UserPresenter {
			return &UserPresenter{
				Name:     u.Name,
				Username: u.Username,
				Password: u.Password,
				Type:     string(u.Type),
				Email:    u.Email,
			}
		}),
		ErrorMessage: errorMessage,
	}
}
