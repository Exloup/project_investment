package controller

import (
	"final_project_coins_investment/model"
	"final_project_coins_investment/usecase"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Struct Controller
type PriceController struct {
	router  *gin.Engine
	usecase usecase.PriceUseCase
}

// Function Usecase Insert / Create
func (pr *PriceController) InsertPrice(ctx *gin.Context) {
	var price model.Prices
	err := ctx.ShouldBindJSON(&price)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	p, err := pr.usecase.InsertPrice(&price)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, p)
	fmt.Println(price.Performa)
}

// Function Usecase Get All
func (pr *PriceController) GetAllPrice(ctx *gin.Context) {
	prices, err := pr.usecase.GetAllPrice()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, prices)
}

// Function Usecase Get by ID
func (pr *PriceController) GetPriceById(ctx *gin.Context) {
	id := ctx.Param("id")
	prices, err := pr.usecase.GetPriceById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, prices)
}

// Function Usecase Update by ID
func (pr *PriceController) UpdatePrice(ctx *gin.Context) {
	var price model.Prices
	err := ctx.ShouldBindJSON(&price)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}
	p, err := pr.usecase.UpdatePrice(&price)
	if err != nil {
		panic(err)
	}
	ctx.JSON(http.StatusOK, p)
}

// Function Usecase Delete by ID
func (pr *PriceController) DeletePrice(ctx *gin.Context) {
	var price model.Products
	id := ctx.Param("id")
	err := pr.usecase.DeletePrice(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, price)
}

// Store Controller
func NewPriceController(router *gin.Engine, priceUseCase usecase.PriceUseCase) *PriceController {
	newPriceController := PriceController{
		router:  router,
		usecase: priceUseCase,
	}
	router.POST("/prices/insert", newPriceController.InsertPrice)
	router.GET("/prices/getall", newPriceController.GetAllPrice)
	router.GET("/prices/getall/:id", newPriceController.GetPriceById)
	router.PUT("/prices/update/:id", newPriceController.UpdatePrice)
	router.DELETE("/prices/delete/:id", newPriceController.DeletePrice)
	return &newPriceController
}
