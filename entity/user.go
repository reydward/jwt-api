package entity

type User struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
}
