package models

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserCreate struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}
