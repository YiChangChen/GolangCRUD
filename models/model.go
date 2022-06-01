package models

import "github.com/golang-jwt/jwt"

type User struct {
	ID       string `json:"id" db:"Id"`
	Username string `json:"username" db:"Username"`
	Email    string `json:"email" db:"Email"`
	Phone    string `json:"phone" db:"Phone"`
}

// custom claims
type Claims struct {
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}
