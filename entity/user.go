package entity

type User struct {
	Id       string `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Tel      string `json:"tel"`
}
