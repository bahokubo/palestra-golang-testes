package user

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

func (s *Service) List() (users []*User, err error) {
	if _, err = s.repo.List(); err != nil {
		return nil, err
	}

	return users, nil
}
