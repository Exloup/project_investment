package authenticator

import "github.com/golang-jwt/jwt/v4"

type TokenClaims struct {
	jwt.StandardClaims
	Email string `json:"Email"`
	Name string `json:"Name"`
}