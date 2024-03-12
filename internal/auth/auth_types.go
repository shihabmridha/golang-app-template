package auth

type UserLogin struct {
	Email    string `json:"email"`
	Passowrd string `json:"password"`
}

type AuthToken struct {
	AccessToken string `json:"accessToken"`
}
