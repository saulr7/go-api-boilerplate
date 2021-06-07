package models

import "github.com/dgrijalva/jwt-go"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claim struct {
	jwt.StandardClaims
	Email string `json:"email"`
}
