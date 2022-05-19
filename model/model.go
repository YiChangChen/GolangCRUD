package model

type User struct {
	ID       string `json:"id" db:"Id"`
	Username string `json:"username" db:"Username"`
	Email    string `json:"email" db:"Email"`
	Phone    string `json:"phone" db:"Phone"`
}
