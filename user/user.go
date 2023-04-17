package user

type userType string

const (
	ADMIN = "ADMIN"
	DBA   = "DBA"
)

type User struct {
	ID       string   `json:"id"`
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
	Update(user *User) (*User, error)
	Delete(id string) (int, error)
}

// UseCase interface
type UseCase interface {
	Create(u []*User) ([]*User, error)
	List() (users []*User, err error)
	Update(user *User) (*User, error)
	Delete(id string) error
}
