package user

type userType string

type User struct {
	Name     string   `json:"name"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Type     userType `json:"type"`
	Email    string   `json:"email"`
}

// Repository interface
type Repository interface {
	Create(users []*User) ([]*User, error)
	List() ([]*User, error)
}

// UseCase interface
type UseCase interface {
	Create(u []*User) ([]*User, error)
	List() (users []*User, err error)
}
