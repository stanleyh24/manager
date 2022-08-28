package models

type User struct {
	Id       int64  `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Role struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
