package user

import (
	"context"
	"fmt"

	"github.com/shihabmridha/golang-app-template/pkg/database"
	"github.com/shihabmridha/golang-app-template/pkg/logging"
	"go.uber.org/zap"
)

type Repository struct {
	db     *database.Sql
	logger *zap.SugaredLogger
}

func NewRepo(ctx context.Context, sql *database.Sql) Repository {
	return Repository{
		db:     sql,
		logger: logging.FromContext(ctx),
	}
}

func (r *Repository) GetAll() ([]User, error) {
	users := []User{}
	err := r.db.Select(&users, "SELECT * FROM user")

	if err != nil {
		r.logger.Errorf("query failed to fetch users. error: %w", err)
		return nil, fmt.Errorf("query failed to fetch users. error: %w", err)
	}

	return users, nil
}

func (r *Repository) GetByUsername(username string) (*User, error) {
	user := &User{}
	err := r.db.Get(user, "SELECT * FROM user WHERE username=?", username)

	if err != nil {
		r.logger.Errorf("query failed to fetch user (username = %s). error: %w", username, err)
		return nil, fmt.Errorf("query failed to fetch user (username = %s). error: %w", username, err)
	}

	return user, nil
}

func (r *Repository) GetByEmail(email string) (*User, error) {
	user := &User{}
	err := r.db.Get(user, "SELECT * FROM user WHERE email=?", email)

	if err != nil {
		r.logger.Errorf("query failed to fetch user (email = %s). error: %w", email, err)
		return nil, fmt.Errorf("query failed to fetch user (email = %s). error: %w", email, err)
	}

	return user, nil
}

func (r *Repository) GetIdUsingActivationCode(code string) (int64, error) {
	var id int64 = 0
	err := r.db.Get(&id, "SELECT id FROM user WHERE activationCode=?", code)

	if err != nil {
		r.logger.Errorf("query failed to fetch user by activation code. error: %w", err)
		return 0, fmt.Errorf("query failed to fetch user by activation code. error: %w", err)
	}

	if id == 0 {
		return 0, fmt.Errorf("failed to fetch user by activation code")
	}

	return id, nil
}

func (r *Repository) Create(user User) (int64, error) {
	res, err := r.db.Exec(`
        INSERT INTO user (firstName, lastName, username, email, password, birthDate, isActive, activationCode) VALUES (?, ?, ?, ?, ?, ?, false, ?)
    `, user.FirstName, user.LastName, user.Username, user.Email, user.Passowrd, user.BirthDate, user.ActivationCode)

	if err != nil {
		r.logger.Errorf("query failed to create user. error: %w", err)
		return 0, fmt.Errorf("query failed to create user. error: %w", err)
	}

	id, _ := res.LastInsertId()

	return id, nil
}

func (r *Repository) AlreadyExists(username, email string) (bool, error) {
	id := 0
	if err := r.db.Get(&id, "SELECT id FROM user WHERE username=? OR email=?", username, email); err != nil {
		r.logger.Errorf("query failed to fetch existing user. error: %w", err)
		return false, fmt.Errorf("query failed to fetch existing user. error: %w", err)
	}

	return id != 0, nil
}

func (r *Repository) Activate(id int64) error {
	_, err := r.db.Exec(`UPDATE user SET isActive = true WHERE id = ?`, id)

	if err != nil {
		r.logger.Errorf("query failed activate user. error: %w", err)
		return fmt.Errorf("query failed activate user. error: %w", err)
	}

	return nil
}
