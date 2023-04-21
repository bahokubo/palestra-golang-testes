package user

const (
	ADMIN = "ADMIN"
	DBA   = "DBA"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Type     string `json:"type"`
	Email    string `json:"email"`
}

// Repository interface
type Repository interface {
	Create([]*User) ([]*User, error)
	List() ([]*User, error)
	Update(*User) (*User, error)
	Delete(id string) (int, error)
}

// UseCase interface
type UseCase interface {
	Create([]*User) ([]*User, error)
	List() ([]*User, error)
	Update(*User) (*User, error)
	Delete(id string) error
}
