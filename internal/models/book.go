package models

type Book struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BookCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
