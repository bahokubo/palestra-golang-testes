package user

var UserNotFound = "Error user not found"

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(users []*User) ([]*User, error) {
	if _, err := s.repo.Create(users); err != nil {
		return nil, err
	}

	return users, nil
}

// func (s *Service) List() (users []*User, err error) {
// 	if users, err = s.repo.List(); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// func (s *Service) Update(user *User) (*User, error) {
// 	if _, err := s.repo.Update(user); err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }

// func (s *Service) Delete(id string) error {
// 	i, err := s.repo.Delete(id)

// 	if i == 0 {
// 		return errors.New(UserNotFound)
// 	}

// 	return err
// }
