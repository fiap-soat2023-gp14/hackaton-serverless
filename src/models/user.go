package models

type UserForm struct {
	Name     string `json:"username"`
	Registry string `json:"registry"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
