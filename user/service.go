package user

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Create(u []*User) ([]*User, error) {
	if _, err := s.repo.Create(u); err != nil {
		return nil, err
	}

	return u, nil
}
