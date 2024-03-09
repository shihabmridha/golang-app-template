package user

import "time"

type User struct {
	Id             int        `db:"id" json:"id"`
	FirstName      string     `db:"firstName" json:"firstName"`
	LastName       string     `db:"lastName" json:"lastName"`
	Username       string     `db:"username" json:"username"`
	Email          string     `db:"email" json:"email"`
	Passowrd       string     `db:"password"`
	IsActive       bool       `db:"isActive" json:"isActive"`
	ActivationCode string     `db:"activationCode"`
	BirthDate      *time.Time `db:"birthDate" json:"birthDate"`
	CreatedAt      *time.Time `db:"createdAt" json:"createdAt"`
	UpdatedAt      *time.Time `db:"updatedAt" json:"updatedAt"`
}
