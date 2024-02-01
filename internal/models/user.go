package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

type UserCreate struct {
	Username string `json:"username"`
	Age      *int   `json:"age"`
	Email    string `json:"email"`
}
