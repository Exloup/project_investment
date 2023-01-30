package middleware

import (
	"final_project_coins_investment/utils/authenticator"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
}

type AuthTokenMiddleware interface {
	RequireTokken() gin.HandlerFunc
}

type authTokenMiddleware struct {
	acsToken authenticator.AccessToken
}

// Function for get Token
func (atm *authTokenMiddleware) RequireTokken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h := authHeader{}
		if err := ctx.ShouldBindHeader(&h); err != nil {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		// Make Token
		tokenString := strings.Replace(h.AuthorizationHeader, "Bearer", "", -1)
		fmt.Println(tokenString)
		if tokenString == "" {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
		// make sure token is verify
		token, err := atm.acsToken.VerifyAccessToken(tokenString)
		if err != nil {
			ctx.JSON(200, gin.H{
				"message": "User Account",
			})
			ctx.Abort()
			return
		}
		// If token is empty
		if token != nil {
			ctx.Next()
		} else {
			ctx.JSON(401, gin.H{
				"message": "Unauthorized",
			})
			ctx.Abort()
			return
		}
	}
}

// Store in New token
func NewTokenValidator(acsToken authenticator.AccessToken) AuthTokenMiddleware {
	return &authTokenMiddleware{
		acsToken: acsToken,
	}
}
