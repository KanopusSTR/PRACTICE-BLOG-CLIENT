package models

import "github.com/golang-jwt/jwt/v5"

type Token struct {
	Mail string `json:"mail"`
	jwt.RegisteredClaims
}
