package auth

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/shihabmridha/golang-app-template/internal/user"
	"github.com/shihabmridha/golang-app-template/pkg/config"
	"github.com/shihabmridha/golang-app-template/pkg/render"
	"golang.org/x/crypto/bcrypt"
)

const UserId TokenUserId = iota

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

func (s *Service) Login(login UserLogin) (*AuthToken, error) {
	user, _ := s.usrRepo.GetByEmail(login.Email)
	if user == nil {
		return nil, fmt.Errorf("invalid usename or password")
	}

	isValidPwd := isValidPassword(login.Passowrd, user.Passowrd)
	if !isValidPwd {
		return nil, fmt.Errorf("invalid usename or password")
	}

	if !user.IsActive {
		return nil, fmt.Errorf("inactive user")
	}

	token, _ := createToken(s.appCfg.JwtSecret(), user.Id)

	return token, nil
}

// Middleware to verify JWT token and set userId to context
func (s *Service) Verify(render *render.Renderer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			jwtToken, err := getTokenFromHeader(r)

			if err != nil {
				render.RenderJSON(w, http.StatusUnauthorized, nil)
				return
			}

			userId, err := parseJwt(s.appCfg.JwtSecret(), jwtToken)

			if err != nil {
				render.RenderJSON(w, http.StatusUnauthorized, nil)
				return
			}

			ctx := context.WithValue(r.Context(), UserId, userId)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func isValidPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func newJwt(secret []byte, id int64) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", id),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
	}).SignedString(secret)
}

func createToken(secret []byte, userId int64) (*AuthToken, error) {
	jwt, err := newJwt(secret, userId)
	if err != nil {
		return nil, fmt.Errorf("failed to generate jwt token. error: %w", err)
	}

	token := &AuthToken{
		AccessToken: jwt,
	}

	return token, nil
}

func getTokenFromHeader(r *http.Request) (string, error) {
	authHeader := r.Header.Get("authorization")
	parts := strings.Split(authHeader, " ")

	if len(parts) < 2 {
		return "", fmt.Errorf("invalid authorization header")
	}

	return parts[1], nil
}

func parseJwt(secret []byte, jwtToken string) (int64, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("failed to validate signing method")
		}

		return secret, nil
	})

	if err != nil {
		return 0, fmt.Errorf("failed parsing jwt. error: %w", err)
	}

	if !token.Valid {
		return 0, fmt.Errorf("invalid token")
	}

	tokenSub, err := token.Claims.GetSubject()

	if err != nil {
		return 0, fmt.Errorf("failed to fetch token sub")
	}

	tokenUserId, err := strconv.ParseInt(tokenSub, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("failed to parse user id from token")
	}

	return tokenUserId, nil
}
