package controller

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerAccountController struct {
	router  *gin.Engine
	usecase usecase.AccountUseCase
}

// Function Get All Account
func (ac *CustomerAccountController) GetAllAccount(ctx *gin.Context) {
	accounts, err := ac.usecase.GetAllAccount()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

// Function Get ID Account
func (ac *CustomerAccountController) GetAccountById(ctx *gin.Context) {
	id := ctx.Param("id")
	accounts, err := ac.usecase.GetAccountById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

// Function Registration Account
func (ac *CustomerAccountController) RegisterAccount(ctx *gin.Context) {
	var accounts model.CustomerAccounts
	err := ctx.ShouldBindJSON(&accounts)
	if err != nil {
		log.Println("err", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	a, err := ac.usecase.RegisterAccount(&accounts)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, a)
}

// Funtion Update Account
func (ac *CustomerAccountController) UpdateAccount(ctx *gin.Context) {
	var accounts model.CustomerAccounts
	err := ctx.ShouldBindJSON(&accounts)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	a, err := ac.usecase.UpdateAccount(&accounts)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, a)
}

// Function Activated by ID
func (ac *CustomerAccountController) ActivatedAccount(ctx *gin.Context) {
	var accounts model.CustomerAccounts
	err := ctx.ShouldBindJSON(&accounts)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	a, err := ac.usecase.ActivatedAccount(&accounts)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, a)
}

// Function Deactivated by ID
func (ac *CustomerAccountController) DeactivatedAccount(ctx *gin.Context) {
	var accounts model.CustomerAccounts
	err := ctx.ShouldBindJSON(&accounts)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	a, err := ac.usecase.DeactivatedAccount(&accounts)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, a)
}

// Fucntion Delete Account
func (ac *CustomerAccountController) DeleteAccount(ctx *gin.Context) {
	var accounts model.CustomerAccounts
	id := ctx.Param("id")
	err := ac.usecase.DeleteAccount(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

// Store Account Controller
func NewAccountController(router *gin.Engine, accountUseCase usecase.AccountUseCase) *CustomerAccountController {
	newAccountController := CustomerAccountController{
		router:  router,
		usecase: accountUseCase,
	}
	router.POST("/account/register", newAccountController.RegisterAccount)
	router.GET("/account/getall", newAccountController.GetAllAccount)
	router.GET("/account/getall/:id", newAccountController.GetAccountById)
	router.PUT("/account/update/:id", newAccountController.UpdateAccount)
	router.PUT("/account/activated/:id", newAccountController.ActivatedAccount)
	router.PUT("/account/deactivated/:id", newAccountController.DeactivatedAccount)
	router.DELETE("/account/delete/:id", newAccountController.DeleteAccount)
	return &newAccountController
}
