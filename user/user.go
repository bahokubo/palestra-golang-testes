package user

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"usrname"`
	Password string `json:"password"`
	Type     string `json:"type"`
}

// Repository interface
type Repository interface {
	Create(users []*User) ([]*User, error)
}

// UseCase interface
type UseCase interface {
	Create(u []*User) ([]*User, error)
}
