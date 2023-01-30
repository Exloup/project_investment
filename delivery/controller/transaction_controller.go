package controller

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct Controller
type TransactionProductController struct {
	router  *gin.Engine
	usecase usecase.TransactionProductUseCase
}

// Function Usecase Insert / Create
func (tpc *TransactionProductController) CreateTransactionProduct(ctx *gin.Context) {
	var product model.TransactionProducts
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	p, err := tpc.usecase.CreateTransactionProduct(&product)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, p)
}

// Function Usecase Get All
func (tpc *TransactionProductController) GetAllTransactionProduct(ctx *gin.Context) {
	product, err := tpc.usecase.GetAllTransactionProduct()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// Function Usecase Get by ID
func (tpc *TransactionProductController) GetTransactionProductById(ctx *gin.Context) {
	id := ctx.Param("id")
	product, err := tpc.usecase.GetTransactionProductById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// Function Usecase Update by ID
func (tpc *TransactionProductController) UpdateTransactionProduct(ctx *gin.Context) {
	var product model.TransactionProducts
	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	p, err := tpc.usecase.UpdateTransactionProduct(&product)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, p)
}

// Function Usecase Delete by ID
func (tpc *TransactionProductController) DeleteTransactionProduct(ctx *gin.Context) {
	var product model.TransactionProducts
	id := ctx.Param("id")
	err := tpc.usecase.DeleteTransactionProduct(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, product)
}

// Store Controller
func NewTransactionProductController(router *gin.Engine, transactionproductUseCase usecase.TransactionProductUseCase) *TransactionProductController {
	newTransactionProductController := TransactionProductController{
		router:  router,
		usecase: transactionproductUseCase,
	}
	router.POST("/trx/insert", newTransactionProductController.CreateTransactionProduct)
	router.GET("/trx/getall", newTransactionProductController.GetAllTransactionProduct)
	router.GET("/trx/getall/:id", newTransactionProductController.GetTransactionProductById)
	router.PUT("/trx/update/:id", newTransactionProductController.UpdateTransactionProduct)
	router.DELETE("/trx/delete/:id", newTransactionProductController.DeleteTransactionProduct)
	return &newTransactionProductController
}
