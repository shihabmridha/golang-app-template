package user

import (
	"fmt"

	"github.com/shihabmridha/golang-app-template/pkg/database"
)

type Repository struct {
	db *database.Sql
}

func NewRepo(sql *database.Sql) *Repository {
	return &Repository{db: sql}
}

func (r *Repository) Get() (*[]User, error) {
	users := []User{}
	err := r.db.Select(&users, "SELECT * FROM user")

	if err != nil {
		return nil, fmt.Errorf("UserRepository - Get - r.db.Select: %w", err)
	}

	return &users, nil
}

func (r *Repository) Create(user User) (int64, error) {
	res, err := r.db.Exec(`
        INSERT INTO user (firstName, lastName, username, email, password, birthDate, isActive, activationCode) VALUES (?, ?, ?, ?, ?, ?, false, ?)
    `, user.FirstName, user.LastName, user.Username, user.Email, user.Passowrd, user.BirthDate, user.ActivationCode)

	if err != nil {
		return 0, fmt.Errorf("UserRepository - Create - r.db.Exec: %w", err)
	}

	id, _ := res.LastInsertId()

	return id, nil
}
