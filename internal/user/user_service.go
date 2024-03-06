package user

import (
	"fmt"
)

type Service struct {
	usrRepo Repository
}

func NewSvc(r Repository) *Service {
	return &Service{usrRepo: r}
}

func (s *Service) Get() (*[]User, error) {
	users, err := s.usrRepo.Get()

	if err != nil {
		return nil, fmt.Errorf("UserService - Get - s.repo.Get: %w", err)
	}

	return users, nil
}

func (s *Service) Create() error {
	_, err := s.usrRepo.Get()

	if err != nil {
		return fmt.Errorf("UserService - Create - s.repo.Create: %w", err)
	}

	return nil
}
