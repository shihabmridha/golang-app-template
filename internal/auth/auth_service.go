package auth

import (
	"fmt"

	"github.com/shihabmridha/golang-app-template/internal/user"
	"github.com/shihabmridha/golang-app-template/pkg/config"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	appCfg  *config.App
	usrRepo *user.Repository
}

func NewService(cfg *config.App, r *user.Repository) *Service {
	return &Service{
		appCfg:  cfg,
		usrRepo: r,
	}
}

func isValidPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (s *Service) GetToken(login UserLogin) (*AuthToken, error) {
	user, _ := s.usrRepo.GetByEmail(login.Email)
	if user == nil {
		return nil, fmt.Errorf("invalid usename or password")
	}

	isValidPwd := isValidPassword(login.Passowrd, user.Passowrd)
	if !isValidPwd {
		return nil, fmt.Errorf("invalid usename or password")
	}

	token := &AuthToken{
		AccessToken: "a nice jwt",
	}

	return token, nil
}
