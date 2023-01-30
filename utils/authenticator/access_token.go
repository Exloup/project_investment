package authenticator

import (
	"final_project_coins_investment/config"
	"final_project_coins_investment/model"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Interface Token
type AccessToken interface {
	CreateAccessToken(uc *model.UserAccountCredential) (string, error)
	VerifyAccessToken(tokenString string) (jwt.MapClaims, error)
}

type accessToken struct {
	conf config.TokenConfig
}

// Create Access Token
func (at *accessToken) CreateAccessToken(ucd *model.UserAccountCredential) (string, error) {
	start := time.Now().Local()
	end := start.Add(at.conf.AccessTokenLifeTime)
	claims := TokenClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer: at.conf.ApplicationName,
		},
		Email: ucd.Email,
		Name:  ucd.Name,
	}
	claims.IssuedAt = start.Unix()
	claims.ExpiresAt = end.Unix()
	token := jwt.NewWithClaims(
		at.conf.JwtSigningMethod, claims)
	return token.SignedString([]byte(at.conf.JwtSignatureKey))
}

// Create Verify Token
func (at *accessToken) VerifyAccessToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("signing method invalid")
		} else if method != at.conf.JwtSigningMethod {
			return nil, fmt.Errorf("signing method invalid")
		}
		return []byte(at.conf.JwtSignatureKey), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	return claims, nil
}

// Store Repo Token
func NewAccessToken(config config.TokenConfig) AccessToken {
	return &accessToken{
		conf: config,
	}
}
