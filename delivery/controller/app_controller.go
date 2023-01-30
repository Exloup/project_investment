package controller

import (
	"final_project_coins_investment/delivery/middleware"
	"final_project_coins_investment/model"
	"final_project_coins_investment/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppController struct {
	rg          *gin.RouterGroup
	authUseCase usecase.AuthUseCase
}

// Func for User Authentification
func (ac *AppController) userAuth(ctx *gin.Context) {
	var user model.UserAccountCredential
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Can't Bind Struct",
		})
		return
	}
	token, err := ac.authUseCase.UserAuth(user)
	if err != nil {
		ctx.AbortWithStatus(401)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (ac *AppController) getAccount(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "User Account",
	})
}

func NewAppController(routerGroup *gin.RouterGroup, authUseCase usecase.AuthUseCase, tokenMiddleware middleware.AuthTokenMiddleware) *AppController {
	controller := AppController{
		rg:          routerGroup,
		authUseCase: authUseCase,
	}
	controller.rg.POST("/auth", controller.userAuth)
	protectedGroup := controller.rg.Group("/protected", tokenMiddleware.RequireTokken())
	protectedGroup.GET("/user", controller.getAccount)
	return &controller
}
