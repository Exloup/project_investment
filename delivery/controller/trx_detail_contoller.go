package controller

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct Controller
type TransactionDetailController struct {
	router  *gin.Engine
	usecase usecase.TransactionDetailUseCase
}

// Function Usecase Insert / Create
func (tdc *TransactionDetailController) CreateTransactionDetail(ctx *gin.Context) {
	var products model.TransactionDetails
	err := ctx.ShouldBindJSON(&products)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	p, err := tdc.usecase.CreateTransactionDetail(&products)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, p)
}

// Function Usecase Get All
func (tdc *TransactionDetailController) GetAllTransactionDetail(ctx *gin.Context) {
	products, err := tdc.usecase.GetAllTransactionDetail()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// Function Usecase Get by ID
func (tdc *TransactionDetailController) GetTransactionDetailById(ctx *gin.Context) {
	id := ctx.Param("id")
	products, err := tdc.usecase.GetTransactionDetailById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// Function Usecase Update by ID
func (tdc *TransactionDetailController) UpdateTransactionDetail(ctx *gin.Context) {
	var products model.TransactionDetails
	err := ctx.ShouldBindJSON(&products)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	p, err := tdc.usecase.UpdateTransactionDetail(&products)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, p)
}

// Function Usecase Delete by ID
func (tdc *TransactionDetailController) DeleteTransactionDetail(ctx *gin.Context) {
	var products model.TransactionDetails
	id := ctx.Param("id")
	err := tdc.usecase.DeleteTransactionDetail(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, products)
}

// Store Controller
func NewTransactionDetailController(router *gin.Engine, transactiondetailUseCase usecase.TransactionDetailUseCase) *TransactionDetailController {
	newTransactionDetailController := TransactionDetailController{
		router:  router,
		usecase: transactiondetailUseCase,
	}
	router.POST("/trx-detail/insert", newTransactionDetailController.CreateTransactionDetail)
	router.GET("/trx-detail/getall", newTransactionDetailController.GetAllTransactionDetail)
	router.GET("/trx-detail/getall/:id", newTransactionDetailController.GetTransactionDetailById)
	router.PUT("/trx-detail/update/:id", newTransactionDetailController.UpdateTransactionDetail)
	router.DELETE("/trx-detail/delete/:id", newTransactionDetailController.DeleteTransactionDetail)
	return &newTransactionDetailController
}
