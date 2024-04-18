package user

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"

	"github.com/shihabmridha/golang-app-template/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	cfg     config.Config
	usrRepo Repository
}

func NewService(cfg config.Config, r Repository) Service {
	return Service{
		cfg:     cfg,
		usrRepo: r,
	}
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func newActivationCode(username, secret string) (string, error) {
	token := make([]byte, 4)
	rand.Read([]byte(fmt.Sprintf("%s%s", username, secret)))

	hasher := sha1.New()
	hasher.Write(token)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	return sha, nil
}

func (s *Service) GetAll() ([]User, error) {
	users, err := s.usrRepo.GetAll()

	if err != nil {
		return nil, fmt.Errorf("failed to get all user. error: %w", err)
	}

	for i := range users {
		users[i].Normalize()
	}

	return users, nil
}

func (s *Service) Create(user User) error {
	exists, err := s.usrRepo.AlreadyExists(user.Username, user.Email)
	if err != nil {
		return fmt.Errorf("failed to check if user exists. error: %w", err)
	}

	if exists {
		return fmt.Errorf("user already exists")
	}

	hashedPwd, err := hashPassword(user.Passowrd)
	if err != nil {
		return fmt.Errorf("failed to create password hash. error: %w", err)
	}

	activationCode, err := newActivationCode(user.Username, s.cfg.ActivationCodeSecret)
	if err != nil {
		return fmt.Errorf("failed to create password hash. error: %w", err)
	}

	user.Passowrd = hashedPwd
	user.ActivationCode = activationCode

	_, err = s.usrRepo.Create(user)
	if err != nil {
		return fmt.Errorf("failed to create user. error: %w", err)
	}

	return nil
}

func (s *Service) Activate(code string) error {
	id, err := s.usrRepo.GetIdUsingActivationCode(code)

	if err != nil {
		return fmt.Errorf("failed to fetch id using activation code. error: %w", err)
	}

	if err := s.usrRepo.Activate(id); err != nil {
		return fmt.Errorf("failed to activate user. error: %w", err)
	}

	return nil
}
