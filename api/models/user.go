package model

type User struct {
	UserId   int    `json:"id"`
	UserName string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
