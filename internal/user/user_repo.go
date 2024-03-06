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
		return nil, fmt.Errorf("UserRepository - Get - r.Sql.Select: %w", err)
	}

	return &users, nil
}
