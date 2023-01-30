package usecase

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/utils/authenticator"
)

// Interface untuk Authentifikator
type AuthUseCase interface {
	UserAuth(user model.UserAccountCredential) (token string, err error)
}

// Struct Authentificator
type authUseCase struct {
	tokenService authenticator.AccessToken
}

// User Authentificator
func (au *authUseCase) UserAuth(user model.UserAccountCredential) (token string, err error) {
	if user.Email == user.Email && user.Password == user.Password {
		token, err := au.tokenService.CreateAccessToken(&user)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		return "", nil
	}
}

func NewAuthUseCase(service authenticator.AccessToken) AuthUseCase {
	verifyAuth := new(authUseCase)
	verifyAuth.tokenService = service
	return verifyAuth
}
